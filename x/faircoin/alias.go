package faircoin

import (
	"github.com/sarojchintakrindi/fairchain/tree/master/x/faircoin/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)



var (
	NewMsgSendCoins = types.NewMsgBuyName
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
)


type (
	MsgSendCoins     = types.MsgSetName
)