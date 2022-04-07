package entities

import "fmt"

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
