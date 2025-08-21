// Package base provides core utilities for interacting with the blockchain,
// including transaction setup, contract calls, and error handling.
package base

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/Thektonic/eth-interfaces/customerrors"
	"github.com/Thektonic/eth-interfaces/hex"
	"github.com/Thektonic/eth-interfaces/inferences"
	"github.com/Thektonic/eth-interfaces/transaction"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
)

// Interactions holds the context, client, sender address, private key, disperse contract, and explorer URL.
type Interactions struct {
	Ctx      context.Context
	Client   simulated.Client
	Address  common.Address
	pk       *ecdsa.PrivateKey
	disperse *inferences.Disperse
	explorer *string
	TxOptsFn transaction.TxOptsBuilderFunc
	safe     bool
}

// Session holds call options and bound contract instance for contract interactions
type Session struct {
	callOpts *bind.CallOpts
	instance *bind.BoundContract
}

// CallOpts returns the call options for contract calls
func (s *Session) CallOpts() *bind.CallOpts {
	return s.callOpts
}

// Instance returns the bound contract instance
func (s *Session) Instance() *bind.BoundContract {
	return s.instance
}

// Safe returns whether the interactions are in safe mode
func (i *Interactions) Safe() bool {
	return i.safe
}

// IBaseInteractions defines the interface for verifying transactions.
type IBaseInteractions interface {
	VerifyTransaction(ctx context.Context, to common.Address, data []byte, value int64) error
}

// NewBaseInteractions creates a new instance of BaseInteractions for blockchain interaction.
func NewBaseInteractions(
	client simulated.Client,
	pk *ecdsa.PrivateKey,
	explorer *string,
	safe bool,
	txOptsFn ...transaction.TxOptsBuilderFunc,
) *Interactions {
	ctx := context.TODO()
	_, err := client.BlockNumber(ctx)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	var txOptFn transaction.TxOptsBuilderFunc

	if len(txOptsFn) != 0 {
		txOptFn = txOptsFn[0]
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return &Interactions{ctx, client, fromAddress, pk, nil, explorer, txOptFn, safe}
}

// SetDisperse initializes the disperse contract for multi-address fund transfers.
func (i *Interactions) SetDisperse(_ string) error {
	var err error
	i.disperse = inferences.NewDisperse()
	return err
}

// BaseTxSetup sets up transaction options (nonce, gas price, chain ID, etc.) for sending a transaction.
func (i *Interactions) BaseTxSetup() (*bind.TransactOpts, error) {
	if i.TxOptsFn != nil {
		return i.TxOptsFn()
	}

	gasPrice, err := i.Client.SuggestGasPrice(i.Ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to suggest gas price: %v", err)
	}
	nonce, err := i.Client.PendingNonceAt(i.Ctx, i.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to get user nonce: %v", err)
	}

	chainID, err := i.Client.ChainID(i.Ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %v", err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(i.pk, chainID)
	if err != nil {
		return nil, err
	}

	opts.From = i.Address
	opts.GasPrice = gasPrice

	opts.From = i.Address
	opts.Nonce = new(big.Int).SetUint64(nonce)

	return opts, nil
}

// BaseCallSetup returns the call options for read-only contract operations.
func (i *Interactions) BaseCallSetup() *bind.CallOpts {
	return &bind.CallOpts{
		From:    i.Address,
		Pending: false,
	}
}

// CatchTx waits for a transaction to be mined and returns its hash or an error message.
func (i *Interactions) CatchTx(tx *ethTypes.Transaction, err error) (string, error) {
	if err != nil {
		return FailedTx(err)
	}
	receipt, err := bind.WaitMined(context.Background(), i.Client, tx)
	if receipt == nil {
		return FailedTx(err)
	}
	if i.explorer != nil {
		return SuccessTx(fmt.Sprintf(*i.explorer, "/tx/", receipt.TxHash.Hex()))
	}
	return SuccessTx(receipt.TxHash.Hex())
}

// VerifyTransaction simulates a contract call to verify transaction validity.
func (i *Interactions) VerifyTransaction(ctx context.Context, to common.Address, data []byte, value int64) error {
	callMsg := ethereum.CallMsg{
		From:  i.Address,
		To:    &to,
		Data:  data,
		Value: big.NewInt(value),
	}

	_, err := i.Client.CallContract(ctx, callMsg, nil)
	return err
}

// Disperse uses the disperse contract to send funds to multiple addresses.
func (i *Interactions) Disperse(addresses []common.Address, totalValue uint) (string, error) {
	if i.disperse == nil {
		return FailedTx(fmt.Errorf("disperse contract not initialized"))
	}
	opts, err := i.BaseTxSetup()
	if err != nil {
		return FailedTx(err)
	}
	amounts := []*big.Int{}
	for range addresses {
		amounts = append(amounts, new(big.Int).SetUint64(uint64(totalValue)/uint64(len(addresses))))
	}

	tmpValue := *opts.Value

	opts.Value = new(big.Int).SetUint64(uint64(totalValue))

	i.TxOptsFn = func() (*bind.TransactOpts, error) {
		return opts, nil
	}

	fmt.Println("Dispersing...")

	instance := i.disperse.Instance(i.Client, i.Address)

	tx, err := transaction.Transact(
		i,
		&Session{callOpts: i.BaseCallSetup(), instance: instance},
		i.disperse.PackDisperseEther(addresses, amounts), transaction.DefaultUnpacker,
	)

	if err != nil {
		return FailedTx(err)
	}

	opts.Value = &tmpValue

	return i.CatchTx(tx, err)
}

// SendAllFunds transfers the entire balance to a designated address after fee estimation.
func (i *Interactions) SendAllFunds(to common.Address) (*ethTypes.Transaction, error) {
	bn, err := i.Client.BlockNumber(i.Ctx)
	if err != nil {
		return nil, err
	}
	balance, err := i.Client.BalanceAt(i.Ctx, i.Address, new(big.Int).SetUint64(bn))
	if err != nil {
		return nil, err
	}

	msg := ethereum.CallMsg{
		From:  i.Address,
		To:    &to,
		Value: balance,
		Data:  nil,
	}

	gasLimit, err := i.Client.EstimateGas(i.Ctx, msg)
	if err != nil {
		return nil, err
	}

	gasPrice, err := i.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	gasCost := new(big.Int).SetUint64(gasLimit)
	gasCost.Mul(gasCost, gasPrice)
	value := big.NewInt(0).Add(balance, gasCost)
	if value.Int64() < balance.Int64() {
		return i.TransferETH(to, value)
	}
	return nil, fmt.Errorf(
		"fees exceed balances\nfees : %f ETH\nbalance : %f ETH",
		hex.ParseEther(big.NewInt(0).Sub(balance, value)),
		hex.ParseEther(balance),
	)
}

// TransferETH transfers Ether to the specified address, ensuring sufficient balance and proper fee estimation.
func (i *Interactions) TransferETH(to common.Address, value *big.Int) (*ethTypes.Transaction, error) {
	balance, err := i.Client.BalanceAt(i.Ctx, to, nil)
	if err != nil {
		return nil, err
	}

	msg := ethereum.CallMsg{From: i.Address, To: &to, Value: balance, Data: nil}

	gasLimit, err := i.Client.EstimateGas(i.Ctx, msg)
	if err != nil {
		return nil, err
	}

	gasPrice, err := i.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	gasCost := new(big.Int).SetUint64(gasLimit)
	gasCost.Mul(gasCost, gasPrice)
	txCost := big.NewInt(0).Add(value, gasCost)
	if txCost.Uint64() < balance.Uint64() {
		return nil, fmt.Errorf(
			"unsufficient balance for the transfer\n value + fees : %f ETH\nbalance : %f ETH",
			hex.ParseEther(txCost),
			hex.ParseEther(balance),
		)
	}

	nonce, err := i.Client.PendingNonceAt(context.Background(), i.Address)
	if err != nil {
		return nil, err
	}

	// Create the transaction
	tx := ethTypes.NewTx(&ethTypes.LegacyTx{
		Nonce:    nonce,
		To:       &to,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
	})
	// Get the chain ID
	chainID, err := i.Client.ChainID(i.Ctx)
	if err != nil {
		return nil, err
	}

	// Sign the transaction
	signedTx, err := ethTypes.SignTx(tx, ethTypes.NewEIP155Signer(chainID), i.pk)
	if err != nil {
		return nil, fmt.Errorf("failed to sign the tx: %w", err)
	}

	// Broadcast the transaction
	err = i.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to send the tx: %w", err)
	}

	return signedTx, nil
}

// CheckSignatures checks if a contract supports specific function signatures.
func (i *Interactions) CheckSignatures(contractAddress common.Address, signatures []hex.Signature) error {
	// Get proxy bytecode
	byteCode, err := i.Client.CodeAt(i.Ctx, contractAddress, nil)
	if err != nil {
		return fmt.Errorf("failed to get contract bytecode: %w", err)
	}
	byteCodeHex := common.Bytes2Hex(byteCode)

	// Check signatures against proxy bytecode first
	var notFoundSigs []hex.Signature
	for _, signature := range signatures {
		sigHex := signature.GetHex()
		if !strings.Contains(byteCodeHex, sigHex[:8]) {
			notFoundSigs = append(notFoundSigs, signature)
		}
	}

	// If all signatures found, return success
	if len(notFoundSigs) == 0 {
		return nil
	}

	// Check implementation contract if any signatures not found
	implAddress, err := hex.GetImplementationAddress(i.Ctx, i.Client, contractAddress)
	if err != nil {
		return customerrors.WrapInterfacingError("CheckSignatures", err)
	}

	notSupported := ""

	// If implementation found, check remaining sigs against it
	if implAddress != (common.Address{}) {
		implCode, err := i.Client.CodeAt(i.Ctx, implAddress, nil)
		if err != nil {
			// If can't get implementation code, mark all remaining as not supported
			for _, sig := range notFoundSigs {
				notSupported += fmt.Sprintf("%s: %s\n", sig, sig.GetHex()[:8])
			}
		} else {
			implCodeHex := common.Bytes2Hex(implCode)
			for _, sig := range notFoundSigs {
				sigHex := sig.GetHex()
				if !strings.Contains(implCodeHex, sigHex[:8]) {
					notSupported += fmt.Sprintf("%s: %s\n", sig, sigHex[:8])
				}
			}
		}
	} else {
		// Check diamond facets for remaining signatures
		for _, sig := range notFoundSigs {
			supported, err := hex.CheckDiamondFunction(i.Ctx, i.Client, contractAddress, sig.GetSelector())
			if err != nil || !supported {
				notSupported += fmt.Sprintf("%s: %s\n", sig, sig.GetHex()[:8])
			}
		}
	}

	if len(notSupported) > 0 {
		return customerrors.WrapInterfacingError("CheckSignatures", fmt.Errorf("not supported functions: %s", notSupported))
	}
	return nil
}

// ManageCustomContractError handles custom contract errors.
func (i *Interactions) ManageCustomContractError(abiString string, err error) error {
	if len(abiString) == 0 {
		return err
	}
	errBytes, success := ethclient.RevertErrorData(err)
	if success {
		customErr, err := i.MatchErrors(abiString, errBytes)
		if err != nil {
			return customerrors.WrapInterfacingError("MatchErrors", err)
		}
		return errors.New(customErr)
	}
	return nil
}

// MatchErrors matches error bytes to custom error names.
func (i *Interactions) MatchErrors(abiString string, errBytes []byte) (string, error) {
	abi, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		return "", err
	}
	if len(errBytes) < hex.ErrorMethodIDLength {
		panic("invalid error data")
	}
	methodID := errBytes[:hex.ErrorMethodIDLength]

	// Find matching custom error
	var errorName string
	for _, abiError := range abi.Errors {
		abierr := abiError.ID.Bytes()[:4]
		if bytes.Equal(abierr[:4], methodID) {
			errorName = abiError.Name
			break
		}
	}
	return errorName, nil
}
