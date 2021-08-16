package main

import "fmt"

func main() {
	var lenOfNumber int = 100
	var allNumber []string
	incrementForFizz := 3
	incrementForBuzz := 5
	allNumber = initializeAllNumber(allNumber, lenOfNumber)
	allNumber = addFizzBuzz(allNumber, " Fizz", incrementForFizz)
	allNumber = addFizzBuzz(allNumber, " Buzz", incrementForBuzz)
	printSlice(allNumber)
}

func initializeAllNumber(sliceToInitialize []string, numberToAdd int) []string {
	for i := 0; i < numberToAdd; i++ {
		sliceToInitialize = append(sliceToInitialize, " ")
		println("initialised", i, sliceToInitialize[i])
	}
	return sliceToInitialize
}

func printSlice(sliceToPrint []string) {

	for i := 0; i < len(sliceToPrint); i++ {
		fmt.Println("value at ", i, "=", sliceToPrint[i])
	}
}

func addFizzBuzz(sliceToFizzBuzz []string, termToAdd string, incrementAmmount int) []string {

	for lastIndex := incrementAmmount; lastIndex < len(sliceToFizzBuzz); lastIndex = lastIndex + incrementAmmount {

		sliceToFizzBuzz[lastIndex] = sliceToFizzBuzz[lastIndex] + termToAdd
	}
	return sliceToFizzBuzz
}
