package handlers

import (
	// "time"
	// "flag"
	"fmt"
	"net/http"
	"strings"
	"log"
	"io/ioutil"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	_ "strconv"
	_ "app/model"
	


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
	fmt.Fprintf(w, "<h1>Hello Godemy<h1/>") // Send data to client side.
}

// Assuming this is how we generate data, maybe in the future it would be from a db.
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// data := map[string]interface{}{
	// 	"id":"12345",
	// 	"ts":time.Now().Format(time.RFC3339),
	// }

	// b, err := json.Marshal(data)
	// if err != nil {
	// 	http.Error(w, err.Error(), 400)
	// 	return
	// }

	// w.Write(b)
}

// Index handler to serve our React.js front end
func IndexHandler (w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Println(path)

	if path == "" {
		path = "./static/index.html"
	}
	log.Println(path)
	data, err := ioutil.ReadFile(path)

	if err == nil {
		var contentType string

		contentType = "text/html"

		w.Header().Add("Content Type", contentType)
		w.Write(data)
	}

}

func StaticHandler (w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	

	file := vars["file"]
	
	var contentType string 
	var path string
	if strings.HasSuffix(file, ".js") {
		log.Println("JS")
		contentType = "text/javascript"
		path = "static/js/" + file
	} else if strings.HasSuffix(file, ".html") {
		log.Println(file)
		contentType = "text/html"
		path = "static/" + file
	} else {
		contentType = "text/plain"
		log.Println("PLAIN")
		path = r.URL.Path[1:]
	}

	data, err := ioutil.ReadFile(path)

	if err == nil {
		log.Println("Content type: " + contentType + " | " + 
			"File: " + path)
		w.Header().Add("Content-Type", contentType)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}

}

func LoginHandler (w http.ResponseWriter, r *http.Request) {
	path := "./static/login.html"
	data, err := ioutil.ReadFile(path)

	if err == nil {
		var contentType string

		contentType = "text/html"

		w.Header().Add("Content Type", contentType)
		w.Write(data)
	}

}

func Authenticate (w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("username")
	pass := r.FormValue("password")

	log.Printf("UserName: %s | Password: %s\n", name, pass)
	w.WriteHeader(http.StatusAccepted)
}