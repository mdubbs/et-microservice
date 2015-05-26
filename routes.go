package et

import (
  "net/http"
)

type Route struct {
  Name        string
  Method      string
  Pattern     string
  HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
  Route{
    "Index",
    "GET",
    "/",
    Index,
  },
  Route{
    "GetWeather",
    "GET",
    "/weather/{zip}",
    GetWeather,
  },
  Route{
    "TimMood",
    "GET",
    "/tim/mood",
    Tim,
  },
  Route{
    "Food",
    "GET",
    "/food",
    GetFood,
  },
}