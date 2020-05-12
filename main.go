package main

/**
encoding/json - T
his package contains methods that are used to convert Go types to JSON and
vice-versa (This conversion is called as encode/decode in Go,
	serialization/de-serialization or
	marshall/unmarshall in other languages).

fmt - T
his package implements formatted I/O functions similar to scanf and printf in C

io/ioutil - T
his package implements some I/O utility functions
(For instance, reading the contents of a file , reading from a io.Reader, etc)

log -
Has methods for formatting and printing log messages.

net/http - Contains methods for performing operations over HTTP.
It provides HTTP server and client implementations and has abstractions for HTTP request, response, headers, etc.
*/
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:completed`
}

func get() {
	// Example for calling an API with a know data structure
	fmt.Println("1. Performing HTTP Get...")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// convert response body to Todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("API Response as struck %+v\n", todoStruct)
}

func getStock() {
	// unknown data structure
	fmt.Println("1. Get Stock API")
	url := "https://alpha-vantage.p.rapidapi.com/query?datatype=json&symbol=MSFT&function=TIME_SERIES_MONTHLY"
	api_key := "8a443e481emshe2c3e916c5d31d6p12c4f1jsn8ee582ea634f"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-rapidapi-host", "alpha-vantage.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", api_key)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var jsonData interface{}
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		log.Fatalln(err)
	}
	data := jsonData.(map[string]interface{})

	fmt.Println(res)
	//fmt.Println(string(body))
	fmt.Printf("%+v\n", data)

}

func main() {
	get()
	getStock()
}
