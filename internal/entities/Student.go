package entities

import (
	"encoding/json"
	"fmt"
)

var RandomStudents = []Student{
	{Id: 1, Firstname: "Joe", Lastname: "Doe", Age: 20, LanguageCode: "js"},
	{Id: 2, Firstname: "Bob", Lastname: "Doe", Age: 21, LanguageCode: "ts"},
	{Id: 3, Firstname: "Jahn", Lastname: "Doe", Age: 21, LanguageCode: "go"},
}

type Student struct {
	Id           int    `json:"id"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Age          int    `json:"age"`
	LanguageCode string `json:"languageCode"`
}

func NewStudent(id int, firstname, lastname string, age int, languageCode string) Student {
	return Student{id, firstname, lastname, age, languageCode}
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
