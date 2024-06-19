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

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: go run main.go <decode> <ARG>")
        return
    }

    mode := os.Args[1]
    message := os.Args[2]

    switch mode {
    case "decode":
        fmt.Println(decodeMorse(message))
    default:
        fmt.Println("Usage: go run main.go <decode> <ARG>")
    }
}
