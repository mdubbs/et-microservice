package et

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "math/rand"

  "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  fmt.Fprintln(w, "Welcome!")
}

func Tim(w http.ResponseWriter, r *http.Request) {

  x := []string{
    "kicking someone's ass!",
    "crying..",
    "giving someone a nice tall glass of shut the fuck up!",
    "getting some Panda Express.",
    "a SharePoint god!",
    "you should go outside and play hide and go fuck yourself.",
  }

  w.Header().Set("Content-Type", "text/html; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  fmt.Fprintln(w, "Tim is currently feeling like: ", x[rand.Intn(len(x))])
}

func GetWeather(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  content, err := GetWeatherRecord(vars["zip"])
  if err != nil {
    log.Fatal(err)
    json.NewEncoder(w).Encode("Error! Sorry bro!")
  } else {
    //convert temp from K to F
    content.Main.Temp = (((content.Main.Temp-273.15)*1.8)+32)
    content.Main.TempHigh = (((content.Main.TempHigh-273.15)*1.8)+32)
    content.Main.TempLow = (((content.Main.TempLow-273.15)*1.8)+32)

    format := r.URL.Query().Get("format")

    if format == "json" {
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusOK)
      json.NewEncoder(w).Encode(content)
    } else {
      w.Header().Set("Content-Type", "text/html; charset=UTF-8")
      w.WriteHeader(http.StatusOK)
      var x int = int(content.Main.Temp)
      fmt.Fprintf(w, "In %s, it is %v° F and the current conditions are: %s", content.Name, x, content.Weather[0].Main)
    }
  }
}

func GetFood(w http.ResponseWriter, r *http.Request) {
  x, err := getFood()

  if err != nil {
    panic(err)
  }

  w.Header().Set("Content-Type", "text/html; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  fmt.Fprintf(w, "Test: %s", x)
}