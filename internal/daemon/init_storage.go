package daemon

import (
	"fmt"

	"github.com/ulibaysya/krona/internal/config"
	"github.com/ulibaysya/krona/internal/storage"
	"github.com/ulibaysya/krona/internal/storage/cachebased"
	"github.com/ulibaysya/krona/internal/storage/postgres"
)

// func newStorage(cfg config.Storage) (storage.Storage, error) {
// 	const f = "github.com/ulibaysya/krona/internal/storage.newStorage"
//
// 	var err error
//
// 	var rdbms storage.Storage
// 	switch cfg.RDBMS.Engine {
// 	case "postgres":
// 		rdbms, err = postgres.New(cfg.RDBMS)
// 	default:
// 		return nil, fmt.Errorf("%s: %w", f, storage.NewBadEngine(cfg.RDBMS.Engine))
// 	}
// 	if err != nil {
// 		return nil, fmt.Errorf("%s: %w", f, err)
// 	}
//
// 	var cache storage.Storage
// 	switch cfg.Cache.Engine {
// 	case "":
// 		return rdbms, nil
// 	case "valkey":
// 		_ = cache // TODO implement valkey cache :)
// 		_ = err
// 	default:
// 		return nil, fmt.Errorf("%s: %w", f, storage.NewBadEngine(cfg.Cache.Engine))
// 	}
// 	if err != nil {
// 		return nil, fmt.Errorf("%s: %w", f, err)
// 	}
//
// 	strg, err := cachebased.New(rdbms, cache)
// 	if err != nil {
// 		return nil, fmt.Errorf("%s: %w", f, err)
// 	}
//
// 	return strg, nil
// }

func newStorage(cfg config.Storage) (storage.Storage, error) {
	const f = "github.com/ulibaysya/krona/internal/storage.newStorage"

	switch cfg.Method {
	case "rdbms":
		rdbms, err := newRDBMS(cfg)
		if err != nil {
			return nil, err
		}
		return rdbms, nil
	case "cachebased":
		rdbms, err := newRDBMS(cfg)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", f, err)
		}
		cache, err := newCache(cfg)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", f, err)
		}
		strg, err := cachebased.New(rdbms, cache)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", f, err)
		}
		return strg, nil
	default:
		return nil, fmt.Errorf("%s: %w", f, storage.NewBadMethod(cfg.Method))
	}
}

func newRDBMS(cfg config.Storage) (storage.Storage, error) {
	switch cfg.RDBMS.Engine {
	case "postgres":
		return postgres.New(cfg.RDBMS)
	default:
		return nil, storage.NewBadEngine(cfg.RDBMS.Engine)
	}
}

func newCache(cfg config.Storage) (storage.Storage, error) {
	switch cfg.Cache.Engine {
	// case "valkey":
		// return postgres.New(cfg.RDBMS)
	default:
		return nil, fmt.Errorf(`cache is not yet implemented. use "rdbms" method`)
	}
}
