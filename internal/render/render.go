package render

import (
	con "dshusdock/go_project/internal/constants"
	"encoding/json"
	"html/template"
	"log/slog"
	"net/http"
)

var files = []string{
	"./ui/html/pages/base.tmpl.html",
	"./ui/html/pages/layout.tmpl.html",
	"./ui/html/pages/header.tmpl.html",
	"./ui/html/pages/test/page1.tmpl.html",
	"./ui/html/pages/sidenav.tmpl.html",
	"./ui/html/pages/system-list.tmpl.html",
	"./ui/html/pages/test-modal.tmpl.html",
}

type Payload struct {
    Server string
}

func JSONResponse(w http.ResponseWriter, data string) {
	t := Payload{Server: data}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(t)
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, r *http.Request, d any) {
	tmpl := template.Must(template.ParseFiles(files...))

	tmpl.ExecuteTemplate(w, "base", d)
}

// RenderTemplate renders a template
func RenderAppTemplate(w http.ResponseWriter, r *http.Request, data any, _type int) {
	info := con.GetRenderInfo(_type)
	slog.Debug("RenderAppTemplate: ","template",  info.TemplateName)	
	template.Must(template.ParseFiles(info.TemplateFiles...)).ExecuteTemplate(w, info.TemplateName, data)
}
