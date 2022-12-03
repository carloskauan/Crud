package models

import (
  "fmt"
  "crud/database"
  "database/sql"
)

type Product struct{
  Id int
  Nome, Desc string
  Price float64
  Quant int
}

var db = database.ConnectDb()

func Get(typeD string, id string)[]*Product{
  var res *sql.Rows
  var err error
  products := []*Product{}
  switch typeD {
    case "one":
      res, err = db.Query(fmt.Sprintf("SELECT * FROM products WHERE id=%s", id))
    case "all":
      res, err = db.Query("SELECT * FROM products ORDER BY nome asc")
  }
  if err != nil{
    fmt.Println("Erro ao ao realizar query\n", err)
    return nil
  }
  for res.Next(){
    var product Product
    if err := res.Scan(&product.Id, &product.Nome, &product.Desc, &product.Price, &product.Quant); err != nil{
      fmt.Println("Erro no scan\n",err)
      return nil
    }
    products = append(products, &product)
  }
  return products
}
func Insert(name string, descr string, price float64, quant int) {
  _, err := db.Exec(fmt.Sprintf("INSERT INTO products (nome, descr, price, quant) VALUES ('%s', '%s', %f, %d)", name, descr, price, quant))
  if err != nil{
    fmt.Println("Erro ao inserir dados na tabela\n",err)
  }
}
func Delete(id string){
  _, err := db.Exec(fmt.Sprintf("DELETE FROM products WHERE id=%s", id))
  if err != nil{
    fmt.Println("Erro ao deletar produto",err)
  }
}
func Update(id int, nome string, descr string, price float64, quant int){
  fmt.Println(id, nome, descr, price, quant)
  res, err := db.Prepare("update products set nome=?, descr=?, price=?, quant=? where id=?")
  if err != nil{
    fmt.Println("Erro ao ataualizar produtos", err)
  }
  _, err = res.Exec(nome, descr, price, quant, id)
  if err != nil{
    fmt.Println("Erro ao ataualizar produtos", err)
  }
}
