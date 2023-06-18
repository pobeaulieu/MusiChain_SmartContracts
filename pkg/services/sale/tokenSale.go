// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sale

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

// TokenSaleMetaData contains all meta data concerning the TokenSale contract.
var TokenSaleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Sell\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numberOfTokens\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"buyTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensSold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5060405161060638038061060683398101604081905261002e91610064565b60018054336001600160a01b0319918216179091555f80549091166001600160a01b03939093169290921790915560025561009b565b5f8060408385031215610075575f80fd5b82516001600160a01b038116811461008b575f80fd5b6020939093015192949293505050565b61055e806100a85f395ff3fe608060405260043610610054575f3560e01c8063380d831b146100585780633d00b8e31461006e578063518ab2a81461008157806355a373d6146100a95780637ff9b596146100df578063f851a440146100f4575b5f80fd5b348015610063575f80fd5b5061006c610113565b005b61006c61007c36600461039b565b6101a3565b34801561008c575f80fd5b5061009660035481565b6040519081526020015b60405180910390f35b3480156100b4575f80fd5b505f546100c7906001600160a01b031681565b6040516001600160a01b0390911681526020016100a0565b3480156100ea575f80fd5b5061009660025481565b3480156100ff575f80fd5b506001546100c7906001600160a01b031681565b6001546001600160a01b0316331461016a5760405162461bcd60e51b81526020600482015260156024820152742cb7ba9030b932903737ba103a34329030b236b4b760591b60448201526064015b60405180910390fd5b6001546040516001600160a01b03909116904780156108fc02915f818181858888f193505050501580156101a0573d5f803e3d5ffd5b50565b6002546101b0908361046c565b34146102135760405162461bcd60e51b815260206004820152602c60248201527f596f75206e65656420746f2073656e642074686520636f727265637420616d6f60448201526b3ab73a1037b31022ba3432b960a11b6064820152608401610161565b5f54604051627eeac760e11b81523060048201526024810185905283916001600160a01b03169062fdd58e90604401602060405180830381865afa15801561025d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102819190610489565b10156102cf5760405162461bcd60e51b815260206004820152601b60248201527f4e6f7420656e6f75676820746f6b656e7320617661696c61626c6500000000006044820152606401610161565b5f54604051637921219560e11b81526001600160a01b039091169063f242432a9061030690309033908890889088906004016104a0565b5f604051808303815f87803b15801561031d575f80fd5b505af115801561032f573d5f803e3d5ffd5b505050508160035f8282546103449190610515565b909155505060408051338152602081018490527f5e5e995ce3133561afceaa51a9a154d5db228cd7525d34df5185582c18d3df09910160405180910390a1505050565b634e487b7160e01b5f52604160045260245ffd5b5f805f606084860312156103ad575f80fd5b8335925060208401359150604084013567ffffffffffffffff808211156103d2575f80fd5b818601915086601f8301126103e5575f80fd5b8135818111156103f7576103f7610387565b604051601f8201601f19908116603f0116810190838211818310171561041f5761041f610387565b81604052828152896020848701011115610437575f80fd5b826020860160208301375f6020848301015280955050505050509250925092565b634e487b7160e01b5f52601160045260245ffd5b808202811582820484141761048357610483610458565b92915050565b5f60208284031215610499575f80fd5b5051919050565b5f60018060a01b03808816835260208188168185015286604085015285606085015260a06080850152845191508160a08501525f5b828110156104f15785810182015185820160c0015281016104d5565b50505f60c0828501015260c0601f19601f8301168401019150509695505050505050565b808201808211156104835761048361045856fea26469706673582212207c5088de70f3f59932fafbf64f3aaf970ed00589ba42051d9602b8af64b4686a64736f6c63430008140033",
}

// TokenSaleABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenSaleMetaData.ABI instead.
var TokenSaleABI = TokenSaleMetaData.ABI

// TokenSaleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenSaleMetaData.Bin instead.
var TokenSaleBin = TokenSaleMetaData.Bin

// DeployTokenSale deploys a new Ethereum contract, binding an instance of TokenSale to it.
func DeployTokenSale(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenContract common.Address, _tokenPrice *big.Int) (common.Address, *types.Transaction, *TokenSale, error) {
	parsed, err := TokenSaleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenSaleBin), backend, _tokenContract, _tokenPrice)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenSale{TokenSaleCaller: TokenSaleCaller{contract: contract}, TokenSaleTransactor: TokenSaleTransactor{contract: contract}, TokenSaleFilterer: TokenSaleFilterer{contract: contract}}, nil
}

// TokenSale is an auto generated Go binding around an Ethereum contract.
type TokenSale struct {
	TokenSaleCaller     // Read-only binding to the contract
	TokenSaleTransactor // Write-only binding to the contract
	TokenSaleFilterer   // Log filterer for contract events
}

// TokenSaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenSaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenSaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenSaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSaleSession struct {
	Contract     *TokenSale        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenSaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenSaleCallerSession struct {
	Contract *TokenSaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TokenSaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenSaleTransactorSession struct {
	Contract     *TokenSaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TokenSaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenSaleRaw struct {
	Contract *TokenSale // Generic contract binding to access the raw methods on
}

// TokenSaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenSaleCallerRaw struct {
	Contract *TokenSaleCaller // Generic read-only contract binding to access the raw methods on
}

// TokenSaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenSaleTransactorRaw struct {
	Contract *TokenSaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenSale creates a new instance of TokenSale, bound to a specific deployed contract.
func NewTokenSale(address common.Address, backend bind.ContractBackend) (*TokenSale, error) {
	contract, err := bindTokenSale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenSale{TokenSaleCaller: TokenSaleCaller{contract: contract}, TokenSaleTransactor: TokenSaleTransactor{contract: contract}, TokenSaleFilterer: TokenSaleFilterer{contract: contract}}, nil
}

// NewTokenSaleCaller creates a new read-only instance of TokenSale, bound to a specific deployed contract.
func NewTokenSaleCaller(address common.Address, caller bind.ContractCaller) (*TokenSaleCaller, error) {
	contract, err := bindTokenSale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenSaleCaller{contract: contract}, nil
}

// NewTokenSaleTransactor creates a new write-only instance of TokenSale, bound to a specific deployed contract.
func NewTokenSaleTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenSaleTransactor, error) {
	contract, err := bindTokenSale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenSaleTransactor{contract: contract}, nil
}

// NewTokenSaleFilterer creates a new log filterer instance of TokenSale, bound to a specific deployed contract.
func NewTokenSaleFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenSaleFilterer, error) {
	contract, err := bindTokenSale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenSaleFilterer{contract: contract}, nil
}

// bindTokenSale binds a generic wrapper to an already deployed contract.
func bindTokenSale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenSaleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenSale *TokenSaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenSale.Contract.TokenSaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenSale *TokenSaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenSale.Contract.TokenSaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenSale *TokenSaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenSale.Contract.TokenSaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenSale *TokenSaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenSale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenSale *TokenSaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenSale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenSale *TokenSaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenSale.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TokenSale *TokenSaleCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenSale.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TokenSale *TokenSaleSession) Admin() (common.Address, error) {
	return _TokenSale.Contract.Admin(&_TokenSale.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TokenSale *TokenSaleCallerSession) Admin() (common.Address, error) {
	return _TokenSale.Contract.Admin(&_TokenSale.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() view returns(address)
func (_TokenSale *TokenSaleCaller) TokenContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenSale.contract.Call(opts, &out, "tokenContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() view returns(address)
func (_TokenSale *TokenSaleSession) TokenContract() (common.Address, error) {
	return _TokenSale.Contract.TokenContract(&_TokenSale.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() view returns(address)
func (_TokenSale *TokenSaleCallerSession) TokenContract() (common.Address, error) {
	return _TokenSale.Contract.TokenContract(&_TokenSale.CallOpts)
}

// TokenPrice is a free data retrieval call binding the contract method 0x7ff9b596.
//
// Solidity: function tokenPrice() view returns(uint256)
func (_TokenSale *TokenSaleCaller) TokenPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenSale.contract.Call(opts, &out, "tokenPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenPrice is a free data retrieval call binding the contract method 0x7ff9b596.
//
// Solidity: function tokenPrice() view returns(uint256)
func (_TokenSale *TokenSaleSession) TokenPrice() (*big.Int, error) {
	return _TokenSale.Contract.TokenPrice(&_TokenSale.CallOpts)
}

// TokenPrice is a free data retrieval call binding the contract method 0x7ff9b596.
//
// Solidity: function tokenPrice() view returns(uint256)
func (_TokenSale *TokenSaleCallerSession) TokenPrice() (*big.Int, error) {
	return _TokenSale.Contract.TokenPrice(&_TokenSale.CallOpts)
}

// TokensSold is a free data retrieval call binding the contract method 0x518ab2a8.
//
// Solidity: function tokensSold() view returns(uint256)
func (_TokenSale *TokenSaleCaller) TokensSold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenSale.contract.Call(opts, &out, "tokensSold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensSold is a free data retrieval call binding the contract method 0x518ab2a8.
//
// Solidity: function tokensSold() view returns(uint256)
func (_TokenSale *TokenSaleSession) TokensSold() (*big.Int, error) {
	return _TokenSale.Contract.TokensSold(&_TokenSale.CallOpts)
}

// TokensSold is a free data retrieval call binding the contract method 0x518ab2a8.
//
// Solidity: function tokensSold() view returns(uint256)
func (_TokenSale *TokenSaleCallerSession) TokensSold() (*big.Int, error) {
	return _TokenSale.Contract.TokensSold(&_TokenSale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0x3d00b8e3.
//
// Solidity: function buyTokens(uint256 _tokenId, uint256 _numberOfTokens, bytes _data) payable returns()
func (_TokenSale *TokenSaleTransactor) BuyTokens(opts *bind.TransactOpts, _tokenId *big.Int, _numberOfTokens *big.Int, _data []byte) (*types.Transaction, error) {
	return _TokenSale.contract.Transact(opts, "buyTokens", _tokenId, _numberOfTokens, _data)
}

// BuyTokens is a paid mutator transaction binding the contract method 0x3d00b8e3.
//
// Solidity: function buyTokens(uint256 _tokenId, uint256 _numberOfTokens, bytes _data) payable returns()
func (_TokenSale *TokenSaleSession) BuyTokens(_tokenId *big.Int, _numberOfTokens *big.Int, _data []byte) (*types.Transaction, error) {
	return _TokenSale.Contract.BuyTokens(&_TokenSale.TransactOpts, _tokenId, _numberOfTokens, _data)
}

// BuyTokens is a paid mutator transaction binding the contract method 0x3d00b8e3.
//
// Solidity: function buyTokens(uint256 _tokenId, uint256 _numberOfTokens, bytes _data) payable returns()
func (_TokenSale *TokenSaleTransactorSession) BuyTokens(_tokenId *big.Int, _numberOfTokens *big.Int, _data []byte) (*types.Transaction, error) {
	return _TokenSale.Contract.BuyTokens(&_TokenSale.TransactOpts, _tokenId, _numberOfTokens, _data)
}

// EndSale is a paid mutator transaction binding the contract method 0x380d831b.
//
// Solidity: function endSale() returns()
func (_TokenSale *TokenSaleTransactor) EndSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenSale.contract.Transact(opts, "endSale")
}

// EndSale is a paid mutator transaction binding the contract method 0x380d831b.
//
// Solidity: function endSale() returns()
func (_TokenSale *TokenSaleSession) EndSale() (*types.Transaction, error) {
	return _TokenSale.Contract.EndSale(&_TokenSale.TransactOpts)
}

// EndSale is a paid mutator transaction binding the contract method 0x380d831b.
//
// Solidity: function endSale() returns()
func (_TokenSale *TokenSaleTransactorSession) EndSale() (*types.Transaction, error) {
	return _TokenSale.Contract.EndSale(&_TokenSale.TransactOpts)
}

// TokenSaleSellIterator is returned from FilterSell and is used to iterate over the raw logs and unpacked data for Sell events raised by the TokenSale contract.
type TokenSaleSellIterator struct {
	Event *TokenSaleSell // Event containing the contract specifics and raw log

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
func (it *TokenSaleSellIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSaleSell)
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
		it.Event = new(TokenSaleSell)
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
func (it *TokenSaleSellIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSaleSellIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSaleSell represents a Sell event raised by the TokenSale contract.
type TokenSaleSell struct {
	Buyer  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSell is a free log retrieval operation binding the contract event 0x5e5e995ce3133561afceaa51a9a154d5db228cd7525d34df5185582c18d3df09.
//
// Solidity: event Sell(address _buyer, uint256 _amount)
func (_TokenSale *TokenSaleFilterer) FilterSell(opts *bind.FilterOpts) (*TokenSaleSellIterator, error) {

	logs, sub, err := _TokenSale.contract.FilterLogs(opts, "Sell")
	if err != nil {
		return nil, err
	}
	return &TokenSaleSellIterator{contract: _TokenSale.contract, event: "Sell", logs: logs, sub: sub}, nil
}

// WatchSell is a free log subscription operation binding the contract event 0x5e5e995ce3133561afceaa51a9a154d5db228cd7525d34df5185582c18d3df09.
//
// Solidity: event Sell(address _buyer, uint256 _amount)
func (_TokenSale *TokenSaleFilterer) WatchSell(opts *bind.WatchOpts, sink chan<- *TokenSaleSell) (event.Subscription, error) {

	logs, sub, err := _TokenSale.contract.WatchLogs(opts, "Sell")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSaleSell)
				if err := _TokenSale.contract.UnpackLog(event, "Sell", log); err != nil {
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

// ParseSell is a log parse operation binding the contract event 0x5e5e995ce3133561afceaa51a9a154d5db228cd7525d34df5185582c18d3df09.
//
// Solidity: event Sell(address _buyer, uint256 _amount)
func (_TokenSale *TokenSaleFilterer) ParseSell(log types.Log) (*TokenSaleSell, error) {
	event := new(TokenSaleSell)
	if err := _TokenSale.contract.UnpackLog(event, "Sell", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
