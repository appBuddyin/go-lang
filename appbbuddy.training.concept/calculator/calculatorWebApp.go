package calculator

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserInput struct {
	FirstNumber  float32
	SecondNumber float32
	Operator     string
}

func CalculatorWebApp() {
	var num1 float32
	var num2 float32
	var operator string
	var err error
	var message string
	var u1 UserInput

	router := gin.Default()
	router.GET("/*path", func(c *gin.Context) {
		path := c.Param("path")
		parameter := strings.Split(path, "/")

		log.Printf("parameter type = %T ", parameter)
		log.Println("parameter length = ", len(parameter))
		log.Println("parameter  = ", parameter[:])

		num1, num2, operator, err = valuesFromString(parameter)
		if err == nil {
			u1.FirstNumber = num1
			u1.SecondNumber = num2
			u1.Operator = operator
			answer, err := calculate(u1)
			if err == nil {
				message = fmt.Sprintf("%f %s %f= %f", num1, operator, num2, answer)
			} else {
				message = fmt.Sprintf("an error occurred %e", err)
			}
		} else {
			message = fmt.Sprint("an error occurred -", err)
		}
		c.String(http.StatusOK, message)
	})

	router.Run(":8080")
}

func valuesFromString(parameter []string) (num1, num2 float32, operator string, err error) {
	switch {
	case len(parameter) == 4:

		n1, err := strconv.ParseFloat(parameter[1], 32)
		if err == nil {
			num1 = float32(n1)
			log.Printf("Type: %T \n", num1)
			log.Println("Value:", num1)
		} else {
			log.Println("error in first number conversion:", err)
			return num1, num2, operator, err
		}

		n2, err := strconv.ParseFloat(parameter[3], 64)
		if err == nil {
			num2 = float32(n2)
			log.Printf("Type: %T \n", num2)
			log.Println("Value:", num2)
		} else {
			log.Println("error in first number conversion:", err)
			return num1, num2, operator, err
		}

		operator = parameter[2]
		err = checkOperator(operator)
		log.Println("operator:", operator)

		return num1, num2, operator, err

	case len(parameter) == 5:
		n1, err := strconv.ParseFloat(parameter[1], 32)
		if err == nil {
			num1 = float32(n1)
			log.Printf("Type: %T \n", num1)
			log.Println("Value:", num1)
		} else {
			log.Println("error in first number conversion:", err)
			return num1, num2, operator, err
		}

		n2, err := strconv.ParseFloat(parameter[4], 64)
		if err == nil {
			num2 = float32(n2)
			log.Printf("Type: %T \n", num2)
			log.Println("Value:", num2)
		} else {
			log.Println("error in second number conversion:", err)
			return num1, num2, operator, err
		}

		operator = "/"

		log.Println("operator:", operator)
		err = checkOperator(operator)
		return num1, num2, operator, err

	default:
		log.Println("error: wrong number of arguments")
		err = fmt.Errorf("wrong number of arguments")
		return num1, num2, operator, err
	}

}

func checkOperator(operator string) (err error) {
	if operator == "+" || operator == "-" || operator == "/" || operator == "*" {
		err = nil
		return err
	} else {
		err = fmt.Errorf("the operator %s is not defined", operator)
		return err
	}
}

func calculate(u UserInput) (float32, error) {
	switch u.Operator {
	case "+":
		return u.AddValues()
	case "-":
		return u.SubtractValues()
	case "/":
		return u.DevideValues()
	case "*":
		return u.MultiplyValues()
	default:
		return 0, fmt.Errorf("the operator %s is not defined", u.Operator)
	}
}
