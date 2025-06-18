// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// GatewayMetaData contains all meta data concerning the Gateway contract.
var GatewayMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addressGeneralID\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decryptionKeys\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"encryptionKeyExists\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fids\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generalDecryptionKeys\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generalIDRequested\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generalKeyRequested\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRandomnessByHeight\",\"inputs\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestEncryptionKey\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestRandomness\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestRandomnessHashOnly\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestGeneralDecryptionKey\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestGeneralID\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRequestGeneralFee\",\"inputs\":[{\"name\":\"newFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitDecryptionKey\",\"inputs\":[{\"name\":\"encryptionKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"decryptionKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitEncryptionKey\",\"inputs\":[{\"name\":\"encryptionKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitFID\",\"inputs\":[{\"name\":\"_requester\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_fid\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitGeneralDecryptionKeys\",\"inputs\":[{\"name\":\"requester\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decryptionKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FIDSubmitted\",\"inputs\":[{\"name\":\"_requester\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fid\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestGeneralID\",\"inputs\":[{\"name\":\"requester\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestGeneralKey\",\"inputs\":[{\"name\":\"requester\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayMetaData.ABI instead.
var GatewayABI = GatewayMetaData.ABI

// Gateway is an auto generated Go binding around an Ethereum contract.
type Gateway struct {
	GatewayCaller     // Read-only binding to the contract
	GatewayTransactor // Write-only binding to the contract
	GatewayFilterer   // Log filterer for contract events
}

// GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewaySession struct {
	Contract     *Gateway          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayCallerSession struct {
	Contract *GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayTransactorSession struct {
	Contract     *GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayRaw struct {
	Contract *Gateway // Generic contract binding to access the raw methods on
}

// GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayCallerRaw struct {
	Contract *GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayTransactorRaw struct {
	Contract *GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGateway creates a new instance of Gateway, bound to a specific deployed contract.
func NewGateway(address common.Address, backend bind.ContractBackend) (*Gateway, error) {
	contract, err := bindGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gateway{GatewayCaller: GatewayCaller{contract: contract}, GatewayTransactor: GatewayTransactor{contract: contract}, GatewayFilterer: GatewayFilterer{contract: contract}}, nil
}

// NewGatewayCaller creates a new read-only instance of Gateway, bound to a specific deployed contract.
func NewGatewayCaller(address common.Address, caller bind.ContractCaller) (*GatewayCaller, error) {
	contract, err := bindGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayCaller{contract: contract}, nil
}

// NewGatewayTransactor creates a new write-only instance of Gateway, bound to a specific deployed contract.
func NewGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayTransactor, error) {
	contract, err := bindGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayTransactor{contract: contract}, nil
}

// NewGatewayFilterer creates a new log filterer instance of Gateway, bound to a specific deployed contract.
func NewGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayFilterer, error) {
	contract, err := bindGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayFilterer{contract: contract}, nil
}

// bindGateway binds a generic wrapper to an already deployed contract.
func bindGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gateway *GatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gateway.Contract.GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gateway *GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.Contract.GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gateway *GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gateway.Contract.GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gateway *GatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gateway *GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gateway *GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gateway.Contract.contract.Transact(opts, method, params...)
}

// AddressGeneralID is a free data retrieval call binding the contract method 0x64d0f5d0.
//
// Solidity: function addressGeneralID(address ) view returns(uint256)
func (_Gateway *GatewayCaller) AddressGeneralID(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "addressGeneralID", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressGeneralID is a free data retrieval call binding the contract method 0x64d0f5d0.
//
// Solidity: function addressGeneralID(address ) view returns(uint256)
func (_Gateway *GatewaySession) AddressGeneralID(arg0 common.Address) (*big.Int, error) {
	return _Gateway.Contract.AddressGeneralID(&_Gateway.CallOpts, arg0)
}

// AddressGeneralID is a free data retrieval call binding the contract method 0x64d0f5d0.
//
// Solidity: function addressGeneralID(address ) view returns(uint256)
func (_Gateway *GatewayCallerSession) AddressGeneralID(arg0 common.Address) (*big.Int, error) {
	return _Gateway.Contract.AddressGeneralID(&_Gateway.CallOpts, arg0)
}

// DecryptionKeys is a free data retrieval call binding the contract method 0x111e512e.
//
// Solidity: function decryptionKeys(uint256 ) view returns(bytes)
func (_Gateway *GatewayCaller) DecryptionKeys(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "decryptionKeys", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DecryptionKeys is a free data retrieval call binding the contract method 0x111e512e.
//
// Solidity: function decryptionKeys(uint256 ) view returns(bytes)
func (_Gateway *GatewaySession) DecryptionKeys(arg0 *big.Int) ([]byte, error) {
	return _Gateway.Contract.DecryptionKeys(&_Gateway.CallOpts, arg0)
}

// DecryptionKeys is a free data retrieval call binding the contract method 0x111e512e.
//
// Solidity: function decryptionKeys(uint256 ) view returns(bytes)
func (_Gateway *GatewayCallerSession) DecryptionKeys(arg0 *big.Int) ([]byte, error) {
	return _Gateway.Contract.DecryptionKeys(&_Gateway.CallOpts, arg0)
}

// EncryptionKeyExists is a free data retrieval call binding the contract method 0x9c533b8a.
//
// Solidity: function encryptionKeyExists(bytes ) view returns(bool)
func (_Gateway *GatewayCaller) EncryptionKeyExists(opts *bind.CallOpts, arg0 []byte) (bool, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "encryptionKeyExists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EncryptionKeyExists is a free data retrieval call binding the contract method 0x9c533b8a.
//
// Solidity: function encryptionKeyExists(bytes ) view returns(bool)
func (_Gateway *GatewaySession) EncryptionKeyExists(arg0 []byte) (bool, error) {
	return _Gateway.Contract.EncryptionKeyExists(&_Gateway.CallOpts, arg0)
}

// EncryptionKeyExists is a free data retrieval call binding the contract method 0x9c533b8a.
//
// Solidity: function encryptionKeyExists(bytes ) view returns(bool)
func (_Gateway *GatewayCallerSession) EncryptionKeyExists(arg0 []byte) (bool, error) {
	return _Gateway.Contract.EncryptionKeyExists(&_Gateway.CallOpts, arg0)
}

// Fids is a free data retrieval call binding the contract method 0x6d4c0212.
//
// Solidity: function fids(address , uint256 ) view returns(string)
func (_Gateway *GatewayCaller) Fids(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "fids", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Fids is a free data retrieval call binding the contract method 0x6d4c0212.
//
// Solidity: function fids(address , uint256 ) view returns(string)
func (_Gateway *GatewaySession) Fids(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _Gateway.Contract.Fids(&_Gateway.CallOpts, arg0, arg1)
}

// Fids is a free data retrieval call binding the contract method 0x6d4c0212.
//
// Solidity: function fids(address , uint256 ) view returns(string)
func (_Gateway *GatewayCallerSession) Fids(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _Gateway.Contract.Fids(&_Gateway.CallOpts, arg0, arg1)
}

// GeneralDecryptionKeys is a free data retrieval call binding the contract method 0x80c110e8.
//
// Solidity: function generalDecryptionKeys(address , uint256 ) view returns(bytes)
func (_Gateway *GatewayCaller) GeneralDecryptionKeys(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "generalDecryptionKeys", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GeneralDecryptionKeys is a free data retrieval call binding the contract method 0x80c110e8.
//
// Solidity: function generalDecryptionKeys(address , uint256 ) view returns(bytes)
func (_Gateway *GatewaySession) GeneralDecryptionKeys(arg0 common.Address, arg1 *big.Int) ([]byte, error) {
	return _Gateway.Contract.GeneralDecryptionKeys(&_Gateway.CallOpts, arg0, arg1)
}

// GeneralDecryptionKeys is a free data retrieval call binding the contract method 0x80c110e8.
//
// Solidity: function generalDecryptionKeys(address , uint256 ) view returns(bytes)
func (_Gateway *GatewayCallerSession) GeneralDecryptionKeys(arg0 common.Address, arg1 *big.Int) ([]byte, error) {
	return _Gateway.Contract.GeneralDecryptionKeys(&_Gateway.CallOpts, arg0, arg1)
}

// GeneralIDRequested is a free data retrieval call binding the contract method 0xfb42e356.
//
// Solidity: function generalIDRequested(address , uint256 ) view returns(bool)
func (_Gateway *GatewayCaller) GeneralIDRequested(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "generalIDRequested", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GeneralIDRequested is a free data retrieval call binding the contract method 0xfb42e356.
//
// Solidity: function generalIDRequested(address , uint256 ) view returns(bool)
func (_Gateway *GatewaySession) GeneralIDRequested(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Gateway.Contract.GeneralIDRequested(&_Gateway.CallOpts, arg0, arg1)
}

// GeneralIDRequested is a free data retrieval call binding the contract method 0xfb42e356.
//
// Solidity: function generalIDRequested(address , uint256 ) view returns(bool)
func (_Gateway *GatewayCallerSession) GeneralIDRequested(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Gateway.Contract.GeneralIDRequested(&_Gateway.CallOpts, arg0, arg1)
}

// GeneralKeyRequested is a free data retrieval call binding the contract method 0x3095ffdf.
//
// Solidity: function generalKeyRequested(address , uint256 ) view returns(bool)
func (_Gateway *GatewayCaller) GeneralKeyRequested(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "generalKeyRequested", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GeneralKeyRequested is a free data retrieval call binding the contract method 0x3095ffdf.
//
// Solidity: function generalKeyRequested(address , uint256 ) view returns(bool)
func (_Gateway *GatewaySession) GeneralKeyRequested(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Gateway.Contract.GeneralKeyRequested(&_Gateway.CallOpts, arg0, arg1)
}

// GeneralKeyRequested is a free data retrieval call binding the contract method 0x3095ffdf.
//
// Solidity: function generalKeyRequested(address , uint256 ) view returns(bool)
func (_Gateway *GatewayCallerSession) GeneralKeyRequested(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Gateway.Contract.GeneralKeyRequested(&_Gateway.CallOpts, arg0, arg1)
}

// GetRandomnessByHeight is a free data retrieval call binding the contract method 0xee4d43cf.
//
// Solidity: function getRandomnessByHeight(uint256 height) view returns(bytes32)
func (_Gateway *GatewayCaller) GetRandomnessByHeight(opts *bind.CallOpts, height *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getRandomnessByHeight", height)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRandomnessByHeight is a free data retrieval call binding the contract method 0xee4d43cf.
//
// Solidity: function getRandomnessByHeight(uint256 height) view returns(bytes32)
func (_Gateway *GatewaySession) GetRandomnessByHeight(height *big.Int) ([32]byte, error) {
	return _Gateway.Contract.GetRandomnessByHeight(&_Gateway.CallOpts, height)
}

// GetRandomnessByHeight is a free data retrieval call binding the contract method 0xee4d43cf.
//
// Solidity: function getRandomnessByHeight(uint256 height) view returns(bytes32)
func (_Gateway *GatewayCallerSession) GetRandomnessByHeight(height *big.Int) ([32]byte, error) {
	return _Gateway.Contract.GetRandomnessByHeight(&_Gateway.CallOpts, height)
}

// LatestEncryptionKey is a free data retrieval call binding the contract method 0x74236c86.
//
// Solidity: function latestEncryptionKey() view returns(bytes)
func (_Gateway *GatewayCaller) LatestEncryptionKey(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "latestEncryptionKey")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// LatestEncryptionKey is a free data retrieval call binding the contract method 0x74236c86.
//
// Solidity: function latestEncryptionKey() view returns(bytes)
func (_Gateway *GatewaySession) LatestEncryptionKey() ([]byte, error) {
	return _Gateway.Contract.LatestEncryptionKey(&_Gateway.CallOpts)
}

// LatestEncryptionKey is a free data retrieval call binding the contract method 0x74236c86.
//
// Solidity: function latestEncryptionKey() view returns(bytes)
func (_Gateway *GatewayCallerSession) LatestEncryptionKey() ([]byte, error) {
	return _Gateway.Contract.LatestEncryptionKey(&_Gateway.CallOpts)
}

// LatestRandomness is a free data retrieval call binding the contract method 0x3b5cd3a8.
//
// Solidity: function latestRandomness() view returns(bytes32, uint256)
func (_Gateway *GatewayCaller) LatestRandomness(opts *bind.CallOpts) ([32]byte, *big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "latestRandomness")

	if err != nil {
		return *new([32]byte), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LatestRandomness is a free data retrieval call binding the contract method 0x3b5cd3a8.
//
// Solidity: function latestRandomness() view returns(bytes32, uint256)
func (_Gateway *GatewaySession) LatestRandomness() ([32]byte, *big.Int, error) {
	return _Gateway.Contract.LatestRandomness(&_Gateway.CallOpts)
}

// LatestRandomness is a free data retrieval call binding the contract method 0x3b5cd3a8.
//
// Solidity: function latestRandomness() view returns(bytes32, uint256)
func (_Gateway *GatewayCallerSession) LatestRandomness() ([32]byte, *big.Int, error) {
	return _Gateway.Contract.LatestRandomness(&_Gateway.CallOpts)
}

// LatestRandomnessHashOnly is a free data retrieval call binding the contract method 0xd0d343a7.
//
// Solidity: function latestRandomnessHashOnly() view returns(bytes32)
func (_Gateway *GatewayCaller) LatestRandomnessHashOnly(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "latestRandomnessHashOnly")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LatestRandomnessHashOnly is a free data retrieval call binding the contract method 0xd0d343a7.
//
// Solidity: function latestRandomnessHashOnly() view returns(bytes32)
func (_Gateway *GatewaySession) LatestRandomnessHashOnly() ([32]byte, error) {
	return _Gateway.Contract.LatestRandomnessHashOnly(&_Gateway.CallOpts)
}

// LatestRandomnessHashOnly is a free data retrieval call binding the contract method 0xd0d343a7.
//
// Solidity: function latestRandomnessHashOnly() view returns(bytes32)
func (_Gateway *GatewayCallerSession) LatestRandomnessHashOnly() ([32]byte, error) {
	return _Gateway.Contract.LatestRandomnessHashOnly(&_Gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gateway *GatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gateway *GatewaySession) Owner() (common.Address, error) {
	return _Gateway.Contract.Owner(&_Gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gateway *GatewayCallerSession) Owner() (common.Address, error) {
	return _Gateway.Contract.Owner(&_Gateway.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gateway *GatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gateway *GatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _Gateway.Contract.RenounceOwnership(&_Gateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gateway *GatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Gateway.Contract.RenounceOwnership(&_Gateway.TransactOpts)
}

// RequestGeneralDecryptionKey is a paid mutator transaction binding the contract method 0xedf73999.
//
// Solidity: function requestGeneralDecryptionKey(uint256 id) returns()
func (_Gateway *GatewayTransactor) RequestGeneralDecryptionKey(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "requestGeneralDecryptionKey", id)
}

// RequestGeneralDecryptionKey is a paid mutator transaction binding the contract method 0xedf73999.
//
// Solidity: function requestGeneralDecryptionKey(uint256 id) returns()
func (_Gateway *GatewaySession) RequestGeneralDecryptionKey(id *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.RequestGeneralDecryptionKey(&_Gateway.TransactOpts, id)
}

// RequestGeneralDecryptionKey is a paid mutator transaction binding the contract method 0xedf73999.
//
// Solidity: function requestGeneralDecryptionKey(uint256 id) returns()
func (_Gateway *GatewayTransactorSession) RequestGeneralDecryptionKey(id *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.RequestGeneralDecryptionKey(&_Gateway.TransactOpts, id)
}

// RequestGeneralID is a paid mutator transaction binding the contract method 0x6702e0f5.
//
// Solidity: function requestGeneralID() returns()
func (_Gateway *GatewayTransactor) RequestGeneralID(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "requestGeneralID")
}

// RequestGeneralID is a paid mutator transaction binding the contract method 0x6702e0f5.
//
// Solidity: function requestGeneralID() returns()
func (_Gateway *GatewaySession) RequestGeneralID() (*types.Transaction, error) {
	return _Gateway.Contract.RequestGeneralID(&_Gateway.TransactOpts)
}

// RequestGeneralID is a paid mutator transaction binding the contract method 0x6702e0f5.
//
// Solidity: function requestGeneralID() returns()
func (_Gateway *GatewayTransactorSession) RequestGeneralID() (*types.Transaction, error) {
	return _Gateway.Contract.RequestGeneralID(&_Gateway.TransactOpts)
}

// SetRequestGeneralFee is a paid mutator transaction binding the contract method 0xb09f544a.
//
// Solidity: function setRequestGeneralFee(uint256 newFee) returns()
func (_Gateway *GatewayTransactor) SetRequestGeneralFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "setRequestGeneralFee", newFee)
}

// SetRequestGeneralFee is a paid mutator transaction binding the contract method 0xb09f544a.
//
// Solidity: function setRequestGeneralFee(uint256 newFee) returns()
func (_Gateway *GatewaySession) SetRequestGeneralFee(newFee *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.SetRequestGeneralFee(&_Gateway.TransactOpts, newFee)
}

// SetRequestGeneralFee is a paid mutator transaction binding the contract method 0xb09f544a.
//
// Solidity: function setRequestGeneralFee(uint256 newFee) returns()
func (_Gateway *GatewayTransactorSession) SetRequestGeneralFee(newFee *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.SetRequestGeneralFee(&_Gateway.TransactOpts, newFee)
}

// SubmitDecryptionKey is a paid mutator transaction binding the contract method 0xaa878cf2.
//
// Solidity: function submitDecryptionKey(bytes encryptionKey, bytes decryptionKey, uint256 height) returns()
func (_Gateway *GatewayTransactor) SubmitDecryptionKey(opts *bind.TransactOpts, encryptionKey []byte, decryptionKey []byte, height *big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "submitDecryptionKey", encryptionKey, decryptionKey, height)
}

// SubmitDecryptionKey is a paid mutator transaction binding the contract method 0xaa878cf2.
//
// Solidity: function submitDecryptionKey(bytes encryptionKey, bytes decryptionKey, uint256 height) returns()
func (_Gateway *GatewaySession) SubmitDecryptionKey(encryptionKey []byte, decryptionKey []byte, height *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.SubmitDecryptionKey(&_Gateway.TransactOpts, encryptionKey, decryptionKey, height)
}

// SubmitDecryptionKey is a paid mutator transaction binding the contract method 0xaa878cf2.
//
// Solidity: function submitDecryptionKey(bytes encryptionKey, bytes decryptionKey, uint256 height) returns()
func (_Gateway *GatewayTransactorSession) SubmitDecryptionKey(encryptionKey []byte, decryptionKey []byte, height *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.SubmitDecryptionKey(&_Gateway.TransactOpts, encryptionKey, decryptionKey, height)
}

// SubmitEncryptionKey is a paid mutator transaction binding the contract method 0x0d006aa2.
//
// Solidity: function submitEncryptionKey(bytes encryptionKey) returns()
func (_Gateway *GatewayTransactor) SubmitEncryptionKey(opts *bind.TransactOpts, encryptionKey []byte) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "submitEncryptionKey", encryptionKey)
}

// SubmitEncryptionKey is a paid mutator transaction binding the contract method 0x0d006aa2.
//
// Solidity: function submitEncryptionKey(bytes encryptionKey) returns()
func (_Gateway *GatewaySession) SubmitEncryptionKey(encryptionKey []byte) (*types.Transaction, error) {
	return _Gateway.Contract.SubmitEncryptionKey(&_Gateway.TransactOpts, encryptionKey)
}

// SubmitEncryptionKey is a paid mutator transaction binding the contract method 0x0d006aa2.
//
// Solidity: function submitEncryptionKey(bytes encryptionKey) returns()
func (_Gateway *GatewayTransactorSession) SubmitEncryptionKey(encryptionKey []byte) (*types.Transaction, error) {
	return _Gateway.Contract.SubmitEncryptionKey(&_Gateway.TransactOpts, encryptionKey)
}

// SubmitFID is a paid mutator transaction binding the contract method 0x9cc8c1e8.
//
// Solidity: function submitFID(address _requester, string _fid, uint256 _id) returns()
func (_Gateway *GatewayTransactor) SubmitFID(opts *bind.TransactOpts, _requester common.Address, _fid string, _id *big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "submitFID", _requester, _fid, _id)
}

// SubmitFID is a paid mutator transaction binding the contract method 0x9cc8c1e8.
//
// Solidity: function submitFID(address _requester, string _fid, uint256 _id) returns()
func (_Gateway *GatewaySession) SubmitFID(_requester common.Address, _fid string, _id *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.SubmitFID(&_Gateway.TransactOpts, _requester, _fid, _id)
}

// SubmitFID is a paid mutator transaction binding the contract method 0x9cc8c1e8.
//
// Solidity: function submitFID(address _requester, string _fid, uint256 _id) returns()
func (_Gateway *GatewayTransactorSession) SubmitFID(_requester common.Address, _fid string, _id *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.SubmitFID(&_Gateway.TransactOpts, _requester, _fid, _id)
}

// SubmitGeneralDecryptionKeys is a paid mutator transaction binding the contract method 0xda538e41.
//
// Solidity: function submitGeneralDecryptionKeys(address requester, uint256 id, bytes decryptionKey) returns()
func (_Gateway *GatewayTransactor) SubmitGeneralDecryptionKeys(opts *bind.TransactOpts, requester common.Address, id *big.Int, decryptionKey []byte) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "submitGeneralDecryptionKeys", requester, id, decryptionKey)
}

// SubmitGeneralDecryptionKeys is a paid mutator transaction binding the contract method 0xda538e41.
//
// Solidity: function submitGeneralDecryptionKeys(address requester, uint256 id, bytes decryptionKey) returns()
func (_Gateway *GatewaySession) SubmitGeneralDecryptionKeys(requester common.Address, id *big.Int, decryptionKey []byte) (*types.Transaction, error) {
	return _Gateway.Contract.SubmitGeneralDecryptionKeys(&_Gateway.TransactOpts, requester, id, decryptionKey)
}

// SubmitGeneralDecryptionKeys is a paid mutator transaction binding the contract method 0xda538e41.
//
// Solidity: function submitGeneralDecryptionKeys(address requester, uint256 id, bytes decryptionKey) returns()
func (_Gateway *GatewayTransactorSession) SubmitGeneralDecryptionKeys(requester common.Address, id *big.Int, decryptionKey []byte) (*types.Transaction, error) {
	return _Gateway.Contract.SubmitGeneralDecryptionKeys(&_Gateway.TransactOpts, requester, id, decryptionKey)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gateway *GatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gateway *GatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.TransferOwnership(&_Gateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gateway *GatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.TransferOwnership(&_Gateway.TransactOpts, newOwner)
}

// GatewayFIDSubmittedIterator is returned from FilterFIDSubmitted and is used to iterate over the raw logs and unpacked data for FIDSubmitted events raised by the Gateway contract.
type GatewayFIDSubmittedIterator struct {
	Event *GatewayFIDSubmitted // Event containing the contract specifics and raw log

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
func (it *GatewayFIDSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayFIDSubmitted)
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
		it.Event = new(GatewayFIDSubmitted)
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
func (it *GatewayFIDSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayFIDSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayFIDSubmitted represents a FIDSubmitted event raised by the Gateway contract.
type GatewayFIDSubmitted struct {
	Requester common.Address
	Fid       common.Hash
	Id        *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFIDSubmitted is a free log retrieval operation binding the contract event 0x1cb0a3e71b09453c250a1f0ee46e90780c08d698ebed04ac4f0509b168fce7df.
//
// Solidity: event FIDSubmitted(address indexed _requester, string indexed fid, uint256 indexed id)
func (_Gateway *GatewayFilterer) FilterFIDSubmitted(opts *bind.FilterOpts, _requester []common.Address, fid []string, id []*big.Int) (*GatewayFIDSubmittedIterator, error) {

	var _requesterRule []interface{}
	for _, _requesterItem := range _requester {
		_requesterRule = append(_requesterRule, _requesterItem)
	}
	var fidRule []interface{}
	for _, fidItem := range fid {
		fidRule = append(fidRule, fidItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "FIDSubmitted", _requesterRule, fidRule, idRule)
	if err != nil {
		return nil, err
	}
	return &GatewayFIDSubmittedIterator{contract: _Gateway.contract, event: "FIDSubmitted", logs: logs, sub: sub}, nil
}

// WatchFIDSubmitted is a free log subscription operation binding the contract event 0x1cb0a3e71b09453c250a1f0ee46e90780c08d698ebed04ac4f0509b168fce7df.
//
// Solidity: event FIDSubmitted(address indexed _requester, string indexed fid, uint256 indexed id)
func (_Gateway *GatewayFilterer) WatchFIDSubmitted(opts *bind.WatchOpts, sink chan<- *GatewayFIDSubmitted, _requester []common.Address, fid []string, id []*big.Int) (event.Subscription, error) {

	var _requesterRule []interface{}
	for _, _requesterItem := range _requester {
		_requesterRule = append(_requesterRule, _requesterItem)
	}
	var fidRule []interface{}
	for _, fidItem := range fid {
		fidRule = append(fidRule, fidItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "FIDSubmitted", _requesterRule, fidRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayFIDSubmitted)
				if err := _Gateway.contract.UnpackLog(event, "FIDSubmitted", log); err != nil {
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

// ParseFIDSubmitted is a log parse operation binding the contract event 0x1cb0a3e71b09453c250a1f0ee46e90780c08d698ebed04ac4f0509b168fce7df.
//
// Solidity: event FIDSubmitted(address indexed _requester, string indexed fid, uint256 indexed id)
func (_Gateway *GatewayFilterer) ParseFIDSubmitted(log types.Log) (*GatewayFIDSubmitted, error) {
	event := new(GatewayFIDSubmitted)
	if err := _Gateway.contract.UnpackLog(event, "FIDSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Gateway contract.
type GatewayOwnershipTransferredIterator struct {
	Event *GatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayOwnershipTransferred)
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
		it.Event = new(GatewayOwnershipTransferred)
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
func (it *GatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayOwnershipTransferred represents a OwnershipTransferred event raised by the Gateway contract.
type GatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gateway *GatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayOwnershipTransferredIterator{contract: _Gateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gateway *GatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayOwnershipTransferred)
				if err := _Gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseOwnershipTransferred(log types.Log) (*GatewayOwnershipTransferred, error) {
	event := new(GatewayOwnershipTransferred)
	if err := _Gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayRequestGeneralIDIterator is returned from FilterRequestGeneralID and is used to iterate over the raw logs and unpacked data for RequestGeneralID events raised by the Gateway contract.
type GatewayRequestGeneralIDIterator struct {
	Event *GatewayRequestGeneralID // Event containing the contract specifics and raw log

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
func (it *GatewayRequestGeneralIDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayRequestGeneralID)
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
		it.Event = new(GatewayRequestGeneralID)
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
func (it *GatewayRequestGeneralIDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayRequestGeneralIDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayRequestGeneralID represents a RequestGeneralID event raised by the Gateway contract.
type GatewayRequestGeneralID struct {
	Requester common.Address
	Id        *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestGeneralID is a free log retrieval operation binding the contract event 0x600710de1dac8e9a1eae6575978854662c36f5f0c8243b13db0ace2f0bf4de4f.
//
// Solidity: event RequestGeneralID(address indexed requester, uint256 id)
func (_Gateway *GatewayFilterer) FilterRequestGeneralID(opts *bind.FilterOpts, requester []common.Address) (*GatewayRequestGeneralIDIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "RequestGeneralID", requesterRule)
	if err != nil {
		return nil, err
	}
	return &GatewayRequestGeneralIDIterator{contract: _Gateway.contract, event: "RequestGeneralID", logs: logs, sub: sub}, nil
}

// WatchRequestGeneralID is a free log subscription operation binding the contract event 0x600710de1dac8e9a1eae6575978854662c36f5f0c8243b13db0ace2f0bf4de4f.
//
// Solidity: event RequestGeneralID(address indexed requester, uint256 id)
func (_Gateway *GatewayFilterer) WatchRequestGeneralID(opts *bind.WatchOpts, sink chan<- *GatewayRequestGeneralID, requester []common.Address) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "RequestGeneralID", requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayRequestGeneralID)
				if err := _Gateway.contract.UnpackLog(event, "RequestGeneralID", log); err != nil {
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

// ParseRequestGeneralID is a log parse operation binding the contract event 0x600710de1dac8e9a1eae6575978854662c36f5f0c8243b13db0ace2f0bf4de4f.
//
// Solidity: event RequestGeneralID(address indexed requester, uint256 id)
func (_Gateway *GatewayFilterer) ParseRequestGeneralID(log types.Log) (*GatewayRequestGeneralID, error) {
	event := new(GatewayRequestGeneralID)
	if err := _Gateway.contract.UnpackLog(event, "RequestGeneralID", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayRequestGeneralKeyIterator is returned from FilterRequestGeneralKey and is used to iterate over the raw logs and unpacked data for RequestGeneralKey events raised by the Gateway contract.
type GatewayRequestGeneralKeyIterator struct {
	Event *GatewayRequestGeneralKey // Event containing the contract specifics and raw log

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
func (it *GatewayRequestGeneralKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayRequestGeneralKey)
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
		it.Event = new(GatewayRequestGeneralKey)
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
func (it *GatewayRequestGeneralKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayRequestGeneralKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayRequestGeneralKey represents a RequestGeneralKey event raised by the Gateway contract.
type GatewayRequestGeneralKey struct {
	Requester common.Address
	Id        *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestGeneralKey is a free log retrieval operation binding the contract event 0x901589c3a1797e9706a6e47144e3ef3ef75e1f57718e89f6036c98ee6a6754f7.
//
// Solidity: event RequestGeneralKey(address indexed requester, uint256 id)
func (_Gateway *GatewayFilterer) FilterRequestGeneralKey(opts *bind.FilterOpts, requester []common.Address) (*GatewayRequestGeneralKeyIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "RequestGeneralKey", requesterRule)
	if err != nil {
		return nil, err
	}
	return &GatewayRequestGeneralKeyIterator{contract: _Gateway.contract, event: "RequestGeneralKey", logs: logs, sub: sub}, nil
}

// WatchRequestGeneralKey is a free log subscription operation binding the contract event 0x901589c3a1797e9706a6e47144e3ef3ef75e1f57718e89f6036c98ee6a6754f7.
//
// Solidity: event RequestGeneralKey(address indexed requester, uint256 id)
func (_Gateway *GatewayFilterer) WatchRequestGeneralKey(opts *bind.WatchOpts, sink chan<- *GatewayRequestGeneralKey, requester []common.Address) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "RequestGeneralKey", requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayRequestGeneralKey)
				if err := _Gateway.contract.UnpackLog(event, "RequestGeneralKey", log); err != nil {
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

// ParseRequestGeneralKey is a log parse operation binding the contract event 0x901589c3a1797e9706a6e47144e3ef3ef75e1f57718e89f6036c98ee6a6754f7.
//
// Solidity: event RequestGeneralKey(address indexed requester, uint256 id)
func (_Gateway *GatewayFilterer) ParseRequestGeneralKey(log types.Log) (*GatewayRequestGeneralKey, error) {
	event := new(GatewayRequestGeneralKey)
	if err := _Gateway.contract.UnpackLog(event, "RequestGeneralKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
