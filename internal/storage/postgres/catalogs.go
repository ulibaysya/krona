package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/ulibaysya/krona/internal/storage"
	"github.com/ulibaysya/krona/internal/storage/types"
)

func (db postgres) AddCatalogs(catalogs []types.CatalogRow) ([]types.CatalogRow, error) {
	ctx := context.Background()
	const sql = "INSERT INTO catalogs(alias, img, ru_name) VALUES ($1, $2, $3) RETURNING id, addition_date;"

	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return nil, wrapError("beginning transaction (while adding catalogs)", err)
	}
	defer tx.Rollback(ctx)

	for i, ctlg := range catalogs {
		err := tx.QueryRow(ctx, sql,
			ctlg.Alias,
			ctlg.Img,
			ctlg.RuName,
		).Scan(&ctlg.ID, &ctlg.AdditionDate)
		if err != nil {
			// TODO instead of this we should implement feature that would tell user what was wrong when inserting: such catalog already exists, image already exists, name is not russian, etc
			msg := fmt.Sprintf("inserting catalog(%v)", ctlg)
			if isUniqueViolation(err) {
				return nil, wrapError(msg, storage.ErrUniqueViolation)
			}
			return nil, wrapError(msg, err)
		}
		catalogs[i] = ctlg
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, wrapError("commiting transaction (while adding catalogs)", err)
	}

	return catalogs, nil
}

func (db postgres) GetCatalogs(ids []int64) ([]types.CatalogRow, error) {
	ctx := context.Background()
	const sql = "SELECT id, alias, img, ru_name, addition_date FROM catalogs WHERE id = $1 LIMIT 1;"

	batch := pgx.Batch{}

	catalogs := make([]types.CatalogRow, len(ids))
	for i, id := range ids {
		batch.Queue(sql, id).QueryRow(func(row pgx.Row) error {
			err := row.Scan(
				&catalogs[i].ID,
				&catalogs[i].Alias,
				&catalogs[i].Img,
				&catalogs[i].RuName,
				&catalogs[i].AdditionDate,
			)
			if err != nil {
				return err
			}
			return nil
		})
	}

	if err := db.pool.SendBatch(ctx, &batch).Close(); err != nil {
		wrapError("getting catalogs", err)
	}

	return catalogs, nil
}

func (db postgres) GetCatalogsAliases(aliases []string) ([]types.CatalogRow, error) {
	ctx := context.Background()
	const sql = "SELECT id, alias, img, ru_name, addition_date FROM catalogs WHERE alias = $1 LIMIT 1;"

	batch := pgx.Batch{}

	catalogs := make([]types.CatalogRow, len(aliases))
	for i, id := range aliases {
		batch.Queue(sql, id).QueryRow(func(row pgx.Row) error {
			err := row.Scan(
				&catalogs[i].ID,
				&catalogs[i].Alias,
				&catalogs[i].Img,
				&catalogs[i].RuName,
				&catalogs[i].AdditionDate,
			)
			if err != nil {
				return err
			}
			return nil
		})
	}

	if err := db.pool.SendBatch(ctx, &batch).Close(); err != nil {
		wrapError("getting catalogs by aliases", err)
	}

	return catalogs, nil
}

func (db postgres) GetAllCatalogs() ([]types.CatalogRow, error) {
	ctx := context.Background()
	const sql = "SELECT id, alias, img, ru_name, addition_date FROM catalogs ORDER BY id;"

	catalogs, err := queryAndCollect[types.CatalogRow](ctx, db, sql, nil)
	if err != nil {
		return nil, wrapError("getting all catalogs", err)
	}

	return catalogs, nil
}

func (db postgres) DelCatalogs(ids []int64) error {
	ctx := context.Background()
	const sql = "DELETE FROM catalogs WHERE id = $1;"

	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return wrapError("beginning transaction (while deleting catalogs)", err)
	}
	defer tx.Rollback(ctx)

	for _, id := range ids {
		res, err := tx.Exec(ctx, sql, id)
		if err != nil {
			return wrapError("deleting catalogs", err)
		}
		if aff := res.RowsAffected(); aff == 0 {
			return wrapError("no deleted catalogs", storage.ErrNoneAffected)
		} else if aff > 1 {
			return wrapError("several catalogs deleted when one expected(should panic)", nil)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return wrapError("commiting transaction (while deleting catalogs)", err)
	}

	return nil
}

func (db postgres) DelCatalogsAliases(aliases []string) error {
	ctx := context.Background()
	const sql = "DELETE FROM catalogs WHERE alias = $1;"

	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return wrapError("beginning transaction (while deleting catalogs by aliases)", err)
	}
	defer tx.Rollback(ctx)

	for _, alias := range aliases {
		res, err := tx.Exec(ctx, sql, alias)
		if err != nil {
			return wrapError("deleting catalogs", err)
		}
		if aff := res.RowsAffected(); aff == 0 {
			return wrapError("no deleted catalogs", storage.ErrNoneAffected)
		} else if aff > 1 {
			return wrapError("several catalogs deleted when one expected(should panic)", nil)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return wrapError("commiting transaction (while deleting catalogs by aliases)", err)
	}

	return nil
}

// func scanCatalog(row pgx.CollectableRow) (types.Catalog, error) {
// 	fmt.Println("WORKS???", row.FieldDescriptions())
// 	var catalog types.Catalog
// 	if err := row.Scan(&catalog.ID, &catalog.Alias, &catalog.Img, &catalog.RuName, &catalog.AdditionDate); err != nil {
// 		return types.Catalog{}, err
// 	}
// 	return catalog, nil
// }
