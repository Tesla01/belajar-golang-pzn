package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Success")

	server := http.Server{
		Addr: "0.0.0.0:9000",
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error")
	}
}
