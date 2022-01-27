package client

import (
	govclient "github.com/MonOsmosis/cosmos-sdk/x/gov/client"
	"github.com/MonOsmosis/cosmos-sdk/x/upgrade/client/cli"
	"github.com/MonOsmosis/cosmos-sdk/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
var CancelProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitCancelUpgradeProposal, rest.ProposalCancelRESTHandler)
