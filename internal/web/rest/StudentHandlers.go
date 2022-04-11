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

var studentDAOMemory = persistence.NewStudentDAOBolt()

func GetStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetStudents")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := studentDAOMemory.GetAll()

	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println(resources.MarshalingError, err)
		fmt.Fprintf(w, resources.InternalErrorJson)
		return
	}

	fmt.Fprintf(w, "%s", j)
}

func GetStudentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	intId, _ := strconv.Atoi(id)
	fmt.Printf("GetStudentById (%s)\n", id)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := studentDAOMemory.Get(intId)

	if data.Id == -1 {
		fmt.Fprintf(w, resources.NotFoundResourceJson)
		return
	}

	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println(resources.MarshalingError, err)
		fmt.Fprintf(w, resources.InternalErrorJson)
		return
	}

	fmt.Fprintf(w, "%s", j)
}

func PostStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PostStudent")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var newStudent entities.Student
	err := json.NewDecoder(r.Body).Decode(&newStudent)

	if err != nil {
		fmt.Fprintf(w, resources.MarshalingErrorJson)
		return
	}

	hasBeenSaved := studentDAOMemory.Save(newStudent)

	if hasBeenSaved {
		fmt.Fprintf(w, resources.SuccessfulAdditionJson)
	} else {
		fmt.Fprintf(w, resources.UnsuccessfulAdditionJson)
	}
}

func PutStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PutStudent")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var student entities.Student
	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {
		fmt.Fprintf(w, resources.MarshalingErrorJson)
		return
	}

	hasBeenUpdated := studentDAOMemory.Update(student)

	if hasBeenUpdated {
		fmt.Fprintf(w, resources.SuccessfulUpdateJson)
	} else {
		fmt.Fprintf(w, resources.UnsuccessfulUpdateJson)
	}
}

func DeleteStudentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	intId, _ := strconv.Atoi(id)
	fmt.Printf("DeleteLanguageById (%s)\n", id)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	hasBeenDeleted := studentDAOMemory.Delete(intId)

	if hasBeenDeleted {
		fmt.Fprintf(w, resources.SuccessfulDeletionJson)
	} else {
		fmt.Fprintf(w, resources.UnsuccessfulDeletionJson)
	}
}
