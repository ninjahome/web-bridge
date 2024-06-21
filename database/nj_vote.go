package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
	"github.com/ninjahome/web-bridge/util"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type TweetVoteAction struct {
	CreateTime      int64  `json:"create_time"`
	VoteCount       int    `json:"vote_count"`
	VoteForTheTweet int    `json:"user_vote_count"`
	Slogan          string `json:"slogan"`
	TxHash          string `json:"tx_hash"`
}

type TweetVotePersonalRecord struct {
	CreateTime    int64 `json:"create_time" firestore:"create_time"`
	FirstVoteTime int64 `json:"first_vote_time" firestore:"first_vote_time"`
	VoteCount     int   `json:"vote_count" firestore:"vote_count"`
}

func (dm *DbManager) queryVoteRecord(createTime, voter string, tx *firestore.Transaction, vote *TweetVoteAction) (*firestore.DocumentRef, *TweetVotePersonalRecord, error) {

	recordDoc := dm.fileCli.Collection(DBTableTweetsVoted).Doc(voter).Collection(DBTableTweetsSubStatus).Doc(createTime)

	voteSnapshot, voteErr := tx.Get(recordDoc)
	var votedObj TweetVotePersonalRecord
	if voteErr != nil {
		if status.Code(voteErr) != codes.NotFound {
			util.LogInst().Err(voteErr).Int64("create_time", vote.CreateTime).Msg("Failed to get tweet vote record document")
			return nil, nil, voteErr
		}
		votedObj = TweetVotePersonalRecord{
			FirstVoteTime: time.Now().UnixMilli(),
			CreateTime:    vote.CreateTime,
			VoteCount:     vote.VoteCount,
		}
	} else {
		var err = voteSnapshot.DataTo(&votedObj)
		if err != nil {
			util.LogInst().Err(err).Int64("create_time", vote.CreateTime).Msg("parse tweet vote obj failed")
			return nil, nil, err
		}

		votedObj.VoteCount += vote.VoteCount
		if votedObj.FirstVoteTime == 0 {
			votedObj.FirstVoteTime = time.Now().UnixMilli()
		}
	}

	return recordDoc, &votedObj, nil
}

func (dm *DbManager) queryVoteStatus(voter, sameOwner bool, target string, tx *firestore.Transaction, vote *TweetVoteAction) (*firestore.DocumentRef, *NinjaUsrInfo, error) {
	njDoc := dm.fileCli.Collection(DBTableNJUser).Doc(target)

	njUsrObj, err := tx.Get(njDoc)
	if err != nil {
		util.LogInst().Err(err).Str("web3-id", target).
			Int64("create_time", vote.CreateTime).
			Msg("failed to get nj user")
		return nil, nil, err
	}

	var nu NinjaUsrInfo
	err = njUsrObj.DataTo(&nu)
	if err != nil {
		util.LogInst().Err(err).Str("web3-id", target).
			Int64("create_time", vote.CreateTime).Msg("parse nj user failed")
		return nil, nil, err
	}
	multiPly := 1
	if nu.IsElder {
		multiPly = 2
	}

	var points float64
	if sameOwner {
		points += __dbConf.PointForVote * float64(vote.VoteCount)
		points += __dbConf.PointForBeVote * float64(vote.VoteCount*multiPly)
		nu.VoteCount += vote.VoteCount
		nu.BeVotedCount += vote.VoteCount
	} else {
		if voter {
			nu.VoteCount += vote.VoteCount
			points += __dbConf.PointForVote * float64(vote.VoteCount)
		} else {
			nu.BeVotedCount += vote.VoteCount
			points += __dbConf.PointForBeVote * float64(vote.VoteCount*multiPly)
		}
	}

	go dm.ProcSystemPoints(target, func(sp *SysPoints, _ bool) {
		pointsWithReferrerBonus(sp, points)
	})

	return njDoc, &nu, nil
}

type voteStatusForDb struct {
	statusDoc *firestore.DocumentRef
	voterDoc  *firestore.DocumentRef
	votedDoc  *firestore.DocumentRef
	statusObj *TweetVotePersonalRecord
	voterObj  *NinjaUsrInfo
	votedObj  *NinjaUsrInfo
}

func (dm *DbManager) queryStatus(createTime, voter, voted string, tx *firestore.Transaction, vote *TweetVoteAction) (*voteStatusForDb, error) {
	statusDoc, statusObj, err := dm.queryVoteRecord(createTime, voter, tx, vote)
	if err != nil {
		return nil, err
	}

	sameOwner := voter == voted
	voterDoc, voterObj, err := dm.queryVoteStatus(true, sameOwner, voter, tx, vote)
	if err != nil {
		return nil, err
	}

	if sameOwner {
		return &voteStatusForDb{
			statusDoc: statusDoc,
			voterDoc:  voterDoc,
			votedDoc:  voterDoc,
			statusObj: statusObj,
			voterObj:  voterObj,
			votedObj:  voterObj,
		}, nil
	}

	votedDoc, votedObj, err := dm.queryVoteStatus(false, false, voted, tx, vote)
	if err != nil {
		return nil, err
	}

	return &voteStatusForDb{
		statusDoc: statusDoc,
		voterDoc:  voterDoc,
		votedDoc:  votedDoc,
		statusObj: statusObj,
		voterObj:  voterObj,
		votedObj:  votedObj,
	}, nil
}

func (dm *DbManager) updateStatus(status *voteStatusForDb, tx *firestore.Transaction) error {
	var err = tx.Set(status.statusDoc, status.statusObj)
	if err != nil {
		util.LogInst().Err(err).Msg("update vote status doc err")
		return err
	}

	voterUpdates := []firestore.Update{
		{Path: "vote_count", Value: status.voterObj.VoteCount},
	}

	err = tx.Update(status.voterDoc, voterUpdates)
	if err != nil {
		util.LogInst().Err(err).Msg("update voter doc err")
		return err
	}

	votedUpdates := []firestore.Update{
		{Path: "be_voted_count", Value: status.votedObj.BeVotedCount},
	}

	err = tx.Update(status.votedDoc, votedUpdates)
	if err != nil {
		util.LogInst().Err(err).Msg("update voted  doc err")
		return err
	}

	return nil
}

func (dm *DbManager) UpdatePointsForSingleBets(vote *TweetVoteAction, voter string) {
	dm.ProcSystemPoints(voter, func(sp *SysPoints, _ bool) {
		points := __dbConf.PointForVote * float64(vote.VoteCount)
		pointsWithReferrerBonus(sp, points)
	})
}

func (dm *DbManager) UpdateTweetVoteStatic(vote *TweetVoteAction, voter string) error {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut*3)
	defer cancel()
	stat := &TweetPaymentStatus{
		CreateTime: vote.CreateTime,
		TxHash:     vote.TxHash,
		Status:     TxStSuccess,
	}
	err := dm.checkTransactionStatus(opCtx, stat)
	if err != nil {
		return err
	}

	err = dm.fileCli.RunTransaction(opCtx, func(ctx context.Context, tx *firestore.Transaction) error {
		createTime := fmt.Sprintf("%d", vote.CreateTime)
		tweetDoc := dm.fileCli.Collection(DBTableTweetsPosted).Doc(createTime)

		tweetSnapshot, err := tx.Get(tweetDoc)
		if err != nil {
			util.LogInst().Err(err).Int64("createAt", vote.CreateTime).Msg("Failed to get tweet document to update vote")
			return err
		}

		var existingTweet NinjaTweet
		err = tweetSnapshot.DataTo(&existingTweet)
		if err != nil {
			util.LogInst().Err(err).Int64("createAt", vote.CreateTime).Msg("Failed to decode tweet document data")
			return err
		}

		queryStatus, err := dm.queryStatus(createTime, voter, existingTweet.Web3ID, tx, vote)
		if err != nil {
			util.LogInst().Err(err).Msg("query status for vote failed")
			return err
		}

		var newFieldValue = existingTweet.VoteCount + vote.VoteCount
		err = tx.Update(tweetDoc, []firestore.Update{{Path: "vote_count", Value: newFieldValue}})
		if err != nil {
			util.LogInst().Err(err).Msg("update vote count on tweet failed")
			return err
		}

		err = dm.updateStatus(queryStatus, tx)
		if err != nil {
			util.LogInst().Err(err).Msg("update status for vote failed")
			return err
		}

		vote.VoteCount = newFieldValue
		vote.VoteForTheTweet = queryStatus.statusObj.VoteCount
		return nil
	})

	return err
}

func (dm *DbManager) QueryVotedTweetIDByMe(pageSize int, startID int64, voter string) ([]*TweetVotePersonalRecord, error) {

	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()

	voteRef := dm.fileCli.Collection(DBTableTweetsVoted).Doc(voter).Collection(DBTableTweetsSubStatus)
	var query = voteRef.Limit(pageSize)

	if startID == 0 {
		query = query.OrderBy("create_time", firestore.Desc)
	} else {
		query = query.Where("create_time", "<", startID).OrderBy("create_time", firestore.Desc)
	}

	iter := query.Documents(opCtx)
	defer iter.Stop()
	var result = make([]*TweetVotePersonalRecord, 0)
	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			return result, nil

		}
		if err != nil {
			util.LogInst().Err(err).Msg("query vote status failed")
			return nil, err
		}

		var voteStatus TweetVotePersonalRecord
		err = doc.DataTo(&voteStatus)
		if err != nil {
			util.LogInst().Err(err).Msg("data to obj failed")
			continue
		}
		result = append(result, &voteStatus)
	}
}

func (dm *DbManager) QueryMostVotedTweets(pageSize int, startID int64) ([]*NinjaTweet, error) {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()

	voteRef := dm.fileCli.Collection(DBTableTweetsPosted)
	var query = voteRef.Limit(pageSize).Where("payment_status", "==", TxStSuccess)
	if startID == 0 {
		query = query.OrderBy("vote_count", firestore.Desc)
	} else {
		query = query.Where("vote_count", "<", startID).OrderBy("vote_count", firestore.Desc)
	}

	var iter = query.Documents(opCtx)
	defer iter.Stop()

	var tweets = make([]*NinjaTweet, 0)
	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			return tweets, nil
		}
		if err != nil {
			util.LogInst().Err(err).Msgf("Failed to iterate: %v", err)
			return nil, err
		}

		var tweet NinjaTweet
		err = doc.DataTo(&tweet)
		if err != nil {
			util.LogInst().Err(err).Msgf("Failed to convert document to nj tweet: %v", err)
			return nil, err
		}
		tweets = append(tweets, &tweet)
	}
}
