// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "C:/Users/eliot/Desktop/MGL850/MusiChain/node_modules/@openzeppelin/contracts/token/ERC1155/ERC1155.sol";


contract MusiChain is ERC1155 {

    mapping(address => uint256) public revenue;

    constructor() ERC1155("https://myapi.com/api/token/{id}.json") {
    }

    function mint(address to, uint256 tokenId, uint256 amount, bytes memory data) public {
        _mint(to, tokenId, amount, data);
    }

    function distributeRevenue(address recipient, uint256 amount) public {
        revenue[recipient] += amount;
    }

    function getRevenue(address recipient) public view returns (uint256) {
        return revenue[recipient];
    }
}