package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("task")
var db *bolt.DB //package level variable

type Task struct {
	Key   int
	Value string
}

func Init(dbPath string) error {

	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error { // to start read-write transaction.
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}
func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error { // to start read-write transaction.
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)

		key := itob(id)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}
func AllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error { // to start only read transaction.
		b := tx.Bucket(taskBucket)

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return tasks, nil
}
func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error { // to start read-write transaction.
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}
func itob(v int) []byte { //convert int to byte slices
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
func btoi(b []byte) int { //convert byte slices to int
	return int(binary.BigEndian.Uint64(b))
}
