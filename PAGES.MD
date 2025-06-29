# 🔍 GOFS/e-VND Explorer - Draft Pages

Draft page layouts for customizing Ethernal to showcase the GOFS/e-VND CBDC system.

## PLEASE REFER TO THE [contract-call-example](./dapps/contract-call-example/README.md) for the example of how to call the contracts.

_Note: This is draft minimal information_

### 🏠 1. Homepage - System Dashboard

```
┌─────────────────────────────────────────────────────────────┐
│ 🏛️ VIETNAM e-VND CBDC SYSTEM STATUS                        │
├─────────────────────────────────────────────────────────────┤
│ 💵 e-VND Token        │ 0xDc64...C6C9         │
│ 🔄 Exchange Portal    │ 0x0B30...7016         │
│ 🛂 Entity Registry    │ 0x5FbD...0aa3         │
│ ⚖️ Compliance Registry│ 0xCf7E...0Fc9         │
│ 💰 mUSD               │ 0x9A67...7508          │
├─────────────────────────────────────────────────────────────┤
│ 📈 System Metrics                                           │
│ Total e-VND Supply: 1,000,000,000 VND                      │
│ Verified Entities: 1,247 verified                            │
│ Daily Transactions: 12,450                                  │
│ Compliance Rules: 5 active                                  │
│ Exchange Rate: 1 USD = 24,000 VND                          │
└─────────────────────────────────────────────────────────────┘
```

---

### 💵 2. e-VND Token Details

#### 👥 Verifier Section

```
┌─────────────────────────────────────────────────────────────┐
│ 🔎 AUTHORIZED VERIFIERS                                     │
├─────────────────────────────────────────────────────────────┤
│ Address                    │ Name              │ Status      │
│ 0x7099...79C8             │ Ministry of Finance│ ✅ Active   │
│ 0x3C44...93BC             │ State Bank Vietnam│ ✅ Active   │
└─────────────────────────────────────────────────────────────┘
```

#### 📋 Typed Transactions (Optional)

```
┌─────────────────────────────────────────────────────────────┐
│ 📋 TYPED TRANSACTIONS                                       │
├─────────────────────────────────────────────────────────────┤
│ Classified Transactions: 100,000                            │
│ Buy-Sell: 100,000                                          │
│ FDI: 2,000,000                                           │
│ Transfer: 100,000                                          │
│ Deposit: 100,000                                           │
│ Withdraw: 100,000                                          │
│ ................                                           │
│ Unclassified Transactions: 100,000                          │
│ Total Transactions: 200,000                                 │
└─────────────────────────────────────────────────────────────┘
```

#### 🔍 Enhanced Transaction Display

```
┌─────────────────────────────────────────────────────────────┐
│ 📋 e-VND TRANSACTION DETAILS                               │
├─────────────────────────────────────────────────────────────┤
│ Hash: 0xabc123...                                          │
│ From: Nguyen Van A (Individual) ✅ Verified by SBV        │
│ To:   ABC Corporation (Business) ✅ Verified by MOF        │
│ Amount: 50,000 e-VND                                       │
│ Purpose: Business Payment                                   │
│ Transaction type: BUY/SAFE                              │
└─────────────────────────────────────────────────────────────┘
```

---

### 🔄 3. Exchange Portal Details

#### 📊 Real-Time Exchange Dashboard

```
┌─────────────────────────────────────────────────────────────┐
│ 💱 EXCHANGE PORTAL STATUS                                   │
├─────────────────────────────────────────────────────────────┤
│ Current Rates (Government Controlled):                      │
│ • 1 USD = 24,000 e-VND  (Updated: 2 min ago)              │
├─────────────────────────────────────────────────────────────┤
│ 📈 Money Flow Analytics                                     │
│ Today:    +125M e-VND inflow  │ 2,340 transactions        │
│ This Month: +2.1B e-VND net   │ 45,600 transactions       │
│ This Year:  +12.5B e-VND net  │ 891,200 transactions      │
└─────────────────────────────────────────────────────────────┘
```

**_+ other charts, graphs, etc (optional)_**

---

### 🛡️ 4. Fault Detection System (FDS)

#### ⚖️ Active Compliance Rules

```
┌─────────────────────────────────────────────────────────────┐
│ 📋 ACTIVE COMPLIANCE RULES                                  │
├─────────────────────────────────────────────────────────────┤
│ Rule ID │ Name                    │ Status  │ Violations/24h │
│ R001    │ Daily Transfer Limit    │ ✅ Active│ 23             │
│ R002    │ AML Risk Scoring        │ ✅ Active│ 5              │
│ R003    │ Sanctions Screening     │ ✅ Active│ 0              │
│ R004    │ Entity Verification     │ ✅ Active│ 12             │
│ R005    │ Cross-Border Reporting  │ ✅ Active│ 8              │
└─────────────────────────────────────────────────────────────┘
```

#### 🚫 Blacklist/Whitelist Management

```
┌─────────────────────────────────────────────────────────────┐
│ 🚨 ENFORCEMENT ACTIONS                                      │
├─────────────────────────────────────────────────────────────┤
│ Type      │ Address    │ Reason              │ Action Date    │
│ Blacklist │ 0x1234...  │ Sanctions Violation │ Dec 15, 2024   │
│ Blacklist │ 0x5678...  │ AML High Risk       │ Dec 14, 2024   │
│ Whitelist │ 0x9abc...  │ Government Entity   │ Dec 10, 2024   │
│ Removed   │ 0xdef0...  │ Appeal Successful   │ Dec 12, 2024   │
└─────────────────────────────────────────────────────────────┘
```

---

### ⚖️ 5. Compliance Registry

#### 📜 Registered Compliance Modules

```
┌─────────────────────────────────────────────────────────────┐
│ 🏛️ COMPLIANCE MODULES REGISTRY                             │
├─────────────────────────────────────────────────────────────┤
│ Module Name              │ Address      │ Status           │
│ AddressRestriction       │ 0x5FC8...    │ ✅ Active        │
│ VerificationCompliance   │ 0x0165...    │ ✅ Active        │
│ SupplyCompliance         │ 0xa513...    │ ✅ Active        │
│ ................         │ 0x1234...    │ ✅ Active        │
└─────────────────────────────────────────────────────────────┘
```
