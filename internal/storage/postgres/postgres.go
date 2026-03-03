package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ulibaysya/krona/internal/config"
	"github.com/ulibaysya/krona/internal/storage/types"
)

type postgres struct {
	pool *pgxpool.Pool
}

func New(cfg config.RDBMS) (postgres, error) {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, cfg.Connstr)
	if err != nil {
		return postgres{}, wrapError("initializing pool", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return postgres{}, wrapError("pinging pool", err)
	}

	return postgres{pool: pool}, nil
}

func queryAndCollect[T any](ctx context.Context, db postgres, sql string, fn pgx.RowToFunc[T]) ([]T, error) {
	rows, err := db.pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}

	var objects []T
	if fn == nil {
		objects, err = pgx.CollectRows(rows, pgx.RowToStructByPos[T]) // TODO maybe we should force caller to pass pgx.RowToStructByPos[T], but if nil - return error?
	} else {
		objects, err = pgx.CollectRows(rows, fn)
	}
	if err != nil {
		return nil, err
	}

	return objects, nil
}

func scanParameter(row pgx.CollectableRow) (types.ParameterRow, error) {
	var param types.ParameterRow
	// if err := row.Scan(&param.ID, &param.Key, &param.Value, &param.RuKey, &param.RuValue, &param.CatalogsID, &param.AdditionDate); err != nil {
	if err := row.Scan(nil, nil, nil, nil, nil, &param.ID, &param.Key, &param.Value, &param.RuKey, &param.RuValue, &param.CatalogsID, &param.AdditionDate); err != nil {
		return types.ParameterRow{}, err
	}

	return param, nil
}

func wrapError(msg string, err error) error {
	if err == nil {
		return fmt.Errorf("postgres: %s", msg)
	}
	return fmt.Errorf("postgres: %s: %w", msg, err)
}

func isUniqueViolation(err error) bool {
	pgErr, ok := errors.AsType[*pgconn.PgError](err)
	if ok {
		if pgErr.Code == pgerrcode.UniqueViolation {
			return true
		}
	}
	// var pgErr *pgconn.PgError
	// if errors.As(err, &pgErr) {
	// 	if pgErr.Code == pgerrcode.UniqueViolation {
	// 		return true
	// 	}
	// }
	return false
}
