package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	frankfurterAPIBaseURL = "https://api.frankfurter.app/latest"
	apiRequestTimeout     = 5 * time.Second
)

type ConversionResponse struct {
	Amount float64            `json:"amount"`
	Base   string             `json:"base"`
	Date   string             `json:"date"`
	Rates  map[string]float64 `json:"rates"`
}

type UserInput struct {
	Amount       float64
	FromCurrency string
	ToCurrency   string
}

func readStringInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("erro ao ler entrada: %w", err)
	}

	return strings.ToUpper(strings.TrimSpace(input)), nil
}

func getUserInput(reader *bufio.Reader) (UserInput, error) {
	var input UserInput
	var err error

	fmt.Print("Digite o valor a ser convertido: ")

	amountStr, _ := reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)

	input.Amount, err = strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return UserInput{}, fmt.Errorf("valor inválido: %w", err)
	}

	input.FromCurrency, err = readStringInput(
		"Digite a moeda de origem (ex: USD, BRL, EUR): ",
		reader,
	)

	if err != nil {
		return UserInput{}, err
	}

	if input.FromCurrency == "" {
		return UserInput{}, fmt.Errorf("moeda de origem não pode ser vazia")
	}

	input.ToCurrency, err = readStringInput(
		"Digite a moeda de destino (ex: USD, BRL, EUR): ",
		reader,
	)

	if err != nil {
		return UserInput{}, err
	}

	if input.ToCurrency == "" {
		return UserInput{}, fmt.Errorf("moeda de destino não pode ser vazia")
	}

	return input, nil
}

// fetchConversionData busca os dados de conversão da API.
// Esta função agora assume que fromCurrency != toCurrency,
// pois isso é tratado pelo chamador (main).
func fetchConversionData(
	amount float64,
	fromCurrency string,
	toCurrency string,
) (*ConversionResponse, error) {
	apiURL := fmt.Sprintf(
		"%s?amount=%.2f&from=%s&to=%s",
		frankfurterAPIBaseURL,
		amount,
		fromCurrency,
		toCurrency,
	)

	fmt.Printf(
		"\nBuscando cotação de %s para %s...\n",
		fromCurrency,
		toCurrency,
	)

	client := http.Client{
		Timeout: apiRequestTimeout,
	}
	
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf(
			"erro ao fazer a requisição para a API: %w",
			err,
		)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"erro da API: Status %s (%d)",
			resp.Status,
			resp.StatusCode,
		)
	}

	var conversionData ConversionResponse
	if err := json.NewDecoder(resp.Body).Decode(&conversionData); err != nil {
		return nil, fmt.Errorf(
			"erro ao decodificar JSON da API: %w",
			err,
		)
	}

	return &conversionData, nil
}

// displayResult exibe o resultado da conversão.
// Assume que fromCurrency != toCurrency.
func displayResult(
	originalAmount float64,
	fromCurrency string,
	toCurrency string,
	data *ConversionResponse,
) {
	convertedAmount, ok := data.Rates[toCurrency]
	if !ok {
		// Se !ok, significa que a API não retornou a taxa para toCurrency.
		// Isso pode ser devido a uma moeda de destino inválida ou não suportada
		// em relação à moeda base.
		log.Printf(
			"Aviso: A API não retornou a cotação para %s baseada em %s. Verifique se as moedas são válidas e suportadas.",
			toCurrency,
			data.Base, // data.Base deve ser igual a fromCurrency
		)

		fmt.Printf(
			"Não foi possível obter a taxa de conversão de %s para %s.\n",
			fromCurrency,
			toCurrency,
		)
		return
	}

	fmt.Printf(
		"\n%.2f %s equivale a %.2f %s (cotação de %s)\n",
		originalAmount,
		fromCurrency,
		convertedAmount,
		toCurrency,
		data.Date,
	)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	userInput, err := getUserInput(reader)
	if err != nil {
		log.Fatalf("Erro ao obter entrada do usuário: %v", err)
	}

	if userInput.FromCurrency == userInput.ToCurrency {
		fmt.Printf(
			"\n%.2f %s = %.2f %s\n",
			userInput.Amount,
			userInput.FromCurrency,
			userInput.Amount,
			userInput.ToCurrency,
		)
		return // Encerra aqui se as moedas forem iguais
	}

	// A partir daqui, sabemos que userInput.FromCurrency != userInput.ToCurrency
	conversionData, err := fetchConversionData(
		userInput.Amount,
		userInput.FromCurrency,
		userInput.ToCurrency,
	)
	if err != nil {
		log.Fatalf("Erro ao buscar dados de conversão: %v", err)
	}

	displayResult(
		userInput.Amount,
		userInput.FromCurrency,
		userInput.ToCurrency,
		conversionData,
	)
}