package pingpong

import (
	"github.com/deweppro/core/pkg/provider/server/http"
)

func (pp *PingPong) Pong(message *http.Message) error {
	message.Encode(func() (int, map[string]string, interface{}) {
		return 200, nil, "pong"
	})

	return nil
}
