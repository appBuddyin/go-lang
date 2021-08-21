package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Welcome to the playground!")

	// will be usefull in improving effecincy
	fmt.Println("The time is", time.Now())

	rand.Seed(int64(time.Now().UTC().Nanosecond()))
	fmt.Println("a random generated number=", rand.Intn(100))

}
