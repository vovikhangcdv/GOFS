// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

type EntityType is uint8;
type TxType is uint8;

function compareEntityType(EntityType a, EntityType b) pure returns (bool) {
    return EntityType.unwrap(a) == EntityType.unwrap(b);
}

function compareTxType(TxType a, TxType b) pure returns (bool) {
    return TxType.unwrap(a) == TxType.unwrap(b);
}

using {compareEntityType as ==} for EntityType global;
using {compareTxType as ==} for TxType global;

struct Entity {
    address entityAddress;
    EntityType entityType;
    bytes entityData;
    address verifier;
}
