// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

abstract contract Owner {
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

abstract contract ServiceFeeForWithdraw is Owner {
    uint256 private __serviceFeeReceived;
    uint256 private __withdrawFeeRate = 2;
    uint256 public constant __minValCheck = 1 gwei;
    mapping(address => bool) public __admins;

    modifier onlyAdmin() {
        require(__admins[msg.sender] == true, "only admins operation");
        _;
    }

    event AdminOperation(address admin, bool opType);
    event ServiceFeeChanged(uint256 newSerficeFeeRate);
    event UpgradeToNewRule(address newContract, uint256 balance);

    constructor() {
        __admins[msg.sender] = true;
    }

    function adminServiceFeeWithdraw() public isOwner noReentrant {
        require(__serviceFeeReceived > 0, "insufficient service fee");
        payable(this.getOwner()).transfer(__serviceFeeReceived);
        __serviceFeeReceived = 0;
    }

    function adminSetWithdrawFeeRate(uint256 newRate) public isOwner {
        require(newRate >= 0 && newRate <= 100, "rate invalid");
        __withdrawFeeRate = newRate;
        emit ServiceFeeChanged(newRate);
    }

    function recordServiceFee(uint256 fee) internal {
        __serviceFeeReceived += fee;
    }

    function minusWithdrawFee(uint256 val) internal returns (uint256) {
        if (__withdrawFeeRate == 0) {
            return val;
        }
        uint256 fee = (val / 100) * __withdrawFeeRate;
        __serviceFeeReceived += fee;
        return val - fee;
    }

    function serviceFeeReceived() public view returns (uint256) {
        return __serviceFeeReceived;
    }

    function withdrawFeeRate() public view returns (uint256) {
        return __withdrawFeeRate;
    }

    function adminUpgradeToNewRule(address payable recipient)
    public
    isOwner
    noReentrant
    isValidAddress(recipient)
    {
        if (__serviceFeeReceived > 0) {
            payable(this.getOwner()).transfer(__serviceFeeReceived);
            __serviceFeeReceived = 0;
        }

        uint256 balance = address(this).balance;
        if (balance > 0) {
            recipient.transfer(balance);
        }
        emit UpgradeToNewRule(recipient, balance);
    }

    function adminOperation(address admin, bool isDelete) public isOwner {
        if (isDelete) {
            delete __admins[admin];
        } else {
            __admins[admin] = true;
        }

        emit AdminOperation(admin, isDelete);
    }
}

interface TweetVotePlugInI {
    function tweetBought(
        bytes32 tweetHash,
        address owner,
        address buyer,
        uint256 voteNo
    ) external payable;

    function checkPluginInterface() external pure returns (bool);
}
