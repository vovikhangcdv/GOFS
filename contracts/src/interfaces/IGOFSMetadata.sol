// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title IGOFSMetadata
 * @dev Interface for the GOFS Metadata contract which stores string representations
 * and additional metadata for entity and transaction types.
 * This allows for greater flexibility as new types can be added without
 * modifying the core contracts.
 */
interface IGOFSMetadata {
    type TxType is uint8;
    type EntityType is uint8;

    /**
     * @dev Events emitted by the contract
     */
    event EntityTypeMetadataSet(
        EntityType indexed typeId,
        string name,
        string description
    );
    event TxTypeMetadataSet(
        TxType indexed typeId,
        string name,
        string description
    );

    /**
     * @dev Set metadata for an entity type
     * @param typeId The ID of the entity type
     * @param name The name of the entity type
     * @param description A description of the entity type
     */
    function setEntityTypeMetadata(
        EntityType typeId,
        string memory name,
        string memory description
    ) external;

    /**
     * @dev Set metadata for a transaction type
     * @param typeId The ID of the transaction type
     * @param name The name of the transaction type
     * @param description A description of the transaction type
     */
    function setTxTypeMetadata(
        TxType typeId,
        string memory name,
        string memory description
    ) external;

    /**
     * @dev Get metadata for an entity type
     * @param typeId The ID of the entity type
     * @return name The name of the entity type
     * @return description A description of the entity type
     * @return exists Whether metadata exists for this type
     */
    function getEntityTypeMetadata(
        EntityType typeId
    )
        external
        view
        returns (string memory name, string memory description, bool exists);

    /**
     * @dev Get metadata for a transaction type
     * @param typeId The ID of the transaction type
     * @return name The name of the transaction type
     * @return description A description of the transaction type
     * @return exists Whether metadata exists for this type
     */
    function getTxTypeMetadata(
        TxType typeId
    )
        external
        view
        returns (string memory name, string memory description, bool exists);

    /**
     * @dev Get the name of an entity type
     * @param typeId The ID of the entity type
     * @return name The name of the entity type
     */
    function getEntityTypeName(
        EntityType typeId
    ) external view returns (string memory);

    /**
     * @dev Get the name of a transaction type
     * @param typeId The ID of the transaction type
     * @return name The name of the transaction type
     */
    function getTxTypeName(TxType typeId) external view returns (string memory);

    /**
     * @dev Check if entity type metadata exists
     * @param typeId The ID of the entity type
     * @return exists Whether metadata exists for this type
     */
    function entityTypeExists(EntityType typeId) external view returns (bool);

    /**
     * @dev Check if transaction type metadata exists
     * @param typeId The ID of the transaction type
     * @return exists Whether metadata exists for this type
     */
    function txTypeExists(TxType typeId) external view returns (bool);
}
