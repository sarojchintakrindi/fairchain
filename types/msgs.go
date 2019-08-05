package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName // from key.go


type MsgSendCoins struct {
	Value string `json:"value"`
	Sender sdk.AccAddress `json: "sender"`
}

func NewMsgSendCoins(recip sdk.AccAddress, value string,  sender sdk.AccAddress) MsgSendCoins {
	return MsgSendCoins{
		Value: value,
		Sender: sender,
		Recip: recip,
	}
}


//route returns name of module
func (msg MsgSendCoins) Route() string { return RouterKey }

//type returns action
func (msg MsgSendCoins) Type() string { return "send_coins" }

func (msg MsgSendCoins) ValidateBasic() sdk.Error {
	if msg.sender.Empty() || msg.recip.Empty() {
		return sdk.ErrInvalidAddress(msg.Recip.String())
	}
	if len(msg.Value) == 0 {
		return sdk.ErrUnknownRequest("Value cannot be empty")
	}
	return nil
}

// encodes for signing
func (msg MsgSendCoins) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgSendCoins) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.sender}
}



