// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";

contract Sale {
    struct Listing {
        uint256 tokenId;
        address seller;
        uint256 price;
        uint256 amount;
        bool isForSale;
    }

    IERC1155 public tokenContract;
    uint256[] public tokenIds;
    mapping(uint256 => Listing) public listings;

    constructor(IERC1155 _tokenContract) {
        tokenContract = _tokenContract;
    }

    function convertToWei(uint256 etherValue) public pure returns(uint256) {
        return etherValue * 10**18;
    }

    function listToken(uint256 tokenId, uint256 priceInEther, uint256 amount) public {
        require(tokenId > 0, "Token ID should be greater than zero");
        require(priceInEther > 0, "Price should be greater than zero");
        require(amount > 0, "Amount should be greater than zero");

        uint256 priceInWei = convertToWei(priceInEther);

        listings[tokenId] = Listing({
            tokenId: tokenId,
            seller: msg.sender,
            price: priceInWei,
            amount: amount,
            isForSale: true
        });

        tokenIds.push(tokenId);
    }

    function buyToken(uint256 tokenId) public payable {
        require(listings[tokenId].isForSale, "This item is not for sale");
        Listing memory listing = listings[tokenId];
        require(msg.value >= listing.price, "Sent value is less than the listing price");
        require(tokenContract.balanceOf(listing.seller, listing.tokenId) >= listing.amount, "Seller does not have enough tokens for sale");

        tokenContract.safeTransferFrom(listing.seller, msg.sender, listing.tokenId, listing.amount, "");

        (bool success, ) = payable(listing.seller).call{value: msg.value}("");
        require(success, "Failed to transfer Ether to the seller");

        listings[tokenId].isForSale = false;
    }

    function getAllListings() public view returns (Listing[] memory) {
        uint256 count = 0;
        for (uint256 i = 0; i < tokenIds.length; i++) {
            if (listings[tokenIds[i]].isForSale) {
                count++;
            }
        }

        Listing[] memory activeListings = new Listing[](count);
        uint256 index = 0;
        for (uint256 i = 0; i < tokenIds.length; i++) {
            if (listings[tokenIds[i]].isForSale) {
                activeListings[index] = listings[tokenIds[i]];
                index++;
            }
        }

        return activeListings;
    }

    function removeListing(uint256 tokenId) public {
        require(listings[tokenId].seller == msg.sender, "Only the seller can remove the listing");
        listings[tokenId].isForSale = false;
    }

    function getListingsByUser(address user) public view returns (Listing[] memory) {
        uint count = 0;
        for (uint i = 0; i < tokenIds.length; i++) {
            if (listings[tokenIds[i]].seller == user && listings[tokenIds[i]].isForSale) {
                count++;
            }
        }

        Listing[] memory userlistings = new Listing[](count);
        uint index = 0;
        for (uint i = 0; i < tokenIds.length; i++) {
            if (listings[tokenIds[i]].seller == user && listings[tokenIds[i]].isForSale) {
                userlistings[index] = listings[tokenIds[i]];
                index++;
            }
        }

        return userlistings;
    }

    function getOwnedTokens(address owner) public view returns (uint256[] memory) {
        uint256 count = 0;
        for (uint256 i = 0; i < tokenIds.length; i++) {
            if (tokenContract.balanceOf(owner, tokenIds[i]) > 0) {
                count++;
            }
        }

        uint256[] memory ownerTokens = new uint256[](count);
        uint256 index = 0;
        for (uint256 i = 0; i < tokenIds.length; i++) {
            if (tokenContract.balanceOf(owner, tokenIds[i]) > 0) {
                ownerTokens[index] = tokenIds[i];
                index++;
            }
        }

        return ownerTokens;
    }
}