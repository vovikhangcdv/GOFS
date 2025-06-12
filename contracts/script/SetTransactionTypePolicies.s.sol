// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script, console2} from "forge-std/Script.sol";
import {BaseScript} from "./BaseScript.s.sol";
import {EntityType, TxType} from "../src/interfaces/ITypes.sol";
import {EntityLibrary} from "../src/libraries/EntityLibrary.sol";

/**
 * @title SetTransactionTypePolicies
 * @dev Script to set up transaction type transfer policies using the admin role
 */
contract SetTransactionTypePolicies is BaseScript {
    // Transaction Types
    TxType constant UNKNOWN_TYPE = TxType.wrap(0);
    TxType constant PAYMENT_TYPE = TxType.wrap(1);
    TxType constant GIVE_TYPE = TxType.wrap(2);
    TxType constant LOAN_TYPE = TxType.wrap(3);
    TxType constant INVESTMENT_TYPE = TxType.wrap(4);
    TxType constant EXCHANGE_TYPE = TxType.wrap(5);

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
        setupTransactionTypePolicies();

        vm.stopBroadcast();

        // Log the completion
        console2.log(
            "Transaction type transfer policies have been set up successfully"
        );
        logTransactionTypePolicies();
        logTransactionTypeMetadata();
    }

    function setupTransactionTypePolicies() public {
        // Setup metadata for each transaction type
        setupTransactionTypeMetadata();
        // Setup policies for each transaction type
        setupUnknownPolicies();
        setupPaymentPolicies();
        setupGivePolicies();
        setupLoanPolicies();
        setupInvestmentPolicies();
        setupExchangePolicies();
    }

    function setupTransactionTypeMetadata() internal {
        // Set metadata for each transaction type
        transactionTypeCompliance.setTransactionTypeMetadata(
            UNKNOWN_TYPE,
            "UNKNOWN: Default transaction type for unclassified transfers"
        );
        transactionTypeCompliance.setTransactionTypeMetadata(
            PAYMENT_TYPE,
            "PAYMENT: Commercial transactions for goods and services"
        );
        transactionTypeCompliance.setTransactionTypeMetadata(
            GIVE_TYPE,
            "GIVE: Non-commercial transfers and gifts between entities"
        );
        transactionTypeCompliance.setTransactionTypeMetadata(
            LOAN_TYPE,
            "LOAN: Financial transactions involving lending and borrowing between authorized entities"
        );
        transactionTypeCompliance.setTransactionTypeMetadata(
            INVESTMENT_TYPE,
            "INVESTMENT: Capital investment transactions including securities and business investments"
        );
        transactionTypeCompliance.setTransactionTypeMetadata(
            EXCHANGE_TYPE,
            "EXCHANGE: Currency exchange and trading transactions through authorized portals"
        );
    }

    function setupUnknownPolicies() internal {
        // UNKNOWN type can be used by most basic domestic entities
        EntityType[] memory fromTypes = new EntityType[](5);
        fromTypes[0] = EntityLibrary.INDIVIDUAL;
        fromTypes[1] = EntityLibrary.PRIVATE_ENTERPRISE;
        fromTypes[2] = EntityLibrary.HOUSEHOLD_BUSINESS;
        fromTypes[3] = EntityLibrary.UNKNOWN;
        fromTypes[4] = EntityLibrary.EXCHANGE_PORTAL;

        transactionTypeCompliance.setTransactionTypePolicy(
            UNKNOWN_TYPE,
            fromTypes,
            fromTypes
        );
    }

    function setupPaymentPolicies() internal {
        // All registered entities except UNKNOWN can participate in payments
        EntityType[] memory fromTypes = new EntityType[](14);
        fromTypes[0] = EntityLibrary.INDIVIDUAL;
        fromTypes[1] = EntityLibrary.PRIVATE_ENTERPRISE;
        fromTypes[2] = EntityLibrary.LLC_ONE_MEMBER;
        fromTypes[3] = EntityLibrary.LLC_MULTI_MEMBER;
        fromTypes[4] = EntityLibrary.JOINT_STOCK_COMPANY;
        fromTypes[5] = EntityLibrary.PARTNERSHIP;
        fromTypes[6] = EntityLibrary.COOPERATIVE;
        fromTypes[7] = EntityLibrary.UNION_OF_COOPERATIVES;
        fromTypes[8] = EntityLibrary.COOPERATIVE_GROUP;
        fromTypes[9] = EntityLibrary.HOUSEHOLD_BUSINESS;
        fromTypes[10] = EntityLibrary.FOREIGN_INDIVIDUAL_INVESTOR;
        fromTypes[11] = EntityLibrary.FOREIGN_ORGANIZATION_INVESTOR;
        fromTypes[12] = EntityLibrary.FOREIGN_INVESTED_ECONOMIC_ORGANIZATION;
        fromTypes[13] = EntityLibrary.UNKNOWN;

        transactionTypeCompliance.setTransactionTypePolicy(
            PAYMENT_TYPE,
            fromTypes,
            fromTypes // Same types can receive payments
        );
    }

    function setupGivePolicies() internal {
        // Individuals and organizations can give
        EntityType[] memory fromTypes = new EntityType[](14);
        fromTypes[0] = EntityLibrary.INDIVIDUAL;
        fromTypes[1] = EntityLibrary.PRIVATE_ENTERPRISE;
        fromTypes[2] = EntityLibrary.LLC_ONE_MEMBER;
        fromTypes[3] = EntityLibrary.LLC_MULTI_MEMBER;
        fromTypes[4] = EntityLibrary.JOINT_STOCK_COMPANY;
        fromTypes[5] = EntityLibrary.PARTNERSHIP;
        fromTypes[6] = EntityLibrary.COOPERATIVE;
        fromTypes[7] = EntityLibrary.UNION_OF_COOPERATIVES;
        fromTypes[8] = EntityLibrary.COOPERATIVE_GROUP;
        fromTypes[9] = EntityLibrary.HOUSEHOLD_BUSINESS;
        fromTypes[10] = EntityLibrary.FOREIGN_INDIVIDUAL_INVESTOR;
        fromTypes[11] = EntityLibrary.FOREIGN_ORGANIZATION_INVESTOR;
        fromTypes[12] = EntityLibrary.FOREIGN_INVESTED_ECONOMIC_ORGANIZATION;
        fromTypes[13] = EntityLibrary.EXCHANGE_PORTAL;

        // Most entities can receive gifts
        EntityType[] memory toTypes = new EntityType[](3);
        toTypes[0] = EntityLibrary.INDIVIDUAL;
        toTypes[1] = EntityLibrary.HOUSEHOLD_BUSINESS;
        toTypes[2] = EntityLibrary.FOREIGN_INDIVIDUAL_INVESTOR;

        transactionTypeCompliance.setTransactionTypePolicy(
            GIVE_TYPE,
            fromTypes,
            toTypes
        );
    }

    function setupLoanPolicies() internal {
        EntityType[] memory fromTypes = new EntityType[](5);
        fromTypes[0] = EntityLibrary.INDIVIDUAL; // Individuals can lend
        fromTypes[1] = EntityLibrary.PRIVATE_ENTERPRISE; // Private banks
        fromTypes[2] = EntityLibrary.LLC_ONE_MEMBER; // Bank subsidiaries
        fromTypes[3] = EntityLibrary.LLC_MULTI_MEMBER; // Bank consortiums
        fromTypes[4] = EntityLibrary.JOINT_STOCK_COMPANY; // Joint stock banks

        // Most business entities and individuals can receive loans
        EntityType[] memory toTypes = new EntityType[](7);
        toTypes[0] = EntityLibrary.INDIVIDUAL;
        toTypes[1] = EntityLibrary.PRIVATE_ENTERPRISE;
        toTypes[2] = EntityLibrary.LLC_ONE_MEMBER;
        toTypes[3] = EntityLibrary.LLC_MULTI_MEMBER;
        toTypes[4] = EntityLibrary.JOINT_STOCK_COMPANY;
        toTypes[5] = EntityLibrary.PARTNERSHIP;
        toTypes[6] = EntityLibrary.HOUSEHOLD_BUSINESS;

        transactionTypeCompliance.setTransactionTypePolicy(
            LOAN_TYPE,
            fromTypes,
            toTypes
        );
    }

    function setupInvestmentPolicies() internal {
        // Organizations and foreign investors can make investments
        EntityType[] memory fromTypes = new EntityType[](9);
        fromTypes[0] = EntityLibrary.INDIVIDUAL;
        fromTypes[1] = EntityLibrary.PRIVATE_ENTERPRISE;
        fromTypes[2] = EntityLibrary.LLC_ONE_MEMBER;
        fromTypes[3] = EntityLibrary.LLC_MULTI_MEMBER;
        fromTypes[4] = EntityLibrary.JOINT_STOCK_COMPANY;
        fromTypes[5] = EntityLibrary.PARTNERSHIP;
        fromTypes[6] = EntityLibrary.FOREIGN_INDIVIDUAL_INVESTOR;
        fromTypes[7] = EntityLibrary.FOREIGN_ORGANIZATION_INVESTOR;
        fromTypes[8] = EntityLibrary.FOREIGN_INVESTED_ECONOMIC_ORGANIZATION;

        // Most business entities can receive investments
        EntityType[] memory toTypes = new EntityType[](7);
        toTypes[0] = EntityLibrary.PRIVATE_ENTERPRISE;
        toTypes[1] = EntityLibrary.LLC_ONE_MEMBER;
        toTypes[2] = EntityLibrary.LLC_MULTI_MEMBER;
        toTypes[3] = EntityLibrary.JOINT_STOCK_COMPANY;
        toTypes[4] = EntityLibrary.PARTNERSHIP;
        toTypes[5] = EntityLibrary.COOPERATIVE;
        toTypes[6] = EntityLibrary.UNION_OF_COOPERATIVES;

        transactionTypeCompliance.setTransactionTypePolicy(
            INVESTMENT_TYPE,
            fromTypes,
            toTypes
        );
    }

    function setupExchangePolicies() internal {
        // Exchange transactions can only be initiated by exchange portals
        EntityType[] memory fromTypes = new EntityType[](14);
        fromTypes[0] = EntityLibrary.INDIVIDUAL;
        fromTypes[1] = EntityLibrary.PRIVATE_ENTERPRISE;
        fromTypes[2] = EntityLibrary.LLC_ONE_MEMBER;
        fromTypes[3] = EntityLibrary.LLC_MULTI_MEMBER;
        fromTypes[4] = EntityLibrary.JOINT_STOCK_COMPANY;
        fromTypes[5] = EntityLibrary.PARTNERSHIP;
        fromTypes[6] = EntityLibrary.COOPERATIVE;
        fromTypes[7] = EntityLibrary.UNION_OF_COOPERATIVES;
        fromTypes[8] = EntityLibrary.COOPERATIVE_GROUP;
        fromTypes[9] = EntityLibrary.HOUSEHOLD_BUSINESS;
        fromTypes[10] = EntityLibrary.FOREIGN_INDIVIDUAL_INVESTOR;
        fromTypes[11] = EntityLibrary.FOREIGN_ORGANIZATION_INVESTOR;
        fromTypes[12] = EntityLibrary.FOREIGN_INVESTED_ECONOMIC_ORGANIZATION;
        fromTypes[13] = EntityLibrary.EXCHANGE_PORTAL;
        transactionTypeCompliance.setTransactionTypePolicy(
            EXCHANGE_TYPE,
            fromTypes,
            fromTypes
        );
    }

    function logTransactionTypePolicies() internal view {
        TxType[] memory txTypes = new TxType[](6);
        txTypes[0] = UNKNOWN_TYPE;
        txTypes[1] = PAYMENT_TYPE;
        txTypes[2] = GIVE_TYPE;
        txTypes[3] = LOAN_TYPE;
        txTypes[4] = INVESTMENT_TYPE;
        txTypes[5] = EXCHANGE_TYPE;

        string[] memory txTypeNames = new string[](6);
        txTypeNames[0] = "UNKNOWN";
        txTypeNames[1] = "PAYMENT";
        txTypeNames[2] = "GIVE";
        txTypeNames[3] = "LOAN";
        txTypeNames[4] = "INVESTMENT";
        txTypeNames[5] = "EXCHANGE";

        console2.log(
            "\n========== Transaction Type Transfer Policies =========="
        );
        console2.log("------------------------");

        for (uint i = 0; i < txTypes.length; i++) {
            console2.log(string(abi.encodePacked("\n", txTypeNames[i], ":")));

            EntityType[] memory fromTypes = transactionTypeCompliance
                .getAllowedFromEntityTypes(txTypes[i]);
            EntityType[] memory toTypes = transactionTypeCompliance
                .getAllowedToEntityTypes(txTypes[i]);

            console2.log("From Types:");
            for (uint j = 0; j < fromTypes.length; j++) {
                console2.log("\t", getEntityTypeName(fromTypes[j]));
            }

            console2.log("To Types:");
            for (uint j = 0; j < toTypes.length; j++) {
                console2.log("\t", getEntityTypeName(toTypes[j]));
            }

            console2.log(
                string(
                    abi.encodePacked(
                        "Usable: ",
                        transactionTypeCompliance.isTransactionTypeUsable(
                            txTypes[i]
                        )
                            ? "Yes"
                            : "No"
                    )
                )
            );
        }
    }

    function logTransactionTypeMetadata() internal view {
        TxType[] memory txTypes = new TxType[](6);
        txTypes[0] = UNKNOWN_TYPE;
        txTypes[1] = PAYMENT_TYPE;
        txTypes[2] = GIVE_TYPE;
        txTypes[3] = LOAN_TYPE;
        txTypes[4] = INVESTMENT_TYPE;
        txTypes[5] = EXCHANGE_TYPE;

        console2.log("\nTransaction Type Metadata:");
        console2.log("-------------------------");

        for (uint i = 0; i < txTypes.length; i++) {
            console2.log(
                transactionTypeCompliance.getTransactionTypeMetadata(txTypes[i])
            );
        }
    }

    function getEntityTypeName(
        EntityType entityType
    ) internal pure returns (string memory) {
        uint8 id = EntityType.unwrap(entityType);
        if (id == 0) return "UNKNOWN";
        if (id == 1) return "INDIVIDUAL";
        if (id == 2) return "PRIVATE_ENTERPRISE";
        if (id == 3) return "LLC_ONE_MEMBER";
        if (id == 4) return "LLC_MULTI_MEMBER";
        if (id == 5) return "JOINT_STOCK_COMPANY";
        if (id == 6) return "PARTNERSHIP";
        if (id == 7) return "COOPERATIVE";
        if (id == 8) return "UNION_OF_COOPERATIVES";
        if (id == 9) return "COOPERATIVE_GROUP";
        if (id == 10) return "HOUSEHOLD_BUSINESS";
        if (id == 11) return "FOREIGN_INDIVIDUAL_INVESTOR";
        if (id == 12) return "FOREIGN_ORGANIZATION_INVESTOR";
        if (id == 13) return "FOREIGN_INVESTED_ECONOMIC_ORGANIZATION";
        if (id == 14) return "EXCHANGE_PORTAL";
        return "UNKNOWN";
    }
}
