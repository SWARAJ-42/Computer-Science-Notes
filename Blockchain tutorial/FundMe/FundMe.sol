// SPDX-License-Identifier: MIT
pragma solidity ^0.8.8;

import "FundMe/PriceConverter.sol";

// defining an exclusive error rather than using require for gas efficiency: check line 41
error NotOwner();

contract FundMe {
    // defining owner of the contract
    address public immutable i_owner; // immutable means the variable can only be changed inside the constructor, this tweak is again for gas efficiency

    constructor() {
        i_owner = msg.sender; // the sender who deployed the contract
    }

    // mentioning that PriceConvertor library can be used on any variable of datatype uint256
    using PriceConverter for uint256;

    // set a minimum value to fund in USD
    uint256 public constant MIN_USD = 50 * 1e18; // constant keyword for gas efficiency
    address[] public funders; // Array of funders
    mapping(address => uint256) public AddressToMoneyFunded; // Tracking the amount funded by each funder

    function fund() public payable {
        // condition to send atleast one ETH
        require(
            msg.value.getConversionRate() >= MIN_USD,
            "Funds not sufficient, minimum deposit must be 1 ETH"
        ); // here msg.value is this amount in Wei units and 1e18 means 1 * 10 ** 18 wei i.e. 1 eth

        // Reverting - Undoing any actions before, and sending the remaining gas back

        // Adding the funder to the Array and update map too
        funders.push(msg.sender);
        AddressToMoneyFunded[msg.sender] = msg.value;
    }

    // creating a modifier to check if the withdrawer is the owner
    modifier OnlyOwner() {
        // require(msg.sender == i_owner, "Withdraw failed: Only the owner can withdraw the fund.");
        if (msg.sender != i_owner) {
            revert NotOwner();
        }
        _; // This all the code of the function on which the modifier is used must be run after the code in the modifier
    }

    function withdraw() public OnlyOwner {
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
        // bool sendSuccess = payable(msg.sender).send(address(this).balance);
        // require(sendSuccess, "send failed");

        // // call: its a low level function which very powerfull
        (bool callSuccess, ) = payable(msg.sender).call{
            value: address(this).balance
        }("");
        require(callSuccess, "call failed");
    }

    // for handling transaction from outside the contract scope functions
    // this is for the senders to donate the fund via "fund" function only and we can track them via fund function
    receive() external payable {
        fund();
    }

    fallback() external payable {
        fund();
    }
}
