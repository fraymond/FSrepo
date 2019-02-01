// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"stormcatcher/lib/common"
	"stormcatcher/vnode/accounts/abi"
	"stormcatcher/vnode/accounts/abi/bind"
	"stormcatcher/vnode/core/types"
	"strings"
)

// MainABI is the input ABI used to generate the binding from.
const MainABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"shardMapping\",\"outputs\":[{\"name\":\"shardId\",\"type\":\"uint256\"},{\"name\":\"nodeCount\",\"type\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\"},{\"name\":\"size\",\"type\":\"uint256\"},{\"name\":\"availableSize\",\"type\":\"uint256\"},{\"name\":\"percentage\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recentlyUsedList\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"shardCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[]\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"have\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pos\",\"type\":\"uint256\"},{\"name\":\"tosend\",\"type\":\"address[]\"},{\"name\":\"amount\",\"type\":\"uint256[]\"},{\"name\":\"times\",\"type\":\"uint256[]\"}],\"name\":\"postFlush\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"recentlyUsed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fileMapping\",\"outputs\":[{\"name\":\"fileId\",\"type\":\"uint256\"},{\"name\":\"fileHash\",\"type\":\"string\"},{\"name\":\"fileName\",\"type\":\"string\"},{\"name\":\"fileSize\",\"type\":\"uint256\"},{\"name\":\"fileOwner\",\"type\":\"address\"},{\"name\":\"createTime\",\"type\":\"uint256\"},{\"name\":\"verifiedCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fileList\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"shardNodeList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newlist\",\"type\":\"address[]\"}],\"name\":\"updateNodeList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllShards\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"setBlockVerificationInterval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fileCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddr\",\"type\":\"address\"},{\"name\":\"pos\",\"type\":\"uint256\"}],\"name\":\"getRedeemMapping\",\"outputs\":[{\"name\":\"redeemingAddr\",\"type\":\"address[]\"},{\"name\":\"redeemingAmt\",\"type\":\"uint256[]\"},{\"name\":\"redeemingtime\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeFromAddressArray\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurNodeList\",\"outputs\":[{\"name\":\"nodeList\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"curNodeList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"getEnterRecords\",\"outputs\":[{\"name\":\"enterAmt\",\"type\":\"uint256[]\"},{\"name\":\"entertime\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"shardList\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"array\",\"type\":\"uint256[]\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeFromArray\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"verifyGroupId\",\"type\":\"address\"},{\"name\":\"verifyNodeId\",\"type\":\"address\"},{\"name\":\"votingNodeId\",\"type\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"fileHash\",\"type\":\"string\"},{\"name\":\"shardId\",\"type\":\"uint256\"}],\"name\":\"voteVerifyTransaction\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"fileId\",\"type\":\"uint256\"},{\"name\":\"ipfsId\",\"type\":\"string\"}],\"name\":\"readFile\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"verifyGroupId\",\"type\":\"address\"},{\"name\":\"verifyNodeId\",\"type\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"fileHash\",\"type\":\"string\"},{\"name\":\"shardId\",\"type\":\"uint256\"}],\"name\":\"submitVerifyTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val1\",\"type\":\"uint256\"},{\"name\":\"val2\",\"type\":\"uint256\"}],\"name\":\"min\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"fileHash\",\"type\":\"string\"},{\"name\":\"fileName\",\"type\":\"string\"},{\"name\":\"fileSize\",\"type\":\"uint256\"},{\"name\":\"createTime\",\"type\":\"uint256\"}],\"name\":\"addFile\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"scsId\",\"type\":\"address\"},{\"name\":\"beneficiary\",\"type\":\"address\"},{\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"addNode\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"redeemFromMicroChain\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"fileHash\",\"type\":\"string\"},{\"name\":\"fileName\",\"type\":\"string\"},{\"name\":\"fileSize\",\"type\":\"uint256\"},{\"name\":\"createTime\",\"type\":\"uint256\"},{\"name\":\"ipfsId\",\"type\":\"string\"}],\"name\":\"addFile\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"shardSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"addShard\",\"outputs\":[{\"name\":\"shardId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"fileId\",\"type\":\"uint256\"}],\"name\":\"removeFile\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"myAddr\",\"type\":\"address\"}],\"name\":\"getMyFileHashes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"shardId\",\"type\":\"uint256\"}],\"name\":\"getAllFilesByShard\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"scsId\",\"type\":\"address\"}],\"name\":\"removeNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"unassignedNoteList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setAwardAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"setShardSize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"weight\",\"type\":\"uint256\"},{\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"setCapacity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"fileId\",\"type\":\"uint256\"}],\"name\":\"getFileById\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"capacityMapping\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"enterPos\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeShardMapping\",\"outputs\":[{\"name\":\"shardId\",\"type\":\"uint256\"},{\"name\":\"nodeCount\",\"type\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\"},{\"name\":\"size\",\"type\":\"uint256\"},{\"name\":\"availableSize\",\"type\":\"uint256\"},{\"name\":\"percentage\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeMapping\",\"outputs\":[{\"name\":\"shardId\",\"type\":\"uint256\"},{\"name\":\"scsId\",\"type\":\"address\"},{\"name\":\"beneficiary\",\"type\":\"address\"},{\"name\":\"size\",\"type\":\"uint256\"},{\"name\":\"lastVerifiedBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"award\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"verifyGroupMapping\",\"outputs\":[{\"name\":\"scsId\",\"type\":\"address\"},{\"name\":\"verifyNodeId\",\"type\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"fileHash\",\"type\":\"string\"},{\"name\":\"totalCount\",\"type\":\"uint256\"},{\"name\":\"votedCount\",\"type\":\"uint256\"},{\"name\":\"affirmCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"}]"

// Main is an auto generated Go binding around an MoacNode contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
}

// MainCaller is an auto generated read-only Go binding around an MoacNode contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an MoacNode contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an MoacNode contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an MoacNode contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an MoacNode contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an MoacNode contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an MoacNode contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an MoacNode contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins( address) constant returns(uint256)
func (_Main *MainCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "admins", arg0)
	return *ret0, err
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins( address) constant returns(uint256)
func (_Main *MainSession) Admins(arg0 common.Address) (*big.Int, error) {
	return _Main.Contract.Admins(&_Main.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins( address) constant returns(uint256)
func (_Main *MainCallerSession) Admins(arg0 common.Address) (*big.Int, error) {
	return _Main.Contract.Admins(&_Main.CallOpts, arg0)
}

// CapacityMapping is a free data retrieval call binding the contract method 0xe687929e.
//
// Solidity: function capacityMapping( uint256) constant returns(uint256)
func (_Main *MainCaller) CapacityMapping(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "capacityMapping", arg0)
	return *ret0, err
}

// CapacityMapping is a free data retrieval call binding the contract method 0xe687929e.
//
// Solidity: function capacityMapping( uint256) constant returns(uint256)
func (_Main *MainSession) CapacityMapping(arg0 *big.Int) (*big.Int, error) {
	return _Main.Contract.CapacityMapping(&_Main.CallOpts, arg0)
}

// CapacityMapping is a free data retrieval call binding the contract method 0xe687929e.
//
// Solidity: function capacityMapping( uint256) constant returns(uint256)
func (_Main *MainCallerSession) CapacityMapping(arg0 *big.Int) (*big.Int, error) {
	return _Main.Contract.CapacityMapping(&_Main.CallOpts, arg0)
}

// CurNodeList is a free data retrieval call binding the contract method 0x5b1222b4.
//
// Solidity: function curNodeList( uint256) constant returns(address)
func (_Main *MainCaller) CurNodeList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "curNodeList", arg0)
	return *ret0, err
}

// CurNodeList is a free data retrieval call binding the contract method 0x5b1222b4.
//
// Solidity: function curNodeList( uint256) constant returns(address)
func (_Main *MainSession) CurNodeList(arg0 *big.Int) (common.Address, error) {
	return _Main.Contract.CurNodeList(&_Main.CallOpts, arg0)
}

// CurNodeList is a free data retrieval call binding the contract method 0x5b1222b4.
//
// Solidity: function curNodeList( uint256) constant returns(address)
func (_Main *MainCallerSession) CurNodeList(arg0 *big.Int) (common.Address, error) {
	return _Main.Contract.CurNodeList(&_Main.CallOpts, arg0)
}

// EnterPos is a free data retrieval call binding the contract method 0xf2096c3b.
//
// Solidity: function enterPos() constant returns(uint256)
func (_Main *MainCaller) EnterPos(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "enterPos")
	return *ret0, err
}

// EnterPos is a free data retrieval call binding the contract method 0xf2096c3b.
//
// Solidity: function enterPos() constant returns(uint256)
func (_Main *MainSession) EnterPos() (*big.Int, error) {
	return _Main.Contract.EnterPos(&_Main.CallOpts)
}

// EnterPos is a free data retrieval call binding the contract method 0xf2096c3b.
//
// Solidity: function enterPos() constant returns(uint256)
func (_Main *MainCallerSession) EnterPos() (*big.Int, error) {
	return _Main.Contract.EnterPos(&_Main.CallOpts)
}

// FileCount is a free data retrieval call binding the contract method 0x43953519.
//
// Solidity: function fileCount() constant returns(uint256)
func (_Main *MainCaller) FileCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "fileCount")
	return *ret0, err
}

// FileCount is a free data retrieval call binding the contract method 0x43953519.
//
// Solidity: function fileCount() constant returns(uint256)
func (_Main *MainSession) FileCount() (*big.Int, error) {
	return _Main.Contract.FileCount(&_Main.CallOpts)
}

// FileCount is a free data retrieval call binding the contract method 0x43953519.
//
// Solidity: function fileCount() constant returns(uint256)
func (_Main *MainCallerSession) FileCount() (*big.Int, error) {
	return _Main.Contract.FileCount(&_Main.CallOpts)
}

// FileList is a free data retrieval call binding the contract method 0x273ccaa2.
//
// Solidity: function fileList( uint256) constant returns(uint256)
func (_Main *MainCaller) FileList(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "fileList", arg0)
	return *ret0, err
}

// FileList is a free data retrieval call binding the contract method 0x273ccaa2.
//
// Solidity: function fileList( uint256) constant returns(uint256)
func (_Main *MainSession) FileList(arg0 *big.Int) (*big.Int, error) {
	return _Main.Contract.FileList(&_Main.CallOpts, arg0)
}

// FileList is a free data retrieval call binding the contract method 0x273ccaa2.
//
// Solidity: function fileList( uint256) constant returns(uint256)
func (_Main *MainCallerSession) FileList(arg0 *big.Int) (*big.Int, error) {
	return _Main.Contract.FileList(&_Main.CallOpts, arg0)
}

// FileMapping is a free data retrieval call binding the contract method 0x209217af.
//
// Solidity: function fileMapping( uint256) constant returns(fileId uint256, fileHash string, fileName string, fileSize uint256, fileOwner address, createTime uint256, verifiedCount uint256)
func (_Main *MainCaller) FileMapping(opts *bind.CallOpts, arg0 *big.Int) (struct {
	FileId        *big.Int
	FileHash      string
	FileName      string
	FileSize      *big.Int
	FileOwner     common.Address
	CreateTime    *big.Int
	VerifiedCount *big.Int
}, error) {
	ret := new(struct {
		FileId        *big.Int
		FileHash      string
		FileName      string
		FileSize      *big.Int
		FileOwner     common.Address
		CreateTime    *big.Int
		VerifiedCount *big.Int
	})
	out := ret
	err := _Main.contract.Call(opts, out, "fileMapping", arg0)
	return *ret, err
}

// FileMapping is a free data retrieval call binding the contract method 0x209217af.
//
// Solidity: function fileMapping( uint256) constant returns(fileId uint256, fileHash string, fileName string, fileSize uint256, fileOwner address, createTime uint256, verifiedCount uint256)
func (_Main *MainSession) FileMapping(arg0 *big.Int) (struct {
	FileId        *big.Int
	FileHash      string
	FileName      string
	FileSize      *big.Int
	FileOwner     common.Address
	CreateTime    *big.Int
	VerifiedCount *big.Int
}, error) {
	return _Main.Contract.FileMapping(&_Main.CallOpts, arg0)
}

// FileMapping is a free data retrieval call binding the contract method 0x209217af.
//
// Solidity: function fileMapping( uint256) constant returns(fileId uint256, fileHash string, fileName string, fileSize uint256, fileOwner address, createTime uint256, verifiedCount uint256)
func (_Main *MainCallerSession) FileMapping(arg0 *big.Int) (struct {
	FileId        *big.Int
	FileHash      string
	FileName      string
	FileSize      *big.Int
	FileOwner     common.Address
	CreateTime    *big.Int
	VerifiedCount *big.Int
}, error) {
	return _Main.Contract.FileMapping(&_Main.CallOpts, arg0)
}

// GetAllFilesByShard is a free data retrieval call binding the contract method 0xa3dc2c22.
//
// Solidity: function getAllFilesByShard(shardId uint256) constant returns(uint256[])
func (_Main *MainCaller) GetAllFilesByShard(opts *bind.CallOpts, shardId *big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "getAllFilesByShard", shardId)
	return *ret0, err
}

// GetAllFilesByShard is a free data retrieval call binding the contract method 0xa3dc2c22.
//
// Solidity: function getAllFilesByShard(shardId uint256) constant returns(uint256[])
func (_Main *MainSession) GetAllFilesByShard(shardId *big.Int) ([]*big.Int, error) {
	return _Main.Contract.GetAllFilesByShard(&_Main.CallOpts, shardId)
}

// GetAllFilesByShard is a free data retrieval call binding the contract method 0xa3dc2c22.
//
// Solidity: function getAllFilesByShard(shardId uint256) constant returns(uint256[])
func (_Main *MainCallerSession) GetAllFilesByShard(shardId *big.Int) ([]*big.Int, error) {
	return _Main.Contract.GetAllFilesByShard(&_Main.CallOpts, shardId)
}

// GetAllShards is a free data retrieval call binding the contract method 0x3ec43aad.
//
// Solidity: function getAllShards() constant returns(uint256[])
func (_Main *MainCaller) GetAllShards(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "getAllShards")
	return *ret0, err
}

// GetAllShards is a free data retrieval call binding the contract method 0x3ec43aad.
//
// Solidity: function getAllShards() constant returns(uint256[])
func (_Main *MainSession) GetAllShards() ([]*big.Int, error) {
	return _Main.Contract.GetAllShards(&_Main.CallOpts)
}

// GetAllShards is a free data retrieval call binding the contract method 0x3ec43aad.
//
// Solidity: function getAllShards() constant returns(uint256[])
func (_Main *MainCallerSession) GetAllShards() ([]*big.Int, error) {
	return _Main.Contract.GetAllShards(&_Main.CallOpts)
}

// GetCurNodeList is a free data retrieval call binding the contract method 0x57058a33.
//
// Solidity: function getCurNodeList() constant returns(nodeList address[])
func (_Main *MainCaller) GetCurNodeList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "getCurNodeList")
	return *ret0, err
}

// GetCurNodeList is a free data retrieval call binding the contract method 0x57058a33.
//
// Solidity: function getCurNodeList() constant returns(nodeList address[])
func (_Main *MainSession) GetCurNodeList() ([]common.Address, error) {
	return _Main.Contract.GetCurNodeList(&_Main.CallOpts)
}

// GetCurNodeList is a free data retrieval call binding the contract method 0x57058a33.
//
// Solidity: function getCurNodeList() constant returns(nodeList address[])
func (_Main *MainCallerSession) GetCurNodeList() ([]common.Address, error) {
	return _Main.Contract.GetCurNodeList(&_Main.CallOpts)
}

// GetEnterRecords is a free data retrieval call binding the contract method 0x5b15270c.
//
// Solidity: function getEnterRecords(userAddr address) constant returns(enterAmt uint256[], entertime uint256[])
func (_Main *MainCaller) GetEnterRecords(opts *bind.CallOpts, userAddr common.Address) (struct {
	EnterAmt  []*big.Int
	Entertime []*big.Int
}, error) {
	ret := new(struct {
		EnterAmt  []*big.Int
		Entertime []*big.Int
	})
	out := ret
	err := _Main.contract.Call(opts, out, "getEnterRecords", userAddr)
	return *ret, err
}

// GetEnterRecords is a free data retrieval call binding the contract method 0x5b15270c.
//
// Solidity: function getEnterRecords(userAddr address) constant returns(enterAmt uint256[], entertime uint256[])
func (_Main *MainSession) GetEnterRecords(userAddr common.Address) (struct {
	EnterAmt  []*big.Int
	Entertime []*big.Int
}, error) {
	return _Main.Contract.GetEnterRecords(&_Main.CallOpts, userAddr)
}

// GetEnterRecords is a free data retrieval call binding the contract method 0x5b15270c.
//
// Solidity: function getEnterRecords(userAddr address) constant returns(enterAmt uint256[], entertime uint256[])
func (_Main *MainCallerSession) GetEnterRecords(userAddr common.Address) (struct {
	EnterAmt  []*big.Int
	Entertime []*big.Int
}, error) {
	return _Main.Contract.GetEnterRecords(&_Main.CallOpts, userAddr)
}

// GetFileById is a free data retrieval call binding the contract method 0xe4e6025e.
//
// Solidity: function getFileById(fileId uint256) constant returns(string, string, uint256, address, uint256, uint256)
func (_Main *MainCaller) GetFileById(opts *bind.CallOpts, fileId *big.Int) (string, string, *big.Int, common.Address, *big.Int, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(*big.Int)
		ret3 = new(common.Address)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _Main.contract.Call(opts, out, "getFileById", fileId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetFileById is a free data retrieval call binding the contract method 0xe4e6025e.
//
// Solidity: function getFileById(fileId uint256) constant returns(string, string, uint256, address, uint256, uint256)
func (_Main *MainSession) GetFileById(fileId *big.Int) (string, string, *big.Int, common.Address, *big.Int, *big.Int, error) {
	return _Main.Contract.GetFileById(&_Main.CallOpts, fileId)
}

// GetFileById is a free data retrieval call binding the contract method 0xe4e6025e.
//
// Solidity: function getFileById(fileId uint256) constant returns(string, string, uint256, address, uint256, uint256)
func (_Main *MainCallerSession) GetFileById(fileId *big.Int) (string, string, *big.Int, common.Address, *big.Int, *big.Int, error) {
	return _Main.Contract.GetFileById(&_Main.CallOpts, fileId)
}

// GetMyFileHashes is a free data retrieval call binding the contract method 0xa358b17e.
//
// Solidity: function getMyFileHashes(myAddr address) constant returns(uint256[])
func (_Main *MainCaller) GetMyFileHashes(opts *bind.CallOpts, myAddr common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "getMyFileHashes", myAddr)
	return *ret0, err
}

// GetMyFileHashes is a free data retrieval call binding the contract method 0xa358b17e.
//
// Solidity: function getMyFileHashes(myAddr address) constant returns(uint256[])
func (_Main *MainSession) GetMyFileHashes(myAddr common.Address) ([]*big.Int, error) {
	return _Main.Contract.GetMyFileHashes(&_Main.CallOpts, myAddr)
}

// GetMyFileHashes is a free data retrieval call binding the contract method 0xa358b17e.
//
// Solidity: function getMyFileHashes(myAddr address) constant returns(uint256[])
func (_Main *MainCallerSession) GetMyFileHashes(myAddr common.Address) ([]*big.Int, error) {
	return _Main.Contract.GetMyFileHashes(&_Main.CallOpts, myAddr)
}

// GetRedeemMapping is a free data retrieval call binding the contract method 0x442e95b1.
//
// Solidity: function getRedeemMapping(userAddr address, pos uint256) constant returns(redeemingAddr address[], redeemingAmt uint256[], redeemingtime uint256[])
func (_Main *MainCaller) GetRedeemMapping(opts *bind.CallOpts, userAddr common.Address, pos *big.Int) (struct {
	RedeemingAddr []common.Address
	RedeemingAmt  []*big.Int
	Redeemingtime []*big.Int
}, error) {
	ret := new(struct {
		RedeemingAddr []common.Address
		RedeemingAmt  []*big.Int
		Redeemingtime []*big.Int
	})
	out := ret
	err := _Main.contract.Call(opts, out, "getRedeemMapping", userAddr, pos)
	return *ret, err
}

// GetRedeemMapping is a free data retrieval call binding the contract method 0x442e95b1.
//
// Solidity: function getRedeemMapping(userAddr address, pos uint256) constant returns(redeemingAddr address[], redeemingAmt uint256[], redeemingtime uint256[])
func (_Main *MainSession) GetRedeemMapping(userAddr common.Address, pos *big.Int) (struct {
	RedeemingAddr []common.Address
	RedeemingAmt  []*big.Int
	Redeemingtime []*big.Int
}, error) {
	return _Main.Contract.GetRedeemMapping(&_Main.CallOpts, userAddr, pos)
}

// GetRedeemMapping is a free data retrieval call binding the contract method 0x442e95b1.
//
// Solidity: function getRedeemMapping(userAddr address, pos uint256) constant returns(redeemingAddr address[], redeemingAmt uint256[], redeemingtime uint256[])
func (_Main *MainCallerSession) GetRedeemMapping(userAddr common.Address, pos *big.Int) (struct {
	RedeemingAddr []common.Address
	RedeemingAmt  []*big.Int
	Redeemingtime []*big.Int
}, error) {
	return _Main.Contract.GetRedeemMapping(&_Main.CallOpts, userAddr, pos)
}

// Have is a free data retrieval call binding the contract method 0x10ba0499.
//
// Solidity: function have(addrs address[], addr address) constant returns(bool)
func (_Main *MainCaller) Have(opts *bind.CallOpts, addrs []common.Address, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "have", addrs, addr)
	return *ret0, err
}

// Have is a free data retrieval call binding the contract method 0x10ba0499.
//
// Solidity: function have(addrs address[], addr address) constant returns(bool)
func (_Main *MainSession) Have(addrs []common.Address, addr common.Address) (bool, error) {
	return _Main.Contract.Have(&_Main.CallOpts, addrs, addr)
}

// Have is a free data retrieval call binding the contract method 0x10ba0499.
//
// Solidity: function have(addrs address[], addr address) constant returns(bool)
func (_Main *MainCallerSession) Have(addrs []common.Address, addr common.Address) (bool, error) {
	return _Main.Contract.Have(&_Main.CallOpts, addrs, addr)
}

// NodeMapping is a free data retrieval call binding the contract method 0xfbd1b4ce.
//
// Solidity: function nodeMapping( address) constant returns(shardId uint256, scsId address, beneficiary address, size uint256, lastVerifiedBlock uint256)
func (_Main *MainCaller) NodeMapping(opts *bind.CallOpts, arg0 common.Address) (struct {
	ShardId           *big.Int
	ScsId             common.Address
	Beneficiary       common.Address
	Size              *big.Int
	LastVerifiedBlock *big.Int
}, error) {
	ret := new(struct {
		ShardId           *big.Int
		ScsId             common.Address
		Beneficiary       common.Address
		Size              *big.Int
		LastVerifiedBlock *big.Int
	})
	out := ret
	err := _Main.contract.Call(opts, out, "nodeMapping", arg0)
	return *ret, err
}

// NodeMapping is a free data retrieval call binding the contract method 0xfbd1b4ce.
//
// Solidity: function nodeMapping( address) constant returns(shardId uint256, scsId address, beneficiary address, size uint256, lastVerifiedBlock uint256)
func (_Main *MainSession) NodeMapping(arg0 common.Address) (struct {
	ShardId           *big.Int
	ScsId             common.Address
	Beneficiary       common.Address
	Size              *big.Int
	LastVerifiedBlock *big.Int
}, error) {
	return _Main.Contract.NodeMapping(&_Main.CallOpts, arg0)
}

// NodeMapping is a free data retrieval call binding the contract method 0xfbd1b4ce.
//
// Solidity: function nodeMapping( address) constant returns(shardId uint256, scsId address, beneficiary address, size uint256, lastVerifiedBlock uint256)
func (_Main *MainCallerSession) NodeMapping(arg0 common.Address) (struct {
	ShardId           *big.Int
	ScsId             common.Address
	Beneficiary       common.Address
	Size              *big.Int
	LastVerifiedBlock *big.Int
}, error) {
	return _Main.Contract.NodeMapping(&_Main.CallOpts, arg0)
}

// NodeShardMapping is a free data retrieval call binding the contract method 0xf501c8eb.
//
// Solidity: function nodeShardMapping( address) constant returns(shardId uint256, nodeCount uint256, weight uint256, size uint256, availableSize uint256, percentage uint256)
func (_Main *MainCaller) NodeShardMapping(opts *bind.CallOpts, arg0 common.Address) (struct {
	ShardId       *big.Int
	NodeCount     *big.Int
	Weight        *big.Int
	Size          *big.Int
	AvailableSize *big.Int
	Percentage    *big.Int
}, error) {
	ret := new(struct {
		ShardId       *big.Int
		NodeCount     *big.Int
		Weight        *big.Int
		Size          *big.Int
		AvailableSize *big.Int
		Percentage    *big.Int
	})
	out := ret
	err := _Main.contract.Call(opts, out, "nodeShardMapping", arg0)
	return *ret, err
}

// NodeShardMapping is a free data retrieval call binding the contract method 0xf501c8eb.
//
// Solidity: function nodeShardMapping( address) constant returns(shardId uint256, nodeCount uint256, weight uint256, size uint256, availableSize uint256, percentage uint256)
func (_Main *MainSession) NodeShardMapping(arg0 common.Address) (struct {
	ShardId       *big.Int
	NodeCount     *big.Int
	Weight        *big.Int
	Size          *big.Int
	AvailableSize *big.Int
	Percentage    *big.Int
}, error) {
	return _Main.Contract.NodeShardMapping(&_Main.CallOpts, arg0)
}

// NodeShardMapping is a free data retrieval call binding the contract method 0xf501c8eb.
//
// Solidity: function nodeShardMapping( address) constant returns(shardId uint256, nodeCount uint256, weight uint256, size uint256, availableSize uint256, percentage uint256)
func (_Main *MainCallerSession) NodeShardMapping(arg0 common.Address) (struct {
	ShardId       *big.Int
	NodeCount     *big.Int
	Weight        *big.Int
	Size          *big.Int
	AvailableSize *big.Int
	Percentage    *big.Int
}, error) {
	return _Main.Contract.NodeShardMapping(&_Main.CallOpts, arg0)
}

// RecentlyUsedList is a free data retrieval call binding the contract method 0x04500519.
//
// Solidity: function recentlyUsedList( uint256) constant returns(uint256)
func (_Main *MainCaller) RecentlyUsedList(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "recentlyUsedList", arg0)
	return *ret0, err
}

// RecentlyUsedList is a free data retrieval call binding the contract method 0x04500519.
//
// Solidity: function recentlyUsedList( uint256) constant returns(uint256)
func (_Main *MainSession) RecentlyUsedList(arg0 *big.Int) (*big.Int, error) {
	return _Main.Contract.RecentlyUsedList(&_Main.CallOpts, arg0)
}

// RecentlyUsedList is a free data retrieval call binding the contract method 0x04500519.
//
// Solidity: function recentlyUsedList( uint256) constant returns(uint256)
func (_Main *MainCallerSession) RecentlyUsedList(arg0 *big.Int) (*big.Int, error) {
	return _Main.Contract.RecentlyUsedList(&_Main.CallOpts, arg0)
}

// ShardCount is a free data retrieval call binding the contract method 0x04e9c77a.
//
// Solidity: function shardCount() constant returns(uint256)
func (_Main *MainCaller) ShardCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "shardCount")
	return *ret0, err
}

// ShardCount is a free data retrieval call binding the contract method 0x04e9c77a.
//
// Solidity: function shardCount() constant returns(uint256)
func (_Main *MainSession) ShardCount() (*big.Int, error) {
	return _Main.Contract.ShardCount(&_Main.CallOpts)
}

// ShardCount is a free data retrieval call binding the contract method 0x04e9c77a.
//
// Solidity: function shardCount() constant returns(uint256)
func (_Main *MainCallerSession) ShardCount() (*big.Int, error) {
	return _Main.Contract.ShardCount(&_Main.CallOpts)
}

// ShardList is a free data retrieval call binding the contract method 0x608c2db5.
//
// Solidity: function shardList( uint256) constant returns(uint256)
func (_Main *MainCaller) ShardList(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "shardList", arg0)
	return *ret0, err
}

// ShardList is a free data retrieval call binding the contract method 0x608c2db5.
//
// Solidity: function shardList( uint256) constant returns(uint256)
func (_Main *MainSession) ShardList(arg0 *big.Int) (*big.Int, error) {
	return _Main.Contract.ShardList(&_Main.CallOpts, arg0)
}

// ShardList is a free data retrieval call binding the contract method 0x608c2db5.
//
// Solidity: function shardList( uint256) constant returns(uint256)
func (_Main *MainCallerSession) ShardList(arg0 *big.Int) (*big.Int, error) {
	return _Main.Contract.ShardList(&_Main.CallOpts, arg0)
}

// ShardMapping is a free data retrieval call binding the contract method 0x02671018.
//
// Solidity: function shardMapping( uint256) constant returns(shardId uint256, nodeCount uint256, weight uint256, size uint256, availableSize uint256, percentage uint256)
func (_Main *MainCaller) ShardMapping(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ShardId       *big.Int
	NodeCount     *big.Int
	Weight        *big.Int
	Size          *big.Int
	AvailableSize *big.Int
	Percentage    *big.Int
}, error) {
	ret := new(struct {
		ShardId       *big.Int
		NodeCount     *big.Int
		Weight        *big.Int
		Size          *big.Int
		AvailableSize *big.Int
		Percentage    *big.Int
	})
	out := ret
	err := _Main.contract.Call(opts, out, "shardMapping", arg0)
	return *ret, err
}

// ShardMapping is a free data retrieval call binding the contract method 0x02671018.
//
// Solidity: function shardMapping( uint256) constant returns(shardId uint256, nodeCount uint256, weight uint256, size uint256, availableSize uint256, percentage uint256)
func (_Main *MainSession) ShardMapping(arg0 *big.Int) (struct {
	ShardId       *big.Int
	NodeCount     *big.Int
	Weight        *big.Int
	Size          *big.Int
	AvailableSize *big.Int
	Percentage    *big.Int
}, error) {
	return _Main.Contract.ShardMapping(&_Main.CallOpts, arg0)
}

// ShardMapping is a free data retrieval call binding the contract method 0x02671018.
//
// Solidity: function shardMapping( uint256) constant returns(shardId uint256, nodeCount uint256, weight uint256, size uint256, availableSize uint256, percentage uint256)
func (_Main *MainCallerSession) ShardMapping(arg0 *big.Int) (struct {
	ShardId       *big.Int
	NodeCount     *big.Int
	Weight        *big.Int
	Size          *big.Int
	AvailableSize *big.Int
	Percentage    *big.Int
}, error) {
	return _Main.Contract.ShardMapping(&_Main.CallOpts, arg0)
}

// ShardNodeList is a free data retrieval call binding the contract method 0x27a49763.
//
// Solidity: function shardNodeList( uint256,  uint256) constant returns(address)
func (_Main *MainCaller) ShardNodeList(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "shardNodeList", arg0, arg1)
	return *ret0, err
}

// ShardNodeList is a free data retrieval call binding the contract method 0x27a49763.
//
// Solidity: function shardNodeList( uint256,  uint256) constant returns(address)
func (_Main *MainSession) ShardNodeList(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Main.Contract.ShardNodeList(&_Main.CallOpts, arg0, arg1)
}

// ShardNodeList is a free data retrieval call binding the contract method 0x27a49763.
//
// Solidity: function shardNodeList( uint256,  uint256) constant returns(address)
func (_Main *MainCallerSession) ShardNodeList(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Main.Contract.ShardNodeList(&_Main.CallOpts, arg0, arg1)
}

// ShardSize is a free data retrieval call binding the contract method 0x96f33858.
//
// Solidity: function shardSize() constant returns(uint256)
func (_Main *MainCaller) ShardSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "shardSize")
	return *ret0, err
}

// ShardSize is a free data retrieval call binding the contract method 0x96f33858.
//
// Solidity: function shardSize() constant returns(uint256)
func (_Main *MainSession) ShardSize() (*big.Int, error) {
	return _Main.Contract.ShardSize(&_Main.CallOpts)
}

// ShardSize is a free data retrieval call binding the contract method 0x96f33858.
//
// Solidity: function shardSize() constant returns(uint256)
func (_Main *MainCallerSession) ShardSize() (*big.Int, error) {
	return _Main.Contract.ShardSize(&_Main.CallOpts)
}

// UnassignedNoteList is a free data retrieval call binding the contract method 0xd1e3887c.
//
// Solidity: function unassignedNoteList( uint256) constant returns(address)
func (_Main *MainCaller) UnassignedNoteList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "unassignedNoteList", arg0)
	return *ret0, err
}

// UnassignedNoteList is a free data retrieval call binding the contract method 0xd1e3887c.
//
// Solidity: function unassignedNoteList( uint256) constant returns(address)
func (_Main *MainSession) UnassignedNoteList(arg0 *big.Int) (common.Address, error) {
	return _Main.Contract.UnassignedNoteList(&_Main.CallOpts, arg0)
}

// UnassignedNoteList is a free data retrieval call binding the contract method 0xd1e3887c.
//
// Solidity: function unassignedNoteList( uint256) constant returns(address)
func (_Main *MainCallerSession) UnassignedNoteList(arg0 *big.Int) (common.Address, error) {
	return _Main.Contract.UnassignedNoteList(&_Main.CallOpts, arg0)
}

// VerifyGroupMapping is a free data retrieval call binding the contract method 0xfc54fc21.
//
// Solidity: function verifyGroupMapping( address,  uint256) constant returns(scsId address, verifyNodeId address, blockNumber uint256, fileHash string, totalCount uint256, votedCount uint256, affirmCount uint256)
func (_Main *MainCaller) VerifyGroupMapping(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	ScsId        common.Address
	VerifyNodeId common.Address
	BlockNumber  *big.Int
	FileHash     string
	TotalCount   *big.Int
	VotedCount   *big.Int
	AffirmCount  *big.Int
}, error) {
	ret := new(struct {
		ScsId        common.Address
		VerifyNodeId common.Address
		BlockNumber  *big.Int
		FileHash     string
		TotalCount   *big.Int
		VotedCount   *big.Int
		AffirmCount  *big.Int
	})
	out := ret
	err := _Main.contract.Call(opts, out, "verifyGroupMapping", arg0, arg1)
	return *ret, err
}

// VerifyGroupMapping is a free data retrieval call binding the contract method 0xfc54fc21.
//
// Solidity: function verifyGroupMapping( address,  uint256) constant returns(scsId address, verifyNodeId address, blockNumber uint256, fileHash string, totalCount uint256, votedCount uint256, affirmCount uint256)
func (_Main *MainSession) VerifyGroupMapping(arg0 common.Address, arg1 *big.Int) (struct {
	ScsId        common.Address
	VerifyNodeId common.Address
	BlockNumber  *big.Int
	FileHash     string
	TotalCount   *big.Int
	VotedCount   *big.Int
	AffirmCount  *big.Int
}, error) {
	return _Main.Contract.VerifyGroupMapping(&_Main.CallOpts, arg0, arg1)
}

// VerifyGroupMapping is a free data retrieval call binding the contract method 0xfc54fc21.
//
// Solidity: function verifyGroupMapping( address,  uint256) constant returns(scsId address, verifyNodeId address, blockNumber uint256, fileHash string, totalCount uint256, votedCount uint256, affirmCount uint256)
func (_Main *MainCallerSession) VerifyGroupMapping(arg0 common.Address, arg1 *big.Int) (struct {
	ScsId        common.Address
	VerifyNodeId common.Address
	BlockNumber  *big.Int
	FileHash     string
	TotalCount   *big.Int
	VotedCount   *big.Int
	AffirmCount  *big.Int
}, error) {
	return _Main.Contract.VerifyGroupMapping(&_Main.CallOpts, arg0, arg1)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(admin address) returns()
func (_Main *MainTransactor) AddAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addAdmin", admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(admin address) returns()
func (_Main *MainSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddAdmin(&_Main.TransactOpts, admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(admin address) returns()
func (_Main *MainTransactorSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddAdmin(&_Main.TransactOpts, admin)
}

// AddFile is a paid mutator transaction binding the contract method 0x90414068.
//
// Solidity: function addFile(fileHash string, fileName string, fileSize uint256, createTime uint256, ipfsId string) returns(uint256, uint256)
func (_Main *MainTransactor) AddFile(opts *bind.TransactOpts, fileHash string, fileName string, fileSize *big.Int, createTime *big.Int, ipfsId string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addFile", fileHash, fileName, fileSize, createTime, ipfsId)
}

// AddFile is a paid mutator transaction binding the contract method 0x90414068.
//
// Solidity: function addFile(fileHash string, fileName string, fileSize uint256, createTime uint256, ipfsId string) returns(uint256, uint256)
func (_Main *MainSession) AddFile(fileHash string, fileName string, fileSize *big.Int, createTime *big.Int, ipfsId string) (*types.Transaction, error) {
	return _Main.Contract.AddFile(&_Main.TransactOpts, fileHash, fileName, fileSize, createTime, ipfsId)
}

// AddFile is a paid mutator transaction binding the contract method 0x90414068.
//
// Solidity: function addFile(fileHash string, fileName string, fileSize uint256, createTime uint256, ipfsId string) returns(uint256, uint256)
func (_Main *MainTransactorSession) AddFile(fileHash string, fileName string, fileSize *big.Int, createTime *big.Int, ipfsId string) (*types.Transaction, error) {
	return _Main.Contract.AddFile(&_Main.TransactOpts, fileHash, fileName, fileSize, createTime, ipfsId)
}

// AddNode is a paid mutator transaction binding the contract method 0x841fd1db.
//
// Solidity: function addNode(scsId address, beneficiary address, weight uint256) returns(uint256)
func (_Main *MainTransactor) AddNode(opts *bind.TransactOpts, scsId common.Address, beneficiary common.Address, weight *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addNode", scsId, beneficiary, weight)
}

// AddNode is a paid mutator transaction binding the contract method 0x841fd1db.
//
// Solidity: function addNode(scsId address, beneficiary address, weight uint256) returns(uint256)
func (_Main *MainSession) AddNode(scsId common.Address, beneficiary common.Address, weight *big.Int) (*types.Transaction, error) {
	return _Main.Contract.AddNode(&_Main.TransactOpts, scsId, beneficiary, weight)
}

// AddNode is a paid mutator transaction binding the contract method 0x841fd1db.
//
// Solidity: function addNode(scsId address, beneficiary address, weight uint256) returns(uint256)
func (_Main *MainTransactorSession) AddNode(scsId common.Address, beneficiary common.Address, weight *big.Int) (*types.Transaction, error) {
	return _Main.Contract.AddNode(&_Main.TransactOpts, scsId, beneficiary, weight)
}

// AddShard is a paid mutator transaction binding the contract method 0x97319ed2.
//
// Solidity: function addShard(weight uint256) returns(shardId uint256)
func (_Main *MainTransactor) AddShard(opts *bind.TransactOpts, weight *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addShard", weight)
}

// AddShard is a paid mutator transaction binding the contract method 0x97319ed2.
//
// Solidity: function addShard(weight uint256) returns(shardId uint256)
func (_Main *MainSession) AddShard(weight *big.Int) (*types.Transaction, error) {
	return _Main.Contract.AddShard(&_Main.TransactOpts, weight)
}

// AddShard is a paid mutator transaction binding the contract method 0x97319ed2.
//
// Solidity: function addShard(weight uint256) returns(shardId uint256)
func (_Main *MainTransactorSession) AddShard(weight *big.Int) (*types.Transaction, error) {
	return _Main.Contract.AddShard(&_Main.TransactOpts, weight)
}

// Award is a paid mutator transaction binding the contract method 0xfbdd25e4.
//
// Solidity: function award(beneficiary address) returns()
func (_Main *MainTransactor) Award(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "award", beneficiary)
}

// Award is a paid mutator transaction binding the contract method 0xfbdd25e4.
//
// Solidity: function award(beneficiary address) returns()
func (_Main *MainSession) Award(beneficiary common.Address) (*types.Transaction, error) {
	return _Main.Contract.Award(&_Main.TransactOpts, beneficiary)
}

// Award is a paid mutator transaction binding the contract method 0xfbdd25e4.
//
// Solidity: function award(beneficiary address) returns()
func (_Main *MainTransactorSession) Award(beneficiary common.Address) (*types.Transaction, error) {
	return _Main.Contract.Award(&_Main.TransactOpts, beneficiary)
}

// Min is a paid mutator transaction binding the contract method 0x7ae2b5c7.
//
// Solidity: function min(val1 uint256, val2 uint256) returns(uint256)
func (_Main *MainTransactor) Min(opts *bind.TransactOpts, val1 *big.Int, val2 *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "min", val1, val2)
}

// Min is a paid mutator transaction binding the contract method 0x7ae2b5c7.
//
// Solidity: function min(val1 uint256, val2 uint256) returns(uint256)
func (_Main *MainSession) Min(val1 *big.Int, val2 *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Min(&_Main.TransactOpts, val1, val2)
}

// Min is a paid mutator transaction binding the contract method 0x7ae2b5c7.
//
// Solidity: function min(val1 uint256, val2 uint256) returns(uint256)
func (_Main *MainTransactorSession) Min(val1 *big.Int, val2 *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Min(&_Main.TransactOpts, val1, val2)
}

// PostFlush is a paid mutator transaction binding the contract method 0x12df9412.
//
// Solidity: function postFlush(pos uint256, tosend address[], amount uint256[], times uint256[]) returns()
func (_Main *MainTransactor) PostFlush(opts *bind.TransactOpts, pos *big.Int, tosend []common.Address, amount []*big.Int, times []*big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "postFlush", pos, tosend, amount, times)
}

// PostFlush is a paid mutator transaction binding the contract method 0x12df9412.
//
// Solidity: function postFlush(pos uint256, tosend address[], amount uint256[], times uint256[]) returns()
func (_Main *MainSession) PostFlush(pos *big.Int, tosend []common.Address, amount []*big.Int, times []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.PostFlush(&_Main.TransactOpts, pos, tosend, amount, times)
}

// PostFlush is a paid mutator transaction binding the contract method 0x12df9412.
//
// Solidity: function postFlush(pos uint256, tosend address[], amount uint256[], times uint256[]) returns()
func (_Main *MainTransactorSession) PostFlush(pos *big.Int, tosend []common.Address, amount []*big.Int, times []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.PostFlush(&_Main.TransactOpts, pos, tosend, amount, times)
}

// ReadFile is a paid mutator transaction binding the contract method 0x6616174a.
//
// Solidity: function readFile(fileId uint256, ipfsId string) returns(bool)
func (_Main *MainTransactor) ReadFile(opts *bind.TransactOpts, fileId *big.Int, ipfsId string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "readFile", fileId, ipfsId)
}

// ReadFile is a paid mutator transaction binding the contract method 0x6616174a.
//
// Solidity: function readFile(fileId uint256, ipfsId string) returns(bool)
func (_Main *MainSession) ReadFile(fileId *big.Int, ipfsId string) (*types.Transaction, error) {
	return _Main.Contract.ReadFile(&_Main.TransactOpts, fileId, ipfsId)
}

// ReadFile is a paid mutator transaction binding the contract method 0x6616174a.
//
// Solidity: function readFile(fileId uint256, ipfsId string) returns(bool)
func (_Main *MainTransactorSession) ReadFile(fileId *big.Int, ipfsId string) (*types.Transaction, error) {
	return _Main.Contract.ReadFile(&_Main.TransactOpts, fileId, ipfsId)
}

// RecentlyUsed is a paid mutator transaction binding the contract method 0x1feaa5dc.
//
// Solidity: function recentlyUsed(value uint256) returns(bool)
func (_Main *MainTransactor) RecentlyUsed(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "recentlyUsed", value)
}

// RecentlyUsed is a paid mutator transaction binding the contract method 0x1feaa5dc.
//
// Solidity: function recentlyUsed(value uint256) returns(bool)
func (_Main *MainSession) RecentlyUsed(value *big.Int) (*types.Transaction, error) {
	return _Main.Contract.RecentlyUsed(&_Main.TransactOpts, value)
}

// RecentlyUsed is a paid mutator transaction binding the contract method 0x1feaa5dc.
//
// Solidity: function recentlyUsed(value uint256) returns(bool)
func (_Main *MainTransactorSession) RecentlyUsed(value *big.Int) (*types.Transaction, error) {
	return _Main.Contract.RecentlyUsed(&_Main.TransactOpts, value)
}

// RedeemFromMicroChain is a paid mutator transaction binding the contract method 0x89739c5b.
//
// Solidity: function redeemFromMicroChain() returns()
func (_Main *MainTransactor) RedeemFromMicroChain(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "redeemFromMicroChain")
}

// RedeemFromMicroChain is a paid mutator transaction binding the contract method 0x89739c5b.
//
// Solidity: function redeemFromMicroChain() returns()
func (_Main *MainSession) RedeemFromMicroChain() (*types.Transaction, error) {
	return _Main.Contract.RedeemFromMicroChain(&_Main.TransactOpts)
}

// RedeemFromMicroChain is a paid mutator transaction binding the contract method 0x89739c5b.
//
// Solidity: function redeemFromMicroChain() returns()
func (_Main *MainTransactorSession) RedeemFromMicroChain() (*types.Transaction, error) {
	return _Main.Contract.RedeemFromMicroChain(&_Main.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(admin address) returns()
func (_Main *MainTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(admin address) returns()
func (_Main *MainSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _Main.Contract.RemoveAdmin(&_Main.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(admin address) returns()
func (_Main *MainTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _Main.Contract.RemoveAdmin(&_Main.TransactOpts, admin)
}

// RemoveFile is a paid mutator transaction binding the contract method 0x9926b03f.
//
// Solidity: function removeFile(fileId uint256) returns(bool)
func (_Main *MainTransactor) RemoveFile(opts *bind.TransactOpts, fileId *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeFile", fileId)
}

// RemoveFile is a paid mutator transaction binding the contract method 0x9926b03f.
//
// Solidity: function removeFile(fileId uint256) returns(bool)
func (_Main *MainSession) RemoveFile(fileId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveFile(&_Main.TransactOpts, fileId)
}

// RemoveFile is a paid mutator transaction binding the contract method 0x9926b03f.
//
// Solidity: function removeFile(fileId uint256) returns(bool)
func (_Main *MainTransactorSession) RemoveFile(fileId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveFile(&_Main.TransactOpts, fileId)
}

// RemoveFromAddressArray is a paid mutator transaction binding the contract method 0x4cedb37d.
//
// Solidity: function removeFromAddressArray(index uint256) returns()
func (_Main *MainTransactor) RemoveFromAddressArray(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeFromAddressArray", index)
}

// RemoveFromAddressArray is a paid mutator transaction binding the contract method 0x4cedb37d.
//
// Solidity: function removeFromAddressArray(index uint256) returns()
func (_Main *MainSession) RemoveFromAddressArray(index *big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveFromAddressArray(&_Main.TransactOpts, index)
}

// RemoveFromAddressArray is a paid mutator transaction binding the contract method 0x4cedb37d.
//
// Solidity: function removeFromAddressArray(index uint256) returns()
func (_Main *MainTransactorSession) RemoveFromAddressArray(index *big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveFromAddressArray(&_Main.TransactOpts, index)
}

// RemoveFromArray is a paid mutator transaction binding the contract method 0x628d08b8.
//
// Solidity: function removeFromArray(array uint256[], index uint256) returns(uint256[])
func (_Main *MainTransactor) RemoveFromArray(opts *bind.TransactOpts, array []*big.Int, index *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeFromArray", array, index)
}

// RemoveFromArray is a paid mutator transaction binding the contract method 0x628d08b8.
//
// Solidity: function removeFromArray(array uint256[], index uint256) returns(uint256[])
func (_Main *MainSession) RemoveFromArray(array []*big.Int, index *big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveFromArray(&_Main.TransactOpts, array, index)
}

// RemoveFromArray is a paid mutator transaction binding the contract method 0x628d08b8.
//
// Solidity: function removeFromArray(array uint256[], index uint256) returns(uint256[])
func (_Main *MainTransactorSession) RemoveFromArray(array []*big.Int, index *big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveFromArray(&_Main.TransactOpts, array, index)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xb2b99ec9.
//
// Solidity: function removeNode(scsId address) returns(bool)
func (_Main *MainTransactor) RemoveNode(opts *bind.TransactOpts, scsId common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeNode", scsId)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xb2b99ec9.
//
// Solidity: function removeNode(scsId address) returns(bool)
func (_Main *MainSession) RemoveNode(scsId common.Address) (*types.Transaction, error) {
	return _Main.Contract.RemoveNode(&_Main.TransactOpts, scsId)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xb2b99ec9.
//
// Solidity: function removeNode(scsId address) returns(bool)
func (_Main *MainTransactorSession) RemoveNode(scsId common.Address) (*types.Transaction, error) {
	return _Main.Contract.RemoveNode(&_Main.TransactOpts, scsId)
}

// SetAwardAmount is a paid mutator transaction binding the contract method 0xd2706c19.
//
// Solidity: function setAwardAmount(amount uint256) returns()
func (_Main *MainTransactor) SetAwardAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setAwardAmount", amount)
}

// SetAwardAmount is a paid mutator transaction binding the contract method 0xd2706c19.
//
// Solidity: function setAwardAmount(amount uint256) returns()
func (_Main *MainSession) SetAwardAmount(amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetAwardAmount(&_Main.TransactOpts, amount)
}

// SetAwardAmount is a paid mutator transaction binding the contract method 0xd2706c19.
//
// Solidity: function setAwardAmount(amount uint256) returns()
func (_Main *MainTransactorSession) SetAwardAmount(amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetAwardAmount(&_Main.TransactOpts, amount)
}

// SetBlockVerificationInterval is a paid mutator transaction binding the contract method 0x40d46f2c.
//
// Solidity: function setBlockVerificationInterval(num uint256) returns()
func (_Main *MainTransactor) SetBlockVerificationInterval(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setBlockVerificationInterval", num)
}

// SetBlockVerificationInterval is a paid mutator transaction binding the contract method 0x40d46f2c.
//
// Solidity: function setBlockVerificationInterval(num uint256) returns()
func (_Main *MainSession) SetBlockVerificationInterval(num *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetBlockVerificationInterval(&_Main.TransactOpts, num)
}

// SetBlockVerificationInterval is a paid mutator transaction binding the contract method 0x40d46f2c.
//
// Solidity: function setBlockVerificationInterval(num uint256) returns()
func (_Main *MainTransactorSession) SetBlockVerificationInterval(num *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetBlockVerificationInterval(&_Main.TransactOpts, num)
}

// SetCapacity is a paid mutator transaction binding the contract method 0xdee2b058.
//
// Solidity: function setCapacity(weight uint256, size uint256) returns()
func (_Main *MainTransactor) SetCapacity(opts *bind.TransactOpts, weight *big.Int, size *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setCapacity", weight, size)
}

// SetCapacity is a paid mutator transaction binding the contract method 0xdee2b058.
//
// Solidity: function setCapacity(weight uint256, size uint256) returns()
func (_Main *MainSession) SetCapacity(weight *big.Int, size *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetCapacity(&_Main.TransactOpts, weight, size)
}

// SetCapacity is a paid mutator transaction binding the contract method 0xdee2b058.
//
// Solidity: function setCapacity(weight uint256, size uint256) returns()
func (_Main *MainTransactorSession) SetCapacity(weight *big.Int, size *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetCapacity(&_Main.TransactOpts, weight, size)
}

// SetShardSize is a paid mutator transaction binding the contract method 0xd9f13414.
//
// Solidity: function setShardSize(size uint256) returns()
func (_Main *MainTransactor) SetShardSize(opts *bind.TransactOpts, size *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setShardSize", size)
}

// SetShardSize is a paid mutator transaction binding the contract method 0xd9f13414.
//
// Solidity: function setShardSize(size uint256) returns()
func (_Main *MainSession) SetShardSize(size *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetShardSize(&_Main.TransactOpts, size)
}

// SetShardSize is a paid mutator transaction binding the contract method 0xd9f13414.
//
// Solidity: function setShardSize(size uint256) returns()
func (_Main *MainTransactorSession) SetShardSize(size *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SetShardSize(&_Main.TransactOpts, size)
}

// SubmitVerifyTransaction is a paid mutator transaction binding the contract method 0x681f3d3d.
//
// Solidity: function submitVerifyTransaction(verifyGroupId address, verifyNodeId address, blockNumber uint256, fileHash string, shardId uint256) returns()
func (_Main *MainTransactor) SubmitVerifyTransaction(opts *bind.TransactOpts, verifyGroupId common.Address, verifyNodeId common.Address, blockNumber *big.Int, fileHash string, shardId *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "submitVerifyTransaction", verifyGroupId, verifyNodeId, blockNumber, fileHash, shardId)
}

// SubmitVerifyTransaction is a paid mutator transaction binding the contract method 0x681f3d3d.
//
// Solidity: function submitVerifyTransaction(verifyGroupId address, verifyNodeId address, blockNumber uint256, fileHash string, shardId uint256) returns()
func (_Main *MainSession) SubmitVerifyTransaction(verifyGroupId common.Address, verifyNodeId common.Address, blockNumber *big.Int, fileHash string, shardId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SubmitVerifyTransaction(&_Main.TransactOpts, verifyGroupId, verifyNodeId, blockNumber, fileHash, shardId)
}

// SubmitVerifyTransaction is a paid mutator transaction binding the contract method 0x681f3d3d.
//
// Solidity: function submitVerifyTransaction(verifyGroupId address, verifyNodeId address, blockNumber uint256, fileHash string, shardId uint256) returns()
func (_Main *MainTransactorSession) SubmitVerifyTransaction(verifyGroupId common.Address, verifyNodeId common.Address, blockNumber *big.Int, fileHash string, shardId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.SubmitVerifyTransaction(&_Main.TransactOpts, verifyGroupId, verifyNodeId, blockNumber, fileHash, shardId)
}

// UpdateNodeList is a paid mutator transaction binding the contract method 0x31a06771.
//
// Solidity: function updateNodeList(newlist address[]) returns()
func (_Main *MainTransactor) UpdateNodeList(opts *bind.TransactOpts, newlist []common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "updateNodeList", newlist)
}

// UpdateNodeList is a paid mutator transaction binding the contract method 0x31a06771.
//
// Solidity: function updateNodeList(newlist address[]) returns()
func (_Main *MainSession) UpdateNodeList(newlist []common.Address) (*types.Transaction, error) {
	return _Main.Contract.UpdateNodeList(&_Main.TransactOpts, newlist)
}

// UpdateNodeList is a paid mutator transaction binding the contract method 0x31a06771.
//
// Solidity: function updateNodeList(newlist address[]) returns()
func (_Main *MainTransactorSession) UpdateNodeList(newlist []common.Address) (*types.Transaction, error) {
	return _Main.Contract.UpdateNodeList(&_Main.TransactOpts, newlist)
}

// VoteVerifyTransaction is a paid mutator transaction binding the contract method 0x6446c074.
//
// Solidity: function voteVerifyTransaction(verifyGroupId address, verifyNodeId address, votingNodeId address, blockNumber uint256, fileHash string, shardId uint256) returns(bool)
func (_Main *MainTransactor) VoteVerifyTransaction(opts *bind.TransactOpts, verifyGroupId common.Address, verifyNodeId common.Address, votingNodeId common.Address, blockNumber *big.Int, fileHash string, shardId *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "voteVerifyTransaction", verifyGroupId, verifyNodeId, votingNodeId, blockNumber, fileHash, shardId)
}

// VoteVerifyTransaction is a paid mutator transaction binding the contract method 0x6446c074.
//
// Solidity: function voteVerifyTransaction(verifyGroupId address, verifyNodeId address, votingNodeId address, blockNumber uint256, fileHash string, shardId uint256) returns(bool)
func (_Main *MainSession) VoteVerifyTransaction(verifyGroupId common.Address, verifyNodeId common.Address, votingNodeId common.Address, blockNumber *big.Int, fileHash string, shardId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.VoteVerifyTransaction(&_Main.TransactOpts, verifyGroupId, verifyNodeId, votingNodeId, blockNumber, fileHash, shardId)
}

// VoteVerifyTransaction is a paid mutator transaction binding the contract method 0x6446c074.
//
// Solidity: function voteVerifyTransaction(verifyGroupId address, verifyNodeId address, votingNodeId address, blockNumber uint256, fileHash string, shardId uint256) returns(bool)
func (_Main *MainTransactorSession) VoteVerifyTransaction(verifyGroupId common.Address, verifyNodeId common.Address, votingNodeId common.Address, blockNumber *big.Int, fileHash string, shardId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.VoteVerifyTransaction(&_Main.TransactOpts, verifyGroupId, verifyNodeId, votingNodeId, blockNumber, fileHash, shardId)
}
