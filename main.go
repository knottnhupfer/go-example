package main

import (
	"log"
	"net/http"

	"service/handlers"
)

func main() {

	var number uint64 = 56

	handlers.RegisterHandlers()

	log.Println("Starting server...")
	http.ListenAndServe("localhost:7080", nil)

	println("Number is: ", number)
}
