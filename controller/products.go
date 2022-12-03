package controller

import(
  "fmt"
  "reflect"
  "strconv"
  "net/http"
  "html/template"
  "crud/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Home(res http.ResponseWriter, req *http.Request){
  products := models.Get("all", "")

  temp.ExecuteTemplate(res, "index", products)
}
func New(res http.ResponseWriter, req *http.Request){
  temp.ExecuteTemplate(res, "new", nil)
}
func Insert(res http.ResponseWriter, req *http.Request){
  if req.Method == "POST"{
    nome := req.FormValue("nome")
    descr := req.FormValue("descricao")
    price, _ := strconv.ParseFloat(req.FormValue("preco"), 64)
    quant, _ := strconv.Atoi(req.FormValue("quantidade"))

    models.Insert(nome, descr, price, quant)
  }
  http.Redirect(res, req, "/", 301)
}
func Delete(res http.ResponseWriter, req *http.Request){
  id := req.URL.Query().Get("id")
  models.Delete(id)

  http.Redirect(res, req, "/", 301)
}
func Edit(res http.ResponseWriter, req *http.Request){
  id := req.URL.Query().Get("id")
  product := models.Get("one", id)
  temp.ExecuteTemplate(res, "edit", product)
}
func Update(res http.ResponseWriter, req *http.Request){
  if req.Method == "POST"{
    id, _ := strconv.Atoi(req.FormValue("id"))
    nome := req.FormValue("nome")
    descr := req.FormValue("descricao")
    price, _ := strconv.ParseFloat(req.FormValue("preco"), 64)
    quant, _ := strconv.Atoi(req.FormValue("quantidade"))
    fmt.Println(reflect.TypeOf(nome), descr)
    models.Update(id, nome, descr, price, quant)
  }
  http.Redirect(res, req, "/", 301)
}
