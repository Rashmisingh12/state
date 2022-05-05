package main

import (
	// "fmt"

	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	// http.HandleFunc("/read",read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9091", nil)

}
func set(w http.ResponseWriter, req *http.Request) {
	a, err := req.Cookie("my-cookies")
	if err == http.ErrNoCookie {
		a = &http.Cookie{
			Name:  "my-cookies",
			Value: "0",
		}
	}

	// if err==http.ErrNoCookies{
	// cookie=&http.Cookie{
	// 	Name:"cookies",
	// 	Value:"0",
	// }

	count, _ := strconv.Atoi(a.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	a.Value = strconv.Itoa(count)
	http.SetCookie(w, a)
	io.WriteString(w, a.Value)

}
