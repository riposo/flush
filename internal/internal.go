package internal

import (
	"net/http"

	"github.com/riposo/riposo/pkg/api"
)

// Handler is an api.HandlerFunc.
func Handler(out http.Header, r *http.Request) interface{} {
	txn := api.GetTxn(r)

	if err := txn.Store.Flush(); err != nil {
		return err
	}
	if err := txn.Perms.Flush(); err != nil {
		return err
	}
	if err := txn.Cache.Flush(); err != nil {
		return err
	}

	return responseType{}
}

type responseType struct{}

func (responseType) HTTPStatus() int { return http.StatusAccepted }
