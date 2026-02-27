package postgres

import (
	"crypto/rand"
	"os"
	"testing"

	"github.com/ulibaysya/krona/internal/config"
	"github.com/ulibaysya/krona/internal/storage/types"
)

func TestCatalogs(t *testing.T) {
	db, err := New(config.RDBMS{Connstr: os.Getenv("KRONA_TEST_PGCONNSTR")})
	if err != nil {
		t.Fatalf("expected new postgres database: %v", err)
	}

	// testCatalog := types.Catalog{Alias: "test_catalog", RuName: "Тестовый каталог", Img: "test_catalog.jpg"}
	testCatalog := types.Catalog{Alias: rand.Text(), RuName: rand.Text(), Img: rand.Text()}

	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{"insertCatalog", func(t *testing.T) {
			tmp, err := db.InsertCatalog(testCatalog)
			if err != nil {
				t.Fatalf("expected catalog with id and addition date: %#v", err)
			}
			t.Logf("received catalog: %#v", tmp)
			testCatalog = tmp
		}},
		{"getCatalog", func(t *testing.T) {
			tmp, err := db.GetCatalog(testCatalog.ID)
			if err != nil {
				t.Fatalf("expected catalog: %#v", err)
			}
			if tmp != testCatalog {
				t.Fatalf("expected catalog: %#v; received catalog: %#v", testCatalog, tmp)
			}
			t.Logf("received expected catalog: %#v", tmp)
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
