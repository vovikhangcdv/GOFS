# GOFS/e-VND Development Guide

## Setup Development Environment
1. Use anvil for local development:
   ```bash
   anvil
   ```
2. Deploy all contracts using the provided script:
   ```bash
   cp contracts/env.example .env # Please get the keys from env
   forge script script/SimpleDeployAll.s.sol:SimpleDeployAll  --rpc-url 127.0.0.1:8545 --broadcast
    ```
We will get the following output:
   ```
    == Logs ==
    Deployment completed:
    Deployer Address (ADMIN FOR ALL): 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266  (DeriveKey 0)
    EntityRegistry: 0x5FbDB2315678afecb367f032d93F642f64180aa3
    ComplianceRegistry: 0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9
    eVND Token: 0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9
    mUSD Token: 0x9A676e781A523b5d0C0e43731313A708CB607508
    Exchange Portal: 0x0B306BF915C4d645ff596e518fAf3F9669b97016
    AddressRestrictionCompliance: 0x5FC8d32690cc91D4c39d9d3abcBD16989F875707
    VerificationCompliance: 0x0165878A594ca255338adfa4d48449f69242Eb8F
    SupplyCompliance: 0xa513E6E4b8f2a923D98304ec87F64353C4D5C853
    Verifier 1 Address: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 (DeriveKey 1)
    Verifier 2 Address: 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC (DeriveKey 2)
    Blacklister 1 Address: 0x90F79bf6EB2c4f870365E785982E1f101E93b906 (DeriveKey 3)
    Blacklister 2 Address: 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65 (DeriveKey 4)
   ```
Please use the above addresses for further development.