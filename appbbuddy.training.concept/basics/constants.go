package basics

import "fmt"

const Pi = 3.14

func constants() {
	const World = "language world"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
