package flush

import (
	"net/http"

	"github.com/riposo/flush/internal"
	"github.com/riposo/riposo/pkg/api"
	"github.com/riposo/riposo/pkg/plugin"
)

func init() {
	plugin.Register(plugin.New(
		"flush",
		map[string]interface{}{
			"description": "The __flush__ endpoint can be used to remove all data from all backends.",
			"url":         "https://github.com/riposo/flush",
		},
		func(rts *api.Routes) error {
			rts.Method(http.MethodPost, "/__flush__", api.HandlerFunc(internal.Handler))
			return nil
		},
		nil,
	))
}
