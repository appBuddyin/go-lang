package calculator

import "testing"

func TestDivision(t *testing.T) {
	for _, tt := range testsForDevide {
		var u UserInput
		u.FirstNumber = tt.dividend
		u.SecondNumber = tt.divisor
		got, err := u.DevideValues()
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
			t.Errorf("Expected %f %T but got %f %T", tt.expected, tt.expected, got, got)
		} 
	}
}
