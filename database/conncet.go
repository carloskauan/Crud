package database

import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func ConnectDb() *sql.DB{
  db, err := sql.Open("mysql", "root:1597@tcp(127.0.0.1:3306)/cadastro")
  if err == nil{
    fmt.Println("CONEXÃO ESTABELICIDA COM SUCESSO")
  }else{
    panic(err)
  }
  return db
}
