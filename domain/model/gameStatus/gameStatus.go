package gamestatus

import (
	"fmt"
	"github.com/google/uuid"
	"mahjong/domain/model/baKyokuHonba"
	"mahjong/domain/model/playerId"
	"mahjong/domain/model/tonpuOrHanchan"
)

type (
	GameStatus struct {
		gameStatusId  	uuid.UUID
		baKyokuHonba	bakyokuhonba.BaKyokuHonba
		tonpuOrHanchan 	tonpuorhanchan.TonpuOrHanchan
		scoreId			uuid.UUID
		playerIds		[4]playerid.PlayerId
		isActive		bool
	}
)

// factoryによって生成するので、privateにする
func newGameStatus(
	gameStatusId	uuid.UUID,
	baKyokuHonba	bakyokuhonba.BaKyokuHonba,
	tonpuOrHanchan	tonpuorhanchan.TonpuOrHanchan,
	scoreId			uuid.UUID,
	playerIds		[4]playerid.PlayerId,
	isActive		bool) (*GameStatus, error) {
	
	gameStatus := &GameStatus {
		gameStatusId: gameStatusId,
		baKyokuHonba: baKyokuHonba,
		tonpuOrHanchan: tonpuOrHanchan,
		scoreId: scoreId,
		playerIds: playerIds,
		isActive: isActive,
	}

	return gameStatus, nil
}

func (gameStatus GameStatus) IsOlast() (bool) {
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

func (gameStatus *GameStatus) AdvanceGameBaKyoku() (error) {
	if gameStatus.IsOlast() {
		return fmt.Errorf("オーラス時にBaKyokuを進めることはできません。")
	}

	nextBaKyokuHonba, err := gameStatus.baKyokuHonba.IncrementBaKyoku()
	
	if err != nil { return err }

	gameStatus.baKyokuHonba = *nextBaKyokuHonba

	return nil
}

func (gameStatus *GameStatus) AdvanceGameHonba() (error) {
	nextBaKyokuHonba, err := gameStatus.baKyokuHonba.IncrementHonba()

	if err != nil { return err }

	gameStatus.baKyokuHonba = *nextBaKyokuHonba

	return nil
}
