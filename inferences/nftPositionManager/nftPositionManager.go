// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nftPositionManager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// INonfungiblePositionManagerCollectParams is an auto generated low-level Go binding around an user-defined struct.
type INonfungiblePositionManagerCollectParams struct {
	TokenId    *big.Int
	Recipient  common.Address
	Amount0Max *big.Int
	Amount1Max *big.Int
}

// INonfungiblePositionManagerDecreaseLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type INonfungiblePositionManagerDecreaseLiquidityParams struct {
	TokenId    *big.Int
	Liquidity  *big.Int
	Amount0Min *big.Int
	Amount1Min *big.Int
	Deadline   *big.Int
}

// INonfungiblePositionManagerIncreaseLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type INonfungiblePositionManagerIncreaseLiquidityParams struct {
	TokenId        *big.Int
	Amount0Desired *big.Int
	Amount1Desired *big.Int
	Amount0Min     *big.Int
	Amount1Min     *big.Int
	Deadline       *big.Int
}

// INonfungiblePositionManagerMintParams is an auto generated low-level Go binding around an user-defined struct.
type INonfungiblePositionManagerMintParams struct {
	Token0         common.Address
	Token1         common.Address
	Fee            *big.Int
	TickLower      *big.Int
	TickUpper      *big.Int
	Amount0Desired *big.Int
	Amount1Desired *big.Int
	Amount0Min     *big.Int
	Amount1Min     *big.Int
	Recipient      common.Address
	Deadline       *big.Int
}

// NftPositionManagerMetaData contains all meta data concerning the NftPositionManager contract.
var NftPositionManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Collect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"DecreaseLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"IncreaseLiquidity\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount0Max\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Max\",\"type\":\"uint128\"}],\"internalType\":\"structINonfungiblePositionManager.CollectParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structINonfungiblePositionManager.DecreaseLiquidityParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"decreaseLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structINonfungiblePositionManager.IncreaseLiquidityParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"increaseLiquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint256\",\"name\":\"amount0Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structINonfungiblePositionManager.MintParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"nonce\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside0LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside1LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NftPositionManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use NftPositionManagerMetaData.ABI instead.
var NftPositionManagerABI = NftPositionManagerMetaData.ABI

// NftPositionManager is an auto generated Go binding around an Ethereum contract.
type NftPositionManager struct {
	NftPositionManagerCaller     // Read-only binding to the contract
	NftPositionManagerTransactor // Write-only binding to the contract
	NftPositionManagerFilterer   // Log filterer for contract events
}

// NftPositionManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NftPositionManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftPositionManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NftPositionManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftPositionManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NftPositionManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftPositionManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NftPositionManagerSession struct {
	Contract     *NftPositionManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// NftPositionManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NftPositionManagerCallerSession struct {
	Contract *NftPositionManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// NftPositionManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NftPositionManagerTransactorSession struct {
	Contract     *NftPositionManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// NftPositionManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NftPositionManagerRaw struct {
	Contract *NftPositionManager // Generic contract binding to access the raw methods on
}

// NftPositionManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NftPositionManagerCallerRaw struct {
	Contract *NftPositionManagerCaller // Generic read-only contract binding to access the raw methods on
}

// NftPositionManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NftPositionManagerTransactorRaw struct {
	Contract *NftPositionManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNftPositionManager creates a new instance of NftPositionManager, bound to a specific deployed contract.
func NewNftPositionManager(address common.Address, backend bind.ContractBackend) (*NftPositionManager, error) {
	contract, err := bindNftPositionManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NftPositionManager{NftPositionManagerCaller: NftPositionManagerCaller{contract: contract}, NftPositionManagerTransactor: NftPositionManagerTransactor{contract: contract}, NftPositionManagerFilterer: NftPositionManagerFilterer{contract: contract}}, nil
}

// NewNftPositionManagerCaller creates a new read-only instance of NftPositionManager, bound to a specific deployed contract.
func NewNftPositionManagerCaller(address common.Address, caller bind.ContractCaller) (*NftPositionManagerCaller, error) {
	contract, err := bindNftPositionManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NftPositionManagerCaller{contract: contract}, nil
}

// NewNftPositionManagerTransactor creates a new write-only instance of NftPositionManager, bound to a specific deployed contract.
func NewNftPositionManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NftPositionManagerTransactor, error) {
	contract, err := bindNftPositionManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NftPositionManagerTransactor{contract: contract}, nil
}

// NewNftPositionManagerFilterer creates a new log filterer instance of NftPositionManager, bound to a specific deployed contract.
func NewNftPositionManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NftPositionManagerFilterer, error) {
	contract, err := bindNftPositionManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NftPositionManagerFilterer{contract: contract}, nil
}

// bindNftPositionManager binds a generic wrapper to an already deployed contract.
func bindNftPositionManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NftPositionManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NftPositionManager *NftPositionManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NftPositionManager.Contract.NftPositionManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NftPositionManager *NftPositionManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftPositionManager.Contract.NftPositionManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NftPositionManager *NftPositionManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NftPositionManager.Contract.NftPositionManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NftPositionManager *NftPositionManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NftPositionManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NftPositionManager *NftPositionManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftPositionManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NftPositionManager *NftPositionManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NftPositionManager.Contract.contract.Transact(opts, method, params...)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 tokenId) view returns(uint96 nonce, address operator, address token0, address token1, uint24 fee, int24 tickLower, int24 tickUpper, uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_NftPositionManager *NftPositionManagerCaller) Positions(opts *bind.CallOpts, tokenId *big.Int) (struct {
	Nonce                    *big.Int
	Operator                 common.Address
	Token0                   common.Address
	Token1                   common.Address
	Fee                      *big.Int
	TickLower                *big.Int
	TickUpper                *big.Int
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	var out []interface{}
	err := _NftPositionManager.contract.Call(opts, &out, "positions", tokenId)

	outstruct := new(struct {
		Nonce                    *big.Int
		Operator                 common.Address
		Token0                   common.Address
		Token1                   common.Address
		Fee                      *big.Int
		TickLower                *big.Int
		TickUpper                *big.Int
		Liquidity                *big.Int
		FeeGrowthInside0LastX128 *big.Int
		FeeGrowthInside1LastX128 *big.Int
		TokensOwed0              *big.Int
		TokensOwed1              *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Nonce = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Operator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Token0 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TickLower = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TickUpper = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Liquidity = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside0LastX128 = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside1LastX128 = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed0 = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed1 = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 tokenId) view returns(uint96 nonce, address operator, address token0, address token1, uint24 fee, int24 tickLower, int24 tickUpper, uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_NftPositionManager *NftPositionManagerSession) Positions(tokenId *big.Int) (struct {
	Nonce                    *big.Int
	Operator                 common.Address
	Token0                   common.Address
	Token1                   common.Address
	Fee                      *big.Int
	TickLower                *big.Int
	TickUpper                *big.Int
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _NftPositionManager.Contract.Positions(&_NftPositionManager.CallOpts, tokenId)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 tokenId) view returns(uint96 nonce, address operator, address token0, address token1, uint24 fee, int24 tickLower, int24 tickUpper, uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_NftPositionManager *NftPositionManagerCallerSession) Positions(tokenId *big.Int) (struct {
	Nonce                    *big.Int
	Operator                 common.Address
	Token0                   common.Address
	Token1                   common.Address
	Fee                      *big.Int
	TickLower                *big.Int
	TickUpper                *big.Int
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _NftPositionManager.Contract.Positions(&_NftPositionManager.CallOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) payable returns()
func (_NftPositionManager *NftPositionManagerTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _NftPositionManager.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) payable returns()
func (_NftPositionManager *NftPositionManagerSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _NftPositionManager.Contract.Burn(&_NftPositionManager.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) payable returns()
func (_NftPositionManager *NftPositionManagerTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _NftPositionManager.Contract.Burn(&_NftPositionManager.TransactOpts, tokenId)
}

// Collect is a paid mutator transaction binding the contract method 0xfc6f7865.
//
// Solidity: function collect((uint256,address,uint128,uint128) params) payable returns(uint256 amount0, uint256 amount1)
func (_NftPositionManager *NftPositionManagerTransactor) Collect(opts *bind.TransactOpts, params INonfungiblePositionManagerCollectParams) (*types.Transaction, error) {
	return _NftPositionManager.contract.Transact(opts, "collect", params)
}

// Collect is a paid mutator transaction binding the contract method 0xfc6f7865.
//
// Solidity: function collect((uint256,address,uint128,uint128) params) payable returns(uint256 amount0, uint256 amount1)
func (_NftPositionManager *NftPositionManagerSession) Collect(params INonfungiblePositionManagerCollectParams) (*types.Transaction, error) {
	return _NftPositionManager.Contract.Collect(&_NftPositionManager.TransactOpts, params)
}

// Collect is a paid mutator transaction binding the contract method 0xfc6f7865.
//
// Solidity: function collect((uint256,address,uint128,uint128) params) payable returns(uint256 amount0, uint256 amount1)
func (_NftPositionManager *NftPositionManagerTransactorSession) Collect(params INonfungiblePositionManagerCollectParams) (*types.Transaction, error) {
	return _NftPositionManager.Contract.Collect(&_NftPositionManager.TransactOpts, params)
}

// DecreaseLiquidity is a paid mutator transaction binding the contract method 0x0c49ccbe.
//
// Solidity: function decreaseLiquidity((uint256,uint128,uint256,uint256,uint256) params) payable returns(uint256 amount0, uint256 amount1)
func (_NftPositionManager *NftPositionManagerTransactor) DecreaseLiquidity(opts *bind.TransactOpts, params INonfungiblePositionManagerDecreaseLiquidityParams) (*types.Transaction, error) {
	return _NftPositionManager.contract.Transact(opts, "decreaseLiquidity", params)
}

// DecreaseLiquidity is a paid mutator transaction binding the contract method 0x0c49ccbe.
//
// Solidity: function decreaseLiquidity((uint256,uint128,uint256,uint256,uint256) params) payable returns(uint256 amount0, uint256 amount1)
func (_NftPositionManager *NftPositionManagerSession) DecreaseLiquidity(params INonfungiblePositionManagerDecreaseLiquidityParams) (*types.Transaction, error) {
	return _NftPositionManager.Contract.DecreaseLiquidity(&_NftPositionManager.TransactOpts, params)
}

// DecreaseLiquidity is a paid mutator transaction binding the contract method 0x0c49ccbe.
//
// Solidity: function decreaseLiquidity((uint256,uint128,uint256,uint256,uint256) params) payable returns(uint256 amount0, uint256 amount1)
func (_NftPositionManager *NftPositionManagerTransactorSession) DecreaseLiquidity(params INonfungiblePositionManagerDecreaseLiquidityParams) (*types.Transaction, error) {
	return _NftPositionManager.Contract.DecreaseLiquidity(&_NftPositionManager.TransactOpts, params)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0x219f5d17.
//
// Solidity: function increaseLiquidity((uint256,uint256,uint256,uint256,uint256,uint256) params) payable returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NftPositionManager *NftPositionManagerTransactor) IncreaseLiquidity(opts *bind.TransactOpts, params INonfungiblePositionManagerIncreaseLiquidityParams) (*types.Transaction, error) {
	return _NftPositionManager.contract.Transact(opts, "increaseLiquidity", params)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0x219f5d17.
//
// Solidity: function increaseLiquidity((uint256,uint256,uint256,uint256,uint256,uint256) params) payable returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NftPositionManager *NftPositionManagerSession) IncreaseLiquidity(params INonfungiblePositionManagerIncreaseLiquidityParams) (*types.Transaction, error) {
	return _NftPositionManager.Contract.IncreaseLiquidity(&_NftPositionManager.TransactOpts, params)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0x219f5d17.
//
// Solidity: function increaseLiquidity((uint256,uint256,uint256,uint256,uint256,uint256) params) payable returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NftPositionManager *NftPositionManagerTransactorSession) IncreaseLiquidity(params INonfungiblePositionManagerIncreaseLiquidityParams) (*types.Transaction, error) {
	return _NftPositionManager.Contract.IncreaseLiquidity(&_NftPositionManager.TransactOpts, params)
}

// Mint is a paid mutator transaction binding the contract method 0x88316456.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,uint256,uint256,address,uint256) params) payable returns(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NftPositionManager *NftPositionManagerTransactor) Mint(opts *bind.TransactOpts, params INonfungiblePositionManagerMintParams) (*types.Transaction, error) {
	return _NftPositionManager.contract.Transact(opts, "mint", params)
}

// Mint is a paid mutator transaction binding the contract method 0x88316456.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,uint256,uint256,address,uint256) params) payable returns(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NftPositionManager *NftPositionManagerSession) Mint(params INonfungiblePositionManagerMintParams) (*types.Transaction, error) {
	return _NftPositionManager.Contract.Mint(&_NftPositionManager.TransactOpts, params)
}

// Mint is a paid mutator transaction binding the contract method 0x88316456.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,uint256,uint256,address,uint256) params) payable returns(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NftPositionManager *NftPositionManagerTransactorSession) Mint(params INonfungiblePositionManagerMintParams) (*types.Transaction, error) {
	return _NftPositionManager.Contract.Mint(&_NftPositionManager.TransactOpts, params)
}

// NftPositionManagerCollectIterator is returned from FilterCollect and is used to iterate over the raw logs and unpacked data for Collect events raised by the NftPositionManager contract.
type NftPositionManagerCollectIterator struct {
	Event *NftPositionManagerCollect // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftPositionManagerCollectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftPositionManagerCollect)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftPositionManagerCollect)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftPositionManagerCollectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftPositionManagerCollectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftPositionManagerCollect represents a Collect event raised by the NftPositionManager contract.
type NftPositionManagerCollect struct {
	TokenId   *big.Int
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollect is a free log retrieval operation binding the contract event 0x40d0efd1a53d60ecbf40971b9daf7dc90178c3aadc7aab1765632738fa8b8f01.
//
// Solidity: event Collect(uint256 indexed tokenId, address recipient, uint256 amount0, uint256 amount1)
func (_NftPositionManager *NftPositionManagerFilterer) FilterCollect(opts *bind.FilterOpts, tokenId []*big.Int) (*NftPositionManagerCollectIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NftPositionManager.contract.FilterLogs(opts, "Collect", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftPositionManagerCollectIterator{contract: _NftPositionManager.contract, event: "Collect", logs: logs, sub: sub}, nil
}

// WatchCollect is a free log subscription operation binding the contract event 0x40d0efd1a53d60ecbf40971b9daf7dc90178c3aadc7aab1765632738fa8b8f01.
//
// Solidity: event Collect(uint256 indexed tokenId, address recipient, uint256 amount0, uint256 amount1)
func (_NftPositionManager *NftPositionManagerFilterer) WatchCollect(opts *bind.WatchOpts, sink chan<- *NftPositionManagerCollect, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NftPositionManager.contract.WatchLogs(opts, "Collect", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftPositionManagerCollect)
				if err := _NftPositionManager.contract.UnpackLog(event, "Collect", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollect is a log parse operation binding the contract event 0x40d0efd1a53d60ecbf40971b9daf7dc90178c3aadc7aab1765632738fa8b8f01.
//
// Solidity: event Collect(uint256 indexed tokenId, address recipient, uint256 amount0, uint256 amount1)
func (_NftPositionManager *NftPositionManagerFilterer) ParseCollect(log types.Log) (*NftPositionManagerCollect, error) {
	event := new(NftPositionManagerCollect)
	if err := _NftPositionManager.contract.UnpackLog(event, "Collect", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftPositionManagerDecreaseLiquidityIterator is returned from FilterDecreaseLiquidity and is used to iterate over the raw logs and unpacked data for DecreaseLiquidity events raised by the NftPositionManager contract.
type NftPositionManagerDecreaseLiquidityIterator struct {
	Event *NftPositionManagerDecreaseLiquidity // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftPositionManagerDecreaseLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftPositionManagerDecreaseLiquidity)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftPositionManagerDecreaseLiquidity)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftPositionManagerDecreaseLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftPositionManagerDecreaseLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftPositionManagerDecreaseLiquidity represents a DecreaseLiquidity event raised by the NftPositionManager contract.
type NftPositionManagerDecreaseLiquidity struct {
	TokenId   *big.Int
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDecreaseLiquidity is a free log retrieval operation binding the contract event 0x26f6a048ee9138f2c0ce266f322cb99228e8d619ae2bff30c67f8dcf9d2377b4.
//
// Solidity: event DecreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NftPositionManager *NftPositionManagerFilterer) FilterDecreaseLiquidity(opts *bind.FilterOpts, tokenId []*big.Int) (*NftPositionManagerDecreaseLiquidityIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NftPositionManager.contract.FilterLogs(opts, "DecreaseLiquidity", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftPositionManagerDecreaseLiquidityIterator{contract: _NftPositionManager.contract, event: "DecreaseLiquidity", logs: logs, sub: sub}, nil
}

// WatchDecreaseLiquidity is a free log subscription operation binding the contract event 0x26f6a048ee9138f2c0ce266f322cb99228e8d619ae2bff30c67f8dcf9d2377b4.
//
// Solidity: event DecreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NftPositionManager *NftPositionManagerFilterer) WatchDecreaseLiquidity(opts *bind.WatchOpts, sink chan<- *NftPositionManagerDecreaseLiquidity, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NftPositionManager.contract.WatchLogs(opts, "DecreaseLiquidity", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftPositionManagerDecreaseLiquidity)
				if err := _NftPositionManager.contract.UnpackLog(event, "DecreaseLiquidity", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDecreaseLiquidity is a log parse operation binding the contract event 0x26f6a048ee9138f2c0ce266f322cb99228e8d619ae2bff30c67f8dcf9d2377b4.
//
// Solidity: event DecreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NftPositionManager *NftPositionManagerFilterer) ParseDecreaseLiquidity(log types.Log) (*NftPositionManagerDecreaseLiquidity, error) {
	event := new(NftPositionManagerDecreaseLiquidity)
	if err := _NftPositionManager.contract.UnpackLog(event, "DecreaseLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftPositionManagerIncreaseLiquidityIterator is returned from FilterIncreaseLiquidity and is used to iterate over the raw logs and unpacked data for IncreaseLiquidity events raised by the NftPositionManager contract.
type NftPositionManagerIncreaseLiquidityIterator struct {
	Event *NftPositionManagerIncreaseLiquidity // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftPositionManagerIncreaseLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftPositionManagerIncreaseLiquidity)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftPositionManagerIncreaseLiquidity)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftPositionManagerIncreaseLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftPositionManagerIncreaseLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftPositionManagerIncreaseLiquidity represents a IncreaseLiquidity event raised by the NftPositionManager contract.
type NftPositionManagerIncreaseLiquidity struct {
	TokenId   *big.Int
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncreaseLiquidity is a free log retrieval operation binding the contract event 0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f.
//
// Solidity: event IncreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NftPositionManager *NftPositionManagerFilterer) FilterIncreaseLiquidity(opts *bind.FilterOpts, tokenId []*big.Int) (*NftPositionManagerIncreaseLiquidityIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NftPositionManager.contract.FilterLogs(opts, "IncreaseLiquidity", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftPositionManagerIncreaseLiquidityIterator{contract: _NftPositionManager.contract, event: "IncreaseLiquidity", logs: logs, sub: sub}, nil
}

// WatchIncreaseLiquidity is a free log subscription operation binding the contract event 0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f.
//
// Solidity: event IncreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NftPositionManager *NftPositionManagerFilterer) WatchIncreaseLiquidity(opts *bind.WatchOpts, sink chan<- *NftPositionManagerIncreaseLiquidity, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NftPositionManager.contract.WatchLogs(opts, "IncreaseLiquidity", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftPositionManagerIncreaseLiquidity)
				if err := _NftPositionManager.contract.UnpackLog(event, "IncreaseLiquidity", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncreaseLiquidity is a log parse operation binding the contract event 0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f.
//
// Solidity: event IncreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NftPositionManager *NftPositionManagerFilterer) ParseIncreaseLiquidity(log types.Log) (*NftPositionManagerIncreaseLiquidity, error) {
	event := new(NftPositionManagerIncreaseLiquidity)
	if err := _NftPositionManager.contract.UnpackLog(event, "IncreaseLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
