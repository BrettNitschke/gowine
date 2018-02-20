package models

import (
  "fmt"
)
type Wine struct {
  Id string `json:"id"`
  Winery string `json:"winery"`
  WineName string `json:"wine name"`
  Vintage string `json:"vintage"`
}

func (db *DB) GetWines()([]Wine, error){
  rows, err := db.Query("SELECT * FROM wines")
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  var wines []Wine
  var wine Wine
  for rows.Next(){
    err = rows.Scan(&wine.Id, &wine.Winery, &wine.WineName, &wine.Vintage)
    if err != nil{
      return nil, err
    }
    wines = append(wines, wine)
  }

  if err = rows.Err(); err != nil {
    return nil, err
  }
  
  return wines, nil
}

func (db *DB) AddWine(winery string, wineName string, vintage string){
  var wine Wine
  wine.Winery = winery
  wine.WineName = wineName
  wine.Vintage = vintage
  fmt.Println(wine)

  _, err := db.Exec("INSERT INTO wines(winery_name, wine_name, vintage) VALUES($1, $2, $3)",wine.Winery, wine.WineName, wine.Vintage)
  if err != nil {
    fmt.Println("exec error")
    panic(err)
  }

  return
}
