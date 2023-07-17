// SPDX-License-Identifier: MIT

pragma solidity ^0.8.7;

import "./1-SimpleStorage.sol";

contract ExtraStorage is SimpleStorage {
    function store(uint256 _favoriteNumber) public override {
        FavouriteNumber = _favoriteNumber + 5;
    }
}