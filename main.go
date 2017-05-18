package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprint(w, "To get in touch, please send an email "+
    "to <a href=\"mailto:support@usegolang.com\">"+
    "support@usegolang.com</a>.")
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
  r := mux.NewRouter()
  r.HandleFunc("/", home)
  r.HandleFunc("/contact", contact)
  r.HandleFunc("/faq", faq)
  r.NotFoundHandler = http.HandlerFunc(notFound)
  http.ListenAndServe(":3000", r)
}