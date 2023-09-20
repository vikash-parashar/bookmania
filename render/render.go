package render

import (
	"html/template"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderTemplate(w io.Writer, tmpl string, data interface{}) error {
	temp, err := template.ParseFiles("./templates/"+tmpl+".page.tmpl", "./templates/base.layout.tmpl")
	if err != nil {
		return err
	}
	err = temp.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}
func RenderSuccessPage(c *gin.Context) {
	c.HTML(http.StatusOK, "success.reg.html", nil)
	c.HTML(http.StatusOK, "success.log.html", nil)
	// You can render multiple success pages here if needed
}

func RenderErrorPage(c *gin.Context) {
	c.HTML(http.StatusOK, "error.reg.html", nil)
	c.HTML(http.StatusOK, "error.log.html", nil)
	// You can render multiple error pages here if needed
}
