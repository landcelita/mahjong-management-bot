package gamestatus

import (
	"fmt"

	"github.com/google/uuid"
	bkh "github.com/landcelita/mahjong-management-bot/domain/model/baKyokuHonba"
	pid "github.com/landcelita/mahjong-management-bot/domain/model/playerId"
	"github.com/landcelita/mahjong-management-bot/domain/model/tonpuOrHanchan"
	jc "github.com/landcelita/mahjong-management-bot/domain/model/jicha"
)

type GameStatusId uuid.UUID

type (
	GameStatus struct {
		gameStatusId  	GameStatusId
		baKyokuHonba	bkh.BaKyokuHonba
		tonpuOrHanchan 	tonpuorhanchan.TonpuOrHanchan
		playerIds		map[jc.Jicha]pid.PlayerId
		isActive		bool
	}
)

func NewGameStatus(
	gameStatusId	GameStatusId,
	baKyokuHonba	bkh.BaKyokuHonba,
	tonpuOrHanchan	tonpuorhanchan.TonpuOrHanchan,
	playerIds		map[jc.Jicha]pid.PlayerId,
	isActive		bool,
) (*GameStatus, error) {

	if len(playerIds) != 4 {
		return nil, fmt.Errorf("playerは四人である必要があります。")
	}

	for _, jicha := range []jc.Jicha{jc.Toncha, jc.Nancha, jc.Shacha, jc.Pecha} {
		if _, exist := playerIds[jicha]; !exist {
			return nil, fmt.Errorf(string(jicha) + "が指定されていません。")
		}
	}

	if nan10, _ := bkh.NewBaKyokuHonba(bkh.Nan, 1, 0);
	tonpuOrHanchan == tonpuorhanchan.Tonpu &&
	baKyokuHonba.IsLaterThanOrSameFor(*nan10) {
		return nil, fmt.Errorf("東風戦で南場に入ることはできません。")
	}

	gameStatus := &GameStatus{
		gameStatusId:   gameStatusId,
		baKyokuHonba:   baKyokuHonba,
		tonpuOrHanchan: tonpuOrHanchan,
		playerIds:      playerIds,
		isActive:       isActive,
	}

	return gameStatus, nil
}

func (gameStatus GameStatus) IsOlast() bool {
	if gameStatus.tonpuOrHanchan == tonpuorhanchan.Tonpu {
		if last, err := bkh.NewBaKyokuHonba(
			bkh.Ton,
			4,
			0,
		); gameStatus.baKyokuHonba.EqualsBaKyoku(*last) && err == nil {
			return true
		} else {
			return false
		}
	} else if gameStatus.tonpuOrHanchan == tonpuorhanchan.Hanchan {
		if last, err := bkh.NewBaKyokuHonba(
			bkh.Nan,
			4,
			0,
		); gameStatus.baKyokuHonba.EqualsBaKyoku(*last) && err == nil {
			return true
		} else {
			return false
		}
	}

	return false
}

func (gameStatus *GameStatus) AdvanceGameBaKyoku() error {
	if gameStatus.IsOlast() {
		return fmt.Errorf("オーラス時にBaKyokuを進めることはできません。")
	}

	playerIds_copy := gameStatus.playerIds
	gameStatus.playerIds[jc.Toncha] = playerIds_copy[jc.Nancha]
	gameStatus.playerIds[jc.Nancha] = playerIds_copy[jc.Shacha]
	gameStatus.playerIds[jc.Shacha] = playerIds_copy[jc.Pecha]
	gameStatus.playerIds[jc.Pecha] = playerIds_copy[jc.Toncha]

	nextBaKyokuHonba, err := gameStatus.baKyokuHonba.IncrementBaKyoku()

	if err != nil {
		return err
	}

	gameStatus.baKyokuHonba = *nextBaKyokuHonba

	return nil
}

func (gameStatus *GameStatus) AdvanceGameHonba() error {
	nextBaKyokuHonba, err := gameStatus.baKyokuHonba.IncrementHonba()

	if err != nil {
		return err
	}

	gameStatus.baKyokuHonba = *nextBaKyokuHonba

	return nil
}

func (gameStatus *GameStatus) ID() GameStatusId {
	return gameStatus.gameStatusId
}
