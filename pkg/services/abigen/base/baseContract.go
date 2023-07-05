// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package base

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

// BaseMetaData contains all meta data concerning the Base contract.
var BaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405180606001604052806029815260200162002baa602991396200003d816200004460201b60201c565b50620003ba565b8060029081620000559190620002d3565b5050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620000db57607f821691505b602082108103620000f157620000f062000093565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026200015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200011c565b6200016786836200011c565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620001b4620001ae620001a8846200017f565b62000189565b6200017f565b9050919050565b6000819050919050565b620001d08362000193565b620001e8620001df82620001bb565b84845462000129565b825550505050565b600090565b620001ff620001f0565b6200020c818484620001c5565b505050565b5b81811015620002345762000228600082620001f5565b60018101905062000212565b5050565b601f82111562000283576200024d81620000f7565b62000258846200010c565b8101602085101562000268578190505b6200028062000277856200010c565b83018262000211565b50505b505050565b600082821c905092915050565b6000620002a86000198460080262000288565b1980831691505092915050565b6000620002c3838362000295565b9150826002028217905092915050565b620002de8262000059565b67ffffffffffffffff811115620002fa57620002f962000064565b5b620003068254620000c2565b6200031382828562000238565b600060209050601f8311600181146200034b576000841562000336578287015190505b620003428582620002b5565b865550620003b2565b601f1984166200035b86620000f7565b60005b8281101562000385578489015182556001820191506020850194506020810190506200035e565b86831015620003a55784890151620003a1601f89168262000295565b8355505b6001600288020188555050505b505050505050565b6127e080620003ca6000396000f3fe608060405234801561001057600080fd5b50600436106100925760003560e01c80632eb2c2d6116100665780632eb2c2d6146101435780634e1273f41461015f578063a22cb4651461018f578063e985e9c5146101ab578063f242432a146101db57610092565b8062fdd58e1461009757806301ffc9a7146100c757806308dc9f42146100f75780630e89341c14610113575b600080fd5b6100b160048036038101906100ac9190611598565b6101f7565b6040516100be91906115e7565b60405180910390f35b6100e160048036038101906100dc919061165a565b6102bf565b6040516100ee91906116a2565b60405180910390f35b610111600480360381019061010c9190611803565b6103a1565b005b61012d60048036038101906101289190611872565b6103b2565b60405161013a919061191e565b60405180910390f35b61015d60048036038101906101589190611a08565b610446565b005b61017960048036038101906101749190611b9a565b6104e7565b6040516101869190611cd0565b60405180910390f35b6101a960048036038101906101a49190611d1e565b610600565b005b6101c560048036038101906101c09190611d5e565b610616565b6040516101d291906116a2565b60405180910390f35b6101f560048036038101906101f09190611d9e565b6106aa565b005b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610267576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025e90611ea7565b60405180910390fd5b60008083815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b60007fd9b67a26000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061038a57507f0e89341c000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b8061039a57506103998261074b565b5b9050919050565b6103ad338484846107b5565b505050565b6060600280546103c190611ef6565b80601f01602080910402602001604051908101604052809291908181526020018280546103ed90611ef6565b801561043a5780601f1061040f5761010080835404028352916020019161043a565b820191906000526020600020905b81548152906001019060200180831161041d57829003601f168201915b50505050509050919050565b61044e610965565b73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16148061049457506104938561048e610965565b610616565b5b6104d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ca90611f99565b60405180910390fd5b6104e0858585858561096d565b5050505050565b6060815183511461052d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105249061202b565b60405180910390fd5b6000835167ffffffffffffffff81111561054a576105496116d8565b5b6040519080825280602002602001820160405280156105785781602001602082028036833780820191505090505b50905060005b84518110156105f5576105c585828151811061059d5761059c61204b565b5b60200260200101518583815181106105b8576105b761204b565b5b60200260200101516101f7565b8282815181106105d8576105d761204b565b5b602002602001018181525050806105ee906120a9565b905061057e565b508091505092915050565b61061261060b610965565b8383610c8e565b5050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6106b2610965565b73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614806106f857506106f7856106f2610965565b610616565b5b610737576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161072e90611f99565b60405180910390fd5b6107448585858585610dfa565b5050505050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610824576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161081b90612163565b60405180910390fd5b600061082e610965565b9050600061083b85611095565b9050600061084885611095565b90506108598360008985858961110f565b8460008088815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546108b89190612183565b925050819055508673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f6289896040516109369291906121b7565b60405180910390a461094d83600089858589611117565b61095c8360008989898961111f565b50505050505050565b600033905090565b81518351146109b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109a890612252565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610a20576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a17906122e4565b60405180910390fd5b6000610a2a610965565b9050610a3a81878787878761110f565b60005b8451811015610beb576000858281518110610a5b57610a5a61204b565b5b602002602001015190506000858381518110610a7a57610a7961204b565b5b60200260200101519050600080600084815260200190815260200160002060008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610b1b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b1290612376565b60405180910390fd5b81810360008085815260200190815260200160002060008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508160008085815260200190815260200160002060008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610bd09190612183565b9250508190555050505080610be4906120a9565b9050610a3d565b508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8787604051610c62929190612396565b60405180910390a4610c78818787878787611117565b610c868187878787876112f6565b505050505050565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610cfc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cf39061243f565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051610ded91906116a2565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610e69576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e60906122e4565b60405180910390fd5b6000610e73610965565b90506000610e8085611095565b90506000610e8d85611095565b9050610e9d83898985858961110f565b600080600088815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905085811015610f34576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f2b90612376565b60405180910390fd5b85810360008089815260200190815260200160002060008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508560008089815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610fe99190612183565b925050819055508773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f628a8a6040516110669291906121b7565b60405180910390a461107c848a8a86868a611117565b61108a848a8a8a8a8a61111f565b505050505050505050565b60606000600167ffffffffffffffff8111156110b4576110b36116d8565b5b6040519080825280602002602001820160405280156110e25781602001602082028036833780820191505090505b50905082816000815181106110fa576110f961204b565b5b60200260200101818152505080915050919050565b505050505050565b505050505050565b61113e8473ffffffffffffffffffffffffffffffffffffffff166114cd565b156112ee578373ffffffffffffffffffffffffffffffffffffffff1663f23a6e6187878686866040518663ffffffff1660e01b81526004016111849594939291906124c3565b6020604051808303816000875af19250505080156111c057506040513d601f19601f820116820180604052508101906111bd9190612532565b60015b611265576111cc61256c565b806308c379a00361122857506111e061258e565b806111eb575061122a565b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161121f919061191e565b60405180910390fd5b505b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161125c90612690565b60405180910390fd5b63f23a6e6160e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146112ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112e390612722565b60405180910390fd5b505b505050505050565b6113158473ffffffffffffffffffffffffffffffffffffffff166114cd565b156114c5578373ffffffffffffffffffffffffffffffffffffffff1663bc197c8187878686866040518663ffffffff1660e01b815260040161135b959493929190612742565b6020604051808303816000875af192505050801561139757506040513d601f19601f820116820180604052508101906113949190612532565b60015b61143c576113a361256c565b806308c379a0036113ff57506113b761258e565b806113c25750611401565b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113f6919061191e565b60405180910390fd5b505b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161143390612690565b60405180910390fd5b63bc197c8160e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146114c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114ba90612722565b60405180910390fd5b505b505050505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061152f82611504565b9050919050565b61153f81611524565b811461154a57600080fd5b50565b60008135905061155c81611536565b92915050565b6000819050919050565b61157581611562565b811461158057600080fd5b50565b6000813590506115928161156c565b92915050565b600080604083850312156115af576115ae6114fa565b5b60006115bd8582860161154d565b92505060206115ce85828601611583565b9150509250929050565b6115e181611562565b82525050565b60006020820190506115fc60008301846115d8565b92915050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61163781611602565b811461164257600080fd5b50565b6000813590506116548161162e565b92915050565b6000602082840312156116705761166f6114fa565b5b600061167e84828501611645565b91505092915050565b60008115159050919050565b61169c81611687565b82525050565b60006020820190506116b76000830184611693565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611710826116c7565b810181811067ffffffffffffffff8211171561172f5761172e6116d8565b5b80604052505050565b60006117426114f0565b905061174e8282611707565b919050565b600067ffffffffffffffff82111561176e5761176d6116d8565b5b611777826116c7565b9050602081019050919050565b82818337600083830152505050565b60006117a66117a184611753565b611738565b9050828152602081018484840111156117c2576117c16116c2565b5b6117cd848285611784565b509392505050565b600082601f8301126117ea576117e96116bd565b5b81356117fa848260208601611793565b91505092915050565b60008060006060848603121561181c5761181b6114fa565b5b600061182a86828701611583565b935050602061183b86828701611583565b925050604084013567ffffffffffffffff81111561185c5761185b6114ff565b5b611868868287016117d5565b9150509250925092565b600060208284031215611888576118876114fa565b5b600061189684828501611583565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156118d95780820151818401526020810190506118be565b60008484015250505050565b60006118f08261189f565b6118fa81856118aa565b935061190a8185602086016118bb565b611913816116c7565b840191505092915050565b6000602082019050818103600083015261193881846118e5565b905092915050565b600067ffffffffffffffff82111561195b5761195a6116d8565b5b602082029050602081019050919050565b600080fd5b600061198461197f84611940565b611738565b905080838252602082019050602084028301858111156119a7576119a661196c565b5b835b818110156119d057806119bc8882611583565b8452602084019350506020810190506119a9565b5050509392505050565b600082601f8301126119ef576119ee6116bd565b5b81356119ff848260208601611971565b91505092915050565b600080600080600060a08688031215611a2457611a236114fa565b5b6000611a328882890161154d565b9550506020611a438882890161154d565b945050604086013567ffffffffffffffff811115611a6457611a636114ff565b5b611a70888289016119da565b935050606086013567ffffffffffffffff811115611a9157611a906114ff565b5b611a9d888289016119da565b925050608086013567ffffffffffffffff811115611abe57611abd6114ff565b5b611aca888289016117d5565b9150509295509295909350565b600067ffffffffffffffff821115611af257611af16116d8565b5b602082029050602081019050919050565b6000611b16611b1184611ad7565b611738565b90508083825260208201905060208402830185811115611b3957611b3861196c565b5b835b81811015611b625780611b4e888261154d565b845260208401935050602081019050611b3b565b5050509392505050565b600082601f830112611b8157611b806116bd565b5b8135611b91848260208601611b03565b91505092915050565b60008060408385031215611bb157611bb06114fa565b5b600083013567ffffffffffffffff811115611bcf57611bce6114ff565b5b611bdb85828601611b6c565b925050602083013567ffffffffffffffff811115611bfc57611bfb6114ff565b5b611c08858286016119da565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b611c4781611562565b82525050565b6000611c598383611c3e565b60208301905092915050565b6000602082019050919050565b6000611c7d82611c12565b611c878185611c1d565b9350611c9283611c2e565b8060005b83811015611cc3578151611caa8882611c4d565b9750611cb583611c65565b925050600181019050611c96565b5085935050505092915050565b60006020820190508181036000830152611cea8184611c72565b905092915050565b611cfb81611687565b8114611d0657600080fd5b50565b600081359050611d1881611cf2565b92915050565b60008060408385031215611d3557611d346114fa565b5b6000611d438582860161154d565b9250506020611d5485828601611d09565b9150509250929050565b60008060408385031215611d7557611d746114fa565b5b6000611d838582860161154d565b9250506020611d948582860161154d565b9150509250929050565b600080600080600060a08688031215611dba57611db96114fa565b5b6000611dc88882890161154d565b9550506020611dd98882890161154d565b9450506040611dea88828901611583565b9350506060611dfb88828901611583565b925050608086013567ffffffffffffffff811115611e1c57611e1b6114ff565b5b611e28888289016117d5565b9150509295509295909350565b7f455243313135353a2061646472657373207a65726f206973206e6f742061207660008201527f616c6964206f776e657200000000000000000000000000000000000000000000602082015250565b6000611e91602a836118aa565b9150611e9c82611e35565b604082019050919050565b60006020820190508181036000830152611ec081611e84565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680611f0e57607f821691505b602082108103611f2157611f20611ec7565b5b50919050565b7f455243313135353a2063616c6c6572206973206e6f7420746f6b656e206f776e60008201527f6572206f7220617070726f766564000000000000000000000000000000000000602082015250565b6000611f83602e836118aa565b9150611f8e82611f27565b604082019050919050565b60006020820190508181036000830152611fb281611f76565b9050919050565b7f455243313135353a206163636f756e747320616e6420696473206c656e67746860008201527f206d69736d617463680000000000000000000000000000000000000000000000602082015250565b60006120156029836118aa565b915061202082611fb9565b604082019050919050565b6000602082019050818103600083015261204481612008565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006120b482611562565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036120e6576120e561207a565b5b600182019050919050565b7f455243313135353a206d696e7420746f20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b600061214d6021836118aa565b9150612158826120f1565b604082019050919050565b6000602082019050818103600083015261217c81612140565b9050919050565b600061218e82611562565b915061219983611562565b92508282019050808211156121b1576121b061207a565b5b92915050565b60006040820190506121cc60008301856115d8565b6121d960208301846115d8565b9392505050565b7f455243313135353a2069647320616e6420616d6f756e7473206c656e6774682060008201527f6d69736d61746368000000000000000000000000000000000000000000000000602082015250565b600061223c6028836118aa565b9150612247826121e0565b604082019050919050565b6000602082019050818103600083015261226b8161222f565b9050919050565b7f455243313135353a207472616e7366657220746f20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b60006122ce6025836118aa565b91506122d982612272565b604082019050919050565b600060208201905081810360008301526122fd816122c1565b9050919050565b7f455243313135353a20696e73756666696369656e742062616c616e636520666f60008201527f72207472616e7366657200000000000000000000000000000000000000000000602082015250565b6000612360602a836118aa565b915061236b82612304565b604082019050919050565b6000602082019050818103600083015261238f81612353565b9050919050565b600060408201905081810360008301526123b08185611c72565b905081810360208301526123c48184611c72565b90509392505050565b7f455243313135353a2073657474696e6720617070726f76616c2073746174757360008201527f20666f722073656c660000000000000000000000000000000000000000000000602082015250565b60006124296029836118aa565b9150612434826123cd565b604082019050919050565b600060208201905081810360008301526124588161241c565b9050919050565b61246881611524565b82525050565b600081519050919050565b600082825260208201905092915050565b60006124958261246e565b61249f8185612479565b93506124af8185602086016118bb565b6124b8816116c7565b840191505092915050565b600060a0820190506124d8600083018861245f565b6124e5602083018761245f565b6124f260408301866115d8565b6124ff60608301856115d8565b8181036080830152612511818461248a565b90509695505050505050565b60008151905061252c8161162e565b92915050565b600060208284031215612548576125476114fa565b5b60006125568482850161251d565b91505092915050565b60008160e01c9050919050565b600060033d111561258b5760046000803e61258860005161255f565b90505b90565b600060443d1061261b576125a06114f0565b60043d036004823e80513d602482011167ffffffffffffffff821117156125c857505061261b565b808201805167ffffffffffffffff8111156125e6575050505061261b565b80602083010160043d03850181111561260357505050505061261b565b61261282602001850186611707565b82955050505050505b90565b7f455243313135353a207472616e7366657220746f206e6f6e2d4552433131353560008201527f526563656976657220696d706c656d656e746572000000000000000000000000602082015250565b600061267a6034836118aa565b91506126858261261e565b604082019050919050565b600060208201905081810360008301526126a98161266d565b9050919050565b7f455243313135353a204552433131353552656365697665722072656a6563746560008201527f6420746f6b656e73000000000000000000000000000000000000000000000000602082015250565b600061270c6028836118aa565b9150612717826126b0565b604082019050919050565b6000602082019050818103600083015261273b816126ff565b9050919050565b600060a082019050612757600083018861245f565b612764602083018761245f565b81810360408301526127768186611c72565b9050818103606083015261278a8185611c72565b9050818103608083015261279e818461248a565b9050969550505050505056fea2646970667358221220f6b77c76dc014d845094c8f7b3f54344009d065ebd536b2df22c5fe79d5e0c5464736f6c6343000812003368747470733a2f2f6d757369636861696e2e636f6d2f6170692f746f6b656e2f7b69647d2e6a736f6e",
}

// BaseABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseMetaData.ABI instead.
var BaseABI = BaseMetaData.ABI

// BaseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BaseMetaData.Bin instead.
var BaseBin = BaseMetaData.Bin

// DeployBase deploys a new Ethereum contract, binding an instance of Base to it.
func DeployBase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Base, error) {
	parsed, err := BaseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Base{BaseCaller: BaseCaller{contract: contract}, BaseTransactor: BaseTransactor{contract: contract}, BaseFilterer: BaseFilterer{contract: contract}}, nil
}

// Base is an auto generated Go binding around an Ethereum contract.
type Base struct {
	BaseCaller     // Read-only binding to the contract
	BaseTransactor // Write-only binding to the contract
	BaseFilterer   // Log filterer for contract events
}

// BaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseSession struct {
	Contract     *Base             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseCallerSession struct {
	Contract *BaseCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseTransactorSession struct {
	Contract     *BaseTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseRaw struct {
	Contract *Base // Generic contract binding to access the raw methods on
}

// BaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseCallerRaw struct {
	Contract *BaseCaller // Generic read-only contract binding to access the raw methods on
}

// BaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseTransactorRaw struct {
	Contract *BaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBase creates a new instance of Base, bound to a specific deployed contract.
func NewBase(address common.Address, backend bind.ContractBackend) (*Base, error) {
	contract, err := bindBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Base{BaseCaller: BaseCaller{contract: contract}, BaseTransactor: BaseTransactor{contract: contract}, BaseFilterer: BaseFilterer{contract: contract}}, nil
}

// NewBaseCaller creates a new read-only instance of Base, bound to a specific deployed contract.
func NewBaseCaller(address common.Address, caller bind.ContractCaller) (*BaseCaller, error) {
	contract, err := bindBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseCaller{contract: contract}, nil
}

// NewBaseTransactor creates a new write-only instance of Base, bound to a specific deployed contract.
func NewBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseTransactor, error) {
	contract, err := bindBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseTransactor{contract: contract}, nil
}

// NewBaseFilterer creates a new log filterer instance of Base, bound to a specific deployed contract.
func NewBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseFilterer, error) {
	contract, err := bindBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseFilterer{contract: contract}, nil
}

// bindBase binds a generic wrapper to an already deployed contract.
func bindBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Base *BaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Base.Contract.BaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Base *BaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Base.Contract.BaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Base *BaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Base.Contract.BaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Base *BaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Base.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Base *BaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Base.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Base *BaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Base.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Base *BaseCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Base.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Base *BaseSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Base.Contract.BalanceOf(&_Base.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Base *BaseCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Base.Contract.BalanceOf(&_Base.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Base *BaseCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Base.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Base *BaseSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Base.Contract.BalanceOfBatch(&_Base.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Base *BaseCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Base.Contract.BalanceOfBatch(&_Base.CallOpts, accounts, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Base *BaseCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Base.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Base *BaseSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Base.Contract.IsApprovedForAll(&_Base.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Base *BaseCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Base.Contract.IsApprovedForAll(&_Base.CallOpts, account, operator)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Base *BaseCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Base.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Base *BaseSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Base.Contract.SupportsInterface(&_Base.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Base *BaseCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Base.Contract.SupportsInterface(&_Base.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Base *BaseCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Base.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Base *BaseSession) Uri(arg0 *big.Int) (string, error) {
	return _Base.Contract.Uri(&_Base.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Base *BaseCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _Base.Contract.Uri(&_Base.CallOpts, arg0)
}

// Mint is a paid mutator transaction binding the contract method 0x08dc9f42.
//
// Solidity: function mint(uint256 tokenId, uint256 amount, bytes data) returns()
func (_Base *BaseTransactor) Mint(opts *bind.TransactOpts, tokenId *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Base.contract.Transact(opts, "mint", tokenId, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x08dc9f42.
//
// Solidity: function mint(uint256 tokenId, uint256 amount, bytes data) returns()
func (_Base *BaseSession) Mint(tokenId *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Base.Contract.Mint(&_Base.TransactOpts, tokenId, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x08dc9f42.
//
// Solidity: function mint(uint256 tokenId, uint256 amount, bytes data) returns()
func (_Base *BaseTransactorSession) Mint(tokenId *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Base.Contract.Mint(&_Base.TransactOpts, tokenId, amount, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Base *BaseTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Base.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Base *BaseSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Base.Contract.SafeBatchTransferFrom(&_Base.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Base *BaseTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Base.Contract.SafeBatchTransferFrom(&_Base.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Base *BaseTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Base.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Base *BaseSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Base.Contract.SafeTransferFrom(&_Base.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Base *BaseTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Base.Contract.SafeTransferFrom(&_Base.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Base *BaseTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Base.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Base *BaseSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Base.Contract.SetApprovalForAll(&_Base.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Base *BaseTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Base.Contract.SetApprovalForAll(&_Base.TransactOpts, operator, approved)
}

// BaseApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Base contract.
type BaseApprovalForAllIterator struct {
	Event *BaseApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaseApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BaseApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BaseApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseApprovalForAll represents a ApprovalForAll event raised by the Base contract.
type BaseApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Base *BaseFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*BaseApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Base.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BaseApprovalForAllIterator{contract: _Base.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Base *BaseFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BaseApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Base.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseApprovalForAll)
				if err := _Base.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Base *BaseFilterer) ParseApprovalForAll(log types.Log) (*BaseApprovalForAll, error) {
	event := new(BaseApprovalForAll)
	if err := _Base.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Base contract.
type BaseTransferBatchIterator struct {
	Event *BaseTransferBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaseTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseTransferBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BaseTransferBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BaseTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseTransferBatch represents a TransferBatch event raised by the Base contract.
type BaseTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Base *BaseFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*BaseTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Base.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BaseTransferBatchIterator{contract: _Base.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Base *BaseFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *BaseTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Base.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseTransferBatch)
				if err := _Base.contract.UnpackLog(event, "TransferBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Base *BaseFilterer) ParseTransferBatch(log types.Log) (*BaseTransferBatch, error) {
	event := new(BaseTransferBatch)
	if err := _Base.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Base contract.
type BaseTransferSingleIterator struct {
	Event *BaseTransferSingle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaseTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseTransferSingle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BaseTransferSingle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BaseTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseTransferSingle represents a TransferSingle event raised by the Base contract.
type BaseTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Base *BaseFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*BaseTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Base.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BaseTransferSingleIterator{contract: _Base.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Base *BaseFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *BaseTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Base.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseTransferSingle)
				if err := _Base.contract.UnpackLog(event, "TransferSingle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Base *BaseFilterer) ParseTransferSingle(log types.Log) (*BaseTransferSingle, error) {
	event := new(BaseTransferSingle)
	if err := _Base.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Base contract.
type BaseURIIterator struct {
	Event *BaseURI // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaseURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseURI)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BaseURI)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BaseURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseURI represents a URI event raised by the Base contract.
type BaseURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Base *BaseFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*BaseURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Base.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &BaseURIIterator{contract: _Base.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Base *BaseFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *BaseURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Base.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseURI)
				if err := _Base.contract.UnpackLog(event, "URI", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Base *BaseFilterer) ParseURI(log types.Log) (*BaseURI, error) {
	event := new(BaseURI)
	if err := _Base.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
