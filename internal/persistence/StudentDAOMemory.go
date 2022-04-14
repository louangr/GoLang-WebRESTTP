package persistence

import "internal/entities"

type StudentDAOMemory struct{}

var _ StudentDAO = (*StudentDAOMemory)(nil) // to check if StudentDAOMemory implements StudentDAO

func NewStudentDAOMemory() *StudentDAOMemory {
	return new(StudentDAOMemory)
}

var students = entities.RandomStudents

func (dao *StudentDAOMemory) Get(id int) entities.Student {
	for _, v := range students {
		if v.Id == id {
			return v
		}
	}
	return entities.Student{Id: -1}
}

func (dao *StudentDAOMemory) GetAll() []entities.Student {
	return students
}

func (dao *StudentDAOMemory) Save(s entities.Student) bool {
	students = append(students, s)
	return true
}

func (dao *StudentDAOMemory) Update(s entities.Student) bool {
	for i, v := range students {
		if v.Id == s.Id {
			students[i].Firstname = s.Firstname
			students[i].Lastname = s.Lastname
			students[i].Age = s.Age
			students[i].LanguageCode = s.LanguageCode
			return true
		}
	}
	return false
}

func (dao *StudentDAOMemory) Delete(id int) bool {
	for i, v := range students {
		if v.Id == id {
			students = append(students[:i], students[i+1:]...)
			return true
		}
	}
	return false
}
