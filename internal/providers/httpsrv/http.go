package httpsrv

import "github.com/deweppro/core/pkg/server/http"

type HTTPModule struct {
	cfg     *ConfigHttp
	httpsrv *http.Server
}

func NewHTTPModule(cfg *ConfigHttp) *HTTPModule {
	return &HTTPModule{
		cfg:     cfg,
		httpsrv: http.NewServer(),
	}
}

func (h *HTTPModule) Inject(i http.ModuleInjecter) {
	h.httpsrv.AddRoute(i)
}

func (h *HTTPModule) Up() error {
	h.httpsrv.SetAddr(h.cfg.Http.Addr)

	return h.httpsrv.Up()
}

func (h *HTTPModule) Down() error {
	return h.httpsrv.Down()
}
