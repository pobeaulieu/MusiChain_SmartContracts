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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyAmount\",\"type\":\"uint256\"}],\"name\":\"buyToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllListings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForSale\",\"type\":\"bool\"}],\"internalType\":\"structSale.Listing[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getListingsByUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForSale\",\"type\":\"bool\"}],\"internalType\":\"structSale.Listing[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getOwnedTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"listToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForSale\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"removeListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001f7138038062001f718339818101604052810190620000379190620000fc565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200012e565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000b08262000083565b9050919050565b6000620000c482620000a3565b9050919050565b620000d681620000b7565b8114620000e257600080fd5b50565b600081519050620000f681620000cb565b92915050565b6000602082840312156200011557620001146200007e565b5b60006200012584828501620000e5565b91505092915050565b611e33806200013e6000396000f3fe6080604052600436106100865760003560e01c8063ae73ccec11610059578063ae73ccec14610138578063bce64a7d14610163578063d58778d61461018c578063d9d61655146101c9578063de74e57b1461020657610086565b8063057466ea1461008b578063479ad4c3146100a757806355a373d6146100d057806369df8327146100fb575b600080fd5b6100a560048036038101906100a09190611257565b610247565b005b3480156100b357600080fd5b506100ce60048036038101906100c99190611297565b610662565b005b3480156100dc57600080fd5b506100e5610738565b6040516100f29190611343565b60405180910390f35b34801561010757600080fd5b50610122600480360381019061011d919061139c565b61075c565b60405161012f9190611519565b60405180910390f35b34801561014457600080fd5b5061014d610ab6565b60405161015a9190611519565b60405180910390f35b34801561016f57600080fd5b5061018a6004803603810190610185919061153b565b610cf8565b005b34801561019857600080fd5b506101b360048036038101906101ae9190611297565b610eca565b6040516101c0919061159d565b60405180910390f35b3480156101d557600080fd5b506101f060048036038101906101eb919061139c565b610eee565b6040516101fd9190611667565b60405180910390f35b34801561021257600080fd5b5061022d60048036038101906102289190611297565b611172565b60405161023e9594939291906116a7565b60405180910390f35b6002600083815260200190815260200160002060040160009054906101000a900460ff166102aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a190611757565b60405180910390fd5b6000600260008481526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201548152602001600382015481526020016004820160009054906101000a900460ff1615151515815250509050600082826040015161036d91906117a6565b9050803410156103b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103a99061185a565b60405180910390fd5b8260008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1662fdd58e846020015185600001516040518363ffffffff1660e01b815260040161041592919061187a565b602060405180830381865afa158015610432573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061045691906118b8565b1015610497576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048e90611957565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f242432a8360200151338560000151876040518563ffffffff1660e01b81526004016104fe94939291906119ae565b600060405180830381600087803b15801561051857600080fd5b505af115801561052c573d6000803e3d6000fd5b505050506000826020015173ffffffffffffffffffffffffffffffffffffffff168260405161055a90611a34565b60006040518083038185875af1925050503d8060008114610597576040519150601f19603f3d011682016040523d82523d6000602084013e61059c565b606091505b50509050806105e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105d790611abb565b60405180910390fd5b836002600087815260200190815260200160002060030160008282546106069190611adb565b92505081905550600060026000878152602001908152602001600020600301540361065b5760006002600087815260200190815260200160002060040160006101000a81548160ff0219169083151502179055505b5050505050565b3373ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610706576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106fd90611b81565b60405180910390fd5b60006002600083815260200190815260200160002060040160006101000a81548160ff02191690831515021790555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606000805b600180549050811015610863578373ffffffffffffffffffffffffffffffffffffffff16600260006001848154811061079e5761079d611ba1565b5b9060005260206000200154815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614801561083c5750600260006001838154811061081157610810611ba1565b5b9060005260206000200154815260200190815260200160002060040160009054906101000a900460ff165b1561085057818061084c90611bd0565b9250505b808061085b90611bd0565b915050610762565b5060008167ffffffffffffffff8111156108805761087f611c18565b5b6040519080825280602002602001820160405280156108b957816020015b6108a66111d5565b81526020019060019003908161089e5790505b5090506000805b600180549050811015610aaa578573ffffffffffffffffffffffffffffffffffffffff1660026000600184815481106108fc576108fb611ba1565b5b9060005260206000200154815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614801561099a5750600260006001838154811061096f5761096e611ba1565b5b9060005260206000200154815260200190815260200160002060040160009054906101000a900460ff165b15610a975760026000600183815481106109b7576109b6611ba1565b5b906000526020600020015481526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201548152602001600382015481526020016004820160009054906101000a900460ff161515151581525050838381518110610a7d57610a7c611ba1565b5b60200260200101819052508180610a9390611bd0565b9250505b8080610aa290611bd0565b9150506108c0565b50819350505050919050565b60606000805b600180549050811015610b32576002600060018381548110610ae157610ae0611ba1565b5b9060005260206000200154815260200190815260200160002060040160009054906101000a900460ff1615610b1f578180610b1b90611bd0565b9250505b8080610b2a90611bd0565b915050610abc565b5060008167ffffffffffffffff811115610b4f57610b4e611c18565b5b604051908082528060200260200182016040528015610b8857816020015b610b756111d5565b815260200190600190039081610b6d5790505b5090506000805b600180549050811015610cee576002600060018381548110610bb457610bb3611ba1565b5b9060005260206000200154815260200190815260200160002060040160009054906101000a900460ff1615610cdb576002600060018381548110610bfb57610bfa611ba1565b5b906000526020600020015481526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201548152602001600382015481526020016004820160009054906101000a900460ff161515151581525050838381518110610cc157610cc0611ba1565b5b60200260200101819052508180610cd790611bd0565b9250505b8080610ce690611bd0565b915050610b8f565b5081935050505090565b60008311610d3b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3290611cb9565b60405180910390fd5b60008211610d7e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d7590611d4b565b60405180910390fd5b60008111610dc1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610db890611ddd565b60405180910390fd5b6040518060a001604052808481526020013373ffffffffffffffffffffffffffffffffffffffff16815260200183815260200182815260200160011515815250600260008581526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201556060820151816003015560808201518160040160006101000a81548160ff0219169083151502179055509050506001839080600181540180825580915050600190039060005260206000200160009091909190915055505050565b60018181548110610eda57600080fd5b906000526020600020016000915090505481565b60606000805b600180549050811015610fe55760008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1662fdd58e8660018581548110610f5457610f53611ba1565b5b90600052602060002001546040518363ffffffff1660e01b8152600401610f7c92919061187a565b602060405180830381865afa158015610f99573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fbd91906118b8565b1115610fd2578180610fce90611bd0565b9250505b8080610fdd90611bd0565b915050610ef4565b5060008167ffffffffffffffff81111561100257611001611c18565b5b6040519080825280602002602001820160405280156110305781602001602082028036833780820191505090505b5090506000805b6001805490508110156111665760008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1662fdd58e886001858154811061109757611096611ba1565b5b90600052602060002001546040518363ffffffff1660e01b81526004016110bf92919061187a565b602060405180830381865afa1580156110dc573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061110091906118b8565b1115611153576001818154811061111a57611119611ba1565b5b906000526020600020015483838151811061113857611137611ba1565b5b602002602001018181525050818061114f90611bd0565b9250505b808061115e90611bd0565b915050611037565b50819350505050919050565b60026020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030154908060040160009054906101000a900460ff16905085565b6040518060a0016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081526020016000151581525090565b600080fd5b6000819050919050565b61123481611221565b811461123f57600080fd5b50565b6000813590506112518161122b565b92915050565b6000806040838503121561126e5761126d61121c565b5b600061127c85828601611242565b925050602061128d85828601611242565b9150509250929050565b6000602082840312156112ad576112ac61121c565b5b60006112bb84828501611242565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006113096113046112ff846112c4565b6112e4565b6112c4565b9050919050565b600061131b826112ee565b9050919050565b600061132d82611310565b9050919050565b61133d81611322565b82525050565b60006020820190506113586000830184611334565b92915050565b6000611369826112c4565b9050919050565b6113798161135e565b811461138457600080fd5b50565b60008135905061139681611370565b92915050565b6000602082840312156113b2576113b161121c565b5b60006113c084828501611387565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6113fe81611221565b82525050565b61140d8161135e565b82525050565b60008115159050919050565b61142881611413565b82525050565b60a08201600082015161144460008501826113f5565b5060208201516114576020850182611404565b50604082015161146a60408501826113f5565b50606082015161147d60608501826113f5565b506080820151611490608085018261141f565b50505050565b60006114a2838361142e565b60a08301905092915050565b6000602082019050919050565b60006114c6826113c9565b6114d081856113d4565b93506114db836113e5565b8060005b8381101561150c5781516114f38882611496565b97506114fe836114ae565b9250506001810190506114df565b5085935050505092915050565b6000602082019050818103600083015261153381846114bb565b905092915050565b6000806000606084860312156115545761155361121c565b5b600061156286828701611242565b935050602061157386828701611242565b925050604061158486828701611242565b9150509250925092565b61159781611221565b82525050565b60006020820190506115b2600083018461158e565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60006115f083836113f5565b60208301905092915050565b6000602082019050919050565b6000611614826115b8565b61161e81856115c3565b9350611629836115d4565b8060005b8381101561165a57815161164188826115e4565b975061164c836115fc565b92505060018101905061162d565b5085935050505092915050565b600060208201905081810360008301526116818184611609565b905092915050565b6116928161135e565b82525050565b6116a181611413565b82525050565b600060a0820190506116bc600083018861158e565b6116c96020830187611689565b6116d6604083018661158e565b6116e3606083018561158e565b6116f06080830184611698565b9695505050505050565b600082825260208201905092915050565b7f54686973206974656d206973206e6f7420666f722073616c6500000000000000600082015250565b60006117416019836116fa565b915061174c8261170b565b602082019050919050565b6000602082019050818103600083015261177081611734565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006117b182611221565b91506117bc83611221565b92508282026117ca81611221565b915082820484148315176117e1576117e0611777565b5b5092915050565b7f53656e742076616c7565206973206c657373207468616e2074686520746f746160008201527f6c20636f73740000000000000000000000000000000000000000000000000000602082015250565b60006118446026836116fa565b915061184f826117e8565b604082019050919050565b6000602082019050818103600083015261187381611837565b9050919050565b600060408201905061188f6000830185611689565b61189c602083018461158e565b9392505050565b6000815190506118b28161122b565b92915050565b6000602082840312156118ce576118cd61121c565b5b60006118dc848285016118a3565b91505092915050565b7f53656c6c657220646f6573206e6f74206861766520656e6f75676820746f6b6560008201527f6e7320666f722073616c65000000000000000000000000000000000000000000602082015250565b6000611941602b836116fa565b915061194c826118e5565b604082019050919050565b6000602082019050818103600083015261197081611934565b9050919050565b600082825260208201905092915050565b50565b6000611998600083611977565b91506119a382611988565b600082019050919050565b600060a0820190506119c36000830187611689565b6119d06020830186611689565b6119dd604083018561158e565b6119ea606083018461158e565b81810360808301526119fb8161198b565b905095945050505050565b600081905092915050565b6000611a1e600083611a06565b9150611a2982611988565b600082019050919050565b6000611a3f82611a11565b9150819050919050565b7f4661696c656420746f207472616e7366657220457468657220746f207468652060008201527f73656c6c65720000000000000000000000000000000000000000000000000000602082015250565b6000611aa56026836116fa565b9150611ab082611a49565b604082019050919050565b60006020820190508181036000830152611ad481611a98565b9050919050565b6000611ae682611221565b9150611af183611221565b9250828203905081811115611b0957611b08611777565b5b92915050565b7f4f6e6c79207468652073656c6c65722063616e2072656d6f766520746865206c60008201527f697374696e670000000000000000000000000000000000000000000000000000602082015250565b6000611b6b6026836116fa565b9150611b7682611b0f565b604082019050919050565b60006020820190508181036000830152611b9a81611b5e565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000611bdb82611221565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c0d57611c0c611777565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f546f6b656e2049442073686f756c642062652067726561746572207468616e2060008201527f7a65726f00000000000000000000000000000000000000000000000000000000602082015250565b6000611ca36024836116fa565b9150611cae82611c47565b604082019050919050565b60006020820190508181036000830152611cd281611c96565b9050919050565b7f50726963652073686f756c642062652067726561746572207468616e207a657260008201527f6f00000000000000000000000000000000000000000000000000000000000000602082015250565b6000611d356021836116fa565b9150611d4082611cd9565b604082019050919050565b60006020820190508181036000830152611d6481611d28565b9050919050565b7f416d6f756e742073686f756c642062652067726561746572207468616e207a6560008201527f726f000000000000000000000000000000000000000000000000000000000000602082015250565b6000611dc76022836116fa565b9150611dd282611d6b565b604082019050919050565b60006020820190508181036000830152611df681611dba565b905091905056fea2646970667358221220cee7e464b6993ff82260b5c1c4c67b746f8830d861fdc24fb43f1b2c64dfd85d64736f6c63430008120033",
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

// BuyToken is a paid mutator transaction binding the contract method 0x057466ea.
//
// Solidity: function buyToken(uint256 tokenId, uint256 buyAmount) payable returns()
func (_Sale *SaleTransactor) BuyToken(opts *bind.TransactOpts, tokenId *big.Int, buyAmount *big.Int) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "buyToken", tokenId, buyAmount)
}

// BuyToken is a paid mutator transaction binding the contract method 0x057466ea.
//
// Solidity: function buyToken(uint256 tokenId, uint256 buyAmount) payable returns()
func (_Sale *SaleSession) BuyToken(tokenId *big.Int, buyAmount *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.BuyToken(&_Sale.TransactOpts, tokenId, buyAmount)
}

// BuyToken is a paid mutator transaction binding the contract method 0x057466ea.
//
// Solidity: function buyToken(uint256 tokenId, uint256 buyAmount) payable returns()
func (_Sale *SaleTransactorSession) BuyToken(tokenId *big.Int, buyAmount *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.BuyToken(&_Sale.TransactOpts, tokenId, buyAmount)
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
