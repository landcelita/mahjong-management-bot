package gamestatus

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/landcelita/mahjong-management-bot/domain/model/baKyokuHonba"
	"github.com/landcelita/mahjong-management-bot/domain/model/playerId"
	"github.com/landcelita/mahjong-management-bot/domain/model/tonpuOrHanchan"
)

type GameStatusId uuid.UUID

type Jicha uint

const (
	Toncha Jicha = iota + 1
	Nancha
	Shacha
	Pecha
)

type (
	GameStatus struct {
		gameStatusId  	GameStatusId
		baKyokuHonba	bakyokuhonba.BaKyokuHonba
		tonpuOrHanchan 	tonpuorhanchan.TonpuOrHanchan
		playerIds		map[Jicha]playerid.PlayerId
		isActive		bool
	}
)

func NewGameStatus(
	gameStatusId	GameStatusId,
	baKyokuHonba	bakyokuhonba.BaKyokuHonba,
	tonpuOrHanchan	tonpuorhanchan.TonpuOrHanchan,
	playerIds		map[Jicha]playerid.PlayerId,
	isActive		bool) (*GameStatus, error) {

	if len(playerIds) != 4 {
		return nil, fmt.Errorf("playerは四人である必要があります。")
	}

	if _, exist := playerIds[Toncha]; !exist { return nil, fmt.Errorf("Tonchaが指定されていません。") }
	if _, exist := playerIds[Nancha]; !exist { return nil, fmt.Errorf("Nanchaが指定されていません。") }
	if _, exist := playerIds[Shacha]; !exist { return nil, fmt.Errorf("Shachaが指定されていません。") }
	if _, exist := playerIds[Pecha]; !exist { return nil, fmt.Errorf("Pechaが指定されていません。") }

	if nan10, _ := bakyokuhonba.NewBaKyokuHonba(bakyokuhonba.Nan, 1, 0); 
	tonpuOrHanchan == tonpuorhanchan.Tonpu &&
	baKyokuHonba.IsLaterThanOrSameFor(*nan10) {
		return nil, fmt.Errorf("東風戦で南場に入ることはできません。")
	}
	
	gameStatus := &GameStatus {
		gameStatusId: gameStatusId,
		baKyokuHonba: baKyokuHonba,
		tonpuOrHanchan: tonpuOrHanchan,
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

	playerIds_copy := gameStatus.playerIds
	gameStatus.playerIds[Toncha] = playerIds_copy[Nancha]
	gameStatus.playerIds[Nancha] = playerIds_copy[Shacha]
	gameStatus.playerIds[Shacha] = playerIds_copy[Pecha]
	gameStatus.playerIds[Pecha] = playerIds_copy[Toncha]

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
