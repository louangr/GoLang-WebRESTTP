package persistence

import "internal/entities"

type LanguageDAOMemory struct{}

var _ LanguageDAO = (*LanguageDAOMemory)(nil) // to check if LanguageDAOMemory implements LanguageDAO

func NewLanguageDAOMemory() *LanguageDAOMemory {
	return new(LanguageDAOMemory)
}

var languages = []entities.Language{
	{Code: "fra", Name: "French"},
	{Code: "eng", Name: "English"},
	{Code: "spa", Name: "Spanish"},
	{Code: "fin", Name: "Finnish"},
}

func (dao *LanguageDAOMemory) Get(code string) entities.Language {
	for _, v := range languages {
		if v.Code == code {
			return v
		}
	}
	return entities.Language{Code: "nil"}
}

func (dao *LanguageDAOMemory) GetAll() []entities.Language {
	return languages
}

func (dao *LanguageDAOMemory) Save(l entities.Language) bool {
	languages = append(languages, l)
	return true
}

func (dao *LanguageDAOMemory) Update(l entities.Language) bool {
	for i, v := range languages {
		if v.Code == l.Code {
			languages[i].Name = l.Name
			return true
		}
	}
	return false
}

func (dao *LanguageDAOMemory) Delete(code string) bool {
	for i, v := range languages {
		if v.Code == code {
			languages = append(languages[:i], languages[i+1:]...)
			return true
		}
	}
	return false
}
