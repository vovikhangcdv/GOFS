// GOFS/e-VND CBDC System - Frontend Contract Interaction Demo
// Inspired by .vscode/getEntityName-minimal.js but expanded for full system
const { ethers } = require('ethers');


// Contract addresses (update these with your deployed addresses)
// == Logs ==
// Deployer Address (ADMIN FOR ALL): 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 (DeriveKey 0)
// ProxyAdmin: 0x5FbDB2315678afecb367f032d93F642f64180aa3
// EntityRegistry: 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0
// ComplianceRegistry: 0x0165878A594ca255338adfa4d48449f69242Eb8F
// eVND Token: 0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6
// Exchange Portal: 0x67d269191c92Caf3cD7723F116c85e6E9bf55933
// AddressRestrictionCompliance: 0x610178dA211FEF7D417bC0e6FeD39F05609AD788
// VerificationCompliance: 0xA51c1fc2f0D1a1b8494Ed1FE312d7C3a78Ed91C0
// SupplyCompliance: 0x9A676e781A523b5d0C0e43731313A708CB607508
// TransactionTypeCompliance: 0x959922bE3CAee4b8Cd9a407cc3ac1C251C2007B1
// EntityTypeCompliance: 0x68B1D87F95878fE05B998F19b66F4baba5De1aed
// Verifier 1 Address: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 (DeriveKey 1)
// Verifier 2 Address: 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC (DeriveKey 2)
// Blacklister 1 Address: 0x90F79bf6EB2c4f870365E785982E1f101E93b906 (DeriveKey 3)
// Blacklister 2 Address: 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65 (DeriveKey 4)
const CONTRACT_ADDRESSES = {
    COMPLIANT_TOKEN: '0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6',        // e-VND Token
    EXCHANGE_PORTAL: '0x67d269191c92Caf3cD7723F116c85e6E9bf55933',        // Exchange Portal
    ENTITY_REGISTRY: '0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0',        // Entity Registry
    COMPLIANCE_REGISTRY: '0x0165878A594ca255338adfa4d48449f69242Eb8F',    // Compliance Registry
    PROXY_ADMIN: '0x5FbDB2315678afecb367f032d93F642f64180aa3',             // Proxy Admin
    COMPLIANCE_MODULES: {
        ADDRESS_RESTRICTION_COMPLIANCE: '0x610178dA211FEF7D417bC0e6FeD39F05609AD788',
        VERIFICATION_COMPLIANCE: '0xA51c1fc2f0D1a1b8494Ed1FE312d7C3a78Ed91C0',
        SUPPLY_COMPLIANCE: '0x9A676e781A523b5d0C0e43731313A708CB607508',
        TRANSACTION_TYPE_COMPLIANCE: '0x959922bE3CAee4b8Cd9a407cc3ac1C251C2007B1',
        ENTITY_TYPE_COMPLIANCE: '0x68B1D87F95878fE05B998F19b66F4baba5De1aed'
    }
};

// Load ABIs from contracts/out directory
const loadABI = (contractName) => {
    try {
        const fs = require('fs');
        const path = require('path');

        // Try the correct path structure: contracts/out/ContractName.sol/ContractName.json
        const abiPath = path.join(__dirname, '..', '..', 'contracts', 'out', `${contractName}.sol`, `${contractName}.json`);

        if (!fs.existsSync(abiPath)) {
            console.error(`ABI file not found at ${abiPath}. Please run 'forge build' in \`contracts\/\` folder to generate the ABI.`);
        }

        const contractData = JSON.parse(fs.readFileSync(abiPath, 'utf8'));

        if (!contractData.abi) {
            console.warn(`Warning: No ABI found in ${abiPath}`);
            return getMinimalABI(contractName);
        }

        // console.log(`Successfully loaded ABI for ${contractName} from ${abiPath}`);
        return contractData.abi;
    } catch (error) {
        console.warn(`Warning: Could not load ABI for ${contractName}: ${error.message}`);
        console.warn(`Falling back to minimal ABI for ${contractName}`);
        return getMinimalABI(contractName);
    }
};

// Main contract interaction class
class GOFSContractInteractions {
    constructor(rpcUrl, signer = null) {
        this.provider = new ethers.JsonRpcProvider(rpcUrl);
        this.signer = signer || this.provider;

        // Initialize contracts
        this.compliantToken = new ethers.Contract(
            CONTRACT_ADDRESSES.COMPLIANT_TOKEN,
            loadABI('CompliantToken'),
            this.signer
        );

        this.entityRegistry = new ethers.Contract(
            CONTRACT_ADDRESSES.ENTITY_REGISTRY,
            loadABI('EntityRegistry'),
            this.signer
        );

        this.exchangePortal = new ethers.Contract(
            CONTRACT_ADDRESSES.EXCHANGE_PORTAL,
            loadABI('ExchangePortal'),
            this.signer
        );

        this.complianceRegistry = new ethers.Contract(
            CONTRACT_ADDRESSES.COMPLIANCE_REGISTRY,
            loadABI('ComplianceRegistry'),
            this.signer
        );

        this.transactionTypeCompliance = new ethers.Contract(
            CONTRACT_ADDRESSES.COMPLIANCE_MODULES.TRANSACTION_TYPE_COMPLIANCE,
            loadABI('TransactionTypeCompliance'),
            this.signer
        );

        this.entityTypeCompliance = new ethers.Contract(
            CONTRACT_ADDRESSES.COMPLIANCE_MODULES.ENTITY_TYPE_COMPLIANCE,
            loadABI('EntityTypeCompliance'),
            this.signer
        );
    }

    // ============================================================================
    // 1. SYSTEM DASHBOARD FUNCTIONS (Homepage)
    // ============================================================================

    async getSystemOverview() {
        try {
            const [tokenName, tokenSymbol, totalSupply, decimals] = await Promise.all([
                this.compliantToken.name(),
                this.compliantToken.symbol(),
                this.compliantToken.totalSupply(),
                this.compliantToken.decimals()
            ]);

            const exchangeRate = (await this.exchangePortal.getExchangeRate()) / (await this.exchangePortal.EXCHANGE_RATE_DECIMALS());

            return {
                tokenInfo: {
                    name: tokenName,
                    symbol: tokenSymbol,
                    totalSupply: ethers.formatUnits(totalSupply, decimals),
                    decimals: decimals,
                    address: CONTRACT_ADDRESSES.COMPLIANT_TOKEN
                },
                exchangeRate: {
                    rate: exchangeRate,
                    portalAddress: CONTRACT_ADDRESSES.EXCHANGE_PORTAL
                },
                systemAddresses: CONTRACT_ADDRESSES
            };
        } catch (error) {
            console.error('Error getting system overview:', error);
            throw error;
        }
    }

    // ============================================================================
    // 2. e-VND TOKEN FUNCTIONS
    // ============================================================================

    async getTokenBalance(address) {
        try {
            const balance = await this.compliantToken.balanceOf(address);
            const decimals = await this.compliantToken.decimals();
            return ethers.formatUnits(balance, decimals);
        } catch (error) {
            console.error('Error getting token balance:', error);
            throw error;
        }
    }

    async transferTokens(to, amount, txType = 0) {
        try {
            const decimals = await this.compliantToken.decimals();
            const amountWei = ethers.parseUnits(amount.toString(), decimals);

            let tx;
            if (txType === 0) {
                tx = await this.compliantToken.transfer(to, amountWei);
            } else {
                tx = await this.compliantToken.transferWithType(to, amountWei, txType);
            }

            return await tx.wait();
        } catch (error) {
            console.error('Error transferring tokens:', error);
            throw error;
        }
    }

    async getEntityName(entityAddress) {
        try {
            const entity = await this.entityRegistry.getEntity(entityAddress);
            if (entity.entityAddress === ethers.ZeroAddress) {
                return null; // Entity not found
            }

            // Decode entity data (assuming it contains name as first field)
            const decoded = ethers.AbiCoder.defaultAbiCoder().decode(['string', 'bytes32'], entity.entityData);
            return decoded[0];
        } catch (error) {
            console.error('Error getting entity name:', error);
            return null;
        }
    }

    async isEntityVerified(entityAddress) {
        try {
            return await this.entityRegistry.isVerifiedEntity(entityAddress);
        } catch (error) {
            console.error('Error checking entity verification:', error);
            return false;
        }
    }

    // ============================================================================
    // 3. EXCHANGE PORTAL FUNCTIONS
    // ============================================================================

    async getExchangeRate() {
        try {
            return (await this.exchangePortal.getExchangeRate()) / (await this.exchangePortal.EXCHANGE_RATE_DECIMALS());
        } catch (error) {
            console.error('Error getting exchange rate:', error);
            throw error;
        }
    }

    async calculateExchangeAmount(amount, txType = 0) {
        try {
            const decimals = await this.compliantToken.decimals();
            const amountWei = ethers.parseUnits(amount.toString(), decimals);

            let exchangeAmount;
            if (txType === 0) {
                exchangeAmount = await this.exchangePortal.getExchangeAmount(amountWei, 0);
            } else {
                exchangeAmount = await this.exchangePortal.getExchangeAmountWithType(amountWei, 0, txType);
            }

            return ethers.formatUnits(exchangeAmount, decimals);
        } catch (error) {
            console.error('Error calculating exchange amount:', error);
            throw error;
        }
    }

    async performExchange(amount, txType = 0) {
        try {
            const decimals = await this.compliantToken.decimals();
            const amountWei = ethers.parseUnits(amount.toString(), decimals);

            let tx;
            if (txType === 0) {
                tx = await this.exchangePortal.exchange(this.signer.address, amountWei, 0);
            } else {
                tx = await this.exchangePortal.exchangeWithType(this.signer.address, amountWei, 0, txType);
            }

            return await tx.wait();
        } catch (error) {
            console.error('Error performing exchange:', error);
            throw error;
        }
    }

    // ============================================================================
    // 4. COMPLIANCE & ENTITY REGISTRY FUNCTIONS
    // ============================================================================

    async getComplianceModules() {
        try {
            const modules = await this.complianceRegistry.getModules();
            const moduleDetails = [];

            for (const moduleAddress of modules) {
                // call the moduleAddress .name()
                const module = new ethers.Contract(moduleAddress, loadABI('ICompliance'), this.signer);
                const moduleName = await module.name();

                moduleDetails.push({
                    name: moduleName,
                    address: moduleAddress,
                    isRegistered: true
                });
            }

            return moduleDetails;
        } catch (error) {
            console.error('Error getting compliance modules:', error);
            throw error;
        }
    }

    async getVerifierInfo(verifierAddress) {
        try {
            const allowedTypes = await this.entityRegistry.getVerifierAllowedTypes(verifierAddress);
            return {
                address: verifierAddress,
                allowedEntityTypes: allowedTypes,
                isActive: allowedTypes.length > 0
            };
        } catch (error) {
            console.error('Error getting verifier info:', error);
            throw error;
        }
    }

    async registerEntity(entityData, verifierSignature) {
        try {
            const entity = {
                entityAddress: entityData.address,
                entityType: entityData.type,
                entityData: entityData.encodedData,
                verifier: entityData.verifier
            };

            const tx = await this.entityRegistry.register(entity, verifierSignature);
            return await tx.wait();
        } catch (error) {
            console.error('Error registering entity:', error);
            throw error;
        }
    }

    // ============================================================================
    // 5. TRANSACTION MONITORING & EVENTS
    // ============================================================================

    async getTransactionHistory(address, fromBlock = 'latest', toBlock = 'latest') {
        try {
            const filter = this.compliantToken.filters.Transfer(address);
            const events = await this.compliantToken.queryFilter(filter, fromBlock, toBlock);

            const transactions = await Promise.all(events.map(async (event) => {
                const block = await event.getBlock();
                const entityName = await this.getEntityName(event.args.to);

                return {
                    hash: event.transactionHash,
                    from: event.args.from,
                    to: event.args.to,
                    toName: entityName,
                    amount: ethers.formatUnits(event.args.value, await this.compliantToken.decimals()),
                    blockNumber: event.blockNumber,
                    timestamp: block.timestamp
                };
            }));

            return transactions;
        } catch (error) {
            console.error('Error getting transaction history:', error);
            throw error;
        }
    }

    async getTypedTransactionHistory(address, fromBlock = 'latest', toBlock = 'latest') {
        try {
            const filter = this.compliantToken.filters.TypedTransfer(address);
            const events = await this.compliantToken.queryFilter(filter, fromBlock, toBlock);

            const transactions = await Promise.all(events.map(async (event) => {
                const block = await event.getBlock();
                const entityName = await this.getEntityName(event.args.to);

                return {
                    hash: event.transactionHash,
                    from: event.args.from,
                    to: event.args.to,
                    toName: entityName,
                    amount: ethers.formatUnits(event.args.value, await this.compliantToken.decimals()),
                    transactionType: event.args.txType,
                    blockNumber: event.blockNumber,
                    timestamp: block.timestamp
                };
            }));

            return transactions;
        } catch (error) {
            console.error('Error getting typed transaction history:', error);
            throw error;
        }
    }

    // ============================================================================
    // 6. UTILITY FUNCTIONS
    // ============================================================================

    async getEntityDetails(entityAddress) {
        try {
            const entity = await this.entityRegistry.getEntity(entityAddress);
            if (entity.entityAddress === ethers.ZeroAddress) {
                return null; // Entity not found
            }

            const isVerified = await this.entityRegistry.isVerifiedEntity(entityAddress);
            const balance = await this.getTokenBalance(entityAddress);

            // Decode entity data
            const decoded = ethers.AbiCoder.defaultAbiCoder().decode(['string', 'bytes32'], entity.entityData);
            const name = decoded[0];
            const infoHash = decoded[1];

            return {
                address: entityAddress,
                name: name,
                entityType: Number(entity.entityType),
                entityTypeName: this.getEntityTypeName(Number(entity.entityType)),
                infoHash: infoHash,
                verifier: entity.verifier,
                isVerified: isVerified,
                balance: balance
            };
        } catch (error) {
            console.error('Error getting entity details:', error);
            return null;
        }
    }

    async getSystemMetrics() {
        try {
            const [totalSupply, decimals, exchangeRate, rateDecimals] = await Promise.all([
                this.compliantToken.totalSupply(),
                this.compliantToken.decimals(),
                this.exchangePortal.exchangeRate(),
            ]);

            return {
                totalSupply: ethers.formatUnits(totalSupply, decimals),
                exchangeRate: ethers.formatUnits(exchangeRate, decimals),
                tokenDecimals: decimals,
            };
        } catch (error) {
            console.error('Error getting system metrics:', error);
            throw error;
        }
    }

    // ============================================================================
    // 4. METADATA FUNCTIONS
    // ============================================================================

    async getAllEntityTypeMetadata() {
        try {
            const entityTypes = [1, 2, 3, 10]; // BANK, CORPORATE, INDIVIDUAL, EXCHANGE_PORTAL
            const metadata = {};

            for (const typeId of entityTypes) {
                try {
                    const metadataString = await this.entityRegistry.getEntityTypeMetadata(typeId);
                    metadata[typeId] = {
                        id: typeId,
                        name: this.getEntityTypeName(typeId),
                        metadata: metadataString
                    };
                } catch (error) {
                    console.warn(`Could not get metadata for entity type ${typeId}:`, error.message);
                    metadata[typeId] = {
                        id: typeId,
                        name: this.getEntityTypeName(typeId),
                        metadata: "No metadata available"
                    };
                }
            }

            return metadata;
        } catch (error) {
            console.error('Error getting entity type metadata:', error);
            throw error;
        }
    }

    async getAllTransactionTypeMetadata() {
        try {
            const transactionTypes = [1, 2, 3, 4]; // LOAN, PAYMENT, INVESTMENT, EXCHANGE
            const metadata = {};

            for (const typeId of transactionTypes) {
                try {
                    const metadataString = await this.transactionTypeCompliance.getTransactionTypeMetadata(typeId);
                    metadata[typeId] = {
                        id: typeId,
                        name: this.getTransactionTypeName(typeId),
                        metadata: metadataString
                    };
                } catch (error) {
                    console.warn(`Could not get metadata for transaction type ${typeId}:`, error.message);
                    metadata[typeId] = {
                        id: typeId,
                        name: this.getTransactionTypeName(typeId),
                        metadata: "No metadata available"
                    };
                }
            }

            return metadata;
        } catch (error) {
            console.error('Error getting transaction type metadata:', error);
            throw error;
        }
    }

    async getEntityTypeMetadata(entityTypeId) {
        try {
            const metadata = await this.entityRegistry.getEntityTypeMetadata(entityTypeId);
            return {
                id: entityTypeId,
                name: this.getEntityTypeName(entityTypeId),
                metadata: metadata
            };
        } catch (error) {
            console.error(`Error getting entity type metadata for ${entityTypeId}:`, error);
            throw error;
        }
    }

    async getTransactionTypeMetadata(transactionTypeId) {
        try {
            const metadata = await this.transactionTypeCompliance.getTransactionTypeMetadata(transactionTypeId);
            return {
                id: transactionTypeId,
                name: this.getTransactionTypeName(transactionTypeId),
                metadata: metadata
            };
        } catch (error) {
            console.error(`Error getting transaction type metadata for ${transactionTypeId}:`, error);
            throw error;
        }
    }

    // Helper functions for type names
    getEntityTypeName(typeId) {
        const entityTypeNames = {
            1: 'BANK',
            2: 'CORPORATE',
            3: 'INDIVIDUAL',
            10: 'EXCHANGE_PORTAL'
        };
        return entityTypeNames[typeId] || `UNKNOWN_${typeId}`;
    }

    getTransactionTypeName(typeId) {
        const transactionTypeNames = {
            1: 'LOAN',
            2: 'PAYMENT',
            3: 'INVESTMENT',
            4: 'EXCHANGE'
        };
        return transactionTypeNames[typeId] || `UNKNOWN_${typeId}`;
    }

    // ============================================================================
    // 5. COMPLIANCE POLICY FUNCTIONS
    // ============================================================================

    async getTransactionTypePolicies() {
        try {
            const transactionTypes = [1, 2, 3, 4]; // LOAN, PAYMENT, INVESTMENT, EXCHANGE
            const policies = {};

            for (const typeId of transactionTypes) {
                try {
                    const [fromTypes, toTypes] = await Promise.all([
                        this.transactionTypeCompliance.getAllowedFromEntityTypes(typeId),
                        this.transactionTypeCompliance.getAllowedToEntityTypes(typeId)
                    ]);

                    policies[typeId] = {
                        id: typeId,
                        name: this.getTransactionTypeName(typeId),
                        allowedFromEntityTypes: fromTypes.map(type => ({
                            id: Number(type),
                            name: this.getEntityTypeName(Number(type))
                        })),
                        allowedToEntityTypes: toTypes.map(type => ({
                            id: Number(type),
                            name: this.getEntityTypeName(Number(type))
                        })),
                        isUsable: await this.transactionTypeCompliance.isTransactionTypeUsable(typeId)
                    };
                } catch (error) {
                    console.warn(`Could not get policies for transaction type ${typeId}:`, error.message);
                }
            }

            return policies;
        } catch (error) {
            console.error('Error getting transaction type policies:', error);
            throw error;
        }
    }

    async getEntityTypeTransferPolicies() {
        try {
            const entityTypes = [1, 2, 3, 10]; // BANK, CORPORATE, INDIVIDUAL, EXCHANGE_PORTAL
            const policies = {};

            for (const fromTypeId of entityTypes) {
                policies[fromTypeId] = {
                    id: fromTypeId,
                    name: this.getEntityTypeName(fromTypeId),
                    allowedTransfersTo: {}
                };

                for (const toTypeId of entityTypes) {
                    try {
                        const isAllowed = await this.entityTypeCompliance.isTransferAllowed(fromTypeId, toTypeId);
                        policies[fromTypeId].allowedTransfersTo[toTypeId] = {
                            id: toTypeId,
                            name: this.getEntityTypeName(toTypeId),
                            allowed: isAllowed
                        };
                    } catch (error) {
                        console.warn(`Could not check transfer policy from ${fromTypeId} to ${toTypeId}:`, error.message);
                        policies[fromTypeId].allowedTransfersTo[toTypeId] = {
                            id: toTypeId,
                            name: this.getEntityTypeName(toTypeId),
                            allowed: false
                        };
                    }
                }
            }

            return policies;
        } catch (error) {
            console.error('Error getting entity type transfer policies:', error);
            throw error;
        }
    }

    async demoUsage(rpcUrl) {
        // Initialize with your RPC URL
        // const rpcUrl = 'http://localhost:8545'; // Update with your RPC
        if (!rpcUrl) {
            rpcUrl = 'http://localhost:8545';
        }
        const gofs = new GOFSContractInteractions(rpcUrl);

        try {
            // 1. System Dashboard
            console.log('=== 1. Homepage - System Dashboard ===');
            console.log('Total eVND Supply:', await gofs.compliantToken.totalSupply());
            console.log('Verified Entities:', '<NEED TO IMPLEMENT>');
            console.log('Daily Transactions:', '<NEED TO IMPLEMENT>');
            console.log('Compliance Rules:', (await gofs.getComplianceModules()).length);
            console.log('Exchange Rate:', await gofs.getExchangeRate());

            // 2. e-VND Token Details
            console.log('\n=== 2. e-VND Token Details ===');
            console.log('Token Name:', await gofs.compliantToken.name());
            console.log('Token Symbol:', await gofs.compliantToken.symbol());
            console.log('Token Address:', CONTRACT_ADDRESSES.COMPLIANT_TOKEN);
            console.log('Token Decimals:', await gofs.compliantToken.decimals());
            console.log('Token Total Supply:', await gofs.compliantToken.totalSupply());

            // 3. Entity Type Metadata
            console.log('\n=== 3. Entity Type Metadata ===');
            const entityTypeMetadata = await gofs.getAllEntityTypeMetadata();
            console.log('Entity Type Metadata:');
            Object.values(entityTypeMetadata).forEach(type => {
                console.log(`  ${type.name} (ID: ${type.id}): ${type.metadata}`);
            });

            // 4. Transaction Type Metadata
            console.log('\n=== 4. Transaction Type Metadata ===');
            const transactionTypeMetadata = await gofs.getAllTransactionTypeMetadata();
            console.log('Transaction Type Metadata:');
            Object.values(transactionTypeMetadata).forEach(type => {
                console.log(`  ${type.name} (ID: ${type.id}): ${type.metadata}`);
            });

            // 5. Transaction Type Policies
            console.log('\n=== 5. Transaction Type Policies ===');
            const transactionTypePolicies = await gofs.getTransactionTypePolicies();
            console.log('Transaction Type Policies:');
            Object.values(transactionTypePolicies).forEach(policy => {
                console.log(`\n  ${policy.name} (ID: ${policy.id}):`);
                console.log(`    Usable: ${policy.isUsable}`);
                console.log(`    Allowed From Entity Types: ${policy.allowedFromEntityTypes.map(t => t.name).join(', ')}`);
                console.log(`    Allowed To Entity Types: ${policy.allowedToEntityTypes.map(t => t.name).join(', ')}`);
            });

            // 6. Entity Type Transfer Policies
            console.log('\n=== 6. Entity Type Transfer Policies ===');
            const entityTypePolicies = await gofs.getEntityTypeTransferPolicies();
            console.log('Entity Type Transfer Policies:');
            Object.values(entityTypePolicies).forEach(fromPolicy => {
                console.log(`\n  ${fromPolicy.name} can transfer to:`);
                Object.values(fromPolicy.allowedTransfersTo).forEach(toPolicy => {
                    const status = toPolicy.allowed ? '✅' : '❌';
                    console.log(`    ${status} ${toPolicy.name}`);
                });
            });

            // 7. Individual Entity Details
            console.log('\n=== 7. Individual Entity Details ===');
            const example_user = CONTRACT_ADDRESSES.EXCHANGE_PORTAL;
            const entityDetails = await gofs.getEntityDetails(example_user);
            if (entityDetails) {
                console.log('Entity Details:');
                console.log(`  Address: ${entityDetails.address}`);
                console.log(`  Name: ${entityDetails.name}`);
                console.log(`  Entity Type: ${entityDetails.entityTypeName} (ID: ${entityDetails.entityType})`);
                console.log(`  Verified: ${entityDetails.isVerified ? 'Yes' : 'No'}`);
                console.log(`  Balance: ${entityDetails.balance} e-VND`);
                console.log(`  Verifier: ${entityDetails.verifier}`);
            } else {
                console.log('Entity not found');
            }

            // 8. Individual Metadata Examples
            console.log('\n=== 8. Individual Metadata Examples ===');
            const bankMetadata = await gofs.getEntityTypeMetadata(1);
            console.log(`Bank Entity Type: ${bankMetadata.metadata}`);

            const loanMetadata = await gofs.getTransactionTypeMetadata(1);
            console.log(`Loan Transaction Type: ${loanMetadata.metadata}`);

        } catch (error) {
            console.error('Demo error:', error);
        }
    }

}

// ============================================================================
// USAGE EXAMPLES
// ============================================================================



// Export for use in other modules
module.exports = {
    GOFSContractInteractions,
    CONTRACT_ADDRESSES,
};

// Run demo if this file is executed directly
if (require.main === module) {
    const rpcUrl = process.argv[2];
    const gofs = new GOFSContractInteractions(rpcUrl);
    gofs.demoUsage(rpcUrl).catch(console.error);
} 