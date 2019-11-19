// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BASManagerSimpleABI is the input ABI used to generate the binding from.
const BASManagerSimpleABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"y\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"rent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bc\",\"type\":\"bytes32\"}],\"name\":\"queryByBCAddress\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"},{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"},{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newPrice\",\"type\":\"uint16\"}],\"name\":\"changePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashKey\",\"type\":\"bytes32\"}],\"name\":\"queryByHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"},{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"},{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"queryByString\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"},{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"},{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"DataRecords\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"IPv4\",\"type\":\"bytes4\"},{\"internalType\":\"bytes16\",\"name\":\"IPv6\",\"type\":\"bytes16\"},{\"internalType\":\"bytes1\",\"name\":\"BCLength\",\"type\":\"bytes1\"},{\"internalType\":\"bytes32\",\"name\":\"BCAddress\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"EthAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"y\",\"type\":\"uint8\"}],\"name\":\"checkRent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"change\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// BASManagerSimple is an auto generated Go binding around an Ethereum contract.
type BASManagerSimple struct {
	BASManagerSimpleCaller     // Read-only binding to the contract
	BASManagerSimpleTransactor // Write-only binding to the contract
	BASManagerSimpleFilterer   // Log filterer for contract events
}

// BASManagerSimpleCaller is an auto generated read-only Go binding around an Ethereum contract.
type BASManagerSimpleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BASManagerSimpleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BASManagerSimpleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BASManagerSimpleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BASManagerSimpleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BASManagerSimpleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BASManagerSimpleSession struct {
	Contract     *BASManagerSimple // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BASManagerSimpleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BASManagerSimpleCallerSession struct {
	Contract *BASManagerSimpleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BASManagerSimpleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BASManagerSimpleTransactorSession struct {
	Contract     *BASManagerSimpleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BASManagerSimpleRaw is an auto generated low-level Go binding around an Ethereum contract.
type BASManagerSimpleRaw struct {
	Contract *BASManagerSimple // Generic contract binding to access the raw methods on
}

// BASManagerSimpleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BASManagerSimpleCallerRaw struct {
	Contract *BASManagerSimpleCaller // Generic read-only contract binding to access the raw methods on
}

// BASManagerSimpleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BASManagerSimpleTransactorRaw struct {
	Contract *BASManagerSimpleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBASManagerSimple creates a new instance of BASManagerSimple, bound to a specific deployed contract.
func NewBASManagerSimple(address common.Address, backend bind.ContractBackend) (*BASManagerSimple, error) {
	contract, err := bindBASManagerSimple(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BASManagerSimple{BASManagerSimpleCaller: BASManagerSimpleCaller{contract: contract}, BASManagerSimpleTransactor: BASManagerSimpleTransactor{contract: contract}, BASManagerSimpleFilterer: BASManagerSimpleFilterer{contract: contract}}, nil
}

// NewBASManagerSimpleCaller creates a new read-only instance of BASManagerSimple, bound to a specific deployed contract.
func NewBASManagerSimpleCaller(address common.Address, caller bind.ContractCaller) (*BASManagerSimpleCaller, error) {
	contract, err := bindBASManagerSimple(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BASManagerSimpleCaller{contract: contract}, nil
}

// NewBASManagerSimpleTransactor creates a new write-only instance of BASManagerSimple, bound to a specific deployed contract.
func NewBASManagerSimpleTransactor(address common.Address, transactor bind.ContractTransactor) (*BASManagerSimpleTransactor, error) {
	contract, err := bindBASManagerSimple(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BASManagerSimpleTransactor{contract: contract}, nil
}

// NewBASManagerSimpleFilterer creates a new log filterer instance of BASManagerSimple, bound to a specific deployed contract.
func NewBASManagerSimpleFilterer(address common.Address, filterer bind.ContractFilterer) (*BASManagerSimpleFilterer, error) {
	contract, err := bindBASManagerSimple(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BASManagerSimpleFilterer{contract: contract}, nil
}

// bindBASManagerSimple binds a generic wrapper to an already deployed contract.
func bindBASManagerSimple(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BASManagerSimpleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BASManagerSimple *BASManagerSimpleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BASManagerSimple.Contract.BASManagerSimpleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BASManagerSimple *BASManagerSimpleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BASManagerSimple.Contract.BASManagerSimpleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BASManagerSimple *BASManagerSimpleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BASManagerSimple.Contract.BASManagerSimpleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BASManagerSimple *BASManagerSimpleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BASManagerSimple.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BASManagerSimple *BASManagerSimpleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BASManagerSimple.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BASManagerSimple *BASManagerSimpleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BASManagerSimple.Contract.contract.Transact(opts, method, params...)
}

// DataRecords is a free data retrieval call binding the contract method 0xb62242a3.
//
// Solidity: function DataRecords(bytes32 ) constant returns(bytes4 IPv4, bytes16 IPv6, bytes1 BCLength, bytes32 BCAddress, address EthAddress)
func (_BASManagerSimple *BASManagerSimpleCaller) DataRecords(opts *bind.CallOpts, arg0 [32]byte) (struct {
	IPv4       [4]byte
	IPv6       [16]byte
	BCLength   [1]byte
	BCAddress  [32]byte
	EthAddress common.Address
}, error) {
	ret := new(struct {
		IPv4       [4]byte
		IPv6       [16]byte
		BCLength   [1]byte
		BCAddress  [32]byte
		EthAddress common.Address
	})
	out := ret
	err := _BASManagerSimple.contract.Call(opts, out, "DataRecords", arg0)
	return *ret, err
}

// DataRecords is a free data retrieval call binding the contract method 0xb62242a3.
//
// Solidity: function DataRecords(bytes32 ) constant returns(bytes4 IPv4, bytes16 IPv6, bytes1 BCLength, bytes32 BCAddress, address EthAddress)
func (_BASManagerSimple *BASManagerSimpleSession) DataRecords(arg0 [32]byte) (struct {
	IPv4       [4]byte
	IPv6       [16]byte
	BCLength   [1]byte
	BCAddress  [32]byte
	EthAddress common.Address
}, error) {
	return _BASManagerSimple.Contract.DataRecords(&_BASManagerSimple.CallOpts, arg0)
}

// DataRecords is a free data retrieval call binding the contract method 0xb62242a3.
//
// Solidity: function DataRecords(bytes32 ) constant returns(bytes4 IPv4, bytes16 IPv6, bytes1 BCLength, bytes32 BCAddress, address EthAddress)
func (_BASManagerSimple *BASManagerSimpleCallerSession) DataRecords(arg0 [32]byte) (struct {
	IPv4       [4]byte
	IPv6       [16]byte
	BCLength   [1]byte
	BCAddress  [32]byte
	EthAddress common.Address
}, error) {
	return _BASManagerSimple.Contract.DataRecords(&_BASManagerSimple.CallOpts, arg0)
}

// CheckAllowance is a free data retrieval call binding the contract method 0x9405de64.
//
// Solidity: function checkAllowance() constant returns(uint256)
func (_BASManagerSimple *BASManagerSimpleCaller) CheckAllowance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BASManagerSimple.contract.Call(opts, out, "checkAllowance")
	return *ret0, err
}

// CheckAllowance is a free data retrieval call binding the contract method 0x9405de64.
//
// Solidity: function checkAllowance() constant returns(uint256)
func (_BASManagerSimple *BASManagerSimpleSession) CheckAllowance() (*big.Int, error) {
	return _BASManagerSimple.Contract.CheckAllowance(&_BASManagerSimple.CallOpts)
}

// CheckAllowance is a free data retrieval call binding the contract method 0x9405de64.
//
// Solidity: function checkAllowance() constant returns(uint256)
func (_BASManagerSimple *BASManagerSimpleCallerSession) CheckAllowance() (*big.Int, error) {
	return _BASManagerSimple.Contract.CheckAllowance(&_BASManagerSimple.CallOpts)
}

// CheckRent is a free data retrieval call binding the contract method 0xd5dc648a.
//
// Solidity: function checkRent(uint8 y) constant returns(uint256)
func (_BASManagerSimple *BASManagerSimpleCaller) CheckRent(opts *bind.CallOpts, y uint8) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BASManagerSimple.contract.Call(opts, out, "checkRent", y)
	return *ret0, err
}

// CheckRent is a free data retrieval call binding the contract method 0xd5dc648a.
//
// Solidity: function checkRent(uint8 y) constant returns(uint256)
func (_BASManagerSimple *BASManagerSimpleSession) CheckRent(y uint8) (*big.Int, error) {
	return _BASManagerSimple.Contract.CheckRent(&_BASManagerSimple.CallOpts, y)
}

// CheckRent is a free data retrieval call binding the contract method 0xd5dc648a.
//
// Solidity: function checkRent(uint8 y) constant returns(uint256)
func (_BASManagerSimple *BASManagerSimpleCallerSession) CheckRent(y uint8) (*big.Int, error) {
	return _BASManagerSimple.Contract.CheckRent(&_BASManagerSimple.CallOpts, y)
}

// Hash is a free data retrieval call binding the contract method 0xb411ee94.
//
// Solidity: function hash(string key) constant returns(bytes32)
func (_BASManagerSimple *BASManagerSimpleCaller) Hash(opts *bind.CallOpts, key string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BASManagerSimple.contract.Call(opts, out, "hash", key)
	return *ret0, err
}

// Hash is a free data retrieval call binding the contract method 0xb411ee94.
//
// Solidity: function hash(string key) constant returns(bytes32)
func (_BASManagerSimple *BASManagerSimpleSession) Hash(key string) ([32]byte, error) {
	return _BASManagerSimple.Contract.Hash(&_BASManagerSimple.CallOpts, key)
}

// Hash is a free data retrieval call binding the contract method 0xb411ee94.
//
// Solidity: function hash(string key) constant returns(bytes32)
func (_BASManagerSimple *BASManagerSimpleCallerSession) Hash(key string) ([32]byte, error) {
	return _BASManagerSimple.Contract.Hash(&_BASManagerSimple.CallOpts, key)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BASManagerSimple *BASManagerSimpleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BASManagerSimple.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BASManagerSimple *BASManagerSimpleSession) Owner() (common.Address, error) {
	return _BASManagerSimple.Contract.Owner(&_BASManagerSimple.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BASManagerSimple *BASManagerSimpleCallerSession) Owner() (common.Address, error) {
	return _BASManagerSimple.Contract.Owner(&_BASManagerSimple.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_BASManagerSimple *BASManagerSimpleCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BASManagerSimple.contract.Call(opts, out, "price")
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_BASManagerSimple *BASManagerSimpleSession) Price() (*big.Int, error) {
	return _BASManagerSimple.Contract.Price(&_BASManagerSimple.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_BASManagerSimple *BASManagerSimpleCallerSession) Price() (*big.Int, error) {
	return _BASManagerSimple.Contract.Price(&_BASManagerSimple.CallOpts)
}

// QueryByBCAddress is a free data retrieval call binding the contract method 0x6e5c0a4b.
//
// Solidity: function queryByBCAddress(bytes32 bc) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32, address)
func (_BASManagerSimple *BASManagerSimpleCaller) QueryByBCAddress(opts *bind.CallOpts, bc [32]byte) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, common.Address, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([4]byte)
		ret2 = new([16]byte)
		ret3 = new([1]byte)
		ret4 = new([32]byte)
		ret5 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _BASManagerSimple.contract.Call(opts, out, "queryByBCAddress", bc)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// QueryByBCAddress is a free data retrieval call binding the contract method 0x6e5c0a4b.
//
// Solidity: function queryByBCAddress(bytes32 bc) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32, address)
func (_BASManagerSimple *BASManagerSimpleSession) QueryByBCAddress(bc [32]byte) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, common.Address, error) {
	return _BASManagerSimple.Contract.QueryByBCAddress(&_BASManagerSimple.CallOpts, bc)
}

// QueryByBCAddress is a free data retrieval call binding the contract method 0x6e5c0a4b.
//
// Solidity: function queryByBCAddress(bytes32 bc) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32, address)
func (_BASManagerSimple *BASManagerSimpleCallerSession) QueryByBCAddress(bc [32]byte) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, common.Address, error) {
	return _BASManagerSimple.Contract.QueryByBCAddress(&_BASManagerSimple.CallOpts, bc)
}

// QueryByHash is a free data retrieval call binding the contract method 0xa2cb7c95.
//
// Solidity: function queryByHash(bytes32 hashKey) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32, address)
func (_BASManagerSimple *BASManagerSimpleCaller) QueryByHash(opts *bind.CallOpts, hashKey [32]byte) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, common.Address, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([4]byte)
		ret2 = new([16]byte)
		ret3 = new([1]byte)
		ret4 = new([32]byte)
		ret5 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _BASManagerSimple.contract.Call(opts, out, "queryByHash", hashKey)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// QueryByHash is a free data retrieval call binding the contract method 0xa2cb7c95.
//
// Solidity: function queryByHash(bytes32 hashKey) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32, address)
func (_BASManagerSimple *BASManagerSimpleSession) QueryByHash(hashKey [32]byte) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, common.Address, error) {
	return _BASManagerSimple.Contract.QueryByHash(&_BASManagerSimple.CallOpts, hashKey)
}

// QueryByHash is a free data retrieval call binding the contract method 0xa2cb7c95.
//
// Solidity: function queryByHash(bytes32 hashKey) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32, address)
func (_BASManagerSimple *BASManagerSimpleCallerSession) QueryByHash(hashKey [32]byte) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, common.Address, error) {
	return _BASManagerSimple.Contract.QueryByHash(&_BASManagerSimple.CallOpts, hashKey)
}

// QueryByString is a free data retrieval call binding the contract method 0xab5e6aa7.
//
// Solidity: function queryByString(string key) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32, address)
func (_BASManagerSimple *BASManagerSimpleCaller) QueryByString(opts *bind.CallOpts, key string) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, common.Address, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([4]byte)
		ret2 = new([16]byte)
		ret3 = new([1]byte)
		ret4 = new([32]byte)
		ret5 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _BASManagerSimple.contract.Call(opts, out, "queryByString", key)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// QueryByString is a free data retrieval call binding the contract method 0xab5e6aa7.
//
// Solidity: function queryByString(string key) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32, address)
func (_BASManagerSimple *BASManagerSimpleSession) QueryByString(key string) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, common.Address, error) {
	return _BASManagerSimple.Contract.QueryByString(&_BASManagerSimple.CallOpts, key)
}

// QueryByString is a free data retrieval call binding the contract method 0xab5e6aa7.
//
// Solidity: function queryByString(string key) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32, address)
func (_BASManagerSimple *BASManagerSimpleCallerSession) QueryByString(key string) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, common.Address, error) {
	return _BASManagerSimple.Contract.QueryByString(&_BASManagerSimple.CallOpts, key)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BASManagerSimple *BASManagerSimpleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BASManagerSimple.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BASManagerSimple *BASManagerSimpleSession) Token() (common.Address, error) {
	return _BASManagerSimple.Contract.Token(&_BASManagerSimple.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BASManagerSimple *BASManagerSimpleCallerSession) Token() (common.Address, error) {
	return _BASManagerSimple.Contract.Token(&_BASManagerSimple.CallOpts)
}

// Change is a paid mutator transaction binding the contract method 0xda6459af.
//
// Solidity: function change(string key, bytes data) returns()
func (_BASManagerSimple *BASManagerSimpleTransactor) Change(opts *bind.TransactOpts, key string, data []byte) (*types.Transaction, error) {
	return _BASManagerSimple.contract.Transact(opts, "change", key, data)
}

// Change is a paid mutator transaction binding the contract method 0xda6459af.
//
// Solidity: function change(string key, bytes data) returns()
func (_BASManagerSimple *BASManagerSimpleSession) Change(key string, data []byte) (*types.Transaction, error) {
	return _BASManagerSimple.Contract.Change(&_BASManagerSimple.TransactOpts, key, data)
}

// Change is a paid mutator transaction binding the contract method 0xda6459af.
//
// Solidity: function change(string key, bytes data) returns()
func (_BASManagerSimple *BASManagerSimpleTransactorSession) Change(key string, data []byte) (*types.Transaction, error) {
	return _BASManagerSimple.Contract.Change(&_BASManagerSimple.TransactOpts, key, data)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x7085b50a.
//
// Solidity: function changePrice(uint16 newPrice) returns(uint256)
func (_BASManagerSimple *BASManagerSimpleTransactor) ChangePrice(opts *bind.TransactOpts, newPrice uint16) (*types.Transaction, error) {
	return _BASManagerSimple.contract.Transact(opts, "changePrice", newPrice)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x7085b50a.
//
// Solidity: function changePrice(uint16 newPrice) returns(uint256)
func (_BASManagerSimple *BASManagerSimpleSession) ChangePrice(newPrice uint16) (*types.Transaction, error) {
	return _BASManagerSimple.Contract.ChangePrice(&_BASManagerSimple.TransactOpts, newPrice)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x7085b50a.
//
// Solidity: function changePrice(uint16 newPrice) returns(uint256)
func (_BASManagerSimple *BASManagerSimpleTransactorSession) ChangePrice(newPrice uint16) (*types.Transaction, error) {
	return _BASManagerSimple.Contract.ChangePrice(&_BASManagerSimple.TransactOpts, newPrice)
}

// Rent is a paid mutator transaction binding the contract method 0x0d944db7.
//
// Solidity: function rent(string key, uint8 y, bytes data) returns()
func (_BASManagerSimple *BASManagerSimpleTransactor) Rent(opts *bind.TransactOpts, key string, y uint8, data []byte) (*types.Transaction, error) {
	return _BASManagerSimple.contract.Transact(opts, "rent", key, y, data)
}

// Rent is a paid mutator transaction binding the contract method 0x0d944db7.
//
// Solidity: function rent(string key, uint8 y, bytes data) returns()
func (_BASManagerSimple *BASManagerSimpleSession) Rent(key string, y uint8, data []byte) (*types.Transaction, error) {
	return _BASManagerSimple.Contract.Rent(&_BASManagerSimple.TransactOpts, key, y, data)
}

// Rent is a paid mutator transaction binding the contract method 0x0d944db7.
//
// Solidity: function rent(string key, uint8 y, bytes data) returns()
func (_BASManagerSimple *BASManagerSimpleTransactorSession) Rent(key string, y uint8, data []byte) (*types.Transaction, error) {
	return _BASManagerSimple.Contract.Rent(&_BASManagerSimple.TransactOpts, key, y, data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BASManagerSimple *BASManagerSimpleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BASManagerSimple.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BASManagerSimple *BASManagerSimpleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BASManagerSimple.Contract.TransferOwnership(&_BASManagerSimple.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BASManagerSimple *BASManagerSimpleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BASManagerSimple.Contract.TransferOwnership(&_BASManagerSimple.TransactOpts, newOwner)
}
