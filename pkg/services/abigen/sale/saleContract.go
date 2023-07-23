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

// SaleListing is an auto generated low-level Go binding around an user-defined struct.
type SaleListing struct {
	Seller  common.Address
	TokenId *big.Int
	Price   *big.Int
	Amount  *big.Int
}

// SaleMetaData contains all meta data concerning the Sale contract.
var SaleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"buyToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getListings\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structSale.Listing[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getListingsByUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structSale.Listing[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"listToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001487380380620014878339818101604052810190620000379190620000fc565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200012e565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000b08262000083565b9050919050565b6000620000c482620000a3565b9050919050565b620000d681620000b7565b8114620000e257600080fd5b50565b600081519050620000f681620000cb565b92915050565b6000602082840312156200011557620001146200007e565b5b60006200012584828501620000e5565b91505092915050565b611349806200013e6000396000f3fe6080604052600436106100555760003560e01c80632d296bf11461005a57806355a373d61461007657806369df8327146100a1578063bce64a7d146100de578063de74e57b14610107578063f1b2d6a314610147575b600080fd5b610074600480360381019061006f9190610b3f565b610172565b005b34801561008257600080fd5b5061008b610625565b6040516100989190610beb565b60405180910390f35b3480156100ad57600080fd5b506100c860048036038101906100c39190610c44565b610649565b6040516100d59190610d93565b60405180910390f35b3480156100ea57600080fd5b5061010560048036038101906101009190610db5565b6108c5565b005b34801561011357600080fd5b5061012e60048036038101906101299190610b3f565b610993565b60405161013e9493929190610e26565b60405180910390f35b34801561015357600080fd5b5061015c6109f3565b6040516101699190610d93565b60405180910390f35b60018054905081106101b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101b090610ec8565b60405180910390fd5b6000600182815481106101cf576101ce610ee8565b5b90600052602060002090600402016040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815260200160028201548152602001600382015481525050905080604001513410156102a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161029c90610f89565b60405180910390fd5b806060015160008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1662fdd58e836000015184602001516040518363ffffffff1660e01b815260040161030c929190610fa9565b602060405180830381865afa158015610329573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061034d9190610fe7565b101561038e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161038590611086565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f242432a826000015133846020015185606001516040518563ffffffff1660e01b81526004016103f994939291906110dd565b600060405180830381600087803b15801561041357600080fd5b505af1158015610427573d6000803e3d6000fd5b505050506000816000015173ffffffffffffffffffffffffffffffffffffffff163460405161045590611163565b60006040518083038185875af1925050503d8060008114610492576040519150601f19603f3d011682016040523d82523d6000602084013e610497565b606091505b50509050806104db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d2906111ea565b60405180910390fd5b60018080805490506104ed9190611239565b815481106104fe576104fd610ee8565b5b9060005260206000209060040201600184815481106105205761051f610ee8565b5b90600052602060002090600402016000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018201548160010155600282015481600201556003820154816003015590505060018054806105c8576105c761126d565b5b6001900381819060005260206000209060040201600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560018201600090556002820160009055600382016000905550509055505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606000805b6001805490508110156106f6578373ffffffffffffffffffffffffffffffffffffffff166001828154811061068757610686610ee8565b5b906000526020600020906004020160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036106e35781806106df9061129c565b9250505b80806106ee9061129c565b91505061064f565b5060008167ffffffffffffffff811115610713576107126112e4565b5b60405190808252806020026020018201604052801561074c57816020015b610739610ac6565b8152602001906001900390816107315790505b5090506000805b6001805490508110156108b9578573ffffffffffffffffffffffffffffffffffffffff166001828154811061078b5761078a610ee8565b5b906000526020600020906004020160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036108a657600181815481106107ec576107eb610ee8565b5b90600052602060002090600402016040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152505083838151811061088c5761088b610ee8565b5b602002602001018190525081806108a29061129c565b9250505b80806108b19061129c565b915050610753565b50819350505050919050565b600160405180608001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200185815260200184815260200183815250908060018154018082558091505060019003906000526020600020906004020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155606082015181600301555050505050565b600181815481106109a357600080fd5b90600052602060002090600402016000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154908060030154905084565b60606001805480602002602001604051908101604052809291908181526020016000905b82821015610abd57838290600052602060002090600402016040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152505081526020019060010190610a17565b50505050905090565b6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081525090565b600080fd5b6000819050919050565b610b1c81610b09565b8114610b2757600080fd5b50565b600081359050610b3981610b13565b92915050565b600060208284031215610b5557610b54610b04565b5b6000610b6384828501610b2a565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610bb1610bac610ba784610b6c565b610b8c565b610b6c565b9050919050565b6000610bc382610b96565b9050919050565b6000610bd582610bb8565b9050919050565b610be581610bca565b82525050565b6000602082019050610c006000830184610bdc565b92915050565b6000610c1182610b6c565b9050919050565b610c2181610c06565b8114610c2c57600080fd5b50565b600081359050610c3e81610c18565b92915050565b600060208284031215610c5a57610c59610b04565b5b6000610c6884828501610c2f565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610ca681610c06565b82525050565b610cb581610b09565b82525050565b608082016000820151610cd16000850182610c9d565b506020820151610ce46020850182610cac565b506040820151610cf76040850182610cac565b506060820151610d0a6060850182610cac565b50505050565b6000610d1c8383610cbb565b60808301905092915050565b6000602082019050919050565b6000610d4082610c71565b610d4a8185610c7c565b9350610d5583610c8d565b8060005b83811015610d86578151610d6d8882610d10565b9750610d7883610d28565b925050600181019050610d59565b5085935050505092915050565b60006020820190508181036000830152610dad8184610d35565b905092915050565b600080600060608486031215610dce57610dcd610b04565b5b6000610ddc86828701610b2a565b9350506020610ded86828701610b2a565b9250506040610dfe86828701610b2a565b9150509250925092565b610e1181610c06565b82525050565b610e2081610b09565b82525050565b6000608082019050610e3b6000830187610e08565b610e486020830186610e17565b610e556040830185610e17565b610e626060830184610e17565b95945050505050565b600082825260208201905092915050565b7f496e76616c6964206c697374696e672049440000000000000000000000000000600082015250565b6000610eb2601283610e6b565b9150610ebd82610e7c565b602082019050919050565b60006020820190508181036000830152610ee181610ea5565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f53656e742076616c7565206973206c657373207468616e20746865206c69737460008201527f696e672070726963650000000000000000000000000000000000000000000000602082015250565b6000610f73602983610e6b565b9150610f7e82610f17565b604082019050919050565b60006020820190508181036000830152610fa281610f66565b9050919050565b6000604082019050610fbe6000830185610e08565b610fcb6020830184610e17565b9392505050565b600081519050610fe181610b13565b92915050565b600060208284031215610ffd57610ffc610b04565b5b600061100b84828501610fd2565b91505092915050565b7f53656c6c657220646f6573206e6f74206861766520656e6f75676820746f6b6560008201527f6e7320666f722073616c65000000000000000000000000000000000000000000602082015250565b6000611070602b83610e6b565b915061107b82611014565b604082019050919050565b6000602082019050818103600083015261109f81611063565b9050919050565b600082825260208201905092915050565b50565b60006110c76000836110a6565b91506110d2826110b7565b600082019050919050565b600060a0820190506110f26000830187610e08565b6110ff6020830186610e08565b61110c6040830185610e17565b6111196060830184610e17565b818103608083015261112a816110ba565b905095945050505050565b600081905092915050565b600061114d600083611135565b9150611158826110b7565b600082019050919050565b600061116e82611140565b9150819050919050565b7f4661696c656420746f207472616e7366657220457468657220746f207468652060008201527f73656c6c65720000000000000000000000000000000000000000000000000000602082015250565b60006111d4602683610e6b565b91506111df82611178565b604082019050919050565b60006020820190508181036000830152611203816111c7565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061124482610b09565b915061124f83610b09565b92508282039050818111156112675761126661120a565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60006112a782610b09565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036112d9576112d861120a565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea26469706673582212202646a5bdda11094a65bc1e78d6dfebfca8bfaa6971a352bc99e682f11a9401eb64736f6c63430008120033",
}

// SaleABI is the input ABI used to generate the binding from.
// Deprecated: Use SaleMetaData.ABI instead.
var SaleABI = SaleMetaData.ABI

// SaleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SaleMetaData.Bin instead.
var SaleBin = SaleMetaData.Bin

// DeploySale deploys a new Ethereum contract, binding an instance of Sale to it.
func DeploySale(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenContract common.Address) (common.Address, *types.Transaction, *Sale, error) {
	parsed, err := SaleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SaleBin), backend, _tokenContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sale{SaleCaller: SaleCaller{contract: contract}, SaleTransactor: SaleTransactor{contract: contract}, SaleFilterer: SaleFilterer{contract: contract}}, nil
}

// Sale is an auto generated Go binding around an Ethereum contract.
type Sale struct {
	SaleCaller     // Read-only binding to the contract
	SaleTransactor // Write-only binding to the contract
	SaleFilterer   // Log filterer for contract events
}

// SaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SaleSession struct {
	Contract     *Sale             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SaleCallerSession struct {
	Contract *SaleCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SaleTransactorSession struct {
	Contract     *SaleTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SaleRaw struct {
	Contract *Sale // Generic contract binding to access the raw methods on
}

// SaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SaleCallerRaw struct {
	Contract *SaleCaller // Generic read-only contract binding to access the raw methods on
}

// SaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SaleTransactorRaw struct {
	Contract *SaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSale creates a new instance of Sale, bound to a specific deployed contract.
func NewSale(address common.Address, backend bind.ContractBackend) (*Sale, error) {
	contract, err := bindSale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sale{SaleCaller: SaleCaller{contract: contract}, SaleTransactor: SaleTransactor{contract: contract}, SaleFilterer: SaleFilterer{contract: contract}}, nil
}

// NewSaleCaller creates a new read-only instance of Sale, bound to a specific deployed contract.
func NewSaleCaller(address common.Address, caller bind.ContractCaller) (*SaleCaller, error) {
	contract, err := bindSale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SaleCaller{contract: contract}, nil
}

// NewSaleTransactor creates a new write-only instance of Sale, bound to a specific deployed contract.
func NewSaleTransactor(address common.Address, transactor bind.ContractTransactor) (*SaleTransactor, error) {
	contract, err := bindSale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SaleTransactor{contract: contract}, nil
}

// NewSaleFilterer creates a new log filterer instance of Sale, bound to a specific deployed contract.
func NewSaleFilterer(address common.Address, filterer bind.ContractFilterer) (*SaleFilterer, error) {
	contract, err := bindSale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SaleFilterer{contract: contract}, nil
}

// bindSale binds a generic wrapper to an already deployed contract.
func bindSale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SaleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sale *SaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sale.Contract.SaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sale *SaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sale.Contract.SaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sale *SaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sale.Contract.SaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sale *SaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sale *SaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sale *SaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sale.Contract.contract.Transact(opts, method, params...)
}

// GetListings is a free data retrieval call binding the contract method 0xf1b2d6a3.
//
// Solidity: function getListings() view returns((address,uint256,uint256,uint256)[])
func (_Sale *SaleCaller) GetListings(opts *bind.CallOpts) ([]SaleListing, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "getListings")

	if err != nil {
		return *new([]SaleListing), err
	}

	out0 := *abi.ConvertType(out[0], new([]SaleListing)).(*[]SaleListing)

	return out0, err

}

// GetListings is a free data retrieval call binding the contract method 0xf1b2d6a3.
//
// Solidity: function getListings() view returns((address,uint256,uint256,uint256)[])
func (_Sale *SaleSession) GetListings() ([]SaleListing, error) {
	return _Sale.Contract.GetListings(&_Sale.CallOpts)
}

// GetListings is a free data retrieval call binding the contract method 0xf1b2d6a3.
//
// Solidity: function getListings() view returns((address,uint256,uint256,uint256)[])
func (_Sale *SaleCallerSession) GetListings() ([]SaleListing, error) {
	return _Sale.Contract.GetListings(&_Sale.CallOpts)
}

// GetListingsByUser is a free data retrieval call binding the contract method 0x69df8327.
//
// Solidity: function getListingsByUser(address user) view returns((address,uint256,uint256,uint256)[])
func (_Sale *SaleCaller) GetListingsByUser(opts *bind.CallOpts, user common.Address) ([]SaleListing, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "getListingsByUser", user)

	if err != nil {
		return *new([]SaleListing), err
	}

	out0 := *abi.ConvertType(out[0], new([]SaleListing)).(*[]SaleListing)

	return out0, err

}

// GetListingsByUser is a free data retrieval call binding the contract method 0x69df8327.
//
// Solidity: function getListingsByUser(address user) view returns((address,uint256,uint256,uint256)[])
func (_Sale *SaleSession) GetListingsByUser(user common.Address) ([]SaleListing, error) {
	return _Sale.Contract.GetListingsByUser(&_Sale.CallOpts, user)
}

// GetListingsByUser is a free data retrieval call binding the contract method 0x69df8327.
//
// Solidity: function getListingsByUser(address user) view returns((address,uint256,uint256,uint256)[])
func (_Sale *SaleCallerSession) GetListingsByUser(user common.Address) ([]SaleListing, error) {
	return _Sale.Contract.GetListingsByUser(&_Sale.CallOpts, user)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(address seller, uint256 tokenId, uint256 price, uint256 amount)
func (_Sale *SaleCaller) Listings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seller  common.Address
	TokenId *big.Int
	Price   *big.Int
	Amount  *big.Int
}, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "listings", arg0)

	outstruct := new(struct {
		Seller  common.Address
		TokenId *big.Int
		Price   *big.Int
		Amount  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(address seller, uint256 tokenId, uint256 price, uint256 amount)
func (_Sale *SaleSession) Listings(arg0 *big.Int) (struct {
	Seller  common.Address
	TokenId *big.Int
	Price   *big.Int
	Amount  *big.Int
}, error) {
	return _Sale.Contract.Listings(&_Sale.CallOpts, arg0)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(address seller, uint256 tokenId, uint256 price, uint256 amount)
func (_Sale *SaleCallerSession) Listings(arg0 *big.Int) (struct {
	Seller  common.Address
	TokenId *big.Int
	Price   *big.Int
	Amount  *big.Int
}, error) {
	return _Sale.Contract.Listings(&_Sale.CallOpts, arg0)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() view returns(address)
func (_Sale *SaleCaller) TokenContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "tokenContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() view returns(address)
func (_Sale *SaleSession) TokenContract() (common.Address, error) {
	return _Sale.Contract.TokenContract(&_Sale.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() view returns(address)
func (_Sale *SaleCallerSession) TokenContract() (common.Address, error) {
	return _Sale.Contract.TokenContract(&_Sale.CallOpts)
}

// BuyToken is a paid mutator transaction binding the contract method 0x2d296bf1.
//
// Solidity: function buyToken(uint256 listingId) payable returns()
func (_Sale *SaleTransactor) BuyToken(opts *bind.TransactOpts, listingId *big.Int) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "buyToken", listingId)
}

// BuyToken is a paid mutator transaction binding the contract method 0x2d296bf1.
//
// Solidity: function buyToken(uint256 listingId) payable returns()
func (_Sale *SaleSession) BuyToken(listingId *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.BuyToken(&_Sale.TransactOpts, listingId)
}

// BuyToken is a paid mutator transaction binding the contract method 0x2d296bf1.
//
// Solidity: function buyToken(uint256 listingId) payable returns()
func (_Sale *SaleTransactorSession) BuyToken(listingId *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.BuyToken(&_Sale.TransactOpts, listingId)
}

// ListToken is a paid mutator transaction binding the contract method 0xbce64a7d.
//
// Solidity: function listToken(uint256 tokenId, uint256 price, uint256 amount) returns()
func (_Sale *SaleTransactor) ListToken(opts *bind.TransactOpts, tokenId *big.Int, price *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "listToken", tokenId, price, amount)
}

// ListToken is a paid mutator transaction binding the contract method 0xbce64a7d.
//
// Solidity: function listToken(uint256 tokenId, uint256 price, uint256 amount) returns()
func (_Sale *SaleSession) ListToken(tokenId *big.Int, price *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.ListToken(&_Sale.TransactOpts, tokenId, price, amount)
}

// ListToken is a paid mutator transaction binding the contract method 0xbce64a7d.
//
// Solidity: function listToken(uint256 tokenId, uint256 price, uint256 amount) returns()
func (_Sale *SaleTransactorSession) ListToken(tokenId *big.Int, price *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.ListToken(&_Sale.TransactOpts, tokenId, price, amount)
}
