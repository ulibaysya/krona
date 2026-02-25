package service

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/ulibaysya/krona/internal/config"

	"github.com/ulibaysya/krona/internal/daemon/service/handlers"
	"github.com/ulibaysya/krona/internal/log"
	"github.com/ulibaysya/krona/internal/storage"
)

type Service struct {
	mux *chi.Mux
}

func New(cfg config.Service, log log.Logger, strg storage.Storage) (Service, error) {
	const f = "github.com/ulibaysya/krona/internal/daemon/service.New"

	baseTemplates := []string{
		filepath.Join(cfg.TemplatesPath, "base.tmpl"),
		filepath.Join(cfg.TemplatesPath, "head.tmpl"),
		filepath.Join(cfg.TemplatesPath, "header.tmpl"),
		filepath.Join(cfg.TemplatesPath, "footer.tmpl"),
	}
	baseTemplate, err := template.ParseFiles(baseTemplates...)
	if err != nil {
		return Service{}, fmt.Errorf("%s: %w", f, err)
	}

	mux := chi.NewMux()

	objects := []struct {
		path     string
		aliases  []string
		handler  func(strg storage.Storage, tmpl *template.Template) http.HandlerFunc
		tmplFile string
		tmpl     *template.Template
	}{
		{
			path:     "GET /",
			handler:  handlers.GetRoot,
			tmplFile: "root.tmpl",
		},
		{
			path:     "GET /catalogs",
			aliases:  []string{"GET /catalog"},
			handler:  handlers.GetCatalogs,
			tmplFile: "catalogs.tmpl",
		},
		{
			path: "GET /catalogs/{catalogsID}",
			aliases: []string{
				"GET /catalog/{catalogsID}",
			},
			handler:  handlers.GetCatalogsID,
			tmplFile: "catalog.tmpl",
		},
		{
			path: "GET /catalogs/{catalogsID}/{productID}",
			aliases: []string{
				"GET /catalog/{catalogsID}/{productID}",
			},
			handler:  handlers.GetProductID,
			tmplFile: "product.tmpl",
		},
	}

	for _, obj := range objects {
		c, err := baseTemplate.Clone()
		if err != nil {
			return Service{}, fmt.Errorf("%s: %w", f, err)
		}
		_, err = c.ParseFiles(filepath.Join(cfg.TemplatesPath, obj.tmplFile))
		if err != nil {
			return Service{}, fmt.Errorf("%s: %w", f, err)
		}

		mux.HandleFunc(obj.path, obj.handler(strg, c))
		for _, alias := range obj.aliases {
			mux.HandleFunc(alias, obj.handler(strg, c))
		}
	}

	if cfg.Static.Serve {
		h := http.FileServer(http.Dir(cfg.Static.Path))
		mux.Handle("/*", h)
	}

	return Service{mux: mux}, nil
}

func (s Service) GetMux() chi.Mux {
	return *s.mux
}
