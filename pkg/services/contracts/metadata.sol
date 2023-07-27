// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Metadata {
    mapping(uint256 => string) public tokenNames;
    mapping(uint256 => string) public ipfsPaths;

    uint256[] public tokenIds;

    function setMetadata(uint256 tokenId, string memory tokenName, string memory ipfsPath) external {
        tokenNames[tokenId] = tokenName;
        ipfsPaths[tokenId] = ipfsPath;
        tokenIds.push(tokenId);
    }

    function getTokenIds() external view returns(uint256[] memory) {
        return tokenIds;
    }
}