package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/ninjahome/web-bridge/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type SysPoints struct {
	EthAddr    string `json:"eth_addr" firestore:"eth_addr"`
	Points     int    `json:"points"  firestore:"points"`
	BonusToWin int    `json:"bonus_to_win" firestore:"bonus_to_win"`
}

func (dm *DbManager) ProcSystemPoints(ethAddr string, points, bonus int) {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()

	err := dm.fileCli.RunTransaction(opCtx, func(ctx context.Context, tx *firestore.Transaction) error {
		docRef := dm.fileCli.Collection(DBTableUserPoints).Doc(strings.ToLower(ethAddr))
		doc, err := tx.Get(docRef)

		if err != nil {
			if status.Code(err) == codes.NotFound {
				sp := &SysPoints{EthAddr: ethAddr, Points: points, BonusToWin: bonus}
				return tx.Set(docRef, sp)
			}
			return err
		}

		var sp SysPoints
		if err := doc.DataTo(&sp); err != nil {
			return err
		}
		newPoints := sp.Points + points
		newBonus := sp.BonusToWin - bonus

		return tx.Update(docRef, []firestore.Update{
			{Path: "points", Value: newPoints},
			{Path: "bonus_to_win", Value: newBonus},
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
