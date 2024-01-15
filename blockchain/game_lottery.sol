// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "./common.sol";

contract TweetLotteryGame is ServiceFeeForWithdraw, TweetVotePlugInI {
    uint256 public __lotteryGameRoundTime = 48 hours;
    uint256 public __currentLotteryTicketID = 100000;
    uint8 public __bonusRateToWinner = 50;
    bool public __openToOuterPlayer = true;

    uint256 public __ticketPriceForOuter = 1e6 gwei;
    uint8 public __serviceFeeRateForTicketBuy = 10;

    uint256 public totalBonus = 0 gwei;
    uint256 public currentRoundNo = 0;

    struct GameInfoOneRound {
        bytes32 randomHash;
        uint256 discoverTime;
        address winner;
        bytes32 winTeam;
        uint256 winTicketID;
        uint256 bonus;
        uint256 randomVal;
    }
    struct TweetTeam {
        mapping(address => uint256) memVotes;
        mapping(uint256 => address) memIndex;
        uint256 voteNo;
        uint256 memCount;
    }

    struct BuyerInfo {
        address addr;
        bytes32 team;
    }

    mapping(uint256 => bytes32) public buyerInfoIdxForTickets;
    mapping(bytes32 => BuyerInfo) public buyerInfoRecords;

    mapping(address => GameInfoOneRound[]) public winnerGameInfo;
    mapping(uint256 => GameInfoOneRound) public gameInfoRecord;
    mapping(uint256 => uint256[]) public ticketsRecords;

    mapping(uint256 => mapping(bytes32 => TweetTeam)) private tweetTeamMap;
    mapping(uint256 => mapping(address => uint256[])) public ticketsOfBuyer;
    mapping(uint256 => bytes32[]) public teamList;

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
        bytes32 winnerTeam,
        uint256 ticketID,
        uint256 bonus,
        uint256 bonusToTeam,
        uint256 random,
        bytes32 nextRandomHash
    );

    constructor(address[] memory admins, bytes32 hash) payable {
        require(hash != bytes32(0), "invalid random hash");
        GameInfoOneRound memory newRoundInfo = GameInfoOneRound({
            randomHash: hash,
            discoverTime: block.timestamp + __lotteryGameRoundTime,
            winner: address(0),
            winTeam: bytes32(0),
            winTicketID: 0,
            bonus: msg.value,
            randomVal: 0
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

    /********************************************************************************
     *                       lottery admin
     *********************************************************************************/
    /**
     * @dev Initiates the next round of the lottery game.
     * @param hash The initial random hash for the new round.
     * This function transitions the game to a new round. It carries over any unclaimed bonus from the current round to the next round.
     * The provided hash is used as the starting point for the random number generation in the new round.
     * Only callable by administrators of the contract to ensure controlled progression of game rounds.
     * Emits a SkipToNewRound event upon successfully initiating a new round.
     *
     * Requirements:
     * - The provided hash must not be the zero value.
     * - The function must be called by an administrator.
     */
    function skip(bytes32 hash) private {
        GameInfoOneRound memory newRoundInfo = GameInfoOneRound({
            randomHash: hash,
            discoverTime: block.timestamp + __lotteryGameRoundTime,
            winner: address(0),
            winTeam: bytes32(0),
            winTicketID: 0,
            bonus: 0,
            randomVal: 0
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

    /*
     * @dev Distributes the bonus to the team members of the winning ticket.
     * @param val The total bonus value to be distributed.
     * @param winner The buyer information of the winning ticket.
     * @return The team hash of the winning team.
     * This function calculates and distributes the bonus to each team member based on their votes.
     * If the winner has no team (team hash is zero), the full bonus is allocated to the winner's account.
     */
    function dispatchBonusToTeam(uint256 val, BuyerInfo memory winner)
    internal
    returns (bytes32 teamHash)
    {
        if (winner.team == bytes32(0)) {
            balance[winner.addr] += val;
            return bytes32(0);
        }
        TweetTeam storage team = tweetTeamMap[currentRoundNo][winner.team];
        uint256 totalVote = team.voteNo;
        if (totalVote <= 1) {
            balance[winner.addr] += val;
            return winner.team;
        }

        uint256 bonusPerVote = val / (totalVote - 1);

        for (uint256 i = 0; i < team.memCount; i++) {
            address teamMember = team.memIndex[i];
            uint256 memberVoteCount = team.memVotes[teamMember];

            if (teamMember == address(0) || memberVoteCount == 0) {
                continue;
            }
            if (teamMember == winner.addr) {
                memberVoteCount -= 1;
            }
            balance[teamMember] += bonusPerVote * memberVoteCount;
        }

        return winner.team;
    }

    /**
     * @dev Generates the winning ticket ID for the current lottery round based on a random number.
     * @param random The random number provided to determine the winning ticket.
     * @param currentHash A hash representing the current state of randomness.
     * @return The ID of the winning ticket.
     * This function calculates the winner of the current lottery round by using the provided random number.
     * It ensures the integrity of the random number by checking it against the currentHash.
     * The function computes a new hash from the given random number and other blockchain parameters, then uses this hash to select a winning ticket from all the tickets in the current round.
     *
     * Requirements:
     * - There must be at least one ticket in the current round.
     * - The provided hash must match the expected hash, ensuring the random number is valid.
     */
    function generateWiner(uint256 random, bytes32 currentHash)
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

    /**
     * @dev Discovers and announces the winner of the current lottery round based on a random number.
     * @param random The random number used to select the winning ticket.
     * @param nextRoundRandomHash The hash for the next round's random number.
     * This function determines the winning ticket based on the provided random number and assigns the winnings.
     * It also prepares the game for the next round using the provided next round hash.
     * Emits a DiscoverWinner event upon finding a winner.
     */
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

        uint256 ticketId = generateWiner(random, gInfo.randomHash);

        bytes32 buyerHash = buyerInfoIdxForTickets[ticketId];
        require(buyerHash != bytes32(0), "invalid winner hash");
        BuyerInfo memory winner = buyerInfoRecords[buyerHash];
        require(winner.addr != address(0), "invalid winner address");

        gInfo.randomVal = random;
        gInfo.winner = winner.addr;
        gInfo.winTicketID = ticketId;

        uint256 bonusToWinner = ((gInfo.bonus / 100) * __bonusRateToWinner);
        balance[winner.addr] += bonusToWinner;

        uint256 bonusToTeam = gInfo.bonus - bonusToWinner;
        gInfo.winTeam = dispatchBonusToTeam(bonusToTeam, winner);

        currentRoundNo += 1;
        totalBonus += gInfo.bonus;

        winnerGameInfo[winner.addr].push(gInfo);

        gameInfoRecord[currentRoundNo] = GameInfoOneRound({
            randomHash: nextRoundRandomHash,
            discoverTime: block.timestamp + __lotteryGameRoundTime,
            winner: address(0),
            winTeam: bytes32(0),
            winTicketID: 0,
            bonus: 0,
            randomVal: 0
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
    /**
     * @dev Generates lottery tickets for a buyer in the current lottery round.
     * @param no The number of tickets to be generated.
     * @param buyer The address of the buyer purchasing the tickets.
     * @param tweetHash The hash of the tweet associated with the ticket purchase.
     * This function is responsible for generating a specified number of lottery tickets for a buyer.
     * It records each ticket's information and associates it with the buyer's information.
     * For each ticket, a unique ID is generated and mapped to the buyer's address and the corresponding tweet hash.
     * The function increments the global ticket ID counter and updates mappings to track tickets of the buyer and tickets in the current round.
     *
     * Requirements:
     * - The buyer's address and tweet hash must be valid.
     * - The number of tickets requested must be greater than zero.
     */
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

    /**
     * @dev Processes the purchase of votes for a specific tweet, as part of the lottery game.
     * @param tweetHash The hash of the tweet being voted on.
     * @param tweetOwner The owner of the tweet.
     * @param buyer The address of the buyer who is purchasing votes.
     * @param voteNo The number of votes being purchased.
     * This function increases the game's bonus pool with the value sent to the contract and generates lottery tickets for the buyer.
     * It also updates the vote count in the relevant tweet team.
     * Emits a TweetBought event upon successful execution.
     */
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

        updateTweetTeam(tweetHash, buyer, voteNo);

        emit TweetBought(tweetHash, tweetOwner, buyer, val, voteNo);
    }

    function updateTweetTeam(
        bytes32 tweetHash,
        address buyer,
        uint256 voteNo
    ) internal {
        TweetTeam storage team = tweetTeamMap[currentRoundNo][tweetHash];
        if (team.memCount == 0) {
            teamList[currentRoundNo].push(tweetHash);
        }
        if (team.memVotes[buyer] == 0) {
            team.memIndex[team.memCount] = buyer;
            team.memCount += 1;
        }
        team.memVotes[buyer] += voteNo;
        team.voteNo += voteNo;
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

    /**
     * @dev Allows external users to buy lottery tickets when the game is open to the public.
     * @param ticketNo The number of tickets the external user wants to purchase.
     * This function enables external users (non-platform users) to participate in the lottery game by purchasing tickets.
     * It verifies if the game is open to external players and if the sent Ether matches the price of the requested number of tickets.
     * The function calculates the service fee, records it, and adds the remainder of the Ether sent to the game's bonus pool.
     * It then generates the requested number of tickets for the external user.
     * Emits a TicketSold event upon successful ticket purchase.
     *
     * Requirements:
     * - The number of tickets requested must be greater than zero.
     * - The game must be open to external players.
     * - The sent Ether must exactly match the total price of the requested tickets.
     */
    function buyTicketFromOuter(uint256 ticketNo)
    public
    payable
    noReentrant
    inRun
    {
        require(ticketNo > 0, "invalid ticket number");
        require(__openToOuterPlayer, "not open now");
        uint256 b = msg.value;
        require(b == __ticketPriceForOuter, "ticket price change");

        uint256 serFee = (b / 100) * __serviceFeeRateForTicketBuy;
        recordServiceFee(serFee);
        b -= serFee;

        gameInfoRecord[currentRoundNo].bonus += b;
        generateTicket(ticketNo, msg.sender, bytes32(0));

        emit TicketSold(msg.sender, ticketNo, serFee);
    }

    function checkPluginInterface() external pure returns (bool) {
        return true;
    }

    /********************************************************************************
     *                       basic query
     *********************************************************************************/

    function allTeamInfo(uint256 roundNo)
    public
    view
    returns (
        bytes32[] memory tweets,
        uint256[] memory memCounts,
        uint256[] memory voteCounts
    )
    {
        tweets = teamList[roundNo];
        memCounts = new uint256[](tweets.length);
        voteCounts = new uint256[](tweets.length);
        for (uint256 idx; idx < tweets.length; idx++) {
            TweetTeam storage team = tweetTeamMap[roundNo][tweets[idx]];
            memCounts[idx] = team.memCount;
            voteCounts[idx] = team.voteNo;
        }
        return (tweets, memCounts, voteCounts);
    }

    function tweetList(uint256 roundNo)
    public
    view
    returns (bytes32[] memory tweets)
    {
        return teamList[roundNo];
    }

    function teamMembers(uint256 roundNo, bytes32 tweet)
    public
    view
    returns (
        uint256 voteNo,
        uint256 memNo,
        uint256[] memory voteNos,
        address[] memory members
    )
    {
        TweetTeam storage team = tweetTeamMap[roundNo][tweet];

        members = new address[](team.memCount);
        voteNos = new uint256[](team.memCount);

        for (uint256 idx = 0; idx < team.memCount; idx++) {
            address voter = team.memIndex[idx];
            members[idx] = voter;
            voteNos[idx] = team.memVotes[voter];
        }
        return (team.voteNo, team.memCount, voteNos, members);
    }

    function voteNoOfTeamate(
        uint256 roundNo,
        bytes32 tweet,
        address memAddr
    )
    public
    view
    returns (
        uint256,
        uint256,
        uint256
    )
    {
        TweetTeam storage team = tweetTeamMap[roundNo][tweet];
        return (team.memCount, team.voteNo, team.memVotes[memAddr]);
    }

    function historyRoundInfo(uint256 from, uint256 to)
    public
    view
    returns (GameInfoOneRound[] memory infos)
    {
        infos = new GameInfoOneRound[](to - from + 1);
        for (uint256 i = from; i <= to; i++) {
            infos[i] = gameInfoRecord[i];
        }
        return infos;
    }

    function tickList(uint256 round, address owner)
    public
    view
    returns (uint256[] memory, bytes32[] memory)
    {
        uint256[] memory list = ticketsOfBuyer[round][owner];
        bytes32[] memory teamHash = new bytes32[](list.length);
        for (uint256 x = 0; x < list.length; x++) {
            uint256 tid = list[x];
            bytes32 buyerHash = buyerInfoIdxForTickets[tid];
            BuyerInfo memory bi = buyerInfoRecords[buyerHash];
            teamHash[x] = bi.team;
        }
        return (list, teamHash);
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

    function userWinnerData(address winner)
    public
    view
    returns (GameInfoOneRound[] memory winInfos)
    {
        return winnerGameInfo[winner];
    }
}
