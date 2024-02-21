// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "./common.sol";

contract TweetLotteryGame is ServiceFeeForWithdraw, TweetVotePlugInI {
    uint256 public __lotteryGameRoundTime = 48 hours;
    uint256 public __currentLotteryTicketID = 100000;
    uint8 public __bonusRateToWinner = 50;
    bool public __openToOuterPlayer = true;

    uint256 public __ticketPriceForOuter = 5 * 1e6 gwei;
    uint8 public __serviceFeeRateForTicketBuy = 10;

    uint256 public totalBonus = 0 gwei;
    uint256 public currentRoundNo = 0;
    uint256 public bonusForPoints = 0;

    struct GameInfoOneRound {
        bytes32 randomHash;
        uint256 discoverTime;
        address winner;
        uint256 winTicketID;
        uint256 bonus;
        uint256 bonusForWinner;
        uint256 randomVal;
    }

    mapping(uint256 => address) public buyerInfoIdxForTickets;
    mapping(uint256 => GameInfoOneRound) public gameInfoRecord;
    mapping(uint256 => uint256[]) public ticketsRecords;
    mapping(uint256 => mapping(address => uint256[])) public ticketsOfBuyer;

    event TweetBought(
        bytes32 thash,
        address owner,
        address buyer,
        uint256 val,
        uint256 no
    );

    event AdminOperated(uint256 newTimeInMinutes, string opName);
    event SkipToNewRound(bytes32 hash, uint256 round);
    event TicketSold(address buyer, uint256 no, uint256 serviceFee);
    event DiscoverWinner(
        address winner,
        uint256 ticketID,
        uint256 bonus,
        uint256 bonusToPoints,
        uint256 random,
        bytes32 nextRandomHash
    );

    constructor(address[] memory admins, bytes32 hash) payable {
        require(hash != bytes32(0), "invalid random hash");
        GameInfoOneRound memory newRoundInfo = GameInfoOneRound({
            randomHash: hash,
            discoverTime: block.timestamp + __lotteryGameRoundTime,
            winner: address(0),
            winTicketID: 0,
            bonus: msg.value,
            randomVal: 0,
            bonusForWinner: 0
        });

        gameInfoRecord[currentRoundNo] = newRoundInfo;
        for (uint256 idx; idx < admins.length; idx++) {
            __admins[admins[idx]] = true;
        }
        __admins[msg.sender] = true;
    }

    receive() external payable {
        gameInfoRecord[currentRoundNo].bonus += msg.value;
        emit AdminOperated(msg.value, "received_bonus_from_outer");
    }

    /********************************************************************************
     *                       admin operation
     *********************************************************************************/
    function adminOpenToOuterPlayer(bool isOpen) public isOwner {
        require(__openToOuterPlayer != isOpen, "no need change");
        __openToOuterPlayer = isOpen;
        emit AdminOperated(isOpen ? 1 : 0, "game_open_to_outer_player");
    }

    function adminSetTicketPriceForOuter(uint256 priceInFinney) public isOwner {
        require(priceInFinney > __minValCheck, "invalid ticket price");
        require(
            __ticketPriceForOuter != priceInFinney * 1e6 gwei,
            "no need change"
        );
        __ticketPriceForOuter = priceInFinney * 1e6 gwei;
    }

    function adminSetServiceFeeRateForTicketBuy(uint8 newRate) public isOwner {
        require(newRate >= 0 && newRate <= 100, "invalid rate param");
        require(__serviceFeeRateForTicketBuy != newRate, "no need change");

        __serviceFeeRateForTicketBuy = newRate;
        emit AdminOperated(newRate, "rate_of_service_fee_for_tciket_buy");
    }

    function adminChangeRoundTime(uint256 newTimeInMinutes) public isOwner {
        require(newTimeInMinutes > 0, "invalid time");
        require(
            __lotteryGameRoundTime != newTimeInMinutes * 1 minutes,
            "no need change"
        );

        __lotteryGameRoundTime = newTimeInMinutes * 1 minutes;

        emit AdminOperated(newTimeInMinutes, "game_round_time_in_minitues");
    }

    function adminChangeBonusRateToWinner(uint8 newRate) public isOwner {
        require(newRate >= 0 && newRate <= 100, "invalid rate");
        require(__bonusRateToWinner != newRate, "no need change");

        __bonusRateToWinner = newRate;

        emit AdminOperated(newRate, "rate_for_bonus_winner");
    }

    function finishPoint(address payable tokenContract) public isOwner {
        require(address(0) != tokenContract, "invalid address");
        require(bonusForPoints >= __minValCheck, "too small point eth");
        tokenContract.transfer(bonusForPoints);
    }

    /********************************************************************************
     *                       lottery admin
     *********************************************************************************/

    function skip(bytes32 hash) private {
        GameInfoOneRound memory newRoundInfo = GameInfoOneRound({
            randomHash: hash,
            discoverTime: block.timestamp + __lotteryGameRoundTime,
            winner: address(0),
            winTicketID: 0,
            bonus: 0,
            randomVal: 0,
            bonusForWinner: 0
        });

        newRoundInfo.bonus += gameInfoRecord[currentRoundNo].bonus;
        gameInfoRecord[currentRoundNo].bonus = 0;
        currentRoundNo += 1;

        gameInfoRecord[currentRoundNo] = newRoundInfo;
    }

    function skipToNextRound(bytes32 hash) public onlyAdmin noReentrant {
        require(hash != bytes32(0), "Hash cannot be the zero value");

        skip(hash);

        emit SkipToNewRound(hash, currentRoundNo);
    }

    function generateWinner(uint256 random, bytes32 currentHash)
    internal
    view
    returns (uint256)
    {
        uint256[] memory allTickets = ticketsRecords[currentRoundNo];
        require(allTickets.length > 0, "no tickets");

        bytes32 hash = keccak256(abi.encodePacked(random));
        require(hash == currentHash, "invalid random data");

        bytes32 newRandom = keccak256(
            abi.encodePacked(
                uint256(blockhash(block.number - 1)),
                block.timestamp,
                block.prevrandao,
                random
            )
        );

        uint256 idx = uint256(newRandom) % allTickets.length;
        return allTickets[idx];
    }

    function discoverWinner(uint256 random, bytes32 nextRoundRandomHash)
    public
    onlyAdmin
    noReentrant
    {
        GameInfoOneRound storage gInfo = gameInfoRecord[currentRoundNo];

        require(gInfo.randomHash != bytes32(0), "random not set");
        require(gInfo.winner == address(0), "can't have winner before game");
        require(
            block.timestamp >= (gInfo.discoverTime - 10 minutes),
            "not time"
        );

        if (gInfo.bonus <= __minValCheck) {
            skip(nextRoundRandomHash);
            return;
        }

        uint256 ticketId = generateWinner(random, gInfo.randomHash);

        address winnerAddr = buyerInfoIdxForTickets[ticketId];
        require(winnerAddr != address(0), "invalid winner address");

        gInfo.randomVal = random;
        gInfo.winner = winnerAddr;
        gInfo.winTicketID = ticketId;

        uint256 bonusToWinner = ((gInfo.bonus / 100) * __bonusRateToWinner);
        balance[winnerAddr] += bonusToWinner;
        gInfo.bonusForWinner = bonusToWinner;
        bonusForPoints += gInfo.bonus - bonusToWinner;

        currentRoundNo += 1;
        totalBonus += gInfo.bonus;

        gameInfoRecord[currentRoundNo] = GameInfoOneRound({
            randomHash: nextRoundRandomHash,
            discoverTime: block.timestamp + __lotteryGameRoundTime,
            winner: address(0),
            winTicketID: 0,
            bonus: 0,
            randomVal: 0,
            bonusForWinner: 0
        });

        emit DiscoverWinner(
            winnerAddr,
            ticketId,
            bonusToWinner,
            bonusForPoints,
            random,
            nextRoundRandomHash
        );
    }

    /********************************************************************************
     *                       lottery operation
     *********************************************************************************/

    function generateTicket(uint256 no, address buyer) internal {
        for (uint256 idx = 1; idx <= no; idx++) {
            uint256 newTid = __currentLotteryTicketID + idx;

            ticketsRecords[currentRoundNo].push(newTid);

            buyerInfoIdxForTickets[newTid] = buyer;

            ticketsOfBuyer[currentRoundNo][buyer].push(newTid);
        }

        __currentLotteryTicketID += no;
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

        gameInfoRecord[currentRoundNo].bonus += val;

        generateTicket(voteNo, buyer);

        emit TweetBought(tweetHash, tweetOwner, buyer, val, voteNo);
    }

    function withdraw(uint256 amount, bool all) public noReentrant inRun {
        uint256 _curBalance = balance[msg.sender];
        if (all) {
            amount = _curBalance;
        }
        require(amount > __minValCheck, "too small amount");
        require(_curBalance >= amount, "more than balance");
        require(_curBalance <= address(this).balance, "insufficient founds");

        balance[msg.sender] -= amount;

        uint256 reminders = minusWithdrawFee(amount);

        payable(msg.sender).transfer(reminders);

        emit WithdrawService(msg.sender, reminders);
    }

    function buyTicketFromOuter(uint256 ticketNo)
    public
    payable
    noReentrant
    inRun
    {
        require(ticketNo > 0, "invalid ticket number");
        require(__openToOuterPlayer, "not open now");
        uint256 b = msg.value;
        require(b == __ticketPriceForOuter * ticketNo, "ticket price change");

        uint256 serFee = (b / 100) * __serviceFeeRateForTicketBuy;
        recordServiceFee(serFee);
        b -= serFee;

        gameInfoRecord[currentRoundNo].bonus += b;
        generateTicket(ticketNo, msg.sender);

        emit TicketSold(msg.sender, ticketNo, serFee);
    }

    function checkPluginInterface() external pure returns (bool) {
        return true;
    }

    /********************************************************************************
     *                       basic query
     *********************************************************************************/

    function historyRoundInfo(uint256 from, uint256 to)
    public
    view
    returns (GameInfoOneRound[] memory infos)
    {
        require(to >= from, "invalid param");
        uint256 size = to - from + 1;
        infos = new GameInfoOneRound[](size);
        for (uint256 i = 0; i < size; i++) {
            infos[i] = gameInfoRecord[i + from];
        }
        return infos;
    }

    function tickList(uint256 round, address owner)
    public
    view
    returns (uint256[] memory)
    {
        return ticketsOfBuyer[round][owner];
    }

    function systemSettings()
    public
    view
    returns (
        uint256,
        uint256,
        uint256,
        uint256,
        bool
    )
    {
        return (
            currentRoundNo,
            totalBonus,
            ticketsRecords[currentRoundNo].length,
            __ticketPriceForOuter,
            __openToOuterPlayer
        );
    }
}
