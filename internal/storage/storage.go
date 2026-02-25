package storage

import "github.com/ulibaysya/krona/internal/storage/types"

type Storage interface {
	Getter
}

type Getter interface {
	GetCatalog(id int64) (types.Catalog, error)

	GetCatalogs() ([]types.Catalog, error)
	GetBanners() ([]types.Banner, error)

	GetCatalogParameter(id int64) types.CatalogParameter
	GetProductByID(id int64) types.Product
	GetProductByName(name string) types.Product
}
