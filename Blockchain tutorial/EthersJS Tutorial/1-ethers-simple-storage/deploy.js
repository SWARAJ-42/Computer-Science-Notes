const ethers = require("ethers");
const fs = require("fs-extra");

async function main() {
    // compile them in our code
    // compile them separately
    // http://127.0.0.1:7545
    // 0x76c2e4fcb6804878d8ae59528a61235707c7ac083045012800f65605c59efcdf

    // Setting up the network(ganache) and user
    const provider = new ethers.providers.JsonRpcProvider("http://127.0.0.1:7545");
    const wallet = new ethers.Wallet("0x76c2e4fcb6804878d8ae59528a61235707c7ac083045012800f65605c59efcdf", provider);

    // Defining the abi and bin files created from compiling
    const abi = fs.readFileSync("./SimpleStorage_sol_SimpleStorage.abi", "utf-8");
    const binary = fs.readFileSync("./SimpleStorage_sol_SimpleStorage.bin", "utf-8");

    // Assembling all the requires files to deploy the contract
    const contractFactory = new ethers.ContractFactory(abi, binary, wallet);
    console.log("Deploying please wait...");

    // arguments of deploy are called overrides
    const contract = await contractFactory.deploy(/*{gasPrice: 1000000000}*/); // STOP here!! and wait for the deployment
    // console.log(contract);
    
    console.log("Here is the deployment transaction (transaction response): ");
    console.log(contract.deployTransaction);

    const transactionReceipt = await contract.deployTransaction.wait(1);
    console.log("Here is the transaction receipt: ");
    console.log(transactionReceipt);

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

main().then(() => process.exit(0)).catch((error) => {
    console.error(error);
    process.exit(1);
});