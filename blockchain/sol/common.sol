// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract Owner {
    address private owner;

    event OwnerSet(address indexed oldOwner, address indexed newOwner);

    modifier isOwner() {
        require(msg.sender == owner, "Caller is not owner");
        _;
    }

    constructor() {
        owner = msg.sender;
        emit OwnerSet(address(0), owner);
    }

    function changeOwner(address newOwner) public isOwner {
        emit OwnerSet(owner, newOwner);
        owner = newOwner;
    }

    function getOwner() external view returns (address) {
        return owner;
    }

    bool private locked;

    modifier noReentrant() {
        require(!locked, "No re-entrancy");
        locked = true;
        _;
        locked = false;
    }

    modifier isValidAddress(address addr) {
        require(addr != address(0), "invalid address");
        _;
    }
}

interface PlugInI {
    function tweetBought(
        bytes32 tweetHash,
        address owner,
        address buyer,
        uint256 leftVal,
        uint256 voteNo
    ) external;

    function checkPluginInterface() external pure returns (bool);
}
