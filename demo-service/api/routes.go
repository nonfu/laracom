package api

import (
    "encoding/json"
    "net/http"
)

type Route struct {
    Name string
    Method string
    Pattern string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "SayHello",
        "GET",
        "/hello",
        func(writer http.ResponseWriter, request *http.Request) {
            writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
            dict := map[string]string{
                "message": "hello world!",
            }
            data, _ := json.Marshal(dict)
            writer.Write(data)
        },
    },
}