# 🏗️ GOFS/e-VND Development Guide

This guide provides a high-level overview of the GOFS project structure and components.

## 📁 Project Structure

```
gofs/
├── contracts/          # Smart contract implementation
├── chainspammer/       # Transaction generation & testing tool
├── fds/               # Fault Detection System
├── dapps/             # Frontend applications
│   ├── explorer/      # Blockchain explorer
```

## 🔧 Basic Setup

```bash
# starting blockchain
anvil
# starting contracts
cd contracts/ && forge script script/SimpleDeployAll.s.sol:SimpleDeployAll --rpc-url 127.0.0.1:8545 --broadcast --slow

# starting chainspammer
# starting dapps/explorer
# starting fds
```

---

## 📜 Contracts

```
== Logs ==
  Deployment completed:
  Deployer Address (ADMIN FOR ALL): 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266  (DeriveKey 0)
  ProxyAdmin: 0x5FbDB2315678afecb367f032d93F642f64180aa3
  EntityRegistry: 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0
  ComplianceRegistry: 0x610178dA211FEF7D417bC0e6FeD39F05609AD788
  eVND Token: 0xA51c1fc2f0D1a1b8494Ed1FE312d7C3a78Ed91C0
  Exchange Portal: 0x09635F643e140090A9A8Dcd712eD6285858ceBef
  AddressRestrictionCompliance: 0x9A676e781A523b5d0C0e43731313A708CB607508
  VerificationCompliance: 0x959922bE3CAee4b8Cd9a407cc3ac1C251C2007B1
  SupplyCompliance: 0x68B1D87F95878fE05B998F19b66F4baba5De1aed
  TransactionTypeCompliance: 0x84eA74d481Ee0A5332c457a4d796187F6Ba67fEB
  Verifier 1 Address: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 (DeriveKey 1)
  Verifier 2 Address: 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC (DeriveKey 2)
  Blacklister 1 Address: 0x90F79bf6EB2c4f870365E785982E1f101E93b906 (DeriveKey 3)
  Blacklister 2 Address: 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65 (DeriveKey 4)
  Transaction Type Admin: 0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc (DeriveKey 5)
  EntityTypeCompliance: 0x8f86403A4DE0BB5791fa46B8e795C547942fE4Cf
```

---

## ⚡ ChainSpammer

See [chainspammer/README.md](chainspammer/README.md)

---

## 🛡️ FDS (Fault Detection System)

See [fds/README.md](fds/README.md)

---

## 🔍 Explorer

Bulild on top of [Ethernal](https://github.com/ethernal-labs/ethernal)

See [explorer/README.md](explorer/README.md)

---
