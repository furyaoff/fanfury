package client

import (
	"github.com/furyaoff/fanfury/x/fantoken/client/cli"
	"github.com/cosmos/cosmos-sdk/client"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govrest "github.com/cosmos/cosmos-sdk/x/gov/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdUpdateFantokenFees, ProposalRESTHandler)

func ProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{}
}
