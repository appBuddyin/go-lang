package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	for _, tt := range testsForAdd {
		var u UserInput
		u.FirstNumber = tt.number1
		u.SecondNumber = tt.number2
		got, err := u.AddValues()
		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error, but got one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, got)
		}
	}
}
