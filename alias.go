package faircoin

import (
	"https://github.com/sarojchintakrindi/faircoin/tree/master/types"
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