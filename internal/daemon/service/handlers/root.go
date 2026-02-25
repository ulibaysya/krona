package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ulibaysya/krona/internal/storage"
	"github.com/ulibaysya/krona/internal/storage/types"
)

func GetRoot(strg storage.Storage, tmpl *template.Template) http.HandlerFunc {
	type templateData struct {
		Head
		Catalogs []types.Catalog
		Banners  []types.Banner
	}

	const title = `Мебель по самым низким ценам`
	const descpription = `ООО «Крона» предоставляет мебель по самым низким ценам в Туапсе. Предоставялем большой ассортимент мебели, отделочных материалов, низкие цены, розничные и оптовые продажи!`
	const keywords = `шкаф мебель кухня стол стул кровать диван гостиная комод спальня ванная кресло пуф матрас дверь двери дверь-купе двери-купе прихожая детская офис`

	head := Head{
		Title:       title,
		Description: descpription,
		Keywords:    keywords,
	}

	return func(w http.ResponseWriter, r *http.Request) {
		catalogs, err := strg.GetCatalogs()
		if err != nil {
			fmt.Println(err)
			HandleInternalServerError(w, r)
			return
		}

		banners, err := strg.GetBanners()
		if err != nil {
			fmt.Println(err)
			HandleInternalServerError(w, r)
			return
		}

		if err := tmpl.Execute(w, templateData{Head: head, Catalogs: catalogs, Banners: banners}); err != nil {
			fmt.Println(err)
			HandleInternalServerError(w, r)
			return
		}
	}
}
