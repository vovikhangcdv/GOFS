# Upgradeable Contracts Guide

This guide explains how to deploy and manage the contracts using OpenZeppelin's TransparentProxy pattern for upgradeability.

## Overview

All contracts have been updated to support the TransparentProxy pattern, which allows you to:

-   Upgrade contract logic while maintaining the same address
-   Preserve state across upgrades
-   Maintain contract interactions without changing addresses

## Prerequisites

### Required Dependencies

You need to install the OpenZeppelin upgradeable contracts:

```bash
# For Foundry
forge install OpenZeppelin/openzeppelin-contracts-upgradeable
forge install OpenZeppelin/openzeppelin-foundry-upgrades

# For Hardhat/npm
npm install @openzeppelin/contracts-upgradeable
npm install @openzeppelin/hardhat-upgrades
```

### Updated Contracts

The following contracts have been converted to upgradeable versions:

1. **Core Contracts:**

    - `ComplianceRegistry.sol`
    - `CompliantToken.sol`
    - `EntityRegistry.sol`
    - `ExchangePortal.sol`
    - `TypedToken.sol` (abstract base)

2. **Compliance Modules:**
    - `AddressRestrictionCompliance.sol`
    - `EntityTypeCompliance.sol`
    - `SupplyCompliance.sol`
    - `TransactionTypeCompliance.sol`
    - `VerificationCompliance.sol`

## Key Changes Made

### 1. Constructor to Initializer Pattern

-   Replaced constructors with `initialize()` functions
-   Added `_disableInitializers()` in constructors to prevent direct implementation calls
-   Used `initializer` modifier to ensure one-time initialization

### 2. Inheritance Updates

-   Changed from regular OpenZeppelin contracts to upgradeable versions:
    -   `AccessControl` → `AccessControlUpgradeable`
    -   `ERC20` → `ERC20Upgradeable`
    -   `EIP712` → `EIP712Upgradeable`

### 3. Immutable Variables Removed

-   Converted `immutable` variables to regular state variables
-   This was necessary because immutable variables are stored in bytecode, not state

### 4. Storage Gaps Added

-   Added `__gap` arrays to all contracts for future upgrade compatibility
-   Gaps reserve storage slots for future versions

## Deployment Guide

### Using the Provided Script

A deployment script is provided at `script/DeployUpgradeableContracts.s.sol`:

```bash
forge script script/DeployUpgradeableContracts.s.sol:DeployUpgradeableContracts \
  --fork-url <YOUR_RPC_URL> \
  --broadcast \
  --verify
```

### Manual Deployment Example

```solidity
// Deploy EntityRegistry
address entityRegistryProxy = Upgrades.deployTransparentProxy(
    "EntityRegistry.sol",
    admin, // ProxyAdmin owner
    abi.encodeCall(EntityRegistry.initialize, ())
);

// Deploy ComplianceRegistry
address complianceRegistryProxy = Upgrades.deployTransparentProxy(
    "ComplianceRegistry.sol",
    admin,
    abi.encodeCall(ComplianceRegistry.initialize, ())
);

// Deploy CompliantToken
address tokenProxy = Upgrades.deployTransparentProxy(
    "CompliantToken.sol",
    admin,
    abi.encodeCall(CompliantToken.initialize, (
        "Government Digital Currency",
        "GDC",
        complianceRegistryProxy
    ))
);
```

## Upgrading Contracts

### Prepare New Implementation

1. Create new contract version inheriting from the same base
2. Ensure no storage layout changes (unless carefully managed)
3. Add new storage variables at the end if needed
4. Test thoroughly with upgrade simulation

### Perform Upgrade

```solidity
// Using OpenZeppelin Upgrades plugin
Upgrades.upgradeProxy(proxyAddress, "NewContractVersion.sol");
```

### Storage Layout Considerations

⚠️ **Important**: When upgrading contracts, be careful about storage layout:

-   Don't change the order of existing state variables
-   Don't change the type of existing state variables
-   Don't remove state variables
-   Add new variables at the end
-   Use storage gaps for future expansion

## Contract Initialization

### Order of Deployment

Follow this order to avoid dependency issues:

1. **EntityRegistry** (no dependencies)
2. **ComplianceRegistry** (no dependencies)
3. **Compliance Modules** (may depend on EntityRegistry)
4. **CompliantToken** (depends on ComplianceRegistry)
5. **SupplyCompliance** (depends on CompliantToken)
6. **ExchangePortal** (depends on CompliantToken)

### Configuration After Deployment

After deployment, you'll need to configure the system:

```solidity
// Add compliance modules to ComplianceRegistry
ComplianceRegistry registry = ComplianceRegistry(complianceRegistryProxy);
registry.addModule(verificationComplianceProxy);
registry.addModule(entityTypeComplianceProxy);
// ... add other modules

// Configure SupplyCompliance
SupplyCompliance supply = SupplyCompliance(supplyComplianceProxy);
supply.setMaxSupply(1e9 * 1e18); // 1 billion tokens

// Add verifiers to EntityRegistry
EntityRegistry entityReg = EntityRegistry(entityRegistryProxy);
entityReg.addVerifier(verifierAddress, [EntityType.wrap(1), EntityType.wrap(2)]);
```

## Security Considerations

### Admin Keys

-   The ProxyAdmin owner has upgrade privileges
-   Consider using a multisig wallet or governance contract for admin
-   Implement timelock for upgrades in production

### Testing Upgrades

Always test upgrades on testnet:

```bash
# Deploy original version
forge script script/DeployUpgradeableContracts.s.sol --fork-url $TESTNET_RPC

# Test upgrade
forge script script/UpgradeContracts.s.sol --fork-url $TESTNET_RPC
```

### Verification

After deployment, verify that:

-   All contracts are properly initialized
-   Proxy admin is set correctly
-   All compliance modules are registered
-   Dependencies between contracts work as expected

## Common Issues and Solutions

### Issue: "Contract instance has been previously initialized"

**Solution**: Ensure you're not calling `initialize()` twice. This can happen if you call it directly after deployment.

### Issue: Storage collision during upgrade

**Solution**: Use `forge inspect ContractName storage-layout` to compare storage layouts before upgrading.

### Issue: Interface detection failing

**Solution**: Ensure `supportsInterface` is properly implemented in upgradeable contracts.

## Monitoring and Maintenance

### Events to Monitor

-   `Upgraded` events from proxy contracts
-   `AdminChanged` events from ProxyAdmin
-   Initialization events from each contract

### Regular Checks

-   Verify proxy admin ownership
-   Check implementation addresses
-   Monitor for any failed transactions due to upgrade issues

## Examples

See the deployment script for complete examples of:

-   Deploying all contracts with proper initialization
-   Setting up compliance modules
-   Configuring token parameters
-   Establishing entity verification system

This upgrade to the TransparentProxy pattern provides a robust foundation for maintaining and evolving your smart contract system while preserving addresses and state across upgrades.
