package handlers

import (
	"html/template"
	"net/http"

	"github.com/ulibaysya/krona/internal/storage"
)

func GetProductID(strg storage.Storage, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("we are at product")
		// catalogs, err := strg.GetCatalogs()
		// if err != nil {
		// 	w.WriteHeader(500)
		// 	json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		// }
		// w.WriteHeader(200)
		// json.NewEncoder(w).Encode(catalogs)
	}
}
