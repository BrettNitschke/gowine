package models

import (
  "database/sql"
  _ "github.com/lib/pq"
  "fmt"
)

type Datastore interface {
  GetWines()([]Wine, error)
  AddWine(winery string, wineName string, vintage string)
}

type DB struct {
  *sql.DB
}

func MakeDB(connectionString string)(*DB, error){
  
  db, err := sql.Open("postgres", connectionString)

  if err != nil{
    return nil, err
  }

  err = db.Ping()
  if err != nil {
    return nil, err
  }

  fmt.Println("connection weee!")
  return &DB{db}, nil
}
