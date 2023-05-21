package client_test

import (
	"fmt"
	"testing"

	"github.com/bandprotocol/chain/v2/cylinder/client"
	"github.com/bandprotocol/chain/v2/pkg/tss"
	"github.com/bandprotocol/chain/v2/x/tss/types"
	"github.com/stretchr/testify/assert"
)

func TestGetRound1Data(t *testing.T) {
	tests := []struct {
		name               string
		queryGroupResponse *types.QueryGroupResponse
		memberID           tss.MemberID
		expectedData       types.Round1Data
		expectedError      error
	}{
		{
			name: "Existing MemberID",
			queryGroupResponse: &types.QueryGroupResponse{
				AllRound1Data: []types.Round1Data{
					{
						MemberID: 1,
					},
				},
			},
			memberID:      1,
			expectedData:  types.Round1Data{MemberID: 1},
			expectedError: nil,
		},
		{
			name: "No data from MemberID",
			queryGroupResponse: &types.QueryGroupResponse{
				AllRound1Data: []types.Round1Data{
					{
						MemberID: 1,
					},
				},
			},
			memberID:      2,
			expectedData:  types.Round1Data{},
			expectedError: fmt.Errorf("No Round1Data from MemberID(2)"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			groupResponse := client.NewGroupResponse(test.queryGroupResponse)
			data, err := groupResponse.GetRound1Data(test.memberID)
			assert.Equal(t, test.expectedError, err)
			assert.Equal(t, test.expectedData, data)
		})
	}
}

func TestGetRound2Shares(t *testing.T) {
	tests := []struct {
		name               string
		queryGroupResponse *types.QueryGroupResponse
		memberID           tss.MemberID
		expectedShare      *types.Round2Share
		expectedError      error
	}{
		{
			name: "Existing Member ID",
			queryGroupResponse: &types.QueryGroupResponse{
				Round2Shares: []types.Round2Share{
					{},
				},
			},
			memberID:      1,
			expectedShare: &types.Round2Share{},
			expectedError: nil,
		},
		{
			name: "Invalid Member ID",
			queryGroupResponse: &types.QueryGroupResponse{
				Round2Shares: []types.Round2Share{
					{},
				},
			},
			memberID:      2,
			expectedShare: nil,
			expectedError: fmt.Errorf("No Round2Shares from MemberID(2)"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			groupResponse := client.NewGroupResponse(test.queryGroupResponse)

			share, err := groupResponse.GetRound2Shares(test.memberID)

			assert.Equal(t, test.expectedError, err)
			assert.Equal(t, test.expectedShare, share)
		})
	}
}

func TestGetEncryptedSecretShare(t *testing.T) {
	tests := []struct {
		name               string
		queryGroupResponse *types.QueryGroupResponse
		memberIDJ          tss.MemberID
		memberIDI          tss.MemberID
		expectedShare      tss.Scalar
		expectedError      error
	}{
		{
			name: "Existing Member IDs",
			queryGroupResponse: &types.QueryGroupResponse{
				Round2Shares: []types.Round2Share{
					{
						EncryptedSecretShares: []tss.Scalar{[]byte("share1"), []byte("share2")},
					},
				},
			},
			memberIDJ:     1,
			memberIDI:     1,
			expectedShare: []byte("share1"),
			expectedError: nil,
		},
		{
			name: "Invalid Member J ID",
			queryGroupResponse: &types.QueryGroupResponse{
				Round2Shares: []types.Round2Share{
					{
						EncryptedSecretShares: []tss.Scalar{[]byte("share1"), []byte("share2")},
					},
				},
			},
			memberIDJ:     2,
			memberIDI:     1,
			expectedShare: nil,
			expectedError: fmt.Errorf("No Round2Shares from MemberID(2)"),
		},
		{
			name: "Invalid Member I ID",
			queryGroupResponse: &types.QueryGroupResponse{
				Round2Shares: []types.Round2Share{
					{
						EncryptedSecretShares: []tss.Scalar{[]byte("share1"), []byte("share2")},
					},
				},
			},
			memberIDJ:     1,
			memberIDI:     4,
			expectedShare: nil,
			expectedError: fmt.Errorf("No Round2Shares from MemberID(1) for MemberID(4)"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			groupResponse := client.NewGroupResponse(test.queryGroupResponse)

			share, err := groupResponse.GetEncryptedSecretShare(test.memberIDJ, test.memberIDI)

			assert.Equal(t, test.expectedError, err)
			assert.Equal(t, test.expectedShare, share)
		})
	}
}
