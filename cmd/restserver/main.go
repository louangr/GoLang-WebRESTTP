// Package classification Students and Languages APIs
//
// Students and Languages APIs
//
// Terms Of Service:
//
//	Schemes: http, https
//	Version: 1.0.0
//	BasePath: /
//	Contact: Louan <by@carrier.pigeon>
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

import (
	"fmt"
	"internal/persistence"
	"internal/web/rest"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	dbFileName := os.Args[2]
	persistence.InitBoldDB(dbFileName)
	defer persistence.BoltDBInstance.DB.Close()

	portNumber := os.Args[1]
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/students", rest.GetStudents).Methods("GET")
	myRouter.HandleFunc("/students/{id:[0-9]+}", rest.GetStudentById).Methods("GET")
	myRouter.HandleFunc("/students", rest.PostStudent).Methods("POST")
	myRouter.HandleFunc("/students", rest.PutStudent).Methods("PUT")
	myRouter.HandleFunc("/students/{id:[0-9]+}", rest.DeleteStudentById).Methods("DELETE")

	myRouter.HandleFunc("/languages", rest.GetLanguages).Methods("GET")
	myRouter.HandleFunc("/languages/{code}", rest.GetLanguageById).Methods("GET")
	myRouter.HandleFunc("/languages", rest.PostLanguage).Methods("POST")
	myRouter.HandleFunc("/languages", rest.PutLanguage).Methods("PUT")
	myRouter.HandleFunc("/languages/{code}", rest.DeleteLanguageById).Methods("DELETE")

	fmt.Println("Server started on port", portNumber)
	log.Fatal(http.ListenAndServe(":"+portNumber, myRouter))
}
