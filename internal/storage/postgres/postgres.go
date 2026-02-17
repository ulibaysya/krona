package postgres

import (
	"context"
	"fmt"

	// "github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ulibaysya/krona/internal/config"
	"github.com/ulibaysya/krona/internal/storage/types"
)

type postgres struct {
	pool *pgxpool.Pool
}

func New(cfg config.RDBMS) (postgres, error) {
	pool, err := pgxpool.New(context.Background(), cfg.Connstr)
	if err != nil {
		return postgres{}, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return postgres{}, err
	}

	return postgres{pool: pool}, nil
}

func (db postgres) GetCatalog(id int32) (types.Catalog, error) {
	const f = "github.com/ulibaysya/krona/internal/storage/postgres.GetCatalog"

	const sql = "SELECT id, name, ru_name FROM catalogs WHERE id = $1"

	ctlg := types.Catalog{}
	// pgId := pgtype.Int4{}
	if err := db.pool.QueryRow(context.Background(), sql, id).Scan(&ctlg.ID, &ctlg.Name, &ctlg.RuName); err != nil {
		return types.Catalog{}, fmt.Errorf("%s: %w", f, err)
	}

	// ctlg.ID = pgId.Int32

	return ctlg, nil
}

func (db postgres) GetCatalogParameter(id int32) types.CatalogParameter {
	return types.CatalogParameter{}
}

func (db postgres) GetProductByID(id int32) types.Product {
	return types.Product{}
}

func (db postgres) GetProductByName(name string) types.Product {
	return types.Product{}
}

