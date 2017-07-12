package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "usegolang.com/views"
  "usegolang.com/controllers"
)

var homeView *views.View
var contactView *views.View
var faqView *views.View
var usersC *controllers.Users

func home(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  homeView.Render(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  contactView.Render(w, nil)
}

func faq(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  faqView.Render(w, nil)
}

func notFound(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprint(w, "<p>"+
    "<h2>"+
    "404</h2>")
}

func main() {
  homeView = views.NewView("bootstrap", "views/home.gohtml")
  contactView = views.NewView("bootstrap", "views/contact.gohtml")
  faqView = views.NewView("bootstrap", "views/FAQ.gohtml")
  usersC = controllers.NewUsers()

  r := mux.NewRouter()
  r.HandleFunc("/", home)
  r.HandleFunc("/contact", contact).Methods("GET")
  r.HandleFunc("/faq", faq).Methods("GET")
  r.HandleFunc("/signup", usersC.New).Methods("GET")
  r.HandleFunc("/signup", usersC.Create).Methods("POST")
  r.NotFoundHandler = http.HandlerFunc(notFound)
  http.ListenAndServe(":3000", r)
}