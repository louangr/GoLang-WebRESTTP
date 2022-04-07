package persistence

import "internal/entities"

type LanguageDAO interface {
	Get(code string) entities.Language
	GetAll() []entities.Language
	Save(l entities.Language) bool
	Update(l entities.Language) bool
	Delete(code string) bool
}
