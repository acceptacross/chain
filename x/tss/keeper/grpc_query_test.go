package keeper_test

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/bandprotocol/chain/v2/pkg/tss"
	"github.com/bandprotocol/chain/v2/x/tss/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

func (s *KeeperTestSuite) TestGRPCQueryGroup() {
	ctx, msgSrvr, q, k := s.ctx, s.msgSrvr, s.querier, s.app.TSSKeeper
	groupID, memberID1, memberID2 := tss.GroupID(1), tss.MemberID(1), tss.MemberID(2)
	members := []string{
		"band18gtd9xgw6z5fma06fxnhet7z2ctrqjm3z4k7ad",
		"band1s743ydr36t6p29jsmrxm064guklgthsn3t90ym",
		"band1p08slm6sv2vqy4j48hddkd6hpj8yp6vlw3pf8p",
		"band1p08slm6sv2vqy4j48hddkd6hpj8yp6vlw3pf8p",
		"band12jf07lcaj67mthsnklngv93qkeuphhmxst9mh8",
	}
	round1Data1 := types.Round1Data{
		MemberID: memberID1,
		CoefficientsCommit: []tss.Point{
			[]byte("point1"),
			[]byte("point2"),
			[]byte("point3"),
		},
		OneTimePubKey: []byte("OneTimePubKeySample"),
		A0Sig:         []byte("A0SigSample"),
		OneTimeSig:    []byte("OneTimeSigSample"),
	}
	round1Data2 := types.Round1Data{
		MemberID: memberID2,
		CoefficientsCommit: []tss.Point{
			[]byte("point1"),
			[]byte("point2"),
			[]byte("point3"),
		},
		OneTimePubKey: []byte("OneTimePubKeySample"),
		A0Sig:         []byte("A0SigSample"),
		OneTimeSig:    []byte("OneTimeSigSample"),
	}
	round2Data1 := types.Round2Data{
		MemberID: memberID1,
		EncryptedSecretShares: tss.Scalars{
			[]byte("scalar1"),
			[]byte("scalar2"),
		},
	}
	round2Data2 := types.Round2Data{
		MemberID: memberID2,
		EncryptedSecretShares: tss.Scalars{
			[]byte("scalar1"),
			[]byte("scalar2"),
		},
	}
	complainWithStatus1 := types.ComplainsWithStatus{
		MemberID: memberID1,
		ComplainsWithStatus: []types.ComplainWithStatus{
			{
				Complain: types.Complain{
					I:         1,
					J:         2,
					KeySym:    []byte("key_sym"),
					Signature: []byte("signature"),
					NonceSym:  []byte("nonce_sym"),
				},
				ComplainStatus: types.SUCCESS,
			},
		},
	}
	complainWithStatus2 := types.ComplainsWithStatus{
		MemberID: memberID2,
		ComplainsWithStatus: []types.ComplainWithStatus{
			{
				Complain: types.Complain{
					I:         1,
					J:         2,
					KeySym:    []byte("key_sym"),
					Signature: []byte("signature"),
					NonceSym:  []byte("nonce_sym"),
				},
				ComplainStatus: types.SUCCESS,
			},
		},
	}
	confirm1 := types.Confirm{
		MemberID:     memberID1,
		OwnPubKeySig: []byte("own_pub_key_sig"),
	}
	confirm2 := types.Confirm{
		MemberID:     memberID2,
		OwnPubKeySig: []byte("own_pub_key_sig"),
	}

	msgSrvr.CreateGroup(ctx, &types.MsgCreateGroup{
		Members:   members,
		Threshold: 3,
		Sender:    members[0],
	})

	// Set round1 data
	k.SetRound1Data(ctx, groupID, round1Data1)
	k.SetRound1Data(ctx, groupID, round1Data2)

	// Set round 2 data
	k.SetRound2Data(ctx, groupID, round2Data1)
	k.SetRound2Data(ctx, groupID, round2Data2)

	// Set complain
	k.SetComplainsWithStatus(ctx, groupID, complainWithStatus1)
	k.SetComplainsWithStatus(ctx, groupID, complainWithStatus2)

	// Set confirm
	k.SetConfirm(ctx, groupID, confirm1)
	k.SetConfirm(ctx, groupID, confirm2)

	var req types.QueryGroupRequest
	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
		postTest func(res *types.QueryGroupResponse)
	}{
		{
			"non existing group",
			func() {
				req = types.QueryGroupRequest{
					GroupId: 2,
				}
			},
			false,
			func(res *types.QueryGroupResponse) {},
		},
		{
			"success",
			func() {
				req = types.QueryGroupRequest{
					GroupId: uint64(groupID),
				}
			},
			true,
			func(res *types.QueryGroupResponse) {
				dkgContextB, _ := hex.DecodeString("a1cdd234702bbdbd8a4fa9fc17f2a83d569f553ae4bd1755985e5039532d108c")

				s.Require().Equal(&types.QueryGroupResponse{
					Group: types.Group{
						Size_:     5,
						Threshold: 3,
						PubKey:    nil,
						Status:    types.ROUND_1,
					},
					DKGContext: dkgContextB,
					Members: []types.Member{
						{
							Member:      "band18gtd9xgw6z5fma06fxnhet7z2ctrqjm3z4k7ad",
							PubKey:      tss.PublicKey(nil),
							IsMalicious: false,
						},
						{
							Member:      "band1s743ydr36t6p29jsmrxm064guklgthsn3t90ym",
							PubKey:      tss.PublicKey(nil),
							IsMalicious: false,
						},
						{
							Member:      "band1p08slm6sv2vqy4j48hddkd6hpj8yp6vlw3pf8p",
							PubKey:      tss.PublicKey(nil),
							IsMalicious: false,
						},
						{
							Member:      "band1p08slm6sv2vqy4j48hddkd6hpj8yp6vlw3pf8p",
							PubKey:      tss.PublicKey(nil),
							IsMalicious: false,
						},
						{
							Member:      "band12jf07lcaj67mthsnklngv93qkeuphhmxst9mh8",
							PubKey:      tss.PublicKey(nil),
							IsMalicious: false,
						},
					},
					AllRound1Data: []types.Round1Data{
						round1Data1,
						round1Data2,
					},
					AllRound2Data: []types.Round2Data{
						round2Data1,
						round2Data2,
					},
					AllComplainsWithStatus: []types.ComplainsWithStatus{
						complainWithStatus1,
						complainWithStatus2,
					},
					AllConfirm: []types.Confirm{
						confirm1,
						confirm2,
					},
				}, res)
			},
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			res, err := q.Group(ctx, &req)
			if tc.expPass {
				s.Require().NoError(err)
			} else {
				s.Require().Error(err)
			}

			tc.postTest(res)
		})
	}
}

func (s *KeeperTestSuite) TestGRPCQueryMembers() {
	ctx, q, k := s.ctx, s.querier, s.app.TSSKeeper
	members := []types.Member{
		{
			Member:      "band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs",
			PubKey:      tss.PublicKey(nil),
			IsMalicious: false,
		},
		{
			Member:      "band1p40yh3zkmhcv0ecqp3mcazy83sa57rgjp07dun",
			PubKey:      tss.PublicKey(nil),
			IsMalicious: false,
		},
	}

	// set members
	for i, m := range members {
		k.SetMember(ctx, tss.GroupID(1), tss.MemberID(i+1), m)
	}

	var req types.QueryMembersRequest
	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
		postTest func(res *types.QueryMembersResponse)
	}{
		{
			"non existing member",
			func() {
				req = types.QueryMembersRequest{
					GroupId: 2,
				}
			},
			false,
			func(res *types.QueryMembersResponse) {},
		},
		{
			"success",
			func() {
				req = types.QueryMembersRequest{
					GroupId: 1,
				}
			},
			true,
			func(res *types.QueryMembersResponse) {
				s.Require().Equal(members, res.Members)
			},
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			_, err := q.Members(ctx, &req)
			if tc.expPass {
				s.Require().NoError(err)
			} else {
				s.Require().Error(err)
			}
		})
	}
}

func (s *KeeperTestSuite) TestGRPCQueryIsGrantee() {
	ctx, q, authzKeeper := s.ctx, s.querier, s.app.AuthzKeeper
	expTime := time.Unix(0, 0)

	// Init grantee address
	grantee, _ := sdk.AccAddressFromBech32("band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs")

	// Init granter address
	granter, _ := sdk.AccAddressFromBech32("band1p40yh3zkmhcv0ecqp3mcazy83sa57rgjp07dun")

	// Save grant msgs to grantee
	for _, m := range types.MsgGrants {
		authzKeeper.SaveGrant(s.ctx, grantee, granter, authz.NewGenericAuthorization(m), &expTime)
	}

	var req types.QueryIsGranteeRequest
	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
		postTest func(res *types.QueryIsGranteeResponse)
	}{
		{
			"address is not bech32",
			func() {
				req = types.QueryIsGranteeRequest{
					Granter: "asdsd1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs",
					Grantee: "padads40yh3zkmhcv0ecqp3mcazy83sa57rgjp07dun",
				}
			},
			false,
			func(res *types.QueryIsGranteeResponse) {},
		},
		{
			"address is empty",
			func() {
				req = types.QueryIsGranteeRequest{
					Granter: "",
					Grantee: "",
				}
			},
			false,
			func(res *types.QueryIsGranteeResponse) {},
		},
		{
			"grantee address is not grant by granter",
			func() {
				req = types.QueryIsGranteeRequest{
					Granter: "band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs",
					Grantee: "band17eplw6tga7wqgruqdtalw3rky4njkx6vngxjlt",
				}
			},
			true,
			func(res *types.QueryIsGranteeResponse) {
				s.Require().False(res.IsGrantee)
			},
		},
		{
			"success",
			func() {
				req = types.QueryIsGranteeRequest{
					Granter: "band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs",
					Grantee: "band1p40yh3zkmhcv0ecqp3mcazy83sa57rgjp07dun",
				}
			},
			true,
			func(res *types.QueryIsGranteeResponse) {
				s.Require().True(res.IsGrantee)
			},
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			_, err := q.IsGrantee(ctx, &req)
			if tc.expPass {
				s.Require().NoError(err)
			} else {
				s.Require().Error(err)
			}
		})
	}
}
