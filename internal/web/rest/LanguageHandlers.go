package rest

import (
	"encoding/json"
	"fmt"
	"internal/entities"
	"internal/persistence"
	"internal/resources"
	"net/http"

	"github.com/gorilla/mux"
)

var languageDAO = persistence.NewLanguageDAOBolt()

func GetLanguages(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetLanguages")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := languageDAO.GetAll()

	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println(resources.MarshalingError, err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, resources.InternalErrorJson)
		return
	}

	fmt.Fprintf(w, "%s", j)
}

func GetLanguageById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]
	fmt.Printf("GetLanguageById (%s)\n", code)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := languageDAO.Get(code)

	if data.Code == "nil" {
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

func PostLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PostLanguage")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var newLanguage entities.Language
	err := json.NewDecoder(r.Body).Decode(&newLanguage)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, resources.MarshalingErrorJson)
		return
	}

	hasBeenSaved := languageDAO.Save(newLanguage)

	if hasBeenSaved {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, resources.SuccessfulAdditionJson)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, resources.UnsuccessfulAdditionJson)
	}
}

func PutLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PutLanguage")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var language entities.Language
	err := json.NewDecoder(r.Body).Decode(&language)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, resources.MarshalingErrorJson)
		return
	}

	hasBeenUpdated := languageDAO.Update(language)

	if hasBeenUpdated {
		fmt.Fprintf(w, resources.SuccessfulUpdateJson)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, resources.UnsuccessfulUpdateJson)
	}
}

func DeleteLanguageById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]
	fmt.Printf("DeleteLanguageById (%s)\n", code)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	hasBeenDeleted := languageDAO.Delete(code)

	if hasBeenDeleted {
		fmt.Fprintf(w, resources.SuccessfulDeletionJson)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, resources.UnsuccessfulDeletionJson)
	}
}
