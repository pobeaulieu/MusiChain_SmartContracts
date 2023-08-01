// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";

contract Sale {
    struct Listing {
        uint256 listingId;
        uint256 tokenId;
        address seller;
        uint256 price;
        uint256 amount;
        bool isForSale;
    }

    uint256 public currentListingId=1;
    IERC1155 public tokenContract;
    uint256[] public listingIds;
    mapping(uint256 => Listing) public listings;

    constructor(IERC1155 _tokenContract) {
        tokenContract = _tokenContract;
    }

    function listToken(uint256 tokenId, uint256 priceInWei, uint256 amount) public {
        require(tokenId > 0, "Token ID should be greater than zero");
        require(priceInWei > 0, "Price should be greater than zero");
        require(amount > 0, "Amount should be greater than zero");
        uint256 newListingId = currentListingId++;
        listings[newListingId] = Listing({
            listingId : newListingId,
            tokenId: tokenId,
            seller: msg.sender,
            price: priceInWei,
            amount: amount,
            isForSale: true
        });

        listingIds.push(newListingId);

    }

    function buyToken(uint256 listingId, uint256 buyAmount) public payable {
        require(listings[listingId].isForSale, "This item is not for sale");
        Listing memory listing = listings[listingId];
        uint256 totalCost = listing.price * buyAmount;
        require(msg.value >= totalCost, "Sent value is less than the total cost");
        require(tokenContract.balanceOf(listing.seller, listing.tokenId) >= buyAmount, "Seller does not have enough tokens for sale");

        tokenContract.safeTransferFrom(listing.seller, msg.sender, listing.tokenId, buyAmount, "");

        (bool success, ) = payable(listing.seller).call{value: totalCost}("");
        require(success, "Failed to transfer Ether to the seller");

        listings[listingId].amount -= buyAmount;

        if(listings[listingId].amount == 0) {
            listings[listingId].isForSale = false;
        }
    }

    function getAllListings() public view returns (Listing[] memory) {
        uint256 count = 0;
        for (uint256 i = 0; i < listingIds.length; i++) {
            if (listings[listingIds[i]].isForSale) {
                count++;
            }
        }

        Listing[] memory activeListings = new Listing[](count);
        uint256 index = 0;
        for (uint256 i = 0; i < listingIds.length; i++) {
            if (listings[listingIds[i]].isForSale) {
                activeListings[index] = listings[listingIds[i]];
                index++;
            }
        }

        return activeListings;
    }

    function removeListing(uint256 listingId) public {
        require(listings[listingId].seller == msg.sender, "Only the seller can remove the listing");
        listings[listingId].isForSale = false;
    }

    function getListingsByUser(address user) public view returns (Listing[] memory) {
        uint count = 0;
        for (uint i = 0; i < listingIds.length; i++) {
            if (listings[listingIds[i]].seller == user && listings[listingIds[i]].isForSale) {
                count++;
            }
        }

        Listing[] memory userlistings = new Listing[](count);
        uint index = 0;
        for (uint i = 0; i < listingIds.length; i++) {
            if (listings[listingIds[i]].seller == user && listings[listingIds[i]].isForSale) {
                userlistings[index] = listings[listingIds[i]];
                index++;
            }
        }

        return userlistings;
    }

    function getOwnedTokens(address owner) public view returns (uint256[] memory) {
        uint256 count = 0;
        for (uint256 i = 0; i < listingIds.length; i++) {
            if (tokenContract.balanceOf(owner, listingIds[i]) > 0) {
                count++;
            }
        }

        uint256[] memory ownerTokens = new uint256[](count);
        uint256 index = 0;
        for (uint256 i = 0; i < listingIds.length; i++) {
            if (tokenContract.balanceOf(owner, listingIds[i]) > 0) {
                ownerTokens[index] = listingIds[i];
                index++;
            }
        }

        return ownerTokens;
    }
}