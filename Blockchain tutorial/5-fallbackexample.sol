// SPDX-License-Identifier: MIT

pragma solidity ^0.8.7;

contract FallbackExample {
    uint256 public result; 

    // receive
    // receive function only gets triggered if no data/instruction is sent along with the transaction
    // receive function for transaction event is analogous to constructor for contract deployment
    receive() external payable {
        result = 1;
    }
    // now try sending ETH to this contract externally without instructions, maybe via calldata in localVM or Metamask in test network
    // The receive function will trigger automatically and the result value will be modified to 1


    // fallback
    // fallback function get triggered if the instruction/data is not interpreted by the contract
    fallback() external payable {
        result = 2;
    }
    // now try sending ETH to this contract externally with some gibberous hexadecimal code, maybe via calldata in local VM or Metamask in test network
    // The fallbacl function will trigger automatically and the result value will be modified to 2

    // Ether is sent to contract
    //     is msg.data empty?
    //           / \
    //         yes  no
    //         /     \
    //     recieve()? fallback()
    //       /    \
    //     yes     no
    //     /         \
    // receive()   fallback()
}