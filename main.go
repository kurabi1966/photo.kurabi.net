package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "<h1>Welcome to my awesome site!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3333", nil)
	// fmt.Println("Learnning GoLang!")
}
