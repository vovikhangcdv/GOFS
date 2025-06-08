// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/TokenX.sol";

contract DeployTokenX is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        // Deploy TokenX with initial parameters
        TokenX tokenX = new TokenX(
            "Token X",
            "TKX",
            1000000 * 10**18  // 1 million tokens with 18 decimals
        );

        vm.stopBroadcast();
    }
} 