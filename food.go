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

func getFood() (*FoodRecord, error) {
  x, err := getApiKey()
  if err != nil {
    panic(err)
  }
  content, err := getContent("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=42.742135,-84.570538&radius=3219&types=restaurant&key="+x)
  var record FoodRecord
  err = json.Unmarshal(content, &record)
  if err != nil {
    return nil, err
  }
  return &record, err
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