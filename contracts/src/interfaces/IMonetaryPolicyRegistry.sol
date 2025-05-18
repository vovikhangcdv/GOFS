// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title IMonetaryPolicyRegistry
 * @dev Interface for the Monetary Policy Registry that defines
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
     * @param minValue Minimum transaction value allowed
     * @param maxValue Maximum transaction value allowed
     * @param settlementPeriod Time in seconds until transaction settlement (0 = immediate)
     */
    struct TxPolicy {
        uint256 allowedFromEntityTypesBitmap; // Bitmap for sender entity types
        uint256 allowedToEntityTypesBitmap; // Bitmap for receiver entity types
        address[] allowedFromEntities; // Explicit whitelist for senders
        address[] allowedToEntities; // Explicit whitelist for receivers
        address[] blockedFromEntities; // Explicit blacklist for senders
        address[] blockedToEntities; // Explicit blacklist for receivers
        uint256 minValue; // Minimum transaction value allowed
        uint256 maxValue; // Maximum transaction value allowed
        uint256 settlementPeriod; // Time in seconds until settlement (0 = immediate)
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
        EntityType indexed entityType,
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
     * @param minValue Minimum transaction value allowed (0 for no minimum)
     * @param maxValue Maximum transaction value allowed (0 for no maximum)
     * @param settlementPeriod Time in seconds until settlement (0 for immediate)
     */
    function setTxPolicy(
        TxType txType,
        EntityType[] memory allowedFromEntityTypes,
        EntityType[] memory allowedToEntityTypes,
        address[] memory allowedFromEntities,
        address[] memory allowedToEntities,
        address[] memory blockedFromEntities,
        address[] memory blockedToEntities,
        uint256 minValue,
        uint256 maxValue,
        uint256 settlementPeriod
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
}
