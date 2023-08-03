// SPDX-License-Identifier: MIT

pragma solidity ^0.8.7;

//importing a remote ABI for usage
import "@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";

// All the functions in the library are internal
library PriceConverter {
     // fetching the ETH/USD price ratio
    function getPrice() internal view returns (uint256){
        // we require an ABI
        // Address to provide the conversion value: 0x694AA1769357215DE4FAC081bf1f309aDC325306
        // Sepolia test network
        AggregatorV3Interface priceFeed = AggregatorV3Interface(0x694AA1769357215DE4FAC081bf1f309aDC325306);
        (,int256 price,,,) = priceFeed.latestRoundData();
        return uint256(price * 1e10); // to make it upto 8 decimal places
    }

    // checking the version of the AggregatorV3Interface being used........ not necessary to find though
    function getVersion() internal view returns (uint256) {
        AggregatorV3Interface priceFeed = AggregatorV3Interface(0x694AA1769357215DE4FAC081bf1f309aDC325306);
        return priceFeed.version();
    }

    // getting the total ETH amount in USD
    function getConversionRate(uint256 ethAmount) internal view returns(uint256) {
        uint256 ethPrice = getPrice();
        uint256 ethAmountUSD = ethPrice * ethAmount / 1e18;
        return ethAmountUSD;
    }
}