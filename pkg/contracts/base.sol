// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "./Metadata.sol";

// Contract representing the tokens created by the creator
contract Base is ERC1155 {
    // This metadata contract will be used to store information about the tokens
    Metadata private metadata;
    uint256 public _currentTokenId = 1;
    event Mint(address indexed from, uint256 tokenId);

    constructor(Metadata _metadata) ERC1155("https://musichain.com/api/token/{id}.json") {
        metadata = _metadata;
    }

    // The mint function used to create tokens
    function mint(string memory tokenName, uint256 amount, string memory ipfsPath, uint256 ticketPool, uint256 div, bytes memory data) public {
        uint256 newTokenId = _currentTokenId++;
        _mint(msg.sender, newTokenId, amount, data);
        metadata.setMetadata(newTokenId, tokenName, ipfsPath, ticketPool, div, msg.sender);
        emit Mint(msg.sender, newTokenId);
    }

    // This function provides a way to pay dividends to each owner of a token
    function payDividends(uint256 tokenId, uint256 amount) public payable {
        address originalCreator = metadata.getOriginalCreator(tokenId);
        address[] memory owners = metadata.getTokenOwners(tokenId);

        require(originalCreator == msg.sender, "Only the original creator can send Ether");
        require(owners.length > 0, "Token does not exist or has no owners");
        require(msg.value >= amount * (owners.length - 1), "Sent value is less than the required total");  // subtract 1 for original creator

        uint256 sentAmount = 0;
        for (uint256 i = 0; i < owners.length; i++) {
            if (owners[i] != originalCreator) {
                payable(owners[i]).transfer(amount);
                sentAmount += amount;
            }
        }

        payable(msg.sender).transfer(msg.value - sentAmount);
    }

    function getTokenBalance(address account, uint256 tokenId) public view returns (uint256) {
        return balanceOf(account, tokenId);
    }

}