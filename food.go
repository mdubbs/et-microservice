package main

import (
  //"net/http"
  //"io/ioutil"
  "encoding/json"
  "os"
  "log"
)

type PlacesApiKey struct {
  Places string
}

type FoodRecord struct {
  Results []struct {
    Name string
  }
}

func getFood() (string, error) {
  x, err := getApiKey()
  if err != nil {
    panic(err)
  }
  return x, err
}

func getApiKey() (string, error) {
  file, _ := os.Open("keys.json")
  decoder := json.NewDecoder(file)
  placesApiKey := PlacesApiKey{}
  err := decoder.Decode(&placesApiKey)
  if err != nil {
    log.Printf("ERROR\t%s", err)
    panic(err)
  }
  return placesApiKey.Places, err
}