package main

import (
  //"fmt"
  "net/http"
  "crud/routes"
)

func main(){
  routes.Routes()
  http.ListenAndServe(":7070", nil)
}
