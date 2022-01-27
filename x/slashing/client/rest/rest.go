package rest

import (
	"github.com/gorilla/mux"

	"github.com/MonOsmosis/cosmos-sdk/client/rest"

	"github.com/MonOsmosis/cosmos-sdk/client"
)

func RegisterHandlers(clientCtx client.Context, rtr *mux.Router) {
	r := rest.WithHTTPDeprecationHeaders(rtr)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)
}
