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

func (db cachebased) InsertCatalog(types.CatalogRow) (types.CatalogRow, error) {
	return types.CatalogRow{}, nil
}

func (db cachebased) GetCatalog(id int64) (types.CatalogRow, error) {
	return types.CatalogRow{}, nil
}

func (db cachebased) GetCatalogs() ([]types.CatalogRow, error) {
	return nil, nil
}

func (db cachebased) GetBanners() ([]types.BannerRow, error) {
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
