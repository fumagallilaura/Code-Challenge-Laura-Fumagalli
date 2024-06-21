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

func TestBitsToMorse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"10101010001000111010111011100000001011101110111000101011100011101010001",
			".... . -.--   .--- ..- -.. .",
		},
	}

	for _, test := range tests {
		result := bitsToMorse(test.input)
		if result != test.expected {
			t.Errorf("bitsToMorse(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}