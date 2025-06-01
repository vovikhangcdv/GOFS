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
    EntityRegistry: 0xD8a5a9b31c3C0232E196d518E89Fd8bF83AcAd43
    ComplianceRegistry: 0x36b58F5C1969B7b6591D752ea6F5486D069010AB
    eVND Token: 0x8198f5d8F8CfFE8f9C413d98a0A55aEB8ab9FbB7
    Exchange Portal: 0xC9a43158891282A2B1475592D5719c001986Aaec
    AddressRestrictionCompliance: 0x0355B7B8cb128fA5692729Ab3AAa199C1753f726
    VerificationCompliance: 0x202CCe504e04bEd6fC0521238dDf04Bc9E8E15aB
    SupplyCompliance: 0xf4B146FbA71F41E0592668ffbF264F1D186b2Ca8
    Verifier 1 Address: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 (DeriveKey 1)
    Verifier 2 Address: 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC (DeriveKey 2)
    Blacklister 1 Address: 0x90F79bf6EB2c4f870365E785982E1f101E93b906 (DeriveKey 3)
    Blacklister 2 Address: 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65 (DeriveKey 4)
   ```
Please use the above addresses for further development.