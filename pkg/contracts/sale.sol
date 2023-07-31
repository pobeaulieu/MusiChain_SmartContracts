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

    IERC1155 public tokenContract;
    uint256 public listingCount;
    mapping(uint256 => Listing) public listings;
    mapping(uint256 => uint256[]) public tokenToListings;
    mapping (address => uint256[]) public ownedTokens;
    mapping (address => mapping(uint256 => bool)) public hasToken;

    constructor(IERC1155 _tokenContract) {
        tokenContract = _tokenContract;
    }

    function listToken(uint256 tokenId, uint256 priceInWei, uint256 amount) public {
        require(tokenId > 0, "Token ID should be greater than zero");
        require(priceInWei > 0, "Price should be greater than zero");
        require(amount > 0, "Amount should be greater than zero");

        if (!hasToken[msg.sender][tokenId]) {
            ownedTokens[msg.sender].push(tokenId);
            hasToken[msg.sender][tokenId] = true;
        }

        listings[listingCount] = Listing({
            listingId: listingCount,
            tokenId: tokenId,
            seller: msg.sender,
            price: priceInWei,
            amount: amount,
            isForSale: true
        });

        tokenToListings[tokenId].push(listingCount);
        listingCount++;
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

        ownedTokens[msg.sender].push(listing.tokenId);

        if(listings[listingId].amount == 0) {
            listings[listingId].isForSale = false;
        }
    }

    function getAllListings() public view returns (Listing[] memory) {
        Listing[] memory activeListings = new Listing[](listingCount);
        for (uint256 i = 0; i < listingCount; i++) {
            if (listings[i].isForSale) {
                activeListings[i] = listings[i];
            }
        }
        return activeListings;
    }

    function removeListing(uint256 listingId) public {
        require(listings[listingId].seller == msg.sender, "Only the seller can remove the listing");
        listings[listingId].isForSale = false;
    }

    function getListingsByUser(address user) public view returns (Listing[] memory) {
        uint256[] memory userListingIds = new uint256[](listingCount);
        uint256 count = 0;
        for (uint256 i = 0; i < listingCount; i++) {
            if (listings[i].seller == user && listings[i].isForSale) {
                userListingIds[count] = i;
                count++;
            }
        }

        Listing[] memory userlistings = new Listing[](count);
        for (uint256 i = 0; i < count; i++) {
            userlistings[i] = listings[userListingIds[i]];
        }

        return userlistings;
    }

    function getListingsByToken(uint256 tokenId) public view returns (Listing[] memory) {
        uint256[] memory listingIds = tokenToListings[tokenId];
        uint256 count = listingIds.length;

        Listing[] memory tokenListings = new Listing[](count);
        for (uint256 i = 0; i < count; i++) {
            tokenListings[i] = listings[listingIds[i]];
        }
        return tokenListings;
    }

    function getOwnedTokens(address owner) public view returns (uint256[] memory) {
        return ownedTokens[owner];
    }
}