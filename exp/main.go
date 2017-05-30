package main

import (
  "html/template"
  "os"
)

func main() {
  t, err := template.ParseFiles("hello.gohtml")
  if err != nil {
    panic(err)
  }

  m := make(map[string]int)
  m["Sting"] = 56
  m["Space"] = 83
  
  type rectangle struct {
	x1, y1, x2, y2 float64
  }
  
  r := rectangle{0, 14.5, 54.23, 43}
  
  data := struct {
    Name string
	Dog string
	Age int
	Mp map[string]int
	Rc rectangle
  }{"<script>alert('Howdy!');</script>", "Dog value", 15, m, r}

  err = t.Execute(os.Stdout, data)
  if err != nil {
    panic(err)
  }
}