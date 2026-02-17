package daemon

import (
	"fmt"

	"github.com/ulibaysya/krona/internal/config"
	"github.com/ulibaysya/krona/internal/daemon/server"
	"github.com/ulibaysya/krona/internal/daemon/service"
	"github.com/ulibaysya/krona/internal/log"
	"github.com/ulibaysya/krona/internal/storage"
)

type Daemon struct {
	cfg config.Config
	logger log.Logger
	strg storage.Storage
	srvc service.Service
	serv server.Server
}

func New(cfgPath string) (Daemon, error) {
	const f = "github.com/ulibaysya/krona/internal/daemon.New"

	var d Daemon
	var err error

	d.cfg, err = config.New(cfgPath)
	if err != nil {
		return Daemon{}, fmt.Errorf("%s: %w", f, err)
	}

	d.logger, err = log.New(d.cfg.Log)
	if err != nil {
		return Daemon{}, fmt.Errorf("%s: %w", f, err)
	}

	d.strg, err = newStorage(d.cfg.Storage)
	if err != nil {
		return Daemon{}, fmt.Errorf("%s: %w", f, err)
	}

	d.srvc, err = service.New(d.cfg.Service, d.logger, d.strg)
	if err != nil {
		return Daemon{}, fmt.Errorf("%s: %w", f, err)
	}

	d.serv, err = server.New(d.cfg.Server, d.srvc)
	if err != nil {
		return Daemon{}, fmt.Errorf("%s: %w", f, err)
	}

	return d, nil
}

func (d Daemon) Run() error {
	const f = "github.com/ulibaysya/krona/internal/daemon.Run"

	err := d.serv.Serve()
	if err != nil {
		return fmt.Errorf("%s: %w", f, err)
	}
	
	return nil
}

func (d Daemon) Shutdown() error {
	const f = "github.com/ulibaysya/krona/internal/daemon.Shutdown"

	if err := d.logger.Close(); err != nil {
		return fmt.Errorf("%s: %w", f, err)
	}

	return nil
}
