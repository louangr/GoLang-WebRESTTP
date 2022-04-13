package persistence

import (
	"github.com/stretchr/testify/assert"
	"internal/entities"
	"testing"
)

var dao = NewStudentDAOMemory()

func TestGetAll(t *testing.T) {
	var actual []int

	for _, student := range dao.GetAll() {
		actual = append(actual, student.Id)
	}

	expected := []int{1, 2, 3}
	assert.Equal(t, expected, actual, "Get all the students")
}

func TestGet(t *testing.T) {
	expected := entities.NewStudent(1, "Joe", "Doe", 20, "js")
	actual := dao.Get(1)
	assert.Equal(t, expected, actual, "Get Student by Id")
}

func TestGetFail(t *testing.T) {
	actual := dao.Get(10)
	expectedId := -1
	assert.Equal(t, expectedId, actual.Id, "Get unknow Student by Id must return '-1' Id")
}

func TestDelete(t *testing.T) {
	assert.Equal(t, true, dao.Delete(1), "Delete Student by Id")
}

func TestDeleteFail(t *testing.T) {
	assert.Equal(t, false, dao.Delete(10), "Delete unknow Student must return false")
}

func TestSave(t *testing.T) {
	var expected = entities.NewStudent(4, "TestFirstName", "TestLastName", 30, "go")
	dao.Save(expected)
	actual := dao.Get(4)
	assert.Equal(t, expected, actual, "Add Student")
}
