package renderer

import (
	"html/template"
	"net/http"
)

type Renderer struct {
	viewsLocation         string
	viewsFileExtension    string
	partialsLocation      string
	partialsFileExtension string
	persistCache          bool
	TemplateCache         map[string]*template.Template
}

func NewRenderer(viewsLocation, viewsFileExtension, partialsFileExtension, partialsLocation string, persistCache bool) *Renderer {

	rr := &Renderer{
		viewsLocation:         viewsLocation,
		viewsFileExtension:    viewsFileExtension,
		partialsLocation:      partialsLocation,
		partialsFileExtension: partialsFileExtension,
		persistCache:          persistCache,
	}
	tc, err := rr.CreateTemplateCache()
	if err != nil {
		panic("could not CreateTemplateCache in NewRenderer")
	}
	rr.TemplateCache = tc
	return rr
}

func (rr *Renderer) Render(w http.ResponseWriter, r *http.Request, name string) error {

	if !rr.persistCache {
		tc, err := rr.CreateTemplateCache()
		if err != nil {
			return err
		}
		rr.TemplateCache = tc
	}

	return nil
}

func (rr *Renderer) CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	return cache, nil
}
