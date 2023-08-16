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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"addTokenOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"divs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getOriginalCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getOwnedTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getOwnerOfToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getTokensCreatedBy\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ipfsPaths\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"originalCreators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ownedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipfsPath\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"ticketPool\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"div\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"setMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ticketPools\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenNames\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokensCreatedBy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611738806100206000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80637f1dd80b116100a2578063d58778d611610071578063d58778d61461032a578063d9d616551461035a578063ddc0f3121461038a578063decf3e6d146103ba578063e149f036146103d65761010b565b80637f1dd80b1461026a5780638b3059301461029a578063af416e14146102ca578063caa47fbf146102fa5761010b565b8063543e062c116100de578063543e062c146101bc57806367f718a9146101ec578063681ece9d1461020a5780636d5454781461023a5761010b565b806310bf0717146101105780631749926514610140578063310495ab14610170578063351b3893146101a0575b600080fd5b61012a60048036038101906101259190610cee565b610406565b6040516101379190610dab565b60405180910390f35b61015a60048036038101906101559190610cee565b6104a6565b6040516101679190610ddc565b60405180910390f35b61018a60048036038101906101859190610cee565b6104be565b6040516101979190610dab565b60405180910390f35b6101ba60048036038101906101b59190610e55565b61055e565b005b6101d660048036038101906101d19190610e95565b6105d6565b6040516101e39190610ee4565b60405180910390f35b6101f4610624565b6040516102019190610fbd565b60405180910390f35b610224600480360381019061021f9190610cee565b61067c565b604051610231919061109d565b60405180910390f35b610254600480360381019061024f9190610cee565b61071d565b6040516102619190610ddc565b60405180910390f35b610284600480360381019061027f9190610cee565b610735565b6040516102919190610ee4565b60405180910390f35b6102b460048036038101906102af9190610cee565b61084c565b6040516102c19190610ee4565b60405180910390f35b6102e460048036038101906102df91906110bf565b61087f565b6040516102f19190610ddc565b60405180910390f35b610314600480360381019061030f9190610cee565b6108b0565b6040516103219190610ee4565b60405180910390f35b610344600480360381019061033f9190610cee565b6108ed565b6040516103519190610ddc565b60405180910390f35b610374600480360381019061036f91906110ff565b610911565b6040516103819190610fbd565b60405180910390f35b6103a4600480360381019061039f91906110ff565b6109a8565b6040516103b19190610fbd565b60405180910390f35b6103d460048036038101906103cf9190611261565b610a3f565b005b6103f060048036038101906103eb91906110bf565b610c73565b6040516103fd9190610ddc565b60405180910390f35b6001602052806000526040600020600091509050805461042590611355565b80601f016020809104026020016040519081016040528092919081815260200182805461045190611355565b801561049e5780601f106104735761010080835404028352916020019161049e565b820191906000526020600020905b81548152906001019060200180831161048157829003601f168201915b505050505081565b60036020528060005260406000206000915090505481565b600060205280600052604060002060009150905080546104dd90611355565b80601f016020809104026020016040519081016040528092919081815260200182805461050990611355565b80156105565780601f1061052b57610100808354040283529160200191610556565b820191906000526020600020905b81548152906001019060200180831161053957829003601f168201915b505050505081565b60046000838152602001908152602001600020819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600460205281600052604060002081815481106105f257600080fd5b906000526020600020016000915091509054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6060600880548060200260200160405190810160405280929190818152602001828054801561067257602002820191906000526020600020905b81548152602001906001019080831161065e575b5050505050905090565b60606004600083815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561071157602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116106c7575b50505050509050919050565b60026020528060005260406000206000915090505481565b600080600460008481526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156107cb57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610781575b5050505050905060008151111561080c5780600182516107eb91906113b5565b815181106107fc576107fb6113e9565b5b6020026020010151915050610847565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083e90611464565b60405180910390fd5b919050565b60056020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6007602052816000526040600020818154811061089b57600080fd5b90600052602060002001600091509150505481565b60006005600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b600881815481106108fd57600080fd5b906000526020600020016000915090505481565b6060600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561099c57602002820191906000526020600020905b815481526020019060010190808311610988575b50505050509050919050565b6060600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015610a3357602002820191906000526020600020905b815481526020019060010190808311610a1f575b50505050509050919050565b846000808881526020019081526020016000209081610a5e9190611630565b5083600160008881526020019081526020016000209081610a7f9190611630565b50826002600088815260200190815260200160002081905550816003600088815260200190815260200160002081905550600886908060018154018082558091505060019003906000526020600020016000909190919091505560046000878152602001908152602001600020819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806005600088815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020869080600181540180825580915050600190039060005260206000200160009091909190915055600760008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020869080600181540180825580915050600190039060005260206000200160009091909190915055505050505050565b60066020528160005260406000208181548110610c8f57600080fd5b90600052602060002001600091509150505481565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610ccb81610cb8565b8114610cd657600080fd5b50565b600081359050610ce881610cc2565b92915050565b600060208284031215610d0457610d03610cae565b5b6000610d1284828501610cd9565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610d55578082015181840152602081019050610d3a565b60008484015250505050565b6000601f19601f8301169050919050565b6000610d7d82610d1b565b610d878185610d26565b9350610d97818560208601610d37565b610da081610d61565b840191505092915050565b60006020820190508181036000830152610dc58184610d72565b905092915050565b610dd681610cb8565b82525050565b6000602082019050610df16000830184610dcd565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610e2282610df7565b9050919050565b610e3281610e17565b8114610e3d57600080fd5b50565b600081359050610e4f81610e29565b92915050565b60008060408385031215610e6c57610e6b610cae565b5b6000610e7a85828601610cd9565b9250506020610e8b85828601610e40565b9150509250929050565b60008060408385031215610eac57610eab610cae565b5b6000610eba85828601610cd9565b9250506020610ecb85828601610cd9565b9150509250929050565b610ede81610e17565b82525050565b6000602082019050610ef96000830184610ed5565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610f3481610cb8565b82525050565b6000610f468383610f2b565b60208301905092915050565b6000602082019050919050565b6000610f6a82610eff565b610f748185610f0a565b9350610f7f83610f1b565b8060005b83811015610fb0578151610f978882610f3a565b9750610fa283610f52565b925050600181019050610f83565b5085935050505092915050565b60006020820190508181036000830152610fd78184610f5f565b905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61101481610e17565b82525050565b6000611026838361100b565b60208301905092915050565b6000602082019050919050565b600061104a82610fdf565b6110548185610fea565b935061105f83610ffb565b8060005b83811015611090578151611077888261101a565b975061108283611032565b925050600181019050611063565b5085935050505092915050565b600060208201905081810360008301526110b7818461103f565b905092915050565b600080604083850312156110d6576110d5610cae565b5b60006110e485828601610e40565b92505060206110f585828601610cd9565b9150509250929050565b60006020828403121561111557611114610cae565b5b600061112384828501610e40565b91505092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61116e82610d61565b810181811067ffffffffffffffff8211171561118d5761118c611136565b5b80604052505050565b60006111a0610ca4565b90506111ac8282611165565b919050565b600067ffffffffffffffff8211156111cc576111cb611136565b5b6111d582610d61565b9050602081019050919050565b82818337600083830152505050565b60006112046111ff846111b1565b611196565b9050828152602081018484840111156112205761121f611131565b5b61122b8482856111e2565b509392505050565b600082601f8301126112485761124761112c565b5b81356112588482602086016111f1565b91505092915050565b60008060008060008060c0878903121561127e5761127d610cae565b5b600061128c89828a01610cd9565b965050602087013567ffffffffffffffff8111156112ad576112ac610cb3565b5b6112b989828a01611233565b955050604087013567ffffffffffffffff8111156112da576112d9610cb3565b5b6112e689828a01611233565b94505060606112f789828a01610cd9565b935050608061130889828a01610cd9565b92505060a061131989828a01610e40565b9150509295509295509295565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061136d57607f821691505b6020821081036113805761137f611326565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006113c082610cb8565b91506113cb83610cb8565b92508282039050818111156113e3576113e2611386565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f546f6b656e20646f6573206e6f74206578697374000000000000000000000000600082015250565b600061144e601483610d26565b915061145982611418565b602082019050919050565b6000602082019050818103600083015261147d81611441565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026114e67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826114a9565b6114f086836114a9565b95508019841693508086168417925050509392505050565b6000819050919050565b600061152d61152861152384610cb8565b611508565b610cb8565b9050919050565b6000819050919050565b61154783611512565b61155b61155382611534565b8484546114b6565b825550505050565b600090565b611570611563565b61157b81848461153e565b505050565b5b8181101561159f57611594600082611568565b600181019050611581565b5050565b601f8211156115e4576115b581611484565b6115be84611499565b810160208510156115cd578190505b6115e16115d985611499565b830182611580565b50505b505050565b600082821c905092915050565b6000611607600019846008026115e9565b1980831691505092915050565b600061162083836115f6565b9150826002028217905092915050565b61163982610d1b565b67ffffffffffffffff81111561165257611651611136565b5b61165c8254611355565b6116678282856115a3565b600060209050601f83116001811461169a5760008415611688578287015190505b6116928582611614565b8655506116fa565b601f1984166116a886611484565b60005b828110156116d0578489015182556001820191506020850194506020810190506116ab565b868310156116ed57848901516116e9601f8916826115f6565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220398bfdf56f12e2ff86c987ae6e97cc6d529fdf8cdf7a4dd6809829fefdd646dd64736f6c63430008120033",
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

// GetOriginalCreator is a free data retrieval call binding the contract method 0xcaa47fbf.
//
// Solidity: function getOriginalCreator(uint256 tokenId) view returns(address)
func (_Metadata *MetadataCaller) GetOriginalCreator(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Metadata.contract.Call(opts, &out, "getOriginalCreator", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOriginalCreator is a free data retrieval call binding the contract method 0xcaa47fbf.
//
// Solidity: function getOriginalCreator(uint256 tokenId) view returns(address)
func (_Metadata *MetadataSession) GetOriginalCreator(tokenId *big.Int) (common.Address, error) {
	return _Metadata.Contract.GetOriginalCreator(&_Metadata.CallOpts, tokenId)
}

// GetOriginalCreator is a free data retrieval call binding the contract method 0xcaa47fbf.
//
// Solidity: function getOriginalCreator(uint256 tokenId) view returns(address)
func (_Metadata *MetadataCallerSession) GetOriginalCreator(tokenId *big.Int) (common.Address, error) {
	return _Metadata.Contract.GetOriginalCreator(&_Metadata.CallOpts, tokenId)
}

// GetOwnedTokens is a free data retrieval call binding the contract method 0xd9d61655.
//
// Solidity: function getOwnedTokens(address owner) view returns(uint256[])
func (_Metadata *MetadataCaller) GetOwnedTokens(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Metadata.contract.Call(opts, &out, "getOwnedTokens", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOwnedTokens is a free data retrieval call binding the contract method 0xd9d61655.
//
// Solidity: function getOwnedTokens(address owner) view returns(uint256[])
func (_Metadata *MetadataSession) GetOwnedTokens(owner common.Address) ([]*big.Int, error) {
	return _Metadata.Contract.GetOwnedTokens(&_Metadata.CallOpts, owner)
}

// GetOwnedTokens is a free data retrieval call binding the contract method 0xd9d61655.
//
// Solidity: function getOwnedTokens(address owner) view returns(uint256[])
func (_Metadata *MetadataCallerSession) GetOwnedTokens(owner common.Address) ([]*big.Int, error) {
	return _Metadata.Contract.GetOwnedTokens(&_Metadata.CallOpts, owner)
}

// GetOwnerOfToken is a free data retrieval call binding the contract method 0x7f1dd80b.
//
// Solidity: function getOwnerOfToken(uint256 tokenId) view returns(address)
func (_Metadata *MetadataCaller) GetOwnerOfToken(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Metadata.contract.Call(opts, &out, "getOwnerOfToken", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwnerOfToken is a free data retrieval call binding the contract method 0x7f1dd80b.
//
// Solidity: function getOwnerOfToken(uint256 tokenId) view returns(address)
func (_Metadata *MetadataSession) GetOwnerOfToken(tokenId *big.Int) (common.Address, error) {
	return _Metadata.Contract.GetOwnerOfToken(&_Metadata.CallOpts, tokenId)
}

// GetOwnerOfToken is a free data retrieval call binding the contract method 0x7f1dd80b.
//
// Solidity: function getOwnerOfToken(uint256 tokenId) view returns(address)
func (_Metadata *MetadataCallerSession) GetOwnerOfToken(tokenId *big.Int) (common.Address, error) {
	return _Metadata.Contract.GetOwnerOfToken(&_Metadata.CallOpts, tokenId)
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

// GetTokenOwners is a free data retrieval call binding the contract method 0x681ece9d.
//
// Solidity: function getTokenOwners(uint256 tokenId) view returns(address[])
func (_Metadata *MetadataCaller) GetTokenOwners(opts *bind.CallOpts, tokenId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Metadata.contract.Call(opts, &out, "getTokenOwners", tokenId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenOwners is a free data retrieval call binding the contract method 0x681ece9d.
//
// Solidity: function getTokenOwners(uint256 tokenId) view returns(address[])
func (_Metadata *MetadataSession) GetTokenOwners(tokenId *big.Int) ([]common.Address, error) {
	return _Metadata.Contract.GetTokenOwners(&_Metadata.CallOpts, tokenId)
}

// GetTokenOwners is a free data retrieval call binding the contract method 0x681ece9d.
//
// Solidity: function getTokenOwners(uint256 tokenId) view returns(address[])
func (_Metadata *MetadataCallerSession) GetTokenOwners(tokenId *big.Int) ([]common.Address, error) {
	return _Metadata.Contract.GetTokenOwners(&_Metadata.CallOpts, tokenId)
}

// GetTokensCreatedBy is a free data retrieval call binding the contract method 0xddc0f312.
//
// Solidity: function getTokensCreatedBy(address owner) view returns(uint256[])
func (_Metadata *MetadataCaller) GetTokensCreatedBy(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Metadata.contract.Call(opts, &out, "getTokensCreatedBy", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTokensCreatedBy is a free data retrieval call binding the contract method 0xddc0f312.
//
// Solidity: function getTokensCreatedBy(address owner) view returns(uint256[])
func (_Metadata *MetadataSession) GetTokensCreatedBy(owner common.Address) ([]*big.Int, error) {
	return _Metadata.Contract.GetTokensCreatedBy(&_Metadata.CallOpts, owner)
}

// GetTokensCreatedBy is a free data retrieval call binding the contract method 0xddc0f312.
//
// Solidity: function getTokensCreatedBy(address owner) view returns(uint256[])
func (_Metadata *MetadataCallerSession) GetTokensCreatedBy(owner common.Address) ([]*big.Int, error) {
	return _Metadata.Contract.GetTokensCreatedBy(&_Metadata.CallOpts, owner)
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

// OriginalCreators is a free data retrieval call binding the contract method 0x8b305930.
//
// Solidity: function originalCreators(uint256 ) view returns(address)
func (_Metadata *MetadataCaller) OriginalCreators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Metadata.contract.Call(opts, &out, "originalCreators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OriginalCreators is a free data retrieval call binding the contract method 0x8b305930.
//
// Solidity: function originalCreators(uint256 ) view returns(address)
func (_Metadata *MetadataSession) OriginalCreators(arg0 *big.Int) (common.Address, error) {
	return _Metadata.Contract.OriginalCreators(&_Metadata.CallOpts, arg0)
}

// OriginalCreators is a free data retrieval call binding the contract method 0x8b305930.
//
// Solidity: function originalCreators(uint256 ) view returns(address)
func (_Metadata *MetadataCallerSession) OriginalCreators(arg0 *big.Int) (common.Address, error) {
	return _Metadata.Contract.OriginalCreators(&_Metadata.CallOpts, arg0)
}

// OwnedTokens is a free data retrieval call binding the contract method 0xe149f036.
//
// Solidity: function ownedTokens(address , uint256 ) view returns(uint256)
func (_Metadata *MetadataCaller) OwnedTokens(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Metadata.contract.Call(opts, &out, "ownedTokens", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnedTokens is a free data retrieval call binding the contract method 0xe149f036.
//
// Solidity: function ownedTokens(address , uint256 ) view returns(uint256)
func (_Metadata *MetadataSession) OwnedTokens(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Metadata.Contract.OwnedTokens(&_Metadata.CallOpts, arg0, arg1)
}

// OwnedTokens is a free data retrieval call binding the contract method 0xe149f036.
//
// Solidity: function ownedTokens(address , uint256 ) view returns(uint256)
func (_Metadata *MetadataCallerSession) OwnedTokens(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Metadata.Contract.OwnedTokens(&_Metadata.CallOpts, arg0, arg1)
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

// TokenOwners is a free data retrieval call binding the contract method 0x543e062c.
//
// Solidity: function tokenOwners(uint256 , uint256 ) view returns(address)
func (_Metadata *MetadataCaller) TokenOwners(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Metadata.contract.Call(opts, &out, "tokenOwners", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenOwners is a free data retrieval call binding the contract method 0x543e062c.
//
// Solidity: function tokenOwners(uint256 , uint256 ) view returns(address)
func (_Metadata *MetadataSession) TokenOwners(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Metadata.Contract.TokenOwners(&_Metadata.CallOpts, arg0, arg1)
}

// TokenOwners is a free data retrieval call binding the contract method 0x543e062c.
//
// Solidity: function tokenOwners(uint256 , uint256 ) view returns(address)
func (_Metadata *MetadataCallerSession) TokenOwners(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Metadata.Contract.TokenOwners(&_Metadata.CallOpts, arg0, arg1)
}

// TokensCreatedBy is a free data retrieval call binding the contract method 0xaf416e14.
//
// Solidity: function tokensCreatedBy(address , uint256 ) view returns(uint256)
func (_Metadata *MetadataCaller) TokensCreatedBy(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Metadata.contract.Call(opts, &out, "tokensCreatedBy", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensCreatedBy is a free data retrieval call binding the contract method 0xaf416e14.
//
// Solidity: function tokensCreatedBy(address , uint256 ) view returns(uint256)
func (_Metadata *MetadataSession) TokensCreatedBy(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Metadata.Contract.TokensCreatedBy(&_Metadata.CallOpts, arg0, arg1)
}

// TokensCreatedBy is a free data retrieval call binding the contract method 0xaf416e14.
//
// Solidity: function tokensCreatedBy(address , uint256 ) view returns(uint256)
func (_Metadata *MetadataCallerSession) TokensCreatedBy(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Metadata.Contract.TokensCreatedBy(&_Metadata.CallOpts, arg0, arg1)
}

// AddTokenOwner is a paid mutator transaction binding the contract method 0x351b3893.
//
// Solidity: function addTokenOwner(uint256 tokenId, address newOwner) returns()
func (_Metadata *MetadataTransactor) AddTokenOwner(opts *bind.TransactOpts, tokenId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _Metadata.contract.Transact(opts, "addTokenOwner", tokenId, newOwner)
}

// AddTokenOwner is a paid mutator transaction binding the contract method 0x351b3893.
//
// Solidity: function addTokenOwner(uint256 tokenId, address newOwner) returns()
func (_Metadata *MetadataSession) AddTokenOwner(tokenId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _Metadata.Contract.AddTokenOwner(&_Metadata.TransactOpts, tokenId, newOwner)
}

// AddTokenOwner is a paid mutator transaction binding the contract method 0x351b3893.
//
// Solidity: function addTokenOwner(uint256 tokenId, address newOwner) returns()
func (_Metadata *MetadataTransactorSession) AddTokenOwner(tokenId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _Metadata.Contract.AddTokenOwner(&_Metadata.TransactOpts, tokenId, newOwner)
}

// SetMetadata is a paid mutator transaction binding the contract method 0xdecf3e6d.
//
// Solidity: function setMetadata(uint256 tokenId, string tokenName, string ipfsPath, uint256 ticketPool, uint256 div, address sender) returns()
func (_Metadata *MetadataTransactor) SetMetadata(opts *bind.TransactOpts, tokenId *big.Int, tokenName string, ipfsPath string, ticketPool *big.Int, div *big.Int, sender common.Address) (*types.Transaction, error) {
	return _Metadata.contract.Transact(opts, "setMetadata", tokenId, tokenName, ipfsPath, ticketPool, div, sender)
}

// SetMetadata is a paid mutator transaction binding the contract method 0xdecf3e6d.
//
// Solidity: function setMetadata(uint256 tokenId, string tokenName, string ipfsPath, uint256 ticketPool, uint256 div, address sender) returns()
func (_Metadata *MetadataSession) SetMetadata(tokenId *big.Int, tokenName string, ipfsPath string, ticketPool *big.Int, div *big.Int, sender common.Address) (*types.Transaction, error) {
	return _Metadata.Contract.SetMetadata(&_Metadata.TransactOpts, tokenId, tokenName, ipfsPath, ticketPool, div, sender)
}

// SetMetadata is a paid mutator transaction binding the contract method 0xdecf3e6d.
//
// Solidity: function setMetadata(uint256 tokenId, string tokenName, string ipfsPath, uint256 ticketPool, uint256 div, address sender) returns()
func (_Metadata *MetadataTransactorSession) SetMetadata(tokenId *big.Int, tokenName string, ipfsPath string, ticketPool *big.Int, div *big.Int, sender common.Address) (*types.Transaction, error) {
	return _Metadata.Contract.SetMetadata(&_Metadata.TransactOpts, tokenId, tokenName, ipfsPath, ticketPool, div, sender)
}
