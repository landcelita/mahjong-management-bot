package inmem

import (
	"sync"

	gs "github.com/landcelita/mahjong-management-bot/domain/model/gameStatus"
)

type GameStatusRepoInMem struct {
	data  map[gs.GameStatusId]*gs.GameStatus
	mutex sync.RWMutex
}

func NewGameStatusRepoInMem() *GameStatusRepoInMem {
	return &GameStatusRepoInMem{
		data: make(map[gs.GameStatusId]*gs.GameStatus),
	}
}

func (i *GameStatusRepoInMem) FindById(gameStatusId gs.GameStatusId) (*gs.GameStatus, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()

	return i.data[gameStatusId], nil
}

func (i *GameStatusRepoInMem) Upsert(o *gs.GameStatus) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.data[o.ID()] = o
	return nil
}