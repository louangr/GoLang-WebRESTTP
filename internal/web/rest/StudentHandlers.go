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

// swagger:operation GET /students student GetStudents
// ---
// summary: Return all students
// description: If the are not students, an empty array will be returned
// responses:
//   "200":
//     "$ref": "#/responses/studentStructArray"

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

// swagger:operation GET /students/{id} student GetStudentById
// ---
// summary: Return a student by Id
// description: If the student is not found, a 404 status code will be returned
// parameters:
// - name: id
//   in: path
//   description: student Id
//   type: integer
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/studentStruct"
//   "404":
//     "$ref": "#/responses/genericResponse"

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

// swagger:operation POST /students student PostStudent
// ---
// summary: Create a new student
// description: If the request body format is not correct, a 400 status code will be returned
// produces:
// - application/json
// parameters:
// - name: student
//   in: request body
//   description: a student object
//   required: true
//   type: Student
// responses:
//   "200":
//     "$ref": "#/responses/genericResponse"
//   "400":
//     "$ref": "#/responses/genericResponse"

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

// swagger:operation PUT /students student PutStudent
// ---
// summary: Update an existing student
// description: If the request body format is not correct or the target student Id is not found, a 400 status code will be returned
// produces:
// - application/json
// parameters:
// - name: student
//   in: request body
//   description: a student object
//   required: true
//   type: Student
// responses:
//   "200":
//     "$ref": "#/responses/genericResponse"
//   "400":
//     "$ref": "#/responses/genericResponse"

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

// swagger:operation DELETE /students/{id} student DeleteStudentById
// ---
// summary: Delete a student by Id
// description: If the student is not found, a 404 status code will be returned
// parameters:
// - name: id
//   in: path
//   description: student Id
//   type: integer
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/genericResponse"
//   "404":
//     "$ref": "#/responses/genericResponse"

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
