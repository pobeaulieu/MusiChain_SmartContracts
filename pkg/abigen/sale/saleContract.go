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
	ListingId *big.Int
	TokenId   *big.Int
	Seller    common.Address
	Price     *big.Int
	Amount    *big.Int
	IsForSale bool
}

// SaleMetaData contains all meta data concerning the Sale contract.
var SaleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyAmount\",\"type\":\"uint256\"}],\"name\":\"buyToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentListingId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllListings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForSale\",\"type\":\"bool\"}],\"internalType\":\"structSale.Listing[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getListingsByUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForSale\",\"type\":\"bool\"}],\"internalType\":\"structSale.Listing[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getOwnedTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"listToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listingIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForSale\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"removeListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405260016000553480156200001657600080fd5b50604051620020333803806200203383398181016040528101906200003c919062000102565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000134565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000b68262000089565b9050919050565b6000620000ca82620000a9565b9050919050565b620000dc81620000bd565b8114620000e857600080fd5b50565b600081519050620000fc81620000d1565b92915050565b6000602082840312156200011b576200011a62000084565b5b60006200012b84828501620000eb565b91505092915050565b611eef80620001446000396000f3fe6080604052600436106100915760003560e01c806369df83271161005957806369df83271461016e578063ae73ccec146101ab578063bce64a7d146101d6578063d9d61655146101ff578063de74e57b1461023c57610091565b8063057466ea146100965780631d6a45a7146100b2578063479ad4c3146100dd57806352bea3601461010657806355a373d614610143575b600080fd5b6100b060048036038101906100ab91906112f2565b61027e565b005b3480156100be57600080fd5b506100c76106a7565b6040516100d49190611341565b60405180910390f35b3480156100e957600080fd5b5061010460048036038101906100ff919061135c565b6106ad565b005b34801561011257600080fd5b5061012d6004803603810190610128919061135c565b610783565b60405161013a9190611341565b60405180910390f35b34801561014f57600080fd5b506101586107a7565b6040516101659190611408565b60405180910390f35b34801561017a57600080fd5b5061019560048036038101906101909190611461565b6107cd565b6040516101a291906115f1565b60405180910390f35b3480156101b757600080fd5b506101c0610b31565b6040516101cd91906115f1565b60405180910390f35b3480156101e257600080fd5b506101fd60048036038101906101f89190611613565b610d7d565b005b34801561020b57600080fd5b5061022660048036038101906102219190611461565b610f7a565b6040516102339190611715565b60405180910390f35b34801561024857600080fd5b50610263600480360381019061025e919061135c565b611200565b60405161027596959493929190611755565b60405180910390f35b6003600083815260200190815260200160002060050160009054906101000a900460ff166102e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102d890611813565b60405180910390fd5b6000600360008481526020019081526020016000206040518060c001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201548152602001600482015481526020016005820160009054906101000a900460ff161515151581525050905060008282606001516103ae9190611862565b9050803410156103f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ea90611916565b60405180910390fd5b82600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1662fdd58e846040015185602001516040518363ffffffff1660e01b8152600401610458929190611936565b602060405180830381865afa158015610475573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104999190611974565b10156104da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d190611a13565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f242432a8360400151338560200151876040518563ffffffff1660e01b81526004016105439493929190611a6a565b600060405180830381600087803b15801561055d57600080fd5b505af1158015610571573d6000803e3d6000fd5b505050506000826040015173ffffffffffffffffffffffffffffffffffffffff168260405161059f90611af0565b60006040518083038185875af1925050503d80600081146105dc576040519150601f19603f3d011682016040523d82523d6000602084013e6105e1565b606091505b5050905080610625576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161061c90611b77565b60405180910390fd5b8360036000878152602001908152602001600020600401600082825461064b9190611b97565b9250508190555060006003600087815260200190815260200160002060040154036106a05760006003600087815260200190815260200160002060050160006101000a81548160ff0219169083151502179055505b5050505050565b60005481565b3373ffffffffffffffffffffffffffffffffffffffff166003600083815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610751576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161074890611c3d565b60405180910390fd5b60006003600083815260200190815260200160002060050160006101000a81548160ff02191690831515021790555050565b6002818154811061079357600080fd5b906000526020600020016000915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606000805b6002805490508110156108d4578373ffffffffffffffffffffffffffffffffffffffff16600360006002848154811061080f5761080e611c5d565b5b9060005260206000200154815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161480156108ad5750600360006002838154811061088257610881611c5d565b5b9060005260206000200154815260200190815260200160002060050160009054906101000a900460ff165b156108c15781806108bd90611c8c565b9250505b80806108cc90611c8c565b9150506107d3565b5060008167ffffffffffffffff8111156108f1576108f0611cd4565b5b60405190808252806020026020018201604052801561092a57816020015b610917611269565b81526020019060019003908161090f5790505b5090506000805b600280549050811015610b25578573ffffffffffffffffffffffffffffffffffffffff16600360006002848154811061096d5761096c611c5d565b5b9060005260206000200154815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16148015610a0b575060036000600283815481106109e0576109df611c5d565b5b9060005260206000200154815260200190815260200160002060050160009054906101000a900460ff165b15610b12576003600060028381548110610a2857610a27611c5d565b5b906000526020600020015481526020019081526020016000206040518060c001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201548152602001600482015481526020016005820160009054906101000a900460ff161515151581525050838381518110610af857610af7611c5d565b5b60200260200101819052508180610b0e90611c8c565b9250505b8080610b1d90611c8c565b915050610931565b50819350505050919050565b60606000805b600280549050811015610bad576003600060028381548110610b5c57610b5b611c5d565b5b9060005260206000200154815260200190815260200160002060050160009054906101000a900460ff1615610b9a578180610b9690611c8c565b9250505b8080610ba590611c8c565b915050610b37565b5060008167ffffffffffffffff811115610bca57610bc9611cd4565b5b604051908082528060200260200182016040528015610c0357816020015b610bf0611269565b815260200190600190039081610be85790505b5090506000805b600280549050811015610d73576003600060028381548110610c2f57610c2e611c5d565b5b9060005260206000200154815260200190815260200160002060050160009054906101000a900460ff1615610d60576003600060028381548110610c7657610c75611c5d565b5b906000526020600020015481526020019081526020016000206040518060c001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201548152602001600482015481526020016005820160009054906101000a900460ff161515151581525050838381518110610d4657610d45611c5d565b5b60200260200101819052508180610d5c90611c8c565b9250505b8080610d6b90611c8c565b915050610c0a565b5081935050505090565b60008311610dc0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610db790611d75565b60405180910390fd5b60008211610e03576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dfa90611e07565b60405180910390fd5b60008111610e46576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e3d90611e99565b60405180910390fd5b6000806000815480929190610e5a90611c8c565b9190505590506040518060c001604052808281526020018581526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018381526020016001151581525060036000838152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301556080820151816004015560a08201518160050160006101000a81548160ff021916908315150217905550905050600281908060018154018082558091505060019003906000526020600020016000909190919091505550505050565b60606000805b600280549050811015611072576000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1662fdd58e8660028581548110610fe157610fe0611c5d565b5b90600052602060002001546040518363ffffffff1660e01b8152600401611009929190611936565b602060405180830381865afa158015611026573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061104a9190611974565b111561105f57818061105b90611c8c565b9250505b808061106a90611c8c565b915050610f80565b5060008167ffffffffffffffff81111561108f5761108e611cd4565b5b6040519080825280602002602001820160405280156110bd5781602001602082028036833780820191505090505b5090506000805b6002805490508110156111f4576000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1662fdd58e886002858154811061112557611124611c5d565b5b90600052602060002001546040518363ffffffff1660e01b815260040161114d929190611936565b602060405180830381865afa15801561116a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061118e9190611974565b11156111e157600281815481106111a8576111a7611c5d565b5b90600052602060002001548383815181106111c6576111c5611c5d565b5b60200260200101818152505081806111dd90611c8c565b9250505b80806111ec90611c8c565b9150506110c4565b50819350505050919050565b60036020528060005260406000206000915090508060000154908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154908060040154908060050160009054906101000a900460ff16905086565b6040518060c001604052806000815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081526020016000151581525090565b600080fd5b6000819050919050565b6112cf816112bc565b81146112da57600080fd5b50565b6000813590506112ec816112c6565b92915050565b60008060408385031215611309576113086112b7565b5b6000611317858286016112dd565b9250506020611328858286016112dd565b9150509250929050565b61133b816112bc565b82525050565b60006020820190506113566000830184611332565b92915050565b600060208284031215611372576113716112b7565b5b6000611380848285016112dd565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006113ce6113c96113c484611389565b6113a9565b611389565b9050919050565b60006113e0826113b3565b9050919050565b60006113f2826113d5565b9050919050565b611402816113e7565b82525050565b600060208201905061141d60008301846113f9565b92915050565b600061142e82611389565b9050919050565b61143e81611423565b811461144957600080fd5b50565b60008135905061145b81611435565b92915050565b600060208284031215611477576114766112b7565b5b60006114858482850161144c565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6114c3816112bc565b82525050565b6114d281611423565b82525050565b60008115159050919050565b6114ed816114d8565b82525050565b60c08201600082015161150960008501826114ba565b50602082015161151c60208501826114ba565b50604082015161152f60408501826114c9565b50606082015161154260608501826114ba565b50608082015161155560808501826114ba565b5060a082015161156860a08501826114e4565b50505050565b600061157a83836114f3565b60c08301905092915050565b6000602082019050919050565b600061159e8261148e565b6115a88185611499565b93506115b3836114aa565b8060005b838110156115e45781516115cb888261156e565b97506115d683611586565b9250506001810190506115b7565b5085935050505092915050565b6000602082019050818103600083015261160b8184611593565b905092915050565b60008060006060848603121561162c5761162b6112b7565b5b600061163a868287016112dd565b935050602061164b868287016112dd565b925050604061165c868287016112dd565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600061169e83836114ba565b60208301905092915050565b6000602082019050919050565b60006116c282611666565b6116cc8185611671565b93506116d783611682565b8060005b838110156117085781516116ef8882611692565b97506116fa836116aa565b9250506001810190506116db565b5085935050505092915050565b6000602082019050818103600083015261172f81846116b7565b905092915050565b61174081611423565b82525050565b61174f816114d8565b82525050565b600060c08201905061176a6000830189611332565b6117776020830188611332565b6117846040830187611737565b6117916060830186611332565b61179e6080830185611332565b6117ab60a0830184611746565b979650505050505050565b600082825260208201905092915050565b7f54686973206974656d206973206e6f7420666f722073616c6500000000000000600082015250565b60006117fd6019836117b6565b9150611808826117c7565b602082019050919050565b6000602082019050818103600083015261182c816117f0565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061186d826112bc565b9150611878836112bc565b9250828202611886816112bc565b9150828204841483151761189d5761189c611833565b5b5092915050565b7f53656e742076616c7565206973206c657373207468616e2074686520746f746160008201527f6c20636f73740000000000000000000000000000000000000000000000000000602082015250565b60006119006026836117b6565b915061190b826118a4565b604082019050919050565b6000602082019050818103600083015261192f816118f3565b9050919050565b600060408201905061194b6000830185611737565b6119586020830184611332565b9392505050565b60008151905061196e816112c6565b92915050565b60006020828403121561198a576119896112b7565b5b60006119988482850161195f565b91505092915050565b7f53656c6c657220646f6573206e6f74206861766520656e6f75676820746f6b6560008201527f6e7320666f722073616c65000000000000000000000000000000000000000000602082015250565b60006119fd602b836117b6565b9150611a08826119a1565b604082019050919050565b60006020820190508181036000830152611a2c816119f0565b9050919050565b600082825260208201905092915050565b50565b6000611a54600083611a33565b9150611a5f82611a44565b600082019050919050565b600060a082019050611a7f6000830187611737565b611a8c6020830186611737565b611a996040830185611332565b611aa66060830184611332565b8181036080830152611ab781611a47565b905095945050505050565b600081905092915050565b6000611ada600083611ac2565b9150611ae582611a44565b600082019050919050565b6000611afb82611acd565b9150819050919050565b7f4661696c656420746f207472616e7366657220457468657220746f207468652060008201527f73656c6c65720000000000000000000000000000000000000000000000000000602082015250565b6000611b616026836117b6565b9150611b6c82611b05565b604082019050919050565b60006020820190508181036000830152611b9081611b54565b9050919050565b6000611ba2826112bc565b9150611bad836112bc565b9250828203905081811115611bc557611bc4611833565b5b92915050565b7f4f6e6c79207468652073656c6c65722063616e2072656d6f766520746865206c60008201527f697374696e670000000000000000000000000000000000000000000000000000602082015250565b6000611c276026836117b6565b9150611c3282611bcb565b604082019050919050565b60006020820190508181036000830152611c5681611c1a565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000611c97826112bc565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611cc957611cc8611833565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f546f6b656e2049442073686f756c642062652067726561746572207468616e2060008201527f7a65726f00000000000000000000000000000000000000000000000000000000602082015250565b6000611d5f6024836117b6565b9150611d6a82611d03565b604082019050919050565b60006020820190508181036000830152611d8e81611d52565b9050919050565b7f50726963652073686f756c642062652067726561746572207468616e207a657260008201527f6f00000000000000000000000000000000000000000000000000000000000000602082015250565b6000611df16021836117b6565b9150611dfc82611d95565b604082019050919050565b60006020820190508181036000830152611e2081611de4565b9050919050565b7f416d6f756e742073686f756c642062652067726561746572207468616e207a6560008201527f726f000000000000000000000000000000000000000000000000000000000000602082015250565b6000611e836022836117b6565b9150611e8e82611e27565b604082019050919050565b60006020820190508181036000830152611eb281611e76565b905091905056fea2646970667358221220d9160ef8e81c2dd2f8dccc68a2aca777aedc12f7c762f3e094fb57eb34ca7a5464736f6c63430008120033",
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
// Solidity: function getAllListings() view returns((uint256,uint256,address,uint256,uint256,bool)[])
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
// Solidity: function getAllListings() view returns((uint256,uint256,address,uint256,uint256,bool)[])
func (_Sale *SaleSession) GetAllListings() ([]SaleListing, error) {
	return _Sale.Contract.GetAllListings(&_Sale.CallOpts)
}

// GetAllListings is a free data retrieval call binding the contract method 0xae73ccec.
//
// Solidity: function getAllListings() view returns((uint256,uint256,address,uint256,uint256,bool)[])
func (_Sale *SaleCallerSession) GetAllListings() ([]SaleListing, error) {
	return _Sale.Contract.GetAllListings(&_Sale.CallOpts)
}

// GetListingsByUser is a free data retrieval call binding the contract method 0x69df8327.
//
// Solidity: function getListingsByUser(address user) view returns((uint256,uint256,address,uint256,uint256,bool)[])
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
// Solidity: function getListingsByUser(address user) view returns((uint256,uint256,address,uint256,uint256,bool)[])
func (_Sale *SaleSession) GetListingsByUser(user common.Address) ([]SaleListing, error) {
	return _Sale.Contract.GetListingsByUser(&_Sale.CallOpts, user)
}

// GetListingsByUser is a free data retrieval call binding the contract method 0x69df8327.
//
// Solidity: function getListingsByUser(address user) view returns((uint256,uint256,address,uint256,uint256,bool)[])
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

// ListingIds is a free data retrieval call binding the contract method 0x52bea360.
//
// Solidity: function listingIds(uint256 ) view returns(uint256)
func (_Sale *SaleCaller) ListingIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "listingIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ListingIds is a free data retrieval call binding the contract method 0x52bea360.
//
// Solidity: function listingIds(uint256 ) view returns(uint256)
func (_Sale *SaleSession) ListingIds(arg0 *big.Int) (*big.Int, error) {
	return _Sale.Contract.ListingIds(&_Sale.CallOpts, arg0)
}

// ListingIds is a free data retrieval call binding the contract method 0x52bea360.
//
// Solidity: function listingIds(uint256 ) view returns(uint256)
func (_Sale *SaleCallerSession) ListingIds(arg0 *big.Int) (*big.Int, error) {
	return _Sale.Contract.ListingIds(&_Sale.CallOpts, arg0)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 listingId, uint256 tokenId, address seller, uint256 price, uint256 amount, bool isForSale)
func (_Sale *SaleCaller) Listings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ListingId *big.Int
	TokenId   *big.Int
	Seller    common.Address
	Price     *big.Int
	Amount    *big.Int
	IsForSale bool
}, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "listings", arg0)

	outstruct := new(struct {
		ListingId *big.Int
		TokenId   *big.Int
		Seller    common.Address
		Price     *big.Int
		Amount    *big.Int
		IsForSale bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ListingId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.IsForSale = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 listingId, uint256 tokenId, address seller, uint256 price, uint256 amount, bool isForSale)
func (_Sale *SaleSession) Listings(arg0 *big.Int) (struct {
	ListingId *big.Int
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
// Solidity: function listings(uint256 ) view returns(uint256 listingId, uint256 tokenId, address seller, uint256 price, uint256 amount, bool isForSale)
func (_Sale *SaleCallerSession) Listings(arg0 *big.Int) (struct {
	ListingId *big.Int
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
