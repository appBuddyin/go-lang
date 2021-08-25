package calculator

import "errors"

func (u UserInput)DevideValues() (float32, error) {
	if u.SecondNumber == 0 {
		return 0, errors.New("error cant devide by zero")
	}
	return float32(u.FirstNumber / u.SecondNumber), nil
}
