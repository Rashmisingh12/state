package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func set(w http.ResponseWriter, req *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookies",
		Value: "some rashmi",
	})
	fmt.Fprintln(w, "COOKIES WRITTEN-CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to:dev tool/appliation/cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookies")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "YOUR COOKIES:", c)
}
