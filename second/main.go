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
   w.Header().Set("Content-Type","text/html;charset=utf-8")
   io.WriteString(w,`
   <form method="get">
   <input type="text" name="q">
   <input type="Submit">
   </form>

   <br>`+v)
}