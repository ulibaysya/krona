package storage

import "github.com/ulibaysya/krona/internal/storage/types"

type Storage interface {
	Getter
}

type Getter interface {
	GetCatalog(id int32) (types.Catalog, error)
	GetCatalogParameter(id int32) types.CatalogParameter
	GetProductByID(id int32) types.Product
	GetProductByName(name string) types.Product
}
