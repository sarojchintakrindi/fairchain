/** 
Modules initialized and configred into complete app using sdk.ModuleBasicManager
**/



package app 

import (
	"encoding/json"
	"os"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"

	"github.com/cosmos/cosmos-sdk/x/auth/genaccounts"

	"github.com/cosmos/cosmos-sdk/x/bank"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/staking"
	//"github.com/cosmos/sdk-application-tutorial/x/nameservice"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db"
)

const appName = "faircoin"

var (
	//Might need to change path 

    // default directories for app CLI
	DefaultCLIHome = os.ExpandEnv("$HOME/.nscli")

	// app data and config storage
	DefaultNodeHome = os.ExpandEnv("$HOME/.nsd")

	ModuleBasics sdk.ModuleBasicManager
)

type fairCoinApp struct{
	*bam.BaseApp
}

func NewFairCoinApp(logger log.Logger, db dbm.DB) *fairCoinApp{
	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

	var app = &fairCoinApp{
		BaseApp: bApp,
		cdc:     cdc,
	}


}