package calculator

func (u UserInput) SubtractValues() (float32, error) {
	return float32(u.FirstNumber - u.SecondNumber), nil
}
