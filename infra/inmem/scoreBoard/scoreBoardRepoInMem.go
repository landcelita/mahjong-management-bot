package sbinmem

import (
	"sync"

	sb "github.com/landcelita/mahjong-management-bot/domain/model/scoreBoard"
)

type ScoreBoardRepoInMem struct {
	Data  map[sb.ScoreBoardId]*sb.ScoreBoard
	mutex sync.RWMutex
}

func NewScoreBoardRepoInMem() sb.ScoreBoardIRepo {
	return &ScoreBoardRepoInMem{
		Data: make(map[sb.ScoreBoardId]*sb.ScoreBoard),
	}
}

func (i *ScoreBoardRepoInMem) FindById(scoreBoardId sb.ScoreBoardId) (*sb.ScoreBoard, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()

	return i.Data[scoreBoardId], nil
}

func (i *ScoreBoardRepoInMem) Upsert(o *sb.ScoreBoard) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.Data[o.ID()] = o
	return nil
}