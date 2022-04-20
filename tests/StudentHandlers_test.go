package tests

import (
	"bytes"
	"encoding/json"
	"internal/config"
	"internal/entities"
	"internal/resources"
	"internal/web/rest"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

func init() {
	config.IsMemoryDAONecessary = true
}

func TestGetStudents(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rest.GetStudents)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	result := rr.Body.String()
	var students []entities.Student
	err = json.Unmarshal([]byte(result), &students)
	if err != nil {
		t.Errorf("handler returned unexpected body type: don't get got []Student")
	}

	if !reflect.DeepEqual(students, entities.RandomStudents) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			students, entities.RandomStudents)
	}
}

func TestGetStudentById(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rest.GetStudentById)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	result := rr.Body.String()
	var student entities.Student
	err = json.Unmarshal([]byte(result), &student)
	if err != nil {
		t.Errorf("handler returned unexpected body type: don't get got []Student")
	}

	if !reflect.DeepEqual(student, entities.RandomStudents[0]) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			student, entities.RandomStudents[0])
	}
}

func TestGetStudentByIdWithUndefinedId(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"id": "10",
	})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rest.GetStudentById)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

	result := rr.Body.String()

	if result != resources.NotFoundResourceJson {
		t.Errorf("handler returned unexpected body: got %v want %v",
			result, resources.NotFoundResourceJson)
	}
}

func TestPostStudent(t *testing.T) {
	newStudent := entities.NewStudent(4, 30, "newStudentFirstname", "newStudentLastname", "go")
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(newStudent.Marshal()))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rest.PostStudent)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	result := rr.Body.String()

	if result != resources.SuccessfulAdditionJson {
		t.Errorf("handler returned unexpected body: got %v want %v",
			result, resources.SuccessfulAdditionJson)
	}
}

func TestPostStudentWithWrongBodyContent(t *testing.T) {
	newStudent := `{
        "id": 4,
        "firstname": "newStudentFirstname",
        "lastname": "newStudentLastname",
        "age": 30,
        languageCode: go
    }`

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(newStudent)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rest.PostStudent)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	result := rr.Body.String()

	if result != resources.MarshalingErrorJson {
		t.Errorf("handler returned unexpected body: got %v want %v",
			result, resources.MarshalingErrorJson)
	}
}

func TestPutStudent(t *testing.T) {
	studentToUpdate := entities.NewStudent(1, 20, "updatedFirstname", "updatedLastname", "go")
	req, err := http.NewRequest("PUT", "/", bytes.NewBuffer(studentToUpdate.Marshal()))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rest.PutStudent)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	result := rr.Body.String()

	if result != resources.SuccessfulUpdateJson {
		t.Errorf("handler returned unexpected body: got %v want %v",
			result, resources.SuccessfulUpdateJson)
	}
}

func TestPutStudentWithWrongBodyContent(t *testing.T) {
	studentToUpdate := `{
        "id": 1,
        "firstname": "updatedFirstname",
        "lastname": "updatedLastname",
        "age": 20,
        languageCode: go
    }`

	req, err := http.NewRequest("PUT", "/", bytes.NewBuffer([]byte(studentToUpdate)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rest.PutStudent)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	result := rr.Body.String()

	if result != resources.MarshalingErrorJson {
		t.Errorf("handler returned unexpected body: got %v want %v",
			result, resources.MarshalingErrorJson)
	}
}

func TestPutStudentWhileStudentIsNotExisting(t *testing.T) {
	studentToUpdate := `{
        "id": 10,
        "firstname": "updatedFirstname",
        "lastname": "updatedLastname",
        "age": 20,
        "languageCode": "go"
    }`

	req, err := http.NewRequest("PUT", "/", bytes.NewBuffer([]byte(studentToUpdate)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rest.PutStudent)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	result := rr.Body.String()

	if result != resources.UnsuccessfulUpdateJson {
		t.Errorf("handler returned unexpected body: got %v want %v",
			result, resources.UnsuccessfulUpdateJson)
	}
}

func TestDeleteStudentById(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rest.DeleteStudentById)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	result := rr.Body.String()

	if result != resources.SuccessfulDeletionJson {
		t.Errorf("handler returned unexpected body: got %v want %v",
			result, resources.SuccessfulDeletionJson)
	}
}

func TestDeleteStudentByIdWithUndefinedId(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"id": "10",
	})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rest.DeleteStudentById)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

	result := rr.Body.String()

	if result != resources.NotFoundResourceJson {
		t.Errorf("handler returned unexpected body: got %v want %v",
			result, resources.NotFoundResourceJson)
	}
}
