package persistence

import "internal/entities"

type StudentDAO interface {
	Get(id int) entities.Student
	GetAll() []entities.Student
	Save(s entities.Student) bool
	Update(s entities.Student) bool
	Delete(id int) bool
}
