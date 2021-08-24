package multiply

import "testing"

// table test
var tests = []struct {
	name     string
	number1  float32
	number2  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 1000.0, false},
	{"valid-data", 500.0, 10.0, 5000.0, false},
	{"valid-data", 12.0, 1.0, 12.0, false},
	{"valid-data", 40.0, 5.0, 200.0, false},
	{"valid-data", 80.0, 10.0, 800.0, false},
}

func TestAdd(t *testing.T) {
	for _, tt := range tests {
		got, err := MultiplyValues(tt.number1, tt.number2)
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
