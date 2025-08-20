// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package inferences

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// Ierc20burnableMetaData contains all meta data concerning the Ierc20burnable contract.
var Ierc20burnableMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "Ierc20burnable",
	Bin: "0x608060405234801561001057600080fd5b506040518060400160405280600981526020017f54455354546f6b656e00000000000000000000000000000000000000000000008152506040518060400160405280600281526020017f5454000000000000000000000000000000000000000000000000000000000000815250816003908161008c91906105bc565b50806004908161009c91906105bc565b5050506100ba336a52b7d2dcc80cd2e40000006100bf60201b60201c565b6107ae565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036101315760006040517fec442f0500000000000000000000000000000000000000000000000000000000815260040161012891906106cf565b60405180910390fd5b6101436000838361014760201b60201c565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361019957806002600082825461018d9190610719565b9250508190555061026c565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610225578381836040517fe450d38c00000000000000000000000000000000000000000000000000000000815260040161021c9392919061075c565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036102b55780600260008282540392505081905550610302565b806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161035f9190610793565b60405180910390a3505050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806103ed57607f821691505b602082108103610400576103ff6103a6565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026104687fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261042b565b610472868361042b565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006104b96104b46104af8461048a565b610494565b61048a565b9050919050565b6000819050919050565b6104d38361049e565b6104e76104df826104c0565b848454610438565b825550505050565b600090565b6104fc6104ef565b6105078184846104ca565b505050565b5b8181101561052b576105206000826104f4565b60018101905061050d565b5050565b601f8211156105705761054181610406565b61054a8461041b565b81016020851015610559578190505b61056d6105658561041b565b83018261050c565b50505b505050565b600082821c905092915050565b600061059360001984600802610575565b1980831691505092915050565b60006105ac8383610582565b9150826002028217905092915050565b6105c58261036c565b67ffffffffffffffff8111156105de576105dd610377565b5b6105e882546103d5565b6105f382828561052f565b600060209050601f8311600181146106265760008415610614578287015190505b61061e85826105a0565b865550610686565b601f19841661063486610406565b60005b8281101561065c57848901518255600182019150602085019450602081019050610637565b868310156106795784890151610675601f891682610582565b8355505b6001600288020188555050505b505050505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006106b98261068e565b9050919050565b6106c9816106ae565b82525050565b60006020820190506106e460008301846106c0565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006107248261048a565b915061072f8361048a565b9250828201905080821115610747576107466106ea565b5b92915050565b6107568161048a565b82525050565b600060608201905061077160008301866106c0565b61077e602083018561074d565b61078b604083018461074d565b949350505050565b60006020820190506107a8600083018461074d565b92915050565b610f86806107bd6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806342966c681161007157806342966c681461016857806370a082311461018457806379cc6790146101b457806395d89b41146101d0578063a9059cbb146101ee578063dd62ed3e1461021e576100a9565b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100fc57806323b872dd1461011a578063313ce5671461014a575b600080fd5b6100b661024e565b6040516100c39190610bad565b60405180910390f35b6100e660048036038101906100e19190610c68565b6102e0565b6040516100f39190610cc3565b60405180910390f35b610104610303565b6040516101119190610ced565b60405180910390f35b610134600480360381019061012f9190610d08565b61030d565b6040516101419190610cc3565b60405180910390f35b61015261033c565b60405161015f9190610d77565b60405180910390f35b610182600480360381019061017d9190610d92565b610345565b005b61019e60048036038101906101999190610dbf565b610359565b6040516101ab9190610ced565b60405180910390f35b6101ce60048036038101906101c99190610c68565b6103a1565b005b6101d86103c1565b6040516101e59190610bad565b60405180910390f35b61020860048036038101906102039190610c68565b610453565b6040516102159190610cc3565b60405180910390f35b61023860048036038101906102339190610dec565b610476565b6040516102459190610ced565b60405180910390f35b60606003805461025d90610e5b565b80601f016020809104026020016040519081016040528092919081815260200182805461028990610e5b565b80156102d65780601f106102ab576101008083540402835291602001916102d6565b820191906000526020600020905b8154815290600101906020018083116102b957829003601f168201915b5050505050905090565b6000806102eb6104fd565b90506102f8818585610505565b600191505092915050565b6000600254905090565b6000806103186104fd565b9050610325858285610517565b6103308585856105ab565b60019150509392505050565b60006012905090565b6103566103506104fd565b8261069f565b50565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6103b3826103ad6104fd565b83610517565b6103bd828261069f565b5050565b6060600480546103d090610e5b565b80601f01602080910402602001604051908101604052809291908181526020018280546103fc90610e5b565b80156104495780601f1061041e57610100808354040283529160200191610449565b820191906000526020600020905b81548152906001019060200180831161042c57829003601f168201915b5050505050905090565b60008061045e6104fd565b905061046b8185856105ab565b600191505092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b6105128383836001610721565b505050565b60006105238484610476565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146105a55781811015610595578281836040517ffb8f41b200000000000000000000000000000000000000000000000000000000815260040161058c93929190610e9b565b60405180910390fd5b6105a484848484036000610721565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361061d5760006040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016106149190610ed2565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361068f5760006040517fec442f050000000000000000000000000000000000000000000000000000000081526004016106869190610ed2565b60405180910390fd5b61069a8383836108f8565b505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036107115760006040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016107089190610ed2565b60405180910390fd5b61071d826000836108f8565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16036107935760006040517fe602df0500000000000000000000000000000000000000000000000000000000815260040161078a9190610ed2565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036108055760006040517f94280d620000000000000000000000000000000000000000000000000000000081526004016107fc9190610ed2565b60405180910390fd5b81600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555080156108f2578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516108e99190610ced565b60405180910390a35b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361094a57806002600082825461093e9190610f1c565b92505081905550610a1d565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050818110156109d6578381836040517fe450d38c0000000000000000000000000000000000000000000000000000000081526004016109cd93929190610e9b565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610a665780600260008282540392505081905550610ab3565b806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610b109190610ced565b60405180910390a3505050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610b57578082015181840152602081019050610b3c565b60008484015250505050565b6000601f19601f8301169050919050565b6000610b7f82610b1d565b610b898185610b28565b9350610b99818560208601610b39565b610ba281610b63565b840191505092915050565b60006020820190508181036000830152610bc78184610b74565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610bff82610bd4565b9050919050565b610c0f81610bf4565b8114610c1a57600080fd5b50565b600081359050610c2c81610c06565b92915050565b6000819050919050565b610c4581610c32565b8114610c5057600080fd5b50565b600081359050610c6281610c3c565b92915050565b60008060408385031215610c7f57610c7e610bcf565b5b6000610c8d85828601610c1d565b9250506020610c9e85828601610c53565b9150509250929050565b60008115159050919050565b610cbd81610ca8565b82525050565b6000602082019050610cd86000830184610cb4565b92915050565b610ce781610c32565b82525050565b6000602082019050610d026000830184610cde565b92915050565b600080600060608486031215610d2157610d20610bcf565b5b6000610d2f86828701610c1d565b9350506020610d4086828701610c1d565b9250506040610d5186828701610c53565b9150509250925092565b600060ff82169050919050565b610d7181610d5b565b82525050565b6000602082019050610d8c6000830184610d68565b92915050565b600060208284031215610da857610da7610bcf565b5b6000610db684828501610c53565b91505092915050565b600060208284031215610dd557610dd4610bcf565b5b6000610de384828501610c1d565b91505092915050565b60008060408385031215610e0357610e02610bcf565b5b6000610e1185828601610c1d565b9250506020610e2285828601610c1d565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610e7357607f821691505b602082108103610e8657610e85610e2c565b5b50919050565b610e9581610bf4565b82525050565b6000606082019050610eb06000830186610e8c565b610ebd6020830185610cde565b610eca6040830184610cde565b949350505050565b6000602082019050610ee76000830184610e8c565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610f2782610c32565b9150610f3283610c32565b9250828201905080821115610f4a57610f49610eed565b5b9291505056fea26469706673582212205a23c688a4f16f5643750c76c0aa75547aab035f987a8ddb9b0809126e3994c864736f6c634300081c0033",
}

// Ierc20burnable is an auto generated Go binding around an Ethereum contract.
type Ierc20burnable struct {
	abi abi.ABI
}

// NewIerc20burnable creates a new instance of Ierc20burnable.
func NewIerc20burnable() *Ierc20burnable {
	parsed, err := Ierc20burnableMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Ierc20burnable{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Ierc20burnable) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (ierc20burnable *Ierc20burnable) PackAllowance(owner common.Address, spender common.Address) []byte {
	enc, err := ierc20burnable.abi.Pack("allowance", owner, spender)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (ierc20burnable *Ierc20burnable) TryPackAllowance(owner common.Address, spender common.Address) ([]byte, error) {
	return ierc20burnable.abi.Pack("allowance", owner, spender)
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (ierc20burnable *Ierc20burnable) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := ierc20burnable.abi.Unpack("allowance", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (ierc20burnable *Ierc20burnable) PackApprove(spender common.Address, value *big.Int) []byte {
	enc, err := ierc20burnable.abi.Pack("approve", spender, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (ierc20burnable *Ierc20burnable) TryPackApprove(spender common.Address, value *big.Int) ([]byte, error) {
	return ierc20burnable.abi.Pack("approve", spender, value)
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (ierc20burnable *Ierc20burnable) UnpackApprove(data []byte) (bool, error) {
	out, err := ierc20burnable.abi.Unpack("approve", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (ierc20burnable *Ierc20burnable) PackBalanceOf(account common.Address) []byte {
	enc, err := ierc20burnable.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (ierc20burnable *Ierc20burnable) TryPackBalanceOf(account common.Address) ([]byte, error) {
	return ierc20burnable.abi.Pack("balanceOf", account)
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (ierc20burnable *Ierc20burnable) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := ierc20burnable.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42966c68.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function burn(uint256 value) returns()
func (ierc20burnable *Ierc20burnable) PackBurn(value *big.Int) []byte {
	enc, err := ierc20burnable.abi.Pack("burn", value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42966c68.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function burn(uint256 value) returns()
func (ierc20burnable *Ierc20burnable) TryPackBurn(value *big.Int) ([]byte, error) {
	return ierc20burnable.abi.Pack("burn", value)
}

// PackBurnFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79cc6790.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (ierc20burnable *Ierc20burnable) PackBurnFrom(account common.Address, value *big.Int) []byte {
	enc, err := ierc20burnable.abi.Pack("burnFrom", account, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBurnFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79cc6790.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (ierc20burnable *Ierc20burnable) TryPackBurnFrom(account common.Address, value *big.Int) ([]byte, error) {
	return ierc20burnable.abi.Pack("burnFrom", account, value)
}

// PackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function decimals() view returns(uint8)
func (ierc20burnable *Ierc20burnable) PackDecimals() []byte {
	enc, err := ierc20burnable.abi.Pack("decimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function decimals() view returns(uint8)
func (ierc20burnable *Ierc20burnable) TryPackDecimals() ([]byte, error) {
	return ierc20burnable.abi.Pack("decimals")
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (ierc20burnable *Ierc20burnable) UnpackDecimals(data []byte) (uint8, error) {
	out, err := ierc20burnable.abi.Unpack("decimals", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function name() view returns(string)
func (ierc20burnable *Ierc20burnable) PackName() []byte {
	enc, err := ierc20burnable.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function name() view returns(string)
func (ierc20burnable *Ierc20burnable) TryPackName() ([]byte, error) {
	return ierc20burnable.abi.Pack("name")
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (ierc20burnable *Ierc20burnable) UnpackName(data []byte) (string, error) {
	out, err := ierc20burnable.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function symbol() view returns(string)
func (ierc20burnable *Ierc20burnable) PackSymbol() []byte {
	enc, err := ierc20burnable.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function symbol() view returns(string)
func (ierc20burnable *Ierc20burnable) TryPackSymbol() ([]byte, error) {
	return ierc20burnable.abi.Pack("symbol")
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (ierc20burnable *Ierc20burnable) UnpackSymbol(data []byte) (string, error) {
	out, err := ierc20burnable.abi.Unpack("symbol", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18160ddd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function totalSupply() view returns(uint256)
func (ierc20burnable *Ierc20burnable) PackTotalSupply() []byte {
	enc, err := ierc20burnable.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18160ddd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function totalSupply() view returns(uint256)
func (ierc20burnable *Ierc20burnable) TryPackTotalSupply() ([]byte, error) {
	return ierc20burnable.abi.Pack("totalSupply")
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (ierc20burnable *Ierc20burnable) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := ierc20burnable.abi.Unpack("totalSupply", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9059cbb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (ierc20burnable *Ierc20burnable) PackTransfer(to common.Address, value *big.Int) []byte {
	enc, err := ierc20burnable.abi.Pack("transfer", to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9059cbb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (ierc20burnable *Ierc20burnable) TryPackTransfer(to common.Address, value *big.Int) ([]byte, error) {
	return ierc20burnable.abi.Pack("transfer", to, value)
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (ierc20burnable *Ierc20burnable) UnpackTransfer(data []byte) (bool, error) {
	out, err := ierc20burnable.abi.Unpack("transfer", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (ierc20burnable *Ierc20burnable) PackTransferFrom(from common.Address, to common.Address, value *big.Int) []byte {
	enc, err := ierc20burnable.abi.Pack("transferFrom", from, to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (ierc20burnable *Ierc20burnable) TryPackTransferFrom(from common.Address, to common.Address, value *big.Int) ([]byte, error) {
	return ierc20burnable.abi.Pack("transferFrom", from, to, value)
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (ierc20burnable *Ierc20burnable) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := ierc20burnable.abi.Unpack("transferFrom", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// Ierc20burnableApproval represents a Approval event raised by the Ierc20burnable contract.
type Ierc20burnableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const Ierc20burnableApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (Ierc20burnableApproval) ContractEventName() string {
	return Ierc20burnableApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (ierc20burnable *Ierc20burnable) UnpackApprovalEvent(log *types.Log) (*Ierc20burnableApproval, error) {
	event := "Approval"
	if log.Topics[0] != ierc20burnable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Ierc20burnableApproval)
	if len(log.Data) > 0 {
		if err := ierc20burnable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ierc20burnable.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// Ierc20burnableTransfer represents a Transfer event raised by the Ierc20burnable contract.
type Ierc20burnableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const Ierc20burnableTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (Ierc20burnableTransfer) ContractEventName() string {
	return Ierc20burnableTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (ierc20burnable *Ierc20burnable) UnpackTransferEvent(log *types.Log) (*Ierc20burnableTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != ierc20burnable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Ierc20burnableTransfer)
	if len(log.Data) > 0 {
		if err := ierc20burnable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ierc20burnable.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (ierc20burnable *Ierc20burnable) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], ierc20burnable.abi.Errors["ERC20InsufficientAllowance"].ID.Bytes()[:4]) {
		return ierc20burnable.UnpackERC20InsufficientAllowanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], ierc20burnable.abi.Errors["ERC20InsufficientBalance"].ID.Bytes()[:4]) {
		return ierc20burnable.UnpackERC20InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], ierc20burnable.abi.Errors["ERC20InvalidApprover"].ID.Bytes()[:4]) {
		return ierc20burnable.UnpackERC20InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], ierc20burnable.abi.Errors["ERC20InvalidReceiver"].ID.Bytes()[:4]) {
		return ierc20burnable.UnpackERC20InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], ierc20burnable.abi.Errors["ERC20InvalidSender"].ID.Bytes()[:4]) {
		return ierc20burnable.UnpackERC20InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], ierc20burnable.abi.Errors["ERC20InvalidSpender"].ID.Bytes()[:4]) {
		return ierc20burnable.UnpackERC20InvalidSpenderError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// Ierc20burnableERC20InsufficientAllowance represents a ERC20InsufficientAllowance error raised by the Ierc20burnable contract.
type Ierc20burnableERC20InsufficientAllowance struct {
	Spender   common.Address
	Allowance *big.Int
	Needed    *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func Ierc20burnableERC20InsufficientAllowanceErrorID() common.Hash {
	return common.HexToHash("0xfb8f41b23e99d2101d86da76cdfa87dd51c82ed07d3cb62cbc473e469dbc75c3")
}

// UnpackERC20InsufficientAllowanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func (ierc20burnable *Ierc20burnable) UnpackERC20InsufficientAllowanceError(raw []byte) (*Ierc20burnableERC20InsufficientAllowance, error) {
	out := new(Ierc20burnableERC20InsufficientAllowance)
	if err := ierc20burnable.abi.UnpackIntoInterface(out, "ERC20InsufficientAllowance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Ierc20burnableERC20InsufficientBalance represents a ERC20InsufficientBalance error raised by the Ierc20burnable contract.
type Ierc20burnableERC20InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func Ierc20burnableERC20InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xe450d38cd8d9f7d95077d567d60ed49c7254716e6ad08fc9872816c97e0ffec6")
}

// UnpackERC20InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func (ierc20burnable *Ierc20burnable) UnpackERC20InsufficientBalanceError(raw []byte) (*Ierc20burnableERC20InsufficientBalance, error) {
	out := new(Ierc20burnableERC20InsufficientBalance)
	if err := ierc20burnable.abi.UnpackIntoInterface(out, "ERC20InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Ierc20burnableERC20InvalidApprover represents a ERC20InvalidApprover error raised by the Ierc20burnable contract.
type Ierc20burnableERC20InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidApprover(address approver)
func Ierc20burnableERC20InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xe602df05cc75712490294c6c104ab7c17f4030363910a7a2626411c6d3118847")
}

// UnpackERC20InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidApprover(address approver)
func (ierc20burnable *Ierc20burnable) UnpackERC20InvalidApproverError(raw []byte) (*Ierc20burnableERC20InvalidApprover, error) {
	out := new(Ierc20burnableERC20InvalidApprover)
	if err := ierc20burnable.abi.UnpackIntoInterface(out, "ERC20InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Ierc20burnableERC20InvalidReceiver represents a ERC20InvalidReceiver error raised by the Ierc20burnable contract.
type Ierc20burnableERC20InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func Ierc20burnableERC20InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0xec442f055133b72f3b2f9f0bb351c406b178527de2040a7d1feb4e058771f613")
}

// UnpackERC20InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func (ierc20burnable *Ierc20burnable) UnpackERC20InvalidReceiverError(raw []byte) (*Ierc20burnableERC20InvalidReceiver, error) {
	out := new(Ierc20burnableERC20InvalidReceiver)
	if err := ierc20burnable.abi.UnpackIntoInterface(out, "ERC20InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Ierc20burnableERC20InvalidSender represents a ERC20InvalidSender error raised by the Ierc20burnable contract.
type Ierc20burnableERC20InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSender(address sender)
func Ierc20burnableERC20InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x96c6fd1edd0cd6ef7ff0ecc0facdf53148dc0048b57fe58af65755250a7a96bd")
}

// UnpackERC20InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSender(address sender)
func (ierc20burnable *Ierc20burnable) UnpackERC20InvalidSenderError(raw []byte) (*Ierc20burnableERC20InvalidSender, error) {
	out := new(Ierc20burnableERC20InvalidSender)
	if err := ierc20burnable.abi.UnpackIntoInterface(out, "ERC20InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Ierc20burnableERC20InvalidSpender represents a ERC20InvalidSpender error raised by the Ierc20burnable contract.
type Ierc20burnableERC20InvalidSpender struct {
	Spender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSpender(address spender)
func Ierc20burnableERC20InvalidSpenderErrorID() common.Hash {
	return common.HexToHash("0x94280d62c347d8d9f4d59a76ea321452406db88df38e0c9da304f58b57b373a2")
}

// UnpackERC20InvalidSpenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSpender(address spender)
func (ierc20burnable *Ierc20burnable) UnpackERC20InvalidSpenderError(raw []byte) (*Ierc20burnableERC20InvalidSpender, error) {
	out := new(Ierc20burnableERC20InvalidSpender)
	if err := ierc20burnable.abi.UnpackIntoInterface(out, "ERC20InvalidSpender", raw); err != nil {
		return nil, err
	}
	return out, nil
}
