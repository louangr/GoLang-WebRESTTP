package entities

import "fmt"

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
