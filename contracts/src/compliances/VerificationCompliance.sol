// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ICompliance} from "../interfaces/ICompliance.sol";
import {EntityRegistry} from "../EntityRegistry.sol";
import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";

/**
 * @title VerificationCompliance
 * @dev Compliance module that checks if both sender and receiver are verified entities
 * in the EntityRegistry before allowing transfers.
 */
contract VerificationCompliance is ICompliance {
    // The immutable reference to the Entity Registry
    EntityRegistry public immutable entityRegistry;

    error InvalidEntityRegistryAddress();

    /**
     * @dev Constructor to set the EntityRegistry address
     * @param _entityRegistry The address of the EntityRegistry contract
     */
    constructor(address _entityRegistry) {
        if (_entityRegistry == address(0))
            revert InvalidEntityRegistryAddress();
        entityRegistry = EntityRegistry(_entityRegistry);
    }

    /**
     * @dev Implements IERC165 interface detection
     */
    function supportsInterface(
        bytes4 interfaceId
    ) public view virtual override(IERC165) returns (bool) {
        return interfaceId == type(ICompliance).interfaceId;
    }

    /**
     * @dev Check if both addresses are verified entities
     * @param _from The sender's address
     * @param _to The recipient's address
     * @param _amount The amount to transfer (unused in this module)
     * @return bool Whether the transfer is allowed
     */
    function canTransfer(
        address _from,
        address _to,
        uint256 _amount
    ) public view override returns (bool) {
        (bool success, ) = canTransferWithFailureReason(_from, _to, _amount);
        return success;
    }

    /**
     * @dev Check if both addresses are verified entities and provide failure reason
     * @param _from The sender's address
     * @param _to The recipient's address
     * @param _amount The amount to transfer (unused in this module)
     * @return (bool, string) Whether the transfer is allowed and failure reason if not
     */
    function canTransferWithFailureReason(
        address _from,
        address _to,
        uint256 _amount
    ) public view override returns (bool, string memory) {
        // For minting, only check recipient
        if (_from == address(0)) {
            if (!entityRegistry.isVerifiedEntity(_to)) {
                return (
                    false,
                    "Recipient for minting is not a verified entity"
                );
            }
            return (true, "");
        }

        // For burning, only check sender
        if (_to == address(0)) {
            if (!entityRegistry.isVerifiedEntity(_from)) {
                return (false, "Sender for burning is not a verified entity");
            }
            return (true, "");
        }

        // Check sender for regular transfers
        if (!entityRegistry.isVerifiedEntity(_from)) {
            return (false, "Sender is not a verified entity");
        }

        // Check recipient for regular transfers
        if (!entityRegistry.isVerifiedEntity(_to)) {
            return (false, "Recipient is not a verified entity");
        }

        return (true, "");
    }
}
