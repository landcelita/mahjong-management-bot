package scoreboard

import (
	"github.com/google/uuid"
)

type ScoreBoardId uuid.UUID
type Score int

type (
	ScoreBoard struct {
		scoreBoardId	ScoreBoardId
		scores			[4]Score
		kyotaku			Score
	}
)

// func NewScoreBoard(
// 	scoreBoardId	ScoreBoardId,
// 	scores			[4]Score,
// 	kyotaku			Score,
// 	) (*ScoreBoard, error) {
	
	
// }
