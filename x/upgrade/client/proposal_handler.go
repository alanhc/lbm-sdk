package client

import (
	govclient "github.com/line/lbm-sdk/x/gov/client"
	"github.com/line/lbm-sdk/x/upgrade/client/cli"
	"github.com/line/lbm-sdk/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
var CancelProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitCancelUpgradeProposal, rest.ProposalCancelRESTHandler)
