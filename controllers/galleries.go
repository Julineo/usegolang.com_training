package controllers

import (
	"net/http"
	"fmt"
	
	"usegolang.com/views"
)

func NewGalleries() *Galleries {
  return &Galleries{
    NewView: views.NewView("bootstrap", "users/new"),
  }
}

type Galleries struct {
  NewView *views.View
}

// GET
func (u *Galleries) New(w http.ResponseWriter, r *http.Request) {
  //u.NewView.Render(w, nil)
  fmt.Fprintln(w, "Galleriess")
}