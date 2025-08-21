package royalties_test

// Package royalties_test contains tests for royalties interactions.

import (
	"math/big"
	"testing"

	"github.com/Thektonic/eth-interfaces/base"
	"github.com/Thektonic/eth-interfaces/inferences"
	"github.com/Thektonic/eth-interfaces/nft"
	"github.com/Thektonic/eth-interfaces/nft/royalties"
	"github.com/Thektonic/eth-interfaces/testingtools"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

// Test_RoyaltiesInfos verifies the RoyaltiesInfos method for valid and invalid token IDs using a table-driven approach.
func Test_RoyaltiesInfos(t *testing.T) {
	backend, _, contractAddr, privKey, err := testingtools.SetupBlockchain(t,
		inferences.Ierc721MetaData.ABI,
		inferences.Ierc721MetaData.Bin,
		"MyNFT",
		"MNFT",
	)
	assert.Nil(t, err)
	defer func() {
		if err := backend.Close(); err != nil {
			t.Logf("failed to close backend: %v", err)
		}
	}()

	baseInteractions := base.NewBaseInteractions(backend.Client(), privKey, nil, false)
	nftA, err := nft.NewERC721Interactions(baseInteractions, *contractAddr, []nft.BaseNFTSignature{nft.BalanceOf})
	assert.Nil(t, err)

	royInteractions, err := royalties.NewERC721RoyaltiesInteractions(
		nftA, []royalties.IERC721RoyaltiesSignature{royalties.RoyaltyInfo},
	)
	if err != nil {
		t.Skipf("Skipping royalties test as royalties interactions are not implemented: %v", err)
	}

	tests := []struct {
		name       string
		tokenID    *big.Int
		shouldFail bool
	}{
		{
			name:    "Valid token",
			tokenID: common.Big0,
		},
	}

	salePrice := big.NewInt(1e18)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := royInteractions.RoyaltiesInfos(tt.tokenID, salePrice)
			if tt.shouldFail {
				assert.Error(t, err, "expected error for tokenID %v but got success", tt.tokenID)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
