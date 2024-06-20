# Morse Code Decoder

Este projeto implementa um codificador/decodificador de código Morse em Go.

<hr>

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

## Descrição Técnica

A solução implementada utiliza a linguagem de programação Go para criar um codificador/decodificador de código Morse. O código é estruturado de forma modular para facilitar a compreensão e manutenção.

- Mapa de Código Morse: Foi criado um mapa(ou dicionário) e ele é utilizado para associar cada sequência de código Morse a sua letra ou número correspondente. Esse mapa é definido como um map[string]string onde a chave é a sequência de código Morse e o valor é a letra ou número correspondente.

- textToMorseMap: mapeia cada letra ou número do alfabeto para sua respectiva sequência de código Morse(o inverso da variável anterior). Isso é feito no momento da inicialização do programa, graças à função init(). Essa inicialização permite que tenhamos um mapeamento bidirecional entre letras/números e código Morse. Enquanto morseCodeMap é usado para decodificar código Morse em texto legível, textToMorseMap é usado para codificar texto legível em código Morse. Essa abordagem torna o código mais flexível e eficiente, pois evita duplicação de lógica e facilita a manutenção do programa.

- Função decodeMorse: Esta função recebe uma string de código Morse como entrada e realiza a decodificação para uma mensagem. Ela divide a string de código Morse em palavras separadas por três espaços (" ") e em seguida divide cada palavra em letras separadas por um espaço (" "). Para cada letra, ela verifica no mapa de código Morse e adiciona a letra correspondente à mensagem decodificada.

- Função encodeToMorse: Esta função recebe uma mensagem como entrada e realiza a codificação para uma string de código Morse. Ela converte a mensagem para maiúsculas, divide a mensagem em palavras e em seguida itera sobre cada letra da palavra. Para cada letra, ela verifica no mapa de código Morse e adiciona a sequência de código Morse correspondente à mensagem codificada.

Ambas as funções são projetadas para lidar com espaços extras antes ou depois dos códigos de código Morse, ignorando-os durante o processo de codificação ou decodificação.

Essa abordagem permite que o programa codifique e decodifique mensagens em código Morse de forma precisa e rápida, oferecendo uma solução prática para a conversão entre texto legível e código Morse. Como adicional, permite que a pessoa usuária escolha via CLI se gostaria de codificar ou decodificar a menságem.
