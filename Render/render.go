package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type Renderer struct {
	viewsLocation         string
	viewsFileExtension    string
	partialsLocation      string
	partialsFileExtension string
	persistCache          bool
	TemplateCache         map[string]*template.Template
}

func NewRenderer(viewsLocation, viewsFileExtension, partialsLocation, partialsFileExtension string, persistCache bool) *Renderer {

	rr := &Renderer{
		viewsLocation:         viewsLocation,
		viewsFileExtension:    viewsFileExtension,
		partialsLocation:      partialsLocation,
		partialsFileExtension: partialsFileExtension,
		persistCache:          persistCache,
	}
	rr.CreateTemplateCache()
	return rr
}

func (rr *Renderer) Render(w http.ResponseWriter, r *http.Request, name string, data map[string]interface{}) error {

	if !rr.persistCache {
		rr.CreateTemplateCache()
	}

	t, ok := rr.TemplateCache[name]
	if !ok {
		return errors.New("can't get template from cache")
	}

	buf := new(bytes.Buffer)

	err := t.Execute(buf, data)
	if err != nil {
		log.Fatal(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
		return err
	}
	return nil
}

func (rr *Renderer) CreateTemplateCache() {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*%s", rr.viewsLocation, rr.viewsFileExtension))
	if err != nil {
		rr.TemplateCache = myCache
		return
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			rr.TemplateCache = myCache
			return
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*%s", rr.partialsLocation, rr.partialsFileExtension))
		if err != nil {
			rr.TemplateCache = myCache
			return
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*%s", rr.partialsLocation, rr.partialsFileExtension))
			if err != nil {
				rr.TemplateCache = myCache
				return
			}
		}

		myCache[name] = ts
	}

	rr.TemplateCache = myCache
}
