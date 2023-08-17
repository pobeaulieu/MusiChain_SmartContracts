
# MusiChain Deployment Tool

## Environment Setup: Ganache and Metamask Connection

### Metamask Account Creation
1. Create a Metamask account.
2. Obtain the seed phrase associated with the Metamask wallet.

### Ganache Workspace Setup
1. Create a new Ganache workspace.
2. Provide the Metamask wallet's seed phrase during Ganache setup.

### Adding Test Network to Metamask
1. Open Metamask.
2. Add a custom network with the following details:
    - Network Name: Test Network
    - New RPC URL: http://localhost:7545
    - Chain ID: 1337
    - Currency Symbol: ETH

Now, the test accounts from Ganache should be accessible through Metamask.

## Deployment Process

1. Write the Solidity contracts.
2. Use Solc to generate .bin and .abi files for each contract. Alternatively, Remix compiler can be used for this purpose.
3. Use Abigen to generate .go files using the following command. Repeat this step for all three contracts.
`abigen --bin=pkg/services/build/Base.bin --abi=pkg/services/build/Base.abi --pkg=base --out=pkg/services/abigen/base/baseContract.go`
4. In the `main.go` file, add the private key of the account where the contracts will be deployed.
5. Run the code in `main.go`. This code will deploy the contracts and generate the `.env` file containing contract addresses.
