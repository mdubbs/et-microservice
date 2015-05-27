package weather

import (
  "fmt"
  "encoding/json"
  "github.com/mdubbs/et-microservice/rest"
)

type WeatherRecord struct {
  Weather []struct {
    Id int `json:"id"`
    Main string `json:"main"`
    Icon string `json:"icon"`
    Weather string `json:"description"`
  } `json:"weather"`
  Main struct {
    Temp float64 `json:"temp"`
    Humidity int `json:"humidity"`
    Pressure float64 `json:"pressure"`
    TempHigh float64 `json:"temp_max"`
    TempLow float64 `json:"temp_min"`
  } `json:"main"`
  Name string `json:"name"`
}

func GetWeatherRecord(zip string) (*WeatherRecord, error) {
  content, err := rest.GetContent(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?zip=%s,us", zip))
  if err != nil {
    return nil, err
  }

  var record WeatherRecord
  err = json.Unmarshal(content, &record)
  if err != nil {
    return nil, err
  }
  return &record, err
}