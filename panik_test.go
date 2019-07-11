package panik

import (
	"errors"
	"testing"
)

func TestOnError(t *testing.T) {
	output := ""

	fnPanic = func(v interface{}) {
		output = v.(error).Error()
	}

	OnError(nil)

	if output != "" {
		t.Errorf("Expected nothing; got %s", output)
	}

	OnError(errors.New("Oops"))

	if output == "" {
		t.Error("Expected error; got nothing")
	}
}
