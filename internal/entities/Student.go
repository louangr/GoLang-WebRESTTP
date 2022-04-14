package entities

import (
	"encoding/json"
	"fmt"
)

var RandomStudents = []Student{
	{Id: 1, Firstname: "Joe", Lastname: "Doe", Age: 20, LanguageCode: "go"},
	{Id: 2, Firstname: "Bob", Lastname: "Doe", Age: 21, LanguageCode: "js"},
	{Id: 3, Firstname: "Jahn", Lastname: "Doe", Age: 21, LanguageCode: "ts"},
}

type Student struct {
	Id           int    `json:"id"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Age          int    `json:"age"`
	LanguageCode string `json:"languageCode"`
}

func NewStudent(id, age int, firstname, lastname, languageCode string) Student {
	return Student{id, firstname, lastname, age, languageCode}
}

func (s Student) String() string {
	return fmt.Sprintf("%d - %s %s", s.Id, s.Firstname, s.Lastname)
}

func (s Student) Marshal() []byte {
	j, _ := json.Marshal(s)
	return j
}
