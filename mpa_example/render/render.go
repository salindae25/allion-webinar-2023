package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderPage(w http.ResponseWriter, pageName string, data interface{}) {
	pageFilePath := fmt.Sprintf("./templates/%s.page.tmpl", pageName)

	page, err := template.ParseFiles("./templates/base.layout.tmpl", pageFilePath)
	catch(w, err)

	err = page.Execute(w, data)
	catch(w, err)
}

func catch(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
