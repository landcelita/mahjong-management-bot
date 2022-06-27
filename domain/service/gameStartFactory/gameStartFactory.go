package gamestartfactory

import (
	gs "github.com/landcelita/mahjong-management-bot/domain/model/gameStatus"
	jc "github.com/landcelita/mahjong-management-bot/domain/model/jicha"
	pid "github.com/landcelita/mahjong-management-bot/domain/model/playerId"
	sc "github.com/landcelita/mahjong-management-bot/domain/model/score"
	sb "github.com/landcelita/mahjong-management-bot/domain/model/scoreBoard"
	toh "github.com/landcelita/mahjong-management-bot/domain/model/tonpuOrHanchan"
)

type GameStartFactory struct {
	gameStatusRepo		gs.GameStatusIRepo
	scoreBoardRepo		sb.ScoreBoardIRepo
}

func NewGameStartFactory(
	grepo gs.GameStatusIRepo,
	srepo sb.ScoreBoardIRepo,
) (GameStartIFactory, error) {
	return &GameStartFactory{
		gameStatusRepo: grepo,
		scoreBoardRepo: srepo,
	}, nil
}

func (gamestartfactory *GameStartFactory) StartNewGame(
	tonpuOrHanchan	toh.TonpuOrHanchan,
	tonchaid		pid.PlayerId,
	nanchaid		pid.PlayerId,
	shachaid		pid.PlayerId,
	pechaid			pid.PlayerId,
) (
	*gs.GameStatus,
	*sb.ScoreBoard,
	error,
) {
	playerids := map[jc.Jicha]pid.PlayerId{
		jc.Toncha:		tonchaid,
		jc.Nancha:		nanchaid,
		jc.Shacha:		shachaid,
		jc.Pecha:		pechaid,
	}

	gameStatus, errgs := gs.NewInitGameStatus(
		tonpuOrHanchan,
		playerids,
	)
	if errgs != nil {
		return nil, nil, errgs
	}

	initScore, erris := sc.NewScore(25000)
	if erris != nil {
		return nil, nil, erris
	}

	scoreBoard, errsb := sb.NewInitScoreBoard(*initScore)
	if errsb != nil {
		return nil, nil, errsb
	}

	if e := gamestartfactory.gameStatusRepo.Upsert(gameStatus); e != nil {
		return nil, nil, e
	}
	if e := gamestartfactory.scoreBoardRepo.Upsert(scoreBoard); e != nil {
		return nil, nil, e
	}

	return gameStatus, scoreBoard, nil
}
