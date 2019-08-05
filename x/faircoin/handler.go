package faircoin

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "faircoin" type messages.

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSendCoins:
			return handleMsgSendCoins(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgSendCoins(ctx sdk.Context, keeper Keeper, msg MsgSendCoins) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Incorrect Owner").Result() // If not, throw an error
	}
	keeper.SetName(ctx, msg.Name, msg.Value) // If so, set the name to the value specified in the msg.
	return sdk.Result{}                      // return
}


func handleMsgSendCoins(ctx sdk.Context, keeper Keeper, msg MsgSendCoins) sdk.Result {

	err := keeper.coinKeeper.SendCoins(ctx, msg.Sender, msg.Recip, msg.Value)

	if err != nil {
			return sdk.ErrInsufficientCoins("Sender does not have enough coins").Result()
		}

	return sdk.Result{}
}