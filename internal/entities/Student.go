package entities

import (
	"encoding/json"
	"fmt"
)

var RandomStudents = []Student{
	{Id: 1, Firstname: "Joe", Lastname: "Doe", Age: 20, LanguageCode: "fra"},
	{Id: 2, Firstname: "Bob", Lastname: "Doe", Age: 21, LanguageCode: "fra"},
	{Id: 3, Firstname: "Jahn", Lastname: "Doe", Age: 21, LanguageCode: "eng"},
}

type Student struct {
	Id           int    `json:"id"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Age          int    `json:"age"`
	LanguageCode string `json:"languageCode"`
}

func NewStudent() Student {
	return Student{}
}

func (s Student) String() string {
	return fmt.Sprintf("%d - %s %s", s.Id, s.Firstname, s.Lastname)
}

func (s Student) Marshal() ([]byte, error) {
	j, err := json.Marshal(s)
	if err != nil {
		return nil, &json.MarshalerError{}
	}

	return j, nil
}
