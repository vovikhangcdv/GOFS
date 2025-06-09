# Contract Call Example

This is a simple example of how to call contracts using the ethers.js library.

## Usage

```bash
npm install
```

```bash
node contract-call-example.js <rpc_url>
```

Expected output looks like this:

```
$ node contract-call-example.js
=== 1. Homepage - System Dashboard ===
Total eVND Supply: 1000000000000000000000000n
Verified Entities: <NEED TO IMPLEMENT>
Daily Transactions: <NEED TO IMPLEMENT>
Compliance Rules: 5
Exchange Rate: 24000n

=== 2. e-VND Token Details ===
Token Name: Vietnamese Digital Currency
Token Symbol: eVND
Token Address: 0xA51c1fc2f0D1a1b8494Ed1FE312d7C3a78Ed91C0
Token Decimals: 18n
Token Total Supply: 1000000000000000000000000n

=== 3. Entity Type Metadata ===
Entity Type Metadata:
  BANK (ID: 1): BANK: Financial institutions authorized to operate in Vietnam
  CORPORATE (ID: 2): CORPORATE: Business entities and organizations registered in Vietnam
  INDIVIDUAL (ID: 3): INDIVIDUAL: Vietnamese citizens and residents with verified identity
  EXCHANGE_PORTAL (ID: 10): EXCHANGE_PORTAL: Authorized currency exchange and trading platforms

=== 4. Transaction Type Metadata ===
Transaction Type Metadata:
  LOAN (ID: 1): LOAN: Bank loan transactions between financial institutions and borrowers
  PAYMENT (ID: 2): PAYMENT: Standard payment transactions between verified entities
  INVESTMENT (ID: 3): INVESTMENT: Investment transactions including securities and capital flows
  EXCHANGE (ID: 4): EXCHANGE: Currency exchange transactions through authorized portals

=== 5. Transaction Type Policies ===
Transaction Type Policies:

  LOAN (ID: 1):
    Usable: true
    Allowed From Entity Types: BANK
    Allowed To Entity Types: CORPORATE, INDIVIDUAL

  PAYMENT (ID: 2):
    Usable: true
    Allowed From Entity Types: BANK, CORPORATE, INDIVIDUAL
    Allowed To Entity Types: BANK, CORPORATE, INDIVIDUAL

  INVESTMENT (ID: 3):
    Usable: true
    Allowed From Entity Types: BANK, CORPORATE
    Allowed To Entity Types: CORPORATE, INDIVIDUAL

  EXCHANGE (ID: 4):
    Usable: true
    Allowed From Entity Types: EXCHANGE_PORTAL
    Allowed To Entity Types: EXCHANGE_PORTAL

=== 6. Entity Type Transfer Policies ===
Entity Type Transfer Policies:

  BANK can transfer to:
    ❌ BANK
    ✅ CORPORATE
    ✅ INDIVIDUAL
    ✅ EXCHANGE_PORTAL

  CORPORATE can transfer to:
    ✅ BANK
    ✅ CORPORATE
    ✅ INDIVIDUAL
    ✅ EXCHANGE_PORTAL

  INDIVIDUAL can transfer to:
    ✅ BANK
    ✅ CORPORATE
    ✅ INDIVIDUAL
    ✅ EXCHANGE_PORTAL

  EXCHANGE_PORTAL can transfer to:
    ✅ BANK
    ✅ CORPORATE
    ✅ INDIVIDUAL
    ✅ EXCHANGE_PORTAL

=== 7. Individual Entity Details ===
Entity Details:
  Address: 0x09635F643e140090A9A8Dcd712eD6285858ceBef
  Name: ExchangePortal
  Entity Type: EXCHANGE_PORTAL (ID: 10)
  Verified: Yes
  Balance: 1000000.0 e-VND
  Verifier: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8

=== 8. Individual Metadata Examples ===
Bank Entity Type: BANK: Financial institutions authorized to operate in Vietnam
Loan Transaction Type: LOAN: Bank loan transactions between financial institutions and borrowers
```
