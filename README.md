# Conversor de Moedas em Go

Este é um simples aplicativo de linha de comando (CLI) escrito em Go que converte valores entre diferentes moedas utilizando taxas de câmbio em tempo real da API [Frankfurter.app](https://www.frankfurter.app/).

## Funcionalidades

*   Converte um valor de uma moeda de origem para uma moeda de destino.
*   Utiliza taxas de câmbio atualizadas da API Frankfurter.app.
*   Interface de usuário interativa via linha de comando.
*   Tratamento básico de erros e validação de entrada.
*   Código organizado em funções para melhor legibilidade e manutenção.

## Pré-requisitos

*   **Go**: Versão 1.24.3 ou superior instalada. Você pode baixar em [go.dev](https://go.dev/dl/).
*   **Conexão com a Internet**: Necessário para buscar as taxas de câmbio da API.

## Como Usar

1.  **Clone o repositório (ou baixe o arquivo `main.go`):**
    ```bash
    # Se você tiver o git e o projeto estiver em um repositório:
    # git clone https://github.com/seu_usuario/seu_repositorio.git
    # cd seu_repositorio

    # Ou simplesmente salve o código fornecido como main.go em um diretório.
    ```

2.  **Navegue até o diretório do projeto:**
    ```bash
    cd /caminho/para/o/projeto
    ```

3.  **Execute o programa:**
     ```bash
     go run main.go
     ```

4.  **Siga as instruções no terminal:**
    O programa solicitará que você insira:
    *   O valor a ser convertido.
    *   A moeda de origem (ex: USD, BRL, EUR).
    *   A moeda de destino (ex: USD, BRL, EUR).

## Exemplo de Uso

```bash
$ go run conversor.go
Digite o valor a ser convertido: 100
Digite a moeda de origem (ex: USD, BRL, EUR): USD
Digite a moeda de destino (ex: USD, BRL, EUR): BRL

Buscando cotação de USD para BRL...

100.00 USD equivale a 565.65 BRL (cotação de 2025-05-09)
```
