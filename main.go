package main

import (
	"fmt"
	"net/http"
)

func doNothing(res http.ResponseWriter, req *http.Request) {}

func handlerFunc(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Someone visited our page")
	fmt.Fprint(res, "<h1>Welcome to my awesome site!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/favicon.ico", doNothing)
	http.ListenAndServe(":3333", nil)
}
