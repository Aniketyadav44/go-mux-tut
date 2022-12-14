package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := Router

	HandleRoutes()

	fmt.Println("Server started on port:8000")

	log.Fatal(http.ListenAndServe(":8000", r))
}
