package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateTxPool{}, "sequencer/CreateTxPool", nil)
	cdc.RegisterConcrete(&MsgUpdateTxPool{}, "sequencer/UpdateTxPool", nil)
	cdc.RegisterConcrete(&MsgDeleteTxPool{}, "sequencer/DeleteTxPool", nil)
	cdc.RegisterConcrete(&MsgSendCommitment{}, "sequencer/SendCommitment", nil)
	cdc.RegisterConcrete(&MsgSendPayload{}, "sequencer/SendPayload", nil)
	cdc.RegisterConcrete(&MsgSendTlp{}, "sequencer/SendTlp", nil)
	cdc.RegisterConcrete(&MsgCloseRound{}, "sequencer/CloseRound", nil)
	cdc.RegisterConcrete(&MsgMakeBlock{}, "sequencer/MakeBlock", nil)
	cdc.RegisterConcrete(&MsgCreateBlock{}, "sequencer/CreateBlock", nil)
	cdc.RegisterConcrete(&MsgUpdateBlock{}, "sequencer/UpdateBlock", nil)
	cdc.RegisterConcrete(&MsgDeleteBlock{}, "sequencer/DeleteBlock", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTxPool{},
		&MsgUpdateTxPool{},
		&MsgDeleteTxPool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendCommitment{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendPayload{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendTlp{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCloseRound{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMakeBlock{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateBlock{},
		&MsgUpdateBlock{},
		&MsgDeleteBlock{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
