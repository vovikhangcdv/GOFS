# GOFS/e-VND Development Guide

## 📋 Table of Contents
1. [Task Overview](#task-overview)
2. [Project Structure](#project-structure)

## 📊 Task Overview

| Component | Objectives | Strategy | Owner |
|-----------|-----------|----------|-------|
| **Identity Registry & ERC-4337 Wallets** | • KYC/Onboarding system<br>• Entity type support<br>• KYC/AML checks | • Use ERC-1155 for entity type NFTs<br>• Implement ERC-4337 entry point contract<br>• Create separate KYC verification middleware<br>• Flexible wallet centralization/magenement/upgradable | @hao |
| **CBDC-Token & Policy Registry** | • e-VND token issuance<br>• Monetary policy management<br>• External exchange portal | • Extend ERC-20 implementation<br>• Implement policy hooks for transaction validation<br>• Create atomic swap contract for exchanges | @khang |
| **Fault Detection System** | • Monitor blockchain activities<br>• Anti-Money Laundering checks<br>• Real-time reporting | • Create off-chain monitoring service<br>• Implement rule engine based on [Vietnam's AML law](https://thuvienphapluat.vn/van-ban/Tien-te-Ngan-hang/Luat-14-2022-QH15-Phong-chong-rua-tien-519327.aspx)<br> | @phu |
| **Front-end dApps** | • Demo strategy<br>• Wallet interface<br>• CBDC Explorer<br>• KYC/AML UI | | @thanh |

## 🗂️ Project Structure

```
gofs/
├── contracts
├── dapps
├── fds
├── DEVELOPMENT.MD
└── README.md
```