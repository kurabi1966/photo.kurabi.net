package main

import (
	"fmt"
	"net/http"
)

func doNothing(res http.ResponseWriter, req *http.Request) {}

func handlerFunc(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprint(res, "<h1>Welcome to my awesome site!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/favicon.ico", doNothing)
	http.ListenAndServe(":3333", nil)
}
