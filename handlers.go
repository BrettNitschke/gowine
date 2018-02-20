package main

import (
  "html/template"
  "net/http"
)
func (env *Env) winesRoute(w http.ResponseWriter, r *http.Request){
  if r.Method != "GET" {
    http.Error(w, http.StatusText(405), 405)
    return
  }

  wines, err := env.db.GetWines()
  checkError(err, w)

  temp := template.Must(template.ParseFiles("templates/index.html"))

  checkError(err, w)

  err = temp.Execute(w, wines)
  checkError(err, w)

}

func (env *Env) addWineRoute(w http.ResponseWriter, r *http.Request){
  if r.Method !="POST" {
    http.Error(w, http.StatusText(405), 405)
    return
  }

  err := r.ParseForm()
  checkError(err, w)
  var winery = r.Form.Get("winery")
  var wineName = r.Form.Get("wineName")
  var vintage = r.Form.Get("vintage")


  env.db.AddWine(winery, wineName, vintage)

  http.Redirect(w, r, "/", http.StatusFound)
}

func checkError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
