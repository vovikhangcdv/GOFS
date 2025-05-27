// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

type EntityType is uint8;
type RiskLevel is uint8;
type TxType is uint8;

struct EntityInfo {
    address entityAddress;
    EntityType entityType;
    RiskLevel riskLevel;
    address verifier;
}
