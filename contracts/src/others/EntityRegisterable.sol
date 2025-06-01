// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IEntityRegisterable} from "../interfaces/IEntityRegisterable.sol";
import {IEntityRegistry} from "../interfaces/IEntityRegistry.sol";
import {Entity} from "../interfaces/ITypes.sol";

abstract contract EntityRegisterable is IEntityRegisterable {
    function _authorizeRegister(address forwarder) internal virtual;

    function registerWithRegistry(
        address registryAddress,
        Entity calldata entity,
        bytes calldata verifierSignature
    ) external override {
        _authorizeRegister(msg.sender);
        IEntityRegistry(registryAddress).register(entity, verifierSignature);
    }
}
