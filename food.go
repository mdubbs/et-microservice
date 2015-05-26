package main

import (
  "encoding/json"
  "os"
  "log"
  "strconv"
)

type PlacesConfig struct {
  ApiKey string
  Radius int
  Latitude string
  Longitude string
}

type FoodRecord struct {
  Results []struct {
    Name string
  }
}

func getFood() (*FoodRecord, error) {
  config, err := getConfig()
  if err != nil {
    panic(err)
  }

  key := config.ApiKey
  rad := strconv.Itoa(config.Radius)
  lati := config.Latitude
  longi := config.Longitude

  content, err := getContent("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location="+lati+","+longi+"&radius="+rad+"&types=restaurant&key="+key)
  var record FoodRecord
  err = json.Unmarshal(content, &record)
  if err != nil {
    return nil, err
  }
  return &record, err
}

func getConfig() (*PlacesConfig, error) {
  file, _ := os.Open("places.json")
  decoder := json.NewDecoder(file)
  placesConfig := PlacesConfig{}
  err := decoder.Decode(&placesConfig)

  if err != nil {
    log.Printf("ERROR\t%s", err)
    panic(err)
  }
  return &placesConfig, err
}