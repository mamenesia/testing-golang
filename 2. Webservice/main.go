package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/count", handlerCount)
	http.HandleFunc("/sort", handleSort)

	address := "localhost:9000"
	fmt.Printf("Server started at %s\n", address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
