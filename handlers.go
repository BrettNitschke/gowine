package main

import (
  "html/template"
  "net/http"
  "fmt"
  "github.com/mssola/user_agent"
  "gopkg.in/gomail.v2"
  "crypto/tls"
)
func (env *Env) winesRoute(w http.ResponseWriter, r *http.Request){
  if r.Method != "GET" {
    http.Error(w, http.StatusText(405), 405)
    return
  }
  fmt.Println("load home page")
  ua := user_agent.New(r.UserAgent())
  name, version := ua.Browser()
  isMobile := ua.Mobile()
    fmt.Printf("%v\n", name)
    fmt.Printf("%v\n", version)

    if isMobile {
      fmt.Println("it is a mobile device")
    } else {
      fmt.Println("it is NOT a mobile device")
    }
  //THIS IS A TEST MESSAGE
  //so is this
  m := gomail.NewMessage()
  m.SetHeader("From", "alex@example.com")
  m.SetHeader("To", "brett.r.nitschke@gmail.com")
  //m.SetAddressHeader("Cc", "dan@example.com", "Dan")
  m.SetHeader("Subject", "Hello!")
  //m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
  //m.Attach("/home/Alex/lolcat.jpg")
  body := "test message"
  body += "\n\nBrowser: " + name
  body += "\n\nVersion: " + version
  if isMobile {
    body += "\n\nDevice Type: Mobile"
  } else {
    body += "\n\nDevice Type: Desktop"
  }
  m.SetBody("text/plain", body)
  d := gomail.NewDialer("smtp.gmail.com", 587, "brett.nitschke85@gmail.com", "Charlie17")
  d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

// Send the email to Bob, Cora and Dan.
if err := d.DialAndSend(m); err != nil {
    panic(err)
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
