package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	ID   uint32 `json:"id"`
}

// Used to delete data.
// This method makes a request to delete the data on the server.
func Delete() {
	// create a new HTTP client
	client := &http.Client{}

	// create a new DELETE request
	req, err := http.NewRequest(http.MethodDelete, "https://httpbin.org/delete", nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	// send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	// read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	// print the response body
	fmt.Println(string(body))
}

// Used to retrieve data.
// Used to pass data through parameters in the URL.
// This method is usually used to get information from the server.
func Get() {
	// get methodi yordamida response olish
	resp, err := http.Get("https://httpbin.org/get")

	if err != nil {
		log.Fatal(err)
		return
	}

	defer resp.Body.Close()

	// response ma'lumotlarini o'qib olish
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	// responseni chiqarish
	fmt.Printf(string(data))
}

// Used to truncate data.
// Used to modify data with multiple fields through this method.
func Patch() {
	url := "https://httpbin.org/patch"
	payload := map[string]string{
		"name": "New Name",
	}

	// encode payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
		return
	}

	// create new HTTP PATCH request with JSON payload
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatal(err)
		return
	}

	// set content-type header to JSON
	req.Header.Set("Content-Type", "application/json")

	// create HTTP client and execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	// read the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	// print the response body
	fmt.Println(string(data))
}

// Used to send data to the server.
// Through this method, the browser sends data to the server.
// Data is sent in the "body" part of the package inside the HTTP request.
func Post() {
	// create new struct object
	var u = User{
		Name: "Alex",
		ID:   1,
	}

	// encode payload to JSON
	bytesRepresentation, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
		return
	}

	// create new HTTP POST request with JSON payload
	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatal(err)
		return
	}

	// read the response body
	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	// print the response body
	fmt.Println(string(bytesResp))
}

// Used to change data.
// Through this method, the browser sends data to the server, and the server updates the data.
func Put() {
	url := "https://httpbin.org/put"
	data := []byte(`{"name": "John", "age": 30}`)

	// create new HTTP PUT request with JSON payload
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
		return
	}

	// set content-type header to JSON
	req.Header.Set("Content-Type", "application/json")

	// send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	// read the response body
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	// print the response body
	fmt.Println(string(data))
}

func main() {
	fmt.Println("Delete Method")
	Delete()
	fmt.Println("Get Method")
	Get()
	fmt.Println("Patch Method")
	Patch()
	fmt.Println("Post Method")
	Post()
	fmt.Println("Put Method")
	Put()
}
