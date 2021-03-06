package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	newServer()
}

func getUser(url string) []byte {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return body
}

func bodyToJSON(body []byte) {
	var user map[string]interface{}
	jsonErr := json.Unmarshal(body, &user)

	if jsonErr != nil {
		panic(jsonErr)
	}
	fmt.Println(user["data"])
}

func saveToFile(body []byte) {
	bodyReader := bytes.NewReader(body)
	file, err := os.Create("response.txt")
	if err != nil {
		panic(err)
	}

	io.Copy(file, bodyReader)
}
