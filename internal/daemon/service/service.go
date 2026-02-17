package service

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/ulibaysya/krona/internal/config"

	"github.com/ulibaysya/krona/internal/daemon/service/handlers"
	"github.com/ulibaysya/krona/internal/log"
	"github.com/ulibaysya/krona/internal/storage"
)

type Service struct {
	mux  *chi.Mux
}

func New(cfg config.Service, log log.Logger, strg storage.Storage) (Service, error) {
	const f = "github.com/ulibaysya/krona/internal/daemon/service.New"
	//
	// const (
	// 	root       = "/"
	// 	catalogs   = "/catalogs"
	// 	catalogsID = "/catalogs/{catalogsID}"
	// )
	// templateFiles := []string{
	// 	"index.tmpl",
	// 	"shop.tmpl",
	// }
	// pagesTmpl := make(map[string]*template.Template, len(templateFiles))
	// for _, i := range templateFiles {
	// 	base := *baseTmpl
	// 	t, err := base.ParseFiles(filepath.Join(cfg.TemplatesPath, i))
	// 	if err != nil {
	// 		return Service{}, fmt.Errorf("%s: %w", f, err)
	// 	}
	// 	pagesTmpl[i] = t
	// 	fmt.Println(i, pagesTmpl)
	// }
	//
	templates, err := initTemplates(cfg.TemplatesPath)
	if err != nil {
		return Service{}, fmt.Errorf("%s: %w", f, err)
	}

	obj := []struct {
		path string
		fn   func(strg storage.Storage, tmpl *template.Template) http.HandlerFunc
	}{
		{"GET /", handlers.GetRoot},
		{"GET /catalogs", handlers.GetCatalogs},
		{"GET /catalogs/{catalogsID}", handlers.GetCatalogsID},
	}

	mux := chi.NewMux()

	if cfg.Static.Serve {
		h := http.FileServer(http.Dir(cfg.Static.Path))
		mux.Handle("/*", h)
	}

	for _, i := range obj {
		tmpl, ok := templates[i.path];
		if !ok {
			fmt.Printf("%v: path has nil template, skip\n", i.path)
			continue
		}
		mux.HandleFunc(i.path, i.fn(strg, tmpl))
	}

	return Service{mux: mux}, nil
}

func (s Service) GetMux() chi.Mux {
	return *s.mux
}

// Path: "GET /"

// Templates:
// base.tmpl
// root.tmpl
func initTemplates(dir string) (map[string]*template.Template, error) {

	if dir == "" {
		return nil, errors.New("template path isn't provided")
	}
	fmt.Println(dir)

	// baseFiles := []string{
	// 	filepath.Join(dir, "base.tmpl"),
	// 	filepath.Join(dir, "header.tmpl"),
	// 	filepath.Join(dir, "footer.tmpl"),
	// 	filepath.Join(dir, "js.tmpl"),
	// 	filepath.Join(dir, "right-area.tmpl"),
	// 	filepath.Join(dir, "index.tmpl"),
	// }
	// fmt.Println(baseFiles)
	// baseTmpl, err := template.ParseFiles(baseFiles...)

	baseTmpl, err := template.ParseGlob(filepath.Join(dir, "*"))
	if err != nil {
		return nil, err
	}
	fmt.Println("executing template...")
	baseTmpl.Execute(os.Stdout, nil)
	fmt.Println("...executing template")

	templates := make(map[string]*template.Template, 5)
	templates["GET /"] = baseTmpl

	// files := map[string]string{
	// 	"GET /":                                  "root.tmpl",
	// 	"GET /catalogs":                          "catalogs.tmpl",
	// 	"GET /catalogs/{catalogsID}":             "products.tmpl",
	// 	"GET /catalogs/{catalogsID}/{productID}": "product.tmpl",
	// }
	return templates, nil
}
