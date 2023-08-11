const { ethers } = require("hardhat")
const { expect, chai, assert } = require("chai")

describe('Contract-SimpleStorage', () => { 
    // predefining the contract variables for usage in it()
    let SimpleStorage

    beforeEach( async () => {
        const [owner] = await ethers.getSigners()
        // console.log("Deploying...")
        SimpleStorage = await ethers.deployContract("SimpleStorage")
        await SimpleStorage.waitForDeployment()
        // console.log(`Contract creator address: ${owner.address}`)
        // console.log(`Contract deployed at: ${await SimpleStorage.getAddress()}`)
    })

    /*
        if it.only("", function) is present then only this gets to be tested and others don't
    */

    // Test 1
    it("Should start with a favouriteNumber of value 0", async () => {
        const currentValue = await SimpleStorage.retrieve()
        const expectedValue = "0"
        // assert 
        // expect
        // expect(currentValue.toString()).to.equal(expectedValue) // does exactly the same thing as assert 
        assert.equal(currentValue.toString(), expectedValue)
    })

    // Test 2
    it("Should update when we call store", async () => {
        const expectedValue = "7"
        const transactionResponse = await SimpleStorage.store(expectedValue)
        await transactionResponse.wait(1)
        const recieved_value = await SimpleStorage.retrieve()
        assert.equal(expectedValue, recieved_value.toString())
    })
 })
