package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}

	// GET request
	resp, err := client.Get("https://www.example.com")
	if err != nil {
		fmt.Println("Error ", err)
		return
	}
	// Body response close
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// print body response
	fmt.Println(string(body))
}
