package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const TEMPLATE_DIR string = "D:/Repository/eshare-golang-learning/web-form/src/com/eshare/examples/static/"

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
		template, _ := template.ParseFiles(TEMPLATE_DIR + "/login.html")
		//template.Execute(writer, nil)
		log.Println(template.Execute(writer, nil))
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
