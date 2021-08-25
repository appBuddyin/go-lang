package calculator

type operations interface {
	AddValues() (float32, error)
	SubtractValues() (float32, error)
	MultiplyValues() (float32, error)
	DevideValues() (float32, error)
}
