package categorycontroller

import (
	"go-web/entities"
	"go-web/models/categorymodel"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/category/index.html")
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
	temp, err := template.ParseFiles("views/category/create.html")
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
	var category entities.Category

	category.Name = c.PostForm("name")

	ok := categorymodel.Create(category)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan kategori"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/categories")
}

func EditGet(c *gin.Context) {
	idString := c.Query("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	category := categorymodel.Edit(id)
	data := map[string]any{
		"category": category,
	}

	temp, err := template.ParseFiles("views/category/edit.html")
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
	var category entities.Category

	idString := c.PostForm("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	category.Name = c.PostForm("name")
	category.UpdatedAt = time.Now()

	if ok := categorymodel.Update(id, category); !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate kategori"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/categories")
}

func Delete(c *gin.Context) {
	idString := c.Query("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := categorymodel.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusSeeOther, "/categories")
}
