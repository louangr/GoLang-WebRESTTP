package main

import (
	"fmt"
	"internal/web/rest"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// @title Student and sutend's language API documentation
// @version 1.0.0
// @host localhost:8000
// @BasePath /api

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/api/students", rest.GetStudents).Methods("GET")
	myRouter.HandleFunc("/api/students/{id:[0-9]+}", rest.GetStudentById).Methods("GET")
	myRouter.HandleFunc("/api/students", rest.PostStudent).Methods("POST")
	myRouter.HandleFunc("/api/students", rest.PutStudent).Methods("PUT")
	myRouter.HandleFunc("/api/students/{id:[0-9]+}", rest.DeleteStudentById).Methods("DELETE")

	myRouter.HandleFunc("/api/languages", rest.GetLanguages).Methods("GET")
	myRouter.HandleFunc("/api/languages/{code}", rest.GetLanguageById).Methods("GET")
	myRouter.HandleFunc("/api/languages", rest.PostLanguage).Methods("POST")
	myRouter.HandleFunc("/api/languages", rest.PutLanguage).Methods("PUT")
	myRouter.HandleFunc("/api/languages/{code}", rest.DeleteLanguageById).Methods("DELETE")

	fmt.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
