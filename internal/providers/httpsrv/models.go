package httpsrv

type ConfigHttp struct {
	Http ConfigHttpData `yaml:"http"`
}

type ConfigHttpData struct {
	Addr string `yaml:"addr"`
}
