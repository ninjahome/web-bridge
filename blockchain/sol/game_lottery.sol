// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "./common.sol";

contract LotteryGame is Owner {
    ExchangeI public exchange;
    bytes32 public nextRandomHash;

    constructor(address ex) {
        exchange = ExchangeI(ex);
    }

    function setNextRandomHash(bytes32 newHash) public isOwner {
        nextRandomHash = newHash;
    }
}
