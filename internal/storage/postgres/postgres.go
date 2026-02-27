package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ulibaysya/krona/internal/config"
	"github.com/ulibaysya/krona/internal/storage/types"
)

type postgres struct {
	pool *pgxpool.Pool
}

func New(cfg config.RDBMS) (postgres, error) {
	const f = "internal/storage/postgres.New"

	pool, err := pgxpool.New(context.Background(), cfg.Connstr)
	if err != nil {
		return postgres{}, fmt.Errorf("%s: %w", f, err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return postgres{}, fmt.Errorf("%s: %w", f, err)
	}

	return postgres{pool: pool}, nil
}

func (db postgres) InsertCatalog(catalog types.Catalog) (types.Catalog, error) {
	const f = "internal/storage/postgres.InsertCatalog"

	const sql = "INSERT INTO catalogs(alias, img, ru_name) VALUES ($1, $2, $3) RETURNING id, addition_date;"

	if err := db.pool.QueryRow(context.Background(), sql, catalog.Alias, catalog.Img, catalog.RuName).Scan(&catalog.ID, &catalog.AdditionDate); err != nil {
		return types.Catalog{}, fmt.Errorf("%s: %w", f, err)
	}

	return catalog, nil
}

func (db postgres) GetCatalog(id int64) (types.Catalog, error) {
	const f = "internal/storage/postgres.GetCatalog"

	const sql = "SELECT id, alias, img, ru_name, addition_date FROM catalogs WHERE id = $1;"

	var catalog types.Catalog
	if err := db.pool.QueryRow(context.Background(), sql, id).Scan(&catalog.ID, &catalog.Alias, &catalog.Img, &catalog.RuName, &catalog.AdditionDate); err != nil {
		return types.Catalog{}, fmt.Errorf("%s: %w", f, err)
	}

	return catalog, nil
}

func (db postgres) GetCatalogs() ([]types.Catalog, error) {
	const f = "internal/storage/postgres.GetCatalogs"

	const sql = "SELECT id, alias, img, ru_name, addition_date FROM catalogs;"

	catalogs, err := scanMultiple[types.Catalog](db, sql, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", f, err)
	}
	// rows, err := db.pool.Query(context.Background(), sql)
	// if err != nil {
	// 	return nil, fmt.Errorf("%s: %w", f, err)
	// }
	// catalogs, err := pgx.CollectRows(rows, scanCatalog)
	// if err != nil {
	// 	return nil, fmt.Errorf("%s: %w", f, err)
	// }

	fmt.Println(catalogs)

	return catalogs, nil
}

func (db postgres) GetBanners() ([]types.Banner, error) {
	const f = "internal/storage/postgres.GetBanners"

	const sql = "SELECT id, alias, img, redirect_url, addition_date FROM banners;"

	banners, err := scanMultiple[types.Banner](db, sql, scanBanner)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", f, err)
	}
	// pgx.RowToStructByPos[types.Banner](db.pool.QueryRow(context.Background(), sql))
	// rows, err := db.pool.Query(context.Background(), sql)
	// if err != nil {
	// 	return nil, fmt.Errorf("%s: %w", f, err)
	// }
	// banners, err := pgx.CollectRows(rows, scanBanner)
	// if err != nil {
	// 	return nil, fmt.Errorf("%s: %w", f, err)
	// }

	return banners, nil
}

// func (db postgres) GetCatalogParameter(id int64) types.CatalogParameter {
// 	return types.CatalogParameter{}
// }
//
// func (db postgres) GetProductByID(id int64) types.Product {
// 	return types.Product{}
// }
//
// func (db postgres) GetProductByName(name string) types.Product {
// 	return types.Product{}
// }

func scanMultiple[T any](db postgres, sql string, fn pgx.RowToFunc[T]) ([]T, error) {
	rows, err := db.pool.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}

	var collectable []T
	if fn == nil {
		collectable, err = pgx.CollectRows(rows, pgx.RowToStructByPos[T]) // TODO maybe we should force caller to pass pgx.RowToStructByPos[T], but if nil - return error?
	} else {
		collectable, err = pgx.CollectRows(rows, fn)
	}
	if err != nil {
		return nil, err
	}

	return collectable, nil
}

// func scanCatalog(row pgx.CollectableRow) (types.Catalog, error) {
// 	fmt.Println("WORKS???", row.FieldDescriptions())
// 	var catalog types.Catalog
// 	if err := row.Scan(&catalog.ID, &catalog.Alias, &catalog.Img, &catalog.RuName, &catalog.AdditionDate); err != nil {
// 		return types.Catalog{}, err
// 	}
// 	return catalog, nil
// }

func scanBanner(row pgx.CollectableRow) (types.Banner, error) {
	var banner types.Banner
	var redirectURL zeronull.Text
	if err := row.Scan(&banner.ID, &banner.Alias, &banner.Img, &redirectURL, &banner.AdditionDate); err != nil {
		return types.Banner{}, err
	}

	banner.RedirectURL = string(redirectURL)

	return banner, nil
}
