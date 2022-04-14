package persistence

import (
	"fmt"
	"internal/entities"
	"log"

	"github.com/boltdb/bolt"
)

type BoltDB struct {
	DB *bolt.DB
}

var BoltDBInstance BoltDB

var StudentBucketName string = "Student"
var LanguageBucketName string = "Language"

func InitBoldDB(dbFileName string) {
	db, err := bolt.Open(dbFileName+".db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	BoltDBInstance = BoltDB{DB: db}

	hasToInitValue := false
	err = BoltDBInstance.DB.Update(func(tx *bolt.Tx) error {
		studentBucket := tx.Bucket([]byte(StudentBucketName))
		languageBucket := tx.Bucket([]byte(LanguageBucketName))

		if studentBucket == nil || languageBucket == nil {
			_, err := tx.CreateBucket([]byte(StudentBucketName))
			if err != nil {
				return err
			}

			_, err = tx.CreateBucket([]byte(LanguageBucketName))
			if err != nil {
				return err
			}

			hasToInitValue = true
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	if hasToInitValue {
		BoltDBInstance.fill()
	}
}

func (b *BoltDB) fill() {
	for _, s := range entities.RandomStudents {
		b.Put(StudentBucketName, fmt.Sprintf("%d", s.Id), string(s.Marshal()))
	}

	for _, l := range entities.RandomLanguages {
		b.Put(LanguageBucketName, l.Code, string(l.Marshal()))
	}
}

func (b *BoltDB) Get(bucketName string, key string) string {
	var result string

	err := b.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		result = string(bucket.Get([]byte(key)))
		return nil
	})

	if err != nil || result == "" {
		return "nil"
	}

	return result
}

func (b *BoltDB) GetAll(bucketName string) []string {
	var result []string

	err := b.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			result = append(result, string(v))
		}

		return nil
	})

	if err != nil {
		return nil
	}

	return result
}

func (b *BoltDB) Put(bucketName string, key string, value string) bool {
	element := b.Get(bucketName, key)

	if element != "nil" {
		return false
	}

	err := b.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		bucket.Put([]byte(key), []byte(value))
		return nil
	})

	if err != nil {
		return false
	}

	return true
}

func (b *BoltDB) Delete(bucketName string, key string) bool {
	element := b.Get(bucketName, key)

	if element == "nil" {
		return false
	}

	err := b.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		bucket.Delete([]byte(key))
		return nil
	})

	if err != nil {
		return false
	}

	return true
}
