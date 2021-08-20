package helloWebApp

import (
	"fmt"
	"log"
	"net/http"
)

func HelloApp() {
	// this is a handeler function for http request defined inside it is a inline function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Welcome page at the root")
		if err != nil {
			log.Println("error -", err) // using log good for debuging
		}
		log.Println("number of bytes written is -", n)
	})

	// start listening at http://localhost:8080 and give the result second arg nill as already defind handeler
	_ = http.ListenAndServe(":8080", nil)
}
