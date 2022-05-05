package main

import (
	"io"
	"net/http"
)
func main(){
	http.HandleFunc("/",abc)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func abc(w http.ResponseWriter,req *http.Request){
	v:=req.FormValue("q")
	io.WriteString(w,"hello my name is:"+v)
}