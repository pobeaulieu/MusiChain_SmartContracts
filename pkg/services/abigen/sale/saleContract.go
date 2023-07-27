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
	TokenId   *big.Int
	Seller    common.Address
	Price     *big.Int
	Amount    *big.Int
	IsForSale bool
}

// SaleMetaData contains all meta data concerning the Sale contract.
var SaleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"buyToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"etherValue\",\"type\":\"uint256\"}],\"name\":\"convertToWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllListings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForSale\",\"type\":\"bool\"}],\"internalType\":\"structSale.Listing[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getListingsByUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForSale\",\"type\":\"bool\"}],\"internalType\":\"structSale.Listing[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getOwnedTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInEther\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"listToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForSale\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"removeListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001f1a38038062001f1a8339818101604052810190620000379190620000fc565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200012e565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000b08262000083565b9050919050565b6000620000c482620000a3565b9050919050565b620000d681620000b7565b8114620000e257600080fd5b50565b600081519050620000f681620000cb565b92915050565b6000602082840312156200011557620001146200007e565b5b60006200012584828501620000e5565b91505092915050565b611ddc806200013e6000396000f3fe6080604052600436106100915760003560e01c8063ae73ccec11610059578063ae73ccec14610180578063bce64a7d146101ab578063d58778d6146101d4578063d9d6165514610211578063de74e57b1461024e57610091565b80632d296bf114610096578063479ad4c3146100b257806355a373d6146100db57806369df832714610106578063a60dc38a14610143575b600080fd5b6100b060048036038101906100ab9190611274565b61028f565b005b3480156100be57600080fd5b506100d960048036038101906100d49190611274565b610654565b005b3480156100e757600080fd5b506100f061072a565b6040516100fd9190611320565b60405180910390f35b34801561011257600080fd5b5061012d60048036038101906101289190611379565b61074e565b60405161013a91906114f6565b60405180910390f35b34801561014f57600080fd5b5061016a60048036038101906101659190611274565b610aa8565b6040516101779190611527565b60405180910390f35b34801561018c57600080fd5b50610195610ac5565b6040516101a291906114f6565b60405180910390f35b3480156101b757600080fd5b506101d260048036038101906101cd9190611542565b610d07565b005b3480156101e057600080fd5b506101fb60048036038101906101f69190611274565b610ee7565b6040516102089190611527565b60405180910390f35b34801561021d57600080fd5b5061023860048036038101906102339190611379565b610f0b565b6040516102459190611644565b60405180910390f35b34801561025a57600080fd5b5061027560048036038101906102709190611274565b61118f565b604051610286959493929190611684565b60405180910390f35b6002600082815260200190815260200160002060040160009054906101000a900460ff166102f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102e990611734565b60405180910390fd5b6000600260008381526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201548152602001600382015481526020016004820160009054906101000a900460ff161515151581525050905080604001513410156103ea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e1906117c6565b60405180910390fd5b806060015160008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1662fdd58e836020015184600001516040518363ffffffff1660e01b81526004016104519291906117e6565b602060405180830381865afa15801561046e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104929190611824565b10156104d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ca906118c3565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f242432a826020015133846000015185606001516040518563ffffffff1660e01b815260040161053e949392919061191a565b600060405180830381600087803b15801561055857600080fd5b505af115801561056c573d6000803e3d6000fd5b505050506000816020015173ffffffffffffffffffffffffffffffffffffffff163460405161059a906119a0565b60006040518083038185875af1925050503d80600081146105d7576040519150601f19603f3d011682016040523d82523d6000602084013e6105dc565b606091505b5050905080610620576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161061790611a27565b60405180910390fd5b60006002600085815260200190815260200160002060040160006101000a81548160ff021916908315150217905550505050565b3373ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146106f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ef90611ab9565b60405180910390fd5b60006002600083815260200190815260200160002060040160006101000a81548160ff02191690831515021790555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606000805b600180549050811015610855578373ffffffffffffffffffffffffffffffffffffffff1660026000600184815481106107905761078f611ad9565b5b9060005260206000200154815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614801561082e5750600260006001838154811061080357610802611ad9565b5b9060005260206000200154815260200190815260200160002060040160009054906101000a900460ff165b1561084257818061083e90611b37565b9250505b808061084d90611b37565b915050610754565b5060008167ffffffffffffffff81111561087257610871611b7f565b5b6040519080825280602002602001820160405280156108ab57816020015b6108986111f2565b8152602001906001900390816108905790505b5090506000805b600180549050811015610a9c578573ffffffffffffffffffffffffffffffffffffffff1660026000600184815481106108ee576108ed611ad9565b5b9060005260206000200154815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614801561098c5750600260006001838154811061096157610960611ad9565b5b9060005260206000200154815260200190815260200160002060040160009054906101000a900460ff165b15610a895760026000600183815481106109a9576109a8611ad9565b5b906000526020600020015481526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201548152602001600382015481526020016004820160009054906101000a900460ff161515151581525050838381518110610a6f57610a6e611ad9565b5b60200260200101819052508180610a8590611b37565b9250505b8080610a9490611b37565b9150506108b2565b50819350505050919050565b6000670de0b6b3a764000082610abe9190611bae565b9050919050565b60606000805b600180549050811015610b41576002600060018381548110610af057610aef611ad9565b5b9060005260206000200154815260200190815260200160002060040160009054906101000a900460ff1615610b2e578180610b2a90611b37565b9250505b8080610b3990611b37565b915050610acb565b5060008167ffffffffffffffff811115610b5e57610b5d611b7f565b5b604051908082528060200260200182016040528015610b9757816020015b610b846111f2565b815260200190600190039081610b7c5790505b5090506000805b600180549050811015610cfd576002600060018381548110610bc357610bc2611ad9565b5b9060005260206000200154815260200190815260200160002060040160009054906101000a900460ff1615610cea576002600060018381548110610c0a57610c09611ad9565b5b906000526020600020015481526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201548152602001600382015481526020016004820160009054906101000a900460ff161515151581525050838381518110610cd057610ccf611ad9565b5b60200260200101819052508180610ce690611b37565b9250505b8080610cf590611b37565b915050610b9e565b5081935050505090565b60008311610d4a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d4190611c62565b60405180910390fd5b60008211610d8d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d8490611cf4565b60405180910390fd5b60008111610dd0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dc790611d86565b60405180910390fd5b6000610ddb83610aa8565b90506040518060a001604052808581526020013373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200183815260200160011515815250600260008681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201556060820151816003015560808201518160040160006101000a81548160ff021916908315150217905550905050600184908060018154018082558091505060019003906000526020600020016000909190919091505550505050565b60018181548110610ef757600080fd5b906000526020600020016000915090505481565b60606000805b6001805490508110156110025760008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1662fdd58e8660018581548110610f7157610f70611ad9565b5b90600052602060002001546040518363ffffffff1660e01b8152600401610f999291906117e6565b602060405180830381865afa158015610fb6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fda9190611824565b1115610fef578180610feb90611b37565b9250505b8080610ffa90611b37565b915050610f11565b5060008167ffffffffffffffff81111561101f5761101e611b7f565b5b60405190808252806020026020018201604052801561104d5781602001602082028036833780820191505090505b5090506000805b6001805490508110156111835760008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1662fdd58e88600185815481106110b4576110b3611ad9565b5b90600052602060002001546040518363ffffffff1660e01b81526004016110dc9291906117e6565b602060405180830381865afa1580156110f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061111d9190611824565b1115611170576001818154811061113757611136611ad9565b5b906000526020600020015483838151811061115557611154611ad9565b5b602002602001018181525050818061116c90611b37565b9250505b808061117b90611b37565b915050611054565b50819350505050919050565b60026020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030154908060040160009054906101000a900460ff16905085565b6040518060a0016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081526020016000151581525090565b600080fd5b6000819050919050565b6112518161123e565b811461125c57600080fd5b50565b60008135905061126e81611248565b92915050565b60006020828403121561128a57611289611239565b5b60006112988482850161125f565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006112e66112e16112dc846112a1565b6112c1565b6112a1565b9050919050565b60006112f8826112cb565b9050919050565b600061130a826112ed565b9050919050565b61131a816112ff565b82525050565b60006020820190506113356000830184611311565b92915050565b6000611346826112a1565b9050919050565b6113568161133b565b811461136157600080fd5b50565b6000813590506113738161134d565b92915050565b60006020828403121561138f5761138e611239565b5b600061139d84828501611364565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6113db8161123e565b82525050565b6113ea8161133b565b82525050565b60008115159050919050565b611405816113f0565b82525050565b60a08201600082015161142160008501826113d2565b50602082015161143460208501826113e1565b50604082015161144760408501826113d2565b50606082015161145a60608501826113d2565b50608082015161146d60808501826113fc565b50505050565b600061147f838361140b565b60a08301905092915050565b6000602082019050919050565b60006114a3826113a6565b6114ad81856113b1565b93506114b8836113c2565b8060005b838110156114e95781516114d08882611473565b97506114db8361148b565b9250506001810190506114bc565b5085935050505092915050565b600060208201905081810360008301526115108184611498565b905092915050565b6115218161123e565b82525050565b600060208201905061153c6000830184611518565b92915050565b60008060006060848603121561155b5761155a611239565b5b60006115698682870161125f565b935050602061157a8682870161125f565b925050604061158b8682870161125f565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60006115cd83836113d2565b60208301905092915050565b6000602082019050919050565b60006115f182611595565b6115fb81856115a0565b9350611606836115b1565b8060005b8381101561163757815161161e88826115c1565b9750611629836115d9565b92505060018101905061160a565b5085935050505092915050565b6000602082019050818103600083015261165e81846115e6565b905092915050565b61166f8161133b565b82525050565b61167e816113f0565b82525050565b600060a0820190506116996000830188611518565b6116a66020830187611666565b6116b36040830186611518565b6116c06060830185611518565b6116cd6080830184611675565b9695505050505050565b600082825260208201905092915050565b7f54686973206974656d206973206e6f7420666f722073616c6500000000000000600082015250565b600061171e6019836116d7565b9150611729826116e8565b602082019050919050565b6000602082019050818103600083015261174d81611711565b9050919050565b7f53656e742076616c7565206973206c657373207468616e20746865206c69737460008201527f696e672070726963650000000000000000000000000000000000000000000000602082015250565b60006117b06029836116d7565b91506117bb82611754565b604082019050919050565b600060208201905081810360008301526117df816117a3565b9050919050565b60006040820190506117fb6000830185611666565b6118086020830184611518565b9392505050565b60008151905061181e81611248565b92915050565b60006020828403121561183a57611839611239565b5b60006118488482850161180f565b91505092915050565b7f53656c6c657220646f6573206e6f74206861766520656e6f75676820746f6b6560008201527f6e7320666f722073616c65000000000000000000000000000000000000000000602082015250565b60006118ad602b836116d7565b91506118b882611851565b604082019050919050565b600060208201905081810360008301526118dc816118a0565b9050919050565b600082825260208201905092915050565b50565b60006119046000836118e3565b915061190f826118f4565b600082019050919050565b600060a08201905061192f6000830187611666565b61193c6020830186611666565b6119496040830185611518565b6119566060830184611518565b8181036080830152611967816118f7565b905095945050505050565b600081905092915050565b600061198a600083611972565b9150611995826118f4565b600082019050919050565b60006119ab8261197d565b9150819050919050565b7f4661696c656420746f207472616e7366657220457468657220746f207468652060008201527f73656c6c65720000000000000000000000000000000000000000000000000000602082015250565b6000611a116026836116d7565b9150611a1c826119b5565b604082019050919050565b60006020820190508181036000830152611a4081611a04565b9050919050565b7f4f6e6c79207468652073656c6c65722063616e2072656d6f766520746865206c60008201527f697374696e670000000000000000000000000000000000000000000000000000602082015250565b6000611aa36026836116d7565b9150611aae82611a47565b604082019050919050565b60006020820190508181036000830152611ad281611a96565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611b428261123e565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611b7457611b73611b08565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000611bb98261123e565b9150611bc48361123e565b9250828202611bd28161123e565b91508282048414831517611be957611be8611b08565b5b5092915050565b7f546f6b656e2049442073686f756c642062652067726561746572207468616e2060008201527f7a65726f00000000000000000000000000000000000000000000000000000000602082015250565b6000611c4c6024836116d7565b9150611c5782611bf0565b604082019050919050565b60006020820190508181036000830152611c7b81611c3f565b9050919050565b7f50726963652073686f756c642062652067726561746572207468616e207a657260008201527f6f00000000000000000000000000000000000000000000000000000000000000602082015250565b6000611cde6021836116d7565b9150611ce982611c82565b604082019050919050565b60006020820190508181036000830152611d0d81611cd1565b9050919050565b7f416d6f756e742073686f756c642062652067726561746572207468616e207a6560008201527f726f000000000000000000000000000000000000000000000000000000000000602082015250565b6000611d706022836116d7565b9150611d7b82611d14565b604082019050919050565b60006020820190508181036000830152611d9f81611d63565b905091905056fea264697066735822122018452b0f7c52620be620a2bbc496b6d9f1ff4c2af46d07049b98660d6fd8ff7464736f6c63430008120033",
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

// ConvertToWei is a free data retrieval call binding the contract method 0xa60dc38a.
//
// Solidity: function convertToWei(uint256 etherValue) pure returns(uint256)
func (_Sale *SaleCaller) ConvertToWei(opts *bind.CallOpts, etherValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "convertToWei", etherValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToWei is a free data retrieval call binding the contract method 0xa60dc38a.
//
// Solidity: function convertToWei(uint256 etherValue) pure returns(uint256)
func (_Sale *SaleSession) ConvertToWei(etherValue *big.Int) (*big.Int, error) {
	return _Sale.Contract.ConvertToWei(&_Sale.CallOpts, etherValue)
}

// ConvertToWei is a free data retrieval call binding the contract method 0xa60dc38a.
//
// Solidity: function convertToWei(uint256 etherValue) pure returns(uint256)
func (_Sale *SaleCallerSession) ConvertToWei(etherValue *big.Int) (*big.Int, error) {
	return _Sale.Contract.ConvertToWei(&_Sale.CallOpts, etherValue)
}

// GetAllListings is a free data retrieval call binding the contract method 0xae73ccec.
//
// Solidity: function getAllListings() view returns((uint256,address,uint256,uint256,bool)[])
func (_Sale *SaleCaller) GetAllListings(opts *bind.CallOpts) ([]SaleListing, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "getAllListings")

	if err != nil {
		return *new([]SaleListing), err
	}

	out0 := *abi.ConvertType(out[0], new([]SaleListing)).(*[]SaleListing)

	return out0, err

}

// GetAllListings is a free data retrieval call binding the contract method 0xae73ccec.
//
// Solidity: function getAllListings() view returns((uint256,address,uint256,uint256,bool)[])
func (_Sale *SaleSession) GetAllListings() ([]SaleListing, error) {
	return _Sale.Contract.GetAllListings(&_Sale.CallOpts)
}

// GetAllListings is a free data retrieval call binding the contract method 0xae73ccec.
//
// Solidity: function getAllListings() view returns((uint256,address,uint256,uint256,bool)[])
func (_Sale *SaleCallerSession) GetAllListings() ([]SaleListing, error) {
	return _Sale.Contract.GetAllListings(&_Sale.CallOpts)
}

// GetListingsByUser is a free data retrieval call binding the contract method 0x69df8327.
//
// Solidity: function getListingsByUser(address user) view returns((uint256,address,uint256,uint256,bool)[])
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
// Solidity: function getListingsByUser(address user) view returns((uint256,address,uint256,uint256,bool)[])
func (_Sale *SaleSession) GetListingsByUser(user common.Address) ([]SaleListing, error) {
	return _Sale.Contract.GetListingsByUser(&_Sale.CallOpts, user)
}

// GetListingsByUser is a free data retrieval call binding the contract method 0x69df8327.
//
// Solidity: function getListingsByUser(address user) view returns((uint256,address,uint256,uint256,bool)[])
func (_Sale *SaleCallerSession) GetListingsByUser(user common.Address) ([]SaleListing, error) {
	return _Sale.Contract.GetListingsByUser(&_Sale.CallOpts, user)
}

// GetOwnedTokens is a free data retrieval call binding the contract method 0xd9d61655.
//
// Solidity: function getOwnedTokens(address owner) view returns(uint256[])
func (_Sale *SaleCaller) GetOwnedTokens(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "getOwnedTokens", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOwnedTokens is a free data retrieval call binding the contract method 0xd9d61655.
//
// Solidity: function getOwnedTokens(address owner) view returns(uint256[])
func (_Sale *SaleSession) GetOwnedTokens(owner common.Address) ([]*big.Int, error) {
	return _Sale.Contract.GetOwnedTokens(&_Sale.CallOpts, owner)
}

// GetOwnedTokens is a free data retrieval call binding the contract method 0xd9d61655.
//
// Solidity: function getOwnedTokens(address owner) view returns(uint256[])
func (_Sale *SaleCallerSession) GetOwnedTokens(owner common.Address) ([]*big.Int, error) {
	return _Sale.Contract.GetOwnedTokens(&_Sale.CallOpts, owner)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 tokenId, address seller, uint256 price, uint256 amount, bool isForSale)
func (_Sale *SaleCaller) Listings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenId   *big.Int
	Seller    common.Address
	Price     *big.Int
	Amount    *big.Int
	IsForSale bool
}, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "listings", arg0)

	outstruct := new(struct {
		TokenId   *big.Int
		Seller    common.Address
		Price     *big.Int
		Amount    *big.Int
		IsForSale bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsForSale = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 tokenId, address seller, uint256 price, uint256 amount, bool isForSale)
func (_Sale *SaleSession) Listings(arg0 *big.Int) (struct {
	TokenId   *big.Int
	Seller    common.Address
	Price     *big.Int
	Amount    *big.Int
	IsForSale bool
}, error) {
	return _Sale.Contract.Listings(&_Sale.CallOpts, arg0)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 tokenId, address seller, uint256 price, uint256 amount, bool isForSale)
func (_Sale *SaleCallerSession) Listings(arg0 *big.Int) (struct {
	TokenId   *big.Int
	Seller    common.Address
	Price     *big.Int
	Amount    *big.Int
	IsForSale bool
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

// TokenIds is a free data retrieval call binding the contract method 0xd58778d6.
//
// Solidity: function tokenIds(uint256 ) view returns(uint256)
func (_Sale *SaleCaller) TokenIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "tokenIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIds is a free data retrieval call binding the contract method 0xd58778d6.
//
// Solidity: function tokenIds(uint256 ) view returns(uint256)
func (_Sale *SaleSession) TokenIds(arg0 *big.Int) (*big.Int, error) {
	return _Sale.Contract.TokenIds(&_Sale.CallOpts, arg0)
}

// TokenIds is a free data retrieval call binding the contract method 0xd58778d6.
//
// Solidity: function tokenIds(uint256 ) view returns(uint256)
func (_Sale *SaleCallerSession) TokenIds(arg0 *big.Int) (*big.Int, error) {
	return _Sale.Contract.TokenIds(&_Sale.CallOpts, arg0)
}

// BuyToken is a paid mutator transaction binding the contract method 0x2d296bf1.
//
// Solidity: function buyToken(uint256 tokenId) payable returns()
func (_Sale *SaleTransactor) BuyToken(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "buyToken", tokenId)
}

// BuyToken is a paid mutator transaction binding the contract method 0x2d296bf1.
//
// Solidity: function buyToken(uint256 tokenId) payable returns()
func (_Sale *SaleSession) BuyToken(tokenId *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.BuyToken(&_Sale.TransactOpts, tokenId)
}

// BuyToken is a paid mutator transaction binding the contract method 0x2d296bf1.
//
// Solidity: function buyToken(uint256 tokenId) payable returns()
func (_Sale *SaleTransactorSession) BuyToken(tokenId *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.BuyToken(&_Sale.TransactOpts, tokenId)
}

// ListToken is a paid mutator transaction binding the contract method 0xbce64a7d.
//
// Solidity: function listToken(uint256 tokenId, uint256 priceInEther, uint256 amount) returns()
func (_Sale *SaleTransactor) ListToken(opts *bind.TransactOpts, tokenId *big.Int, priceInEther *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "listToken", tokenId, priceInEther, amount)
}

// ListToken is a paid mutator transaction binding the contract method 0xbce64a7d.
//
// Solidity: function listToken(uint256 tokenId, uint256 priceInEther, uint256 amount) returns()
func (_Sale *SaleSession) ListToken(tokenId *big.Int, priceInEther *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.ListToken(&_Sale.TransactOpts, tokenId, priceInEther, amount)
}

// ListToken is a paid mutator transaction binding the contract method 0xbce64a7d.
//
// Solidity: function listToken(uint256 tokenId, uint256 priceInEther, uint256 amount) returns()
func (_Sale *SaleTransactorSession) ListToken(tokenId *big.Int, priceInEther *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.ListToken(&_Sale.TransactOpts, tokenId, priceInEther, amount)
}

// RemoveListing is a paid mutator transaction binding the contract method 0x479ad4c3.
//
// Solidity: function removeListing(uint256 tokenId) returns()
func (_Sale *SaleTransactor) RemoveListing(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "removeListing", tokenId)
}

// RemoveListing is a paid mutator transaction binding the contract method 0x479ad4c3.
//
// Solidity: function removeListing(uint256 tokenId) returns()
func (_Sale *SaleSession) RemoveListing(tokenId *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.RemoveListing(&_Sale.TransactOpts, tokenId)
}

// RemoveListing is a paid mutator transaction binding the contract method 0x479ad4c3.
//
// Solidity: function removeListing(uint256 tokenId) returns()
func (_Sale *SaleTransactorSession) RemoveListing(tokenId *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.RemoveListing(&_Sale.TransactOpts, tokenId)
}
