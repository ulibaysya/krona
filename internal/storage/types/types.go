package types

import "time"

type Catalog struct {
	ID           int64
	Alias        string
	Img          string
	RuName       string
	AdditionDate time.Time
}

type CatalogParameter struct {
	ID             int32
	Key, RuKey     string
	Value, RuValue string
	CatalogID      int32
}

type Product struct {
	ID           int32
	Name, RuName string
}

type Banner struct {
	ID           int64
	Alias        string
	Img          string
	RedirectURL  string
	AdditionDate time.Time
}
