package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Http client

func main3() {

	client := &http.Client{}

	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response body:", err)
		return
	}

	fmt.Println("response from server:", string(body))
}
