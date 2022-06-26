package gamestatus

import (
	"fmt"

	"github.com/google/uuid"
	bakyokuhonba "github.com/landcelita/mahjong-management-bot/domain/model/baKyokuHonba"
	"github.com/landcelita/mahjong-management-bot/domain/model/jicha"
	playerid "github.com/landcelita/mahjong-management-bot/domain/model/playerId"
	tonpuorhanchan "github.com/landcelita/mahjong-management-bot/domain/model/tonpuOrHanchan"
)

type GameStatusId uuid.UUID

type (
	GameStatus struct {
		gameStatusId   GameStatusId
		baKyokuHonba   bakyokuhonba.BaKyokuHonba
		tonpuOrHanchan tonpuorhanchan.TonpuOrHanchan
		playerIds      map[jicha.Jicha]playerid.PlayerId
		isActive       bool
	}
)

func newGameStatus(
	gameStatusId GameStatusId,
	baKyokuHonba bakyokuhonba.BaKyokuHonba,
	tonpuOrHanchan tonpuorhanchan.TonpuOrHanchan,
	playerIds map[jicha.Jicha]playerid.PlayerId,
	isActive bool,
) (*GameStatus, error) {

	if len(playerIds) != 4 {
		return nil, fmt.Errorf("playerは四人である必要があります。")
	}

	for _, jicha := range []jicha.Jicha{jicha.Toncha, jicha.Nancha, jicha.Shacha, jicha.Pecha} {
		if _, exist := playerIds[jicha]; !exist {
			return nil, fmt.Errorf(string(jicha) + "が指定されていません。")
		}
	}

	if nan10, _ := bakyokuhonba.NewBaKyokuHonba(bakyokuhonba.Nan, 1, 0); tonpuOrHanchan == tonpuorhanchan.Tonpu &&
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
		if last, err := bakyokuhonba.NewBaKyokuHonba(
			bakyokuhonba.Ton,
			4,
			0,
		); gameStatus.baKyokuHonba.EqualsBaKyoku(*last) && err == nil {
			return true
		} else {
			return false
		}
	} else if gameStatus.tonpuOrHanchan == tonpuorhanchan.Hanchan {
		if last, err := bakyokuhonba.NewBaKyokuHonba(
			bakyokuhonba.Nan,
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
	gameStatus.playerIds[jicha.Toncha] = playerIds_copy[jicha.Nancha]
	gameStatus.playerIds[jicha.Nancha] = playerIds_copy[jicha.Shacha]
	gameStatus.playerIds[jicha.Shacha] = playerIds_copy[jicha.Pecha]
	gameStatus.playerIds[jicha.Pecha] = playerIds_copy[jicha.Toncha]

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
