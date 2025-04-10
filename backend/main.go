package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Backend Golang jalan!!! Selamat datang !")
	})
	fmt.Println("Server backend jalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
