package scoreboard

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/landcelita/mahjong-management-bot/domain/model/jicha"
	"github.com/landcelita/mahjong-management-bot/domain/model/score"
)

type ScoreBoardId uuid.UUID

type (
	ScoreBoard struct {
		scoreBoardId	ScoreBoardId
		scores			map[jicha.Jicha]score.Score
		kyotaku			score.Score
	}
)

func NewScoreBoard(
	scoreBoardId	ScoreBoardId,
	scores			map[jicha.Jicha]score.Score,
	kyotaku			score.Score,
	) (*ScoreBoard, error) {
	
	if score100k, _ := score.NewScore(100000);
	!scores[jicha.Toncha].
		Add(scores[jicha.Nancha]).
		Add(scores[jicha.Shacha]).
		Add(scores[jicha.Pecha]).
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
