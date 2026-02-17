package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/ulibaysya/krona/internal/storage"
)

func GetRoot(strg storage.Storage, tmpl *template.Template) http.HandlerFunc {
	fmt.Printf("%v %v\n", tmpl, &tmpl)
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("we are here")
		tmpl.Execute(w, nil)
	}
}

func GetCatalogs(strg storage.Storage, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func GetCatalogsID(strg storage.Storage, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "catalogsID"), 10, 32)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		ctlg, err := strg.GetCatalog(int32(id))
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		fmt.Fprintf(w, "%+v", ctlg)
	}
}
