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

	gameStatus, err := gs.NewInitGameStatus(
		tonpuOrHanchan,
		playerids,
	)
	if err != nil {
		return nil, nil, err
	}

	initScore, err := sc.NewScore(25000)
	if err != nil {
		return nil, nil, err
	}

	scoreBoard, err := sb.NewInitScoreBoard(*initScore)
	if err != nil {
		return nil, nil, err
	}

	if err := gamestartfactory.gameStatusRepo.Upsert(gameStatus); err != nil {
		return nil, nil, err
	}
	if err := gamestartfactory.scoreBoardRepo.Upsert(scoreBoard); err != nil {
		return nil, nil, err
	}

	return gameStatus, scoreBoard, nil
}
