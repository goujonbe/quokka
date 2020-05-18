package password

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	t.Run("password length > 0 but empty set", func(t *testing.T) {
		_, err := Generate(12, false, false, false, false)
		if err != ErrorEmptyCharacterSet {
			t.Errorf("got %q want %q", err, ErrorEmptyCharacterSet)
		}
	})
	t.Run("password length = 0", func(t *testing.T) {
		got, _ := Generate(0, true, false, false, false)
		expected := ""
		if got != expected {
			t.Errorf("got %q want %q", got, expected)
		}
	})
	t.Run("password length > 0", func(t *testing.T) {
		got, _ := Generate(10, true, false, false, false)
		expected := 10
		if len(got) != expected {
			t.Errorf("got %q want %q", len(got), expected)
		}
	})
}

func TestBuildCharacterSet(t *testing.T) {
	testCases := []struct {
		name                string
		includeLowerLetters bool
		includeUpperLetters bool
		includeDigits       bool
		includeSymbols      bool
		expected            string
	}{
		{"empty set", false, false, false, false, ""},
		{"set of symbols only", false, false, false, true, Symbols},
		{"set of digits only", false, false, true, false, Digits},
		{"set of digits and symbols", false, false, true, true, Digits + Symbols},
		{"set of upper letters", false, true, false, false, UpperLetters},
		{"set of upper letters and symbols", false, true, false, true, UpperLetters + Symbols},
		{"set of upper letters and digits", false, true, true, false, UpperLetters + Digits},
		{"set of upper letters and digits and symbols", false, true, true, true, UpperLetters + Digits + Symbols},
		{"set of lower letters ", true, false, false, false, LowerLetters},
		{"set of lower letters and symbols", true, false, false, true, LowerLetters + Symbols},
		{"set of lower letters and digits", true, false, true, false, LowerLetters + Digits},
		{"set of lower letters and digits and symbols", true, false, true, true, LowerLetters + Digits + Symbols},
		{"set of lower letters and upper letters", true, true, false, false, LowerLetters + UpperLetters},
		{"set of lower letters and upper letters and symbols", true, true, false, true, LowerLetters + UpperLetters + Symbols},
		{"set of lower letters and upper letters and digits", true, true, true, false, LowerLetters + UpperLetters + Digits},
		{"set of lower letters and upper letters and symbols and digits", true, true, true, true, LowerLetters + UpperLetters + Digits + Symbols},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := buildCharacterSet(tt.includeLowerLetters, tt.includeUpperLetters, tt.includeDigits, tt.includeSymbols)
			if got != tt.expected {
				t.Errorf("got %q want %q", got, tt.expected)
			}
		})
	}
}
