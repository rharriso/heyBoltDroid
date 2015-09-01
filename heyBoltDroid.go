package heyBoltDroid

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

type DBO struct{}

func (d DBO) DoDataBaseThing() string {
	db, err := bolt.Open("bolton.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create bucket
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("MahBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		v := b.Get([]byte("buddy"))
		fmt.Printf("Buddy before: %v\n", string(v))

		err = b.Put([]byte("buddy"), []byte("guy"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		v = b.Get([]byte("buddy"))
		fmt.Printf("Buddy before: %v\n", string(v))

		return nil
	})

	return "WHATEVER"
}
