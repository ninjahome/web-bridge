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
)

type TweetVoteAction struct {
	CreateTime      int64 `json:"create_time"`
	VoteCount       int   `json:"vote_count"`
	VoteForTheTweet int   `json:"user_vote_count"`
}

type TweetVotePersonalRecord struct {
	CreateTime int64 `json:"create_time" firestore:"create_time"`
	VoteCount  int   `json:"vote_count" firestore:"vote_count"`
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
			CreateTime: vote.CreateTime,
			VoteCount:  vote.VoteCount,
		}
	} else {
		var err = voteSnapshot.DataTo(&votedObj)
		if err != nil {
			util.LogInst().Err(err).Int64("create_time", vote.CreateTime).Msg("parse tweet vote obj failed")
			return nil, nil, err
		}

		votedObj.VoteCount += vote.VoteCount
	}

	return recordDoc, &votedObj, nil
}

func (dm *DbManager) queryVoteStatus(voter bool, target string, tx *firestore.Transaction, vote *TweetVoteAction) (*firestore.DocumentRef, *NinjaUsrInfo, error) {
	njDoc := dm.fileCli.Collection(DBTableNJUser).Doc(target)

	voteSnapshot, err := tx.Get(njDoc)
	if err != nil {
		util.LogInst().Err(err).Str("web3-id", target).
			Int64("create_time", vote.CreateTime).
			Msg("failed to get nj user")
		return nil, nil, err
	}
	var nu NinjaUsrInfo
	err = voteSnapshot.DataTo(&nu)
	if err != nil {
		util.LogInst().Err(err).Str("web3-id", target).
			Int64("create_time", vote.CreateTime).Msg("parse nj user failed")
		return nil, nil, err
	}

	if voter {
		nu.VoteCount += vote.VoteCount
	} else {
		nu.BeVotedCount += vote.VoteCount
	}

	return njDoc, &nu, nil
}

type voteStatusForDb struct {
	recordDoc *firestore.DocumentRef
	voterDoc  *firestore.DocumentRef
	votedDoc  *firestore.DocumentRef
	recordObj *TweetVotePersonalRecord
	voterObj  *NinjaUsrInfo
	votedObj  *NinjaUsrInfo
}

func (dm *DbManager) queryStatus(createTime, voter, voted string, tx *firestore.Transaction, vote *TweetVoteAction) (*voteStatusForDb, error) {
	recordDoc, recordObj, err := dm.queryVoteRecord(createTime, voter, tx, vote)
	if err != nil {
		return nil, err
	}

	voterDoc, voterObj, err := dm.queryVoteStatus(true, voter, tx, vote)
	if err != nil {
		return nil, err
	}

	votedDoc, votedObj, err := dm.queryVoteStatus(false, voted, tx, vote)
	if err != nil {
		return nil, err
	}

	return &voteStatusForDb{
		recordDoc: recordDoc,
		voterDoc:  voterDoc,
		votedDoc:  votedDoc,
		recordObj: recordObj,
		voterObj:  voterObj,
		votedObj:  votedObj,
	}, nil
}

func (dm *DbManager) updateStatus(status *voteStatusForDb, tx *firestore.Transaction) error {
	var err = tx.Set(status.recordDoc, status.recordObj)
	if err != nil {
		util.LogInst().Err(err).Msg("update vote status doc err")
		return err
	}

	err = tx.Update(status.voterDoc, []firestore.Update{{Path: "vote_count", Value: status.voterObj.VoteCount}})
	if err != nil {
		util.LogInst().Err(err).Msg("update voter doc err")
		return err
	}

	err = tx.Update(status.votedDoc, []firestore.Update{{Path: "be_voted_count", Value: status.votedObj.BeVotedCount}})
	if err != nil {
		util.LogInst().Err(err).Msg("update voted  doc err")
		return err
	}

	return nil
}

func (dm *DbManager) UpdateTweetVoteStatic(vote *TweetVoteAction, voter string) error {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()

	err := dm.fileCli.RunTransaction(opCtx, func(ctx context.Context, tx *firestore.Transaction) error {
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
		vote.VoteForTheTweet = queryStatus.recordObj.VoteCount
		return nil
	})

	return err
}

func (dm *DbManager) QueryVotedTweetID(pageSize int, startID int64, voter string) ([]*TweetVotePersonalRecord, error) {

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
