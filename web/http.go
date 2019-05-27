package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	appTmpl      template.Template
	wd           string
	staticDir    string
	templatesDir string
)

func init() {
	var err error
	wd, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	staticDir = fmt.Sprintf("%s/web/ui/dist", wd)
	templatesDir = fmt.Sprintf("%s/web/ui/resources/views", wd)

	appTmpl = view("app.html")
}

func appHandler(w http.ResponseWriter, r *http.Request) {
	_ = appTmpl.Execute(w, nil)
}

func staticHandler() http.Handler {
	return http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir)))
}

func view(path string) template.Template {
	fullPath := fmt.Sprintf("%s/%s", templatesDir, path)
	return *template.Must(template.ParseFiles(fullPath))
}
