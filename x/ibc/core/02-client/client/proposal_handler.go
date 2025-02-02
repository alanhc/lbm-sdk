package client

import (
	"net/http"

	"github.com/line/lbm-sdk/client"
	"github.com/line/lbm-sdk/types/rest"
	govclient "github.com/line/lbm-sdk/x/gov/client"
	govrest "github.com/line/lbm-sdk/x/gov/client/rest"

	"github.com/line/lbm-sdk/x/ibc/core/02-client/client/cli"
)

var (
	UpdateClientProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpdateClientProposal, emptyRestHandler)
	UpgradeProposalHandler      = govclient.NewProposalHandler(cli.NewCmdSubmitUpgradeProposal, emptyRestHandler)
)

func emptyRestHandler(client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "unsupported-ibc-client",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Legacy REST Routes are not supported for IBC proposals")
		},
	}
}
