package routes

import (
  "net/http"
  "crud/controller"
)

func Routes(){
  http.HandleFunc("/", controller.Home)
  http.HandleFunc("/new", controller.New)
  http.HandleFunc("/insert", controller.Insert)
  http.HandleFunc("/delete", controller.Delete)
  http.HandleFunc("/edit", controller.Edit)
  http.HandleFunc("/update", controller.Update)
}
