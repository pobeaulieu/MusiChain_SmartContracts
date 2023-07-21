// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";

contract Base is ERC1155 {
    mapping(uint256 => address[]) private tokenOwners;
    mapping(address => uint256[]) private ownedTokens;
    mapping(uint256 => string) public tokenNames;  // change to public
    uint256 private _currentTokenId = 0;
    mapping(uint256 => address) public originalCreators;  // change to public

    event Mint(address indexed from, uint256 tokenId);

    constructor() ERC1155("https://musichain.com/api/token/{id}.json") {}

    function mint(string memory tokenName, uint256 amount, bytes memory data) public {
        uint256 newTokenId = _currentTokenId++;
        _mint(msg.sender, newTokenId, amount, data);
        tokenNames[newTokenId] = tokenName;
        tokenOwners[newTokenId].push(msg.sender);
        originalCreators[newTokenId] = msg.sender;
        ownedTokens[msg.sender].push(newTokenId);
        emit Mint(msg.sender, newTokenId);
    }

    function balanceOfToken(uint256 tokenId) public view returns (uint256) {
        return tokenOwners[tokenId].length;
    }

    function getOwnerAddresses(uint256 tokenId) public view returns (address[] memory) {
        return tokenOwners[tokenId];
    }

    function getOwnerOfToken(uint256 tokenId) public view returns (address) {
        address[] memory owners = getOwnerAddresses(tokenId);
        if (owners.length > 0) {
            return owners[owners.length - 1];
        } else {
            revert("Token does not exist");
        }
    }

    function getAllTokenOwners(uint256 tokenId) public view returns (address[] memory) {
        return tokenOwners[tokenId];
    }

    function getOwnedTokens(address owner) public view returns (uint256[] memory) {
        return ownedTokens[owner];
    }
}