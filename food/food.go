package food

import (
  "encoding/json"
  "os"
  "log"
  "strconv"
  "github.com/mdubbs/et-microservice/rest"
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
/**
 * Get the places config file attributes
 * @return {PlacesConfig}   Struct of attributs from config JSON
 */
func getConfig() (*PlacesConfig, error) {
  file, _ := os.Open("config/places.json")
  decoder := json.NewDecoder(file)
  placesConfig := PlacesConfig{}
  err := decoder.Decode(&placesConfig)

  if err != nil {
    log.Printf("ERROR\t%s", err)
    panic(err)
  }
  return &placesConfig, err
}
/**
 * Retrieve restaraunts from Places API and parse into Struct
 * @return {FoodRecord}
 */   
func GetFoodRecord() (*FoodRecord, error) {
  config, err := getConfig()
  if err != nil {
    panic(err)
  }

  key := config.ApiKey
  rad := strconv.Itoa(config.Radius)
  lati := config.Latitude
  longi := config.Longitude

  content, err := rest.GetContent("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location="+lati+","+longi+"&radius="+rad+"&types=restaurant&key="+key)
  var record FoodRecord
  err = json.Unmarshal(content, &record)
  if err != nil {
    return nil, err
  }
  return &record, err
}