package main

import (
	"net/http"
	"./objects"
)

func main() {

	http.HandleFunc("/objects/", objects.Handler)

	http.ListenAndServe(":27721", nil)
}