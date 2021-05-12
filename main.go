package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	res.Header().Set("Content-Type", "text/html")
	if req.URL.Path == "/" {
		fmt.Fprint(res, "<h1>Welcome to my awesome site!</h1>")
	} else if req.URL.Path == "/contact" {
		fmt.Fprint(res, "To get in touch, please send an email to <a href=\"mailto:support@photos.kurabi.net\">support@photos.kurabi.net</a>")
	} else if req.URL.Path == "favicon.ico" {
		return
	} else {
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprint(res, "<h1>Ooops... We could not find thepage you were looking for<h1>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3333", nil)
}
