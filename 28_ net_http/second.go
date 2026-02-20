package main

import (
	"fmt"
	"net/http"
)

// http sever with routing

func main2() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Go")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "welcoming users")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("server error:", err)
	}

	fmt.Println("server is running on port 8080")
}
