package maincontroller

import (
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {

	temp, err := template.ParseFiles("views/main/index.html")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"message": "Failed to load template"})
		return
	}

	err = temp.Execute(c.Writer, nil)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"message": "Failed to execute template"})
		return
	}
}
