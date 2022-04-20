package rest

import (
	"encoding/json"
	"fmt"
	"internal/entities"
	"internal/persistence"
	"internal/resources"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetStudents(w http.ResponseWriter, r *http.Request) {
	var studentDAO = *persistence.GetStudentDAOInstance()
	fmt.Println("GetStudents")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := studentDAO.GetAll()

	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println(resources.MarshalingError, err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, resources.InternalErrorJson)
		return
	}

	fmt.Fprintf(w, "%s", j)
}

func GetStudentById(w http.ResponseWriter, r *http.Request) {
	var studentDAO = *persistence.GetStudentDAOInstance()
	vars := mux.Vars(r)
	id := vars["id"]
	intId, _ := strconv.Atoi(id)
	fmt.Printf("GetStudentById (%s)\n", id)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := studentDAO.Get(intId)

	if data.Id == -1 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, resources.NotFoundResourceJson)
		return
	}

	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println(resources.MarshalingError, err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, resources.InternalErrorJson)
		return
	}

	fmt.Fprintf(w, "%s", j)
}

func PostStudent(w http.ResponseWriter, r *http.Request) {
	var studentDAO = *persistence.GetStudentDAOInstance()
	fmt.Println("PostStudent")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var newStudent entities.Student
	err := json.NewDecoder(r.Body).Decode(&newStudent)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, resources.MarshalingErrorJson)
		return
	}

	hasBeenSaved := studentDAO.Save(newStudent)

	if hasBeenSaved {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, resources.SuccessfulAdditionJson)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, resources.UnsuccessfulAdditionJson)
	}
}

func PutStudent(w http.ResponseWriter, r *http.Request) {
	var studentDAO = *persistence.GetStudentDAOInstance()
	fmt.Println("PutStudent")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var student entities.Student
	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, resources.MarshalingErrorJson)
		return
	}

	hasBeenUpdated := studentDAO.Update(student)

	if hasBeenUpdated {
		fmt.Fprintf(w, resources.SuccessfulUpdateJson)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, resources.UnsuccessfulUpdateJson)
	}
}

func DeleteStudentById(w http.ResponseWriter, r *http.Request) {
	var studentDAO = *persistence.GetStudentDAOInstance()
	vars := mux.Vars(r)
	id := vars["id"]
	intId, _ := strconv.Atoi(id)
	fmt.Printf("DeleteLanguageById (%s)\n", id)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	hasBeenDeleted := studentDAO.Delete(intId)

	if hasBeenDeleted {
		fmt.Fprintf(w, resources.SuccessfulDeletionJson)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, resources.NotFoundResourceJson)
	}
}
