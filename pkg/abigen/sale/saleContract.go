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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyAmount\",\"type\":\"uint256\"}],\"name\":\"buyToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllListings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForSale\",\"type\":\"bool\"}],\"internalType\":\"structSale.Listing[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getListingsByToken\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForSale\",\"type\":\"bool\"}],\"internalType\":\"structSale.Listing[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getListingsByUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForSale\",\"type\":\"bool\"}],\"internalType\":\"structSale.Listing[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getOwnedTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hasToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"listToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listingCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForSale\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ownedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"removeListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenToListings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620021f3380380620021f38339818101604052810190620000379190620000fc565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200012e565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000b08262000083565b9050919050565b6000620000c482620000a3565b9050919050565b620000d681620000b7565b8114620000e257600080fd5b50565b600081519050620000f681620000cb565b92915050565b6000602082840312156200011557620001146200007e565b5b60006200012584828501620000e5565b91505092915050565b6120b5806200013e6000396000f3fe6080604052600436106100c25760003560e01c8063a9b07c261161007f578063d9d6165511610059578063d9d616551461026d578063de61a0df146102aa578063de74e57b146102e7578063e149f03614610329576100c2565b8063a9b07c26146101ee578063ae73ccec14610219578063bce64a7d14610244576100c2565b8063057466ea146100c75780631ca17eff146100e3578063479ad4c31461012057806355a373d61461014957806363a2e3cc1461017457806369df8327146101b1575b600080fd5b6100e160048036038101906100dc919061145d565b610366565b005b3480156100ef57600080fd5b5061010a6004803603810190610105919061145d565b6107f5565b60405161011791906114ac565b60405180910390f35b34801561012c57600080fd5b50610147600480360381019061014291906114c7565b610826565b005b34801561015557600080fd5b5061015e6108fc565b60405161016b9190611573565b60405180910390f35b34801561018057600080fd5b5061019b600480360381019061019691906114c7565b610920565b6040516101a89190611703565b60405180910390f35b3480156101bd57600080fd5b506101d860048036038101906101d39190611751565b610aff565b6040516101e59190611703565b60405180910390f35b3480156101fa57600080fd5b50610203610da8565b60405161021091906114ac565b60405180910390f35b34801561022557600080fd5b5061022e610dae565b60405161023b9190611703565b60405180910390f35b34801561025057600080fd5b5061026b6004803603810190610266919061177e565b610f31565b005b34801561027957600080fd5b50610294600480360381019061028f9190611751565b611274565b6040516102a19190611880565b60405180910390f35b3480156102b657600080fd5b506102d160048036038101906102cc91906118a2565b61130b565b6040516102de91906118f1565b60405180910390f35b3480156102f357600080fd5b5061030e600480360381019061030991906114c7565b61133a565b6040516103209695949392919061191b565b60405180910390f35b34801561033557600080fd5b50610350600480360381019061034b91906118a2565b6113a3565b60405161035d91906114ac565b60405180910390f35b6002600083815260200190815260200160002060050160009054906101000a900460ff166103c9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103c0906119d9565b60405180910390fd5b6000600260008481526020019081526020016000206040518060c001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201548152602001600482015481526020016005820160009054906101000a900460ff161515151581525050905060008282606001516104969190611a28565b9050803410156104db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d290611adc565b60405180910390fd5b8260008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1662fdd58e846040015185602001516040518363ffffffff1660e01b815260040161053e929190611afc565b602060405180830381865afa15801561055b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057f9190611b3a565b10156105c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b790611bd9565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f242432a8360400151338560200151876040518563ffffffff1660e01b81526004016106279493929190611c30565b600060405180830381600087803b15801561064157600080fd5b505af1158015610655573d6000803e3d6000fd5b505050506000826040015173ffffffffffffffffffffffffffffffffffffffff168260405161068390611cb6565b60006040518083038185875af1925050503d80600081146106c0576040519150601f19603f3d011682016040523d82523d6000602084013e6106c5565b606091505b5050905080610709576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070090611d3d565b60405180910390fd5b8360026000878152602001908152602001600020600401600082825461072f9190611d5d565b92505081905550600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208360200151908060018154018082558091505060019003906000526020600020016000909190919091505560006002600087815260200190815260200160002060040154036107ee5760006002600087815260200190815260200160002060050160006101000a81548160ff0219169083151502179055505b5050505050565b6003602052816000526040600020818154811061081157600080fd5b90600052602060002001600091509150505481565b3373ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146108ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108c190611e03565b60405180910390fd5b60006002600083815260200190815260200160002060050160006101000a81548160ff02191690831515021790555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606060006003600084815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561098157602002820191906000526020600020905b81548152602001906001019080831161096d575b5050505050905060008151905060008167ffffffffffffffff8111156109aa576109a9611e23565b5b6040519080825280602002602001820160405280156109e357816020015b6109d06113d4565b8152602001906001900390816109c85790505b50905060005b82811015610af35760026000858381518110610a0857610a07611e52565b5b602002602001015181526020019081526020016000206040518060c001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201548152602001600482015481526020016005820160009054906101000a900460ff161515151581525050828281518110610ad557610ad4611e52565b5b60200260200101819052508080610aeb90611e81565b9150506109e9565b50809350505050919050565b6060600060015467ffffffffffffffff811115610b1f57610b1e611e23565b5b604051908082528060200260200182016040528015610b4d5781602001602082028036833780820191505090505b5090506000805b600154811015610c36578473ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16148015610bef57506002600082815260200190815260200160002060050160009054906101000a900460ff165b15610c235780838381518110610c0857610c07611e52565b5b6020026020010181815250508180610c1f90611e81565b9250505b8080610c2e90611e81565b915050610b54565b5060008167ffffffffffffffff811115610c5357610c52611e23565b5b604051908082528060200260200182016040528015610c8c57816020015b610c796113d4565b815260200190600190039081610c715790505b50905060005b82811015610d9c5760026000858381518110610cb157610cb0611e52565b5b602002602001015181526020019081526020016000206040518060c001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201548152602001600482015481526020016005820160009054906101000a900460ff161515151581525050828281518110610d7e57610d7d611e52565b5b60200260200101819052508080610d9490611e81565b915050610c92565b50809350505050919050565b60015481565b6060600060015467ffffffffffffffff811115610dce57610dcd611e23565b5b604051908082528060200260200182016040528015610e0757816020015b610df46113d4565b815260200190600190039081610dec5790505b50905060005b600154811015610f29576002600082815260200190815260200160002060050160009054906101000a900460ff1615610f1657600260008281526020019081526020016000206040518060c001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201548152602001600482015481526020016005820160009054906101000a900460ff161515151581525050828281518110610f0a57610f09611e52565b5b60200260200101819052505b8080610f2190611e81565b915050610e0d565b508091505090565b60008311610f74576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f6b90611f3b565b60405180910390fd5b60008211610fb7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fae90611fcd565b60405180910390fd5b60008111610ffa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ff19061205f565b60405180910390fd5b600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002060009054906101000a900460ff1661112c57600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208390806001815401808255809150506001900390600052602060002001600090919091909150556001600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002060006101000a81548160ff0219169083151502179055505b6040518060c0016040528060015481526020018481526020013373ffffffffffffffffffffffffffffffffffffffff16815260200183815260200182815260200160011515815250600260006001548152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301556080820151816004015560a08201518160050160006101000a81548160ff0219169083151502179055509050506003600084815260200190815260200160002060015490806001815401808255809150506001900390600052602060002001600090919091909150556001600081548092919061126a90611e81565b9190505550505050565b6060600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156112ff57602002820191906000526020600020905b8154815260200190600101908083116112eb575b50505050509050919050565b60056020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b60026020528060005260406000206000915090508060000154908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154908060040154908060050160009054906101000a900460ff16905086565b600460205281600052604060002081815481106113bf57600080fd5b90600052602060002001600091509150505481565b6040518060c001604052806000815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081526020016000151581525090565b600080fd5b6000819050919050565b61143a81611427565b811461144557600080fd5b50565b60008135905061145781611431565b92915050565b6000806040838503121561147457611473611422565b5b600061148285828601611448565b925050602061149385828601611448565b9150509250929050565b6114a681611427565b82525050565b60006020820190506114c1600083018461149d565b92915050565b6000602082840312156114dd576114dc611422565b5b60006114eb84828501611448565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061153961153461152f846114f4565b611514565b6114f4565b9050919050565b600061154b8261151e565b9050919050565b600061155d82611540565b9050919050565b61156d81611552565b82525050565b60006020820190506115886000830184611564565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6115c381611427565b82525050565b60006115d4826114f4565b9050919050565b6115e4816115c9565b82525050565b60008115159050919050565b6115ff816115ea565b82525050565b60c08201600082015161161b60008501826115ba565b50602082015161162e60208501826115ba565b50604082015161164160408501826115db565b50606082015161165460608501826115ba565b50608082015161166760808501826115ba565b5060a082015161167a60a08501826115f6565b50505050565b600061168c8383611605565b60c08301905092915050565b6000602082019050919050565b60006116b08261158e565b6116ba8185611599565b93506116c5836115aa565b8060005b838110156116f65781516116dd8882611680565b97506116e883611698565b9250506001810190506116c9565b5085935050505092915050565b6000602082019050818103600083015261171d81846116a5565b905092915050565b61172e816115c9565b811461173957600080fd5b50565b60008135905061174b81611725565b92915050565b60006020828403121561176757611766611422565b5b60006117758482850161173c565b91505092915050565b60008060006060848603121561179757611796611422565b5b60006117a586828701611448565b93505060206117b686828701611448565b92505060406117c786828701611448565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600061180983836115ba565b60208301905092915050565b6000602082019050919050565b600061182d826117d1565b61183781856117dc565b9350611842836117ed565b8060005b8381101561187357815161185a88826117fd565b975061186583611815565b925050600181019050611846565b5085935050505092915050565b6000602082019050818103600083015261189a8184611822565b905092915050565b600080604083850312156118b9576118b8611422565b5b60006118c78582860161173c565b92505060206118d885828601611448565b9150509250929050565b6118eb816115ea565b82525050565b600060208201905061190660008301846118e2565b92915050565b611915816115c9565b82525050565b600060c082019050611930600083018961149d565b61193d602083018861149d565b61194a604083018761190c565b611957606083018661149d565b611964608083018561149d565b61197160a08301846118e2565b979650505050505050565b600082825260208201905092915050565b7f54686973206974656d206973206e6f7420666f722073616c6500000000000000600082015250565b60006119c360198361197c565b91506119ce8261198d565b602082019050919050565b600060208201905081810360008301526119f2816119b6565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611a3382611427565b9150611a3e83611427565b9250828202611a4c81611427565b91508282048414831517611a6357611a626119f9565b5b5092915050565b7f53656e742076616c7565206973206c657373207468616e2074686520746f746160008201527f6c20636f73740000000000000000000000000000000000000000000000000000602082015250565b6000611ac660268361197c565b9150611ad182611a6a565b604082019050919050565b60006020820190508181036000830152611af581611ab9565b9050919050565b6000604082019050611b11600083018561190c565b611b1e602083018461149d565b9392505050565b600081519050611b3481611431565b92915050565b600060208284031215611b5057611b4f611422565b5b6000611b5e84828501611b25565b91505092915050565b7f53656c6c657220646f6573206e6f74206861766520656e6f75676820746f6b6560008201527f6e7320666f722073616c65000000000000000000000000000000000000000000602082015250565b6000611bc3602b8361197c565b9150611bce82611b67565b604082019050919050565b60006020820190508181036000830152611bf281611bb6565b9050919050565b600082825260208201905092915050565b50565b6000611c1a600083611bf9565b9150611c2582611c0a565b600082019050919050565b600060a082019050611c45600083018761190c565b611c52602083018661190c565b611c5f604083018561149d565b611c6c606083018461149d565b8181036080830152611c7d81611c0d565b905095945050505050565b600081905092915050565b6000611ca0600083611c88565b9150611cab82611c0a565b600082019050919050565b6000611cc182611c93565b9150819050919050565b7f4661696c656420746f207472616e7366657220457468657220746f207468652060008201527f73656c6c65720000000000000000000000000000000000000000000000000000602082015250565b6000611d2760268361197c565b9150611d3282611ccb565b604082019050919050565b60006020820190508181036000830152611d5681611d1a565b9050919050565b6000611d6882611427565b9150611d7383611427565b9250828203905081811115611d8b57611d8a6119f9565b5b92915050565b7f4f6e6c79207468652073656c6c65722063616e2072656d6f766520746865206c60008201527f697374696e670000000000000000000000000000000000000000000000000000602082015250565b6000611ded60268361197c565b9150611df882611d91565b604082019050919050565b60006020820190508181036000830152611e1c81611de0565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000611e8c82611427565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611ebe57611ebd6119f9565b5b600182019050919050565b7f546f6b656e2049442073686f756c642062652067726561746572207468616e2060008201527f7a65726f00000000000000000000000000000000000000000000000000000000602082015250565b6000611f2560248361197c565b9150611f3082611ec9565b604082019050919050565b60006020820190508181036000830152611f5481611f18565b9050919050565b7f50726963652073686f756c642062652067726561746572207468616e207a657260008201527f6f00000000000000000000000000000000000000000000000000000000000000602082015250565b6000611fb760218361197c565b9150611fc282611f5b565b604082019050919050565b60006020820190508181036000830152611fe681611faa565b9050919050565b7f416d6f756e742073686f756c642062652067726561746572207468616e207a6560008201527f726f000000000000000000000000000000000000000000000000000000000000602082015250565b600061204960228361197c565b915061205482611fed565b604082019050919050565b600060208201905081810360008301526120788161203c565b905091905056fea264697066735822122042f70ac84f19ef3d1b0cf7156050d82e8d0ad38eee52c239689865be8c5846c864736f6c63430008120033",
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

// GetListingsByToken is a free data retrieval call binding the contract method 0x63a2e3cc.
//
// Solidity: function getListingsByToken(uint256 tokenId) view returns((uint256,uint256,address,uint256,uint256,bool)[])
func (_Sale *SaleCaller) GetListingsByToken(opts *bind.CallOpts, tokenId *big.Int) ([]SaleListing, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "getListingsByToken", tokenId)

	if err != nil {
		return *new([]SaleListing), err
	}

	out0 := *abi.ConvertType(out[0], new([]SaleListing)).(*[]SaleListing)

	return out0, err

}

// GetListingsByToken is a free data retrieval call binding the contract method 0x63a2e3cc.
//
// Solidity: function getListingsByToken(uint256 tokenId) view returns((uint256,uint256,address,uint256,uint256,bool)[])
func (_Sale *SaleSession) GetListingsByToken(tokenId *big.Int) ([]SaleListing, error) {
	return _Sale.Contract.GetListingsByToken(&_Sale.CallOpts, tokenId)
}

// GetListingsByToken is a free data retrieval call binding the contract method 0x63a2e3cc.
//
// Solidity: function getListingsByToken(uint256 tokenId) view returns((uint256,uint256,address,uint256,uint256,bool)[])
func (_Sale *SaleCallerSession) GetListingsByToken(tokenId *big.Int) ([]SaleListing, error) {
	return _Sale.Contract.GetListingsByToken(&_Sale.CallOpts, tokenId)
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

// HasToken is a free data retrieval call binding the contract method 0xde61a0df.
//
// Solidity: function hasToken(address , uint256 ) view returns(bool)
func (_Sale *SaleCaller) HasToken(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "hasToken", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasToken is a free data retrieval call binding the contract method 0xde61a0df.
//
// Solidity: function hasToken(address , uint256 ) view returns(bool)
func (_Sale *SaleSession) HasToken(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Sale.Contract.HasToken(&_Sale.CallOpts, arg0, arg1)
}

// HasToken is a free data retrieval call binding the contract method 0xde61a0df.
//
// Solidity: function hasToken(address , uint256 ) view returns(bool)
func (_Sale *SaleCallerSession) HasToken(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Sale.Contract.HasToken(&_Sale.CallOpts, arg0, arg1)
}

// ListingCount is a free data retrieval call binding the contract method 0xa9b07c26.
//
// Solidity: function listingCount() view returns(uint256)
func (_Sale *SaleCaller) ListingCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "listingCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ListingCount is a free data retrieval call binding the contract method 0xa9b07c26.
//
// Solidity: function listingCount() view returns(uint256)
func (_Sale *SaleSession) ListingCount() (*big.Int, error) {
	return _Sale.Contract.ListingCount(&_Sale.CallOpts)
}

// ListingCount is a free data retrieval call binding the contract method 0xa9b07c26.
//
// Solidity: function listingCount() view returns(uint256)
func (_Sale *SaleCallerSession) ListingCount() (*big.Int, error) {
	return _Sale.Contract.ListingCount(&_Sale.CallOpts)
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

// OwnedTokens is a free data retrieval call binding the contract method 0xe149f036.
//
// Solidity: function ownedTokens(address , uint256 ) view returns(uint256)
func (_Sale *SaleCaller) OwnedTokens(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "ownedTokens", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnedTokens is a free data retrieval call binding the contract method 0xe149f036.
//
// Solidity: function ownedTokens(address , uint256 ) view returns(uint256)
func (_Sale *SaleSession) OwnedTokens(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Sale.Contract.OwnedTokens(&_Sale.CallOpts, arg0, arg1)
}

// OwnedTokens is a free data retrieval call binding the contract method 0xe149f036.
//
// Solidity: function ownedTokens(address , uint256 ) view returns(uint256)
func (_Sale *SaleCallerSession) OwnedTokens(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Sale.Contract.OwnedTokens(&_Sale.CallOpts, arg0, arg1)
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

// TokenToListings is a free data retrieval call binding the contract method 0x1ca17eff.
//
// Solidity: function tokenToListings(uint256 , uint256 ) view returns(uint256)
func (_Sale *SaleCaller) TokenToListings(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "tokenToListings", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToListings is a free data retrieval call binding the contract method 0x1ca17eff.
//
// Solidity: function tokenToListings(uint256 , uint256 ) view returns(uint256)
func (_Sale *SaleSession) TokenToListings(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Sale.Contract.TokenToListings(&_Sale.CallOpts, arg0, arg1)
}

// TokenToListings is a free data retrieval call binding the contract method 0x1ca17eff.
//
// Solidity: function tokenToListings(uint256 , uint256 ) view returns(uint256)
func (_Sale *SaleCallerSession) TokenToListings(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Sale.Contract.TokenToListings(&_Sale.CallOpts, arg0, arg1)
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
