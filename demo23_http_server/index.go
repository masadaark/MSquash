package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

//build-in package

func main() {

	//if path not found return 404 page not found
	// write to response String + urlPath
	http.HandleFunc("/Home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home,%q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Profile,%q", html.EscapeString(r.URL.Path))
	})

	//handle string query http://localhost:88/login?username=admin&password=1234
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home, %s, %s", r.URL.Query().Get("username"), r.URL.Query().Get("password"))
	})

	log.Fatal(http.ListenAndServe(":88", nil))
}
