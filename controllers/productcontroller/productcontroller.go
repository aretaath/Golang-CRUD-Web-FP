package productcontroller

import (
	"go-web/entities"
	"go-web/models/categorymodel"
	"go-web/models/productmodel"
	"go-web/models/usermodel"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	products := productmodel.Getall()
	data := map[string]any{
		"products": products,
	}

	temp, err := template.ParseFiles("views/product/index.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := temp.Execute(c.Writer, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func AddGet(c *gin.Context) {
	categories := categorymodel.GetAll()
	users := usermodel.GetAll()
	data := map[string]any{
		"categories": categories,
		"users":      users,
	}

	temp, err := template.ParseFiles("views/product/create.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Render the template
	if err := temp.Execute(c.Writer, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func AddPost(c *gin.Context) {
	var product entities.Product

	categoryId, err := strconv.Atoi(c.PostForm("category_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userId, err := strconv.Atoi(c.PostForm("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	quantity, err := strconv.Atoi(c.PostForm("quantity"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	product.Name = c.PostForm("name")
	product.Category.Id = uint(categoryId)
	product.User.Id = uint(userId)
	product.Quantity = int64(quantity)
	product.Description = c.PostForm("description")

	if ok := productmodel.Create(product); !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan produk"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/products")
}

func Detail(c *gin.Context) {
	idString := c.Query("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	product := productmodel.Detail(id)
	data := map[string]any{
		"product": product,
	}

	temp, err := template.ParseFiles("views/product/detail.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Render the template
	if err := temp.Execute(c.Writer, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func EditGet(c *gin.Context) {
	idString := c.Query("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	product := productmodel.Detail(id)
	categories := categorymodel.GetAll()

	data := map[string]any{
		"product":    product,
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/product/edit.html")
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
	var product entities.Product

	idString := c.PostForm("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	categoryId, err := strconv.Atoi(c.PostForm("category_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	quantity, err := strconv.Atoi(c.PostForm("quantity"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	product.Name = c.PostForm("name")
	product.Category.Id = uint(categoryId)
	product.Quantity = int64(quantity)
	product.Description = c.PostForm("description")
	product.UpdatedAt = time.Now()

	if ok, err := productmodel.Update(id, product); err != nil || !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Poduct Update Fail: " + err.Error()})
		return
	}

	c.Redirect(http.StatusSeeOther, "/products")
}

func Delete(c *gin.Context) {
	idString := c.Query("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := productmodel.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusSeeOther, "/products")
}
