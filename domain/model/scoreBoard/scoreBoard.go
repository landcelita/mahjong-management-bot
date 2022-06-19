package scoreboard

import (
	"github.com/google/uuid"
	"mahjong/domain/model/score"
)

type ScoreBoardId uuid.UUID

type (
	ScoreBoard struct {
		scoreBoardId	ScoreBoardId
		scores			[4]score.Score
		kyotaku			score.Score
	}
)

func NewScoreBoard(
	scoreBoardId	ScoreBoardId,
	scores			[4]score.Score,
	kyotaku			score.Score,
	) (*ScoreBoard, error) {
	
	if score100k, _ := score.NewScore(1000000);
	scores[0].Add(scores[1]).Add(scores[2]).Add(scores[3]).Add(kyotaku)
	 {

	}
}
