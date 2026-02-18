package types

type Catalog struct {
	ID           int32
	Name, RuName string
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
