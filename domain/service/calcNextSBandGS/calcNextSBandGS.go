package calcnxsbgs

import (
	"fmt"

	gs "github.com/landcelita/mahjong-management-bot/domain/model/gameStatus"
	"github.com/landcelita/mahjong-management-bot/domain/model/jicha"
	kr "github.com/landcelita/mahjong-management-bot/domain/model/kyokuResult"
	"github.com/landcelita/mahjong-management-bot/domain/model/score"
	sb "github.com/landcelita/mahjong-management-bot/domain/model/scoreBoard"
)

type gameProgressType string

const (
	progressKyoku = gameProgressType("progressKyoku")
	progressHonba = gameProgressType("progressHonba")
	gameIsOver    = gameProgressType("gameIsOver")
)

func CalcNextSBandGS(
	gameStatus *gs.GameStatus,
	scoreBoard *sb.ScoreBoard,
	kyokuResult *kr.KyokuResult,
) (
	nextScoreBoard *sb.ScoreBoard,
	nextGameStatus *gs.GameStatus,
	err error,
) {
	nextScoreBoard, err = calcNextScoreBoard(scoreBoard, kyokuResult)
	if err != nil {
		return
	}
	switch calcGameProgressType(nextScoreBoard, kyokuResult, gameStatus) {
	case progressKyoku:
		gameStatus.AdvanceGameBaKyoku()
		nextScoreBoard.AdvanceGameBaKyoku()
	case progressHonba:
		gameStatus.AdvanceGameHonba()
	case gameIsOver:
		gameStatus.GameOver()
	}

	nextGameStatus = gameStatus

	return
}

func calcNextScoreBoard(
	scoreBoard *sb.ScoreBoard,
	kyokuResult *kr.KyokuResult,
) (
	nextScoreBoard *sb.ScoreBoard,
	err error,
) {
	scoreDiff, kyotakuDiff, err := calcScoreAndKyotakuDiff(kyokuResult, scoreBoard)
	if err != nil {
		return
	}

	nextScoreBoard = scoreBoard.AddedWith(scoreDiff, kyotakuDiff)
	return
}

func calcGameProgressType(
	scoreBoard *sb.ScoreBoard,
	kyokuResult *kr.KyokuResult,
	gameStatus *gs.GameStatus,
) (
	gameProgress gameProgressType,
) {
	if scoreBoard.IsAnyoneTobi() {
		return gameIsOver
	}
	if gameStatus.IsOlast() {
		if kyokuResult.IsTonchaRonOrTsumoOrTenpai() {
			if scoreBoard.No1AtOlast() == jicha.Toncha {
				return gameIsOver
			} else {
				return progressHonba
			}
		} else {
			return gameIsOver
		}
	} else {
		if kyokuResult.IsTonchaRonOrTsumoOrTenpai() {
			return progressHonba
		} else {
			return progressKyoku
		}
	}
}

func ceil100(n int) int {
	return (n + 99) / 100 * 100
}

func calcScoreAndKyotakuDiff(
	kyokuResult *kr.KyokuResult,
	scoreBoard *sb.ScoreBoard,
) (
	scoreDiff map[jicha.Jicha]score.Score,
	kyotakuDiff score.Score,
	err error,
) {
	switch kyokuResult.GetKyokuEndType() {
	case kr.Tsumo:
		scoreDiff, kyotakuDiff, err = calcScoreAndKyotakuDiffWhenTsumo(kyokuResult, scoreBoard)
	case kr.Ron:
		scoreDiff, kyotakuDiff, err = calcScoreAndKyotakuDiffWhenRon(kyokuResult, scoreBoard)
	case kr.Ryukyoku:
		scoreDiff, kyotakuDiff, err = calcScoreAndKyotakuDiffWhenRyukyoku(kyokuResult, scoreBoard)
	default:
		err = fmt.Errorf("KyokuEndTypeが不正です")
	}

	return
}

func calcScoreAndKyotakuDiffWhenTsumo(
	kyokuResult *kr.KyokuResult,
	scoreBoard *sb.ScoreBoard,
) (
	scoreDiff map[jicha.Jicha]score.Score,
	kyotakuDiff score.Score,
	err error,
) {
	scoreDiff = map[jicha.Jicha]score.Score{}
	baseScoreUint, err := kyokuResult.CalcBaseScore()
	if err != nil {
		return
	}
	baseScore, err := score.NewScore(ceil100(int(baseScoreUint)))
	if err != nil {
		return
	}
	baseScore2, err := score.NewScore(ceil100(int(baseScoreUint) * 2))
	if err != nil {
		return
	}
	honbaScore, err := score.NewScore(int(kyokuResult.BKH().Honba()) * 100)
	if err != nil {
		return
	}
	winner, err := kyokuResult.WhoTsumo()
	if err != nil {
		return
	}

	if *winner == jicha.Toncha {
		for _, cha := range []jicha.Jicha{jicha.Toncha, jicha.Nancha, jicha.Shacha, jicha.Pecha} {
			if cha == jicha.Toncha {
				scoreDiff[cha] = (*baseScore).Mul(3).Add((honbaScore).Mul(3))
			} else {
				scoreDiff[cha] = (*baseScore).Mul(-1).Add((honbaScore).Mul(-1))
			}
		}
	} else {
		for _, cha := range []jicha.Jicha{jicha.Toncha, jicha.Nancha, jicha.Shacha, jicha.Pecha} {
			if cha == jicha.Toncha {
				scoreDiff[cha] = (*baseScore2).Mul(-1).Add((honbaScore).Mul(-1))
			} else if cha == *winner {
				scoreDiff[cha] = (*baseScore).Mul(2).Add(*baseScore2).Add((honbaScore).Mul(3))
			} else {
				scoreDiff[cha] = (*baseScore).Mul(-1).Add((honbaScore).Mul(-1))
			}
		}
	}

	err = addRiichiAndKyotakuToDiff(kyokuResult, scoreBoard, &scoreDiff, *winner)
	kyotakuDiff = scoreBoard.GetKyotaku().Mul(-1)

	return
}

func calcScoreAndKyotakuDiffWhenRon(
	kyokuResult *kr.KyokuResult,
	scoreBoard *sb.ScoreBoard,
) (
	scoreDiff map[jicha.Jicha]score.Score,
	kyotakuDiff score.Score,
	err error,
) {
	scoreDiff = map[jicha.Jicha]score.Score{}
	baseScoreUint, err := kyokuResult.CalcBaseScore()
	if err != nil {
		return
	}
	honbaScore, err := score.NewScore(int(kyokuResult.BKH().Honba()) * 100)
	if err != nil {
		return
	}
	winner, err := kyokuResult.WhoRonWinner()
	if err != nil {
		return
	}
	loser, err := kyokuResult.WhoRonLoser()
	if err != nil {
		return
	}
	zeroScore, err := score.NewScore(0)
	if err != nil {
		return
	}
	baseScore3, err := score.NewScore(ceil100(int(baseScoreUint * 3)))
	if err != nil {
		return
	}
	baseScore4, err := score.NewScore(ceil100(int(baseScoreUint * 4)))
	if err != nil {
		return
	}

	for _, cha := range []jicha.Jicha{jicha.Toncha, jicha.Nancha, jicha.Shacha, jicha.Pecha} {
		if cha == *winner {
			if *winner == jicha.Toncha {
				scoreDiff[cha] = (*baseScore3).Add((*honbaScore).Mul(3))
			} else {
				scoreDiff[cha] = (*baseScore4).Add((*honbaScore).Mul(3))
			}
		} else if cha == *loser {
			if *winner == jicha.Toncha {
				scoreDiff[cha] = (*baseScore3).Mul(-1).Add((*honbaScore).Mul(-3))
			} else {
				scoreDiff[cha] = (*baseScore4).Mul(-1).Add((*honbaScore).Mul(-3))
			}
		} else {
			scoreDiff[cha] = *zeroScore
		}
	}

	err = addRiichiAndKyotakuToDiff(kyokuResult, scoreBoard, &scoreDiff, *winner)
	kyotakuDiff = scoreBoard.GetKyotaku().Mul(-1)

	return
}

func calcScoreAndKyotakuDiffWhenRyukyoku(
	kyokuResult *kr.KyokuResult,
	scoreBoard *sb.ScoreBoard,
) (
	scoreDiff map[jicha.Jicha]score.Score,
	kyotakuDiff score.Score,
	err error,
) {
	score0, err := score.NewScore(0)
	if err != nil {
		return
	}
	score1000, err := score.NewScore(1000)
	if err != nil {
		return
	}
	score1500, err := score.NewScore(1500)
	if err != nil {
		return
	}
	score3000, err := score.NewScore(3000)
	if err != nil {
		return
	}
	tenpaiers, err := kyokuResult.WhoTenpai()
	if err != nil {
		return
	}
	riichiers := kyokuResult.WhoRiichi()
	var gainTenpaiScore, loseTenpaiScore score.Score

	switch len(*tenpaiers) {
	case 0, 4:
		gainTenpaiScore = *score0
		loseTenpaiScore = *score0
	case 1:
		gainTenpaiScore = *score3000
		loseTenpaiScore = (*score1000).Mul(-1)
	case 2:
		gainTenpaiScore = *score1500
		loseTenpaiScore = (*score1500).Mul(-1)
	case 3:
		gainTenpaiScore = *score1000
		loseTenpaiScore = (*score3000).Mul(-1)
	default:
		err = fmt.Errorf("tenpaiersの人数が不正です(0-4人)")
		return
	}
	scoreDiff = map[jicha.Jicha]score.Score{}

	for _, cha := range []jicha.Jicha{jicha.Toncha, jicha.Nancha, jicha.Shacha, jicha.Pecha} {
		if _, exist := (*tenpaiers)[cha]; exist {
			scoreDiff[cha] = gainTenpaiScore
		} else {
			scoreDiff[cha] = loseTenpaiScore
		}
	}

	for _, cha := range []jicha.Jicha{jicha.Toncha, jicha.Nancha, jicha.Shacha, jicha.Pecha} {
		if _, exist := (*riichiers)[cha]; exist {
			scoreDiff[cha] = scoreDiff[cha].Add((*score1000).Mul(-1))
		}
	}

	kyotakuDiff = score1000.Mul(len(*riichiers))

	return
}

func addRiichiAndKyotakuToDiff(
	kyokuResult *kr.KyokuResult,
	scoreBoard *sb.ScoreBoard,
	scoreDiff *map[jicha.Jicha]score.Score,
	winner jicha.Jicha,
) error {
	riichiScore, err := score.NewScore(1000)
	if err != nil {
		return err
	}
	riichiers := kyokuResult.WhoRiichi()

	(*scoreDiff)[winner] = (*scoreDiff)[winner].Add(scoreBoard.GetKyotaku())

	for _, cha := range []jicha.Jicha{jicha.Toncha, jicha.Nancha, jicha.Shacha, jicha.Pecha} {
		if cha == winner {
			(*scoreDiff)[cha] = (*scoreDiff)[cha].Add((*riichiScore).Mul(len(*riichiers)))
		}
		if _, exist := (*riichiers)[cha]; exist {
			(*scoreDiff)[cha] = (*scoreDiff)[cha].Add((*riichiScore).Mul(-1))
		}
	}

	return nil
}
