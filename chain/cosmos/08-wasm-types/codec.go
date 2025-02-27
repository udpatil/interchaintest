package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"
)

// RegisterInterfaces registers the tendermint concrete client-related
// implementations and interfaces.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*exported.ClientState)(nil),
		&ClientState{},
	)
	registry.RegisterImplementations(
		(*exported.ConsensusState)(nil),
		&ConsensusState{},
	)
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgStoreCode{},
	)
	registry.RegisterImplementations(
		(*exported.ClientMessage)(nil),
		&Misbehaviour{},
	)
	registry.RegisterImplementations(
		(*exported.ClientMessage)(nil),
		&Header{},
	)
}
