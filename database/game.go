package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"github.com/ninjahome/web-bridge/blockchain/ethapi"
	"github.com/ninjahome/web-bridge/util"
	"google.golang.org/api/iterator"
	"strings"
)

const (
	MaxWinHistoryQuery      = 40
	DBTableGameResult       = "lottery_game_round_info"
	DBTableWinTeamForMember = "win_team_info_for_member"
)

func (dm *DbManager) QueryGameWinner(web3id string) ([]*ethapi.GamInfoOnChain, error) {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()

	query := dm.fileCli.Collection(DBTableGameResult).
		Where("winner", "==", strings.ToLower(web3id)).
		OrderBy("discover_time", firestore.Desc).
		Limit(MaxWinHistoryQuery)

	iter := query.Documents(opCtx)
	defer iter.Stop()

	var gameInfos = make([]*ethapi.GamInfoOnChain, 0)
	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			return gameInfos, nil
		}
		if err != nil {
			util.LogInst().Err(err).Msgf("game info to iterate: %v", err)
			return nil, err
		}

		var gi ethapi.GamInfoOnChain
		err = doc.DataTo(&gi)
		if err != nil {
			util.LogInst().Err(err).Msg("parse to game info failed")
			continue
		}
		gameInfos = append(gameInfos, &gi)
	}
}
