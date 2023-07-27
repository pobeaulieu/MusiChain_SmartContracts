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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyAmount\",\"type\":\"uint256\"}],\"name\":\"buyToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"etherValue\",\"type\":\"uint256\"}],\"name\":\"convertToWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllListings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForSale\",\"type\":\"bool\"}],\"internalType\":\"structSale.Listing[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getListingsByUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForSale\",\"type\":\"bool\"}],\"internalType\":\"structSale.Listing[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getOwnedTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInEther\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"listToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForSale\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"removeListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001fe438038062001fe48339818101604052810190620000379190620000fc565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200012e565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000b08262000083565b9050919050565b6000620000c482620000a3565b9050919050565b620000d681620000b7565b8114620000e257600080fd5b50565b600081519050620000f681620000cb565b92915050565b6000602082840312156200011557620001146200007e565b5b60006200012584828501620000e5565b91505092915050565b611ea6806200013e6000396000f3fe6080604052600436106100915760003560e01c8063ae73ccec11610059578063ae73ccec14610180578063bce64a7d146101ab578063d58778d6146101d4578063d9d6165514610211578063de74e57b1461024e57610091565b8063057466ea14610096578063479ad4c3146100b257806355a373d6146100db57806369df832714610106578063a60dc38a14610143575b600080fd5b6100b060048036038101906100ab91906112ca565b61028f565b005b3480156100be57600080fd5b506100d960048036038101906100d4919061130a565b6106aa565b005b3480156100e757600080fd5b506100f0610780565b6040516100fd91906113b6565b60405180910390f35b34801561011257600080fd5b5061012d6004803603810190610128919061140f565b6107a4565b60405161013a919061158c565b60405180910390f35b34801561014f57600080fd5b5061016a6004803603810190610165919061130a565b610afe565b60405161017791906115bd565b60405180910390f35b34801561018c57600080fd5b50610195610b1b565b6040516101a2919061158c565b60405180910390f35b3480156101b757600080fd5b506101d260048036038101906101cd91906115d8565b610d5d565b005b3480156101e057600080fd5b506101fb60048036038101906101f6919061130a565b610f3d565b60405161020891906115bd565b60405180910390f35b34801561021d57600080fd5b506102386004803603810190610233919061140f565b610f61565b60405161024591906116da565b60405180910390f35b34801561025a57600080fd5b506102756004803603810190610270919061130a565b6111e5565b60405161028695949392919061171a565b60405180910390f35b6002600083815260200190815260200160002060040160009054906101000a900460ff166102f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102e9906117ca565b60405180910390fd5b6000600260008481526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201548152602001600382015481526020016004820160009054906101000a900460ff161515151581525050905060008282604001516103b59190611819565b9050803410156103fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f1906118cd565b60405180910390fd5b8260008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1662fdd58e846020015185600001516040518363ffffffff1660e01b815260040161045d9291906118ed565b602060405180830381865afa15801561047a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061049e919061192b565b10156104df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d6906119ca565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f242432a8360200151338560000151876040518563ffffffff1660e01b81526004016105469493929190611a21565b600060405180830381600087803b15801561056057600080fd5b505af1158015610574573d6000803e3d6000fd5b505050506000826020015173ffffffffffffffffffffffffffffffffffffffff16826040516105a290611aa7565b60006040518083038185875af1925050503d80600081146105df576040519150601f19603f3d011682016040523d82523d6000602084013e6105e4565b606091505b5050905080610628576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161061f90611b2e565b60405180910390fd5b8360026000878152602001908152602001600020600301600082825461064e9190611b4e565b9250508190555060006002600087815260200190815260200160002060030154036106a35760006002600087815260200190815260200160002060040160006101000a81548160ff0219169083151502179055505b5050505050565b3373ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461074e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161074590611bf4565b60405180910390fd5b60006002600083815260200190815260200160002060040160006101000a81548160ff02191690831515021790555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606000805b6001805490508110156108ab578373ffffffffffffffffffffffffffffffffffffffff1660026000600184815481106107e6576107e5611c14565b5b9060005260206000200154815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161480156108845750600260006001838154811061085957610858611c14565b5b9060005260206000200154815260200190815260200160002060040160009054906101000a900460ff165b1561089857818061089490611c43565b9250505b80806108a390611c43565b9150506107aa565b5060008167ffffffffffffffff8111156108c8576108c7611c8b565b5b60405190808252806020026020018201604052801561090157816020015b6108ee611248565b8152602001906001900390816108e65790505b5090506000805b600180549050811015610af2578573ffffffffffffffffffffffffffffffffffffffff16600260006001848154811061094457610943611c14565b5b9060005260206000200154815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161480156109e2575060026000600183815481106109b7576109b6611c14565b5b9060005260206000200154815260200190815260200160002060040160009054906101000a900460ff165b15610adf5760026000600183815481106109ff576109fe611c14565b5b906000526020600020015481526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201548152602001600382015481526020016004820160009054906101000a900460ff161515151581525050838381518110610ac557610ac4611c14565b5b60200260200101819052508180610adb90611c43565b9250505b8080610aea90611c43565b915050610908565b50819350505050919050565b6000670de0b6b3a764000082610b149190611819565b9050919050565b60606000805b600180549050811015610b97576002600060018381548110610b4657610b45611c14565b5b9060005260206000200154815260200190815260200160002060040160009054906101000a900460ff1615610b84578180610b8090611c43565b9250505b8080610b8f90611c43565b915050610b21565b5060008167ffffffffffffffff811115610bb457610bb3611c8b565b5b604051908082528060200260200182016040528015610bed57816020015b610bda611248565b815260200190600190039081610bd25790505b5090506000805b600180549050811015610d53576002600060018381548110610c1957610c18611c14565b5b9060005260206000200154815260200190815260200160002060040160009054906101000a900460ff1615610d40576002600060018381548110610c6057610c5f611c14565b5b906000526020600020015481526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201548152602001600382015481526020016004820160009054906101000a900460ff161515151581525050838381518110610d2657610d25611c14565b5b60200260200101819052508180610d3c90611c43565b9250505b8080610d4b90611c43565b915050610bf4565b5081935050505090565b60008311610da0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d9790611d2c565b60405180910390fd5b60008211610de3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dda90611dbe565b60405180910390fd5b60008111610e26576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e1d90611e50565b60405180910390fd5b6000610e3183610afe565b90506040518060a001604052808581526020013373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200183815260200160011515815250600260008681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201556060820151816003015560808201518160040160006101000a81548160ff021916908315150217905550905050600184908060018154018082558091505060019003906000526020600020016000909190919091505550505050565b60018181548110610f4d57600080fd5b906000526020600020016000915090505481565b60606000805b6001805490508110156110585760008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1662fdd58e8660018581548110610fc757610fc6611c14565b5b90600052602060002001546040518363ffffffff1660e01b8152600401610fef9291906118ed565b602060405180830381865afa15801561100c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611030919061192b565b111561104557818061104190611c43565b9250505b808061105090611c43565b915050610f67565b5060008167ffffffffffffffff81111561107557611074611c8b565b5b6040519080825280602002602001820160405280156110a35781602001602082028036833780820191505090505b5090506000805b6001805490508110156111d95760008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1662fdd58e886001858154811061110a57611109611c14565b5b90600052602060002001546040518363ffffffff1660e01b81526004016111329291906118ed565b602060405180830381865afa15801561114f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611173919061192b565b11156111c6576001818154811061118d5761118c611c14565b5b90600052602060002001548383815181106111ab576111aa611c14565b5b60200260200101818152505081806111c290611c43565b9250505b80806111d190611c43565b9150506110aa565b50819350505050919050565b60026020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030154908060040160009054906101000a900460ff16905085565b6040518060a0016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081526020016000151581525090565b600080fd5b6000819050919050565b6112a781611294565b81146112b257600080fd5b50565b6000813590506112c48161129e565b92915050565b600080604083850312156112e1576112e061128f565b5b60006112ef858286016112b5565b9250506020611300858286016112b5565b9150509250929050565b6000602082840312156113205761131f61128f565b5b600061132e848285016112b5565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061137c61137761137284611337565b611357565b611337565b9050919050565b600061138e82611361565b9050919050565b60006113a082611383565b9050919050565b6113b081611395565b82525050565b60006020820190506113cb60008301846113a7565b92915050565b60006113dc82611337565b9050919050565b6113ec816113d1565b81146113f757600080fd5b50565b600081359050611409816113e3565b92915050565b6000602082840312156114255761142461128f565b5b6000611433848285016113fa565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61147181611294565b82525050565b611480816113d1565b82525050565b60008115159050919050565b61149b81611486565b82525050565b60a0820160008201516114b76000850182611468565b5060208201516114ca6020850182611477565b5060408201516114dd6040850182611468565b5060608201516114f06060850182611468565b5060808201516115036080850182611492565b50505050565b600061151583836114a1565b60a08301905092915050565b6000602082019050919050565b60006115398261143c565b6115438185611447565b935061154e83611458565b8060005b8381101561157f5781516115668882611509565b975061157183611521565b925050600181019050611552565b5085935050505092915050565b600060208201905081810360008301526115a6818461152e565b905092915050565b6115b781611294565b82525050565b60006020820190506115d260008301846115ae565b92915050565b6000806000606084860312156115f1576115f061128f565b5b60006115ff868287016112b5565b9350506020611610868287016112b5565b9250506040611621868287016112b5565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60006116638383611468565b60208301905092915050565b6000602082019050919050565b60006116878261162b565b6116918185611636565b935061169c83611647565b8060005b838110156116cd5781516116b48882611657565b97506116bf8361166f565b9250506001810190506116a0565b5085935050505092915050565b600060208201905081810360008301526116f4818461167c565b905092915050565b611705816113d1565b82525050565b61171481611486565b82525050565b600060a08201905061172f60008301886115ae565b61173c60208301876116fc565b61174960408301866115ae565b61175660608301856115ae565b611763608083018461170b565b9695505050505050565b600082825260208201905092915050565b7f54686973206974656d206973206e6f7420666f722073616c6500000000000000600082015250565b60006117b460198361176d565b91506117bf8261177e565b602082019050919050565b600060208201905081810360008301526117e3816117a7565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061182482611294565b915061182f83611294565b925082820261183d81611294565b91508282048414831517611854576118536117ea565b5b5092915050565b7f53656e742076616c7565206973206c657373207468616e2074686520746f746160008201527f6c20636f73740000000000000000000000000000000000000000000000000000602082015250565b60006118b760268361176d565b91506118c28261185b565b604082019050919050565b600060208201905081810360008301526118e6816118aa565b9050919050565b600060408201905061190260008301856116fc565b61190f60208301846115ae565b9392505050565b6000815190506119258161129e565b92915050565b6000602082840312156119415761194061128f565b5b600061194f84828501611916565b91505092915050565b7f53656c6c657220646f6573206e6f74206861766520656e6f75676820746f6b6560008201527f6e7320666f722073616c65000000000000000000000000000000000000000000602082015250565b60006119b4602b8361176d565b91506119bf82611958565b604082019050919050565b600060208201905081810360008301526119e3816119a7565b9050919050565b600082825260208201905092915050565b50565b6000611a0b6000836119ea565b9150611a16826119fb565b600082019050919050565b600060a082019050611a3660008301876116fc565b611a4360208301866116fc565b611a5060408301856115ae565b611a5d60608301846115ae565b8181036080830152611a6e816119fe565b905095945050505050565b600081905092915050565b6000611a91600083611a79565b9150611a9c826119fb565b600082019050919050565b6000611ab282611a84565b9150819050919050565b7f4661696c656420746f207472616e7366657220457468657220746f207468652060008201527f73656c6c65720000000000000000000000000000000000000000000000000000602082015250565b6000611b1860268361176d565b9150611b2382611abc565b604082019050919050565b60006020820190508181036000830152611b4781611b0b565b9050919050565b6000611b5982611294565b9150611b6483611294565b9250828203905081811115611b7c57611b7b6117ea565b5b92915050565b7f4f6e6c79207468652073656c6c65722063616e2072656d6f766520746865206c60008201527f697374696e670000000000000000000000000000000000000000000000000000602082015250565b6000611bde60268361176d565b9150611be982611b82565b604082019050919050565b60006020820190508181036000830152611c0d81611bd1565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000611c4e82611294565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c8057611c7f6117ea565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f546f6b656e2049442073686f756c642062652067726561746572207468616e2060008201527f7a65726f00000000000000000000000000000000000000000000000000000000602082015250565b6000611d1660248361176d565b9150611d2182611cba565b604082019050919050565b60006020820190508181036000830152611d4581611d09565b9050919050565b7f50726963652073686f756c642062652067726561746572207468616e207a657260008201527f6f00000000000000000000000000000000000000000000000000000000000000602082015250565b6000611da860218361176d565b9150611db382611d4c565b604082019050919050565b60006020820190508181036000830152611dd781611d9b565b9050919050565b7f416d6f756e742073686f756c642062652067726561746572207468616e207a6560008201527f726f000000000000000000000000000000000000000000000000000000000000602082015250565b6000611e3a60228361176d565b9150611e4582611dde565b604082019050919050565b60006020820190508181036000830152611e6981611e2d565b905091905056fea26469706673582212205f40c63dd541085fe4ed4c38c40fc2088230e5e8ca7ab7984b5266903721886064736f6c63430008120033",
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
