package render

import (
	"html/template"
	"io"
	"log"
)

func RenderTemplate(w io.Writer, tmpName string, data interface{}) error {
	tmp, err := template.ParseFiles("./templates/"+tmpName+".page.tmpl", "./templates/base.layout.tmpl")
	if err != nil {
		log.Printf("failed to parse template files .\n %v \n", err)
		return err
	}
	err = tmp.Execute(w, data)
	if err != nil {
		log.Printf("failed to execute template.\n %v \n", err)
		return err
	}
	return nil
}
