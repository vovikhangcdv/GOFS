// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Hashes} from "@openzeppelin/contracts/utils/cryptography/Hashes.sol";

/// @title MerkleLibrary
/// @notice Reference: https://github.com/OpenZeppelin/merkle-tree
/// @notice Library for Merkle tree building and proof generation
/// @notice For testing purposes only. In real production, Merkle root is only calculated in off-chain.
/// @dev Provides functions for calculating the root of a Merkle tree and generating proofs
library MerkleLibrary {
    
    // @notice Calculates the root of a Merkle tree for a given array of infos
    /// @param infos The array of infos to calculate the root for. The length must be a power of 2.
    /// @return The root of the Merkle tree
    function calculateInfoRoot(string[] memory infos) internal pure returns (bytes32) {
        require(_isPowerOf2(infos.length), "MerkleLibrary: Length must be a power of 2");
        bytes32[] memory tree = buildTree(infos);
        return tree[0];
    }

    /// @notice Converts an array of strings to an array of hashed values
    /// @param infos The array of strings to convert to leaves
    /// @return The array of leaves
    function _toLeaves(string[] memory infos) internal pure returns (bytes32[] memory) {
        bytes32[] memory leaves = new bytes32[](infos.length);
        for (uint256 i = 0; i < infos.length; i++) {
            leaves[i] = keccak256(abi.encodePacked(infos[i]));
        }
        return leaves;
    }
    
    /// @notice Checks if a number is a power of 2
    /// @param x The number to check
    /// @return True if the number is a power of 2, false otherwise
    function _isPowerOf2(uint256 x) internal pure returns (bool) {
        return x != 0 && (x & (x - 1)) == 0;
    }

    /// @notice Builds a Merkle tree for a given array of leaves
    /// @param infos The array of data to build the tree for. The length must be a power of 2.
    /// @return The Merkle tree
    function buildTree(string[] memory infos) public pure returns (bytes32[] memory) {
        require(_isPowerOf2(infos.length), "MerkleLibrary: Length must be a power of 2");
        bytes32[] memory leaves = _toLeaves(infos);
        bytes32[] memory tree = new bytes32[](leaves.length * 2 - 1);   
        for (uint256 i = 0; i < leaves.length; i++) {
            tree[tree.length - 1 - i] = leaves[i];
        }

        for (uint256 i = 0; i < tree.length - leaves.length; i++) {
            uint256 idx = tree.length - 1 - leaves.length - i;
            tree[idx] = Hashes.commutativeKeccak256(tree[_getLeftChildIndex(idx)], tree[_getRightChildIndex(idx)]);
        }
        return tree;
    }

    /// @notice Generates a proof for a given index in a Merkle tree
    /// @param tree The Merkle tree to generate the proof for.
    /// @param index The index of the leaf to generate the proof for
    /// @return The proof for the given index
    function getProof(bytes32[] memory tree, uint256 index) internal pure returns (bytes32[] memory) {
        require(index < tree.length, "MerkleLibrary: Index out of bounds");
        require(_isPowerOf2(tree.length + 1));
        index = tree.length - index - 1;
        bytes32[] memory proof = new bytes32[](_getLevels(tree.length) - 1);
        uint256 cnt = 0;
        while (index > 0) {
            proof[cnt++] = tree[_getSiblingIndex(index)];
            index = _getParentIndex(index);
        }
        return proof;
    }

    function _getLevels(uint256 length) internal pure returns (uint256) {
        uint256 levels = 0;
        while (length > 1) {
            length = (length + 1) / 2;
            levels++;
        }
        return levels;
    }
    function _getParentIndex(uint256 index) internal pure returns (uint256) {
        require(index > 0, "MerkleLibrary: Index must be greater than 0");
        return (index - 1) / 2;
    }

    function _getSiblingIndex(uint256 index) internal pure returns (uint256) {
        require(index > 0, "MerkleLibrary: Index must be greater than 0");
        return index % 2 == 0 ? index - 1 : index + 1;
    }

    function _getLeftChildIndex(uint256 index) internal pure returns (uint256) {
        return index * 2 + 1;
    }

    function _getRightChildIndex(uint256 index) internal pure returns (uint256) {
        return index * 2 + 2;
    }
}