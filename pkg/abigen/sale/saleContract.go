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
	ListingId     *big.Int
	TokenId       *big.Int
	Seller        common.Address
	Price         *big.Int
	Amount        *big.Int
	OrignalAmount *big.Int
	IsForSale     bool
}

// SaleMetaData contains all meta data concerning the Sale contract.
var SaleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"_baseContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"baseContract\",\"outputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyAmount\",\"type\":\"uint256\"}],\"name\":\"buyToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentListingId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllListings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orignalAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForSale\",\"type\":\"bool\"}],\"internalType\":\"structSale.Listing[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getListingsByUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orignalAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForSale\",\"type\":\"bool\"}],\"internalType\":\"structSale.Listing[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getOwnedTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"listToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orignalAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForSale\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"removeListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260016000553480156200001657600080fd5b5060405162001ee138038062001ee183398181016040528101906200003c919062000102565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000134565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000b68262000089565b9050919050565b6000620000ca82620000a9565b9050919050565b620000dc81620000bd565b8114620000e857600080fd5b50565b600081519050620000fc81620000d1565b92915050565b6000602082840312156200011b576200011a62000084565b5b60006200012b84828501620000eb565b91505092915050565b611d9d80620001446000396000f3fe6080604052600436106100865760003560e01c806369df83271161005957806369df832714610126578063ae73ccec14610163578063bce64a7d1461018e578063d9d61655146101b7578063de74e57b146101f457610086565b8063057466ea1461008b5780631d6a45a7146100a7578063479ad4c3146100d25780635b32619c146100fb575b600080fd5b6100a560048036038101906100a0919061117f565b610237565b005b3480156100b357600080fd5b506100bc61066a565b6040516100c991906111ce565b60405180910390f35b3480156100de57600080fd5b506100f960048036038101906100f491906111e9565b610670565b005b34801561010757600080fd5b50610110610746565b60405161011d9190611295565b60405180910390f35b34801561013257600080fd5b5061014d600480360381019061014891906112ee565b61076c565b60405161015a9190611491565b60405180910390f35b34801561016f57600080fd5b50610178610a3e565b6040516101859190611491565b60405180910390f35b34801561019a57600080fd5b506101b560048036038101906101b091906114b3565b610c34565b005b3480156101c357600080fd5b506101de60048036038101906101d991906112ee565b610e18565b6040516101eb91906115b5565b60405180910390f35b34801561020057600080fd5b5061021b600480360381019061021691906111e9565b611080565b60405161022e97969594939291906115f5565b60405180910390f35b6002600083815260200190815260200160002060060160009054906101000a900460ff1661029a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610291906116c1565b60405180910390fd5b6000600260008481526020019081526020016000206040518060e001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820154815260200160048201548152602001600582015481526020016006820160009054906101000a900460ff161515151581525050905060008282606001516103719190611710565b9050803410156103b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ad906117c4565b60405180910390fd5b82600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1662fdd58e846040015185602001516040518363ffffffff1660e01b815260040161041b9291906117e4565b602060405180830381865afa158015610438573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061045c9190611822565b101561049d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610494906118c1565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f242432a8360400151338560200151876040518563ffffffff1660e01b81526004016105069493929190611918565b600060405180830381600087803b15801561052057600080fd5b505af1158015610534573d6000803e3d6000fd5b505050506000826040015173ffffffffffffffffffffffffffffffffffffffff16826040516105629061199e565b60006040518083038185875af1925050503d806000811461059f576040519150601f19603f3d011682016040523d82523d6000602084013e6105a4565b606091505b50509050806105e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105df90611a25565b60405180910390fd5b8360026000878152602001908152602001600020600401600082825461060e9190611a45565b9250508190555060006002600087815260200190815260200160002060040154036106635760006002600087815260200190815260200160002060060160006101000a81548160ff0219169083151502179055505b5050505050565b60005481565b3373ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610714576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070b90611aeb565b60405180910390fd5b60006002600083815260200190815260200160002060060160006101000a81548160ff02191690831515021790555050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606000805b600054811015610834578373ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614801561080d57506002600082815260200190815260200160002060060160009054906101000a900460ff165b1561082157818061081d90611b0b565b9250505b808061082c90611b0b565b915050610772565b5060008167ffffffffffffffff81111561085157610850611b53565b5b60405190808252806020026020018201604052801561088a57816020015b6108776110ef565b81526020019060019003908161086f5790505b5090506000805b600054811015610a32578573ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614801561092c57506002600082815260200190815260200160002060060160009054906101000a900460ff165b15610a1f57600260008281526020019081526020016000206040518060e001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820154815260200160048201548152602001600582015481526020016006820160009054906101000a900460ff161515151581525050838381518110610a0557610a04611b82565b5b60200260200101819052508180610a1b90611b0b565b9250505b8080610a2a90611b0b565b915050610891565b50819350505050919050565b60606000805b600054811015610a99576002600082815260200190815260200160002060060160009054906101000a900460ff1615610a86578180610a8290611b0b565b9250505b8080610a9190611b0b565b915050610a44565b5060008167ffffffffffffffff811115610ab657610ab5611b53565b5b604051908082528060200260200182016040528015610aef57816020015b610adc6110ef565b815260200190600190039081610ad45790505b5090506000805b600054811015610c2a576002600082815260200190815260200160002060060160009054906101000a900460ff1615610c1757600260008281526020019081526020016000206040518060e001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820154815260200160048201548152602001600582015481526020016006820160009054906101000a900460ff161515151581525050838381518110610bfd57610bfc611b82565b5b60200260200101819052508180610c1390611b0b565b9250505b8080610c2290611b0b565b915050610af6565b5081935050505090565b60008311610c77576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6e90611c23565b60405180910390fd5b60008211610cba576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cb190611cb5565b60405180910390fd5b60008111610cfd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cf490611d47565b60405180910390fd5b6000806000815480929190610d1190611b0b565b9190505590506040518060e001604052808281526020018581526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018381526020018381526020016001151581525060026000838152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301556080820151816004015560a0820151816005015560c08201518160060160006101000a81548160ff02191690831515021790555090505050505050565b60606000805b600054811015610f05576000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1662fdd58e8660026000868152602001908152602001600020600101546040518363ffffffff1660e01b8152600401610e9c9291906117e4565b602060405180830381865afa158015610eb9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610edd9190611822565b1115610ef2578180610eee90611b0b565b9250505b8080610efd90611b0b565b915050610e1e565b5060008167ffffffffffffffff811115610f2257610f21611b53565b5b604051908082528060200260200182016040528015610f505781602001602082028036833780820191505090505b5090506000805b600054811015611074576000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1662fdd58e8860026000868152602001908152602001600020600101546040518363ffffffff1660e01b8152600401610fd59291906117e4565b602060405180830381865afa158015610ff2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110169190611822565b111561106157600260008281526020019081526020016000206001015483838151811061104657611045611b82565b5b602002602001018181525050818061105d90611b0b565b9250505b808061106c90611b0b565b915050610f57565b50819350505050919050565b60026020528060005260406000206000915090508060000154908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154908060040154908060050154908060060160009054906101000a900460ff16905087565b6040518060e001604052806000815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081526020016000151581525090565b600080fd5b6000819050919050565b61115c81611149565b811461116757600080fd5b50565b60008135905061117981611153565b92915050565b6000806040838503121561119657611195611144565b5b60006111a48582860161116a565b92505060206111b58582860161116a565b9150509250929050565b6111c881611149565b82525050565b60006020820190506111e360008301846111bf565b92915050565b6000602082840312156111ff576111fe611144565b5b600061120d8482850161116a565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061125b61125661125184611216565b611236565b611216565b9050919050565b600061126d82611240565b9050919050565b600061127f82611262565b9050919050565b61128f81611274565b82525050565b60006020820190506112aa6000830184611286565b92915050565b60006112bb82611216565b9050919050565b6112cb816112b0565b81146112d657600080fd5b50565b6000813590506112e8816112c2565b92915050565b60006020828403121561130457611303611144565b5b6000611312848285016112d9565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61135081611149565b82525050565b61135f816112b0565b82525050565b60008115159050919050565b61137a81611365565b82525050565b60e0820160008201516113966000850182611347565b5060208201516113a96020850182611347565b5060408201516113bc6040850182611356565b5060608201516113cf6060850182611347565b5060808201516113e26080850182611347565b5060a08201516113f560a0850182611347565b5060c082015161140860c0850182611371565b50505050565b600061141a8383611380565b60e08301905092915050565b6000602082019050919050565b600061143e8261131b565b6114488185611326565b935061145383611337565b8060005b8381101561148457815161146b888261140e565b975061147683611426565b925050600181019050611457565b5085935050505092915050565b600060208201905081810360008301526114ab8184611433565b905092915050565b6000806000606084860312156114cc576114cb611144565b5b60006114da8682870161116a565b93505060206114eb8682870161116a565b92505060406114fc8682870161116a565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600061153e8383611347565b60208301905092915050565b6000602082019050919050565b600061156282611506565b61156c8185611511565b935061157783611522565b8060005b838110156115a857815161158f8882611532565b975061159a8361154a565b92505060018101905061157b565b5085935050505092915050565b600060208201905081810360008301526115cf8184611557565b905092915050565b6115e0816112b0565b82525050565b6115ef81611365565b82525050565b600060e08201905061160a600083018a6111bf565b61161760208301896111bf565b61162460408301886115d7565b61163160608301876111bf565b61163e60808301866111bf565b61164b60a08301856111bf565b61165860c08301846115e6565b98975050505050505050565b600082825260208201905092915050565b7f54686973206974656d206973206e6f7420666f722073616c6500000000000000600082015250565b60006116ab601983611664565b91506116b682611675565b602082019050919050565b600060208201905081810360008301526116da8161169e565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061171b82611149565b915061172683611149565b925082820261173481611149565b9150828204841483151761174b5761174a6116e1565b5b5092915050565b7f53656e742076616c7565206973206c657373207468616e2074686520746f746160008201527f6c20636f73740000000000000000000000000000000000000000000000000000602082015250565b60006117ae602683611664565b91506117b982611752565b604082019050919050565b600060208201905081810360008301526117dd816117a1565b9050919050565b60006040820190506117f960008301856115d7565b61180660208301846111bf565b9392505050565b60008151905061181c81611153565b92915050565b60006020828403121561183857611837611144565b5b60006118468482850161180d565b91505092915050565b7f53656c6c657220646f6573206e6f74206861766520656e6f75676820746f6b6560008201527f6e7320666f722073616c65000000000000000000000000000000000000000000602082015250565b60006118ab602b83611664565b91506118b68261184f565b604082019050919050565b600060208201905081810360008301526118da8161189e565b9050919050565b600082825260208201905092915050565b50565b60006119026000836118e1565b915061190d826118f2565b600082019050919050565b600060a08201905061192d60008301876115d7565b61193a60208301866115d7565b61194760408301856111bf565b61195460608301846111bf565b8181036080830152611965816118f5565b905095945050505050565b600081905092915050565b6000611988600083611970565b9150611993826118f2565b600082019050919050565b60006119a98261197b565b9150819050919050565b7f4661696c656420746f207472616e7366657220457468657220746f207468652060008201527f73656c6c65720000000000000000000000000000000000000000000000000000602082015250565b6000611a0f602683611664565b9150611a1a826119b3565b604082019050919050565b60006020820190508181036000830152611a3e81611a02565b9050919050565b6000611a5082611149565b9150611a5b83611149565b9250828203905081811115611a7357611a726116e1565b5b92915050565b7f4f6e6c79207468652073656c6c65722063616e2072656d6f766520746865206c60008201527f697374696e670000000000000000000000000000000000000000000000000000602082015250565b6000611ad5602683611664565b9150611ae082611a79565b604082019050919050565b60006020820190508181036000830152611b0481611ac8565b9050919050565b6000611b1682611149565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611b4857611b476116e1565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f546f6b656e2049442073686f756c642062652067726561746572207468616e2060008201527f7a65726f00000000000000000000000000000000000000000000000000000000602082015250565b6000611c0d602483611664565b9150611c1882611bb1565b604082019050919050565b60006020820190508181036000830152611c3c81611c00565b9050919050565b7f50726963652073686f756c642062652067726561746572207468616e207a657260008201527f6f00000000000000000000000000000000000000000000000000000000000000602082015250565b6000611c9f602183611664565b9150611caa82611c43565b604082019050919050565b60006020820190508181036000830152611cce81611c92565b9050919050565b7f416d6f756e742073686f756c642062652067726561746572207468616e207a6560008201527f726f000000000000000000000000000000000000000000000000000000000000602082015250565b6000611d31602283611664565b9150611d3c82611cd5565b604082019050919050565b60006020820190508181036000830152611d6081611d24565b905091905056fea2646970667358221220a94fb957b359defbb86d0919c68f496ad17b9485588d29ee75f419769737179364736f6c63430008120033",
}

// SaleABI is the input ABI used to generate the binding from.
// Deprecated: Use SaleMetaData.ABI instead.
var SaleABI = SaleMetaData.ABI

// SaleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SaleMetaData.Bin instead.
var SaleBin = SaleMetaData.Bin

// DeploySale deploys a new Ethereum contract, binding an instance of Sale to it.
func DeploySale(auth *bind.TransactOpts, backend bind.ContractBackend, _baseContract common.Address) (common.Address, *types.Transaction, *Sale, error) {
	parsed, err := SaleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SaleBin), backend, _baseContract)
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

// BaseContract is a free data retrieval call binding the contract method 0x5b32619c.
//
// Solidity: function baseContract() view returns(address)
func (_Sale *SaleCaller) BaseContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "baseContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseContract is a free data retrieval call binding the contract method 0x5b32619c.
//
// Solidity: function baseContract() view returns(address)
func (_Sale *SaleSession) BaseContract() (common.Address, error) {
	return _Sale.Contract.BaseContract(&_Sale.CallOpts)
}

// BaseContract is a free data retrieval call binding the contract method 0x5b32619c.
//
// Solidity: function baseContract() view returns(address)
func (_Sale *SaleCallerSession) BaseContract() (common.Address, error) {
	return _Sale.Contract.BaseContract(&_Sale.CallOpts)
}

// CurrentListingId is a free data retrieval call binding the contract method 0x1d6a45a7.
//
// Solidity: function currentListingId() view returns(uint256)
func (_Sale *SaleCaller) CurrentListingId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "currentListingId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentListingId is a free data retrieval call binding the contract method 0x1d6a45a7.
//
// Solidity: function currentListingId() view returns(uint256)
func (_Sale *SaleSession) CurrentListingId() (*big.Int, error) {
	return _Sale.Contract.CurrentListingId(&_Sale.CallOpts)
}

// CurrentListingId is a free data retrieval call binding the contract method 0x1d6a45a7.
//
// Solidity: function currentListingId() view returns(uint256)
func (_Sale *SaleCallerSession) CurrentListingId() (*big.Int, error) {
	return _Sale.Contract.CurrentListingId(&_Sale.CallOpts)
}

// GetAllListings is a free data retrieval call binding the contract method 0xae73ccec.
//
// Solidity: function getAllListings() view returns((uint256,uint256,address,uint256,uint256,uint256,bool)[])
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
// Solidity: function getAllListings() view returns((uint256,uint256,address,uint256,uint256,uint256,bool)[])
func (_Sale *SaleSession) GetAllListings() ([]SaleListing, error) {
	return _Sale.Contract.GetAllListings(&_Sale.CallOpts)
}

// GetAllListings is a free data retrieval call binding the contract method 0xae73ccec.
//
// Solidity: function getAllListings() view returns((uint256,uint256,address,uint256,uint256,uint256,bool)[])
func (_Sale *SaleCallerSession) GetAllListings() ([]SaleListing, error) {
	return _Sale.Contract.GetAllListings(&_Sale.CallOpts)
}

// GetListingsByUser is a free data retrieval call binding the contract method 0x69df8327.
//
// Solidity: function getListingsByUser(address user) view returns((uint256,uint256,address,uint256,uint256,uint256,bool)[])
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
// Solidity: function getListingsByUser(address user) view returns((uint256,uint256,address,uint256,uint256,uint256,bool)[])
func (_Sale *SaleSession) GetListingsByUser(user common.Address) ([]SaleListing, error) {
	return _Sale.Contract.GetListingsByUser(&_Sale.CallOpts, user)
}

// GetListingsByUser is a free data retrieval call binding the contract method 0x69df8327.
//
// Solidity: function getListingsByUser(address user) view returns((uint256,uint256,address,uint256,uint256,uint256,bool)[])
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
// Solidity: function listings(uint256 ) view returns(uint256 listingId, uint256 tokenId, address seller, uint256 price, uint256 amount, uint256 orignalAmount, bool isForSale)
func (_Sale *SaleCaller) Listings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ListingId     *big.Int
	TokenId       *big.Int
	Seller        common.Address
	Price         *big.Int
	Amount        *big.Int
	OrignalAmount *big.Int
	IsForSale     bool
}, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "listings", arg0)

	outstruct := new(struct {
		ListingId     *big.Int
		TokenId       *big.Int
		Seller        common.Address
		Price         *big.Int
		Amount        *big.Int
		OrignalAmount *big.Int
		IsForSale     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ListingId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.OrignalAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.IsForSale = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 listingId, uint256 tokenId, address seller, uint256 price, uint256 amount, uint256 orignalAmount, bool isForSale)
func (_Sale *SaleSession) Listings(arg0 *big.Int) (struct {
	ListingId     *big.Int
	TokenId       *big.Int
	Seller        common.Address
	Price         *big.Int
	Amount        *big.Int
	OrignalAmount *big.Int
	IsForSale     bool
}, error) {
	return _Sale.Contract.Listings(&_Sale.CallOpts, arg0)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 listingId, uint256 tokenId, address seller, uint256 price, uint256 amount, uint256 orignalAmount, bool isForSale)
func (_Sale *SaleCallerSession) Listings(arg0 *big.Int) (struct {
	ListingId     *big.Int
	TokenId       *big.Int
	Seller        common.Address
	Price         *big.Int
	Amount        *big.Int
	OrignalAmount *big.Int
	IsForSale     bool
}, error) {
	return _Sale.Contract.Listings(&_Sale.CallOpts, arg0)
}

// BuyToken is a paid mutator transaction binding the contract method 0x057466ea.
//
// Solidity: function buyToken(uint256 listingId, uint256 buyAmount) payable returns()
func (_Sale *SaleTransactor) BuyToken(opts *bind.TransactOpts, listingId *big.Int, buyAmount *big.Int) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "buyToken", listingId, buyAmount)
}

// BuyToken is a paid mutator transaction binding the contract method 0x057466ea.
//
// Solidity: function buyToken(uint256 listingId, uint256 buyAmount) payable returns()
func (_Sale *SaleSession) BuyToken(listingId *big.Int, buyAmount *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.BuyToken(&_Sale.TransactOpts, listingId, buyAmount)
}

// BuyToken is a paid mutator transaction binding the contract method 0x057466ea.
//
// Solidity: function buyToken(uint256 listingId, uint256 buyAmount) payable returns()
func (_Sale *SaleTransactorSession) BuyToken(listingId *big.Int, buyAmount *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.BuyToken(&_Sale.TransactOpts, listingId, buyAmount)
}

// ListToken is a paid mutator transaction binding the contract method 0xbce64a7d.
//
// Solidity: function listToken(uint256 tokenId, uint256 priceInWei, uint256 amount) returns()
func (_Sale *SaleTransactor) ListToken(opts *bind.TransactOpts, tokenId *big.Int, priceInWei *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "listToken", tokenId, priceInWei, amount)
}

// ListToken is a paid mutator transaction binding the contract method 0xbce64a7d.
//
// Solidity: function listToken(uint256 tokenId, uint256 priceInWei, uint256 amount) returns()
func (_Sale *SaleSession) ListToken(tokenId *big.Int, priceInWei *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.ListToken(&_Sale.TransactOpts, tokenId, priceInWei, amount)
}

// ListToken is a paid mutator transaction binding the contract method 0xbce64a7d.
//
// Solidity: function listToken(uint256 tokenId, uint256 priceInWei, uint256 amount) returns()
func (_Sale *SaleTransactorSession) ListToken(tokenId *big.Int, priceInWei *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.ListToken(&_Sale.TransactOpts, tokenId, priceInWei, amount)
}

// RemoveListing is a paid mutator transaction binding the contract method 0x479ad4c3.
//
// Solidity: function removeListing(uint256 listingId) returns()
func (_Sale *SaleTransactor) RemoveListing(opts *bind.TransactOpts, listingId *big.Int) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "removeListing", listingId)
}

// RemoveListing is a paid mutator transaction binding the contract method 0x479ad4c3.
//
// Solidity: function removeListing(uint256 listingId) returns()
func (_Sale *SaleSession) RemoveListing(listingId *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.RemoveListing(&_Sale.TransactOpts, listingId)
}

// RemoveListing is a paid mutator transaction binding the contract method 0x479ad4c3.
//
// Solidity: function removeListing(uint256 listingId) returns()
func (_Sale *SaleTransactorSession) RemoveListing(listingId *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.RemoveListing(&_Sale.TransactOpts, listingId)
}
