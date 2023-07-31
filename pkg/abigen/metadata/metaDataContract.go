// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package metadata

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

// MetadataMetaData contains all meta data concerning the Metadata contract.
var MetadataMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"divs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ipfsPaths\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipfsPath\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"ticketPool\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"div\",\"type\":\"uint256\"}],\"name\":\"setMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ticketPools\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenNames\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610b68806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806367f718a91161005b57806367f718a9146101125780636d54547814610130578063cb4368ca14610160578063d58778d61461017c5761007d565b806310bf07171461008257806317499265146100b2578063310495ab146100e2575b600080fd5b61009c60048036038101906100979190610483565b6101ac565b6040516100a99190610540565b60405180910390f35b6100cc60048036038101906100c79190610483565b61024c565b6040516100d99190610571565b60405180910390f35b6100fc60048036038101906100f79190610483565b610264565b6040516101099190610540565b60405180910390f35b61011a610304565b604051610127919061064a565b60405180910390f35b61014a60048036038101906101459190610483565b61035c565b6040516101579190610571565b60405180910390f35b61017a600480360381019061017591906107a1565b610374565b005b61019660048036038101906101919190610483565b610415565b6040516101a39190610571565b60405180910390f35b600160205280600052604060002060009150905080546101cb90610883565b80601f01602080910402602001604051908101604052809291908181526020018280546101f790610883565b80156102445780601f1061021957610100808354040283529160200191610244565b820191906000526020600020905b81548152906001019060200180831161022757829003601f168201915b505050505081565b60036020528060005260406000206000915090505481565b6000602052806000526040600020600091509050805461028390610883565b80601f01602080910402602001604051908101604052809291908181526020018280546102af90610883565b80156102fc5780601f106102d1576101008083540402835291602001916102fc565b820191906000526020600020905b8154815290600101906020018083116102df57829003601f168201915b505050505081565b6060600480548060200260200160405190810160405280929190818152602001828054801561035257602002820191906000526020600020905b81548152602001906001019080831161033e575b5050505050905090565b60026020528060005260406000206000915090505481565b8360008087815260200190815260200160002090816103939190610a60565b50826001600087815260200190815260200160002090816103b49190610a60565b5081600260008781526020019081526020016000208190555080600360008781526020019081526020016000208190555060048590806001815401808255809150506001900390600052602060002001600090919091909150555050505050565b6004818154811061042557600080fd5b906000526020600020016000915090505481565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b6104608161044d565b811461046b57600080fd5b50565b60008135905061047d81610457565b92915050565b60006020828403121561049957610498610443565b5b60006104a78482850161046e565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156104ea5780820151818401526020810190506104cf565b60008484015250505050565b6000601f19601f8301169050919050565b6000610512826104b0565b61051c81856104bb565b935061052c8185602086016104cc565b610535816104f6565b840191505092915050565b6000602082019050818103600083015261055a8184610507565b905092915050565b61056b8161044d565b82525050565b60006020820190506105866000830184610562565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6105c18161044d565b82525050565b60006105d383836105b8565b60208301905092915050565b6000602082019050919050565b60006105f78261058c565b6106018185610597565b935061060c836105a8565b8060005b8381101561063d57815161062488826105c7565b975061062f836105df565b925050600181019050610610565b5085935050505092915050565b6000602082019050818103600083015261066481846105ec565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6106ae826104f6565b810181811067ffffffffffffffff821117156106cd576106cc610676565b5b80604052505050565b60006106e0610439565b90506106ec82826106a5565b919050565b600067ffffffffffffffff82111561070c5761070b610676565b5b610715826104f6565b9050602081019050919050565b82818337600083830152505050565b600061074461073f846106f1565b6106d6565b9050828152602081018484840111156107605761075f610671565b5b61076b848285610722565b509392505050565b600082601f8301126107885761078761066c565b5b8135610798848260208601610731565b91505092915050565b600080600080600060a086880312156107bd576107bc610443565b5b60006107cb8882890161046e565b955050602086013567ffffffffffffffff8111156107ec576107eb610448565b5b6107f888828901610773565b945050604086013567ffffffffffffffff81111561081957610818610448565b5b61082588828901610773565b93505060606108368882890161046e565b92505060806108478882890161046e565b9150509295509295909350565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061089b57607f821691505b6020821081036108ae576108ad610854565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026109167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826108d9565b61092086836108d9565b95508019841693508086168417925050509392505050565b6000819050919050565b600061095d6109586109538461044d565b610938565b61044d565b9050919050565b6000819050919050565b61097783610942565b61098b61098382610964565b8484546108e6565b825550505050565b600090565b6109a0610993565b6109ab81848461096e565b505050565b5b818110156109cf576109c4600082610998565b6001810190506109b1565b5050565b601f821115610a14576109e5816108b4565b6109ee846108c9565b810160208510156109fd578190505b610a11610a09856108c9565b8301826109b0565b50505b505050565b600082821c905092915050565b6000610a3760001984600802610a19565b1980831691505092915050565b6000610a508383610a26565b9150826002028217905092915050565b610a69826104b0565b67ffffffffffffffff811115610a8257610a81610676565b5b610a8c8254610883565b610a978282856109d3565b600060209050601f831160018114610aca5760008415610ab8578287015190505b610ac28582610a44565b865550610b2a565b601f198416610ad8866108b4565b60005b82811015610b0057848901518255600182019150602085019450602081019050610adb565b86831015610b1d5784890151610b19601f891682610a26565b8355505b6001600288020188555050505b50505050505056fea26469706673582212207fe39d79849e817471d118b2e924f46dd0fe3390899e909a4c2b534191779b3064736f6c63430008120033",
}

// MetadataABI is the input ABI used to generate the binding from.
// Deprecated: Use MetadataMetaData.ABI instead.
var MetadataABI = MetadataMetaData.ABI

// MetadataBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MetadataMetaData.Bin instead.
var MetadataBin = MetadataMetaData.Bin

// DeployMetadata deploys a new Ethereum contract, binding an instance of Metadata to it.
func DeployMetadata(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Metadata, error) {
	parsed, err := MetadataMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MetadataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Metadata{MetadataCaller: MetadataCaller{contract: contract}, MetadataTransactor: MetadataTransactor{contract: contract}, MetadataFilterer: MetadataFilterer{contract: contract}}, nil
}

// Metadata is an auto generated Go binding around an Ethereum contract.
type Metadata struct {
	MetadataCaller     // Read-only binding to the contract
	MetadataTransactor // Write-only binding to the contract
	MetadataFilterer   // Log filterer for contract events
}

// MetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetadataSession struct {
	Contract     *Metadata         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetadataCallerSession struct {
	Contract *MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetadataTransactorSession struct {
	Contract     *MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetadataRaw struct {
	Contract *Metadata // Generic contract binding to access the raw methods on
}

// MetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetadataCallerRaw struct {
	Contract *MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// MetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetadataTransactorRaw struct {
	Contract *MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetadata creates a new instance of Metadata, bound to a specific deployed contract.
func NewMetadata(address common.Address, backend bind.ContractBackend) (*Metadata, error) {
	contract, err := bindMetadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Metadata{MetadataCaller: MetadataCaller{contract: contract}, MetadataTransactor: MetadataTransactor{contract: contract}, MetadataFilterer: MetadataFilterer{contract: contract}}, nil
}

// NewMetadataCaller creates a new read-only instance of Metadata, bound to a specific deployed contract.
func NewMetadataCaller(address common.Address, caller bind.ContractCaller) (*MetadataCaller, error) {
	contract, err := bindMetadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetadataCaller{contract: contract}, nil
}

// NewMetadataTransactor creates a new write-only instance of Metadata, bound to a specific deployed contract.
func NewMetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*MetadataTransactor, error) {
	contract, err := bindMetadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetadataTransactor{contract: contract}, nil
}

// NewMetadataFilterer creates a new log filterer instance of Metadata, bound to a specific deployed contract.
func NewMetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*MetadataFilterer, error) {
	contract, err := bindMetadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetadataFilterer{contract: contract}, nil
}

// bindMetadata binds a generic wrapper to an already deployed contract.
func bindMetadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MetadataMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Metadata *MetadataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Metadata.Contract.MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Metadata *MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Metadata.Contract.MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Metadata *MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Metadata.Contract.MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Metadata *MetadataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Metadata *MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Metadata *MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Metadata.Contract.contract.Transact(opts, method, params...)
}

// Divs is a free data retrieval call binding the contract method 0x17499265.
//
// Solidity: function divs(uint256 ) view returns(uint256)
func (_Metadata *MetadataCaller) Divs(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Metadata.contract.Call(opts, &out, "divs", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Divs is a free data retrieval call binding the contract method 0x17499265.
//
// Solidity: function divs(uint256 ) view returns(uint256)
func (_Metadata *MetadataSession) Divs(arg0 *big.Int) (*big.Int, error) {
	return _Metadata.Contract.Divs(&_Metadata.CallOpts, arg0)
}

// Divs is a free data retrieval call binding the contract method 0x17499265.
//
// Solidity: function divs(uint256 ) view returns(uint256)
func (_Metadata *MetadataCallerSession) Divs(arg0 *big.Int) (*big.Int, error) {
	return _Metadata.Contract.Divs(&_Metadata.CallOpts, arg0)
}

// GetTokenIds is a free data retrieval call binding the contract method 0x67f718a9.
//
// Solidity: function getTokenIds() view returns(uint256[])
func (_Metadata *MetadataCaller) GetTokenIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Metadata.contract.Call(opts, &out, "getTokenIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTokenIds is a free data retrieval call binding the contract method 0x67f718a9.
//
// Solidity: function getTokenIds() view returns(uint256[])
func (_Metadata *MetadataSession) GetTokenIds() ([]*big.Int, error) {
	return _Metadata.Contract.GetTokenIds(&_Metadata.CallOpts)
}

// GetTokenIds is a free data retrieval call binding the contract method 0x67f718a9.
//
// Solidity: function getTokenIds() view returns(uint256[])
func (_Metadata *MetadataCallerSession) GetTokenIds() ([]*big.Int, error) {
	return _Metadata.Contract.GetTokenIds(&_Metadata.CallOpts)
}

// IpfsPaths is a free data retrieval call binding the contract method 0x10bf0717.
//
// Solidity: function ipfsPaths(uint256 ) view returns(string)
func (_Metadata *MetadataCaller) IpfsPaths(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Metadata.contract.Call(opts, &out, "ipfsPaths", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// IpfsPaths is a free data retrieval call binding the contract method 0x10bf0717.
//
// Solidity: function ipfsPaths(uint256 ) view returns(string)
func (_Metadata *MetadataSession) IpfsPaths(arg0 *big.Int) (string, error) {
	return _Metadata.Contract.IpfsPaths(&_Metadata.CallOpts, arg0)
}

// IpfsPaths is a free data retrieval call binding the contract method 0x10bf0717.
//
// Solidity: function ipfsPaths(uint256 ) view returns(string)
func (_Metadata *MetadataCallerSession) IpfsPaths(arg0 *big.Int) (string, error) {
	return _Metadata.Contract.IpfsPaths(&_Metadata.CallOpts, arg0)
}

// TicketPools is a free data retrieval call binding the contract method 0x6d545478.
//
// Solidity: function ticketPools(uint256 ) view returns(uint256)
func (_Metadata *MetadataCaller) TicketPools(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Metadata.contract.Call(opts, &out, "ticketPools", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TicketPools is a free data retrieval call binding the contract method 0x6d545478.
//
// Solidity: function ticketPools(uint256 ) view returns(uint256)
func (_Metadata *MetadataSession) TicketPools(arg0 *big.Int) (*big.Int, error) {
	return _Metadata.Contract.TicketPools(&_Metadata.CallOpts, arg0)
}

// TicketPools is a free data retrieval call binding the contract method 0x6d545478.
//
// Solidity: function ticketPools(uint256 ) view returns(uint256)
func (_Metadata *MetadataCallerSession) TicketPools(arg0 *big.Int) (*big.Int, error) {
	return _Metadata.Contract.TicketPools(&_Metadata.CallOpts, arg0)
}

// TokenIds is a free data retrieval call binding the contract method 0xd58778d6.
//
// Solidity: function tokenIds(uint256 ) view returns(uint256)
func (_Metadata *MetadataCaller) TokenIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Metadata.contract.Call(opts, &out, "tokenIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIds is a free data retrieval call binding the contract method 0xd58778d6.
//
// Solidity: function tokenIds(uint256 ) view returns(uint256)
func (_Metadata *MetadataSession) TokenIds(arg0 *big.Int) (*big.Int, error) {
	return _Metadata.Contract.TokenIds(&_Metadata.CallOpts, arg0)
}

// TokenIds is a free data retrieval call binding the contract method 0xd58778d6.
//
// Solidity: function tokenIds(uint256 ) view returns(uint256)
func (_Metadata *MetadataCallerSession) TokenIds(arg0 *big.Int) (*big.Int, error) {
	return _Metadata.Contract.TokenIds(&_Metadata.CallOpts, arg0)
}

// TokenNames is a free data retrieval call binding the contract method 0x310495ab.
//
// Solidity: function tokenNames(uint256 ) view returns(string)
func (_Metadata *MetadataCaller) TokenNames(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Metadata.contract.Call(opts, &out, "tokenNames", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenNames is a free data retrieval call binding the contract method 0x310495ab.
//
// Solidity: function tokenNames(uint256 ) view returns(string)
func (_Metadata *MetadataSession) TokenNames(arg0 *big.Int) (string, error) {
	return _Metadata.Contract.TokenNames(&_Metadata.CallOpts, arg0)
}

// TokenNames is a free data retrieval call binding the contract method 0x310495ab.
//
// Solidity: function tokenNames(uint256 ) view returns(string)
func (_Metadata *MetadataCallerSession) TokenNames(arg0 *big.Int) (string, error) {
	return _Metadata.Contract.TokenNames(&_Metadata.CallOpts, arg0)
}

// SetMetadata is a paid mutator transaction binding the contract method 0xcb4368ca.
//
// Solidity: function setMetadata(uint256 tokenId, string tokenName, string ipfsPath, uint256 ticketPool, uint256 div) returns()
func (_Metadata *MetadataTransactor) SetMetadata(opts *bind.TransactOpts, tokenId *big.Int, tokenName string, ipfsPath string, ticketPool *big.Int, div *big.Int) (*types.Transaction, error) {
	return _Metadata.contract.Transact(opts, "setMetadata", tokenId, tokenName, ipfsPath, ticketPool, div)
}

// SetMetadata is a paid mutator transaction binding the contract method 0xcb4368ca.
//
// Solidity: function setMetadata(uint256 tokenId, string tokenName, string ipfsPath, uint256 ticketPool, uint256 div) returns()
func (_Metadata *MetadataSession) SetMetadata(tokenId *big.Int, tokenName string, ipfsPath string, ticketPool *big.Int, div *big.Int) (*types.Transaction, error) {
	return _Metadata.Contract.SetMetadata(&_Metadata.TransactOpts, tokenId, tokenName, ipfsPath, ticketPool, div)
}

// SetMetadata is a paid mutator transaction binding the contract method 0xcb4368ca.
//
// Solidity: function setMetadata(uint256 tokenId, string tokenName, string ipfsPath, uint256 ticketPool, uint256 div) returns()
func (_Metadata *MetadataTransactorSession) SetMetadata(tokenId *big.Int, tokenName string, ipfsPath string, ticketPool *big.Int, div *big.Int) (*types.Transaction, error) {
	return _Metadata.Contract.SetMetadata(&_Metadata.TransactOpts, tokenId, tokenName, ipfsPath, ticketPool, div)
}
