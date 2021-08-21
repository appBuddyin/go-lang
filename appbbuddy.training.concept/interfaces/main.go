package main

import "fmt"

// Animal defines the interface for type Animal
type Animal interface {
	Says() string
	NumberOfLegs() int
}

// Dog defines the dog type
type Dog struct {
	Name  string
	Breed string
}

// Gorilla defines the Gorilla type
type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shepherd",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Jock",
		Color:         "grey",
		NumberOfTeeth: 38,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

// Says has a receiver of type *Dog, so it satisfies part of the interface requirements for Animal
// for the Dog type
func (d *Dog) Says() string {
	return "Woof"
}

// NumberOfLegs satisfies the rest of the Animal interface requirements for the Dog type
func (d *Dog) NumberOfLegs() int {
	return 4
}

// Says has a receiver of type *Gorilla, so it satisfies part of the interface requirements for Animal
// for the Gorilla type
func (d *Gorilla) Says() string {
	return "Ugh"
}

// NumberOfLegs satisfies the rest of the Animal interface requirements for the Gorilla type
func (d *Gorilla) NumberOfLegs() int {
	return 2
}
