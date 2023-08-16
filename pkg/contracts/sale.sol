// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";
import "./Metadata.sol";

// This contract aims to provide a way to list token and to buy them
contract Sale {
    // This is the struct used to represent a listing
    struct Listing {
        uint256 listingId;
        uint256 tokenId;
        address seller;
        uint256 price;
        uint256 amount;
        uint256 orignalAmount;
        bool isForSale;
    }

    uint256 public currentListingId=1;
    // The baseContract is used to retrieve information about the tokens and transfer them
    IERC1155 public baseContract;
    // This mapping store the listing we create
    mapping(uint256 => Listing) public listings;

    // The Metadacontract is used to change the ownership of tokens
    Metadata public metadataContract;

    constructor(IERC1155 _baseContract, Metadata _metadataContract) {
        baseContract = _baseContract;
        metadataContract = _metadataContract;
    }

    // This function creates a listing
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
            orignalAmount : amount,
            isForSale: true
        });
    }

    // This function provide a way for a user to buy a token
    function buyToken(uint256 listingId, uint256 buyAmount) public payable {
        require(listings[listingId].isForSale, "This item is not for sale");
        Listing memory listing = listings[listingId];
        uint256 totalCost = listing.price * buyAmount;
        require(msg.value >= totalCost, "Sent value is less than the total cost");
        require(baseContract.balanceOf(listing.seller, listing.tokenId) >= buyAmount, "Seller does not have enough tokens for sale");

        // This function from the ERC1155 standard provides a way to transfer a certain amount of a token from an address
        // to another address
        baseContract.safeTransferFrom(listing.seller, msg.sender, listing.tokenId, buyAmount, "");

        // We update the mapping of the owner
        metadataContract.addTokenOwner(listing.tokenId, msg.sender);

        (bool success, ) = payable(listing.seller).call{value: totalCost}("");
        require(success, "Failed to transfer Ether to the seller");

        listings[listingId].amount -= buyAmount;

        if(listings[listingId].amount == 0) {
            listings[listingId].isForSale = false;
        }
    }

    function getAllListings() public view returns (Listing[] memory) {
        uint256 count = 0;
        for (uint256 i = 0; i < currentListingId; i++) {
            if (listings[i].isForSale) {
                count++;
            }
        }

        Listing[] memory activeListings = new Listing[](count);
        uint256 index = 0;
        for (uint256 i = 0; i < currentListingId; i++) {
            if (listings[i].isForSale) {
                activeListings[index] = listings[i];
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
        uint256 count = 0;
        for (uint256 i = 0; i < currentListingId; i++) {
            if (listings[i].seller == user && listings[i].isForSale) {
                count++;
            }
        }

        Listing[] memory userlistings = new Listing[](count);
        uint256 index = 0;
        for (uint256 i = 0; i < currentListingId; i++) {
            if (listings[i].seller == user && listings[i].isForSale) {
                userlistings[index] = listings[i];
                index++;
            }
        }

        return userlistings;
    }

    function getOwnedTokens(address owner) public view returns (uint256[] memory) {
        uint256 count = 0;
        for (uint256 i = 0; i < currentListingId; i++) {
            if (baseContract.balanceOf(owner, listings[i].tokenId) > 0) {
                count++;
            }
        }

        uint256[] memory ownerTokens = new uint256[](count);
        uint256 index = 0;
        for (uint256 i = 0; i < currentListingId; i++) {
            if (baseContract.balanceOf(owner, listings[i].tokenId) > 0) {
                ownerTokens[index] = listings[i].tokenId;
                index++;
            }
        }

        return ownerTokens;
    }
}