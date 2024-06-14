package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"github.com/ninjahome/web-bridge/util"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
	"strings"
)

type SysPoints struct {
	EthAddr        string  `json:"eth_addr" firestore:"eth_addr"`
	Points         float64 `json:"points"  firestore:"points"`
	BonusToWin     float64 `json:"bonus_to_win" firestore:"bonus_to_win"`
	CurTotalPoints float64 `json:"cur_total_points"`
}

type PointLogic func(sp *SysPoints, isNew bool)

func (dm *DbManager) ProcSystemPoints(ethAddr string, call PointLogic) {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()

	err := dm.fileCli.RunTransaction(opCtx, func(ctx context.Context, tx *firestore.Transaction) error {
		docRef := dm.fileCli.Collection(DBTableUserPoints).Doc(strings.ToLower(ethAddr))
		doc, err := tx.Get(docRef)

		if err != nil {
			if status.Code(err) == codes.NotFound {
				sp := &SysPoints{EthAddr: ethAddr}
				if call != nil {
					call(sp, true)
				}
				return tx.Set(docRef, sp)
			}
			return err
		}

		var sp SysPoints
		if err := doc.DataTo(&sp); err != nil {
			return err
		}
		if call != nil {
			call(&sp, false)
		}

		return tx.Update(docRef, []firestore.Update{
			{Path: "points", Value: sp.Points},
			{Path: "bonus_to_win", Value: sp.BonusToWin},
		})
	})

	if err != nil {
		util.LogInst().Err(err).Str("web3-id", ethAddr).Msg("process system points transaction failed")
	} else {
		util.LogInst().Debug().Str("eth-addr", ethAddr).Msg("process system points transaction success")
	}
}

func (dm *DbManager) QuerySystemPoints(web3ID string) (*SysPoints, error) {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()
	docRef := dm.fileCli.Collection(DBTableUserPoints).Doc(strings.ToLower(web3ID))
	doc, err := docRef.Get(opCtx)
	if err != nil {
		util.LogInst().Err(err).Str("web3-id", web3ID).Msg("query system points failed")
		return nil, err
	}

	var sp SysPoints
	err = doc.DataTo(&sp)
	if err != nil {
		util.LogInst().Err(err).Str("web3-id", web3ID).Msg("data to system points failed")
		return nil, err
	}
	return &sp, nil
}

func pointsWithReferrerBonus(sp *SysPoints, points float64) {
	if sp.BonusToWin > 0 {
		reward := math.Min(sp.BonusToWin, points*2)
		sp.BonusToWin = sp.BonusToWin - reward
		sp.Points += reward
	} else {
		sp.Points += points
	}
}

func (dm *DbManager) RewardForOneRound(totalPoints float64) {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut*10)
	defer cancel()

	err := dm.fileCli.RunTransaction(opCtx, func(ctx context.Context, tx *firestore.Transaction) error {
		iter := tx.Documents(dm.fileCli.Collection(DBTableUserPoints))
		defer iter.Stop()

		for {
			doc, err := iter.Next()
			if errors.Is(err, iterator.Done) {
				break
			}
			if err != nil {
				return err
			}

			var sp SysPoints
			if err := doc.DataTo(&sp); err != nil {
				return err
			}
			if sp.Points <= 0 {
				continue
			}

			pointsDelta := sp.Points / totalPoints * __dbConf.RewardPointsForOneRound
			newPoints := sp.Points + pointsDelta

			err = tx.Update(doc.Ref, []firestore.Update{
				{Path: "points", Value: newPoints},
			})
			if err != nil {
				return err
			}

			util.LogInst().Debug().Str("web3id", sp.EthAddr).
				Float64("newPoints", newPoints).
				Float64("delta", pointsDelta).
				Msg("update reward points success")
		}
		return nil
	})

	if err != nil {
		util.LogInst().Err(err).Msg("pints timer failed")
		return
	}

	util.LogInst().Info().Msg("pints timer transaction succeeded")
}

func (dm *DbManager) PointsAtSnapshot() float64 {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut*10)
	defer cancel()
	var totalPoints float64 = 0

	iter := dm.fileCli.Collection(DBTableUserPoints).Select("points").Documents(opCtx)
	for {
		doc, err := iter.Next()
		if err != nil {
			if errors.Is(err, iterator.Done) {
				break
			}
			util.LogInst().Err(err).Msg("timer:failed to calculate total points")
			return totalPoints
		}
		points := doc.Data()["points"].(float64)
		totalPoints += points
	}

	util.LogInst().Info().Float64("points", totalPoints).Msg("total points process success")
	return totalPoints
}
