package db

import (
	"log"
	"time"

	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB
var err error

const (
	bucket_name string = "taskBucket"
	db_name     string = "task.db"
)

func Init() error {
	db, err = bolt.Open(db_name, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket_name))
		return err
	})
}

func WriteToDB(k, v string) error {
	db, err := bolt.Open(db_name, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket_name))
		err := b.Put([]byte(k), []byte(v))
		return err
	})
	return err
}

func RetrieveFromDB(key string) (string, error) {
	var ret string

	db, err = bolt.Open(db_name, 0666, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket_name))
		ret = string(b.Get([]byte(key)))
		return nil
	})

	return ret, nil
}

func RetrieveAll() string {
	var ret string

	db, err = bolt.Open(db_name, 0666, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket_name))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			ret += string(v) + "\n"
			// fmt.Printf("key=%s, value=%s\n", k, v)
		}
		return nil
	})
	return ret
}

func DelFromDB(key string) error {
	db, err = bolt.Open(db_name, 0666, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(bucket_name)).Delete([]byte(key))
	})
}
