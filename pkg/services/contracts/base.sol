// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";

contract Base is ERC1155 {
    mapping(uint256 => address[]) private tokenOwners;
    mapping(address => uint256[]) private ownedTokens;
    mapping(uint256 => string) public tokenNames;
    uint256 private _currentTokenId = 0;
    mapping(uint256 => address) public originalCreators;
    mapping(uint256 => string) public ipfsPaths;
    mapping(address => uint256[]) private tokensCreatedBy;
    event Mint(address indexed from, uint256 tokenId);

    constructor() ERC1155("https://musichain.com/api/token/{id}.json") {}

    function mint(string memory tokenName, uint256 amount, string memory ipfsPath, bytes memory data) public {
        uint256 newTokenId = _currentTokenId++;
        _mint(msg.sender, newTokenId, amount, data);
        tokenNames[newTokenId] = tokenName;
        ipfsPaths[newTokenId] = ipfsPath;
        tokenOwners[newTokenId].push(msg.sender);
        originalCreators[newTokenId] = msg.sender;
        ownedTokens[msg.sender].push(newTokenId);
        tokensCreatedBy[msg.sender].push(newTokenId);
        emit Mint(msg.sender, newTokenId);
    }

    function getOwnerOfToken(uint256 tokenId) public view returns (address) {
        address[] memory owners = tokenOwners[tokenId];
        if (owners.length > 0) {
            return owners[owners.length - 1];
        } else {
            revert("Token does not exist");
        }
    }
    
    function getOwnedTokens(address owner) public view returns (uint256[] memory) {
        return ownedTokens[owner];
    }

    function getTokensCreatedBy(address owner) public view returns (uint256[] memory) {
        return tokensCreatedBy[owner];
    }
}