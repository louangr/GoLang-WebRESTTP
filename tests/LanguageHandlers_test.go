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

func TestGetLanguages(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rest.GetLanguages)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	result := rr.Body.String()
	var Languages []entities.Language
	err = json.Unmarshal([]byte(result), &Languages)
	if err != nil {
		t.Errorf("handler returned unexpected body type: don't get got []Language")
	}

	if !reflect.DeepEqual(Languages, entities.RandomLanguages) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			Languages, entities.RandomLanguages)
	}
}

func TestGetLanguageById(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"code": "go",
	})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rest.GetLanguageById)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	result := rr.Body.String()
	var language entities.Language
	err = json.Unmarshal([]byte(result), &language)
	if err != nil {
		t.Errorf("handler returned unexpected body type: don't get got []Language")
	}

	if !reflect.DeepEqual(language, entities.RandomLanguages[0]) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			language, entities.RandomLanguages[0])
	}
}

func TestGetLanguageByIdWithUndefinedId(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"code": "notExistingLanguageCode",
	})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rest.GetLanguageById)
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

func TestPostLanguage(t *testing.T) {
	newLanguage := entities.NewLanguage("newLanguageCode", "newLanguageName")
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(newLanguage.Marshal()))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rest.PostLanguage)
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

func TestPostLanguageWithWrongBodyContent(t *testing.T) {
	newLanguage := `{
        "code": "newLanguageCode",
        name: newLanguageName
    }`

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(newLanguage)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rest.PostLanguage)
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

func TestPutLanguage(t *testing.T) {
	languageToUpdate := entities.NewLanguage("go", "updatedLanguageName")
	req, err := http.NewRequest("PUT", "/", bytes.NewBuffer(languageToUpdate.Marshal()))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rest.PutLanguage)
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

func TestPutLanguageWithWrongBodyContent(t *testing.T) {
	languageToUpdate := `{
        "code": "go",
        name: updatedLanguageName
    }`

	req, err := http.NewRequest("PUT", "/", bytes.NewBuffer([]byte(languageToUpdate)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rest.PutLanguage)
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

func TestPutLanguageWhileLanguageIsNotExisting(t *testing.T) {
	languageToUpdate := `{
        "code": "notExistingLanguageCode",
        "name": "updatedLanguageName"
    }`

	req, err := http.NewRequest("PUT", "/", bytes.NewBuffer([]byte(languageToUpdate)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rest.PutLanguage)
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

func TestDeleteLanguageById(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"code": "go",
	})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rest.DeleteLanguageById)
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

func TestDeleteLanguageByIdWithUndefinedId(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"code": "notExistingLanguageCode",
	})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rest.DeleteLanguageById)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	result := rr.Body.String()

	if result != resources.UnsuccessfulDeletionJson {
		t.Errorf("handler returned unexpected body: got %v want %v",
			result, resources.UnsuccessfulDeletionJson)
	}
}
