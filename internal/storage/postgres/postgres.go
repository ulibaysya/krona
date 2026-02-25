package postgres

import (
	"context"
	"fmt"

	// "github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgtype"
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

func (db postgres) GetCatalog(id int64) (types.Catalog, error) {
	const f = "github.com/ulibaysya/krona/internal/storage/postgres.GetCatalog"

	const sql = "SELECT id, name, ru_name FROM catalogs WHERE id = $1"

	ctlg := types.Catalog{}
	// pgId := pgtype.Int4{}
	// if err := db.pool.QueryRow(context.Background(), sql, id).Scan(&ctlg.ID, &ctlg.Name, &ctlg.RuName); err != nil {
		// return types.Catalog{}, fmt.Errorf("%s: %w", f, err)
	// }

	// ctlg.ID = pgId.int64

	return ctlg, nil
}

func (db postgres) GetCatalogs() ([]types.Catalog, error) {
	const f = "github.com/ulibaysya/krona/internal/storage/postgres.GetCatalog"

	rows, err := db.pool.Query(context.Background(), "SELECT id, alias, img, ru_name FROM catalogs")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", f, err)
	}
	defer rows.Close()

	catalogs := make([]types.Catalog, 0, 20)
	fmt.Println(rows.RawValues(), len(rows.RawValues()), cap(rows.RawValues()))
	for rows.Next() {
		tmp := types.Catalog{}
		if err := rows.Scan(&tmp.ID, &tmp.Alias, &tmp.Img, &tmp.RuName); err != nil {
			return nil, fmt.Errorf("%s: %w", f, err)
		}
		catalogs = append(catalogs, tmp)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", f, err)
	}

	fmt.Println(catalogs)

	return catalogs, nil
}

func (db postgres) GetBanners() ([]types.Banner, error) {
	const f = "github.com/ulibaysya/krona/internal/storage/postgres.GetBanners"

	rows, err := db.pool.Query(context.Background(), "SELECT id, alias, img, redirect_url FROM banners")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", f, err)
	}
	defer rows.Close()

	banners := []types.Banner{}
	for rows.Next() {
		tmp := types.Banner{}
		tmpRedirectURL := pgtype.Text{}
		if err := rows.Scan(&tmp.ID, &tmp.Alias, &tmp.Img, &tmpRedirectURL); err != nil {
			return nil, fmt.Errorf("%s: %w", f, err)
		}
		if tmpRedirectURL.Valid {
			tmp.RedirectURL = tmpRedirectURL.String
		}
		banners = append(banners, tmp)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", f, err)
	}

	fmt.Println(banners)

	return banners, nil
}

func (db postgres) GetCatalogParameter(id int64) types.CatalogParameter {
	return types.CatalogParameter{}
}

func (db postgres) GetProductByID(id int64) types.Product {
	return types.Product{}
}

func (db postgres) GetProductByName(name string) types.Product {
	return types.Product{}
}
