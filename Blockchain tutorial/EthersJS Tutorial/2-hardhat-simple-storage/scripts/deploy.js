// imports
// stress on the point that we are import hardhat
const {ethers, run, network} = require("hardhat")

// async main
async function main() {
    const [owner] = await ethers.getSigners()

    const hardhatToken = await ethers.deployContract("SimpleStorage")
    console.log("Deploying...")
    await hardhatToken.waitForDeployment()
    console.log(`Contract creator address: ${owner.address}`)
    const contract_address = await hardhatToken.getAddress()
    console.log(`Deployed contract to: ${contract_address}`)

    // what happens when we deploy the contract in hardhat network?
    // console.log(network.config)

    if (network.config.chainId === 11155111 && process.env.ETHERSCAN_API_KEY) {
      console.log("Waiting for block confirmation...")
      await hardhatToken.deploymentTransaction().wait(6)
      await verify(contract_address, [])
    }

    // interacting with the contract

    // getting the current value
    const currentValue =  await hardhatToken.retrieve()
    console.log(`Current value: ${currentValue}`)
    
    // updating the current value
    const transactionResponse = await hardhatToken.store(7)
    await transactionResponse.wait(1)
    
    // getting the new value 
    const updated_value = await hardhatToken.retrieve()
    console.log(`Updated value: ${updated_value}`)

}

// verify function 
async function verify(contractAddress, args) {
  console.log("Verifying contract...")
  try{
    await run("verify:verify", {
      address: contractAddress,
      constructorArguments: args,
    })
  }catch(e){
    if (e.message.toLowerCase().includes("already verified")){
      console.log("Already verified !")
    }
    else {
      console.log(e)
    }
  }
}

// main
main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error)
        process.exit(1)
    })
