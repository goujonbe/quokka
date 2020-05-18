package password

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	got, _ := Generate(12, true, true, true, true)
	expected := 12
	if len(got) != expected {
		t.Errorf("Expected string of %q characters got %q", expected, len(got))
	}
}

func TestBuildCharacterSet(t *testing.T) {
	testCases := []struct {
		includeLowerLetters bool
		includeUpperLetters bool
		includeDigits       bool
		includeSymbols      bool
		expected            string
	}{
		{false, false, false, false, ""},
		{false, false, false, true, Symbols},
		{false, false, true, false, Digits},
		{false, false, true, true, Digits + Symbols},
		{false, true, false, false, UpperLetters},
		{false, true, false, true, UpperLetters + Symbols},
		{false, true, true, false, UpperLetters + Digits},
		{false, true, true, true, UpperLetters + Digits + Symbols},
		{true, false, false, false, LowerLetters},
		{true, false, false, true, LowerLetters + Symbols},
		{true, false, true, false, LowerLetters + Digits},
		{true, false, true, true, LowerLetters + Digits + Symbols},
		{true, true, false, false, LowerLetters + UpperLetters},
		{true, true, false, true, LowerLetters + UpperLetters + Symbols},
		{true, true, true, false, LowerLetters + UpperLetters + Digits},
		{true, true, true, true, LowerLetters + UpperLetters + Digits + Symbols},
	}

	for _, tt := range testCases {
		got := buildCharacterSet(tt.includeLowerLetters, tt.includeUpperLetters, tt.includeDigits, tt.includeSymbols)
		if got != tt.expected {
			t.Errorf("got %q want %q", got, tt.expected)
		}
	}
}
