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
