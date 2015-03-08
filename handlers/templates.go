package handlers

import (
	"github.com/THUNDERGROOVE/SDETool2/log"
	"html/template"
	"path/filepath"
)

var Templates map[string]*template.Template

func init() {
	if Templates == nil {
		Templates = make(map[string]*template.Template)
	}
}

func LoadTemplates() {
	Templates = make(map[string]*template.Template)

	glob := []string{"template/index.tmpl", "template/login.tmpl",
		"template/fit.tmpl", "template/newfit.tmpl",
		"template/error.tmpl", "template/usermanage.tmpl",
		"template/fits.tmpl"}
	includes := []string{"template/base.tmpl"}
	for _, v := range glob {
		var err error
		log.Info("Attempting to load template", v, "with index of", filepath.Base(v))
		files := append(includes, v)
		Templates[filepath.Base(v)] = template.Must(template.ParseFiles(files...))
		if err != nil {
			log.LogError(err.Error())
		}
	}
}
