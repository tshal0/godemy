package main

import (
	"database/sql"
	// "time"
	// "flag"
	"fmt"
	"net/http"
	"strings"
	"log"
	"io/ioutil"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

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

const (
	DB_USER 		= "postgres"
	DB_PASSWORD 	= "secret"
	DB_NAME 		= "godemy_dev"
)

type User struct {
	userid int64
	name string
}

func main() {

	//////////////////////////////////////////////////////////////
	// Database connection
	//////////////////////////////////////////////////////////////

	
	connStr := "user=postgres dbname=godemy_dev sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	rows, err := db.Query("SELECT * FROM public.users")
	if err != nil {
		log.Println(err.Error())
	}
	
	for rows.Next() {
		var userid int
		var name string
		err = rows.Scan(&userid, &name)
		if err != nil{
			log.Println(err.Error())
		}
		log.Println("userid | name")
		log.Printf("%6v | %6v\n", userid, name)
	}

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


}
