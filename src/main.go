package main

import "net/http"

func main() {
	http.HandleFunc("/api", ApiHandler)
	http.ListenAndServe(":8080", nil)
}
