package types

import "time"

type CatalogRow struct {
	ID           int64
	Alias        string
	Img          string
	RuName       string
	AdditionDate time.Time
}

type BannerRow struct {
	ID           int64
	Alias        string
	Img          string
	RedirectURL  string
	AdditionDate time.Time
}

type ParameterRow struct {
	ID             int64
	Key, Value     string
	RuKey, RuValue string
	CatalogsID     int64
	AdditionDate   time.Time
}
