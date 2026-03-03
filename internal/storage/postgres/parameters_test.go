package postgres

import (
	"errors"
	"os"
	"testing"

	"github.com/jackc/pgerrcode"
	"github.com/ulibaysya/krona/internal/config"
	"github.com/ulibaysya/krona/internal/storage"
	"github.com/ulibaysya/krona/internal/storage/types"
)

func TestParameters(t *testing.T) {
	db, err := New(config.RDBMS{Engine: "postgres", Connstr: os.Getenv("KRONA_TEST_PGCONNSTR")})
	if err != nil {
		t.Fatalf("error while initializing db: %v", err)
	}

	tc := []struct {
		params    []types.ParameterRow
		isFailing bool
		expectedCode string
	}{
		{
			params: []types.ParameterRow{
				{
					Key:        "room",
					RuKey:      "Комната",
					Value:      "hall",
					RuValue:    "В прихожую",
					CatalogsID: 60,
				},
			},
			isFailing: false,
		},
		{
			params: []types.ParameterRow{
				{
					Key:        "room",
					RuKey:      "Комната",
					Value:      "hall",
					RuValue:    "В прихожую",
					CatalogsID: 60,
				},
				{
					Key:        "room",
					RuKey:      "Комната",
					Value:      "bedroom",
					RuValue:    "В спальню",
					CatalogsID: 60,
				},
			},
			isFailing: false,
		},
		{
			params: []types.ParameterRow{
				{
					Key:        "room",
					RuKey:      "Комната",
					Value:      "hall",
					RuValue:    "В прихожую",
					CatalogsID: 60,
				},
				{
					Key:        "room",
					RuKey:      "Комната",
					Value:      "hall",
					RuValue:    "В прихожую",
					CatalogsID: 60,
				},
			},
			isFailing: true,
			expectedCode: pgerrcode.UniqueViolation,
		},
	}

	tcID := 2
	inserted, err := db.AddParameters(tc[tcID].params)

	// pgErr, ok := err.(*pgconn.PgError)
	// if !ok {
	// 	t.Fatalf("error asserting err")
	// }

	if tc[tcID].isFailing {
		if err != nil {
			t.Logf("received expected fail; params: %v; error: %v", tc[tcID].params, err)
			if errors.Is(err, storage.ErrUniqueViolation){
				t.Log("unique")
			} else {
				t.Log("NOT unique")
			}
		} else {
			t.Fatalf("expected for fail, but inserted successfully; inserted: %v", inserted)
		}
	} else {
		if err != nil {
			t.Fatalf("receivued UNexpected fail; params: %v; error: %v", tc[tcID].params, err)
		} else {
			t.Logf("inserted successfully; inserted: %v", inserted)
		}
	}
	// t.Logf("inserted parameters: %v", inserted)
}
