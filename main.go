package main

import (
  "gowine/models"
  "net/http"
  "log"
)

type Env struct {
  db models.Datastore
}


func main(){
  db, err := models.MakeDB(CONNECTION_STRING)
  if err != nil {
    log.Panic(err)
  }

  env := &Env{db}
  http.HandleFunc("/", env.winesRoute)
  http.HandleFunc("/addWine", env.addWineRoute)
  http.ListenAndServe(":3000", nil)

}
