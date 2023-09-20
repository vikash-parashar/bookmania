package render

import (
	"html/template"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func RenderTemplate(w io.Writer, tmpName string, data interface{}) error {
// 	tmp, err := template.ParseFiles("./templates/"+tmpName+".page.tmpl", "./templates/base.layout.tmpl")
// 	if err != nil {
// 		log.Printf("failed to parse template files .\n %v \n", err)
// 		return err
// 	}
// 	err = tmp.Execute(w, data)
// 	if err != nil {
// 		log.Printf("failed to execute template.\n %v \n", err)
// 		return err
// 	}
// 	return nil
// }

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
