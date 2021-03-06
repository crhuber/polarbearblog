package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/akrylysov/algnhsa"
	"github.com/crhuber/polarbearblog/app"
	"github.com/crhuber/polarbearblog/app/lib/timezone"
)

func init() {
	// Verbose logging with file name and line number.
	log.SetFlags(log.Lshortfile)
	// Set the time zone.
	timezone.Set()
}

func main() {
	handler, err := app.Boot()
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Start the web server.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// use lambda
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		fmt.Println("Lambda server running")
		opts := &algnhsa.Options{UseProxyPath: true}
		algnhsa.ListenAndServe(handler, opts)
	} else {
		fmt.Println("Web server running on port:", port)
		log.Fatalln(http.ListenAndServe(":"+port, handler))
	}
}
