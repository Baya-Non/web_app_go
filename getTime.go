package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type PageVariables struct {
	Date string
	Time string
}

// http://localhost:8080/
func main() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	now := time.Now()
	const layout = "2006-01-02"
	const layout2 = "15:04:05"
	HomePageVars := PageVariables{
		Date: now.Format(layout),
		Time: now.Format(layout2),
	}

	t, err := template.ParseFiles("homepage.html")
	if err != nil {
		log.Print("template parsing error: ", err) //log
	}
	err = t.Execute(w, HomePageVars)
	if err != nil {
		log.Print("template executing error: ", err) //log
	}
}
