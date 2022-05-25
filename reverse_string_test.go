package reverse_string

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	input := "he💖llo world"
	expected := "dlrow oll💖eh"

	if result := ReverseString(input); result != expected {
		t.Errorf("ReverseString() = %v, expected %v", result, expected)
	}
}
