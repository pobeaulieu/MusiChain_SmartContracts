// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Metadata {
    mapping(uint256 => string) public tokenNames;
    mapping(uint256 => string) public ipfsPaths;
    mapping(uint256 => uint256) public ticketPools;
    mapping(uint256 => uint256) public divs;


    uint256[] public tokenIds;

    function setMetadata(uint256 tokenId, string memory tokenName, string memory ipfsPath, uint256 ticketPool, uint256 div) external {
        tokenNames[tokenId] = tokenName;
        ipfsPaths[tokenId] = ipfsPath;
        ticketPools[tokenId] = ticketPool;
        divs[tokenId] = div;
        tokenIds.push(tokenId);
    }

    function getTokenIds() external view returns(uint256[] memory) {
        return tokenIds;
    }
}