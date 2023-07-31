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

    function listToken(uint256 tokenId, uint256 priceInWei, uint256 amount) public {
        require(tokenId > 0, "Token ID should be greater than zero");
        require(priceInWei > 0, "Price should be greater than zero");
        require(amount > 0, "Amount should be greater than zero");

        listings[tokenId] = Listing({
            tokenId: tokenId,
            seller: msg.sender,
            price: priceInWei,
            amount: amount,
            isForSale: true
        });

        tokenIds.push(tokenId);
    }

    function buyToken(uint256 tokenId, uint256 buyAmount) public payable {
        require(listings[tokenId].isForSale, "This item is not for sale");
        Listing memory listing = listings[tokenId];
        uint256 totalCost = listing.price * buyAmount;
        require(msg.value >= totalCost, "Sent value is less than the total cost");
        require(tokenContract.balanceOf(listing.seller, listing.tokenId) >= buyAmount, "Seller does not have enough tokens for sale");

        tokenContract.safeTransferFrom(listing.seller, msg.sender, listing.tokenId, buyAmount, "");

        (bool success, ) = payable(listing.seller).call{value: totalCost}("");
        require(success, "Failed to transfer Ether to the seller");

        listings[tokenId].amount -= buyAmount;

        if(listings[tokenId].amount == 0) {
            listings[tokenId].isForSale = false;
        }
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