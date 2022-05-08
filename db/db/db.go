package db

import (
	"errors"
	"fmt"
	"io/fs"

	"github.com/boltdb/bolt"
)

var DEFAULT_BUCKET []byte = []byte("DEFAULT")

type JKZDB struct {
	db *bolt.DB
}

func CreateJKZDB(identifier int) (*JKZDB, error) {
	dbname := fmt.Sprintf("jkzdb-%d.db", identifier)
	db, err := bolt.Open(dbname, fs.ModeAppend, nil)
	if err != nil {
		return nil, err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket(DEFAULT_BUCKET)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &JKZDB{
		db: db,
	}, nil
}

func (jkzdb *JKZDB) CreateNewDBEntry(key, value string) error {
	// TODO: This is for new creation, but we need another function for more fine grained
	// updates with pre-existing keys
	err := jkzdb.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(DEFAULT_BUCKET)
		if b == nil {
			return errors.New("Bucket doesn't exist.")
		}
		err := b.Put([]byte(key), []byte(value))
		return err
	})
	return err
}

func (jkzdb *JKZDB) GetValue(key string) (string, error) {
	value := ""
	err := jkzdb.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(DEFAULT_BUCKET)
		if b == nil {
			return errors.New("Bucket doesn't exist.")
		}
		valueBytes := b.Get([]byte(key))
		if valueBytes == nil {
			return errors.New("Key doesn't exist.")
		}
		value = string(valueBytes)
		return nil
	})
	return value, err
}

func (jkzdb *JKZDB) DeleteKey(key string) error {
	err := jkzdb.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(DEFAULT_BUCKET)
		if b == nil {
			return errors.New("Bucket doesn't exist.")
		}
		err := b.Delete([]byte(key))
		return err
	})
	return err
}

func (jkzdb *JKZDB) isUnique(key string) (bool, error) {

	// check if string is not in db yet
	// can try to call get value
	// TODO: SCAN DB FOR UNIQUE KEYS
	err := jkzdb.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(DEFAULT_BUCKET)
		if b == nil {
			return errors.New("Bucket doesn't exist.")
		}
		valueBytes := b.Get([]byte(key))
		if valueBytes == nil {
			return true, nil
		}
		return nil
	})
	return false, nil
}

func (jkzdb *JKZDB) UpdateEntry(key, value string) error {
	// TODO: This is for new creation, but we need another function for more fine grained
	// updates with pre-existing keys
	err := jkzdb.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(DEFAULT_BUCKET)
		if b == nil {
			return errors.New("Bucket doesn't exist.")
		}
		err := b.Put([]byte(key), []byte(value))
		return err
	})
	return err
}
