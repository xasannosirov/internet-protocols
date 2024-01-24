package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func sendData(data string) {
	resp, err := http.Post("http://127.0.0.1:8181", "text/plain", bytes.NewBufferString(data))
	if err != nil {
		fmt.Println("Error write to server:", err)
		return
	}
	defer resp.Body.Close()
}

func main() {
	for {
		var data string
		fmt.Print("Enter data: ")
		fmt.Scan(&data)

		sendData(data)
	}
}
