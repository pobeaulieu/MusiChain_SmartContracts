// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

//  This contract stores all the data related to a token
contract Metadata {
    // A lot of mapping to keep track of information based on tokenID
    mapping(uint256 => string) public tokenNames;
    mapping(uint256 => string) public ipfsPaths;
    mapping(uint256 => uint256) public ticketPools;
    mapping(uint256 => uint256) public divs;
    mapping(uint256 => address[]) public tokenOwners;
    mapping(uint256 => address) public originalCreators;

    // Those mappings helps to get information about token owned or created by a certain address
    mapping(address => uint256[]) public ownedTokens;
    mapping(address => uint256[]) public tokensCreatedBy;

    uint256[] public tokenIds;

    function setMetadata(uint256 tokenId, string memory tokenName, string memory ipfsPath, uint256 ticketPool, uint256 div, address sender) external {
        tokenNames[tokenId] = tokenName;
        ipfsPaths[tokenId] = ipfsPath;
        ticketPools[tokenId] = ticketPool;
        divs[tokenId] = div;
        tokenIds.push(tokenId);
        tokenOwners[tokenId].push(sender);
        originalCreators[tokenId] = sender;
        ownedTokens[sender].push(tokenId);
        tokensCreatedBy[sender].push(tokenId);
    }

    function getTokenIds() external view returns(uint256[] memory) {
        return tokenIds;
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

    function getOriginalCreator(uint256 tokenId) public view returns (address) {
        return originalCreators[tokenId];
    }

    function getTokenOwners(uint256 tokenId) public view returns (address[] memory) {
        return tokenOwners[tokenId];
    }

}