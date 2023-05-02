package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/appBuddyin/go-lang/appbbuddy.training.concept/dateMicroservice"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	// our dateMicroservice service
	srv := dateMicroservice.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := dateMicroservice.Endpoints{
		GetEndpoint:      dateMicroservice.MakeGetEndpoint(srv),
		StatusEndpoint:   dateMicroservice.MakeStatusEndpoint(srv),
		ValidateEndpoint: dateMicroservice.MakeValidateEndpoint(srv),
	}

	// HTTP transport
	go func() {
		log.Println("dateMicroservice is listening on port:", *httpAddr)
		handler := dateMicroservice.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
