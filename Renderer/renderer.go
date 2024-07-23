package renderer

import (
	"html/template"
	"net/http"
)

type Renderer struct {
	viewsLocation    string
	partialsLocation string
	persistCache     bool
	TemplateCache    map[string]*template.Template
}

func NewRenderer(viewsLocation, partialsLocation string, persistCache bool) *Renderer {
	return &Renderer{
		viewsLocation:    viewsLocation,
		partialsLocation: partialsLocation,
		persistCache:     persistCache,
		TemplateCache:    make(map[string]*template.Template),
	}
}

func Render(w http.ResponseWriter, r *http.Request, name string) {

}
