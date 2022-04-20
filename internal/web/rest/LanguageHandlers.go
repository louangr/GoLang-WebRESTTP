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

// swagger:operation GET /languages language GetLanguages
// ---
// summary: Return all languages
// description: If the are not languages, an empty array will be returned
// responses:
//   "200":
//     "$ref": "#/responses/languageStructArray"

func GetLanguages(w http.ResponseWriter, r *http.Request) {
	var languageDAO = *persistence.GetLanguageDAOInstance()
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

// swagger:operation GET /languages/{code} language GetLanguageById
// ---
// summary: Return a language by Id
// description: If the language is not found, a 404 status code will be returned
// parameters:
// - name: code
//   in: path
//   description: correspond to the language's Id
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/languageStruct"
//   "404":
//     "$ref": "#/responses/genericResponse"

func GetLanguageById(w http.ResponseWriter, r *http.Request) {
	var languageDAO = *persistence.GetLanguageDAOInstance()
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

// swagger:operation POST /languages language PostLanguage
// ---
// summary: Create a new language
// description: If the request body format is not correct, a 400 status code will be returned
// responses:
//   "200":
//     "$ref": "#/responses/genericResponse"
//   "400":
//     "$ref": "#/responses/genericResponse"

func PostLanguage(w http.ResponseWriter, r *http.Request) {
	var languageDAO = *persistence.GetLanguageDAOInstance()
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

// swagger:operation PUT /languages language PutLanguage
// ---
// summary: Update an existing language
// description: If the request body format is not correct or the target language Id is not found, a 400 status code will be returned
// responses:
//   "200":
//     "$ref": "#/responses/genericResponse"
//   "400":
//     "$ref": "#/responses/genericResponse"

func PutLanguage(w http.ResponseWriter, r *http.Request) {
	var languageDAO = *persistence.GetLanguageDAOInstance()
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

// swagger:operation DELETE /languages/{code} language DeleteLanguageById
// ---
// summary: Delete a language by Id
// description: If the language is not found, a 404 status code will be returned
// parameters:
// - name: code
//   in: path
//   description: correspond to the language's Id
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/genericResponse"
//   "404":
//     "$ref": "#/responses/genericResponse"

func DeleteLanguageById(w http.ResponseWriter, r *http.Request) {
	var languageDAO = *persistence.GetLanguageDAOInstance()
	vars := mux.Vars(r)
	code := vars["code"]
	fmt.Printf("DeleteLanguageById (%s)\n", code)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	hasBeenDeleted := languageDAO.Delete(code)

	if hasBeenDeleted {
		fmt.Fprintf(w, resources.SuccessfulDeletionJson)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, resources.NotFoundResourceJson)
	}
}
