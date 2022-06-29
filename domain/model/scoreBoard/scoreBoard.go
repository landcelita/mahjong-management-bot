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
		scoreBoardId ScoreBoardId
		scores       map[jc.Jicha]sc.Score
		kyotaku      sc.Score
	}
)

func NewScoreBoard(
	scoreBoardId ScoreBoardId,
	scores map[jc.Jicha]sc.Score,
	kyotaku sc.Score,
) (*ScoreBoard, error) {

	if len(scores) != 4 {
		return nil, fmt.Errorf("scoreは四人分必要あります。")
	}

	for _, jicha := range []jc.Jicha{jc.Toncha, jc.Nancha, jc.Shacha, jc.Pecha} {
		if _, exist := scores[jicha]; !exist {
			return nil, fmt.Errorf(string(jicha) + "のscoreが指定されていません。")
		}
	}

	if score100k, _ := sc.NewScore(100000); !scores[jc.Toncha].
		Add(scores[jc.Nancha]).
		Add(scores[jc.Shacha]).
		Add(scores[jc.Pecha]).
		Add(kyotaku).
		Equals(*score100k) {

		return nil, fmt.Errorf("合計scoreは100,000と等しくなければいけません。")
	}

	if zero, _ := sc.NewScore(0); kyotaku.LessThan(*zero) {
		return nil, fmt.Errorf("kyoutakuは0を下回ってはいけません。")
	}

	scoreBoard := ScoreBoard{
		scoreBoardId: scoreBoardId,
		scores:       scores,
		kyotaku:      kyotaku,
	}

	return &scoreBoard, nil
}

func NewInitScoreBoard(
	initScore sc.Score,
) (*ScoreBoard, error) {
	initScores := map[jc.Jicha]sc.Score{
		jc.Toncha: initScore,
		jc.Nancha: initScore,
		jc.Shacha: initScore,
		jc.Pecha:  initScore,
	}
	sc0, _ := sc.NewScore(0)

	scoreBoard, err := NewScoreBoard(
		ScoreBoardId(uuid.New()),
		initScores,
		*sc0,
	)

	if err != nil {
		return nil, err
	}

	return scoreBoard, nil
}

func (scoreBoard *ScoreBoard) ID() ScoreBoardId {
	return scoreBoard.scoreBoardId
}

func (scoreBoard *ScoreBoard) AddKyotakuTo(score sc.Score) sc.Score {
	return score.Add(scoreBoard.kyotaku)
}

func (scoreBoard *ScoreBoard) GetKyotaku() sc.Score {
	return scoreBoard.kyotaku
}

func (scoreBoard *ScoreBoard) AddedWith(
	scoreDiff map[jc.Jicha]sc.Score,
	kyotakuDiff sc.Score,
) *ScoreBoard {
	for _, cha := range []jc.Jicha{jc.Toncha, jc.Nancha, jc.Shacha, jc.Pecha} {
		scoreBoard.scores[cha] = scoreBoard.scores[cha].Add(scoreDiff[cha])
	}
	scoreBoard.kyotaku = scoreBoard.kyotaku.Add(kyotakuDiff)

	return scoreBoard
}

func (scoreBoard *ScoreBoard) IsAnyoneTobi() bool {
	sc0, _ := sc.NewScore(0)
	for _, cha := range []jc.Jicha{jc.Toncha, jc.Nancha, jc.Shacha, jc.Pecha} {
		if scoreBoard.scores[cha].LessThan(*sc0) {
			return true
		}
	}
	return false
}

func (scoreBoard *ScoreBoard) No1AtOlast() jc.Jicha {
	firstPlayer := jc.Nancha
	for _, cha := range []jc.Jicha{jc.Nancha, jc.Shacha, jc.Pecha, jc.Toncha} {
		if (scoreBoard.scores[firstPlayer].LessThan(scoreBoard.scores[cha])) {
			firstPlayer = cha
		}
	}
	return firstPlayer
}

func (scoreBoard *ScoreBoard) AdvanceGameBaKyoku() {
	tmp := scoreBoard.scores[jc.Toncha]
	scoreBoard.scores[jc.Toncha] = scoreBoard.scores[jc.Nancha]
	scoreBoard.scores[jc.Nancha] = scoreBoard.scores[jc.Shacha]
	scoreBoard.scores[jc.Shacha] = scoreBoard.scores[jc.Pecha]
	scoreBoard.scores[jc.Pecha] = tmp
}
