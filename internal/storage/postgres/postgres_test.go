package postgres

import (
	"crypto/rand"
	"os"
	"slices"
	"testing"

	"github.com/ulibaysya/krona/internal/config"
	"github.com/ulibaysya/krona/internal/storage/types"
)

func TestCatalogs(t *testing.T) {
	db, err := New(config.RDBMS{Connstr: os.Getenv("KRONA_TEST_PGCONNSTR")})
	if err != nil {
		t.Fatalf("expected new postgres database: %v", err)
	}

	// catalogs := types.Catalog{Alias: "test_catalog", RuName: "Тестовый каталог", Img: "test_catalog.jpg"}
	// catalogs := types.Catalog{Alias: rand.Text(), RuName: rand.Text(), Img: rand.Text()}
	// catalogs := make([]types.Catalog, 3)
	catalogs := make([]types.Catalog, 10)
	for i := range catalogs {
		catalogs[i] = types.Catalog{Alias: rand.Text(), RuName: rand.Text(), Img: rand.Text()}
	}

	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{"InsertCatalog", func(t *testing.T) {
			tmp, err := db.InsertCatalog(catalogs[0])
			if err != nil {
				t.Fatalf("expected catalog with id and addition date: %v", err)
			}
			t.Logf("received catalog: %v", tmp)
			catalogs[0] = tmp
		}},
		{"GetCatalog", func(t *testing.T) {
			tmp, err := db.GetCatalog(catalogs[0].ID)
			if err != nil {
				t.Fatalf("expected catalog: %v", err)
			}
			if tmp != catalogs[0] {
				t.Fatalf("expected catalog: %#v; received catalog: %#v", catalogs, tmp)
			}
			t.Logf("received expected catalog: %#v", tmp)
		}},
		{"GetCatalogAlias", func(t *testing.T) {
			tmp, err := db.GetCatalogAlias(catalogs[0].Alias)
			if err != nil {
				t.Fatalf("expected catalog: %v", err)
			}
			if tmp != catalogs[0] {
				t.Fatalf("expected catalog: %#v; received catalog: %#v", catalogs, tmp)
			}
			t.Logf("received expected catalog: %#v", tmp)
		}},
		{"DeleteCatalog", func(t *testing.T) {
			if err := db.DeleteCatalog(catalogs[0].ID); err != nil {
				t.Fatalf("error while deleting catalog: %v", err)
			}
		}},
		{"GetCatalogs", func(t *testing.T) {
			for i := range catalogs {
				catalogs[i], err = db.InsertCatalog(catalogs[i])
				if err != nil {
					t.Fatalf("error while inserting catalog: %v", err)
				}
			}
			allCatalogs, err := db.GetCatalogs()
			if err != nil {
				t.Fatalf("expected catalog: %v", err)
			}
			if !slices.Equal(allCatalogs, catalogs) {
				t.Fatalf("error while comparing slices")
				// t.Fatalf("error while comparing slices; expected: %#v; received: %#v", catalogs, allCatalogs)
			}
		}},
		{"DeleteCatalogAlias", func(t *testing.T) {
			for _, i := range catalogs {
				if err := db.DeleteCatalogAlias(i.Alias); err != nil {
					t.Fatalf("error while deleting; catalog: %#v; error: %v",  i, err)
				}
			}
		}},
	}

	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}

func insertCatalog(t *testing.T, db postgres, catalog types.Catalog) func(t *testing.T) {
	return func(t *testing.T) {

	}
}

// func TestInsertCatalog(t *testing.T) {
// 	var err error
// 	testCatalog, err = db.InsertCatalog(types.Catalog{Alias: "test_alias", RuName: "Тестовый каталог", Img: "test_catalog.jpg"})
// 	if err != nil {
// 		t.Fatalf("expected catalog: %v", err)
// 	}
// 	t.Logf("inserted catalog: %+v", catalog)
// }
//
// func TestGetCatalog(t *testing.T) {
// 	catalog, err := db.GetCatalog(1)
// 	if err != nil {
// 		t.Fatalf("expected catalog: %v", err)
// 	}
// 	t.Logf("received catalog: %+v", catalog)
// }
