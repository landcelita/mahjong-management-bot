package scoreboard

import (
	"fmt"
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
	
	if score100k, _ := score.NewScore(100000);
	!scores[0].
		Add(scores[1]).
		Add(scores[2]).
		Add(scores[3]).
		Add(kyotaku).
		Equals(*score100k){
		
		return nil, fmt.Errorf("合計scoreは100,000と等しくなければいけません。")
	}

	if zero, _ := score.NewScore(0);
	kyotaku.LessThan(*zero){
		return nil, fmt.Errorf("kyoutakuは0を下回ってはいけません。")
	}

	scoreBoard := ScoreBoard{
		scoreBoardId:	scoreBoardId,
		scores:			scores,
		kyotaku:		kyotaku,
	}

	return &scoreBoard, nil
}
