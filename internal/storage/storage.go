package storage

import "github.com/ulibaysya/krona/internal/storage/types"

type Storage interface {
	Getter
}

type Getter interface {
	InsertCatalog(types.Catalog) (types.Catalog, error)

	GetCatalog(id int64) (types.Catalog, error)
	GetCatalogAlias(alias string) (types.Catalog, error)

	GetCatalogs() ([]types.Catalog, error)

	DeleteCatalog(id int64) error
	DeleteCatalogAlias(alias string) error

	GetBanners() ([]types.Banner, error)

	// GetCatalogParameter(id int64) types.CatalogParameter
	// GetProductByID(id int64) types.Product
	// GetProductByName(name string) types.Product
}
