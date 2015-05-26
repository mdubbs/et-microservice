package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
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
  content, err := getContent(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?zip=%s,us", zip))
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

func getContent(url string) ([]byte, error) {
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return nil, err
  }

  //send request
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return nil, err
  }

  //defer closing
  defer resp.Body.Close()
  //read content
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  return body, nil
}