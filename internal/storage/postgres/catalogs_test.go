package postgres

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"testing"
	"time"

	"github.com/ulibaysya/krona/internal/config"
	"github.com/ulibaysya/krona/internal/storage"
	"github.com/ulibaysya/krona/internal/storage/types"
)

type testObj struct {
	db          postgres
	showObjects bool
}

var catalogsBeds = map[string]types.CatalogRow{
	"beds": {
		Alias:  "beds",
		Img:    "beds.jpg",
		RuName: "Кровати",
	},
	"beds_alias": {
		Alias:  "bed",
		Img:    "beds.jpg",
		RuName: "Кровати",
	},
	"beds_img": {
		Alias:  "beds",
		Img:    "beds.webp",
		RuName: "Кровати",
	},
	"beds_ru_name": {
		Alias:  "beds",
		Img:    "beds.jpg",
		RuName: "Кроваточки",
	},
	"beds_alias_img": {
		Alias:  "bed",
		Img:    "beds.webp",
		RuName: "Кровати",
	},
	"beds_img_ru_name": {
		Alias:  "bed",
		Img:    "beds.webp",
		RuName: "Кровати",
	},
	"beds_alias_ru_name": {
		Alias:  "beds",
		Img:    "beds.webp",
		RuName: "Кроватки",
	},
	"beds_all": {
		Alias:  "bed",
		Img:    "beds.webp",
		RuName: "Кроватки",
	},
}

var catalogsDifferent = map[string]types.CatalogRow{
	"chairs": {
		Alias:  "chairs",
		Img:    "chairs.jpg",
		RuName: "Стулья",
	},
	"wardrobes": {
		Alias:  "wardrobes",
		Img:    "wardrobes.jpg",
		RuName: "Шкафы",
	},
	"tables": {
		Alias:  "tables",
		Img:    "tables.jpg",
		RuName: "Столы",
	},
}

func TestCatalogs(t *testing.T) {
	db, err := New(config.RDBMS{Connstr: os.Getenv("KRONA_TEST_PGCONNSTR")})
	if err != nil {
		t.Fatalf("expected new postgres database: %v", err)
	}
	obj := testObj{db: db}
	if os.Getenv("KRONA_TEST_SHOWOBJECTS") == "1" {
		obj.showObjects = true
	} else {
		obj.showObjects = false
	}

	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{"AddCatalogs", obj.addCatalogs()},
		{"GetCatalogs", obj.getCatalogsAliases()},
	}

	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}

func (obj testObj) addCatalogs() func(t *testing.T) {
	testcases := []struct {
		caseName string
		ctlgs    []types.CatalogRow
		mustFail bool
		failWith error
	}{
		{
			caseName: "beds first",
			ctlgs:    []types.CatalogRow{catalogsBeds["beds"]},
			mustFail: false,
		},
		{
			caseName: "beds second",
			ctlgs:    []types.CatalogRow{catalogsBeds["beds"]},
			mustFail: true,
			failWith: storage.ErrUniqueViolation,
		},
		{
			caseName: "beds new alias",
			ctlgs:    []types.CatalogRow{catalogsBeds["beds_alias"]},
			mustFail: true,
			failWith: storage.ErrUniqueViolation,
		},
		{
			caseName: "beds new img",
			ctlgs:    []types.CatalogRow{catalogsBeds["beds_img"]},
			mustFail: true,
			failWith: storage.ErrUniqueViolation,
		},
		{
			caseName: "beds new ru_name",
			ctlgs:    []types.CatalogRow{catalogsBeds["beds_ru_name"]},
			mustFail: true,
			failWith: storage.ErrUniqueViolation,
		},
		{
			caseName: "beds new alias, img",
			ctlgs:    []types.CatalogRow{catalogsBeds["beds_alias_img"]},
			mustFail: true,
			failWith: storage.ErrUniqueViolation,
		},
		{
			caseName: "beds new img, ru_name",
			ctlgs:    []types.CatalogRow{catalogsBeds["beds_img_ru_name"]},
			mustFail: true,
			failWith: storage.ErrUniqueViolation,
		},
		{
			caseName: "beds new alias, ru_name",
			ctlgs:    []types.CatalogRow{catalogsBeds["beds_alias_ru_name"]},
			mustFail: true,
			failWith: storage.ErrUniqueViolation,
		},
		{
			caseName: "beds new all",
			ctlgs:    []types.CatalogRow{catalogsBeds["beds_all"]},
			mustFail: false,
		},
		{
			caseName: "chairs+beds",
			ctlgs:    []types.CatalogRow{catalogsDifferent["chairs"], catalogsBeds["beds"]},
			mustFail: true,
			failWith: storage.ErrUniqueViolation,
		},
		{
			caseName: "wardrobes+tables",
			ctlgs:    []types.CatalogRow{catalogsDifferent["wardrobes"], catalogsDifferent["tables"]},
			mustFail: false,
		},
		{
			caseName: "wardrobes",
			ctlgs:    []types.CatalogRow{catalogsDifferent["wardrobes"]},
			mustFail: true,
			failWith: storage.ErrUniqueViolation,
		},
		{
			caseName: "tables",
			ctlgs:    []types.CatalogRow{catalogsDifferent["tables"]},
			mustFail: true,
			failWith: storage.ErrUniqueViolation,
		},
	}

	return func(t *testing.T) {
		for _, tc := range testcases {
			_, err := obj.db.AddCatalogs(tc.ctlgs)
			if obj.showObjects {
				t.Logf("%v: adding catalogs: %+v", tc.caseName, tc.ctlgs)
			}
			if msg, yes := isFailed(tc.mustFail, err, tc.failWith); yes {
				t.Fatalf("%v: failed: %v", tc.caseName, msg)
			} else {
				t.Logf("%v: success: %v", tc.caseName, msg)
			}
		}
	}
}

func (obj testObj) getCatalogsAliases() func(t *testing.T) {
	testcases := []struct {
		caseName string
		toQuery []string
		expectingCatalogs    []types.CatalogRow
		failWith error
	}{
		{
			caseName: "beds",
			toQuery: []string{"beds"},
			expectingCatalogs:    []types.CatalogRow{catalogsBeds["beds"]},
		},
		{
			caseName: "beds new alias",
			expectingCatalogs:    []types.CatalogRow{catalogsBeds["beds_alias"]},
			failWith: storage.ErrUniqueViolation,
		},
		{
			caseName: "beds new img",
			expectingCatalogs:    []types.CatalogRow{catalogsBeds["beds_img"]},
			failWith: storage.ErrUniqueViolation,
		},
		{
			caseName: "beds new ru_name",
			expectingCatalogs:    []types.CatalogRow{catalogsBeds["beds_ru_name"]},
			failWith: storage.ErrUniqueViolation,
		},
		{
			caseName: "beds new alias, img",
			expectingCatalogs:    []types.CatalogRow{catalogsBeds["beds_alias_img"]},
			failWith: storage.ErrUniqueViolation,
		},
		{
			caseName: "beds new img, ru_name",
			expectingCatalogs:    []types.CatalogRow{catalogsBeds["beds_img_ru_name"]},
			failWith: storage.ErrUniqueViolation,
		},
		{
			caseName: "beds new alias, ru_name",
			expectingCatalogs:    []types.CatalogRow{catalogsBeds["beds_alias_ru_name"]},
			failWith: storage.ErrUniqueViolation,
		},
		{
			caseName: "beds new all",
			expectingCatalogs:    []types.CatalogRow{catalogsBeds["beds_all"]},
		},
		{
			caseName: "chairs+beds",
			expectingCatalogs:    []types.CatalogRow{catalogsDifferent["chairs"], catalogsBeds["beds"]},
		},
		{
			caseName: "wardrobes+tables",
			expectingCatalogs:    []types.CatalogRow{catalogsDifferent["wardrobes"], catalogsDifferent["tables"]},
		},
		{
			caseName: "wardrobes",
			expectingCatalogs:    []types.CatalogRow{catalogsDifferent["wardrobes"]},
			failWith: storage.ErrUniqueViolation,
		},
		{
			caseName: "tables",
			expectingCatalogs:    []types.CatalogRow{catalogsDifferent["tables"]},
			failWith: storage.ErrUniqueViolation,
		},
	}

	return func(t *testing.T) {
		for _, tc := range testcases {
			if obj.showObjects {
				t.Logf("%v: querying catalogs: %v", tc.caseName, tc.toQuery)
			}
			ctlgs, err := obj.db.GetCatalogsAliases(tc.toQuery)
			if tc.failWith != nil {
				msg, failed := isFailed(true, err, tc.failWith)
				if failed {
					t.Fatalf("failure: %v", msg)
				} else {
					t.Logf("success: %v", msg)
				}
			} else {
				// TODO you should compare somehow these properties to(maybe fill example catalogs while addCatalogs runs)
				for i, ctlg := range ctlgs{
					ctlg.ID = 0
					ctlg.AdditionDate = time.Time{}
					ctlgs[i] = ctlg
				}

				if obj.showObjects {
					t.Logf("%v: received catalogs: %v; comparing with: %v", tc.caseName, ctlgs, tc.expectingCatalogs)
				}
				if slices.Equal(tc.expectingCatalogs, ctlgs) {
					t.Log("success: received expected catalogs")
				} else {
					t.Fatal("failure: received unexpected catalogs")
				}
			}
			// if msg, yes := isFailed(tc.mustFail, err, tc.failWith); yes {
				// t.Fatalf("%v: failed: %v", tc.caseName, msg)
			// } else {
				// t.Logf("%v: success: %v", tc.caseName, msg)
			// }
		}
	}
}

func isFailed(waitingError bool, err, waitingFor error) (string, bool) {
	if waitingError {
		if err != nil {
			if waitingFor != nil {
				if errors.Is(err, waitingFor) {
					return fmt.Sprintf("received expected specified error: %v", err), false
				} else {
					return fmt.Sprintf("received expected error, but unexpected error type: %v; expected type: %v", err, waitingFor), true
				}
			} else {
				return fmt.Sprintf("received expected unspecified error: %v", err), false
			}
		} else {
			if waitingFor != nil {
				return fmt.Sprintf("unexpected no error, expected specified error: %v", waitingFor), true
			} else {
				return "unexpected no error, expected unspecified error", true
			}
		}
	} else {
		if err != nil {
			return fmt.Sprintf("unexpected error, expected no error: %v", err), true
		} else {
			return "expected no error", false
		}
	}
}
