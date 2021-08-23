package helloWebApp

import (
	"encoding/json"
	"log"
	"fmt"
)

type QueryOptions struct {
	FirstNumber  int    `json:"firstNumber"`
	SecondNumber int    `json:"secondNumber"`
	Operator     string `json:"operator"`
}
type AnswerFormet struct {
	Answer  int    `json:"answer"`
	Errors error `json:"error"`
}

func UsingJason() {
	testJson := `
	[
		{
			"firstNumber": 3,
			"secondNumber": 2,
			"operator": "Add"
		},
		{
			"firstNumber": 3,
			"secondNumber": 4,
			"operator": "Subtract"
		},
		{
			"firstNumber": 7,
			"secondNumber": 4,
			"operator": "Multiply"
		}
	]`

	var unmarshaled []QueryOptions

	err := json.Unmarshal([]byte(testJson), &unmarshaled)
	if err != nil {
		log.Println("Error in Unmarshalling jason", err)
	}

	log.Println("unmarshalled:", unmarshaled)

	// write json from a struct
	var mySlice []AnswerFormet

	var m1 AnswerFormet
	m1.Answer = 7
	m1.Errors = nil

	mySlice = append(mySlice, m1)

	var m2 AnswerFormet
	m1.Answer = 7
	m1.Errors = nil

	mySlice = append(mySlice, m2)

	// marshel indent makes it easier to read
	//second arg prefix to json string
	//third spcing to indent
	newJson, err := json.MarshalIndent(mySlice, "", "     ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))

}
