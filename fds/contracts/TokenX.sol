// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract TokenX is ERC20, Ownable {
    mapping(address => bool) public blacklist;
    
    event AddressBlacklisted(address indexed account);
    event AddressRemovedFromBlacklist(address indexed account);
    
    constructor(
        string memory name,
        string memory symbol,
        uint256 initialSupply
    ) ERC20(name, symbol) Ownable(msg.sender) {
        _mint(msg.sender, initialSupply);
    }
    
    function addToBlacklist(address account) external onlyOwner {
        require(!blacklist[account], "Address already blacklisted");
        blacklist[account] = true;
        emit AddressBlacklisted(account);
    }
    
    function removeFromBlacklist(address account) external onlyOwner {
        require(blacklist[account], "Address not blacklisted");
        blacklist[account] = false;
        emit AddressRemovedFromBlacklist(account);
    }
    
    function _update(address from, address to, uint256 value) internal virtual override {
        require(!blacklist[from], "Sender is blacklisted");
        require(!blacklist[to], "Recipient is blacklisted");
        super._update(from, to, value);
    }
} 