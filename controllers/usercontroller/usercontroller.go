package usercontroller

import (
	"go-web/entities"
	"go-web/models/usermodel"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	users := usermodel.GetAll()
	data := map[string]any{
		"users": users,
	}

	temp, err := template.ParseFiles("views/user/index.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Render the template
	if err := temp.Execute(c.Writer, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func AddGet(c *gin.Context) {
	temp, err := template.ParseFiles("views/user/create.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Render the template
	if err := temp.Execute(c.Writer, nil); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func AddPost(c *gin.Context) {
	var user entities.User

	user.Name = c.PostForm("name")

	ok := usermodel.Create(user)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan user"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/users")
}

func EditGet(c *gin.Context) {
	idString := c.Query("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := usermodel.Edit(id)
	data := map[string]any{
		"user": user,
	}

	temp, err := template.ParseFiles("views/user/edit.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Render the template
	if err := temp.Execute(c.Writer, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func EditPost(c *gin.Context) {
	var user entities.User

	idString := c.PostForm("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Name = c.PostForm("name")
	user.UpdatedAt = time.Now()

	if ok := usermodel.Update(id, user); !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate user"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/users")
}

func Delete(c *gin.Context) {
	idString := c.Query("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := usermodel.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusSeeOther, "/users")
}
