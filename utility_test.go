package main

import "testing"

func TestIsValidType(t *testing.T) {

	tests := []struct {
		inputType         string
		expectedValidness bool
	}{
		{"SERIAL", true},
		{"CONCURRENT", true},
		{"", false},
		{"XYZ", false},
		{"serial", false},
		{"SERIALZ", false},
	}

	for _, test := range tests {
		if isValid := isValidType(test.inputType); isValid != test.expectedValidness {
			t.Error("Input Type:", test.inputType, ", Expected Validness:", test.expectedValidness, ", Validness Obtained:", isValid)
		}
	}
}

func TestIsURLsEmpty(t *testing.T) {

	tests := []struct {
		inputURLList      []string
		expectedEmptiness bool
	}{
		{[]string{}, true},
		{[]string{"--/-----/123", "--/---/---/--/--/xyz"}, false},
		{[]string{"/a.com", "/b.com", "c.com"}, false},
	}

	for _, test := range tests {
		if emptiness := isURLsEmpty(test.inputURLList); emptiness != test.expectedEmptiness {
			t.Error("Input URL List:", test.inputURLList, ", Expected Emptiness:", test.expectedEmptiness, ", Emptiness Obtained:", emptiness)
		}
	}
}
