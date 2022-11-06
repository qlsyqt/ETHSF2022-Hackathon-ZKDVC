// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package badge

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
)

// IssuanceInputs is an auto generated low-level Go binding around an user-defined struct.
type IssuanceInputs struct {
	DataCategory string
	SubCategory  string
	PublicInputs [10]*big.Int
}

// RevocationInputs is an auto generated low-level Go binding around an user-defined struct.
type RevocationInputs struct {
	DataCategory string
	SubCategory  string
	PublicInputs [5]*big.Int
}

// SnarkProof is an auto generated low-level Go binding around an user-defined struct.
type SnarkProof struct {
	A [2]*big.Int
	B [2][2]*big.Int
	C [2]*big.Int
}

// BadgeMetaData contains all meta data concerning the Badge contract.
var BadgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"dcpClaimVerifier_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dcpRevocationVerifier_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"DCPBadgeBurnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"DCPBadgeMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from_\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structSnarkProof\",\"name\":\"snarkProof_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"dataCategory\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subCategory\",\"type\":\"string\"},{\"internalType\":\"uint256[5]\",\"name\":\"publicInputs\",\"type\":\"uint256[5]\"}],\"internalType\":\"structRevocationInputs\",\"name\":\"inputs_\",\"type\":\"tuple\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"}],\"name\":\"getTokenCategory\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structSnarkProof\",\"name\":\"snarkProof_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"dataCategory\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subCategory\",\"type\":\"string\"},{\"internalType\":\"uint256[10]\",\"name\":\"publicInputs\",\"type\":\"uint256[10]\"}],\"internalType\":\"structIssuanceInputs\",\"name\":\"inputs_\",\"type\":\"tuple\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dataCategory\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subCategory\",\"type\":\"string\"}],\"name\":\"mintNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dcpClaimVerifier_\",\"type\":\"address\"}],\"name\":\"setClaimVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dcpRevocationVerifier_\",\"type\":\"address\"}],\"name\":\"setRevocationVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0xaaa",
}

// BadgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BadgeMetaData.ABI instead.
var BadgeABI = BadgeMetaData.ABI

// BadgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BadgeMetaData.Bin instead.
var BadgeBin = BadgeMetaData.Bin

// DeployBadge deploys a new Ethereum contract, binding an instance of Badge to it.
func DeployBadge(auth *bind.TransactOpts, backend bind.ContractBackend, uri_ string, dcpClaimVerifier_ common.Address, dcpRevocationVerifier_ common.Address) (common.Address, *types.Transaction, *Badge, error) {
	parsed, err := BadgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BadgeBin), backend, uri_, dcpClaimVerifier_, dcpRevocationVerifier_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Badge{BadgeCaller: BadgeCaller{contract: contract}, BadgeTransactor: BadgeTransactor{contract: contract}, BadgeFilterer: BadgeFilterer{contract: contract}}, nil
}

// Badge is an auto generated Go binding around an Ethereum contract.
type Badge struct {
	BadgeCaller     // Read-only binding to the contract
	BadgeTransactor // Write-only binding to the contract
	BadgeFilterer   // Log filterer for contract events
}

// BadgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BadgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BadgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BadgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BadgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BadgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BadgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BadgeSession struct {
	Contract     *Badge            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BadgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BadgeCallerSession struct {
	Contract *BadgeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BadgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BadgeTransactorSession struct {
	Contract     *BadgeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BadgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BadgeRaw struct {
	Contract *Badge // Generic contract binding to access the raw methods on
}

// BadgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BadgeCallerRaw struct {
	Contract *BadgeCaller // Generic read-only contract binding to access the raw methods on
}

// BadgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BadgeTransactorRaw struct {
	Contract *BadgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBadge creates a new instance of Badge, bound to a specific deployed contract.
func NewBadge(address common.Address, backend bind.ContractBackend) (*Badge, error) {
	contract, err := bindBadge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Badge{BadgeCaller: BadgeCaller{contract: contract}, BadgeTransactor: BadgeTransactor{contract: contract}, BadgeFilterer: BadgeFilterer{contract: contract}}, nil
}

// NewBadgeCaller creates a new read-only instance of Badge, bound to a specific deployed contract.
func NewBadgeCaller(address common.Address, caller bind.ContractCaller) (*BadgeCaller, error) {
	contract, err := bindBadge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BadgeCaller{contract: contract}, nil
}

// NewBadgeTransactor creates a new write-only instance of Badge, bound to a specific deployed contract.
func NewBadgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BadgeTransactor, error) {
	contract, err := bindBadge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BadgeTransactor{contract: contract}, nil
}

// NewBadgeFilterer creates a new log filterer instance of Badge, bound to a specific deployed contract.
func NewBadgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BadgeFilterer, error) {
	contract, err := bindBadge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BadgeFilterer{contract: contract}, nil
}

// bindBadge binds a generic wrapper to an already deployed contract.
func bindBadge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BadgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Badge *BadgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Badge.Contract.BadgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Badge *BadgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Badge.Contract.BadgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Badge *BadgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Badge.Contract.BadgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Badge *BadgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Badge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Badge *BadgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Badge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Badge *BadgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Badge.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Badge *BadgeCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Badge *BadgeSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Badge.Contract.BalanceOf(&_Badge.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Badge *BadgeCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Badge.Contract.BalanceOf(&_Badge.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Badge *BadgeCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Badge *BadgeSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Badge.Contract.GetApproved(&_Badge.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Badge *BadgeCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Badge.Contract.GetApproved(&_Badge.CallOpts, tokenId)
}

// GetTokenCategory is a free data retrieval call binding the contract method 0x2a601df3.
//
// Solidity: function getTokenCategory(uint256 token) view returns(string)
func (_Badge *BadgeCaller) GetTokenCategory(opts *bind.CallOpts, token *big.Int) (string, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "getTokenCategory", token)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTokenCategory is a free data retrieval call binding the contract method 0x2a601df3.
//
// Solidity: function getTokenCategory(uint256 token) view returns(string)
func (_Badge *BadgeSession) GetTokenCategory(token *big.Int) (string, error) {
	return _Badge.Contract.GetTokenCategory(&_Badge.CallOpts, token)
}

// GetTokenCategory is a free data retrieval call binding the contract method 0x2a601df3.
//
// Solidity: function getTokenCategory(uint256 token) view returns(string)
func (_Badge *BadgeCallerSession) GetTokenCategory(token *big.Int) (string, error) {
	return _Badge.Contract.GetTokenCategory(&_Badge.CallOpts, token)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Badge *BadgeCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Badge *BadgeSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Badge.Contract.IsApprovedForAll(&_Badge.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Badge *BadgeCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Badge.Contract.IsApprovedForAll(&_Badge.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Badge *BadgeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Badge *BadgeSession) Name() (string, error) {
	return _Badge.Contract.Name(&_Badge.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Badge *BadgeCallerSession) Name() (string, error) {
	return _Badge.Contract.Name(&_Badge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Badge *BadgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Badge *BadgeSession) Owner() (common.Address, error) {
	return _Badge.Contract.Owner(&_Badge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Badge *BadgeCallerSession) Owner() (common.Address, error) {
	return _Badge.Contract.Owner(&_Badge.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Badge *BadgeCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Badge *BadgeSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Badge.Contract.OwnerOf(&_Badge.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Badge *BadgeCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Badge.Contract.OwnerOf(&_Badge.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Badge *BadgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Badge *BadgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Badge.Contract.SupportsInterface(&_Badge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Badge *BadgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Badge.Contract.SupportsInterface(&_Badge.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Badge *BadgeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Badge *BadgeSession) Symbol() (string, error) {
	return _Badge.Contract.Symbol(&_Badge.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Badge *BadgeCallerSession) Symbol() (string, error) {
	return _Badge.Contract.Symbol(&_Badge.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Badge *BadgeCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Badge *BadgeSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Badge.Contract.TokenByIndex(&_Badge.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Badge *BadgeCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Badge.Contract.TokenByIndex(&_Badge.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Badge *BadgeCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Badge *BadgeSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Badge.Contract.TokenOfOwnerByIndex(&_Badge.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Badge *BadgeCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Badge.Contract.TokenOfOwnerByIndex(&_Badge.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Badge *BadgeCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Badge *BadgeSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Badge.Contract.TokenURI(&_Badge.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Badge *BadgeCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Badge.Contract.TokenURI(&_Badge.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Badge *BadgeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Badge *BadgeSession) TotalSupply() (*big.Int, error) {
	return _Badge.Contract.TotalSupply(&_Badge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Badge *BadgeCallerSession) TotalSupply() (*big.Int, error) {
	return _Badge.Contract.TotalSupply(&_Badge.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Badge *BadgeTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Badge *BadgeSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Badge.Contract.Approve(&_Badge.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Badge *BadgeTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Badge.Contract.Approve(&_Badge.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0xeec349a7.
//
// Solidity: function burn(address from_, (uint256[2],uint256[2][2],uint256[2]) snarkProof_, (string,string,uint256[5]) inputs_) returns()
func (_Badge *BadgeTransactor) Burn(opts *bind.TransactOpts, from_ common.Address, snarkProof_ SnarkProof, inputs_ RevocationInputs) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "burn", from_, snarkProof_, inputs_)
}

// Burn is a paid mutator transaction binding the contract method 0xeec349a7.
//
// Solidity: function burn(address from_, (uint256[2],uint256[2][2],uint256[2]) snarkProof_, (string,string,uint256[5]) inputs_) returns()
func (_Badge *BadgeSession) Burn(from_ common.Address, snarkProof_ SnarkProof, inputs_ RevocationInputs) (*types.Transaction, error) {
	return _Badge.Contract.Burn(&_Badge.TransactOpts, from_, snarkProof_, inputs_)
}

// Burn is a paid mutator transaction binding the contract method 0xeec349a7.
//
// Solidity: function burn(address from_, (uint256[2],uint256[2][2],uint256[2]) snarkProof_, (string,string,uint256[5]) inputs_) returns()
func (_Badge *BadgeTransactorSession) Burn(from_ common.Address, snarkProof_ SnarkProof, inputs_ RevocationInputs) (*types.Transaction, error) {
	return _Badge.Contract.Burn(&_Badge.TransactOpts, from_, snarkProof_, inputs_)
}

// Mint is a paid mutator transaction binding the contract method 0x33a83659.
//
// Solidity: function mint(address to_, (uint256[2],uint256[2][2],uint256[2]) snarkProof_, (string,string,uint256[10]) inputs_) returns()
func (_Badge *BadgeTransactor) Mint(opts *bind.TransactOpts, to_ common.Address, snarkProof_ SnarkProof, inputs_ IssuanceInputs) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "mint", to_, snarkProof_, inputs_)
}

// Mint is a paid mutator transaction binding the contract method 0x33a83659.
//
// Solidity: function mint(address to_, (uint256[2],uint256[2][2],uint256[2]) snarkProof_, (string,string,uint256[10]) inputs_) returns()
func (_Badge *BadgeSession) Mint(to_ common.Address, snarkProof_ SnarkProof, inputs_ IssuanceInputs) (*types.Transaction, error) {
	return _Badge.Contract.Mint(&_Badge.TransactOpts, to_, snarkProof_, inputs_)
}

// Mint is a paid mutator transaction binding the contract method 0x33a83659.
//
// Solidity: function mint(address to_, (uint256[2],uint256[2][2],uint256[2]) snarkProof_, (string,string,uint256[10]) inputs_) returns()
func (_Badge *BadgeTransactorSession) Mint(to_ common.Address, snarkProof_ SnarkProof, inputs_ IssuanceInputs) (*types.Transaction, error) {
	return _Badge.Contract.Mint(&_Badge.TransactOpts, to_, snarkProof_, inputs_)
}

// MintNative is a paid mutator transaction binding the contract method 0x8774015e.
//
// Solidity: function mintNative(string dataCategory, string subCategory) returns()
func (_Badge *BadgeTransactor) MintNative(opts *bind.TransactOpts, dataCategory string, subCategory string) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "mintNative", dataCategory, subCategory)
}

// MintNative is a paid mutator transaction binding the contract method 0x8774015e.
//
// Solidity: function mintNative(string dataCategory, string subCategory) returns()
func (_Badge *BadgeSession) MintNative(dataCategory string, subCategory string) (*types.Transaction, error) {
	return _Badge.Contract.MintNative(&_Badge.TransactOpts, dataCategory, subCategory)
}

// MintNative is a paid mutator transaction binding the contract method 0x8774015e.
//
// Solidity: function mintNative(string dataCategory, string subCategory) returns()
func (_Badge *BadgeTransactorSession) MintNative(dataCategory string, subCategory string) (*types.Transaction, error) {
	return _Badge.Contract.MintNative(&_Badge.TransactOpts, dataCategory, subCategory)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Badge *BadgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Badge *BadgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Badge.Contract.RenounceOwnership(&_Badge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Badge *BadgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Badge.Contract.RenounceOwnership(&_Badge.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Badge *BadgeTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Badge *BadgeSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Badge.Contract.SafeTransferFrom(&_Badge.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Badge *BadgeTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Badge.Contract.SafeTransferFrom(&_Badge.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Badge *BadgeTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Badge *BadgeSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Badge.Contract.SafeTransferFrom0(&_Badge.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Badge *BadgeTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Badge.Contract.SafeTransferFrom0(&_Badge.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Badge *BadgeTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Badge *BadgeSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Badge.Contract.SetApprovalForAll(&_Badge.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Badge *BadgeTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Badge.Contract.SetApprovalForAll(&_Badge.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string uri_) returns()
func (_Badge *BadgeTransactor) SetBaseURI(opts *bind.TransactOpts, uri_ string) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "setBaseURI", uri_)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string uri_) returns()
func (_Badge *BadgeSession) SetBaseURI(uri_ string) (*types.Transaction, error) {
	return _Badge.Contract.SetBaseURI(&_Badge.TransactOpts, uri_)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string uri_) returns()
func (_Badge *BadgeTransactorSession) SetBaseURI(uri_ string) (*types.Transaction, error) {
	return _Badge.Contract.SetBaseURI(&_Badge.TransactOpts, uri_)
}

// SetClaimVerifier is a paid mutator transaction binding the contract method 0x38eb8425.
//
// Solidity: function setClaimVerifier(address dcpClaimVerifier_) returns()
func (_Badge *BadgeTransactor) SetClaimVerifier(opts *bind.TransactOpts, dcpClaimVerifier_ common.Address) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "setClaimVerifier", dcpClaimVerifier_)
}

// SetClaimVerifier is a paid mutator transaction binding the contract method 0x38eb8425.
//
// Solidity: function setClaimVerifier(address dcpClaimVerifier_) returns()
func (_Badge *BadgeSession) SetClaimVerifier(dcpClaimVerifier_ common.Address) (*types.Transaction, error) {
	return _Badge.Contract.SetClaimVerifier(&_Badge.TransactOpts, dcpClaimVerifier_)
}

// SetClaimVerifier is a paid mutator transaction binding the contract method 0x38eb8425.
//
// Solidity: function setClaimVerifier(address dcpClaimVerifier_) returns()
func (_Badge *BadgeTransactorSession) SetClaimVerifier(dcpClaimVerifier_ common.Address) (*types.Transaction, error) {
	return _Badge.Contract.SetClaimVerifier(&_Badge.TransactOpts, dcpClaimVerifier_)
}

// SetRevocationVerifier is a paid mutator transaction binding the contract method 0xd039ba23.
//
// Solidity: function setRevocationVerifier(address dcpRevocationVerifier_) returns()
func (_Badge *BadgeTransactor) SetRevocationVerifier(opts *bind.TransactOpts, dcpRevocationVerifier_ common.Address) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "setRevocationVerifier", dcpRevocationVerifier_)
}

// SetRevocationVerifier is a paid mutator transaction binding the contract method 0xd039ba23.
//
// Solidity: function setRevocationVerifier(address dcpRevocationVerifier_) returns()
func (_Badge *BadgeSession) SetRevocationVerifier(dcpRevocationVerifier_ common.Address) (*types.Transaction, error) {
	return _Badge.Contract.SetRevocationVerifier(&_Badge.TransactOpts, dcpRevocationVerifier_)
}

// SetRevocationVerifier is a paid mutator transaction binding the contract method 0xd039ba23.
//
// Solidity: function setRevocationVerifier(address dcpRevocationVerifier_) returns()
func (_Badge *BadgeTransactorSession) SetRevocationVerifier(dcpRevocationVerifier_ common.Address) (*types.Transaction, error) {
	return _Badge.Contract.SetRevocationVerifier(&_Badge.TransactOpts, dcpRevocationVerifier_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Badge *BadgeTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Badge *BadgeSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Badge.Contract.TransferFrom(&_Badge.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Badge *BadgeTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Badge.Contract.TransferFrom(&_Badge.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Badge *BadgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Badge *BadgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Badge.Contract.TransferOwnership(&_Badge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Badge *BadgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Badge.Contract.TransferOwnership(&_Badge.TransactOpts, newOwner)
}

// BadgeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Badge contract.
type BadgeApprovalIterator struct {
	Event *BadgeApproval // Event containing the contract specifics and raw log

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
func (it *BadgeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeApproval)
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
		it.Event = new(BadgeApproval)
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
func (it *BadgeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeApproval represents a Approval event raised by the Badge contract.
type BadgeApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Badge *BadgeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BadgeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Badge.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BadgeApprovalIterator{contract: _Badge.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Badge *BadgeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BadgeApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Badge.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeApproval)
				if err := _Badge.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Badge *BadgeFilterer) ParseApproval(log types.Log) (*BadgeApproval, error) {
	event := new(BadgeApproval)
	if err := _Badge.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Badge contract.
type BadgeApprovalForAllIterator struct {
	Event *BadgeApprovalForAll // Event containing the contract specifics and raw log

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
func (it *BadgeApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeApprovalForAll)
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
		it.Event = new(BadgeApprovalForAll)
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
func (it *BadgeApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeApprovalForAll represents a ApprovalForAll event raised by the Badge contract.
type BadgeApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Badge *BadgeFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BadgeApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Badge.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BadgeApprovalForAllIterator{contract: _Badge.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Badge *BadgeFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BadgeApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Badge.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeApprovalForAll)
				if err := _Badge.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Badge *BadgeFilterer) ParseApprovalForAll(log types.Log) (*BadgeApprovalForAll, error) {
	event := new(BadgeApprovalForAll)
	if err := _Badge.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeDCPBadgeBurntIterator is returned from FilterDCPBadgeBurnt and is used to iterate over the raw logs and unpacked data for DCPBadgeBurnt events raised by the Badge contract.
type BadgeDCPBadgeBurntIterator struct {
	Event *BadgeDCPBadgeBurnt // Event containing the contract specifics and raw log

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
func (it *BadgeDCPBadgeBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeDCPBadgeBurnt)
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
		it.Event = new(BadgeDCPBadgeBurnt)
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
func (it *BadgeDCPBadgeBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeDCPBadgeBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeDCPBadgeBurnt represents a DCPBadgeBurnt event raised by the Badge contract.
type BadgeDCPBadgeBurnt struct {
	Id       *big.Int
	To       common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDCPBadgeBurnt is a free log retrieval operation binding the contract event 0xd16fb1323274da40cb54d1787af0f6856989a0db6a0d4ddc50ff571c6f2b9227.
//
// Solidity: event DCPBadgeBurnt(uint256 indexed id, address to, address operator)
func (_Badge *BadgeFilterer) FilterDCPBadgeBurnt(opts *bind.FilterOpts, id []*big.Int) (*BadgeDCPBadgeBurntIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Badge.contract.FilterLogs(opts, "DCPBadgeBurnt", idRule)
	if err != nil {
		return nil, err
	}
	return &BadgeDCPBadgeBurntIterator{contract: _Badge.contract, event: "DCPBadgeBurnt", logs: logs, sub: sub}, nil
}

// WatchDCPBadgeBurnt is a free log subscription operation binding the contract event 0xd16fb1323274da40cb54d1787af0f6856989a0db6a0d4ddc50ff571c6f2b9227.
//
// Solidity: event DCPBadgeBurnt(uint256 indexed id, address to, address operator)
func (_Badge *BadgeFilterer) WatchDCPBadgeBurnt(opts *bind.WatchOpts, sink chan<- *BadgeDCPBadgeBurnt, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Badge.contract.WatchLogs(opts, "DCPBadgeBurnt", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeDCPBadgeBurnt)
				if err := _Badge.contract.UnpackLog(event, "DCPBadgeBurnt", log); err != nil {
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

// ParseDCPBadgeBurnt is a log parse operation binding the contract event 0xd16fb1323274da40cb54d1787af0f6856989a0db6a0d4ddc50ff571c6f2b9227.
//
// Solidity: event DCPBadgeBurnt(uint256 indexed id, address to, address operator)
func (_Badge *BadgeFilterer) ParseDCPBadgeBurnt(log types.Log) (*BadgeDCPBadgeBurnt, error) {
	event := new(BadgeDCPBadgeBurnt)
	if err := _Badge.contract.UnpackLog(event, "DCPBadgeBurnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeDCPBadgeMintIterator is returned from FilterDCPBadgeMint and is used to iterate over the raw logs and unpacked data for DCPBadgeMint events raised by the Badge contract.
type BadgeDCPBadgeMintIterator struct {
	Event *BadgeDCPBadgeMint // Event containing the contract specifics and raw log

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
func (it *BadgeDCPBadgeMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeDCPBadgeMint)
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
		it.Event = new(BadgeDCPBadgeMint)
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
func (it *BadgeDCPBadgeMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeDCPBadgeMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeDCPBadgeMint represents a DCPBadgeMint event raised by the Badge contract.
type BadgeDCPBadgeMint struct {
	Id       *big.Int
	To       common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDCPBadgeMint is a free log retrieval operation binding the contract event 0xd0c0225300ec5b4bac0eecf6efa0d4003883f63052c87bb0dffac72bce70f2f4.
//
// Solidity: event DCPBadgeMint(uint256 indexed id, address to, address operator)
func (_Badge *BadgeFilterer) FilterDCPBadgeMint(opts *bind.FilterOpts, id []*big.Int) (*BadgeDCPBadgeMintIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Badge.contract.FilterLogs(opts, "DCPBadgeMint", idRule)
	if err != nil {
		return nil, err
	}
	return &BadgeDCPBadgeMintIterator{contract: _Badge.contract, event: "DCPBadgeMint", logs: logs, sub: sub}, nil
}

// WatchDCPBadgeMint is a free log subscription operation binding the contract event 0xd0c0225300ec5b4bac0eecf6efa0d4003883f63052c87bb0dffac72bce70f2f4.
//
// Solidity: event DCPBadgeMint(uint256 indexed id, address to, address operator)
func (_Badge *BadgeFilterer) WatchDCPBadgeMint(opts *bind.WatchOpts, sink chan<- *BadgeDCPBadgeMint, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Badge.contract.WatchLogs(opts, "DCPBadgeMint", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeDCPBadgeMint)
				if err := _Badge.contract.UnpackLog(event, "DCPBadgeMint", log); err != nil {
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

// ParseDCPBadgeMint is a log parse operation binding the contract event 0xd0c0225300ec5b4bac0eecf6efa0d4003883f63052c87bb0dffac72bce70f2f4.
//
// Solidity: event DCPBadgeMint(uint256 indexed id, address to, address operator)
func (_Badge *BadgeFilterer) ParseDCPBadgeMint(log types.Log) (*BadgeDCPBadgeMint, error) {
	event := new(BadgeDCPBadgeMint)
	if err := _Badge.contract.UnpackLog(event, "DCPBadgeMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Badge contract.
type BadgeOwnershipTransferredIterator struct {
	Event *BadgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BadgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeOwnershipTransferred)
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
		it.Event = new(BadgeOwnershipTransferred)
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
func (it *BadgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeOwnershipTransferred represents a OwnershipTransferred event raised by the Badge contract.
type BadgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Badge *BadgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BadgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Badge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BadgeOwnershipTransferredIterator{contract: _Badge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Badge *BadgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BadgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Badge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeOwnershipTransferred)
				if err := _Badge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Badge *BadgeFilterer) ParseOwnershipTransferred(log types.Log) (*BadgeOwnershipTransferred, error) {
	event := new(BadgeOwnershipTransferred)
	if err := _Badge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Badge contract.
type BadgeTransferIterator struct {
	Event *BadgeTransfer // Event containing the contract specifics and raw log

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
func (it *BadgeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeTransfer)
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
		it.Event = new(BadgeTransfer)
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
func (it *BadgeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeTransfer represents a Transfer event raised by the Badge contract.
type BadgeTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Badge *BadgeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BadgeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Badge.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BadgeTransferIterator{contract: _Badge.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Badge *BadgeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BadgeTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Badge.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeTransfer)
				if err := _Badge.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Badge *BadgeFilterer) ParseTransfer(log types.Log) (*BadgeTransfer, error) {
	event := new(BadgeTransfer)
	if err := _Badge.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
