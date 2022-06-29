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
	if scoreInt % 100 != 0 {
		return nil, fmt.Errorf("scoreは必ず100の倍数でなければなりません。")
	}

	score := Score {
		score:	scoreInt,
	}

	return &score, nil
}

func (score Score) Add(otherScore Score) Score {
	return Score {
		score: score.score + otherScore.score,
	}
}

func (score Score) Mul(n int) Score {
	return Score {
		score: score.score * n,
	}
}

func (score Score) Equals(otherScore Score) bool {
	return score.score == otherScore.score
}

func (score Score) LessThan(otherScore Score) bool {
	return score.score < otherScore.score
}