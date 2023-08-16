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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"_baseContract\",\"type\":\"address\"},{\"internalType\":\"contractMetadata\",\"name\":\"_metadataContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"baseContract\",\"outputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyAmount\",\"type\":\"uint256\"}],\"name\":\"buyToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentListingId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllListings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orignalAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForSale\",\"type\":\"bool\"}],\"internalType\":\"structSale.Listing[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getListingsByUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orignalAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForSale\",\"type\":\"bool\"}],\"internalType\":\"structSale.Listing[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getOwnedTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"listToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orignalAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForSale\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataContract\",\"outputs\":[{\"internalType\":\"contractMetadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"removeListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260016000553480156200001657600080fd5b50604051620020d1380380620020d183398181016040528101906200003c919062000189565b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050620001d0565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000f882620000cb565b9050919050565b60006200010c82620000eb565b9050919050565b6200011e81620000ff565b81146200012a57600080fd5b50565b6000815190506200013e8162000113565b92915050565b60006200015182620000eb565b9050919050565b620001638162000144565b81146200016f57600080fd5b50565b600081519050620001838162000158565b92915050565b60008060408385031215620001a357620001a2620000c6565b5b6000620001b3858286016200012d565b9250506020620001c68582860162000172565b9150509250929050565b611ef180620001e06000396000f3fe6080604052600436106100915760003560e01c806369df83271161005957806369df83271461015c578063ae73ccec14610199578063bce64a7d146101c4578063d9d61655146101ed578063de74e57b1461022a57610091565b8063057466ea146100965780631d6a45a7146100b257806335209821146100dd578063479ad4c3146101085780635b32619c14610131575b600080fd5b6100b060048036038101906100ab919061126e565b61026d565b005b3480156100be57600080fd5b506100c7610733565b6040516100d491906112bd565b60405180910390f35b3480156100e957600080fd5b506100f2610739565b6040516100ff9190611357565b60405180910390f35b34801561011457600080fd5b5061012f600480360381019061012a9190611372565b61075f565b005b34801561013d57600080fd5b50610146610835565b60405161015391906113c0565b60405180910390f35b34801561016857600080fd5b50610183600480360381019061017e9190611419565b61085b565b60405161019091906115bc565b60405180910390f35b3480156101a557600080fd5b506101ae610b2d565b6040516101bb91906115bc565b60405180910390f35b3480156101d057600080fd5b506101eb60048036038101906101e691906115de565b610d23565b005b3480156101f957600080fd5b50610214600480360381019061020f9190611419565b610f07565b60405161022191906116e0565b60405180910390f35b34801561023657600080fd5b50610251600480360381019061024c9190611372565b61116f565b6040516102649796959493929190611720565b60405180910390f35b6002600083815260200190815260200160002060060160009054906101000a900460ff166102d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102c7906117ec565b60405180910390fd5b6000600260008481526020019081526020016000206040518060e001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820154815260200160048201548152602001600582015481526020016006820160009054906101000a900460ff161515151581525050905060008282606001516103a7919061183b565b9050803410156103ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e3906118ef565b60405180910390fd5b82600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1662fdd58e846040015185602001516040518363ffffffff1660e01b815260040161045192919061190f565b602060405180830381865afa15801561046e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610492919061194d565b10156104d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ca906119ec565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f242432a8360400151338560200151876040518563ffffffff1660e01b815260040161053c9493929190611a43565b600060405180830381600087803b15801561055657600080fd5b505af115801561056a573d6000803e3d6000fd5b50505050600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663351b38938360200151336040518363ffffffff1660e01b81526004016105cf929190611a9b565b600060405180830381600087803b1580156105e957600080fd5b505af11580156105fd573d6000803e3d6000fd5b505050506000826040015173ffffffffffffffffffffffffffffffffffffffff168260405161062b90611af2565b60006040518083038185875af1925050503d8060008114610668576040519150601f19603f3d011682016040523d82523d6000602084013e61066d565b606091505b50509050806106b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106a890611b79565b60405180910390fd5b836002600087815260200190815260200160002060040160008282546106d79190611b99565b92505081905550600060026000878152602001908152602001600020600401540361072c5760006002600087815260200190815260200160002060060160006101000a81548160ff0219169083151502179055505b5050505050565b60005481565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b3373ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610803576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107fa90611c3f565b60405180910390fd5b60006002600083815260200190815260200160002060060160006101000a81548160ff02191690831515021790555050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606000805b600054811015610923578373ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161480156108fc57506002600082815260200190815260200160002060060160009054906101000a900460ff165b1561091057818061090c90611c5f565b9250505b808061091b90611c5f565b915050610861565b5060008167ffffffffffffffff8111156109405761093f611ca7565b5b60405190808252806020026020018201604052801561097957816020015b6109666111de565b81526020019060019003908161095e5790505b5090506000805b600054811015610b21578573ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16148015610a1b57506002600082815260200190815260200160002060060160009054906101000a900460ff165b15610b0e57600260008281526020019081526020016000206040518060e001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820154815260200160048201548152602001600582015481526020016006820160009054906101000a900460ff161515151581525050838381518110610af457610af3611cd6565b5b60200260200101819052508180610b0a90611c5f565b9250505b8080610b1990611c5f565b915050610980565b50819350505050919050565b60606000805b600054811015610b88576002600082815260200190815260200160002060060160009054906101000a900460ff1615610b75578180610b7190611c5f565b9250505b8080610b8090611c5f565b915050610b33565b5060008167ffffffffffffffff811115610ba557610ba4611ca7565b5b604051908082528060200260200182016040528015610bde57816020015b610bcb6111de565b815260200190600190039081610bc35790505b5090506000805b600054811015610d19576002600082815260200190815260200160002060060160009054906101000a900460ff1615610d0657600260008281526020019081526020016000206040518060e001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820154815260200160048201548152602001600582015481526020016006820160009054906101000a900460ff161515151581525050838381518110610cec57610ceb611cd6565b5b60200260200101819052508180610d0290611c5f565b9250505b8080610d1190611c5f565b915050610be5565b5081935050505090565b60008311610d66576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d5d90611d77565b60405180910390fd5b60008211610da9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610da090611e09565b60405180910390fd5b60008111610dec576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610de390611e9b565b60405180910390fd5b6000806000815480929190610e0090611c5f565b9190505590506040518060e001604052808281526020018581526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018381526020018381526020016001151581525060026000838152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301556080820151816004015560a0820151816005015560c08201518160060160006101000a81548160ff02191690831515021790555090505050505050565b60606000805b600054811015610ff4576000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1662fdd58e8660026000868152602001908152602001600020600101546040518363ffffffff1660e01b8152600401610f8b92919061190f565b602060405180830381865afa158015610fa8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fcc919061194d565b1115610fe1578180610fdd90611c5f565b9250505b8080610fec90611c5f565b915050610f0d565b5060008167ffffffffffffffff81111561101157611010611ca7565b5b60405190808252806020026020018201604052801561103f5781602001602082028036833780820191505090505b5090506000805b600054811015611163576000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1662fdd58e8860026000868152602001908152602001600020600101546040518363ffffffff1660e01b81526004016110c492919061190f565b602060405180830381865afa1580156110e1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611105919061194d565b111561115057600260008281526020019081526020016000206001015483838151811061113557611134611cd6565b5b602002602001018181525050818061114c90611c5f565b9250505b808061115b90611c5f565b915050611046565b50819350505050919050565b60026020528060005260406000206000915090508060000154908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154908060040154908060050154908060060160009054906101000a900460ff16905087565b6040518060e001604052806000815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081526020016000151581525090565b600080fd5b6000819050919050565b61124b81611238565b811461125657600080fd5b50565b60008135905061126881611242565b92915050565b6000806040838503121561128557611284611233565b5b600061129385828601611259565b92505060206112a485828601611259565b9150509250929050565b6112b781611238565b82525050565b60006020820190506112d260008301846112ae565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061131d611318611313846112d8565b6112f8565b6112d8565b9050919050565b600061132f82611302565b9050919050565b600061134182611324565b9050919050565b61135181611336565b82525050565b600060208201905061136c6000830184611348565b92915050565b60006020828403121561138857611387611233565b5b600061139684828501611259565b91505092915050565b60006113aa82611324565b9050919050565b6113ba8161139f565b82525050565b60006020820190506113d560008301846113b1565b92915050565b60006113e6826112d8565b9050919050565b6113f6816113db565b811461140157600080fd5b50565b600081359050611413816113ed565b92915050565b60006020828403121561142f5761142e611233565b5b600061143d84828501611404565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61147b81611238565b82525050565b61148a816113db565b82525050565b60008115159050919050565b6114a581611490565b82525050565b60e0820160008201516114c16000850182611472565b5060208201516114d46020850182611472565b5060408201516114e76040850182611481565b5060608201516114fa6060850182611472565b50608082015161150d6080850182611472565b5060a082015161152060a0850182611472565b5060c082015161153360c085018261149c565b50505050565b600061154583836114ab565b60e08301905092915050565b6000602082019050919050565b600061156982611446565b6115738185611451565b935061157e83611462565b8060005b838110156115af5781516115968882611539565b97506115a183611551565b925050600181019050611582565b5085935050505092915050565b600060208201905081810360008301526115d6818461155e565b905092915050565b6000806000606084860312156115f7576115f6611233565b5b600061160586828701611259565b935050602061161686828701611259565b925050604061162786828701611259565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60006116698383611472565b60208301905092915050565b6000602082019050919050565b600061168d82611631565b611697818561163c565b93506116a28361164d565b8060005b838110156116d35781516116ba888261165d565b97506116c583611675565b9250506001810190506116a6565b5085935050505092915050565b600060208201905081810360008301526116fa8184611682565b905092915050565b61170b816113db565b82525050565b61171a81611490565b82525050565b600060e082019050611735600083018a6112ae565b61174260208301896112ae565b61174f6040830188611702565b61175c60608301876112ae565b61176960808301866112ae565b61177660a08301856112ae565b61178360c0830184611711565b98975050505050505050565b600082825260208201905092915050565b7f54686973206974656d206973206e6f7420666f722073616c6500000000000000600082015250565b60006117d660198361178f565b91506117e1826117a0565b602082019050919050565b60006020820190508181036000830152611805816117c9565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061184682611238565b915061185183611238565b925082820261185f81611238565b915082820484148315176118765761187561180c565b5b5092915050565b7f53656e742076616c7565206973206c657373207468616e2074686520746f746160008201527f6c20636f73740000000000000000000000000000000000000000000000000000602082015250565b60006118d960268361178f565b91506118e48261187d565b604082019050919050565b60006020820190508181036000830152611908816118cc565b9050919050565b60006040820190506119246000830185611702565b61193160208301846112ae565b9392505050565b60008151905061194781611242565b92915050565b60006020828403121561196357611962611233565b5b600061197184828501611938565b91505092915050565b7f53656c6c657220646f6573206e6f74206861766520656e6f75676820746f6b6560008201527f6e7320666f722073616c65000000000000000000000000000000000000000000602082015250565b60006119d6602b8361178f565b91506119e18261197a565b604082019050919050565b60006020820190508181036000830152611a05816119c9565b9050919050565b600082825260208201905092915050565b50565b6000611a2d600083611a0c565b9150611a3882611a1d565b600082019050919050565b600060a082019050611a586000830187611702565b611a656020830186611702565b611a7260408301856112ae565b611a7f60608301846112ae565b8181036080830152611a9081611a20565b905095945050505050565b6000604082019050611ab060008301856112ae565b611abd6020830184611702565b9392505050565b600081905092915050565b6000611adc600083611ac4565b9150611ae782611a1d565b600082019050919050565b6000611afd82611acf565b9150819050919050565b7f4661696c656420746f207472616e7366657220457468657220746f207468652060008201527f73656c6c65720000000000000000000000000000000000000000000000000000602082015250565b6000611b6360268361178f565b9150611b6e82611b07565b604082019050919050565b60006020820190508181036000830152611b9281611b56565b9050919050565b6000611ba482611238565b9150611baf83611238565b9250828203905081811115611bc757611bc661180c565b5b92915050565b7f4f6e6c79207468652073656c6c65722063616e2072656d6f766520746865206c60008201527f697374696e670000000000000000000000000000000000000000000000000000602082015250565b6000611c2960268361178f565b9150611c3482611bcd565b604082019050919050565b60006020820190508181036000830152611c5881611c1c565b9050919050565b6000611c6a82611238565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c9c57611c9b61180c565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f546f6b656e2049442073686f756c642062652067726561746572207468616e2060008201527f7a65726f00000000000000000000000000000000000000000000000000000000602082015250565b6000611d6160248361178f565b9150611d6c82611d05565b604082019050919050565b60006020820190508181036000830152611d9081611d54565b9050919050565b7f50726963652073686f756c642062652067726561746572207468616e207a657260008201527f6f00000000000000000000000000000000000000000000000000000000000000602082015250565b6000611df360218361178f565b9150611dfe82611d97565b604082019050919050565b60006020820190508181036000830152611e2281611de6565b9050919050565b7f416d6f756e742073686f756c642062652067726561746572207468616e207a6560008201527f726f000000000000000000000000000000000000000000000000000000000000602082015250565b6000611e8560228361178f565b9150611e9082611e29565b604082019050919050565b60006020820190508181036000830152611eb481611e78565b905091905056fea2646970667358221220e27ff2cb57ff49d12bbc592c758476a527dda32f86d76643e049c5ebfda2e7c564736f6c63430008120033",
}

// SaleABI is the input ABI used to generate the binding from.
// Deprecated: Use SaleMetaData.ABI instead.
var SaleABI = SaleMetaData.ABI

// SaleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SaleMetaData.Bin instead.
var SaleBin = SaleMetaData.Bin

// DeploySale deploys a new Ethereum contract, binding an instance of Sale to it.
func DeploySale(auth *bind.TransactOpts, backend bind.ContractBackend, _baseContract common.Address, _metadataContract common.Address) (common.Address, *types.Transaction, *Sale, error) {
	parsed, err := SaleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SaleBin), backend, _baseContract, _metadataContract)
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

// MetadataContract is a free data retrieval call binding the contract method 0x35209821.
//
// Solidity: function metadataContract() view returns(address)
func (_Sale *SaleCaller) MetadataContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "metadataContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MetadataContract is a free data retrieval call binding the contract method 0x35209821.
//
// Solidity: function metadataContract() view returns(address)
func (_Sale *SaleSession) MetadataContract() (common.Address, error) {
	return _Sale.Contract.MetadataContract(&_Sale.CallOpts)
}

// MetadataContract is a free data retrieval call binding the contract method 0x35209821.
//
// Solidity: function metadataContract() view returns(address)
func (_Sale *SaleCallerSession) MetadataContract() (common.Address, error) {
	return _Sale.Contract.MetadataContract(&_Sale.CallOpts)
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
