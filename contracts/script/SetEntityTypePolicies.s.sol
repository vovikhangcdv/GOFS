// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script, console2} from "forge-std/Script.sol";
import {BaseScript} from "./BaseScript.s.sol";
import {EntityType} from "../src/interfaces/ITypes.sol";
import {EntityLibrary} from "../src/libraries/EntityLibrary.sol";

/**
 * @title SetEntityTypePolicies
 * @dev Script to set up entity type transfer policies using the admin role
 */
contract SetEntityTypePolicies is BaseScript {
    // Use constants from EntityLibrary
    using EntityLibrary for EntityType;

    function run() public virtual {
        // Initialize base script
        __BaseScript_init();
        __BaseScript_initializeContracts();

        // Get current nonce to avoid nonce mismatch
        uint256 currentNonce = vm.getNonce(deployer);
        console2.log("Current nonce for admin:", currentNonce);

        // Start recording transactions for deployment
        vm.startBroadcast(vm.deriveKey(mnemonic, 0));

        // Setup all policies
        setupEntityTypePolicies();

        vm.stopBroadcast();

        // Log the completion
        console2.log(
            "Entity type transfer policies have been set up successfully"
        );
        logEntityTypePolicies();
        logEntityTypeMetadata();
    }

    function setupEntityTypePolicies() public {
        // Setup metadata for each entity type
        setupEntityTypeMetadata();

        // Setup policies for each entity type
        setupIndividualPolicies();
        setupPrivateEnterprisePolicies();
        setupLLCPolicies();
        setupJointStockCompanyPolicies();
        setupPartnershipPolicies();
        setupCooperativePolicies();
        setupHouseholdBusinessPolicies();
        setupForeignInvestorPolicies();
        setupExchangePortalPolicies();
    }

    function setupEntityTypeMetadata() internal {
        // Set metadata for each entity type
        entityRegistry.setEntityTypeMetadata(
            EntityLibrary.UNKNOWN,
            "UNKNOWN: Unknown or unclassified entity type"
        );
        entityRegistry.setEntityTypeMetadata(
            EntityLibrary.INDIVIDUAL,
            "INDIVIDUAL: Individual persons (Ca nhan)"
        );
        entityRegistry.setEntityTypeMetadata(
            EntityLibrary.PRIVATE_ENTERPRISE,
            "PRIVATE_ENTERPRISE: Private enterprises (Doanh nghiep tu nhan)"
        );
        entityRegistry.setEntityTypeMetadata(
            EntityLibrary.LLC_ONE_MEMBER,
            "LLC_ONE_MEMBER: Single member limited liability company (TNHH mot thanh vien)"
        );
        entityRegistry.setEntityTypeMetadata(
            EntityLibrary.LLC_MULTI_MEMBER,
            "LLC_MULTI_MEMBER: Multi-member limited liability company (TNHH hai thanh vien tro len)"
        );
        entityRegistry.setEntityTypeMetadata(
            EntityLibrary.JOINT_STOCK_COMPANY,
            "JOINT_STOCK_COMPANY: Joint stock company (Cong ty co phan)"
        );
        entityRegistry.setEntityTypeMetadata(
            EntityLibrary.PARTNERSHIP,
            "PARTNERSHIP: Partnership company (Cong ty hop danh)"
        );
        entityRegistry.setEntityTypeMetadata(
            EntityLibrary.COOPERATIVE,
            "COOPERATIVE: Cooperative (Hop tac xa)"
        );
        entityRegistry.setEntityTypeMetadata(
            EntityLibrary.UNION_OF_COOPERATIVES,
            "UNION_OF_COOPERATIVES: Union of cooperatives (Lien hiep hop tac xa)"
        );
        entityRegistry.setEntityTypeMetadata(
            EntityLibrary.COOPERATIVE_GROUP,
            "COOPERATIVE_GROUP: Cooperative group (To hop tac)"
        );
        entityRegistry.setEntityTypeMetadata(
            EntityLibrary.HOUSEHOLD_BUSINESS,
            "HOUSEHOLD_BUSINESS: Household business (Ho kinh doanh)"
        );
        entityRegistry.setEntityTypeMetadata(
            EntityLibrary.FOREIGN_INDIVIDUAL_INVESTOR,
            "FOREIGN_INDIVIDUAL_INVESTOR: Foreign individual investor (Nha dau tu ca nhan nuoc ngoai)"
        );
        entityRegistry.setEntityTypeMetadata(
            EntityLibrary.FOREIGN_ORGANIZATION_INVESTOR,
            "FOREIGN_ORGANIZATION_INVESTOR: Foreign organization investor (Nha dau tu to chuc nuoc ngoai)"
        );
        entityRegistry.setEntityTypeMetadata(
            EntityLibrary.FOREIGN_INVESTED_ECONOMIC_ORGANIZATION,
            "FOREIGN_INVESTED_ECONOMIC_ORGANIZATION: Foreign invested economic organization (To chuc kinh te co von dau tu nuoc ngoai)"
        );
        entityRegistry.setEntityTypeMetadata(
            EntityLibrary.EXCHANGE_PORTAL,
            "EXCHANGE_PORTAL: Exchange portal (Cong giao dich ngoai te)"
        );
    }

    function setupIndividualPolicies() internal {
        // Individuals can transfer to most domestic entities
        EntityType[] memory allowedTypes = new EntityType[](8);
        allowedTypes[0] = EntityLibrary.INDIVIDUAL;
        allowedTypes[1] = EntityLibrary.PRIVATE_ENTERPRISE;
        allowedTypes[2] = EntityLibrary.LLC_ONE_MEMBER;
        allowedTypes[3] = EntityLibrary.LLC_MULTI_MEMBER;
        allowedTypes[4] = EntityLibrary.JOINT_STOCK_COMPANY;
        allowedTypes[5] = EntityLibrary.PARTNERSHIP;
        allowedTypes[6] = EntityLibrary.COOPERATIVE;
        allowedTypes[7] = EntityLibrary.HOUSEHOLD_BUSINESS;

        bool[] memory alloweds = new bool[](allowedTypes.length);
        for (uint i = 0; i < allowedTypes.length; i++) {
            alloweds[i] = true;
        }

        entityTypeCompliance.setTransferPolicy(
            EntityLibrary.INDIVIDUAL,
            allowedTypes,
            alloweds
        );
    }

    function setupPrivateEnterprisePolicies() internal {
        // Private enterprises can transfer to all entity types
        EntityType[] memory allTypes = getAllEntityTypes();
        bool[] memory alloweds = new bool[](allTypes.length);
        for (uint i = 0; i < allTypes.length; i++) {
            alloweds[i] = true;
        }
        entityTypeCompliance.setTransferPolicy(
            EntityLibrary.PRIVATE_ENTERPRISE,
            allTypes,
            alloweds
        );
    }

    function setupLLCPolicies() internal {
        // Both LLC types can transfer to all entity types
        EntityType[] memory allTypes = getAllEntityTypes();
        bool[] memory alloweds = new bool[](allTypes.length);
        for (uint i = 0; i < allTypes.length; i++) {
            alloweds[i] = true;
        }

        // Set policies for both LLC types
        entityTypeCompliance.setTransferPolicy(
            EntityLibrary.LLC_ONE_MEMBER,
            allTypes,
            alloweds
        );
        entityTypeCompliance.setTransferPolicy(
            EntityLibrary.LLC_MULTI_MEMBER,
            allTypes,
            alloweds
        );
    }

    function setupJointStockCompanyPolicies() internal {
        // Joint stock companies can transfer to all entity types
        EntityType[] memory allTypes = getAllEntityTypes();
        bool[] memory alloweds = new bool[](allTypes.length);
        for (uint i = 0; i < allTypes.length; i++) {
            alloweds[i] = true;
        }
        entityTypeCompliance.setTransferPolicy(
            EntityLibrary.JOINT_STOCK_COMPANY,
            allTypes,
            alloweds
        );
    }

    function setupPartnershipPolicies() internal {
        // Partnerships can transfer to most domestic entities
        EntityType[] memory allowedTypes = new EntityType[](7);
        allowedTypes[0] = EntityLibrary.INDIVIDUAL;
        allowedTypes[1] = EntityLibrary.PRIVATE_ENTERPRISE;
        allowedTypes[2] = EntityLibrary.LLC_ONE_MEMBER;
        allowedTypes[3] = EntityLibrary.LLC_MULTI_MEMBER;
        allowedTypes[4] = EntityLibrary.JOINT_STOCK_COMPANY;
        allowedTypes[5] = EntityLibrary.PARTNERSHIP;
        allowedTypes[6] = EntityLibrary.COOPERATIVE;

        bool[] memory alloweds = new bool[](allowedTypes.length);
        for (uint i = 0; i < allowedTypes.length; i++) {
            alloweds[i] = true;
        }

        entityTypeCompliance.setTransferPolicy(
            EntityLibrary.PARTNERSHIP,
            allowedTypes,
            alloweds
        );
    }

    function setupCooperativePolicies() internal {
        // Setup policies for cooperatives and their unions/groups
        EntityType[] memory cooperativeTypes = new EntityType[](3);
        cooperativeTypes[0] = EntityLibrary.COOPERATIVE;
        cooperativeTypes[1] = EntityLibrary.UNION_OF_COOPERATIVES;
        cooperativeTypes[2] = EntityLibrary.COOPERATIVE_GROUP;

        EntityType[] memory allowedTypes = new EntityType[](4);
        allowedTypes[0] = EntityLibrary.INDIVIDUAL;
        allowedTypes[1] = EntityLibrary.COOPERATIVE;
        allowedTypes[2] = EntityLibrary.UNION_OF_COOPERATIVES;
        allowedTypes[3] = EntityLibrary.COOPERATIVE_GROUP;

        bool[] memory alloweds = new bool[](allowedTypes.length);
        for (uint i = 0; i < allowedTypes.length; i++) {
            alloweds[i] = true;
        }

        // Set policies for each cooperative type
        for (uint i = 0; i < cooperativeTypes.length; i++) {
            entityTypeCompliance.setTransferPolicy(
                cooperativeTypes[i],
                allowedTypes,
                alloweds
            );
        }
    }

    function setupHouseholdBusinessPolicies() internal {
        // Household businesses can transfer to individuals and other household businesses
        EntityType[] memory allowedTypes = new EntityType[](3);
        allowedTypes[0] = EntityLibrary.INDIVIDUAL;
        allowedTypes[1] = EntityLibrary.HOUSEHOLD_BUSINESS;
        allowedTypes[2] = EntityLibrary.PRIVATE_ENTERPRISE;

        bool[] memory alloweds = new bool[](allowedTypes.length);
        for (uint i = 0; i < allowedTypes.length; i++) {
            alloweds[i] = true;
        }

        entityTypeCompliance.setTransferPolicy(
            EntityLibrary.HOUSEHOLD_BUSINESS,
            allowedTypes,
            alloweds
        );
    }

    function setupForeignInvestorPolicies() internal {
        // Setup policies for foreign investors and organizations
        EntityType[] memory foreignTypes = new EntityType[](3);
        foreignTypes[0] = EntityLibrary.FOREIGN_INDIVIDUAL_INVESTOR;
        foreignTypes[1] = EntityLibrary.FOREIGN_ORGANIZATION_INVESTOR;
        foreignTypes[2] = EntityLibrary.FOREIGN_INVESTED_ECONOMIC_ORGANIZATION;

        EntityType[] memory domesticTypes = getAllEntityTypes();
        bool[] memory alloweds = new bool[](domesticTypes.length);
        for (uint i = 0; i < domesticTypes.length; i++) {
            alloweds[i] = true;
        }

        // Foreign entities can transfer to all domestic entities
        for (uint i = 0; i < foreignTypes.length; i++) {
            entityTypeCompliance.setTransferPolicy(
                foreignTypes[i],
                domesticTypes,
                alloweds
            );
        }
    }

    function setupExchangePortalPolicies() internal {
        // Exchange portal can transfer to all entity types
        EntityType[] memory allTypes = getAllEntityTypes();
        bool[] memory alloweds = new bool[](allTypes.length);
        for (uint i = 0; i < allTypes.length; i++) {
            alloweds[i] = true;
        }
        entityTypeCompliance.setTransferPolicy(
            EntityLibrary.EXCHANGE_PORTAL,
            allTypes,
            alloweds
        );
    }

    function getAllEntityTypes() internal pure returns (EntityType[] memory) {
        EntityType[] memory types = new EntityType[](15);
        types[0] = EntityLibrary.UNKNOWN;
        types[1] = EntityLibrary.INDIVIDUAL;
        types[2] = EntityLibrary.PRIVATE_ENTERPRISE;
        types[3] = EntityLibrary.LLC_ONE_MEMBER;
        types[4] = EntityLibrary.LLC_MULTI_MEMBER;
        types[5] = EntityLibrary.JOINT_STOCK_COMPANY;
        types[6] = EntityLibrary.PARTNERSHIP;
        types[7] = EntityLibrary.COOPERATIVE;
        types[8] = EntityLibrary.UNION_OF_COOPERATIVES;
        types[9] = EntityLibrary.COOPERATIVE_GROUP;
        types[10] = EntityLibrary.HOUSEHOLD_BUSINESS;
        types[11] = EntityLibrary.FOREIGN_INDIVIDUAL_INVESTOR;
        types[12] = EntityLibrary.FOREIGN_ORGANIZATION_INVESTOR;
        types[13] = EntityLibrary.FOREIGN_INVESTED_ECONOMIC_ORGANIZATION;
        types[14] = EntityLibrary.EXCHANGE_PORTAL;
        return types;
    }

    function logEntityTypePolicies() internal view {
        EntityType[] memory allTypes = getAllEntityTypes();
        string[] memory typeNames = getTypeNames();

        console2.log("\n========== Entity Type Transfer Policies ==========");

        // Print header
        string memory header = "From \\ To |";
        for (uint j = 0; j < typeNames.length; j++) {
            header = string(abi.encodePacked(header, " ", typeNames[j], " |"));
        }
        console2.log(header);

        // Print separator
        string memory separator = "----------|";
        for (uint j = 0; j < typeNames.length; j++) {
            separator = string(abi.encodePacked(separator, "---------|"));
        }
        console2.log(separator);

        // Print matrix
        for (uint i = 0; i < allTypes.length; i++) {
            string memory line = string(abi.encodePacked(typeNames[i], " |"));
            for (uint j = 0; j < allTypes.length; j++) {
                bool allowed = entityTypeCompliance.isTransferAllowed(
                    allTypes[i],
                    allTypes[j]
                );
                line = string(
                    abi.encodePacked(
                        line,
                        allowed ? "    Y    |" : "    N    |"
                    )
                );
            }
            console2.log(line);
        }
    }

    function logEntityTypeMetadata() internal view {
        EntityType[] memory allTypes = getAllEntityTypes();
        string[] memory typeNames = getTypeNames();

        console2.log("\nEntity Type Metadata:");
        console2.log("-------------------");

        for (uint i = 0; i < allTypes.length; i++) {
            console2.log(
                string(
                    abi.encodePacked(
                        typeNames[i],
                        ": ",
                        entityRegistry.getEntityTypeMetadata(allTypes[i])
                    )
                )
            );
        }
    }

    function getTypeNames() internal pure returns (string[] memory) {
        string[] memory names = new string[](15);
        names[0] = "UNKNOWN";
        names[1] = "INDIV";
        names[2] = "PRIV_ENT";
        names[3] = "LLC_ONE";
        names[4] = "LLC_MULTI";
        names[5] = "JSC";
        names[6] = "PARTNER";
        names[7] = "COOP";
        names[8] = "COOP_UNI";
        names[9] = "COOP_GRP";
        names[10] = "HOUSE_BIZ";
        names[11] = "FOR_INDIV";
        names[12] = "FOR_ORG";
        names[13] = "FOR_ECO";
        names[14] = "EXCHANGE_PORTAL";
        return names;
    }
}
