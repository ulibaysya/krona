package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/ulibaysya/krona/internal/config"
	"github.com/ulibaysya/krona/internal/daemon/service"
)

type Server struct {
	addr, net string
	serv *http.Server
}

func New(cfg config.Server, srvc service.Service) (Server, error) {
	const f = "github.com/ulibaysya/krona/internal/daemon/server.New"

	mux := srvc.GetMux()
	serv := http.Server{Addr: cfg.Address, Handler: &mux}

	return Server{addr: cfg.Address, net: cfg.Network, serv: &serv}, nil
}

func (s Server) Serve() error {
	const f = "github.com/ulibaysya/krona/internal/daemon/server.Serve"

	l, err := net.Listen(s.net, s.addr)
	if err != nil {
		return fmt.Errorf("%s: %w", f, err)
	}
	//
	// if s.net == "unix" {
	// 	if err := os.Chmod(s.addr, 0666); err != nil {
	// 		return fmt.Errorf("%s: %w", f, err)
	// 	}
	// }

	if err := s.serv.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", f, err)
	}

	return nil
}
