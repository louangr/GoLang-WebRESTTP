package tests

import (
	"internal/config"
	"internal/web/rest"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetStudents(t *testing.T) {
	config.IsMemoryDAONecessary = true

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

	expected := `[{"id":1,"firstname":"Joe","lastname":"Doe","age":20,"languageCode":"fra"},{"id":2,"firstname":"Bob","lastname":"Doe","age":21,"languageCode":"fra"},{"id":3,"firstname":"Bob","lastname":"USA","age":21,"languageCode":"eng"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
