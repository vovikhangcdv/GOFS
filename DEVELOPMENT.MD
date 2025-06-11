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
cd contracts/ && forge script script/SimpleDeployAll.s.sol:SimpleDeployAll --rpc-url 127.0.0.1:8545 --broadcast

# starting chainspammer
# starting dapps/explorer
# starting fds
```

---

## 📜 Contracts

Contract Deployment Log: [./contracts/deployment.log](./contracts/deployment.log)

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
