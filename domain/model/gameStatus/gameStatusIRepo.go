package gamestatus

type GameStatusIRepo interface {
	FindById(gameStatusId GameStatusId) (*GameStatus, error)
	Upsert(gameStatus *GameStatus) error
}