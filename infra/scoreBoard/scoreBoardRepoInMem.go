package inmem

import (
	"sync"

	sb "github.com/landcelita/mahjong-management-bot/domain/model/scoreBoard"
)

type ScoreBoardRepoInMem struct {
	data  map[sb.ScoreBoardId]*sb.ScoreBoard
	mutex sync.RWMutex
}

func NewScoreBoardRepoInMem() *ScoreBoardRepoInMem {
	return &ScoreBoardRepoInMem{
		data: make(map[sb.ScoreBoardId]*sb.ScoreBoard),
	}
}

func (i *ScoreBoardRepoInMem) FindById(scoreBoardId sb.ScoreBoardId) (*sb.ScoreBoard, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()

	return i.data[scoreBoardId], nil
}

func (i *ScoreBoardRepoInMem) Upsert(o *sb.ScoreBoard) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.data[o.ID()] = o
	return nil
}