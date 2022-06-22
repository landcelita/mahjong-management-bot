package gamestatus

type GameStatusRepository interface {
	Save(gameStatus *GameStatus) error
	FindByGameStatusId(gameStatusId *GameStatusId) (*GameStatus, error)
}