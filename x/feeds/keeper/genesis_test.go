package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bandprotocol/chain/v2/x/feeds/types"
)

func (suite *KeeperTestSuite) TestExportGenesis() {
	ctx := suite.ctx

	err := suite.feedsKeeper.SetParams(ctx, types.DefaultParams())
	suite.Require().NoError(err)

	err = suite.feedsKeeper.SetReferenceSourceConfig(ctx, types.DefaultReferenceSourceConfig())
	suite.Require().NoError(err)

	delegatorSignals := []types.DelegatorSignals{
		{
			Delegator: ValidDelegator.String(),
			Signals: []types.Signal{
				{
					ID:    "CS:BAND-USD",
					Power: 10000 * 1e6,
				},
				{
					ID:    "CS:BTC-USD",
					Power: 20000 * 1e9,
				},
			},
		},
		{
			Delegator: ValidDelegator2.String(),
			Signals: []types.Signal{
				{
					ID:    "CS:BAND-USD",
					Power: 20000 * 1e6,
				},
				{
					ID:    "CS:BTC-USD",
					Power: 40000 * 1e9,
				},
			},
		},
	}
	suite.feedsKeeper.SetAllDelegatorSignals(ctx, delegatorSignals)

	exportGenesis := suite.feedsKeeper.ExportGenesis(ctx)

	suite.Require().Equal(types.DefaultParams(), exportGenesis.Params)
	suite.Require().Equal(types.DefaultReferenceSourceConfig(), exportGenesis.ReferenceSourceConfig)
	suite.Require().Equal(delegatorSignals, exportGenesis.DelegatorSignals)
}

func (suite *KeeperTestSuite) TestInitGenesis() {
	ctx := suite.ctx
	params := types.NewParams("[NOT_SET]", 30, 30, 60, 3600, 1000_000_000, 100, 30, 5, 300, 28800)

	delegatorSignals := []types.DelegatorSignals{
		{
			Delegator: ValidDelegator.String(),
			Signals: []types.Signal{
				{
					ID:    "CS:BAND-USD",
					Power: 10000 * 1e6,
				},
				{
					ID:    "CS:BTC-USD",
					Power: 20000 * 1e6,
				},
			},
		},
		{
			Delegator: ValidDelegator2.String(),
			Signals: []types.Signal{
				{
					ID:    "CS:BAND-USD",
					Power: 20000 * 1e6,
				},
				{
					ID:    "CS:BTC-USD",
					Power: 40000 * 1e6,
				},
			},
		},
	}

	g := types.DefaultGenesisState()
	g.DelegatorSignals = delegatorSignals
	g.Params = params

	suite.feedsKeeper.InitGenesis(suite.ctx, *g)

	suite.Require().Equal(types.DefaultReferenceSourceConfig(), suite.feedsKeeper.GetReferenceSourceConfig(ctx))
	suite.Require().Equal(params, suite.feedsKeeper.GetParams(ctx))
	for _, ds := range delegatorSignals {
		suite.Require().
			Equal(ds.Signals, suite.feedsKeeper.GetDelegatorSignals(ctx, sdk.MustAccAddressFromBech32(ds.Delegator)))
	}

	stpBand, err := suite.feedsKeeper.GetSignalTotalPower(ctx, "CS:BAND-USD")
	suite.Require().NoError(err)
	suite.Require().Equal(types.Signal{
		ID:    "CS:BAND-USD",
		Power: 30000 * 1e6,
	}, stpBand)

	stpBtc, err := suite.feedsKeeper.GetSignalTotalPower(ctx, "CS:BTC-USD")
	suite.Require().NoError(err)
	suite.Require().Equal(types.Signal{
		ID:    "CS:BTC-USD",
		Power: 60000 * 1e6,
	}, stpBtc)

	suite.Require().Equal(types.CurrentFeeds{
		Feeds: []types.Feed{
			{
				SignalID: "CS:BTC-USD",
				Power:    60000000000,
				Interval: 60,
			},
			{
				SignalID: "CS:BAND-USD",
				Power:    30000000000,
				Interval: 120,
			},
		},
		LastUpdateTimestamp: ctx.BlockTime().Unix(),
		LastUpdateBlock:     ctx.BlockHeight(),
	}, suite.feedsKeeper.GetCurrentFeeds(ctx))
}
