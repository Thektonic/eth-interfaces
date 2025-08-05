package merged_test

import (
	"testing"

	"github.com/Thektonic/eth-interfaces/base"
	"github.com/Thektonic/eth-interfaces/hex"
	"github.com/Thektonic/eth-interfaces/inferences/ERC721A"
	"github.com/Thektonic/eth-interfaces/inferences/ERC721Complete"
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
				abiString:      ERC721A.ERC721AABI,
				byteCodeString: ERC721A.ERC721ABin,
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
				abiString:      ERC721Complete.ERC721CompleteABI,
				byteCodeString: ERC721Complete.ERC721CompleteBin,
				extensions:     []merged.ExtensionEnum{merged.Royalties},
				signatures:     []hex.Signature{royalties.RoyaltyInfo},
			},
		},
		{
			Name: "OK - Instantiate NFT with royalties extension and enumerable extension",
			Args: args{
				abiString:      ERC721Complete.ERC721CompleteABI,
				byteCodeString: ERC721Complete.ERC721CompleteBin,
				extensions:     []merged.ExtensionEnum{merged.Royalties, merged.Enumerable},
				signatures: []hex.Signature{
					royalties.RoyaltyInfo,
					enumerable.TokenByIndex,
					enumerable.TokenOfOwnerByIndex,
				},
			},
		},
		{
			Name: "NOK - Instantiate NFT with enumerable and royalties extensions",
			Args: args{
				abiString:      ERC721A.ERC721AABI,
				byteCodeString: ERC721A.ERC721ABin,
				extensions:     []merged.ExtensionEnum{merged.Enumerable, merged.Royalties},
				signatures:     []hex.Signature{enumerable.TokenOfOwnerByIndex, royalties.RoyaltyInfo},
			},
			ExpectError:   true,
			ExpectedError: "not supported functions: royaltyInfo(uint256,uint256)",
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

			baseInteractions := base.NewBaseInteractions(backend.Client(), privKey, nil)

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
