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
import {MERC20} from "../test/mocks/MockERC20.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

contract SimpleDeployAll is Script {
    EntityRegistry public entityRegistry;
    ComplianceRegistry public complianceRegistry;
    CompliantToken public evndToken;
    ExchangePortal public exchangePortal;
    AddressRestrictionCompliance public addressRestrictionCompliance;
    VerificationCompliance public verificationCompliance;
    SupplyCompliance public supplyCompliance;
    TransactionTypeCompliance public transactionTypeCompliance;
    EntityTypeCompliance public entityTypeCompliance;

    // Single ProxyAdmin for all proxies
    ProxyAdmin public proxyAdmin;

    // Constants for initial setup
    uint256 constant INITIAL_EXCHANGE_RATE = 24_000 * 1e18; // 24:1 initial rate
    uint256 constant INITIAL_EXCHANGE_FEE = 10; // 0.1% fee (in basis points)

    // Transaction Type constants
    TxType constant LOAN_TYPE = TxType.wrap(1);
    TxType constant PAYMENT_TYPE = TxType.wrap(2);
    TxType constant INVESTMENT_TYPE = TxType.wrap(3);
    TxType constant EXCHANGE_TYPE = TxType.wrap(4);

    // Entity Type constants
    EntityType constant BANK = EntityType.wrap(1);
    EntityType constant CORPORATE = EntityType.wrap(2);
    EntityType constant INDIVIDUAL = EntityType.wrap(3);
    EntityType constant EXCHANGE_PORTAL = EntityType.wrap(10);

    // Addresses
    address deployer;
    address verifier_1;
    address verifier_2;
    address blacklister_1;
    address blacklister_2;
    address txTypeAdmin;
    string mnemonic;

    function run() public {
        // Get deployer from mnemonic
        mnemonic = vm.envString("WALLET_MEMO");
        uint256 deployerPrivateKey = vm.deriveKey(mnemonic, 0);
        deployer = vm.addr(deployerPrivateKey);

        // Start recording transactions for deployment
        vm.startBroadcast(deployerPrivateKey);

        // Deploy a single ProxyAdmin
        proxyAdmin = new ProxyAdmin(deployer);

        // Deploy all contracts
        _deployEntityRegistry();
        _deployComplianceRegistry();
        _deployCompliantToken();
        _deployComplianceModules();
        _deployExchangePortal();
        _deployTransactionTypeCompliance();
        _deployEntityTypeCompliance();

        vm.stopBroadcast();

        // Log deployment addresses
        _logDeploymentAddresses();
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
        EntityType[] memory entityTypes = new EntityType[](3);
        entityTypes[0] = EntityType.wrap(1); // INDIVIDUAL
        entityTypes[1] = EntityType.wrap(2); // ORGANIZATION
        entityTypes[2] = EntityType.wrap(10); // EXCHANGE_PORTAL
        verifier_1 = vm.addr(vm.deriveKey(mnemonic, 1)); // First verifier address
        verifier_2 = vm.addr(vm.deriveKey(mnemonic, 2)); // Second verifier address
        entityRegistry.addVerifier(verifier_1, entityTypes);
        entityRegistry.addVerifier(verifier_2, entityTypes);

        // Setup entity type metadata
        _setupEntityTypeMetadata();
    }

    function _setupEntityTypeMetadata() internal {
        // Set metadata for entity types
        entityRegistry.setEntityTypeMetadata(
            BANK,
            "BANK: Financial institutions authorized to operate in Vietnam"
        );
        entityRegistry.setEntityTypeMetadata(
            CORPORATE,
            "CORPORATE: Business entities and organizations registered in Vietnam"
        );
        entityRegistry.setEntityTypeMetadata(
            INDIVIDUAL,
            "INDIVIDUAL: Vietnamese citizens and residents with verified identity"
        );
        entityRegistry.setEntityTypeMetadata(
            EXCHANGE_PORTAL,
            "EXCHANGE_PORTAL: Authorized currency exchange and trading platforms"
        );
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

        // Add compliance modules to registry
        complianceRegistry.addModule(address(addressRestrictionCompliance));
        complianceRegistry.addModule(address(verificationCompliance));
        complianceRegistry.addModule(address(supplyCompliance));

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
            entityType: EntityType.wrap(10), // EXCHANGE_PORTAL type
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

    function _deployTransactionTypeCompliance() internal {
        // 6. Deploy TransactionTypeCompliance with transparent proxy
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

        // Add transaction type compliance to registry
        complianceRegistry.addModule(address(transactionTypeCompliance));

        // Setup transaction type admin
        txTypeAdmin = vm.addr(vm.deriveKey(mnemonic, 5)); // Fifth derived address
        transactionTypeCompliance.grantRole(
            transactionTypeCompliance.TX_TYPE_ADMIN_ROLE(),
            txTypeAdmin
        );

        // Setup basic transaction type policies and metadata
        _setupTransactionTypePolicies();
    }

    function _setupTransactionTypePolicies() internal {
        // Set transaction type metadata
        transactionTypeCompliance.setTransactionTypeMetadata(
            LOAN_TYPE,
            "LOAN: Bank loan transactions between financial institutions and borrowers"
        );
        transactionTypeCompliance.setTransactionTypeMetadata(
            PAYMENT_TYPE,
            "PAYMENT: Standard payment transactions between verified entities"
        );
        transactionTypeCompliance.setTransactionTypeMetadata(
            INVESTMENT_TYPE,
            "INVESTMENT: Investment transactions including securities and capital flows"
        );
        transactionTypeCompliance.setTransactionTypeMetadata(
            EXCHANGE_TYPE,
            "EXCHANGE: Currency exchange transactions through authorized portals"
        );

        // Setup LOAN_TYPE policies
        // Banks can send loans to individuals and corporations
        EntityType[] memory loanFromTypes = new EntityType[](1);
        loanFromTypes[0] = EntityType.wrap(1); // BANK
        EntityType[] memory loanToTypes = new EntityType[](2);
        loanToTypes[0] = EntityType.wrap(2); // CORPORATE
        loanToTypes[1] = EntityType.wrap(3); // INDIVIDUAL
        transactionTypeCompliance.setTransactionTypePolicy(
            LOAN_TYPE,
            loanFromTypes,
            loanToTypes
        );

        // Setup PAYMENT_TYPE policies
        // All entity types can send and receive payments
        EntityType[] memory paymentFromTypes = new EntityType[](3);
        paymentFromTypes[0] = EntityType.wrap(1); // BANK
        paymentFromTypes[1] = EntityType.wrap(2); // CORPORATE
        paymentFromTypes[2] = EntityType.wrap(3); // INDIVIDUAL
        EntityType[] memory paymentToTypes = new EntityType[](3);
        paymentToTypes[0] = EntityType.wrap(1); // BANK
        paymentToTypes[1] = EntityType.wrap(2); // CORPORATE
        paymentToTypes[2] = EntityType.wrap(3); // INDIVIDUAL
        transactionTypeCompliance.setTransactionTypePolicy(
            PAYMENT_TYPE,
            paymentFromTypes,
            paymentToTypes
        );

        // Setup INVESTMENT_TYPE policies
        // Banks and corporations can invest, individuals can receive investments
        EntityType[] memory investmentFromTypes = new EntityType[](2);
        investmentFromTypes[0] = EntityType.wrap(1); // BANK
        investmentFromTypes[1] = EntityType.wrap(2); // CORPORATE
        EntityType[] memory investmentToTypes = new EntityType[](2);
        investmentToTypes[0] = EntityType.wrap(2); // CORPORATE
        investmentToTypes[1] = EntityType.wrap(3); // INDIVIDUAL
        transactionTypeCompliance.setTransactionTypePolicy(
            INVESTMENT_TYPE,
            investmentFromTypes,
            investmentToTypes
        );

        // Setup EXCHANGE_TYPE policies
        // Only exchange portals can handle exchange transactions
        EntityType[] memory exchangeFromTypes = new EntityType[](1);
        exchangeFromTypes[0] = EntityType.wrap(10); // EXCHANGE_PORTAL
        EntityType[] memory exchangeToTypes = new EntityType[](1);
        exchangeToTypes[0] = EntityType.wrap(10); // EXCHANGE_PORTAL
        transactionTypeCompliance.setTransactionTypePolicy(
            EXCHANGE_TYPE,
            exchangeFromTypes,
            exchangeToTypes
        );
    }

    function _deployEntityTypeCompliance() internal {
        // 7. Deploy EntityTypeCompliance with transparent proxy
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

        // Add entity type compliance to registry
        complianceRegistry.addModule(address(entityTypeCompliance));

        // Setup basic entity type transfer policies
        _setupEntityTypePolicies();
    }

    function _setupEntityTypePolicies() internal {
        // Setup transfer policies between entity types

        // Banks can transfer to all entity types
        entityTypeCompliance.setTransferPolicy(BANK, CORPORATE, true);
        entityTypeCompliance.setTransferPolicy(BANK, INDIVIDUAL, true);
        entityTypeCompliance.setTransferPolicy(BANK, EXCHANGE_PORTAL, true);

        // Corporations can transfer to banks, other corporations, and individuals
        entityTypeCompliance.setTransferPolicy(CORPORATE, BANK, true);
        entityTypeCompliance.setTransferPolicy(CORPORATE, CORPORATE, true);
        entityTypeCompliance.setTransferPolicy(CORPORATE, INDIVIDUAL, true);

        // Individuals can transfer to banks, corporations, and other individuals
        entityTypeCompliance.setTransferPolicy(INDIVIDUAL, BANK, true);
        entityTypeCompliance.setTransferPolicy(INDIVIDUAL, CORPORATE, true);
        entityTypeCompliance.setTransferPolicy(INDIVIDUAL, INDIVIDUAL, true);

        // Exchange portals can transfer to all entity types
        entityTypeCompliance.setTransferPolicy(EXCHANGE_PORTAL, BANK, true);
        entityTypeCompliance.setTransferPolicy(
            EXCHANGE_PORTAL,
            CORPORATE,
            true
        );
        entityTypeCompliance.setTransferPolicy(
            EXCHANGE_PORTAL,
            INDIVIDUAL,
            true
        );
        entityTypeCompliance.setTransferPolicy(
            EXCHANGE_PORTAL,
            EXCHANGE_PORTAL,
            true
        );

        // All entity types can receive from exchange portals
        entityTypeCompliance.setTransferPolicy(BANK, EXCHANGE_PORTAL, true);
        entityTypeCompliance.setTransferPolicy(
            CORPORATE,
            EXCHANGE_PORTAL,
            true
        );
        entityTypeCompliance.setTransferPolicy(
            INDIVIDUAL,
            EXCHANGE_PORTAL,
            true
        );
    }

    function _logDeploymentAddresses() internal {
        // Log deployment addresses using console2 from forge-std
        console2.log("Deployment completed:");
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
        console2.log("Verifier 1 Address:", verifier_1, "(DeriveKey 1)");
        console2.log("Verifier 2 Address:", verifier_2, "(DeriveKey 2)");
        console2.log("Blacklister 1 Address:", blacklister_1, "(DeriveKey 3)");
        console2.log("Blacklister 2 Address:", blacklister_2, "(DeriveKey 4)");
        console2.log("Transaction Type Admin:", txTypeAdmin, "(DeriveKey 5)");
        console2.log("EntityTypeCompliance:", address(entityTypeCompliance));
    }
}
