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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"divs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getOriginalCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getOwnedTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getOwnerOfToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getTokensCreatedBy\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ipfsPaths\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"originalCreators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ownedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipfsPath\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"ticketPool\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"div\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"setMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ticketPools\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenNames\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokensCreatedBy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611659806100206000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80638b30593011610097578063d9d6165511610066578063d9d6165514610333578063ddc0f31214610363578063decf3e6d14610393578063e149f036146103af57610100565b80638b30593014610273578063af416e14146102a3578063caa47fbf146102d3578063d58778d61461030357610100565b806367f718a9116100d357806367f718a9146101c5578063681ece9d146101e35780636d545478146102135780637f1dd80b1461024357610100565b806310bf0717146101055780631749926514610135578063310495ab14610165578063543e062c14610195575b600080fd5b61011f600480360381019061011a9190610c4f565b6103df565b60405161012c9190610d0c565b60405180910390f35b61014f600480360381019061014a9190610c4f565b61047f565b60405161015c9190610d3d565b60405180910390f35b61017f600480360381019061017a9190610c4f565b610497565b60405161018c9190610d0c565b60405180910390f35b6101af60048036038101906101aa9190610d58565b610537565b6040516101bc9190610dd9565b60405180910390f35b6101cd610585565b6040516101da9190610eb2565b60405180910390f35b6101fd60048036038101906101f89190610c4f565b6105dd565b60405161020a9190610f92565b60405180910390f35b61022d60048036038101906102289190610c4f565b61067e565b60405161023a9190610d3d565b60405180910390f35b61025d60048036038101906102589190610c4f565b610696565b60405161026a9190610dd9565b60405180910390f35b61028d60048036038101906102889190610c4f565b6107ad565b60405161029a9190610dd9565b60405180910390f35b6102bd60048036038101906102b89190610fe0565b6107e0565b6040516102ca9190610d3d565b60405180910390f35b6102ed60048036038101906102e89190610c4f565b610811565b6040516102fa9190610dd9565b60405180910390f35b61031d60048036038101906103189190610c4f565b61084e565b60405161032a9190610d3d565b60405180910390f35b61034d60048036038101906103489190611020565b610872565b60405161035a9190610eb2565b60405180910390f35b61037d60048036038101906103789190611020565b610909565b60405161038a9190610eb2565b60405180910390f35b6103ad60048036038101906103a89190611182565b6109a0565b005b6103c960048036038101906103c49190610fe0565b610bd4565b6040516103d69190610d3d565b60405180910390f35b600160205280600052604060002060009150905080546103fe90611276565b80601f016020809104026020016040519081016040528092919081815260200182805461042a90611276565b80156104775780601f1061044c57610100808354040283529160200191610477565b820191906000526020600020905b81548152906001019060200180831161045a57829003601f168201915b505050505081565b60036020528060005260406000206000915090505481565b600060205280600052604060002060009150905080546104b690611276565b80601f01602080910402602001604051908101604052809291908181526020018280546104e290611276565b801561052f5780601f106105045761010080835404028352916020019161052f565b820191906000526020600020905b81548152906001019060200180831161051257829003601f168201915b505050505081565b6004602052816000526040600020818154811061055357600080fd5b906000526020600020016000915091509054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606060088054806020026020016040519081016040528092919081815260200182805480156105d357602002820191906000526020600020905b8154815260200190600101908083116105bf575b5050505050905090565b60606004600083815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561067257602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610628575b50505050509050919050565b60026020528060005260406000206000915090505481565b6000806004600084815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561072c57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116106e2575b5050505050905060008151111561076d57806001825161074c91906112d6565b8151811061075d5761075c61130a565b5b60200260200101519150506107a8565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079f90611385565b60405180910390fd5b919050565b60066020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600760205281600052604060002081815481106107fc57600080fd5b90600052602060002001600091509150505481565b60006006600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6008818154811061085e57600080fd5b906000526020600020016000915090505481565b6060600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156108fd57602002820191906000526020600020905b8154815260200190600101908083116108e9575b50505050509050919050565b6060600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561099457602002820191906000526020600020905b815481526020019060010190808311610980575b50505050509050919050565b8460008088815260200190815260200160002090816109bf9190611551565b50836001600088815260200190815260200160002090816109e09190611551565b50826002600088815260200190815260200160002081905550816003600088815260200190815260200160002081905550600886908060018154018082558091505060019003906000526020600020016000909190919091505560046000878152602001908152602001600020819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806006600088815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600560008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020869080600181540180825580915050600190039060005260206000200160009091909190915055600760008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020869080600181540180825580915050600190039060005260206000200160009091909190915055505050505050565b60056020528160005260406000208181548110610bf057600080fd5b90600052602060002001600091509150505481565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610c2c81610c19565b8114610c3757600080fd5b50565b600081359050610c4981610c23565b92915050565b600060208284031215610c6557610c64610c0f565b5b6000610c7384828501610c3a565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610cb6578082015181840152602081019050610c9b565b60008484015250505050565b6000601f19601f8301169050919050565b6000610cde82610c7c565b610ce88185610c87565b9350610cf8818560208601610c98565b610d0181610cc2565b840191505092915050565b60006020820190508181036000830152610d268184610cd3565b905092915050565b610d3781610c19565b82525050565b6000602082019050610d526000830184610d2e565b92915050565b60008060408385031215610d6f57610d6e610c0f565b5b6000610d7d85828601610c3a565b9250506020610d8e85828601610c3a565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610dc382610d98565b9050919050565b610dd381610db8565b82525050565b6000602082019050610dee6000830184610dca565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610e2981610c19565b82525050565b6000610e3b8383610e20565b60208301905092915050565b6000602082019050919050565b6000610e5f82610df4565b610e698185610dff565b9350610e7483610e10565b8060005b83811015610ea5578151610e8c8882610e2f565b9750610e9783610e47565b925050600181019050610e78565b5085935050505092915050565b60006020820190508181036000830152610ecc8184610e54565b905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610f0981610db8565b82525050565b6000610f1b8383610f00565b60208301905092915050565b6000602082019050919050565b6000610f3f82610ed4565b610f498185610edf565b9350610f5483610ef0565b8060005b83811015610f85578151610f6c8882610f0f565b9750610f7783610f27565b925050600181019050610f58565b5085935050505092915050565b60006020820190508181036000830152610fac8184610f34565b905092915050565b610fbd81610db8565b8114610fc857600080fd5b50565b600081359050610fda81610fb4565b92915050565b60008060408385031215610ff757610ff6610c0f565b5b600061100585828601610fcb565b925050602061101685828601610c3a565b9150509250929050565b60006020828403121561103657611035610c0f565b5b600061104484828501610fcb565b91505092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61108f82610cc2565b810181811067ffffffffffffffff821117156110ae576110ad611057565b5b80604052505050565b60006110c1610c05565b90506110cd8282611086565b919050565b600067ffffffffffffffff8211156110ed576110ec611057565b5b6110f682610cc2565b9050602081019050919050565b82818337600083830152505050565b6000611125611120846110d2565b6110b7565b90508281526020810184848401111561114157611140611052565b5b61114c848285611103565b509392505050565b600082601f8301126111695761116861104d565b5b8135611179848260208601611112565b91505092915050565b60008060008060008060c0878903121561119f5761119e610c0f565b5b60006111ad89828a01610c3a565b965050602087013567ffffffffffffffff8111156111ce576111cd610c14565b5b6111da89828a01611154565b955050604087013567ffffffffffffffff8111156111fb576111fa610c14565b5b61120789828a01611154565b945050606061121889828a01610c3a565b935050608061122989828a01610c3a565b92505060a061123a89828a01610fcb565b9150509295509295509295565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061128e57607f821691505b6020821081036112a1576112a0611247565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006112e182610c19565b91506112ec83610c19565b9250828203905081811115611304576113036112a7565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f546f6b656e20646f6573206e6f74206578697374000000000000000000000000600082015250565b600061136f601483610c87565b915061137a82611339565b602082019050919050565b6000602082019050818103600083015261139e81611362565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026114077fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826113ca565b61141186836113ca565b95508019841693508086168417925050509392505050565b6000819050919050565b600061144e61144961144484610c19565b611429565b610c19565b9050919050565b6000819050919050565b61146883611433565b61147c61147482611455565b8484546113d7565b825550505050565b600090565b611491611484565b61149c81848461145f565b505050565b5b818110156114c0576114b5600082611489565b6001810190506114a2565b5050565b601f821115611505576114d6816113a5565b6114df846113ba565b810160208510156114ee578190505b6115026114fa856113ba565b8301826114a1565b50505b505050565b600082821c905092915050565b60006115286000198460080261150a565b1980831691505092915050565b60006115418383611517565b9150826002028217905092915050565b61155a82610c7c565b67ffffffffffffffff81111561157357611572611057565b5b61157d8254611276565b6115888282856114c4565b600060209050601f8311600181146115bb57600084156115a9578287015190505b6115b38582611535565b86555061161b565b601f1984166115c9866113a5565b60005b828110156115f1578489015182556001820191506020850194506020810190506115cc565b8683101561160e578489015161160a601f891682611517565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220a819882db894e852ce389a1fae2ae57071de95c6a6bc766f4f05a7f1fe90eeb464736f6c63430008120033",
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
