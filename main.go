package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "usegolang.com/views"
)

var homeView *views.View
var contactView *views.View

func home(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
}

func faq(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprint(w, "FAQ "+
    "<b>"+
    "FAQ</b>")
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

  r := mux.NewRouter()
  r.HandleFunc("/", home)
  r.HandleFunc("/contact", contact)
  r.HandleFunc("/faq", faq)
  r.NotFoundHandler = http.HandlerFunc(notFound)
  http.ListenAndServe(":3000", r)
}