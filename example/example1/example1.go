// Package example1 demonstrates basic usage of the eth-interfaces library.
package example1

import (
	"log"
	"math/big"
	"os"

	"github.com/Thektonic/eth-interfaces/base"
	"github.com/Thektonic/eth-interfaces/nft"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

// Example1 demonstrates basic NFT contract interactions including connection and transfer
func Example1() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to get environment")
	}

	// Connect to the client
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal("Error getting the client: ", err)
	}

	// Get the private key from the .env file and convert it to an ECDSA key
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal("error reading the PRIVATE_KEY env var: ", err)
	}

	// Defer client close after all potential fatal errors
	defer client.Close()

	// Create a new base interaction object
	baseInteractions := base.NewBaseInteractions(client, privateKey, nil)

	// Create a new ERC721 interaction object from the base interaction
	nftInteractions, err := nft.NewERC721Interactions(
		baseInteractions,
		common.HexToAddress("0"),
		[]nft.BaseNFTSignature{nft.Name, nft.Symbol, nft.TokenURI, nft.TransferFrom},
	)
	if err != nil {
		log.Printf("error creating the NFT interactions: %v", err)
		return
	}

	// Transfer NFT to another address
	_, err = nftInteractions.TransferTo(common.HexToAddress("0"), big.NewInt(0))
	if err != nil {
		log.Printf("error transferring the NFT: %v", err)
		return
	}
}
