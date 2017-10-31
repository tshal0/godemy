package main

import (
	"net/http"

	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	_ "strconv"
	"app/model"
	"app/route"

)

const (
	DB_USER 		= "postgres"
	DB_PASSWORD 	= "secret"
	DB_NAME 		= "godemy_dev"
)

func main() {

	//////////////////////////////////////////////////////////////
	// Database connection
	//////////////////////////////////////////////////////////////

	context, err := db.NewDB("user=postgres password=secret dbname=godemy_dev sslmode=disable")
	if err != nil {
		log.Println(err.Error())
	}
	
	env := &db.Env{Context: context}
	log.Println("Get Users")
	users, err := env.Context.Users()
	
	for _, usr := range users {	
		log.Printf("%d | %s | %s \n", usr.Userid, usr.Name, usr.Passwordhash)
	}
	
	//////////////////////////////////////////////////////////////
	// Original method:
	//////////////////////////////////////////////////////////////

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.IndexHandler)
	r.HandleFunc("/login", handlers.Authenticate).Methods("GET")
	r.HandleFunc("/register", handlers.Register).Methods("GET")
	r.HandleFunc("/static/{file}", handlers.StaticHandler)
	http.Handle("/", r)
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
