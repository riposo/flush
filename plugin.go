package main

import (
	"net/http"

	"github.com/riposo/flush/internal"
	"github.com/riposo/riposo/pkg/api"
	"github.com/riposo/riposo/pkg/plugin"
)

var _ plugin.Factory = Plugin

// Plugin export definition.
func Plugin(rts *api.Routes) (plugin.Plugin, error) {
	rts.Method(http.MethodPost, "/__flush__", api.HandlerFunc(internal.Handler))

	return plugin.New(
		"flush_endpoint",
		map[string]interface{}{
			"description": "The __flush__ endpoint can be used to remove all data from all backends.",
			"url":         "https://github.com/riposo/flush",
		},
		nil,
	), nil
}

func main() {}
