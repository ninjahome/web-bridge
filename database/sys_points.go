package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/ninjahome/web-bridge/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
	"strings"
)

type SysPoints struct {
	EthAddr    string  `json:"eth_addr" firestore:"eth_addr"`
	Points     float32 `json:"points"  firestore:"points"`
	BonusToWin float32 `json:"bonus_to_win" firestore:"bonus_to_win"`
}
type PointLogic func(sp *SysPoints)

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
					call(sp)
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
			call(&sp)
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

func pointsWithReferrerBonus(sp *SysPoints, points float32) {
	if sp.BonusToWin > 0 {
		reward := float32(math.Min(float64(sp.BonusToWin), float64(points*2)))
		sp.BonusToWin = sp.BonusToWin - reward
		sp.Points += reward
	} else {
		sp.Points += points
	}
}
