package devide

import "errors"

func DevideValues(x, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("error cant devide by zero")
	}
	return float32(x / y), nil
}
