package persistence

import (
	"encoding/json"
	"fmt"
	"internal/entities"
)

type StudentDAOBolt struct{}

var _ StudentDAO = (*StudentDAOBolt)(nil) // to check if StudentDAOBolt implements StudentDAO

func NewStudentDAOBolt() StudentDAOBolt {
	return StudentDAOBolt{}
}

func (dao *StudentDAOBolt) Get(id int) entities.Student {
	result := BoltDBInstance.Get(StudentBucketName, fmt.Sprintf("%d", id))
	student := entities.Student{}

	err := json.Unmarshal([]byte(result), &student)
	if err == nil && result != "nil" {
		return student
	}

	return entities.Student{Id: -1}
}

func (dao *StudentDAOBolt) GetAll() []entities.Student {
	students := []entities.Student{}
	results := BoltDBInstance.GetAll(StudentBucketName)

	for _, s := range results {
		student := entities.Student{}
		err := json.Unmarshal([]byte(s), &student)
		if err == nil {
			students = append(students, student)
		}
	}

	return students
}

func (dao *StudentDAOBolt) Save(s entities.Student) bool {
	student, err := json.Marshal(s)
	if err == nil {
		return BoltDBInstance.Put(StudentBucketName, fmt.Sprintf("%d", s.Id), string(student))
	}
	return false
}

func (dao *StudentDAOBolt) Update(s entities.Student) bool {
	return dao.Delete(s.Id) && dao.Save(s)
}

func (dao *StudentDAOBolt) Delete(id int) bool {
	return BoltDBInstance.Delete(StudentBucketName, fmt.Sprintf("%d", id))
}
