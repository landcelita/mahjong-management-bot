package score

import (
	"fmt"
)

type (
	Score struct {
		score	int
	}
)

func NewScore(scoreInt int) (*Score, error){
	if(scoreInt % 100 != 0) {
		return nil, fmt.Errorf("scoreは必ず100の倍数でなければなりません。")
	}

	score := Score {
		score:	scoreInt,
	}

	return &score, nil
}
