// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script, console2} from "forge-std/Script.sol";
import {EntityRegistry} from "../src/EntityRegistry.sol";
import {ComplianceRegistry} from "../src/ComplianceRegistry.sol";
import {CompliantToken} from "../src/CompliantToken.sol";
import {ExchangePortal} from "../src/ExchangePortal.sol";
import {EntityType, Entity} from "../src/interfaces/ITypes.sol"; // Updated import
import {AddressRestrictionCompliance} from "../src/compliances/AddressRestrictionCompliance.sol";
import {VerificationCompliance} from "../src/compliances/VerificationCompliance.sol";
import {SupplyCompliance} from "../src/compliances/SupplyCompliance.sol";
import {MERC20} from "../test/mocks/MockERC20.sol"; // Mock ERC20 for testing
import {Entity, EntityType} from "../src/interfaces/ITypes.sol"; // Updated import
contract SimpleDeployAll is Script {
    EntityRegistry public entityRegistry;
    ComplianceRegistry public complianceRegistry;
    CompliantToken public evndToken;
    ExchangePortal public exchangePortal;
    AddressRestrictionCompliance public addressRestrictionCompliance;
    VerificationCompliance public verificationCompliance;
    SupplyCompliance public supplyCompliance;

    // Constants for initial setup
    uint256 constant INITIAL_EXCHANGE_RATE = 24 * 1e18; // 24:1 initial rate
    uint256 constant INITIAL_EXCHANGE_FEE = 10; // 0.1% fee (in basis points)

    // Addresses
    address deployer;
    address verifier_1;
    address verifier_2;
    address blacklister_1;
    address blacklister_2;

    function run() public {
        // Get deployer from mnemonic
        string memory mnemonic = vm.envString("WALLET_MEMO");
        uint256 deployerPrivateKey = vm.deriveKey(mnemonic, 0);
        deployer = vm.addr(deployerPrivateKey);

        // Start recording transactions for deployment
        vm.startBroadcast(deployerPrivateKey);

        // 1. Deploy EntityRegistry
        entityRegistry = new EntityRegistry();
        EntityType[] memory entityTypes = new EntityType[](3); // Changed from 2 to 3 to include exchange type
        entityTypes[0] = EntityType.wrap(1); // INDIVIDUAL
        entityTypes[1] = EntityType.wrap(2); // ORGANIZATION
        entityTypes[2] = EntityType.wrap(10); // EXCHANGE_PORTAL
        verifier_1 = vm.addr(vm.deriveKey(mnemonic, 1)); // First verifier address
        verifier_2 = vm.addr(vm.deriveKey(mnemonic, 2)); // Second verifier address
        entityRegistry.addVerifier(verifier_1, entityTypes);
        entityRegistry.addVerifier(verifier_2, entityTypes);

        // 2. Deploy ComplianceRegistry and compliances
        complianceRegistry = new ComplianceRegistry();
        // 3. Deploy CompliantToken (e-VND)
        evndToken = new CompliantToken(
            "Vietnamese Digital Currency",
            "eVND",
            address(complianceRegistry)
        );

        // 4. Deploy compliance modules
        addressRestrictionCompliance = new AddressRestrictionCompliance();
        verificationCompliance = new VerificationCompliance(
            address(entityRegistry)
        );
        supplyCompliance = new SupplyCompliance(address(evndToken));
        complianceRegistry.addModule(address(addressRestrictionCompliance));
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
        complianceRegistry.addModule(address(verificationCompliance));
        complianceRegistry.addModule(address(supplyCompliance));
        supplyCompliance.setMaxSupply(
            1_000_000_000 * 10 ** evndToken.decimals() // Set supply limit to 1 billion eVND
        );
        // 5. Deploy ExchangePortal
        MERC20 mUSD = new MERC20("Mock US Dollar", "mUSD");
        exchangePortal = new ExchangePortal(
            address(mUSD),
            address(evndToken),
            INITIAL_EXCHANGE_RATE,
            deployer, // Treasury is set to deployer initially
            INITIAL_EXCHANGE_FEE
        );

        // Register ExchangePortal as an entity
        {
            // Create entity data
            bytes memory entityData = abi.encode(
                "ExchangePortal", // name
                "Official Exchange Portal" // info
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

        mUSD.mint(address(exchangePortal), 1000000 * 10 ** 18); // Mint 1 million mUSD to portal
        evndToken.mint(
            address(exchangePortal),
            1000000 * 10 ** evndToken.decimals()
        ); // Mint 1 million eVND to portal

        vm.stopBroadcast();

        // Log deployment addresses using console2 from forge-std
        console2.log("Deployment completed:");
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
        console2.log("Verifier 1 Address:", verifier_1, "(DeriveKey 1)");
        console2.log("Verifier 2 Address:", verifier_2, "(DeriveKey 2)");
        console2.log("Blacklister 1 Address:", blacklister_1, "(DeriveKey 3)");
        console2.log("Blacklister 2 Address:", blacklister_2, "(DeriveKey 4)");
    }
}
