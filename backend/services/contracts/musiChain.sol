// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

/**
 * @title MusiChain
 * @dev Store & retrieve value in a variable
 * @custom:dev-run-script ./scripts/deploy_with_ethers.ts
 */
contract MusiChain {

    uint256 number;
    uint256 tokenPrice;
    uint256 dividend;
    uint256 initialTicketPool;

    /**
     * @dev Store value in variable
     * @param num value to store
     * @param price of the token
     */
    function store(uint256 num, uint256 price, uint256 divi, uint256 initial) public {
        number = num;
        tokenPrice = price;
        dividend = divi;
        initialTicketPool = initial;
    }

    /**
     * @dev Return value
     * @return value of 'number'
     */
    function retrieve() public view returns (uint256, uint256, uint256, uint256){
        return (number, tokenPrice, dividend, initialTicketPool);
    }
}