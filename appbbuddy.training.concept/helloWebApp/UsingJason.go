package helloWebApp

import (
	"encoding/json"
	"log"
)

type QueryOptions struct {
	firstNumber int 'json:"firstNumber"'
	secondNumber int 'json:"secondNumber"'
	operator string 'json:"operator"'
}

func UsingJason(){
	testJson := '[
		{
			"firstNumber": 3,
			"secondNumber": 2,
			"operator": "Add"
		}
		{
			"firstNumber": 3,
			"secondNumber": 4,
			"operator": "Subtract"
		}
		{
			"firstNumber": 7,
			"secondNumber": 4,
			"operator": "Multiply"
		}
]'


	var unmarshaled []QueryOptions

	err := json.Unmarshal([]byte(testJson),&unmarshaled)
	if err != nil {
		log.println("Error in Unmarshalling jason",err)
	}

	log.println("unmarshalled:",unmarshaled)
}