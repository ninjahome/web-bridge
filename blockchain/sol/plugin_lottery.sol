// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "./common.sol";
import "./tweet_exchange.sol";

contract TweetLotteryGame is ServiceFeeForWithdraw, PlugInI {
    address[] public __winnersInHistory;

    uint256 public __lotteryGameRoundTime = 48 hours;
    uint256 private __currentLotteryTicketID = 0;

    bool public __openToOuterPlayer = false;
    uint256 public __ticketPriceForOuter = 1e6 gwei;
    uint256 public __serviceFeeRateForTicketBuy = 5;

    bytes32 public randomHashForCurrentRound;
    uint256 public nextLotteryDiscoverTime = 0;
    uint256 public currentRoundNo = 1;
    uint256 public bonusOfCurrentRound = 0;

    mapping(address => uint256) public bonusBalance;
    mapping(uint256 => address) public playerAddrOfTicketOwner;
    uint256[] public allTicketsOfThisRound;
    mapping(address => uint256[]) public ticketListOfPlayer;

    event TweetBought(
        bytes32 thash,
        address owner,
        address buyer,
        uint256 val,
        uint256 no
    );

    event RoundTimeChanged(uint256 newTimeInHours);
    event StartLottery(bytes32 hash, uint256 round, uint256 time);
    event WinnerWithdrawBonus(
        address winner,
        uint256 bonus,
        uint256 serviceFee
    );
    event TicketSold(address buyer, uint256 no, uint256 serviceFee);
    event DiscoverWinner(
        address winner,
        uint256 bonus,
        uint256 random,
        bytes32 randomHash,
        uint256 block_number,
        uint256 block_timestamp,
        uint256 block_difficulty,
        uint256 ticketsNo,
        uint256 winnerTicket
    );

    constructor() payable {
        __admins[msg.sender] = true;
        nextLotteryDiscoverTime = block.timestamp;
    }

    receive() external payable {}

    function adminOpenToOuterPlayer(bool isOpen) public isOwner {
        __openToOuterPlayer = isOpen;
    }

    function adminSetTicketPriceForOuter(uint256 priceInFinney) public isOwner {
        require(priceInFinney > __minValCheck, "invalid ticket price");
        __ticketPriceForOuter = priceInFinney * 1e6 gwei;
    }

    function adminSetServiceFeeRateForTicketBuy(uint256 newRate)
    public
    isOwner
    {
        require(newRate >= 0 && newRate <= 100, "invalid rate param");
        __serviceFeeRateForTicketBuy = newRate;
    }

    function adminChangeRoundTime(uint256 newTimeInHours) public isOwner {
        require(newTimeInHours > 10 minutes, "invalid time in hour");

        __lotteryGameRoundTime = newTimeInHours * 1 hours;

        emit RoundTimeChanged(__lotteryGameRoundTime);
    }

    function startNewGameRound(bytes32 hash) public onlyAdmin {
        require(hash != bytes32(0), "Hash cannot be the zero value");

        randomHashForCurrentRound = hash;
        currentRoundNo += 1;

        for (uint256 idx = 0; idx < allTicketsOfThisRound.length; idx++) {
            uint256 tid = allTicketsOfThisRound[idx];

            address player = playerAddrOfTicketOwner[tid];
            delete ticketListOfPlayer[player];

            delete playerAddrOfTicketOwner[tid];
        }

        delete allTicketsOfThisRound;

        emit StartLottery(hash, currentRoundNo, block.timestamp);
    }

    function discoveryWinner(uint256 random) public onlyAdmin noReentrant {
        require(allTicketsOfThisRound.length > 0, "no tickets");
        require(bonusOfCurrentRound > __minValCheck, "no bonus");
        require(
            block.timestamp >= (nextLotteryDiscoverTime - 10 minutes),
            "not time"
        );

        bytes32 hash = keccak256(abi.encodePacked(random));
        require(hash == randomHashForCurrentRound, "invalid random data");

        uint256 idx = generateRandomNumber(random) %
                        allTicketsOfThisRound.length;
        uint256 ticketId = allTicketsOfThisRound[idx];
        address winner = playerAddrOfTicketOwner[ticketId];

        if (winner == address(0)) {
            return;
        }

        uint256 bonus = bonusOfCurrentRound;
        bonusOfCurrentRound = 0;
        bonusBalance[winner] += bonus;

        __winnersInHistory.push(winner);
        nextLotteryDiscoverTime = block.timestamp + __lotteryGameRoundTime;

        emit DiscoverWinner(
            winner,
            bonus,
            random,
            randomHashForCurrentRound,
            block.number,
            block.timestamp,
            block.prevrandao,
            allTicketsOfThisRound.length,
            ticketId
        );
    }

    function buyTicketFromOuter(uint256 ticketNo) public payable noReentrant {
        require(ticketNo > 0, "invalid ticket number");
        require(__openToOuterPlayer, "not open now");
        uint256 balance = msg.value;
        require(balance == __ticketPriceForOuter, "insufficient funds");

        uint256 serFee = (balance / 100) * __serviceFeeRateForTicketBuy;
        serviceFeeInc(serFee);
        balance -= serFee;

        bonusOfCurrentRound += balance;
        generateTicket(ticketNo, msg.sender);

        emit TicketSold(msg.sender, ticketNo, serFee);
    }

    function withdrawByWinner() public payable noReentrant {
        uint256 balance = bonusBalance[msg.sender];

        require(balance > __minValCheck, "no bonus for you");
        require(balance <= address(this).balance, "insufficient founds");

        uint256 serFee = (balance / 100) * serviceFeeRate();
        serviceFeeInc(serFee);
        balance -= serFee;

        bonusBalance[msg.sender] = 0;
        payable(msg.sender).transfer(balance);

        emit WinnerWithdrawBonus(msg.sender, balance, serFee);
    }

    function tweetBought(
        bytes32 tweetHash,
        address tweetOwner,
        address buyer,
        uint256 voteNo
    )
    public
    payable
    onlyAdmin
    isValidAddress(buyer)
    isValidAddress(tweetOwner)
    noReentrant
    {
        uint256 val = msg.value;
        require(val > __minValCheck, "invalid msg value");
        require(voteNo >= 1, "invalid vote no");
        require(tweetHash != bytes32(0));
        bonusOfCurrentRound += val;
        generateTicket(voteNo, buyer);

        emit TweetBought(tweetHash, tweetOwner, buyer, val, voteNo);
    }

    function checkPluginInterface() external pure returns (bool) {
        return true;
    }

    function KolIPRightsBought(
        address kolAddr,
        address buyer,
        uint256 keyNo
    ) external payable {
        return;
    }

    function generateRandomNumber(uint256 random)
    public
    view
    returns (uint256)
    {
        uint256 blockHashNumber = uint256(blockhash(block.number - 1));
        uint256 timestamp = block.timestamp;
        uint256 difficulty = block.prevrandao;
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

    function generateTicket(uint256 no, address buyer) internal {
        for (uint256 idx = 1; idx <= no; idx++) {
            uint256 newTid = __currentLotteryTicketID + idx;

            allTicketsOfThisRound.push(newTid);

            playerAddrOfTicketOwner[newTid] = buyer;

            ticketListOfPlayer[buyer].push(newTid);
        }
        __currentLotteryTicketID += no;
    }
}
