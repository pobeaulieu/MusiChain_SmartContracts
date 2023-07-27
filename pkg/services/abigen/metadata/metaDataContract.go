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
	ABI: "[{\"inputs\":[],\"name\":\"getTokenIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ipfsPaths\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipfsPath\",\"type\":\"string\"}],\"name\":\"setMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenNames\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610a58806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806310bf07171461005c578063310495ab1461008c57806359f9f50f146100bc57806367f718a9146100d8578063d58778d6146100f6575b600080fd5b6100766004803603810190610071919061039b565b610126565b6040516100839190610458565b60405180910390f35b6100a660048036038101906100a1919061039b565b6101c6565b6040516100b39190610458565b60405180910390f35b6100d660048036038101906100d191906105af565b610266565b005b6100e06102d5565b6040516100ed91906106f8565b60405180910390f35b610110600480360381019061010b919061039b565b61032d565b60405161011d9190610729565b60405180910390f35b6001602052806000526040600020600091509050805461014590610773565b80601f016020809104026020016040519081016040528092919081815260200182805461017190610773565b80156101be5780601f10610193576101008083540402835291602001916101be565b820191906000526020600020905b8154815290600101906020018083116101a157829003601f168201915b505050505081565b600060205280600052604060002060009150905080546101e590610773565b80601f016020809104026020016040519081016040528092919081815260200182805461021190610773565b801561025e5780601f106102335761010080835404028352916020019161025e565b820191906000526020600020905b81548152906001019060200180831161024157829003601f168201915b505050505081565b8160008085815260200190815260200160002090816102859190610950565b50806001600085815260200190815260200160002090816102a69190610950565b506002839080600181540180825580915050600190039060005260206000200160009091909190915055505050565b6060600280548060200260200160405190810160405280929190818152602001828054801561032357602002820191906000526020600020905b81548152602001906001019080831161030f575b5050505050905090565b6002818154811061033d57600080fd5b906000526020600020016000915090505481565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b61037881610365565b811461038357600080fd5b50565b6000813590506103958161036f565b92915050565b6000602082840312156103b1576103b061035b565b5b60006103bf84828501610386565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156104025780820151818401526020810190506103e7565b60008484015250505050565b6000601f19601f8301169050919050565b600061042a826103c8565b61043481856103d3565b93506104448185602086016103e4565b61044d8161040e565b840191505092915050565b60006020820190508181036000830152610472818461041f565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6104bc8261040e565b810181811067ffffffffffffffff821117156104db576104da610484565b5b80604052505050565b60006104ee610351565b90506104fa82826104b3565b919050565b600067ffffffffffffffff82111561051a57610519610484565b5b6105238261040e565b9050602081019050919050565b82818337600083830152505050565b600061055261054d846104ff565b6104e4565b90508281526020810184848401111561056e5761056d61047f565b5b610579848285610530565b509392505050565b600082601f8301126105965761059561047a565b5b81356105a684826020860161053f565b91505092915050565b6000806000606084860312156105c8576105c761035b565b5b60006105d686828701610386565b935050602084013567ffffffffffffffff8111156105f7576105f6610360565b5b61060386828701610581565b925050604084013567ffffffffffffffff81111561062457610623610360565b5b61063086828701610581565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61066f81610365565b82525050565b60006106818383610666565b60208301905092915050565b6000602082019050919050565b60006106a58261063a565b6106af8185610645565b93506106ba83610656565b8060005b838110156106eb5781516106d28882610675565b97506106dd8361068d565b9250506001810190506106be565b5085935050505092915050565b60006020820190508181036000830152610712818461069a565b905092915050565b61072381610365565b82525050565b600060208201905061073e600083018461071a565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061078b57607f821691505b60208210810361079e5761079d610744565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026108067fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826107c9565b61081086836107c9565b95508019841693508086168417925050509392505050565b6000819050919050565b600061084d61084861084384610365565b610828565b610365565b9050919050565b6000819050919050565b61086783610832565b61087b61087382610854565b8484546107d6565b825550505050565b600090565b610890610883565b61089b81848461085e565b505050565b5b818110156108bf576108b4600082610888565b6001810190506108a1565b5050565b601f821115610904576108d5816107a4565b6108de846107b9565b810160208510156108ed578190505b6109016108f9856107b9565b8301826108a0565b50505b505050565b600082821c905092915050565b600061092760001984600802610909565b1980831691505092915050565b60006109408383610916565b9150826002028217905092915050565b610959826103c8565b67ffffffffffffffff81111561097257610971610484565b5b61097c8254610773565b6109878282856108c3565b600060209050601f8311600181146109ba57600084156109a8578287015190505b6109b28582610934565b865550610a1a565b601f1984166109c8866107a4565b60005b828110156109f0578489015182556001820191506020850194506020810190506109cb565b86831015610a0d5784890151610a09601f891682610916565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220395b0567c74c5657eb69a21d76c328c625b20dfb13b817481fd8e2a084ba021964736f6c63430008120033",
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

// SetMetadata is a paid mutator transaction binding the contract method 0x59f9f50f.
//
// Solidity: function setMetadata(uint256 tokenId, string tokenName, string ipfsPath) returns()
func (_Metadata *MetadataTransactor) SetMetadata(opts *bind.TransactOpts, tokenId *big.Int, tokenName string, ipfsPath string) (*types.Transaction, error) {
	return _Metadata.contract.Transact(opts, "setMetadata", tokenId, tokenName, ipfsPath)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x59f9f50f.
//
// Solidity: function setMetadata(uint256 tokenId, string tokenName, string ipfsPath) returns()
func (_Metadata *MetadataSession) SetMetadata(tokenId *big.Int, tokenName string, ipfsPath string) (*types.Transaction, error) {
	return _Metadata.Contract.SetMetadata(&_Metadata.TransactOpts, tokenId, tokenName, ipfsPath)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x59f9f50f.
//
// Solidity: function setMetadata(uint256 tokenId, string tokenName, string ipfsPath) returns()
func (_Metadata *MetadataTransactorSession) SetMetadata(tokenId *big.Int, tokenName string, ipfsPath string) (*types.Transaction, error) {
	return _Metadata.Contract.SetMetadata(&_Metadata.TransactOpts, tokenId, tokenName, ipfsPath)
}