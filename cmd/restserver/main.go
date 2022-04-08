package main

import (
	"fmt"
	"internal/web/rest"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
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
