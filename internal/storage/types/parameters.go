package types

import "time"

type Value struct {
	ParamID           int64
	Value, RuValue    string
	ParamAdditionDate time.Time
}

type Parameter struct {
	Key, RuKey string
	CatalogsID int64
	Values     []Value
}

type CatalogParameters struct {
	CatalogRow
	Parameters []Parameter
}
