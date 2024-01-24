package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error on Server", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	fmt.Println(strings.ToUpper(string(body)))
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server is running on 127.0.0.1:8181 port")
	err := http.ListenAndServe("127.0.0.1:8181", nil)

	if err != nil {
		fmt.Println("Error Server:", err)
	}
}
