// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
// Remove /node_modules when building on Remix.
import "node_modules/@openzeppelin/contracts/token/ERC1155/ERC1155.sol";


contract Base is ERC1155 {

    mapping(address => uint256) public revenue;

    constructor() ERC1155("https://musichain.com/api/token/{id}.json") {
    }

    function mint(address to, uint256 tokenId, uint256 amount, bytes memory data) public {
        _mint(to, tokenId, amount, data);
    }

}