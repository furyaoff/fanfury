package merkledrop

import (
	"github.com/furyaoff/fanfury/x/merkledrop/keeper"
	"github.com/furyaoff/fanfury/x/merkledrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// NewHandler handles all messages.
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgCreate:
			res, err := msgServer.Create(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgClaim:
			res, err := msgServer.Claim(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized merkledrop message type: %T", msg)
		}
	}
}

func NewProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.UpdateFeesProposal:
			return handleUpdateFeesProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized merkledrop proposal content type: %T", c)
		}
	}
}

func handleUpdateFeesProposal(ctx sdk.Context, k keeper.Keeper, p *types.UpdateFeesProposal) error {
	ctx.Logger().Info("Updating fantoken fees from proposal")

	if err := p.CreationFee.Validate(); err != nil {
		return err
	}

	params := k.GetParamSet(ctx)
	params.CreationFee = p.CreationFee
	k.SetParamSet(ctx, params)

	return nil
}
