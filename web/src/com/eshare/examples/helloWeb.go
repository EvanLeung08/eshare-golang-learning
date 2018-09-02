package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", helloWeb)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

func helloWeb(w http.ResponseWriter, r *http.Request) {

	//解析参数
	r.ParseForm()
	println(r.Form)
	println(r.URL.Path)
	println(r.URL.Scheme)
	println(r.Form["age"])

	for k, v := range r.Form {
		println("key:", k)
		println("value:", strings.Join(v, ""))

	}

	fmt.Fprint(w, "Evan")
}
