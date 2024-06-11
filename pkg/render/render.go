package render

import (
	"fmt"
	"net/http"
	"text/template"
)


func RenderTemplaate (w http.ResponseWriter, tmpl string) {
	parseTmplate, _ :=  template.ParseFiles("./templates/" + tmpl)

	err := parseTmplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing the template : ", err)
		return
	}
}