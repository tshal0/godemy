package main

import (
	// "time"
	// "flag"
	"fmt"
	"net/http"
	"strings"
	"log"
	"io/ioutil"
	"github.com/gorilla/mux"

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
func indexHandler (w http.ResponseWriter, r *http.Request) {
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

func staticHandler (w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path[1:]
	log.Println("path: " + path)

	data, err := ioutil.ReadFile(path)

	if err == nil {
		var contentType string

		contentType = "text/html"

		w.Header().Add("Content Type", contentType)
		w.Write(data)
	}

}

func loginHandler (w http.ResponseWriter, r *http.Request) {
	path := "./static/login.html"
	data, err := ioutil.ReadFile(path)

	if err == nil {
		var contentType string

		contentType = "text/html"

		w.Header().Add("Content Type", contentType)
		w.Write(data)
	}

}

func main() {

	//////////////////////////////////////////////////////////////
	// Original method:
	//////////////////////////////////////////////////////////////

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/static/{file}", staticHandler)
	http.Handle("/", r)
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	//////////////////////////////////////////////////////////////
	// New method:
	//////////////////////////////////////////////////////////////

	// var dir string

	// flag.StringVar(&dir, "dir", ".", "the directory to serve files from")
	// flag.Parse()

	// r := mux.NewRouter()

	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	// r.HandleFunc("/", indexHandler)

	// srv := &http.Server{
	// 	Handler: 	r,
	// 	Addr:		"127.0.0.1:8080",
	// 	WriteTimeout: 15* time.Second,
	// 	ReadTimeout: 15* time.Second, 
	// }

	// log.Fatal(srv.ListenAndServe())


	// http.HandleFunc("/", sendHello)
	// err := http.ListenAndServe(":8080", nil) // Set listen port
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
}
