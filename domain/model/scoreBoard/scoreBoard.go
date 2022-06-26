package scoreboard

import (
	"fmt"

	"github.com/google/uuid"
	jc "github.com/landcelita/mahjong-management-bot/domain/model/jicha"
	sc "github.com/landcelita/mahjong-management-bot/domain/model/score"
)

type ScoreBoardId uuid.UUID

type (
	ScoreBoard struct {
		scoreBoardId	ScoreBoardId
		scores			map[jc.Jicha]sc.Score
		kyotaku			sc.Score
	}
)

func NewScoreBoard(
	scoreBoardId	ScoreBoardId,
	scores			map[jc.Jicha]sc.Score,
	kyotaku			sc.Score,
) (*ScoreBoard, error) {

	if len(scores) != 4 {
		return nil, fmt.Errorf("scoreは四人分必要あります。")
	}

	for _, jicha := range []jc.Jicha{jc.Toncha, jc.Nancha, jc.Shacha, jc.Pecha} {
		if _, exist := scores[jicha]; !exist {
			return nil, fmt.Errorf(string(jicha) + "のscoreが指定されていません。")
		}
	}
	
	if score100k, _ := sc.NewScore(100000);
	!scores[jc.Toncha].
		Add(scores[jc.Nancha]).
		Add(scores[jc.Shacha]).
		Add(scores[jc.Pecha]).
		Add(kyotaku).
		Equals(*score100k){
		
		return nil, fmt.Errorf("合計scoreは100,000と等しくなければいけません。")
	}

	if zero, _ := sc.NewScore(0);
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
