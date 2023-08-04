const ethers = require("ethers")
const fs = require("fs-extra")
require("dotenv").config()

async function main() {
    // compile them in our code
    // compile them separately

    // Setting up the network(ganache) and user
    const provider = new ethers.providers.JsonRpcProvider(process.env.RPC_URL)
    const wallet = new ethers.Wallet(process.env.PRIVATE_KEY, provider);

    // Defining the abi and bin files created from compiling
    // const encryptedJson = fs.readFileSync("./.encryptedKey.json", "utf-8")
    // let wallet = new ethers.Wallet.fromEncryptedJsonSync(
    //     encryptedJson,
    //     process.env.PRIVATE_KEY_PASSWORD
    // )

    const abi = fs.readFileSync(
        "./SimpleStorage_sol_SimpleStorage.abi",
        "utf-8"
    )
    const binary = fs.readFileSync(
        "./SimpleStorage_sol_SimpleStorage.bin",
        "utf-8"
    )

    // Assembling all the requires files to deploy the contract
    const contractFactory = new ethers.ContractFactory(abi, binary, wallet)
    console.log("Deploying please wait...")

    // arguments of deploy are called overrides
    const contract = await contractFactory.deploy(/*{gasPrice: 1000000000}*/) // STOP here!! and wait for the deployment
    await contract.deployTransaction.wait(1)
    console.log(`Contract creator: ${contract.address}`)
    // console.log(contract);

    // console.log("Here is the deployment transaction (transaction response): ");
    // console.log(contract.deployTransaction);

    // const transactionReceipt = await contract.deployTransaction.wait(1);
    // console.log("Here is the transaction receipt: ");
    // console.log(transactionReceipt);

    // Get number
    const currentFavouriteNumber = await contract.retrieve() // retrieve function is defined in SimpleStorage.sol
    console.log(
        `Current Favourite Number: ${currentFavouriteNumber.toString()}`
    )
    const transactionResponse = await contract.store("7") // Store function is defined in SimpleStorage.sol
    const transactionReceipt = await transactionResponse.wait(1)
    const updatedFavouriteNumber = await contract.retrieve()
    console.log(`Updated favourite number is: ${updatedFavouriteNumber}`)

    // manually creating a transaction data
    // console.log("Let's deploy with only transaction data!");
    // const nonce = await wallet.getTransactionCount();
    // const tx = {
    //     nonce: nonce,
    //     gasPrice: 20000000,
    //     gasLimit: 100000,
    //     to: null,
    //     value:0,
    //     data: "0x00",
    //     chainId: 1337
    // };

    // const signedTxResponse = await wallet.signTransaction(tx);
    // console.log(signedTxResponse);

    // const sendTxResponse = await wallet.sendTransaction(tx);
    // await sendTxResponse.wait(1);
    // console.log(sendTxResponse);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error)
        process.exit(1)
    })
