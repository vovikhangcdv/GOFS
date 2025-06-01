// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Entity} from "./ITypes.sol";

interface IEntityRegisterable {
    /**
     * @dev Function to register an entity with the registry
     * @param registryAddress The address of the EntityRegistry contract
     * @param entity The entity data to register
     * @param verifierSignature The signature from the verifier
     */
    function registerWithRegistry(
        address registryAddress,
        Entity calldata entity,
        bytes calldata verifierSignature
    ) external;
}
