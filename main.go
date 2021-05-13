package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Handle404(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNotFound)
	fmt.Fprint(res, "<h1>Ooops... We could not find thepage you were looking for<h1>")
}

func home(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprint(res, "<h1>Welcome to my awesome site!</h1>")
}

func contact(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprint(res, "To get in touch, please send an email to <a href=\"mailto:support@photos.kurabi.net\">support@photos.kurabi.net</a>")
}

func faq(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprint(res, "<h3>FAQ Page content goes here!</h3>")
}

func doNothing(res http.ResponseWriter, req *http.Request) {}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)

	r.HandleFunc("/favicon", doNothing)
	r.NotFoundHandler = http.HandlerFunc(Handle404)
	http.ListenAndServe(":3333", r)
}
