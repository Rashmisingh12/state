package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/moreValue", moreValue)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8081", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookies",
		Value: "Rashmi Singh",
	})
	fmt.Fprintln(w, "COOKIES WRITTEN-CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to:dev tool/appliation/cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	c1, err := req.Cookie("my-cookies")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIES #1:", c1)
	}
	c2, err := req.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIES #2:", c2)
	}
	c3, err := req.Cookie("original")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIES #3:", c3)
	}
}
func moreValue(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "i am here",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "original",
		Value: "dont change your originality",
	})
	fmt.Fprintln(w, " more COOKIES WRITTEN-CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to:dev tool/appliation/cookies")
}
