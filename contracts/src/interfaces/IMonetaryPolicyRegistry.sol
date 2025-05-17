// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title IMonetaryPolicyRegistry
 * @dev Interface for the  Monetary Policy Registry that defines
 * transaction policies and quotas for the GOFS CBDC system.
 */
interface IMonetaryPolicyRegistry {
    type TxType is uint8;
    type EntityType is uint8;

    /**
     * @dev Transaction policy structure defining restrictions
     * @param allowedFromEntityTypesBitmap Bitmap of entity types allowed as senders
     * @param allowedToEntityTypesBitmap Bitmap of entity types allowed as receivers
     * @param allowedFromEntities Whitelist of specific addresses allowed as senders
     * @param allowedToEntities Whitelist of specific addresses allowed as receivers
     * @param blockedFromEntities Blacklist of specific addresses blocked as senders
     * @param blockedToEntities Blacklist of specific addresses blocked as receivers
     */
    struct TxPolicy {
        uint256 allowedFromEntityTypesBitmap; // Bitmap for sender entity types
        uint256 allowedToEntityTypesBitmap; // Bitmap for receiver entity types
        address[] allowedFromEntities; // Explicit whitelist for senders
        address[] allowedToEntities; // Explicit whitelist for receivers
        address[] blockedFromEntities; // Explicit blacklist for senders
        address[] blockedToEntities; // Explicit blacklist for receivers
    }

    /**
     * @dev Quota policy structure defining transaction limits
     * @param rate Maximum rate of transactions/transfers (e.g., tokens per hour)
     * @param limit Maximum cap for transactions/transfers
     * @param current Current usage amount against the quota
     */
    struct QuotaPolicy {
        uint256 rate; // Rate limit (tokens per time period)
        uint256 limit; // Absolute limit regardless of time
        uint256 current; // Current usage amount
    }

    /**
     * @dev Events emitted by the contract
     */
    event TxPolicyUpdated(TxType indexed txType);
    event QuotaPolicyUpdated(
        TxType indexed txType,
        uint256 rate,
        uint256 limit
    );
    event EntityTypeQuotaPolicyUpdated(
        TxType indexed entityType,
        uint256 rate,
        uint256 limit
    );
    event EntityQuotaPolicyUpdated(
        address indexed entity,
        uint256 rate,
        uint256 limit
    );

    /**
     * @dev Set transaction policy for a specific transaction type
     * @param txType The transaction type to set policy for
     * @param allowedFromEntityTypes Array of entity types allowed to send
     * @param allowedToEntityTypes Array of entity types allowed to receive
     * @param allowedFromEntities Array of specific addresses allowed to send
     * @param allowedToEntities Array of specific addresses allowed to receive
     * @param blockedFromEntities Array of specific addresses blocked from sending
     * @param blockedToEntities Array of specific addresses blocked from receiving
     */
    function setTxPolicy(
        TxType txType,
        EntityType[] memory allowedFromEntityTypes,
        EntityType[] memory allowedToEntityTypes,
        address[] memory allowedFromEntities,
        address[] memory allowedToEntities,
        address[] memory blockedFromEntities,
        address[] memory blockedToEntities
    ) external;

    /**
     * @dev Set quota policy for a specific transaction type
     * @param txType The transaction type to set quota for
     * @param quotaPolicy The quota policy to set
     */
    function setTxQuotaPolicy(
        TxType txType,
        QuotaPolicy calldata quotaPolicy
    ) external;

    /**
     * @dev Set quota policy for a specific entity type
     * @param entityType The entity type to set quota for
     * @param quotaPolicy The quota policy to set
     */
    function setEntityTypeQuotaPolicy(
        EntityType entityType,
        QuotaPolicy calldata quotaPolicy
    ) external;

    /**
     * @dev Set quota policy for a specific entity
     * @param entity The entity address to set quota for
     * @param quotaPolicy The quota policy to set
     */
    function setEntityQuotaPolicy(
        address entity,
        QuotaPolicy calldata quotaPolicy
    ) external;

    /**
     * @dev Get transaction policy for a specific transaction type
     * @param txType The transaction type to get policy for
     * @return policy The transaction policy
     */
    function getTxPolicy(TxType txType) external view returns (TxPolicy memory);

    /**
     * @dev Get quota policy for a specific transaction type
     * @param txType The transaction type to get quota for
     * @return policy The quota policy
     */
    function getTxQuotaPolicy(
        TxType txType
    ) external view returns (QuotaPolicy memory);

    /**
     * @dev Get quota policy for a specific entity type
     * @param entityType The entity type to get quota for
     * @return policy The quota policy
     */
    function getEntityTypeQuotaPolicy(
        EntityType entityType
    ) external view returns (QuotaPolicy memory);

    /**
     * @dev Get quota policy for a specific entity
     * @param entity The entity address to get quota for
     * @return policy The quota policy
     */
    function getEntityQuotaPolicy(
        address entity
    ) external view returns (QuotaPolicy memory);

    /**
     * @dev Check if a transaction is allowed based on policy
     * @param from The sender address
     * @param to The recipient address
     * @param txType The transaction type
     * @return isAllowed True if the transaction is allowed
     */
    function checkTxPolicy(
        address from,
        address to,
        TxType txType
    ) external view returns (bool);

    /**
     * @dev Check if a transaction exceeds quota limits
     * @param from The sender address
     * @param amount The transaction amount
     * @param txType The transaction type
     * @return isExceeding True if the transaction exceeds quota limits
     */
    function checkQuotaPolicy(
        address from,
        uint256 amount,
        TxType txType
    ) external view returns (bool);

    /**
     * @dev Update quota usage for an entity
     * @param txType The transaction type
     * @param from The entity address
     * @param amount The transaction amount
     * @return success True if update successful
     */
    function updateQuotaUsage(
        EntityType txType,
        address from,
        uint256 amount
    ) external returns (bool);
}
