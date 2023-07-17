// SPDX-License-Identifier: MIT
// upper line is optional but some compiler may through an error ( purpose is to make sharing easy )

pragma solidity ^0.8.7; // we have to specify the version of the solidity compiler
// ^ indicates this version or latest than it can compile the file
// pragma solidity >=0.8.7 <0.9.0; (range of compiler we want to use)

// tells the compiler, the next piece of code gonna be a contract
contract SimpleStorage {
    // boolean, uint, int, address, string, bytes (Most basic types)
    bool hasFavouriteNumber = true; // declaring and initiating a variable of type bool
    uint FavouriteNo = 256; // only store positive numbers, same as uint256 bits wise 
    string FavouriteNumberText = "Five";
    int256 Number = 5;
    address myAddress = 0xE3D5EcbfFaff17f7b936A66CCF65121da9d78989 ; // our metamask address
    bytes32 Favouritebyte = "cat"; // strings are secretly bytes object but only for string  ( looks like 0x12345 ) (bytes32 is maximum)


    uint256 public FavouriteNumber; // it gets initialised to 0
    // adding public we are basically creating an getter function for FavouriteNumber 


    // setting the value of the variable, in solidity everything we have to do using functions
    function store(uint256 _favouriteNumber) public virtual  {
        FavouriteNumber = _favouriteNumber;
        uint256 local = 12;// can only be assesed inside these function 
        local = local+1;
    }

    // every contract has its own address, you can see it after deploying it in the network (0xd9145CCE52D386f254917e481eB44e9943F39138)-address of this contract
    // every changes you do or deploy the contract, you have to pay some gas, you can see the trasactions in the terminals

    // the term public you see is the visibility
    // there are four types of visibility - public,private,external,internal (know more in documentation)

    // view, pure - does not cost gases on calling it as it does not change the state of the chain 
    function retrieve() public view returns(uint256) { 
        return FavouriteNo;
    } // same as writing public beside the variable

    function add() public pure returns(uint256){
        return(1+1);
    } // some basic math or algorithm

    struct People {
        uint256 favouriteNumber;
        string name;
    } // creating our own type

    People public person = People({favouriteNumber : 2, name : "patrick"});// declaring and initializing a variable of our type

    People[] public people; //creating an array (we have not given a number inside the bracket so it can be of any size)  
    
    // there are six places in solidity where we can store our data, the three main are - calldata, memory, storage 
    // calldata and memory means the data only stores temporarily but storage is permanent and default
    // You cannot use storage in functions
    // calldata doesnot allow to change itself once defined
    // memory can only be used for struct, array or mappings (string is just an array)
    function addperson(string memory _name, uint256 _favouriteNumber) public {
        people.push(People(_favouriteNumber, _name));
    } // setting the values in people variable


    mapping(string => uint256) public nameToFavouriteNumber; // similar to dictionary

    function setmap(string memory _name, uint256 _favouriteNumber) public {
        nameToFavouriteNumber[_name] = _favouriteNumber;
    } // setting the values in nameToFavouriteNumber
}



