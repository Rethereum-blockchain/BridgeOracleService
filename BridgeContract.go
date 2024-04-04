package main

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

// BridgeContractMetaData contains all meta data concerning the BridgeContract contract.
var BridgeContractMetaData = &bind.MetaData{
    ABI: "[{\"type\":\"event\",\"name\":\"ActionCreated\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"id\",\"internalType\":\"bytes32\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"nonce\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"address\",\"name\":\"receiver\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ActionRequested\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"id\",\"internalType\":\"bytes32\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"actions\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"success\",\"internalType\":\"bool\"}],\"name\":\"authorizeAction\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"action\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"consumedActions\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}]}]",
}

// BridgeContract is an auto generated Go binding around an Ethereum contract.
type BridgeContract struct {
    BridgeContractCaller     // Read-only binding to the contract
    BridgeContractTransactor // Write-only binding to the contract
    BridgeContractFilterer   // Log filterer for contract events
}

// BridgeContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeContractCaller struct {
    contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeContractTransactor struct {
    contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeContractFilterer struct {
    contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeContractSession struct {
    Contract     *BridgeContract   // Generic contract binding to set the session for
    CallOpts     bind.CallOpts     // Call options to use throughout this session
    TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeContractCallerSession struct {
    Contract *BridgeContractCaller // Generic contract caller binding to set the session for
    CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BridgeContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeContractTransactorSession struct {
    Contract     *BridgeContractTransactor // Generic contract transactor binding to set the session for
    TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BridgeContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeContractRaw struct {
    Contract *BridgeContract // Generic contract binding to access the raw methods on
}

// BridgeContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeContractCallerRaw struct {
    Contract *BridgeContractCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeContractTransactorRaw struct {
    Contract *BridgeContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeContract creates a new instance of BridgeContract, bound to a specific deployed contract.
func NewBridgeContract(address common.Address, backend bind.ContractBackend) (*BridgeContract, error) {
    contract, err := bindBridgeContract(address, backend, backend, backend)
    if err != nil {
        return nil, err
    }
    return &BridgeContract{BridgeContractCaller: BridgeContractCaller{contract: contract}, BridgeContractTransactor: BridgeContractTransactor{contract: contract}, BridgeContractFilterer: BridgeContractFilterer{contract: contract}}, nil
}

// NewBridgeContractCaller creates a new read-only instance of BridgeContract, bound to a specific deployed contract.
func NewBridgeContractCaller(address common.Address, caller bind.ContractCaller) (*BridgeContractCaller, error) {
    contract, err := bindBridgeContract(address, caller, nil, nil)
    if err != nil {
        return nil, err
    }
    return &BridgeContractCaller{contract: contract}, nil
}

// NewBridgeContractTransactor creates a new write-only instance of BridgeContract, bound to a specific deployed contract.
func NewBridgeContractTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeContractTransactor, error) {
    contract, err := bindBridgeContract(address, nil, transactor, nil)
    if err != nil {
        return nil, err
    }
    return &BridgeContractTransactor{contract: contract}, nil
}

// NewBridgeContractFilterer creates a new log filterer instance of BridgeContract, bound to a specific deployed contract.
func NewBridgeContractFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeContractFilterer, error) {
    contract, err := bindBridgeContract(address, nil, nil, filterer)
    if err != nil {
        return nil, err
    }
    return &BridgeContractFilterer{contract: contract}, nil
}

// bindBridgeContract binds a generic wrapper to an already deployed contract.
func bindBridgeContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
    parsed, err := BridgeContractMetaData.GetAbi()
    if err != nil {
        return nil, err
    }
    return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeContract *BridgeContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
    return _BridgeContract.Contract.BridgeContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeContract *BridgeContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
    return _BridgeContract.Contract.BridgeContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeContract *BridgeContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
    return _BridgeContract.Contract.BridgeContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeContract *BridgeContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
    return _BridgeContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeContract *BridgeContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
    return _BridgeContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeContract *BridgeContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
    return _BridgeContract.Contract.contract.Transact(opts, method, params...)
}

// Actions is a free data retrieval call binding the contract method 0xf3abde32.
//
// Solidity: function actions(bytes32 ) view returns(bool)
func (_BridgeContract *BridgeContractCaller) Actions(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
    var out []interface{}
    err := _BridgeContract.contract.Call(opts, &out, "actions", arg0)

    if err != nil {
        return *new(bool), err
    }

    out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

    return out0, err

}

// Actions is a free data retrieval call binding the contract method 0xf3abde32.
//
// Solidity: function actions(bytes32 ) view returns(bool)
func (_BridgeContract *BridgeContractSession) Actions(arg0 [32]byte) (bool, error) {
    return _BridgeContract.Contract.Actions(&_BridgeContract.CallOpts, arg0)
}

// Actions is a free data retrieval call binding the contract method 0xf3abde32.
//
// Solidity: function actions(bytes32 ) view returns(bool)
func (_BridgeContract *BridgeContractCallerSession) Actions(arg0 [32]byte) (bool, error) {
    return _BridgeContract.Contract.Actions(&_BridgeContract.CallOpts, arg0)
}

// ConsumedActions is a free data retrieval call binding the contract method 0xe73a771e.
//
// Solidity: function consumedActions(bytes32 ) view returns(bool)
func (_BridgeContract *BridgeContractCaller) ConsumedActions(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
    var out []interface{}
    err := _BridgeContract.contract.Call(opts, &out, "consumedActions", arg0)

    if err != nil {
        return *new(bool), err
    }

    out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

    return out0, err

}

// ConsumedActions is a free data retrieval call binding the contract method 0xe73a771e.
//
// Solidity: function consumedActions(bytes32 ) view returns(bool)
func (_BridgeContract *BridgeContractSession) ConsumedActions(arg0 [32]byte) (bool, error) {
    return _BridgeContract.Contract.ConsumedActions(&_BridgeContract.CallOpts, arg0)
}

// ConsumedActions is a free data retrieval call binding the contract method 0xe73a771e.
//
// Solidity: function consumedActions(bytes32 ) view returns(bool)
func (_BridgeContract *BridgeContractCallerSession) ConsumedActions(arg0 [32]byte) (bool, error) {
    return _BridgeContract.Contract.ConsumedActions(&_BridgeContract.CallOpts, arg0)
}

// AuthorizeAction is a paid mutator transaction binding the contract method 0x7218688e.
//
// Solidity: function authorizeAction(bytes32 action) returns(bool success)
func (_BridgeContract *BridgeContractTransactor) AuthorizeAction(opts *bind.TransactOpts, action [32]byte) (*types.Transaction, error) {
    return _BridgeContract.contract.Transact(opts, "authorizeAction", action)
}

// AuthorizeAction is a paid mutator transaction binding the contract method 0x7218688e.
//
// Solidity: function authorizeAction(bytes32 action) returns(bool success)
func (_BridgeContract *BridgeContractSession) AuthorizeAction(action [32]byte) (*types.Transaction, error) {
    return _BridgeContract.Contract.AuthorizeAction(&_BridgeContract.TransactOpts, action)
}

// AuthorizeAction is a paid mutator transaction binding the contract method 0x7218688e.
//
// Solidity: function authorizeAction(bytes32 action) returns(bool success)
func (_BridgeContract *BridgeContractTransactorSession) AuthorizeAction(action [32]byte) (*types.Transaction, error) {
    return _BridgeContract.Contract.AuthorizeAction(&_BridgeContract.TransactOpts, action)
}

// BridgeContractActionCreatedIterator is returned from FilterActionCreated and is used to iterate over the raw logs and unpacked data for ActionCreated events raised by the BridgeContract contract.
type BridgeContractActionCreatedIterator struct {
    Event *BridgeContractActionCreated // Event containing the contract specifics and raw log

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
func (it *BridgeContractActionCreatedIterator) Next() bool {
    // If the iterator failed, stop iterating
    if it.fail != nil {
        return false
    }
    // If the iterator completed, deliver directly whatever's available
    if it.done {
        select {
        case log := <-it.logs:
            it.Event = new(BridgeContractActionCreated)
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
        it.Event = new(BridgeContractActionCreated)
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
func (it *BridgeContractActionCreatedIterator) Error() error {
    return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeContractActionCreatedIterator) Close() error {
    it.sub.Unsubscribe()
    return nil
}

// BridgeContractActionCreated represents a ActionCreated event raised by the BridgeContract contract.
type BridgeContractActionCreated struct {
    Id       [32]byte
    Nonce    *big.Int
    Receiver common.Address
    Amount   *big.Int
    Raw      types.Log // Blockchain specific contextual infos
}

// FilterActionCreated is a free log retrieval operation binding the contract event 0x8628d7e6a029415f73431bb1589718e9c57fd6c71668eb3ad7afd133d20200cb.
//
// Solidity: event ActionCreated(bytes32 indexed id, uint256 indexed nonce, address receiver, uint256 amount)
func (_BridgeContract *BridgeContractFilterer) FilterActionCreated(opts *bind.FilterOpts, id [][32]byte, nonce []*big.Int) (*BridgeContractActionCreatedIterator, error) {

    var idRule []interface{}
    for _, idItem := range id {
        idRule = append(idRule, idItem)
    }
    var nonceRule []interface{}
    for _, nonceItem := range nonce {
        nonceRule = append(nonceRule, nonceItem)
    }

    logs, sub, err := _BridgeContract.contract.FilterLogs(opts, "ActionCreated", idRule, nonceRule)
    if err != nil {
        return nil, err
    }
    return &BridgeContractActionCreatedIterator{contract: _BridgeContract.contract, event: "ActionCreated", logs: logs, sub: sub}, nil
}

// WatchActionCreated is a free log subscription operation binding the contract event 0x8628d7e6a029415f73431bb1589718e9c57fd6c71668eb3ad7afd133d20200cb.
//
// Solidity: event ActionCreated(bytes32 indexed id, uint256 indexed nonce, address receiver, uint256 amount)
func (_BridgeContract *BridgeContractFilterer) WatchActionCreated(opts *bind.WatchOpts, sink chan<- *BridgeContractActionCreated, id [][32]byte, nonce []*big.Int) (event.Subscription, error) {

    var idRule []interface{}
    for _, idItem := range id {
        idRule = append(idRule, idItem)
    }
    var nonceRule []interface{}
    for _, nonceItem := range nonce {
        nonceRule = append(nonceRule, nonceItem)
    }

    logs, sub, err := _BridgeContract.contract.WatchLogs(opts, "ActionCreated", idRule, nonceRule)
    if err != nil {
        return nil, err
    }
    return event.NewSubscription(func(quit <-chan struct{}) error {
        defer sub.Unsubscribe()
        for {
            select {
            case log := <-logs:
                // New log arrived, parse the event and forward to the user
                event := new(BridgeContractActionCreated)
                if err := _BridgeContract.contract.UnpackLog(event, "ActionCreated", log); err != nil {
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

// ParseActionCreated is a log parse operation binding the contract event 0x8628d7e6a029415f73431bb1589718e9c57fd6c71668eb3ad7afd133d20200cb.
//
// Solidity: event ActionCreated(bytes32 indexed id, uint256 indexed nonce, address receiver, uint256 amount)
func (_BridgeContract *BridgeContractFilterer) ParseActionCreated(log types.Log) (*BridgeContractActionCreated, error) {
    event := new(BridgeContractActionCreated)
    if err := _BridgeContract.contract.UnpackLog(event, "ActionCreated", log); err != nil {
        return nil, err
    }
    event.Raw = log
    return event, nil
}

// BridgeContractActionRequestedIterator is returned from FilterActionRequested and is used to iterate over the raw logs and unpacked data for ActionRequested events raised by the BridgeContract contract.
type BridgeContractActionRequestedIterator struct {
    Event *BridgeContractActionRequested // Event containing the contract specifics and raw log

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
func (it *BridgeContractActionRequestedIterator) Next() bool {
    // If the iterator failed, stop iterating
    if it.fail != nil {
        return false
    }
    // If the iterator completed, deliver directly whatever's available
    if it.done {
        select {
        case log := <-it.logs:
            it.Event = new(BridgeContractActionRequested)
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
        it.Event = new(BridgeContractActionRequested)
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
func (it *BridgeContractActionRequestedIterator) Error() error {
    return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeContractActionRequestedIterator) Close() error {
    it.sub.Unsubscribe()
    return nil
}

// BridgeContractActionRequested represents a ActionRequested event raised by the BridgeContract contract.
type BridgeContractActionRequested struct {
    Id  [32]byte
    Raw types.Log // Blockchain specific contextual infos
}

// FilterActionRequested is a free log retrieval operation binding the contract event 0x9400357dc65c84dba9cea1387bad2eaf611d23b4802e2a16c1f1d57d7ed7b6e5.
//
// Solidity: event ActionRequested(bytes32 indexed id)
func (_BridgeContract *BridgeContractFilterer) FilterActionRequested(opts *bind.FilterOpts, id [][32]byte) (*BridgeContractActionRequestedIterator, error) {

    var idRule []interface{}
    for _, idItem := range id {
        idRule = append(idRule, idItem)
    }

    logs, sub, err := _BridgeContract.contract.FilterLogs(opts, "ActionRequested", idRule)
    if err != nil {
        return nil, err
    }
    return &BridgeContractActionRequestedIterator{contract: _BridgeContract.contract, event: "ActionRequested", logs: logs, sub: sub}, nil
}

// WatchActionRequested is a free log subscription operation binding the contract event 0x9400357dc65c84dba9cea1387bad2eaf611d23b4802e2a16c1f1d57d7ed7b6e5.
//
// Solidity: event ActionRequested(bytes32 indexed id)
func (_BridgeContract *BridgeContractFilterer) WatchActionRequested(opts *bind.WatchOpts, sink chan<- *BridgeContractActionRequested, id [][32]byte) (event.Subscription, error) {

    var idRule []interface{}
    for _, idItem := range id {
        idRule = append(idRule, idItem)
    }

    logs, sub, err := _BridgeContract.contract.WatchLogs(opts, "ActionRequested", idRule)
    if err != nil {
        return nil, err
    }
    return event.NewSubscription(func(quit <-chan struct{}) error {
        defer sub.Unsubscribe()
        for {
            select {
            case log := <-logs:
                // New log arrived, parse the event and forward to the user
                event := new(BridgeContractActionRequested)
                if err := _BridgeContract.contract.UnpackLog(event, "ActionRequested", log); err != nil {
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

// ParseActionRequested is a log parse operation binding the contract event 0x9400357dc65c84dba9cea1387bad2eaf611d23b4802e2a16c1f1d57d7ed7b6e5.
//
// Solidity: event ActionRequested(bytes32 indexed id)
func (_BridgeContract *BridgeContractFilterer) ParseActionRequested(log types.Log) (*BridgeContractActionRequested, error) {
    event := new(BridgeContractActionRequested)
    if err := _BridgeContract.contract.UnpackLog(event, "ActionRequested", log); err != nil {
        return nil, err
    }
    event.Raw = log
    return event, nil
}
