package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "usegolang.com/controllers"
)

func notFound(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprint(w, "<p>"+
    "<h2>"+
    "404</h2>")
}

func main() {
  staticC := controllers.NewStatic()
  usersC := controllers.NewUsers()
  galleriesC := controllers.NewGalleries()

  r := mux.NewRouter()
  r.Handle("/", staticC.Home).Methods("GET")
  r.Handle("/contact", staticC.Contact).Methods("GET")
  r.Handle("/faq", staticC.FAQ).Methods("GET")
  r.HandleFunc("/galleries", galleriesC.New).Methods("GET")
  r.HandleFunc("/signup", usersC.New).Methods("GET")
  r.HandleFunc("/signup", usersC.Create).Methods("POST")
  r.NotFoundHandler = http.HandlerFunc(notFound)
  http.ListenAndServe(":3000", r)
}