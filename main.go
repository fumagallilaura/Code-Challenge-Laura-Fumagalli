package main

import (
    "fmt"
    "os"
    "strings"
)

var morseCodeMap = map[string]string{
    ".-":    "A",
    "-...":  "B",
    "-.-.":  "C",
    "-..":   "D",
    ".":     "E",
    "..-.":  "F",
    "--.":   "G",
    "....":  "H",
    "..":    "I",
    ".---":  "J",
    "-.-":   "K",
    ".-..":  "L",
    "--":    "M",
    "-.":    "N",
    "---":   "O",
    ".--.":  "P",
    "--.-":  "Q",
    ".-.":   "R",
    "...":   "S",
    "-":     "T",
    "..-":   "U",
    "...-":  "V",
    ".--":   "W",
    "-..-":  "X",
    "-.--":  "Y",
    "--..":  "Z",
    "-----": "0",
    ".----": "1",
    "..---": "2",
    "...--": "3",
    "....-": "4",
    ".....": "5",
    "-....": "6",
    "--...": "7",
    "---..": "8",
    "----.": "9",
}

var textToMorseMap = map[string]string{}

func init() {
    for k, v := range morseCodeMap {
        textToMorseMap[v] = k
    }
}

func decodeMorse(morseCode string) string {
    words := strings.Split(morseCode, "   ")
    decodedMessage := ""

    for _, word := range words {
        letters := strings.Split(word, " ")
        for _, letter := range letters {
            if decodedLetter, ok := morseCodeMap[letter]; ok {
                decodedMessage += decodedLetter
            }
        }
        decodedMessage += " "
    }

    return strings.TrimSpace(decodedMessage)
}

func encodeToMorse(text string) string {
    text = strings.ToUpper(text)
    words := strings.Split(text, " ")
    morseMessage := ""

    for _, word := range words {
        for _, letter := range word {
            if morseLetter, ok := textToMorseMap[string(letter)]; ok {
                morseMessage += morseLetter + " "
            }
        }
        morseMessage += "  "
    }

    return strings.TrimSpace(morseMessage)
}

func bitsToMorse(inputBits string) string {
    inputBits = strings.Trim(inputBits, "0")
    words := strings.Split(inputBits, "0000000")
    morseMessage := ""

    for _, word := range words {
        letters := strings.Split(word, "000")
        for _,letter := range letters {
            codes := strings.Split(letter, "0")
            for _, code := range codes {
                if len(code) > 2 {
                    morseMessage += "-"
                } else {
                    morseMessage += "."
                }
            }
            morseMessage += " "
        }
        morseMessage += "  "
    }
    
    return strings.TrimSpace(morseMessage)
}

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: go run main.go <encode/decode> <ARG>")
        return
    }

    mode := os.Args[1]
    message := os.Args[2]

    switch mode {
    case "encodeToMorse":
        fmt.Println(encodeToMorse(message))
    case "decodeMorse":
        fmt.Println(decodeMorse(message))
    case "bitsToMorse":
        fmt.Println(bitsToMorse(message))
    default:
        fmt.Println("Usage: go run main.go <encodeToMorse/decodeMorse/bitsToMorse> <ARG>")
    }
}
