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

var languageDAOMemory = persistence.NewLanguageDAOMemory()

// @Summary Get all languages
// @Description Get all languages
// @Tags Languages
// @Success 200 {array} entities.Language
// @Failure 404 {object} object
// @Router /languages [get]

func GetLanguages(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetLanguages")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := languageDAOMemory.GetAll()

	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println(resources.MarshalingError, err)
		fmt.Fprintf(w, resources.InternalErrorJson)
		return
	}

	fmt.Fprintf(w, "%s", j)
}

// @Summary Get one language
// @Description Get language by Id
// @Tags Languages
// @Param id path number true "Language Id"
// @Success 200 {object} entities.Language
// @Failure 404,500 {object} object
// @Router /languages/{id} [get]

func GetLanguageById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]
	fmt.Printf("GetLanguageById (%s)\n", code)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := languageDAOMemory.Get(code)

	if data.Code == "nil" {
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

// @Summary Create new language based on request body
// @Description Create new language
// @Tags Languages
// @Accept json
// @Success 200 {object} object
// @Failure 400 {object} object
// @Router /languages [post]

func PostLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PostLanguage")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var newLanguage entities.Language
	err := json.NewDecoder(r.Body).Decode(&newLanguage)

	if err != nil {
		fmt.Fprintf(w, resources.MarshalingErrorJson)
		return
	}

	hasBeenSaved := languageDAOMemory.Save(newLanguage)

	if hasBeenSaved {
		fmt.Fprintf(w, resources.SuccessfulAdditionJson)
	} else {
		fmt.Fprintf(w, resources.UnsuccessfulAdditionJson)
	}
}

// @Summary Update language based on request body
// @Description Update language
// @Tags Languages
// @Accept json
// @Success 200 {object} object
// @Failure 400 {object} object
// @Router /languages [put]

func PutLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PutLanguage")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var language entities.Language
	err := json.NewDecoder(r.Body).Decode(&language)

	if err != nil {
		fmt.Fprintf(w, resources.MarshalingErrorJson)
		return
	}

	hasBeenUpdated := languageDAOMemory.Update(language)

	if hasBeenUpdated {
		fmt.Fprintf(w, resources.SuccessfulUpdateJson)
	} else {
		fmt.Fprintf(w, resources.UnsuccessfulUpdateJson)
	}
}

// @Summary Delete one language
// @Description Delete language by Id
// @Tags Languages
// @Param id path number true "Language Id"
// @Success 200 {object} object
// @Failure 404,500 {object} object
// @Router /languages/{id} [delete]

func DeleteLanguageById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]
	fmt.Printf("DeleteLanguageById (%s)\n", code)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	hasBeenDeleted := languageDAOMemory.Delete(code)

	if hasBeenDeleted {
		fmt.Fprintf(w, resources.SuccessfulDeletionJson)
	} else {
		fmt.Fprintf(w, resources.UnsuccessfulDeletionJson)
	}
}
