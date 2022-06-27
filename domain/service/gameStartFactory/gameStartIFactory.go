package gamestartfactory

import (
	gs "github.com/landcelita/mahjong-management-bot/domain/model/gameStatus"
	pid "github.com/landcelita/mahjong-management-bot/domain/model/playerId"
	sb "github.com/landcelita/mahjong-management-bot/domain/model/scoreBoard"
	toh "github.com/landcelita/mahjong-management-bot/domain/model/tonpuOrHanchan"
)

type GameStartIFactory interface {
	StartNewGame(
		tonpuOrHanchan	toh.TonpuOrHanchan,
		tonchaid		pid.PlayerId,
		nanchaid		pid.PlayerId,
		shachaid		pid.PlayerId,
		pechaid			pid.PlayerId,
	) (*gs.GameStatus, *sb.ScoreBoard, error)
}