package web

import (
	"fmt"
	"net/http"
	"time"
	// "io/ioutil"
)

func fooHandler(w http.ResponseWriter,r *http.Request){
	fmt.Println(r.URL)
	w.Header().Set("Location","http://www.baidu.com")
	w.WriteHeader(http.StatusMovedPermanently)
	time.Sleep(1*time.Second)
	w.Write([]byte("<html><body> <h1> Title</h1></body></html>"))
}

func barHandler(w http.ResponseWriter,r *http.Request){
	fmt.Println(r.URL)
	w.Write([]byte("bar"))
}

func StartServer(){

	http.HandleFunc("/foo",fooHandler)
	http.HandleFunc("/bar",barHandler)
	http.ListenAndServe(":8080",nil)
}