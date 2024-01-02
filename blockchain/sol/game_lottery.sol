// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "./common.sol";

contract TweetLotteryGame is ServiceFeeForWithdraw, TweetVotePlugInI {
    uint256 public __lotteryGameRoundTime = 48 hours;
    uint256 public __currentLotteryTicketID = 100000;
    uint256 public __bonusRateToWinner = 50;
    bool public __openToOuterPlayer = false;

    uint256 public __ticketPriceForOuter = 1e6 gwei;
    uint256 public __serviceFeeRateForTicketBuy = 5;

    uint256 public currentRoundNo = 0;
    mapping(address => uint256) public bonusBalance;

    struct GameInfoOneRound {
        bytes32 randomHash;
        uint256 discoverTime;
        address winner;
        bytes32 winTeam;
        uint256 winTicketID;
        uint256 bonus;
    }
    struct TweetTeam {
        mapping(address => uint256) memMap;
        address[] memList;
        uint256 voteNo;
    }

    struct BuyerInfo {
        address addr;
        bytes32 team;
    }

    mapping(uint256 => bytes32) public buyerInfoIdxForTickets;
    mapping(bytes32 => BuyerInfo) public buyerInfoRecords;

    mapping(uint256 => GameInfoOneRound) public gameInfoRecord;
    mapping(uint256 => uint256[]) public ticketsRecords;

    mapping(uint256 => mapping(bytes32 => TweetTeam)) private tweetTeamMap;
    mapping(uint256 => mapping(address => uint256[])) ticketsOfBuyer;

    event TweetBought(
        bytes32 thash,
        address owner,
        address buyer,
        uint256 val,
        uint256 no
    );

    event AdminOperated(uint256 newTimeInHours, string opName);
    event SkipToNewRound(bytes32 hash, uint256 round);
    event WinnerWithdrawBonus(address winner, uint256 bonus);
    event TicketSold(address buyer, uint256 no, uint256 serviceFee);
    event DiscoverWinner(
        address winner,
        bytes32 winnerTeam,
        uint256 ticketID,
        uint256 bonus,
        uint256 bonusToTeam,
        uint256 random,
        bytes32 nextRandomHash
    );
    event KolIpRightBout(address kolAddr, address buyer, uint256 keyNo);

    constructor(bytes32 hash) payable {
        GameInfoOneRound memory newRoundInfo = GameInfoOneRound({
            randomHash: hash,
            discoverTime: block.timestamp + __lotteryGameRoundTime,
            winner: address(0),
            winTeam: bytes32(0),
            winTicketID: 0,
            bonus: msg.value
        });
        gameInfoRecord[currentRoundNo] = newRoundInfo;
    }

    receive() external payable {
        gameInfoRecord[currentRoundNo].bonus += msg.value;
        emit AdminOperated(msg.value, "received_bonus_from_outer");
    }

    /********************************************************************************
     *                       admin operation
     *********************************************************************************/
    function adminOpenToOuterPlayer(bool isOpen) public isOwner {
        __openToOuterPlayer = isOpen;
        emit AdminOperated(isOpen ? 1 : 0, "game_open_to_outer_player");
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
        emit AdminOperated(newRate, "rate_of_service_fee_for_tciket_buy");
    }

    function adminChangeRoundTime(uint256 newTimeInMinutes) public isOwner {
        require(newTimeInMinutes > 0, "invalid time");

        __lotteryGameRoundTime = newTimeInMinutes * 1 minutes;

        emit AdminOperated(newTimeInMinutes, "game_round_time_in_minitues");
    }

    function adminChangeBonusRateToWinner(uint256 newRate) public isOwner {
        require(newRate >= 0 && newRate <= 100, "invalid rate");

        __bonusRateToWinner = newRate;

        emit AdminOperated(newRate, "rate_for_bonus_winner");
    }

    /********************************************************************************
     *                       lottery admin
     *********************************************************************************/

    function skipToNextRound(bytes32 hash) public onlyAdmin noReentrant {
        require(hash != bytes32(0), "Hash cannot be the zero value");

        GameInfoOneRound memory newRoundInfo = GameInfoOneRound({
            randomHash: hash,
            discoverTime: block.timestamp + __lotteryGameRoundTime,
            winner: address(0),
            winTeam: bytes32(0),
            winTicketID: 0,
            bonus: 0
        });

        newRoundInfo.bonus += gameInfoRecord[currentRoundNo].bonus;
        gameInfoRecord[currentRoundNo].bonus = 0;
        currentRoundNo += 1;

        gameInfoRecord[currentRoundNo] = newRoundInfo;

        emit SkipToNewRound(hash, currentRoundNo);
    }

    function dispatchBonusToTeam(uint256 val, BuyerInfo memory winner)
    private
    returns (bytes32 teamHash)
    {
        if (winner.team == bytes32(0)) {
            bonusBalance[winner.addr] += val;
            return bytes32(0);
        }
        TweetTeam storage team = tweetTeamMap[currentRoundNo][winner.team];
        uint256 totalVote = team.voteNo;
        if (totalVote <= 1) {
            bonusBalance[winner.addr] += val;
            return bytes32(0);
        }

        uint256 bonusPerVote = val / (totalVote - 1);

        for (uint256 idx = 0; idx < team.memList.length; idx++) {
            address teamMember = team.memList[idx];
            uint256 vote = team.memMap[teamMember];

            if (teamMember == address(0) || vote == 0) {
                continue;
            }
            if (teamMember == winner.addr) {
                vote -= 1;
            }
            bonusBalance[teamMember] += bonusPerVote * vote;
        }

        return winner.team;
    }

    function generateWiner(uint256 random, bytes32 currentHash)
    public
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
        require(gInfo.bonus > __minValCheck, "no bonus");
        require(
            block.timestamp >= (gInfo.discoverTime - 10 minutes),
            "not time"
        );

        uint256 ticketId = generateWiner(random, gInfo.randomHash);

        bytes32 buyerHash = buyerInfoIdxForTickets[ticketId];
        require(buyerHash != bytes32(0), "invalid winner hash");
        BuyerInfo memory winner = buyerInfoRecords[buyerHash];
        require(winner.addr != address(0), "invalid winner address");

        gInfo.winner = winner.addr;
        gInfo.winTicketID = ticketId;

        uint256 bonusToWinner = ((gInfo.bonus / 100) * __bonusRateToWinner);
        bonusBalance[winner.addr] += bonusToWinner;

        uint256 bonusToTeam = gInfo.bonus - bonusToWinner;
        gInfo.winTeam = dispatchBonusToTeam(bonusToTeam, winner);

        currentRoundNo += 1;

        gameInfoRecord[currentRoundNo] = GameInfoOneRound({
            randomHash: nextRoundRandomHash,
            discoverTime: block.timestamp + __lotteryGameRoundTime,
            winner: address(0),
            winTeam: bytes32(0),
            winTicketID: 0,
            bonus: 0
        });

        emit DiscoverWinner(
            winner.addr,
            winner.team,
            ticketId,
            bonusToWinner,
            bonusToTeam,
            random,
            nextRoundRandomHash
        );
    }

    /********************************************************************************
     *                       lottery operation
     *********************************************************************************/

    function generateTicket(
        uint256 no,
        address buyer,
        bytes32 tweetHash
    ) internal {
        bytes32 buyerHash = keccak256(abi.encodePacked(buyer, tweetHash));
        if (buyerInfoRecords[buyerHash].addr == address(0)) {
            buyerInfoRecords[buyerHash] = BuyerInfo(buyer, tweetHash);
        }

        for (uint256 idx = 1; idx <= no; idx++) {
            uint256 newTid = __currentLotteryTicketID + idx;

            ticketsRecords[currentRoundNo].push(newTid);

            buyerInfoIdxForTickets[newTid] = buyerHash;

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
        generateTicket(voteNo, buyer, tweetHash);

        TweetTeam storage team = tweetTeamMap[currentRoundNo][tweetHash];
        if (team.memMap[buyer] == 0) {
            team.memList.push(buyer);
        }
        team.memMap[buyer] += voteNo;
        team.voteNo += voteNo;

        emit TweetBought(tweetHash, tweetOwner, buyer, val, voteNo);
    }

    function buyTicketFromOuter(uint256 ticketNo) public payable noReentrant {
        require(ticketNo > 0, "invalid ticket number");
        require(__openToOuterPlayer, "not open now");
        uint256 balance = msg.value;
        require(balance == __ticketPriceForOuter, "ticket price change");

        uint256 serFee = (balance / 100) * __serviceFeeRateForTicketBuy;
        recordServiceFee(serFee);
        balance -= serFee;

        gameInfoRecord[currentRoundNo].bonus += balance;
        generateTicket(ticketNo, msg.sender, bytes32(0));

        emit TicketSold(msg.sender, ticketNo, serFee);
    }

    function withdrawByWinner(uint256 amount, bool all) public noReentrant {
        uint256 balance = bonusBalance[msg.sender];
        if (all) {
            amount = balance;
        }
        require(amount > __minValCheck, "too small amount");
        require(balance >= amount, "too much amount");
        require(balance <= address(this).balance, "insufficient founds");

        bonusBalance[msg.sender] -= amount;
        uint256 reminders = minusWithdrawFee(amount);

        payable(msg.sender).transfer(reminders);

        emit WinnerWithdrawBonus(msg.sender, reminders);
    }

    function checkPluginInterface() external pure returns (bool) {
        return true;
    }

    /********************************************************************************
     *                       basic query
     *********************************************************************************/
    function teamMembers(bytes32 tweet)
    public
    view
    returns (address[] memory members)
    {
        TweetTeam storage team = tweetTeamMap[currentRoundNo][tweet];
        if (team.memList.length == 0) {
            return new address[](0);
        }
        members = new address[](team.memList.length);
        for (uint256 idx; idx < team.memList.length; idx++) {
            members[idx] = team.memList[idx];
        }
        return members;
    }

    function teamMembersCountForGame(bytes32 tweet)
    public
    view
    returns (uint256, uint256)
    {
        TweetTeam storage team = tweetTeamMap[currentRoundNo][tweet];
        return (team.memList.length, team.voteNo);
    }

    function teamMemberVoteNo(bytes32 tweet, address memAddr)
    public
    view
    returns (uint256)
    {
        TweetTeam storage team = tweetTeamMap[currentRoundNo][tweet];
        return team.memMap[memAddr];
    }

    function historyTeamMembers(uint256 roundNo, bytes32 tweet)
    public
    view
    returns (address[] memory members)
    {
        TweetTeam storage team = tweetTeamMap[roundNo][tweet];
        if (team.memList.length == 0) {
            return new address[](0);
        }
        members = new address[](team.memList.length);
        for (uint256 idx; idx < team.memList.length; idx++) {
            members[idx] = team.memList[idx];
        }
        return members;
    }

    function historyTeamMembersCountForGame(uint256 roundNo, bytes32 tweet)
    public
    view
    returns (uint256, uint256)
    {
        TweetTeam storage team = tweetTeamMap[roundNo][tweet];
        return (team.memList.length, team.voteNo);
    }

    function historyTeamMemberVoteNo(
        uint256 roundNo,
        bytes32 tweet,
        address memAddr
    ) public view returns (uint256) {
        TweetTeam storage team = tweetTeamMap[roundNo][tweet];
        return team.memMap[memAddr];
    }

    function currentTickets() public view returns (uint256[] memory) {
        return ticketsRecords[currentRoundNo];
    }

    function historyTickets(uint256 round)
    public
    view
    returns (uint256[] memory)
    {
        return ticketsRecords[round];
    }

    function currentTicketNo() public view returns (uint256) {
        return ticketsRecords[currentRoundNo].length;
    }

    function currentBonus() public view returns (uint256) {
        return gameInfoRecord[currentRoundNo].bonus;
    }

    function tickInfos(uint256 tid) public view returns (BuyerInfo memory) {
        bytes32 buyerHash = buyerInfoIdxForTickets[tid];
        return buyerInfoRecords[buyerHash];
    }

    function tickList(address owner) public view returns (uint256[] memory) {
        return ticketsOfBuyer[currentRoundNo][owner];
    }
}
