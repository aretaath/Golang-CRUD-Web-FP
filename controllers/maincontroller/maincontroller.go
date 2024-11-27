package maincontroller

import (
	"net/http"
	"text/template"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/main/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}
