package storage

import "github.com/ulibaysya/krona/internal/storage/types"

type Storage interface {
	Catalogs
	Banners
	Parameters
}

type Catalogs interface {
	AddCatalogs(catalogs []types.CatalogRow) ([]types.CatalogRow, error) // TODO change to AddCatalogs and do the same as in AddParameters

	GetCatalogs(id []int64) ([]types.CatalogRow, error)
	GetCatalogsAliases(aliases []string) ([]types.CatalogRow, error)

	GetAllCatalogs() ([]types.CatalogRow, error)

	DelCatalogs(id []int64) error
	DelCatalogsAliases(alias []string) error
}

type Banners interface {
	GetAllBanners() ([]types.BannerRow, error)
}

type Parameters interface {
	AddParameters(parameters []types.ParameterRow) ([]types.ParameterRow, error)

	GetAllProductParameters() ([]types.CatalogParameters, error)
}
