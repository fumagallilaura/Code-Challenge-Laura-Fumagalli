# Morse Code Decoder

Este projeto implementa um codificador/decodificador de código Morse em Go.

## Descrição Técnica

A solução implementada consiste em um mapa que associa cada sequência de código Morse a sua letra ou número correspondente. A função `decodeMorse` é responsável por decodificar uma string de código Morse em uma mensagem legível. Já a função `encodeToMorse` é responsável por codificar uma uma mensagem em uma string de código Morse. 

## Instruções para Rodar o Programa

> Para rodar este projeto é necessário ter o Golang instalado. Para isso, verifique a [documentação oficial](https://go.dev/doc/install).

1. Clone o repositório:
   ```sh
   git clone git@github.com:fumagallilaura/Code-Challenge-Laura-Fumagalli.git
   cd Code-Challenge-Laura-Fumagalli/
   ```
2. Gere as dependências:
    ```
    go mod init Code-Challenge-Laura-Fumagalli
    ```
3. Compile e execute o programa:
    - Para decodificar:
        ```sh
        go run main.go decode "-..- .- -... .-.. .- ..-"
        ```
        Você deve ver a saída:
        ```
        XABLAU
        ```
    - Para codificar:
        ```sh
        go run main.go encode XABLAU
        ```
        Você deve ver a saída:
        ```
        -..- .- -... .-.. .- ..-
        ```

## Instruções para Rodar os testes

Rode o código

```
go test
```
