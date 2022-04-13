package persistence

import (
	"internal/config"
	"sync"
)

var languageDAOLock = &sync.Mutex{}
var languageDAOInstance *LanguageDAO

func GetLanguageDAOInstance() *LanguageDAO {
	if languageDAOInstance == nil {
		languageDAOLock.Lock()
		defer languageDAOLock.Unlock()
		if languageDAOInstance == nil {
			new := getLanguageDAO()
			languageDAOInstance = &new
		}
	}

	return languageDAOInstance
}

func getLanguageDAO() LanguageDAO {
	if config.IsMemoryDAONecessary {
		return NewLanguageDAOMemory()
	}
	return NewLanguageDAOBolt()
}

var studentDAOLock = &sync.Mutex{}
var studentDAOInstance *StudentDAO

func GetStudentDAOInstance() *StudentDAO {
	if studentDAOInstance == nil {
		studentDAOLock.Lock()
		defer studentDAOLock.Unlock()
		if studentDAOInstance == nil {
			new := getStudentDAO()
			studentDAOInstance = &new
		}
	}

	return studentDAOInstance
}

func getStudentDAO() StudentDAO {
	if config.IsMemoryDAONecessary {
		return NewStudentDAOMemory()
	}
	return NewStudentDAOBolt()
}
