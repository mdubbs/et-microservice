package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "github.com/mdubbs/et-microservice/logger"
)

func NewRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler

        handler = route.HandlerFunc
        handler = logger.Logger(handler, route.Name)

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }

    return router
}