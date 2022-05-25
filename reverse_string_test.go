package reverse_string

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	input := "hello world"
	expected := "dlrow olleh"

	if result := ReverseString(input); result != expected {
		t.Errorf("ReverseString() = %v, expected %v", result, expected)
	}
}
