# The Emerging Tech Microservice API
The greatest of all the ET microservices

![logo](https://raw.githubusercontent.com/mdubbs/et-microservice/master/img/go.png)

* Written in Go, for some reason..
* Uses the OpenWeatherMap API as well as Google's Places API

### Using the service
Just run `go run`, or `go build` and then `./et` or `et.exe`, the service can be reached at `http://localhost:8080`

### Endpoints
```
Route                             Action
-------                           --------
/                                 Returns "Welcome!" (text/html)

/food                             Returns a restaurant within the specified radius of your
                                  location (text/html)

/tim/mood                         Returns Tim's mood

/weather/{ZIP CODE}               Returns that ZIP's weather (text/html)

/weather/{ZIP CODE}?format=json   Returns that ZIP's weather (json)
```
### Using Google Places API

To use the Food endpoint you will need to create a `places.json` file and add the following while substituting in your own Google Places API key (string), specified radius in meters (int), location latitude (string) and logitude (string).

```json
{
  "apikey":     "YOUR-PLACES-API-KEY-HERE",
  "radius":     "DISTANCE_INT",
  "latitude":   "YOUR-LOCATION-LATITUDE",
  "longitude":  "YOUR-LOCATION-LONGITUDE"
}
```
### To-Do
* Refactor (types, etc) big time!
* Add in NoSQL support for keeping track of places you like and don't like
  * Add ability to share a food pick with a link
  * Using MongoDB, that's hip right?
* Add more endpoints!
  * Add some useful endpoints? Maybe..