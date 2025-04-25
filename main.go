package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on [::]:800 (IPv6+IPv4)")
	err := http.ListenAndServe("[::]:800", nil)
	if err != nil {
		panic(err)
	}
}
