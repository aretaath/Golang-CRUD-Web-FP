package usercontroller

import (
	"go-web/entities"
	"go-web/models/usermodel"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	users := usermodel.GetAll()
	data := map[string]any{
		"users": users,
	}

	temp, err := template.ParseFiles("views/user/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/user/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var user entities.User

		user.Name = r.FormValue("name")
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

		ok := usermodel.Create(user)
		if !ok {
			temp, _ := template.ParseFiles("views/user/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/users", http.StatusSeeOther)
	}

}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/user/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		user := usermodel.Edit(id)
		data := map[string]any{
			"user": user,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var user entities.User

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		user.Name = r.FormValue("name")
		user.UpdatedAt = time.Now()

		if ok := usermodel.Update(id, user); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/users", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := usermodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
