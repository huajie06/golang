package db

import (
	"encoding/binary"
	"fmt"
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

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

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

func WriteToDB(v string) error {
	db, err = bolt.Open(db_name, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket_name))
		id, _ := b.NextSequence()
		v_id := itob(int(id))
		err := b.Put(v_id, []byte(v))
		return err
	})
	return err
}

func RetrieveFromDB(key int) (string, error) {
	var ret string

	db, err = bolt.Open(db_name, 0666, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket_name))
		ret = string(b.Get(itob(key)))
		return nil
	})

	return ret, nil
}

func RetrieveAll() {
	db, err = bolt.Open(db_name, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket_name))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("Task #%v: %s\n", btoi(k), v)
		}
		return nil
	})
}

func DelFromDB(key int) error {
	db, err = bolt.Open(db_name, 0666, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(bucket_name)).Delete(itob(key))
	})
}
