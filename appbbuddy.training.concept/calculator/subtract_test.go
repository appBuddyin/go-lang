package calculator

import "testing"

// table test
var testsForSubstring = []struct {
	name     string
	number1 float32
	number2  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 90.0, false},
	{"valid-data", 500.0, 10.0, 490.0, false},
	{"valid-data", 12.0, 1.0, 11.0, false},
	{"valid-data", 40.0, 5.0, 35.0, false},
	{"valid-data", 80.0, 10.0, 70.0, false},
}

func TestSubstraction(t *testing.T) {
	for _, tt := range testsForSubstring {
		var u UserInput
		u.FirstNumber=tt.number1
		u.SecondNumber=tt.number2
		got, err := u.SubtractValues()
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
