package main

import "testing"

func TestDecodeMorse(t *testing.T) {
    tests := []struct {
        input  string
        output string
    }{
        {".... . -.--   .--- ..- -.. .", "HEY JUDE"},
        {".... . .-.. .-.. ---   .-- --- .-. .-.. -..", "HELLO WORLD"},
        {"--. ---   --. --- .-.. .- -. --.", "GO GOLANG"},
    }

    for _, test := range tests {
        result := decodeMorse(test.input)
        if result != test.output {
            t.Errorf("For input '%s', expected '%s', but got '%s'", test.input, test.output, result)
        }
    }
}