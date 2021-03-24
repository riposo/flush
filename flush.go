package flush

import (
	"net/http"

	"github.com/riposo/flush/internal"
	"github.com/riposo/riposo/pkg/api"
	"github.com/riposo/riposo/pkg/plugin"
)

func init() {
	plugin.Register("flush_endpoint", func(rts *api.Routes) (plugin.Plugin, error) {
		rts.Method(http.MethodPost, "/__flush__", api.HandlerFunc(internal.Handler))

		return pin{
			"description": "The __flush__ endpoint can be used to remove all data from all backends.",
			"url":         "https://github.com/riposo/flush",
		}, nil
	})
}

type pin map[string]interface{}

func (p pin) Meta() map[string]interface{} { return map[string]interface{}(p) }
func (pin) Close() error                   { return nil }
