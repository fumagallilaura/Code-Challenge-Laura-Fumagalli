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

func TestEncodeToMorse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"HEY JUDE", ".... . -.--   .--- ..- -.. ."},
		{"A B C", ".-   -...   -.-."},
		{"MORSE CODE", "-- --- .-. ... .   -.-. --- -.. ."},
		{"HELLO", ".... . .-.. .-.. ---"},
	}

	for _, test := range tests {
		result := encodeToMorse(test.input)
		if result != test.expected {
			t.Errorf("encodeToMorse(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}