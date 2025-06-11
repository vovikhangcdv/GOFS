// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script, console2} from "forge-std/Script.sol";
import {EntityRegistry} from "../src/EntityRegistry.sol";
import {ComplianceRegistry} from "../src/ComplianceRegistry.sol";
import {CompliantToken} from "../src/CompliantToken.sol";
import {ExchangePortal} from "../src/ExchangePortal.sol";
import {EntityType, Entity, TxType} from "../src/interfaces/ITypes.sol";
import {AddressRestrictionCompliance} from "../src/compliances/AddressRestrictionCompliance.sol";
import {VerificationCompliance} from "../src/compliances/VerificationCompliance.sol";
import {SupplyCompliance} from "../src/compliances/SupplyCompliance.sol";
import {TransactionTypeCompliance} from "../src/compliances/TransactionTypeCompliance.sol";
import {EntityTypeCompliance} from "../src/compliances/EntityTypeCompliance.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

contract BaseScript is Script {
    // Contract instances
    EntityRegistry public entityRegistry;
    ComplianceRegistry public complianceRegistry;
    CompliantToken public evndToken;
    ExchangePortal public exchangePortal;
    AddressRestrictionCompliance public addressRestrictionCompliance;
    VerificationCompliance public verificationCompliance;
    SupplyCompliance public supplyCompliance;
    TransactionTypeCompliance public transactionTypeCompliance;
    EntityTypeCompliance public entityTypeCompliance;
    ProxyAdmin public proxyAdmin;

    // Addresses
    address public deployer;
    address public verifier_1;
    address public verifier_2;
    address public blacklister_1;
    address public blacklister_2;
    string public mnemonic;

    // Default contract addresses - can be overridden
    address constant DEFAULT_ENTITY_REGISTRY =
        0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0;
    address constant DEFAULT_COMPLIANCE_REGISTRY =
        0x0165878A594ca255338adfa4d48449f69242Eb8F;
    address constant DEFAULT_EVND_TOKEN =
        0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6;
    address constant DEFAULT_EXCHANGE_PORTAL =
        0x67d269191c92Caf3cD7723F116c85e6E9bf55933;
    address constant DEFAULT_ADDRESS_RESTRICTION =
        0x610178dA211FEF7D417bC0e6FeD39F05609AD788;
    address constant DEFAULT_VERIFICATION =
        0xA51c1fc2f0D1a1b8494Ed1FE312d7C3a78Ed91C0;
    address constant DEFAULT_SUPPLY =
        0x9A676e781A523b5d0C0e43731313A708CB607508;
    address constant DEFAULT_TRANSACTION_TYPE =
        0x959922bE3CAee4b8Cd9a407cc3ac1C251C2007B1;
    address constant DEFAULT_ENTITY_TYPE =
        0x68B1D87F95878fE05B998F19b66F4baba5De1aed;

    function __BaseScript_init() internal {
        // Get deployer from mnemonic
        mnemonic = vm.envString("WALLET_MEMO");
        uint256 deployerPrivateKey = vm.deriveKey(mnemonic, 0);
        deployer = vm.addr(deployerPrivateKey);

        // Set up derived addresses
        verifier_1 = vm.addr(vm.deriveKey(mnemonic, 1));
        verifier_2 = vm.addr(vm.deriveKey(mnemonic, 2));
        blacklister_1 = vm.addr(vm.deriveKey(mnemonic, 3));
        blacklister_2 = vm.addr(vm.deriveKey(mnemonic, 4));
    }

    function __BaseScript_initializeContracts() internal {
        // Initialize contract instances with default addresses
        entityRegistry = EntityRegistry(DEFAULT_ENTITY_REGISTRY);
        complianceRegistry = ComplianceRegistry(DEFAULT_COMPLIANCE_REGISTRY);
        evndToken = CompliantToken(DEFAULT_EVND_TOKEN);
        exchangePortal = ExchangePortal(DEFAULT_EXCHANGE_PORTAL);
        addressRestrictionCompliance = AddressRestrictionCompliance(
            DEFAULT_ADDRESS_RESTRICTION
        );
        verificationCompliance = VerificationCompliance(DEFAULT_VERIFICATION);
        supplyCompliance = SupplyCompliance(DEFAULT_SUPPLY);
        transactionTypeCompliance = TransactionTypeCompliance(
            DEFAULT_TRANSACTION_TYPE
        );
        entityTypeCompliance = EntityTypeCompliance(DEFAULT_ENTITY_TYPE);
    }

    function logDeploymentAddresses() public view virtual {
        console2.log("Current contract addresses:");
        console2.log(
            "Deployer Address (ADMIN FOR ALL):",
            deployer,
            " (DeriveKey 0)"
        );
        console2.log("EntityRegistry:", address(entityRegistry));
        console2.log("ComplianceRegistry:", address(complianceRegistry));
        console2.log("eVND Token:", address(evndToken));
        console2.log("Exchange Portal:", address(exchangePortal));
        console2.log(
            "AddressRestrictionCompliance:",
            address(addressRestrictionCompliance)
        );
        console2.log(
            "VerificationCompliance:",
            address(verificationCompliance)
        );
        console2.log("SupplyCompliance:", address(supplyCompliance));
        console2.log(
            "TransactionTypeCompliance:",
            address(transactionTypeCompliance)
        );
        console2.log("EntityTypeCompliance:", address(entityTypeCompliance));
        console2.log("Verifier 1 Address:", verifier_1, "(DeriveKey 1)");
        console2.log("Verifier 2 Address:", verifier_2, "(DeriveKey 2)");
        console2.log("Blacklister 1 Address:", blacklister_1, "(DeriveKey 3)");
        console2.log("Blacklister 2 Address:", blacklister_2, "(DeriveKey 4)");
    }
}
