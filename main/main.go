package main

import (
	"fmt"
	"net/http"

	f "ascii-art-web/func"
)

func main() {
	// Set up the HTTP server routes
	http.HandleFunc("/styles/", f.ServeStyle)
	fmt.Println("The server is working now :")
	http.HandleFunc("/", f.Welcom)
	http.HandleFunc("/ascii-art", f.Last)
	http.HandleFunc("/Download", f.Download)

	// Start the HTTP server
	fmt.Println("the server is running on localhost port 8088")
	fmt.Println("http://localhost:7777")
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		fmt.Println(err)
	}
}
