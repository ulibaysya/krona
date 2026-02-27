package cachebased

import (
	"errors"
	"fmt"

	"github.com/ulibaysya/krona/internal/storage"
	"github.com/ulibaysya/krona/internal/storage/types"
)

type cachebased struct {
	rdbms storage.Storage
	cache storage.Storage
}

func New(rdbms, cache storage.Storage) (cachebased, error) {
	const f = "github.com/ulibaysya/krona/internal/storage/cache.New"

	if rdbms == nil {
		return cachebased{}, fmt.Errorf("%s: %w", f, errors.New("invalid rdbms storage(equal to nil)"))
	}
	if cache == nil {
		return cachebased{}, fmt.Errorf("%s: %w", f, errors.New("invalid cache storage(equal to nil)"))
	}

	return cachebased{rdbms: rdbms, cache: cache}, nil
}

func (db cachebased) InsertCatalog(types.Catalog) (types.Catalog, error) {
	return types.Catalog{}, nil
}

func (db cachebased) GetCatalog(id int64) (types.Catalog, error) {
	return types.Catalog{}, nil
}

func (db cachebased) GetCatalogs() ([]types.Catalog, error) {
	return nil, nil
}

func (db cachebased) GetBanners() ([]types.Banner, error) {
	return nil, nil
}

// func (db cachebased) GetCatalogParameter(id int64) types.CatalogParameter {
// 	return types.CatalogParameter{}
// }
//
// func (db cachebased) GetProductByID(id int64) types.Product {
// 	return types.Product{}
// }
//
// func (db cachebased) GetProductByName(name string) types.Product {
// 	return types.Product{}
// }
