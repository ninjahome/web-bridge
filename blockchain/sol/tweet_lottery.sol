// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "./common.sol";
import "./tweet_exchange.sol";

contract TweetLotteryGame is Owner, PlugInI {
    uint256 lotteryGameRoundTime = (48 hours - 10 minutes);
    mapping(address => bool) admins;

    uint256 currentRound = 1;
    uint256 public gameBalance = 0;
    uint256 public currentLotteryTicketID = 0;
    bytes32 public randomHash;
    uint256 public nextLotteryDrawTime = 0;

    address[] public winners;
    mapping(address => uint256) bonusForWinner;
    mapping(uint256 => address) ticketsOwner;
    uint256[] public allTickets;
    mapping(address => uint256[]) ownersTickets;

    event TweetBought(
        bytes32 thash,
        address owner,
        address buyer,
        uint256 val,
        uint256 no
    );

    modifier onlyAdmin() {
        require(admins[msg.sender] == true, "only admin's operation");
        _;
    }

    event AdminOperation(address admin, bool opType);
    event RoundTimeChanged(uint256 newTime);
    event StartLottery(bytes32 hash, uint256 time);
    event DiscoverWinner(address admin, uint256 bonus);
    event WinnerWithdrawBonus(address winner, uint256 bonus);

    constructor() payable {
        admins[msg.sender] = true;
    }

    receive() external payable {}

    function adminOperation(address admin, bool isDelete) public isOwner {
        if (isDelete) {
            delete admins[admin];
        } else {
            admins[admin] = true;
        }
        emit AdminOperation(admin, isDelete);
    }

    function changeRoundTime(uint256 newTimeInHours) public isOwner {
        require(newTimeInHours > 10 minutes, "invalid time in hour");
        lotteryGameRoundTime = newTimeInHours * 1 days - 10 minutes;
        emit RoundTimeChanged(lotteryGameRoundTime);
    }

    function startLotteryGame(bytes32 hash) public onlyAdmin {
        require(hash != bytes32(0), "Hash cannot be the zero value");
        randomHash = hash;
        for (uint256 idx = 0; idx < allTickets.length; idx++) {
            uint256 tid = allTickets[idx];
            address owner = ticketsOwner[tid];
            delete ownersTickets[owner];
            ticketsOwner[tid] = address(0);
        }
        delete allTickets;

        emit StartLottery(hash, block.timestamp);
    }

    function discoveryWinner(uint256 random) public onlyAdmin noReentrant {
        require(allTickets.length > 0, "no tickets");
        require(gameBalance > 0, "no bonus");
        require(block.timestamp >= nextLotteryDrawTime, "not time");

        bytes32 hash = keccak256(abi.encodePacked(random));
        require(hash == randomHash, "invalid random data");
        currentRound += 1;

        uint256 idx = generateRandomNumber(random) % allTickets.length;
        uint256 ticketId = allTickets[idx];
        address winner = ticketsOwner[ticketId];
        if (winner == address(0)) {
            return;
        }

        winners.push(winner);
        nextLotteryDrawTime += lotteryGameRoundTime;
        uint256 bonus = gameBalance;
        gameBalance = 0;
        bonusForWinner[winner] += bonus;
        emit DiscoverWinner(winner, bonus);
    }

    function withDrawBonus() public noReentrant {
        require(bonusForWinner[msg.sender] > 1 gwei, "no bonus for you");
        uint256 bonus = bonusForWinner[msg.sender];
        bonusForWinner[msg.sender] = 0;
        payable(msg.sender).transfer(bonus);
        emit WinnerWithdrawBonus(msg.sender, bonus);
    }

    function tweetBought(
        bytes32 tweetHash,
        address owner,
        address buyer,
        uint256 leftVal,
        uint256 voteNo
    ) public onlyAdmin {
        gameBalance += leftVal;

        for (uint256 idx = 1; idx <= voteNo; idx++) {
            uint256 newTid = currentLotteryTicketID + idx;
            allTickets.push(newTid);
            ticketsOwner[newTid] = buyer;
            ownersTickets[buyer].push(newTid);
        }

        currentLotteryTicketID += voteNo;

        emit TweetBought(tweetHash, owner, buyer, leftVal, voteNo);
    }

    function checkPluginInterface() external pure returns (bool) {
        return true;
    }

    function generateRandomNumber(uint256 random)
    public
    view
    returns (uint256)
    {
        uint256 blockHashNumber = uint256(blockhash(block.number - 1));
        uint256 timestamp = block.timestamp;
        uint256 difficulty = block.difficulty;
        return
            uint256(
            keccak256(
                abi.encodePacked(
                    blockHashNumber,
                    timestamp,
                    difficulty,
                    random
                )
            )
        );
    }
}
