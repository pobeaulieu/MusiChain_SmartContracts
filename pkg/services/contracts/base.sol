// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
// Remove /node_modules when building on Remix.
import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";


contract Base is ERC1155 {

    constructor() ERC1155("https://musichain.com/api/token/{id}.json") {
    }

    function mint(uint256 tokenId, uint256 amount, bytes memory data) public {
        _mint(msg.sender, tokenId, amount, data);
    }

}