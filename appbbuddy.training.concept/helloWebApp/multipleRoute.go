package helloWebApp

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

//Home is the home page handeler
func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "this is home page")
	if err != nil {
		log.Println("error in Home page handeler -", err) // using log good for debuging
	}
	log.Println("number of bytes written through home page is -", n)
}

//About is the about page handeler
func About(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "this is About page")
	if err != nil {
		log.Println("error in About page handeler -", err) // using log good for debuging
	}
	log.Println("number of bytes written through about page is -", n)
}

func MultipleRoute() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// start listening at http://localhost:8080 and give the result second arg nill as already defind handeler
	log.Println("starting application on portNumber - ", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
