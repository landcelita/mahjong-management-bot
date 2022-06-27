package gsinmem

import (
	"sync"

	gs "github.com/landcelita/mahjong-management-bot/domain/model/gameStatus"
)

type GameStatusRepoInMem struct {
	Data  map[gs.GameStatusId]*gs.GameStatus
	mutex sync.RWMutex
}

func NewGameStatusRepoInMem() gs.GameStatusIRepo {
	return &GameStatusRepoInMem{
		Data: make(map[gs.GameStatusId]*gs.GameStatus),
	}
}

func (i *GameStatusRepoInMem) GetAll() (map[gs.GameStatusId]*gs.GameStatus, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()

	return i.Data, nil
}

func (i *GameStatusRepoInMem) FindById(gameStatusId gs.GameStatusId) (*gs.GameStatus, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()

	return i.Data[gameStatusId], nil
}

func (i *GameStatusRepoInMem) Upsert(o *gs.GameStatus) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.Data[o.ID()] = o
	return nil
}