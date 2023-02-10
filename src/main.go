package main

import "net/http"

func main() {
	http.HandleFunc("/ptlist", PeriodicTaskHandler)
	http.ListenAndServe(":3000", nil)
}