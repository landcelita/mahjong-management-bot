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

	if len(scores) != 4 {
		return nil, fmt.Errorf("scoreは四人分必要あります。")
	}

	for _, jicha := range []jicha.Jicha{jicha.Toncha, jicha.Nancha, jicha.Shacha, jicha.Pecha} {
		if _, exist := scores[jicha]; !exist {
			return nil, fmt.Errorf(string(jicha) + "のscoreが指定されていません。")
		}
	}
	
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
