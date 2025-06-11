// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script, console2} from "forge-std/Script.sol";
import {BaseScript} from "./BaseScript.s.sol";
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
import {EntityLibrary} from "../src/libraries/EntityLibrary.sol";
import {MERC20} from "../test/mocks/MockERC20.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {SetEntityTypePolicies} from "./SetEntityTypePolicies.s.sol";
import {SetTransactionTypePolicies} from "./SetTransactionTypePolicies.s.sol";

contract SimpleDeployAllWithoutSetPolicies is
    BaseScript,
    SetEntityTypePolicies,
    SetTransactionTypePolicies
{
    // Single ProxyAdmin for all proxies

    // Constants for initial setup
    uint256 constant INITIAL_EXCHANGE_RATE = 24_000 * 1e18; // 24:1 initial rate
    uint256 constant INITIAL_EXCHANGE_FEE = 10; // 0.1% fee (in basis points)

    function run()
        public
        override(SetEntityTypePolicies, SetTransactionTypePolicies)
    {
        // Initialize base script
        __BaseScript_init();

        // Start recording transactions for deployment
        vm.startBroadcast(vm.deriveKey(mnemonic, 0));

        // Deploy a single ProxyAdmin
        proxyAdmin = new ProxyAdmin(deployer);

        // Deploy all contracts
        _deployEntityRegistry();
        _deployComplianceRegistry();
        _deployCompliantToken();
        _deployComplianceModules();
        _deployExchangePortal();

        // // Deploy and setup policies using the dedicated scripts
        // setupEntityTypePolicies();
        // setupTransactionTypePolicies();

        vm.stopBroadcast();

        // Log deployment addresses
        logDeploymentAddresses();
        // logEntityTypePolicies();
        // logTransactionTypePolicies();
    }

    function _deployEntityRegistry() internal {
        // 1. Deploy EntityRegistry with transparent proxy
        EntityRegistry entityRegistryImplementation = new EntityRegistry();
        entityRegistry = EntityRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(entityRegistryImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(EntityRegistry.initialize.selector)
                )
            )
        );

        // Setup EntityRegistry
        EntityType[] memory entityTypes = new EntityType[](14);
        entityTypes[0] = EntityLibrary.INDIVIDUAL;
        entityTypes[1] = EntityLibrary.PRIVATE_ENTERPRISE;
        entityTypes[2] = EntityLibrary.LLC_ONE_MEMBER;
        entityTypes[3] = EntityLibrary.LLC_MULTI_MEMBER;
        entityTypes[4] = EntityLibrary.JOINT_STOCK_COMPANY;
        entityTypes[5] = EntityLibrary.PARTNERSHIP;
        entityTypes[6] = EntityLibrary.COOPERATIVE;
        entityTypes[7] = EntityLibrary.UNION_OF_COOPERATIVES;
        entityTypes[8] = EntityLibrary.COOPERATIVE_GROUP;
        entityTypes[9] = EntityLibrary.HOUSEHOLD_BUSINESS;
        entityTypes[10] = EntityLibrary.FOREIGN_INDIVIDUAL_INVESTOR;
        entityTypes[11] = EntityLibrary.FOREIGN_ORGANIZATION_INVESTOR;
        entityTypes[12] = EntityLibrary.FOREIGN_INVESTED_ECONOMIC_ORGANIZATION;
        entityTypes[13] = EntityLibrary.EXCHANGE_PORTAL;

        verifier_1 = vm.addr(vm.deriveKey(mnemonic, 1)); // First verifier address
        verifier_2 = vm.addr(vm.deriveKey(mnemonic, 2)); // Second verifier address
        entityRegistry.addVerifier(verifier_1, entityTypes);
        entityRegistry.addVerifier(verifier_2, entityTypes);
    }

    function _deployComplianceRegistry() internal {
        // 2. Deploy ComplianceRegistry with transparent proxy
        ComplianceRegistry complianceRegistryImplementation = new ComplianceRegistry();
        complianceRegistry = ComplianceRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(complianceRegistryImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        ComplianceRegistry.initialize.selector
                    )
                )
            )
        );
    }

    function _deployCompliantToken() internal {
        // 3. Deploy CompliantToken (e-VND) with transparent proxy
        CompliantToken evndTokenImplementation = new CompliantToken();
        evndToken = CompliantToken(
            address(
                new TransparentUpgradeableProxy(
                    address(evndTokenImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        CompliantToken.initialize.selector,
                        "Vietnamese Digital Currency",
                        "eVND",
                        address(complianceRegistry)
                    )
                )
            )
        );
    }

    function _deployComplianceModules() internal {
        // 4. Deploy compliance modules with transparent proxies
        // AddressRestrictionCompliance
        AddressRestrictionCompliance addressRestrictionImplementation = new AddressRestrictionCompliance();
        addressRestrictionCompliance = AddressRestrictionCompliance(
            address(
                new TransparentUpgradeableProxy(
                    address(addressRestrictionImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        AddressRestrictionCompliance.initialize.selector
                    )
                )
            )
        );

        // VerificationCompliance
        VerificationCompliance verificationImplementation = new VerificationCompliance();
        verificationCompliance = VerificationCompliance(
            address(
                new TransparentUpgradeableProxy(
                    address(verificationImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        VerificationCompliance.initialize.selector,
                        address(entityRegistry)
                    )
                )
            )
        );

        // SupplyCompliance
        SupplyCompliance supplyImplementation = new SupplyCompliance();
        supplyCompliance = SupplyCompliance(
            address(
                new TransparentUpgradeableProxy(
                    address(supplyImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        SupplyCompliance.initialize.selector,
                        address(evndToken)
                    )
                )
            )
        );

        // TransactionTypeCompliance
        TransactionTypeCompliance transactionTypeImplementation = new TransactionTypeCompliance();
        transactionTypeCompliance = TransactionTypeCompliance(
            address(
                new TransparentUpgradeableProxy(
                    address(transactionTypeImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        TransactionTypeCompliance.initialize.selector,
                        address(entityRegistry)
                    )
                )
            )
        );

        // EntityTypeCompliance
        EntityTypeCompliance entityTypeImplementation = new EntityTypeCompliance();
        entityTypeCompliance = EntityTypeCompliance(
            address(
                new TransparentUpgradeableProxy(
                    address(entityTypeImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        EntityTypeCompliance.initialize.selector,
                        address(entityRegistry)
                    )
                )
            )
        );

        // Add compliance modules to registry
        complianceRegistry.addModule(address(addressRestrictionCompliance));
        complianceRegistry.addModule(address(verificationCompliance));
        complianceRegistry.addModule(address(supplyCompliance));
        complianceRegistry.addModule(address(transactionTypeCompliance));
        complianceRegistry.addModule(address(entityTypeCompliance));

        // Setup blacklisters for AddressRestrictionCompliance
        blacklister_1 = vm.addr(vm.deriveKey(mnemonic, 3)); // First blacklister address
        blacklister_2 = vm.addr(vm.deriveKey(mnemonic, 4)); // Second blacklister address
        addressRestrictionCompliance.grantRole(
            addressRestrictionCompliance.BLACKLIST_ADMIN_ROLE(),
            blacklister_1
        );
        addressRestrictionCompliance.grantRole(
            addressRestrictionCompliance.BLACKLIST_ADMIN_ROLE(),
            blacklister_2
        );

        // Set max supply for SupplyCompliance
        supplyCompliance.setMaxSupply(
            1_000_000_000 * 10 ** evndToken.decimals() // Set supply limit to 1 billion eVND
        );
    }

    function _deployExchangePortal() internal {
        // 5. Deploy ExchangePortal with transparent proxy
        MERC20 mUSD = new MERC20("Mock US Dollar", "mUSD");
        ExchangePortal exchangePortalImplementation = new ExchangePortal();
        exchangePortal = ExchangePortal(
            address(
                new TransparentUpgradeableProxy(
                    address(exchangePortalImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        ExchangePortal.initialize.selector,
                        address(mUSD),
                        address(evndToken),
                        INITIAL_EXCHANGE_RATE,
                        deployer, // Treasury is set to deployer initially
                        INITIAL_EXCHANGE_FEE
                    )
                )
            )
        );

        // Register ExchangePortal as an entity
        _registerExchangePortalEntity();

        // Fund the exchange portal
        mUSD.mint(address(exchangePortal), 1000000 * 10 ** 18); // Mint 1 million mUSD to portal
        evndToken.mint(
            address(exchangePortal),
            1000000 * 10 ** evndToken.decimals()
        ); // Mint 1 million eVND to portal
    }

    function _registerExchangePortalEntity() internal {
        // Create entity data
        bytes memory entityData = abi.encode(
            "ExchangePortal", // name
            keccak256("Official Exchange Portal") // info
        );

        Entity memory portalEntity = Entity({
            entityType: EntityLibrary.EXCHANGE_PORTAL,
            entityData: entityData,
            entityAddress: address(exchangePortal),
            verifier: verifier_1
        });

        // Calculate digest using the same method as in tests
        bytes32 structHash = keccak256(
            abi.encode(
                entityRegistry.ENTITY_TYPE_HASH(),
                portalEntity.entityAddress,
                portalEntity.entityType,
                portalEntity.entityData,
                portalEntity.verifier
            )
        );
        bytes32 digest = keccak256(
            abi.encodePacked(
                "\x19\x01",
                entityRegistry.domainSeparator(),
                structHash
            )
        );

        // Sign using verifier's private key
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(
            vm.deriveKey(mnemonic, 1),
            digest
        );
        bytes memory signature = abi.encodePacked(r, s, v);

        // Register from the ExchangePortal's address
        entityRegistry.forwardRegister(portalEntity, signature);
    }

    function logDeploymentAddresses() public view override {
        // Log deployment addresses using console2 from forge-std
        console2.log("========== Addresses ==========");
        console2.log(
            "Deployer Address (ADMIN FOR ALL):",
            deployer,
            " (DeriveKey 0)"
        );
        console2.log("ProxyAdmin:", address(proxyAdmin));
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
