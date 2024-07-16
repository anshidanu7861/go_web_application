package render

import (
	"fmt"
	"net/http"
	"text/template"
)


func RenderTemplate (w http.ResponseWriter, tmpl string) {
	parseTmplate, error :=  template.ParseFiles("../../templates/" + tmpl)

	if error != nil {
		fmt.Println(error, "error for parsing")
	}

	fmt.Println(parseTmplate,"show pareserd templates")

	err := parseTmplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing the template : ", err)
		return
	}
}