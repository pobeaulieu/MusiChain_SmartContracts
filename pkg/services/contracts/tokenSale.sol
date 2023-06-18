pragma solidity ^0.8.0;

import "C:/Users/eliot/Desktop/MGL850/MusiChain/node_modules/@openzeppelin/contracts/token/ERC1155/ERC1155.sol";

contract TokenSale {
    IERC1155 public tokenContract;
    address payable public admin;
    uint256 public tokenPrice;
    uint256 public tokensSold;

    event Sell(address _buyer, uint256 _amount);

    constructor(IERC1155 _tokenContract, uint256 _tokenPrice) {
        admin = payable(msg.sender);
        tokenContract = _tokenContract;
        tokenPrice = _tokenPrice;
    }

    function buyTokens(uint256 _tokenId, uint256 _numberOfTokens, bytes memory _data) public payable {
        require(msg.value == _numberOfTokens * tokenPrice, "You need to send the correct amount of Ether");
        require(tokenContract.balanceOf(address(this), _tokenId) >= _numberOfTokens, "Not enough tokens available");

        tokenContract.safeTransferFrom(address(this), msg.sender, _tokenId, _numberOfTokens, _data);

        tokensSold += _numberOfTokens;
        emit Sell(msg.sender, _numberOfTokens);
    }

    function endSale() public {
        require(msg.sender == admin, "You are not the admin");

        admin.transfer(address(this).balance);
    }
}