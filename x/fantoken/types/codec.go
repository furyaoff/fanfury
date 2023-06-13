package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgIssue{}, "go-fanfury/fantoken/MsgIssue", nil)
	cdc.RegisterConcrete(&MsgMint{}, "go-fanfury/fantoken/MsgMint", nil)
	cdc.RegisterConcrete(&MsgBurn{}, "go-fanfury/fantoken/MsgBurn", nil)
	cdc.RegisterConcrete(&MsgDisableMint{}, "go-fanfury/fantoken/MsgDisableMint", nil)
	cdc.RegisterConcrete(&MsgSetAuthority{}, "go-fanfury/fantoken/MsgSetAuthority", nil)
	cdc.RegisterConcrete(&MsgSetMinter{}, "go-fanfury/fantoken/MsgSetMinter", nil)
	cdc.RegisterConcrete(&MsgSetUri{}, "go-fanfury/fantoken/MsgSetUri", nil)
	cdc.RegisterConcrete(&UpdateFeesProposal{}, "go-fanfury/fantoken/UpdateFeesProposal", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgIssue{},
		&MsgMint{},
		&MsgBurn{},
		&MsgDisableMint{},
		&MsgSetAuthority{},
		&MsgSetMinter{},
		&MsgSetUri{},
	)

	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&UpdateFeesProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
