package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}

}
func login(writer http.ResponseWriter, request *http.Request) {
	println("request method:", request.Method)

	if "GET" == request.Method {
		template, _ := template.ParseFiles("login.html")
		template.Execute(writer, nil)
	} else {
		println("username:", request.Form["username"])
		println("password:", request.Form["password"])
	}
}
func index(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	println(request.Form)
	println("path:", request.URL.Path)
	println("scheme:", request.URL.Scheme)
	println("age:", request.Form["age"])

	for k, v := range request.Form {
		println("key:", k)
		println("value:", v)
	}

	fmt.Fprint(writer, "welcome to my website")
}
