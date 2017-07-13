package controllers

import (
	"net/http"
	"fmt"
	
	"github.com/gorilla/schema"
	
	"usegolang.com/views"
)

func NewUsers() *Users {
  return &Users{
    NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
  }
}

type Users struct {
  NewView *views.View
}

// This is used to render the form where a user can
// create a new user account.
// 
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
  u.NewView.Render(w, nil)
}

// Create is used to process the signup form when a user
// tries to create a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
  if err := r.ParseForm(); err != nil {
    panic(err)
  }
  
  dec := schema.NewDecoder()
  form := SignupForm{}
  if err := dec.Decode(&form, r.PostForm); err != nil {
    panic(err)
  }
  fmt.Fprintln(w, form)
}

type SignupForm struct {
  Email    string `schema:"email"`
  Password string `schema:"password"`
}