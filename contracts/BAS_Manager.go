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

// BASManagerABI is the input ABI used to generate the binding from.
const BASManagerABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"y\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"rent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"IPv6\",\"type\":\"bytes16\"}],\"name\":\"queryByIPv6\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"},{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"},{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bc\",\"type\":\"bytes32\"}],\"name\":\"queryByBCAddress\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"},{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"},{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newPrice\",\"type\":\"uint16\"}],\"name\":\"changePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashKey\",\"type\":\"bytes32\"}],\"name\":\"queryByHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"},{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"},{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"queryByString\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"},{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"},{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"DataRecords\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"IPv4\",\"type\":\"bytes4\"},{\"internalType\":\"bytes16\",\"name\":\"IPv6\",\"type\":\"bytes16\"},{\"internalType\":\"bytes1\",\"name\":\"BCLength\",\"type\":\"bytes1\"},{\"internalType\":\"bytes32\",\"name\":\"BCAddress\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"EthAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"IPv4\",\"type\":\"bytes4\"}],\"name\":\"queryByIPv4\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"},{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"},{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"y\",\"type\":\"uint8\"}],\"name\":\"checkRent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"checkLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"change\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// BASManager is an auto generated Go binding around an Ethereum contract.
type BASManager struct {
	BASManagerCaller     // Read-only binding to the contract
	BASManagerTransactor // Write-only binding to the contract
	BASManagerFilterer   // Log filterer for contract events
}

// BASManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BASManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BASManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BASManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BASManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BASManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BASManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BASManagerSession struct {
	Contract     *BASManager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BASManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BASManagerCallerSession struct {
	Contract *BASManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BASManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BASManagerTransactorSession struct {
	Contract     *BASManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BASManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BASManagerRaw struct {
	Contract *BASManager // Generic contract binding to access the raw methods on
}

// BASManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BASManagerCallerRaw struct {
	Contract *BASManagerCaller // Generic read-only contract binding to access the raw methods on
}

// BASManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BASManagerTransactorRaw struct {
	Contract *BASManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBASManager creates a new instance of BASManager, bound to a specific deployed contract.
func NewBASManager(address common.Address, backend bind.ContractBackend) (*BASManager, error) {
	contract, err := bindBASManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BASManager{BASManagerCaller: BASManagerCaller{contract: contract}, BASManagerTransactor: BASManagerTransactor{contract: contract}, BASManagerFilterer: BASManagerFilterer{contract: contract}}, nil
}

// NewBASManagerCaller creates a new read-only instance of BASManager, bound to a specific deployed contract.
func NewBASManagerCaller(address common.Address, caller bind.ContractCaller) (*BASManagerCaller, error) {
	contract, err := bindBASManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BASManagerCaller{contract: contract}, nil
}

// NewBASManagerTransactor creates a new write-only instance of BASManager, bound to a specific deployed contract.
func NewBASManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*BASManagerTransactor, error) {
	contract, err := bindBASManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BASManagerTransactor{contract: contract}, nil
}

// NewBASManagerFilterer creates a new log filterer instance of BASManager, bound to a specific deployed contract.
func NewBASManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*BASManagerFilterer, error) {
	contract, err := bindBASManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BASManagerFilterer{contract: contract}, nil
}

// bindBASManager binds a generic wrapper to an already deployed contract.
func bindBASManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BASManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BASManager *BASManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BASManager.Contract.BASManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BASManager *BASManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BASManager.Contract.BASManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BASManager *BASManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BASManager.Contract.BASManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BASManager *BASManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BASManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BASManager *BASManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BASManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BASManager *BASManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BASManager.Contract.contract.Transact(opts, method, params...)
}

// DataRecords is a free data retrieval call binding the contract method 0xb62242a3.
//
// Solidity: function DataRecords(bytes32 ) constant returns(bytes4 IPv4, bytes16 IPv6, bytes1 BCLength, bytes32 BCAddress, address EthAddress)
func (_BASManager *BASManagerCaller) DataRecords(opts *bind.CallOpts, arg0 [32]byte) (struct {
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
	err := _BASManager.contract.Call(opts, out, "DataRecords", arg0)
	return *ret, err
}

// DataRecords is a free data retrieval call binding the contract method 0xb62242a3.
//
// Solidity: function DataRecords(bytes32 ) constant returns(bytes4 IPv4, bytes16 IPv6, bytes1 BCLength, bytes32 BCAddress, address EthAddress)
func (_BASManager *BASManagerSession) DataRecords(arg0 [32]byte) (struct {
	IPv4       [4]byte
	IPv6       [16]byte
	BCLength   [1]byte
	BCAddress  [32]byte
	EthAddress common.Address
}, error) {
	return _BASManager.Contract.DataRecords(&_BASManager.CallOpts, arg0)
}

// DataRecords is a free data retrieval call binding the contract method 0xb62242a3.
//
// Solidity: function DataRecords(bytes32 ) constant returns(bytes4 IPv4, bytes16 IPv6, bytes1 BCLength, bytes32 BCAddress, address EthAddress)
func (_BASManager *BASManagerCallerSession) DataRecords(arg0 [32]byte) (struct {
	IPv4       [4]byte
	IPv6       [16]byte
	BCLength   [1]byte
	BCAddress  [32]byte
	EthAddress common.Address
}, error) {
	return _BASManager.Contract.DataRecords(&_BASManager.CallOpts, arg0)
}

// CheckAllowance is a free data retrieval call binding the contract method 0x9405de64.
//
// Solidity: function checkAllowance() constant returns(uint256)
func (_BASManager *BASManagerCaller) CheckAllowance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BASManager.contract.Call(opts, out, "checkAllowance")
	return *ret0, err
}

// CheckAllowance is a free data retrieval call binding the contract method 0x9405de64.
//
// Solidity: function checkAllowance() constant returns(uint256)
func (_BASManager *BASManagerSession) CheckAllowance() (*big.Int, error) {
	return _BASManager.Contract.CheckAllowance(&_BASManager.CallOpts)
}

// CheckAllowance is a free data retrieval call binding the contract method 0x9405de64.
//
// Solidity: function checkAllowance() constant returns(uint256)
func (_BASManager *BASManagerCallerSession) CheckAllowance() (*big.Int, error) {
	return _BASManager.Contract.CheckAllowance(&_BASManager.CallOpts)
}

// CheckLength is a free data retrieval call binding the contract method 0xd79145b4.
//
// Solidity: function checkLength(bytes input) constant returns(uint256)
func (_BASManager *BASManagerCaller) CheckLength(opts *bind.CallOpts, input []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BASManager.contract.Call(opts, out, "checkLength", input)
	return *ret0, err
}

// CheckLength is a free data retrieval call binding the contract method 0xd79145b4.
//
// Solidity: function checkLength(bytes input) constant returns(uint256)
func (_BASManager *BASManagerSession) CheckLength(input []byte) (*big.Int, error) {
	return _BASManager.Contract.CheckLength(&_BASManager.CallOpts, input)
}

// CheckLength is a free data retrieval call binding the contract method 0xd79145b4.
//
// Solidity: function checkLength(bytes input) constant returns(uint256)
func (_BASManager *BASManagerCallerSession) CheckLength(input []byte) (*big.Int, error) {
	return _BASManager.Contract.CheckLength(&_BASManager.CallOpts, input)
}

// CheckRent is a free data retrieval call binding the contract method 0xd5dc648a.
//
// Solidity: function checkRent(uint8 y) constant returns(uint256)
func (_BASManager *BASManagerCaller) CheckRent(opts *bind.CallOpts, y uint8) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BASManager.contract.Call(opts, out, "checkRent", y)
	return *ret0, err
}

// CheckRent is a free data retrieval call binding the contract method 0xd5dc648a.
//
// Solidity: function checkRent(uint8 y) constant returns(uint256)
func (_BASManager *BASManagerSession) CheckRent(y uint8) (*big.Int, error) {
	return _BASManager.Contract.CheckRent(&_BASManager.CallOpts, y)
}

// CheckRent is a free data retrieval call binding the contract method 0xd5dc648a.
//
// Solidity: function checkRent(uint8 y) constant returns(uint256)
func (_BASManager *BASManagerCallerSession) CheckRent(y uint8) (*big.Int, error) {
	return _BASManager.Contract.CheckRent(&_BASManager.CallOpts, y)
}

// Hash is a free data retrieval call binding the contract method 0xb411ee94.
//
// Solidity: function hash(string key) constant returns(bytes32)
func (_BASManager *BASManagerCaller) Hash(opts *bind.CallOpts, key string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BASManager.contract.Call(opts, out, "hash", key)
	return *ret0, err
}

// Hash is a free data retrieval call binding the contract method 0xb411ee94.
//
// Solidity: function hash(string key) constant returns(bytes32)
func (_BASManager *BASManagerSession) Hash(key string) ([32]byte, error) {
	return _BASManager.Contract.Hash(&_BASManager.CallOpts, key)
}

// Hash is a free data retrieval call binding the contract method 0xb411ee94.
//
// Solidity: function hash(string key) constant returns(bytes32)
func (_BASManager *BASManagerCallerSession) Hash(key string) ([32]byte, error) {
	return _BASManager.Contract.Hash(&_BASManager.CallOpts, key)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BASManager *BASManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BASManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BASManager *BASManagerSession) Owner() (common.Address, error) {
	return _BASManager.Contract.Owner(&_BASManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BASManager *BASManagerCallerSession) Owner() (common.Address, error) {
	return _BASManager.Contract.Owner(&_BASManager.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_BASManager *BASManagerCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BASManager.contract.Call(opts, out, "price")
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_BASManager *BASManagerSession) Price() (*big.Int, error) {
	return _BASManager.Contract.Price(&_BASManager.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_BASManager *BASManagerCallerSession) Price() (*big.Int, error) {
	return _BASManager.Contract.Price(&_BASManager.CallOpts)
}

// QueryByBCAddress is a free data retrieval call binding the contract method 0x6e5c0a4b.
//
// Solidity: function queryByBCAddress(bytes32 bc) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32)
func (_BASManager *BASManagerCaller) QueryByBCAddress(opts *bind.CallOpts, bc [32]byte) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([4]byte)
		ret2 = new([16]byte)
		ret3 = new([1]byte)
		ret4 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _BASManager.contract.Call(opts, out, "queryByBCAddress", bc)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// QueryByBCAddress is a free data retrieval call binding the contract method 0x6e5c0a4b.
//
// Solidity: function queryByBCAddress(bytes32 bc) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32)
func (_BASManager *BASManagerSession) QueryByBCAddress(bc [32]byte) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, error) {
	return _BASManager.Contract.QueryByBCAddress(&_BASManager.CallOpts, bc)
}

// QueryByBCAddress is a free data retrieval call binding the contract method 0x6e5c0a4b.
//
// Solidity: function queryByBCAddress(bytes32 bc) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32)
func (_BASManager *BASManagerCallerSession) QueryByBCAddress(bc [32]byte) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, error) {
	return _BASManager.Contract.QueryByBCAddress(&_BASManager.CallOpts, bc)
}

// QueryByHash is a free data retrieval call binding the contract method 0xa2cb7c95.
//
// Solidity: function queryByHash(bytes32 hashKey) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32)
func (_BASManager *BASManagerCaller) QueryByHash(opts *bind.CallOpts, hashKey [32]byte) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([4]byte)
		ret2 = new([16]byte)
		ret3 = new([1]byte)
		ret4 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _BASManager.contract.Call(opts, out, "queryByHash", hashKey)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// QueryByHash is a free data retrieval call binding the contract method 0xa2cb7c95.
//
// Solidity: function queryByHash(bytes32 hashKey) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32)
func (_BASManager *BASManagerSession) QueryByHash(hashKey [32]byte) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, error) {
	return _BASManager.Contract.QueryByHash(&_BASManager.CallOpts, hashKey)
}

// QueryByHash is a free data retrieval call binding the contract method 0xa2cb7c95.
//
// Solidity: function queryByHash(bytes32 hashKey) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32)
func (_BASManager *BASManagerCallerSession) QueryByHash(hashKey [32]byte) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, error) {
	return _BASManager.Contract.QueryByHash(&_BASManager.CallOpts, hashKey)
}

// QueryByIPv4 is a free data retrieval call binding the contract method 0xce246171.
//
// Solidity: function queryByIPv4(bytes4 IPv4) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32)
func (_BASManager *BASManagerCaller) QueryByIPv4(opts *bind.CallOpts, IPv4 [4]byte) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([4]byte)
		ret2 = new([16]byte)
		ret3 = new([1]byte)
		ret4 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _BASManager.contract.Call(opts, out, "queryByIPv4", IPv4)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// QueryByIPv4 is a free data retrieval call binding the contract method 0xce246171.
//
// Solidity: function queryByIPv4(bytes4 IPv4) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32)
func (_BASManager *BASManagerSession) QueryByIPv4(IPv4 [4]byte) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, error) {
	return _BASManager.Contract.QueryByIPv4(&_BASManager.CallOpts, IPv4)
}

// QueryByIPv4 is a free data retrieval call binding the contract method 0xce246171.
//
// Solidity: function queryByIPv4(bytes4 IPv4) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32)
func (_BASManager *BASManagerCallerSession) QueryByIPv4(IPv4 [4]byte) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, error) {
	return _BASManager.Contract.QueryByIPv4(&_BASManager.CallOpts, IPv4)
}

// QueryByIPv6 is a free data retrieval call binding the contract method 0x68228e4b.
//
// Solidity: function queryByIPv6(bytes16 IPv6) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32)
func (_BASManager *BASManagerCaller) QueryByIPv6(opts *bind.CallOpts, IPv6 [16]byte) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([4]byte)
		ret2 = new([16]byte)
		ret3 = new([1]byte)
		ret4 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _BASManager.contract.Call(opts, out, "queryByIPv6", IPv6)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// QueryByIPv6 is a free data retrieval call binding the contract method 0x68228e4b.
//
// Solidity: function queryByIPv6(bytes16 IPv6) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32)
func (_BASManager *BASManagerSession) QueryByIPv6(IPv6 [16]byte) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, error) {
	return _BASManager.Contract.QueryByIPv6(&_BASManager.CallOpts, IPv6)
}

// QueryByIPv6 is a free data retrieval call binding the contract method 0x68228e4b.
//
// Solidity: function queryByIPv6(bytes16 IPv6) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32)
func (_BASManager *BASManagerCallerSession) QueryByIPv6(IPv6 [16]byte) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, error) {
	return _BASManager.Contract.QueryByIPv6(&_BASManager.CallOpts, IPv6)
}

// QueryByString is a free data retrieval call binding the contract method 0xab5e6aa7.
//
// Solidity: function queryByString(string key) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32)
func (_BASManager *BASManagerCaller) QueryByString(opts *bind.CallOpts, key string) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([4]byte)
		ret2 = new([16]byte)
		ret3 = new([1]byte)
		ret4 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _BASManager.contract.Call(opts, out, "queryByString", key)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// QueryByString is a free data retrieval call binding the contract method 0xab5e6aa7.
//
// Solidity: function queryByString(string key) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32)
func (_BASManager *BASManagerSession) QueryByString(key string) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, error) {
	return _BASManager.Contract.QueryByString(&_BASManager.CallOpts, key)
}

// QueryByString is a free data retrieval call binding the contract method 0xab5e6aa7.
//
// Solidity: function queryByString(string key) constant returns(bytes32, bytes4, bytes16, bytes1, bytes32)
func (_BASManager *BASManagerCallerSession) QueryByString(key string) ([32]byte, [4]byte, [16]byte, [1]byte, [32]byte, error) {
	return _BASManager.Contract.QueryByString(&_BASManager.CallOpts, key)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BASManager *BASManagerCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BASManager.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BASManager *BASManagerSession) Token() (common.Address, error) {
	return _BASManager.Contract.Token(&_BASManager.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BASManager *BASManagerCallerSession) Token() (common.Address, error) {
	return _BASManager.Contract.Token(&_BASManager.CallOpts)
}

// Change is a paid mutator transaction binding the contract method 0xda6459af.
//
// Solidity: function change(string key, bytes data) returns()
func (_BASManager *BASManagerTransactor) Change(opts *bind.TransactOpts, key string, data []byte) (*types.Transaction, error) {
	return _BASManager.contract.Transact(opts, "change", key, data)
}

// Change is a paid mutator transaction binding the contract method 0xda6459af.
//
// Solidity: function change(string key, bytes data) returns()
func (_BASManager *BASManagerSession) Change(key string, data []byte) (*types.Transaction, error) {
	return _BASManager.Contract.Change(&_BASManager.TransactOpts, key, data)
}

// Change is a paid mutator transaction binding the contract method 0xda6459af.
//
// Solidity: function change(string key, bytes data) returns()
func (_BASManager *BASManagerTransactorSession) Change(key string, data []byte) (*types.Transaction, error) {
	return _BASManager.Contract.Change(&_BASManager.TransactOpts, key, data)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x7085b50a.
//
// Solidity: function changePrice(uint16 newPrice) returns(uint256)
func (_BASManager *BASManagerTransactor) ChangePrice(opts *bind.TransactOpts, newPrice uint16) (*types.Transaction, error) {
	return _BASManager.contract.Transact(opts, "changePrice", newPrice)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x7085b50a.
//
// Solidity: function changePrice(uint16 newPrice) returns(uint256)
func (_BASManager *BASManagerSession) ChangePrice(newPrice uint16) (*types.Transaction, error) {
	return _BASManager.Contract.ChangePrice(&_BASManager.TransactOpts, newPrice)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x7085b50a.
//
// Solidity: function changePrice(uint16 newPrice) returns(uint256)
func (_BASManager *BASManagerTransactorSession) ChangePrice(newPrice uint16) (*types.Transaction, error) {
	return _BASManager.Contract.ChangePrice(&_BASManager.TransactOpts, newPrice)
}

// Rent is a paid mutator transaction binding the contract method 0x0d944db7.
//
// Solidity: function rent(string key, uint8 y, bytes data) returns()
func (_BASManager *BASManagerTransactor) Rent(opts *bind.TransactOpts, key string, y uint8, data []byte) (*types.Transaction, error) {
	return _BASManager.contract.Transact(opts, "rent", key, y, data)
}

// Rent is a paid mutator transaction binding the contract method 0x0d944db7.
//
// Solidity: function rent(string key, uint8 y, bytes data) returns()
func (_BASManager *BASManagerSession) Rent(key string, y uint8, data []byte) (*types.Transaction, error) {
	return _BASManager.Contract.Rent(&_BASManager.TransactOpts, key, y, data)
}

// Rent is a paid mutator transaction binding the contract method 0x0d944db7.
//
// Solidity: function rent(string key, uint8 y, bytes data) returns()
func (_BASManager *BASManagerTransactorSession) Rent(key string, y uint8, data []byte) (*types.Transaction, error) {
	return _BASManager.Contract.Rent(&_BASManager.TransactOpts, key, y, data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BASManager *BASManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BASManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BASManager *BASManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BASManager.Contract.TransferOwnership(&_BASManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BASManager *BASManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BASManager.Contract.TransferOwnership(&_BASManager.TransactOpts, newOwner)
}
