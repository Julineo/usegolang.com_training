package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "usegolang.com/views"
)

var homeView *views.View
var contactView *views.View
var faqView *views.View
var signupView *views.View

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

func signup(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  signupView.Render(w, nil)
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
  signupView = views.NewView("bootstrap", "views/signup.gohtml")

  r := mux.NewRouter()
  r.HandleFunc("/", home)
  r.HandleFunc("/contact", contact)
  r.HandleFunc("/faq", faq)
  r.HandleFunc("/signup", signup)
  r.NotFoundHandler = http.HandlerFunc(notFound)
  http.ListenAndServe(":3000", r)
}