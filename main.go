package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sendHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse args, you have to call this by yourself
	fmt.Println(r.Form) // Print form information server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Godemy") // Send data to client side.
}

func main() {
	http.HandleFunc("/", sendHello)
	err := http.ListenAndServe(":8080", nil) // Set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
