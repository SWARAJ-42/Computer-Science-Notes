// SPDX-License-Identifier: MIT
pragma solidity ^0.8.8;

import "FundMe/PriceConverter.sol";

contract FundMe {
    // mentioning that PriceConvertor library can be used on any variable of datatype uint256
    using PriceConverter for uint256;

    // set a minimum value to fund in USD
    uint256 public minFundUSD = 50 * 1e18;
    address[] public funders; // Array of funders
    mapping(address => uint256) public AddressToMoneyFunded; // Tracking the amount funded by each funder

    function fund() public payable {
        // condition to send atleast one ETH
        require(
            msg.value.getConversionRate() >= minFundUSD,
            "Funds not sufficient, minimum deposit must be 1 ETH"
        ); // here msg.value is this amount in Wei units and 1e18 means 1 * 10 ** 18 wei i.e. 1 eth

        // Reverting - Undoing any actions before, and sending the remaining gas back

        // Adding the funder to the Array and update map too
        funders.push(msg.sender);
        AddressToMoneyFunded[msg.sender] = msg.value;
    }

    function withdraw() public {
        // updating the values on map AddressToMoneyFunded before actually withdrawing the fund
        for (
            uint256 fundersIndex = 0;
            fundersIndex < funders.length;
            fundersIndex++
        ) {
            address funder = funders[fundersIndex];
            AddressToMoneyFunded[funder] = 0;
        }

        // resetting the whole funders array
        funders = new address[](0); // new array assigned of 0 elements in it

        // actually withdrawing the fund
        /*
            things to note 
            * Three methods to withdraw
                * transfer
                * send
                * call
            * msg.sender is just an address and not payable
            * payable(msg.sender) is a payable address
            * if the transaction failed in transfer method it throws an error which will revert the whole contract
            * if the transaction failed in send method it throws an bool and the contract proceeds 
        */

        // // transfer method
        // payable(msg.sender).transfer(address(this).balance); // this means this contract
        
        // // send method
        bool sendSuccess = payable(msg.sender).send(address(this).balance);
        require(sendSuccess, "send failed");

        // // call: its a low level function which very powerfull
        (bool callSuccess, ) = payable(msg.sender).call{value: address(this).balance}("");
        require(callSuccess, "call failed");
    }
}
