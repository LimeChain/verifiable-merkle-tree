// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package merklelimecontract

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

// MerklelimecontractABI is the input ABI used to generate the binding from.
const MerklelimecontractABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"limeRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"setRoot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"nodes\",\"type\":\"bytes32[]\"},{\"name\":\"leafIndex\",\"type\":\"uint256\"}],\"name\":\"verifyDataInState\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Merklelimecontract is an auto generated Go binding around an Ethereum contract.
type Merklelimecontract struct {
	MerklelimecontractCaller     // Read-only binding to the contract
	MerklelimecontractTransactor // Write-only binding to the contract
	MerklelimecontractFilterer   // Log filterer for contract events
}

// MerklelimecontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerklelimecontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerklelimecontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerklelimecontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerklelimecontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerklelimecontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerklelimecontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerklelimecontractSession struct {
	Contract     *Merklelimecontract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MerklelimecontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerklelimecontractCallerSession struct {
	Contract *MerklelimecontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MerklelimecontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerklelimecontractTransactorSession struct {
	Contract     *MerklelimecontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MerklelimecontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerklelimecontractRaw struct {
	Contract *Merklelimecontract // Generic contract binding to access the raw methods on
}

// MerklelimecontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerklelimecontractCallerRaw struct {
	Contract *MerklelimecontractCaller // Generic read-only contract binding to access the raw methods on
}

// MerklelimecontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerklelimecontractTransactorRaw struct {
	Contract *MerklelimecontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerklelimecontract creates a new instance of Merklelimecontract, bound to a specific deployed contract.
func NewMerklelimecontract(address common.Address, backend bind.ContractBackend) (*Merklelimecontract, error) {
	contract, err := bindMerklelimecontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Merklelimecontract{MerklelimecontractCaller: MerklelimecontractCaller{contract: contract}, MerklelimecontractTransactor: MerklelimecontractTransactor{contract: contract}, MerklelimecontractFilterer: MerklelimecontractFilterer{contract: contract}}, nil
}

// NewMerklelimecontractCaller creates a new read-only instance of Merklelimecontract, bound to a specific deployed contract.
func NewMerklelimecontractCaller(address common.Address, caller bind.ContractCaller) (*MerklelimecontractCaller, error) {
	contract, err := bindMerklelimecontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerklelimecontractCaller{contract: contract}, nil
}

// NewMerklelimecontractTransactor creates a new write-only instance of Merklelimecontract, bound to a specific deployed contract.
func NewMerklelimecontractTransactor(address common.Address, transactor bind.ContractTransactor) (*MerklelimecontractTransactor, error) {
	contract, err := bindMerklelimecontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerklelimecontractTransactor{contract: contract}, nil
}

// NewMerklelimecontractFilterer creates a new log filterer instance of Merklelimecontract, bound to a specific deployed contract.
func NewMerklelimecontractFilterer(address common.Address, filterer bind.ContractFilterer) (*MerklelimecontractFilterer, error) {
	contract, err := bindMerklelimecontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerklelimecontractFilterer{contract: contract}, nil
}

// bindMerklelimecontract binds a generic wrapper to an already deployed contract.
func bindMerklelimecontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerklelimecontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Merklelimecontract *MerklelimecontractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Merklelimecontract.Contract.MerklelimecontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Merklelimecontract *MerklelimecontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merklelimecontract.Contract.MerklelimecontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Merklelimecontract *MerklelimecontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Merklelimecontract.Contract.MerklelimecontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Merklelimecontract *MerklelimecontractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Merklelimecontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Merklelimecontract *MerklelimecontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merklelimecontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Merklelimecontract *MerklelimecontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Merklelimecontract.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Merklelimecontract *MerklelimecontractCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Merklelimecontract.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Merklelimecontract *MerklelimecontractSession) IsOwner() (bool, error) {
	return _Merklelimecontract.Contract.IsOwner(&_Merklelimecontract.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Merklelimecontract *MerklelimecontractCallerSession) IsOwner() (bool, error) {
	return _Merklelimecontract.Contract.IsOwner(&_Merklelimecontract.CallOpts)
}

// LimeRoot is a free data retrieval call binding the contract method 0xb3ca488f.
//
// Solidity: function limeRoot() constant returns(bytes32)
func (_Merklelimecontract *MerklelimecontractCaller) LimeRoot(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Merklelimecontract.contract.Call(opts, out, "limeRoot")
	return *ret0, err
}

// LimeRoot is a free data retrieval call binding the contract method 0xb3ca488f.
//
// Solidity: function limeRoot() constant returns(bytes32)
func (_Merklelimecontract *MerklelimecontractSession) LimeRoot() ([32]byte, error) {
	return _Merklelimecontract.Contract.LimeRoot(&_Merklelimecontract.CallOpts)
}

// LimeRoot is a free data retrieval call binding the contract method 0xb3ca488f.
//
// Solidity: function limeRoot() constant returns(bytes32)
func (_Merklelimecontract *MerklelimecontractCallerSession) LimeRoot() ([32]byte, error) {
	return _Merklelimecontract.Contract.LimeRoot(&_Merklelimecontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Merklelimecontract *MerklelimecontractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Merklelimecontract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Merklelimecontract *MerklelimecontractSession) Owner() (common.Address, error) {
	return _Merklelimecontract.Contract.Owner(&_Merklelimecontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Merklelimecontract *MerklelimecontractCallerSession) Owner() (common.Address, error) {
	return _Merklelimecontract.Contract.Owner(&_Merklelimecontract.CallOpts)
}

// VerifyDataInState is a free data retrieval call binding the contract method 0x305f755c.
//
// Solidity: function verifyDataInState(data bytes, nodes bytes32[], leafIndex uint256) constant returns(bool)
func (_Merklelimecontract *MerklelimecontractCaller) VerifyDataInState(opts *bind.CallOpts, data []byte, nodes [][32]byte, leafIndex *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Merklelimecontract.contract.Call(opts, out, "verifyDataInState", data, nodes, leafIndex)
	return *ret0, err
}

// VerifyDataInState is a free data retrieval call binding the contract method 0x305f755c.
//
// Solidity: function verifyDataInState(data bytes, nodes bytes32[], leafIndex uint256) constant returns(bool)
func (_Merklelimecontract *MerklelimecontractSession) VerifyDataInState(data []byte, nodes [][32]byte, leafIndex *big.Int) (bool, error) {
	return _Merklelimecontract.Contract.VerifyDataInState(&_Merklelimecontract.CallOpts, data, nodes, leafIndex)
}

// VerifyDataInState is a free data retrieval call binding the contract method 0x305f755c.
//
// Solidity: function verifyDataInState(data bytes, nodes bytes32[], leafIndex uint256) constant returns(bool)
func (_Merklelimecontract *MerklelimecontractCallerSession) VerifyDataInState(data []byte, nodes [][32]byte, leafIndex *big.Int) (bool, error) {
	return _Merklelimecontract.Contract.VerifyDataInState(&_Merklelimecontract.CallOpts, data, nodes, leafIndex)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Merklelimecontract *MerklelimecontractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merklelimecontract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Merklelimecontract *MerklelimecontractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Merklelimecontract.Contract.RenounceOwnership(&_Merklelimecontract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Merklelimecontract *MerklelimecontractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Merklelimecontract.Contract.RenounceOwnership(&_Merklelimecontract.TransactOpts)
}

// SetRoot is a paid mutator transaction binding the contract method 0xdab5f340.
//
// Solidity: function setRoot(merkleRoot bytes32) returns()
func (_Merklelimecontract *MerklelimecontractTransactor) SetRoot(opts *bind.TransactOpts, merkleRoot [32]byte) (*types.Transaction, error) {
	return _Merklelimecontract.contract.Transact(opts, "setRoot", merkleRoot)
}

// SetRoot is a paid mutator transaction binding the contract method 0xdab5f340.
//
// Solidity: function setRoot(merkleRoot bytes32) returns()
func (_Merklelimecontract *MerklelimecontractSession) SetRoot(merkleRoot [32]byte) (*types.Transaction, error) {
	return _Merklelimecontract.Contract.SetRoot(&_Merklelimecontract.TransactOpts, merkleRoot)
}

// SetRoot is a paid mutator transaction binding the contract method 0xdab5f340.
//
// Solidity: function setRoot(merkleRoot bytes32) returns()
func (_Merklelimecontract *MerklelimecontractTransactorSession) SetRoot(merkleRoot [32]byte) (*types.Transaction, error) {
	return _Merklelimecontract.Contract.SetRoot(&_Merklelimecontract.TransactOpts, merkleRoot)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Merklelimecontract *MerklelimecontractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Merklelimecontract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Merklelimecontract *MerklelimecontractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Merklelimecontract.Contract.TransferOwnership(&_Merklelimecontract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Merklelimecontract *MerklelimecontractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Merklelimecontract.Contract.TransferOwnership(&_Merklelimecontract.TransactOpts, newOwner)
}

// MerklelimecontractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Merklelimecontract contract.
type MerklelimecontractOwnershipTransferredIterator struct {
	Event *MerklelimecontractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MerklelimecontractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerklelimecontractOwnershipTransferred)
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
		it.Event = new(MerklelimecontractOwnershipTransferred)
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
func (it *MerklelimecontractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerklelimecontractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerklelimecontractOwnershipTransferred represents a OwnershipTransferred event raised by the Merklelimecontract contract.
type MerklelimecontractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Merklelimecontract *MerklelimecontractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MerklelimecontractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Merklelimecontract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MerklelimecontractOwnershipTransferredIterator{contract: _Merklelimecontract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Merklelimecontract *MerklelimecontractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MerklelimecontractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Merklelimecontract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerklelimecontractOwnershipTransferred)
				if err := _Merklelimecontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
