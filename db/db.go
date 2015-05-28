package db

import (
  "github.com/boltdb/bolt"
  "log"
)

func InitializeDb() {
  //open or create BoltDB database
  db, err := bolt.Open("et.db", 0600, nil)
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  //create FoodHistory bucket if it doesn't already exist
  db.Update(func(tx *bolt.Tx) error {
    _, err := tx.CreateBucketIfNotExists([]byte("FoodHistory"))
    if err != nil {
        log.Fatal("create bucket: %s", err)
    }
    return nil
  })

  //create WeatherHistory bucket if it doesn't already exist
  db.Update(func(tx *bolt.Tx) error {
    _, err := tx.CreateBucketIfNotExists([]byte("WeatherHistory"))
    if err != nil {
        log.Fatal("create bucket: %s", err)
    }
    return nil
  })
}