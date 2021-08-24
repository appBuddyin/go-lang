package calculator

import (
	"errors"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

const runAtAddress = "localhost:8080"

type Operation interface {
	Says() string
	NumberOfLegs() int
}
type QueryOptions struct {
	FirstNumber  int    `json:"firstNumber"`
	SecondNumber int    `json:"secondNumber"`
	Operator     string `json:"operator"`
}
type AnswerFormet struct {
	Answer int   `json:"answer"`
	Errors error `json:"error"`
}

func AddValues(x, y int) (int, error) {
	return int(x + y), nil
}

func SubtractValues(x, y int) (int, error) {
	return int(x - y), nil
}

func DevideValues(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("error cant devide by zero")
	}
	return int(x / y), nil
}

func MultiplyValues(x, y int) (int, error) {
	return int(x * y), nil
}

func postQuery(c *gin.Context) {
	var newQuery QueryOptions

	// Call BindJSON to bind the received JSON to newQuery
	if err := c.BindJSON(&newQuery); err != nil {
		return
	}

	var answer AnswerFormet
	switch newQuery.Operator {
	case "AddValues":
		answer.Answer, answer.Errors = AddValues(newQuery.FirstNumber, newQuery.SecondNumber)
	case "SubtractValues":
		answer.Answer, answer.Errors = SubtractValues(newQuery.FirstNumber, newQuery.SecondNumber)
	case "DevideValues":
		answer.Answer, answer.Errors = DevideValues(newQuery.FirstNumber, newQuery.SecondNumber)
	case "MultiplyValues":
		answer.Answer, answer.Errors = MultiplyValues(newQuery.FirstNumber, newQuery.SecondNumber)
	default:
		answer.Answer, answer.Errors = 0, errors.New("did not mach any of the operator")
	}

	c.IndentedJSON(http.StatusCreated, answer)
}

func CalculateApi() {
	router := gin.Default()
	router.POST("/", postQuery)

	log.Println("starting application on adress - ", runAtAddress)
	router.Run(runAtAddress)
}
