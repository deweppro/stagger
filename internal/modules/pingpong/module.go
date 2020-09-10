package pingpong

import (
	"stagger/internal/providers/httpsrv"
	http2 "net/http"

	"github.com/deweppro/core/pkg/provider/server/http"
)

type PingPong struct{}

func New(srv *httpsrv.HTTPModule) *PingPong {
	pp := &PingPong{}
	srv.Inject(pp)
	return pp
}

func (pp *PingPong) Handlers() []http.CallHandler {
	return []http.CallHandler{
		{Method: http2.MethodGet, Path: "/ping", Call: pp.Pong},
	}
}

func (pp *PingPong) Formatter() http.FMT {
	return http.HTTPFormatter
}

func (pp *PingPong) Middelware() http.FN {
	return func(message *http.Message) error {
		return nil
	}
}

func (pp *PingPong) Up() error {
	return nil
}

func (pp *PingPong) Down() error {
	return nil
}
