package persistence

import (
	"encoding/json"
	"internal/entities"
)

type LanguageDAOBolt struct{}

var _ LanguageDAO = (*LanguageDAOBolt)(nil) // to check if LanguageDAOBolt implements LanguageDAO

func NewLanguageDAOBolt() *LanguageDAOBolt {
	return new(LanguageDAOBolt)
}

func (dao *LanguageDAOBolt) Get(code string) entities.Language {
	result := BoltDBInstance.Get(LanguageBucketName, code)
	language := entities.Language{}

	err := json.Unmarshal([]byte(result), &language)
	if err == nil && result != "nil" {
		return language
	}

	return entities.Language{Code: "nil"}
}

func (dao *LanguageDAOBolt) GetAll() []entities.Language {
	languages := []entities.Language{}
	results := BoltDBInstance.GetAll(LanguageBucketName)

	for _, s := range results {
		language := entities.Language{}
		err := json.Unmarshal([]byte(s), &language)
		if err == nil {
			languages = append(languages, language)
		}
	}

	return languages
}

func (dao *LanguageDAOBolt) Save(l entities.Language) bool {
	language, err := json.Marshal(l)
	if err == nil {
		return BoltDBInstance.Put(LanguageBucketName, l.Code, string(language))
	}
	return false
}

func (dao *LanguageDAOBolt) Update(l entities.Language) bool {
	return dao.Delete(l.Code) && dao.Save(l)
}

func (dao *LanguageDAOBolt) Delete(id string) bool {
	return BoltDBInstance.Delete(LanguageBucketName, id)
}
