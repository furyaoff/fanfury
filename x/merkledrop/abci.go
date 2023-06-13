package merkledrop

import (
	"github.com/furyaoff/fanfury/x/merkledrop/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	logger := keeper.Logger(ctx)

	merkledropIDs := keeper.GetMerkleDropsIDByEndHeight(ctx, ctx.BlockHeight())

	for _, merkledropID := range merkledropIDs {
		keeper.Withdraw(ctx, merkledropID)
		keeper.DeleteMerkledropByID(ctx, merkledropID)
		logger.Info("merkledrop deleted", "mdID", merkledropID)
	}
}
