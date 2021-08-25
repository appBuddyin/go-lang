package calculator

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Cal() {
	var num1 float32
	var num2 float32
	var operator string
	var err error
	var message string
	router := gin.Default()
	router.GET("/*path", func(c *gin.Context) {
		path := c.Param("path")
		parameter := strings.Split(path, "/")

		log.Printf("parameter type = %T ", parameter)
		log.Println("parameter length = ", len(parameter))
		log.Println("0 = ", parameter[0])
		log.Println("1 = ", parameter[1])
		log.Println("2 = ", parameter[2])
		log.Println("3 = ", parameter[3])

		num1, num2, operator, err = values(parameter)
		if err == nil {
			message = fmt.Sprintf("%f %s %f=", num1, operator, num2)
		} else {
			message = fmt.Sprint("an error occurred -", err)
		}
		c.String(http.StatusOK, message)
	})

	router.Run(":8080")
}

func values(parameter []string) (num1, num2 float32, operator string, err error) {
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

		return num1, num2, operator, err

	default:
		log.Println("error: wrong number of arguments")
		err = fmt.Errorf("wrong number of arguments")
		return num1, num2, operator, err
	}

}

func checkOperator(operator string)(err error){
	if operator == "+" or
}