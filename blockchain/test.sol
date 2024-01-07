// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

// 合约 A
contract A {
    B b;

    constructor(address payable _contractB) {
        b = B(_contractB);
    }

    function YY() public {
        b.XX(msg.sender);
    }
}

// 合约 B
contract B {
    address public caller;
    uint256 public currentBalance;
    string public smgData;
    event CurrentCaller(address caller, address orig);

    function setAA(address a) public {
        caller = a;
    }

    modifier isOwner() {
        require(msg.sender == caller, "Caller is not owner");
        _;
    }

    function XX(address originalCaller) public isOwner {
        emit CurrentCaller(msg.sender, originalCaller);
    }

    receive() external payable {
        currentBalance += msg.value;
    }

    fallback(bytes calldata _data) external payable returns (bytes memory) {
        smgData = string(_data);
        currentBalance += msg.value;
        return _data;
    }
}

// 合约 A
contract C {
    constructor() payable {}

    function payToLottery(address r, uint256 amount) public {
        payable(r).transfer(amount);
    }

    function payToLottery2(address r, uint256 amount) public {
        (bool success, ) = payable(r).call{value: amount}("test");
        require(success, "Transfer failed");
    }
}
