package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "<h2>404</h2>")	//<h1>404</h1> works, but <h2>404</h2> doesn't
}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/hello/:name", Hello)
	router.NotFound = http.HandlerFunc(notFound)
    log.Fatal(http.ListenAndServe(":8080", router))
}