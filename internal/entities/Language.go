package entities

import (
	"encoding/json"
	"fmt"
)

var RandomLanguages = []Language{
	{Code: "go", Name: "Go"},
	{Code: "js", Name: "JavaScript"},
	{Code: "ts", Name: "TypeScript"},
	{Code: "java", Name: "Java"},
}

type Language struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func NewLanguage() Language {
	return Language{}
}

func (l Language) String() string {
	return fmt.Sprintf("%s - %s", l.Code, l.Name)
}

func (l Language) Marshal() ([]byte, error) {
	j, err := json.Marshal(l)
	if err != nil {
		return nil, &json.MarshalerError{}
	}

	return j, nil
}
