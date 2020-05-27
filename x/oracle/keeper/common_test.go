package keeper_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	bandapp "github.com/bandprotocol/bandchain/chain/app"
	"github.com/bandprotocol/bandchain/chain/pkg/filecache"
	"github.com/bandprotocol/bandchain/chain/simapp"
	me "github.com/bandprotocol/bandchain/chain/x/oracle/keeper"
	"github.com/bandprotocol/bandchain/chain/x/oracle/types"
	connectiontypes "github.com/cosmos/cosmos-sdk/x/ibc/03-connection/types"
	channelexported "github.com/cosmos/cosmos-sdk/x/ibc/04-channel/exported"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/04-channel/types"
	ibctmtypes "github.com/cosmos/cosmos-sdk/x/ibc/07-tendermint/types"
	commitmenttypes "github.com/cosmos/cosmos-sdk/x/ibc/23-commitment/types"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
)

const (
	ChainID         = "bandchain"
	ChainIDA        = "chainA"
	ChainIDB        = "chainB"
	TestClientIDA   = "clientA"
	TestClientIDB   = "clientB"
	TestPortA       = "testporta"
	TestPortB       = "testportb"
	TestChannelA    = "testchannela"
	TestChannelB    = "testchannelb"
	TestConnectionA = "connectionAtoB"
	TestConnectionB = "connectionBtoA"
)

var (
	Owner      = simapp.Owner
	Alice      = simapp.Alice
	Bob        = simapp.Bob
	Carol      = simapp.Carol
	Validator1 = simapp.Validator1
	Validator2 = simapp.Validator2
	Validator3 = simapp.Validator3
)

var (
	BasicName          = "BASIC_NAME"
	BasicDesc          = "BASIC_DESCRIPTION"
	BasicSchema        = "BASIC_SCHEMA"
	BasicSourceCodeURL = "BASIC_SOURCE_CODE_URL"
	BasicFilename      = "BASIC_FILENAME"
	BasicCalldata      = []byte("BASIC_CALLDATA")
	CoinsZero          = sdk.NewCoins()
	Coins10uband       = sdk.NewCoins(sdk.NewInt64Coin("uband", 10))
	Coins20uband       = sdk.NewCoins(sdk.NewInt64Coin("uband", 20))
)

func createTestInput() (*bandapp.BandApp, sdk.Context, me.Keeper) {
	app := simapp.NewSimApp(ChainID, log.NewNopLogger())
	ctx := app.BaseApp.NewContext(false, abci.Header{})
	return app, ctx, app.OracleKeeper
}

func createTestChains(logger log.Logger) (*bandapp.BandApp, *bandapp.BandApp) {
	appA := simapp.NewSimApp(ChainIDA, logger)
	appB := simapp.NewSimApp(ChainIDB, logger)
	return appA, appB
}

func getContext(chain *bandapp.BandApp) sdk.Context {
	now := time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)

	privVal := tmtypes.NewMockPV()
	signers := []tmtypes.PrivValidator{privVal}
	pubKey, err := privVal.GetPubKey()
	if err != nil {
		panic(err)
	}
	validator := tmtypes.NewValidator(pubKey, 1)
	valSet := tmtypes.NewValidatorSet([]*tmtypes.Validator{validator})
	header := ibctmtypes.CreateTestHeader(chain.Name(), 1, now, valSet, signers)

	return chain.NewContext(false, abci.Header{
		ChainID: header.ChainID,
		Height:  header.Height,
	})
}

func createTestChainConnection(chainA *bandapp.BandApp, chainB *bandapp.BandApp) {
	counterParty := connectiontypes.NewCounterparty(TestConnectionA, TestConnectionB, commitmenttypes.NewMerklePrefix(chainA.IBCKeeper.ConnectionKeeper.GetCommitmentPrefix().Bytes()))
	conn := connectiontypes.NewConnectionEnd(3, TestClientIDB, counterParty, connectiontypes.GetCompatibleVersions())
	ctx := chainA.NewContext(false, abci.Header{})
	chainA.IBCKeeper.ConnectionKeeper.SetConnection(ctx, TestConnectionA, conn)
}

func createTestChannel(chainA *bandapp.BandApp, chainB *bandapp.BandApp) {
	counterpart := channeltypes.NewCounterparty(TestPortB, TestChannelB)
	channel := channeltypes.NewChannel(channelexported.OPEN, channelexported.ORDERED, counterpart, []string{TestConnectionA}, "1.0")
	ctx := chainA.NewContext(false, abci.Header{})
	chainA.IBCKeeper.ChannelKeeper.SetChannel(ctx, TestPortA, TestChannelA, channel)
}

func newDefaultRequest() types.Request {
	return types.NewRequest(
		1,
		[]byte("calldata"),
		[]sdk.ValAddress{Validator1.ValAddress, Validator2.ValAddress},
		2,
		0,
		1581503227,
		"clientID",
		nil,
		[]types.ExternalID{42},
	)
}

func deleteFile(path string) {
	err := os.Remove(path)
	if err != nil {
		panic(err)
	}
}

func getTestDataSource() (ds types.DataSource, clear func()) {
	dir, err := ioutil.TempDir("/tmp", "filecache")
	if err != nil {
		panic(err)
	}
	f := filecache.New(dir)
	filename := f.AddFile([]byte("executable"))
	return types.NewDataSource(Owner.Address, "Test data source", "For test only", filename),
		func() { deleteFile(filepath.Join(dir, filename)) }
}

func getTestOracleScript() (os types.OracleScript, clear func()) {
	dir, err := ioutil.TempDir("/tmp", "filecache")
	if err != nil {
		panic(err)
	}
	f := filecache.New(dir)
	filename := f.AddFile([]byte("code"))
	return types.NewOracleScript(Owner.Address, "Test oracle script",
		"For test only", filename, "", "test URL",
	), func() { deleteFile(filepath.Join(dir, filename)) }
}

func loadOracleScriptFromWasm(path string) (os types.OracleScript, clear func()) {
	absPath, _ := filepath.Abs(path)
	code, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	dir, err := ioutil.TempDir("/tmp", "filecache")
	if err != nil {
		panic(err)
	}
	f := filecache.New(dir)
	filename := f.AddFile(code)
	return types.NewOracleScript(
		Owner.Address, "imported script", "description",
		filename, "schema", "sourceCodeURL",
	), func() { deleteFile(filepath.Join(dir, filename)) }
}
