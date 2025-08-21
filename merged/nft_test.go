package merged_test

import (
	"testing"

	"github.com/Thektonic/eth-interfaces/base"
	"github.com/Thektonic/eth-interfaces/hex"
	"github.com/Thektonic/eth-interfaces/inferences"
	"github.com/Thektonic/eth-interfaces/merged"
	"github.com/Thektonic/eth-interfaces/nft"
	"github.com/Thektonic/eth-interfaces/nft/enumerable"
	"github.com/Thektonic/eth-interfaces/nft/royalties"
	"github.com/Thektonic/eth-interfaces/testingtools"
	"github.com/stretchr/testify/assert"
)

// TestAllInfosSuccess tests the successful execution of AllInfos with valid dummyBase and both extensions enabled.
func Test_Instantiation(t *testing.T) {
	type args struct {
		abiString      string
		byteCodeString string
		extensions     []merged.ExtensionEnum
		signatures     []hex.Signature
	}

	testCases := []struct {
		Name          string
		Args          args
		ExpectError   bool
		ExpectedError string
	}{
		{
			Name: "OK - Instantiate NFT with enumerable extension",
			Args: args{
				abiString:      inferences.Ierc721MetaData.ABI,
				byteCodeString: inferences.Ierc721MetaData.Bin,
				extensions:     []merged.ExtensionEnum{merged.Enumerable},
				signatures: []hex.Signature{
					enumerable.TokenByIndex,
					enumerable.TokenOfOwnerByIndex,
				},
			},
		},
		{
			Name: "OK - Instantiate NFT with royalties extension",
			Args: args{
				abiString:      inferences.Ierc721MetaData.ABI,
				byteCodeString: inferences.Ierc721MetaData.Bin,
				extensions:     []merged.ExtensionEnum{merged.Royalties},
				signatures:     []hex.Signature{royalties.RoyaltyInfo},
			},
		},
		{
			Name: "OK - Instantiate NFT with royalties extension and enumerable extension",
			Args: args{
				abiString:      inferences.Ierc721MetaData.ABI,
				byteCodeString: inferences.Ierc721MetaData.Bin,
				extensions:     []merged.ExtensionEnum{merged.Royalties, merged.Enumerable},
				signatures: []hex.Signature{
					royalties.RoyaltyInfo,
					enumerable.TokenByIndex,
					enumerable.TokenOfOwnerByIndex,
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			backend, _, contractAddr, privKey, err := testingtools.SetupBlockchain(t,
				tt.Args.abiString,
				tt.Args.byteCodeString,
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

			nftA, err := nft.NewERC721Interactions(baseInteractions, *contractAddr, []nft.BaseNFTSignature{})
			assert.Nil(t, err)

			// Create a new summed interactions with both Enumerable and Royalties extensions
			_, err = merged.NewERC721SummedInteractions(nftA, tt.Args.signatures, tt.Args.extensions...)
			if tt.ExpectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.ExpectedError)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
