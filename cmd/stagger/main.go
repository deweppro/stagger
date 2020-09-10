package main

import (
	"flag"

	"stagger/internal/modules/pingpong"
	"stagger/internal/providers/httpsrv"

	"github.com/deweppro/core/pkg/db/sqlite"
	"github.com/deweppro/core/pkg/debug"
	"github.com/deweppro/go-app/pkg/app"
)

var cfile = flag.String("config", "/etc/stagger/config.yaml", "path to config file")
var pidfile = flag.String("pid", "/var/run/stagger.pid", "path to pid file")

func main() {
	flag.Parse()

	app.
		New(*cfile).
		ConfigModels(
			// providers
			&debug.ConfigDebug{},
			&sqlite.ConfigSQLite{},
			// modules
			&httpsrv.ConfigHttp{},
		).
		Modules(
			// providers
			debug.New,
			sqlite.MustNew,
			// modules
			httpsrv.NewHTTPModule,
			pingpong.New,
		).
		PidFile(*pidfile).
		Run()

}
