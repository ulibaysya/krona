package postgres

import (
	"context"
	"fmt"

	"github.com/ulibaysya/krona/internal/storage"
	"github.com/ulibaysya/krona/internal/storage/types"
)

func (db postgres) AddParameters(parameters []types.ParameterRow) ([]types.ParameterRow, error) {
	const sql = "INSERT INTO parameters (key, value, ru_key, ru_value, catalogs_id) VALUES ($1, $2, $3, $4, $5) RETURNING id, addition_date;"

	tx, err := db.pool.Begin(context.Background())
	if err != nil {
		return nil, wrapError("beginning transaction", err)
	}
	defer tx.Rollback(context.Background())

	for i, param := range parameters {
		err := tx.QueryRow(context.Background(), sql,
			param.Key,
			param.Value,
			param.RuKey,
			param.RuValue,
			param.CatalogsID).Scan(&param.ID,
			&param.AdditionDate,
		)
		if err != nil {
			 // TODO instead of this we should implement feature that would tell user what was wrong when inserting: wrond catalogs_id, key already exists, key with such value already exists, etc
			msg := fmt.Sprintf("inserting parameter(%v)", param)
			if isUniqueViolation(err) {
				return nil, wrapError(msg, storage.ErrUniqueViolation)
			}
			return nil, wrapError(msg, err)
		}
		parameters[i] = param
	}

	if err := tx.Commit(context.Background()); err != nil {
		return nil, wrapError("commiting transaction", err)
	}

	return parameters, nil
}

func (db postgres) GetAllProductParameters() ([]types.CatalogParameters, error) {
	const f = "postgres.GetAllProductParameters"

	const sql = `SELECT catalogs.id, alias, img, ru_name, catalogs.addition_date, 
		key, ru_key, catalogs_id, parameters.id, value, ru_value, parameters.addition_date 
		FROM catalogs, parameters WHERE catalogs.id = catalogs_id ORDER BY catalogs.id, parameters.key, parameters.value;`

	rows, err := db.pool.Query(context.Background(), sql)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", f, err)
	}
	defer rows.Close()

	rows.Next()

	ctlgs := []types.CatalogParameters{
		types.CatalogParameters{
			Parameters: []types.Parameter{
				types.Parameter{
					Values: make([]types.Value, 1),
				},
			},
		},
	}

	if err = rows.Scan(
		&ctlgs[0].ID,
		&ctlgs[0].Alias,
		&ctlgs[0].Img,
		&ctlgs[0].RuName,
		&ctlgs[0].AdditionDate,
		&ctlgs[0].Parameters[0].Key,
		&ctlgs[0].Parameters[0].RuKey,
		&ctlgs[0].Parameters[0].CatalogsID,
		&ctlgs[0].Parameters[0].Values[0].ParamID,
		&ctlgs[0].Parameters[0].Values[0].Value,
		&ctlgs[0].Parameters[0].Values[0].RuValue,
		&ctlgs[0].Parameters[0].Values[0].ParamAdditionDate,
	); err != nil {
		return nil, fmt.Errorf("%s: %w", f, err)
	}

	latestCatalog := 0
	latestKV := 0

	for rows.Next() {
		var checkKey string
		var checkCatalogsID int64

		val := types.Value{}
		if err := rows.Scan(
			nil,
			nil,
			nil,
			nil,
			nil,
			&checkKey,
			nil,
			&checkCatalogsID,
			&val.ParamID,
			&val.Value,
			&val.RuValue,
			&val.ParamAdditionDate,
		); err != nil {
			return nil, fmt.Errorf("%s: %w", f, err)
		}

		if checkCatalogsID != ctlgs[latestCatalog].ID {
			latestCatalog++
			latestKV = 0
			newCtlg := types.CatalogParameters{
				Parameters: []types.Parameter{
					types.Parameter{
						Key:        checkKey,
						CatalogsID: checkCatalogsID,
						Values: []types.Value{
							val,
						},
					},
				},
			}
			if err := rows.Scan(
				&newCtlg.ID,
				&newCtlg.Alias,
				&newCtlg.Img,
				&newCtlg.RuName,
				&newCtlg.AdditionDate,
				nil,
				&newCtlg.Parameters[0].RuKey,
				nil,
				nil,
				nil,
				nil,
				nil,
			); err != nil {
				return nil, fmt.Errorf("%s: %w", f, err)
			}
			ctlgs = append(ctlgs, newCtlg)
			continue
		}

		if checkKey != ctlgs[latestCatalog].Parameters[latestKV].Key {
			latestKV++
			newKV := types.Parameter{
				Key:        checkKey,
				CatalogsID: checkCatalogsID,
				Values: []types.Value{
					val,
				},
			}
			if err := rows.Scan(
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				&newKV.RuKey,
				nil,
				nil,
				nil,
				nil,
				nil,
			); err != nil {
				return nil, fmt.Errorf("%s: %w", f, err)
			}
			ctlgs[latestCatalog].Parameters = append(ctlgs[latestCatalog].Parameters, newKV)
			continue
		}

		ctlgs[latestCatalog].Parameters[latestKV].Values = append(ctlgs[latestCatalog].Parameters[latestKV].Values, val)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("%s: %w", f, err)
	}

	return ctlgs, nil
}
