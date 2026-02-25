package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ulibaysya/krona/internal/storage"
)

func GetCatalogs(strg storage.Storage, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// catalogs, err := strg.GetCatalogs()
		// if err != nil {
		// 	json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		// 	return
		// }
		//
		// head := Head{
		// 	Title: fmt.Sprintf("Представляем товары из %v типов категорий", len(catalogs)),
		// }
		//
		// if err := tmpl.Execute(w, Default{head, catalogs}); err != nil {
		// 	// .HeadTitle
		// 	json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
		// 	return
		// }
	}
}

func GetCatalogsID(strg storage.Storage, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("we are at catalogsID")
		tmpl.Execute(w, nil)
	}
}

