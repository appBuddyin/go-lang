package calculator

func (u UserInput)MultiplyValues() (float32, error) {
	return float32(u.FirstNumber * u.SecondNumber), nil
}
