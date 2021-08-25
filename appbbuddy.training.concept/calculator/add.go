package calculator

func (u UserInput) AddValues() (float32, error) {
	return float32(u.FirstNumber + u.SecondNumber), nil
}
