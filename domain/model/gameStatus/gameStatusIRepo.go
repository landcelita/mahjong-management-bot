package gamestatus

type GameStatusIRepo interface {
	GetAll() (map[GameStatusId]*GameStatus, error)
	FindById(gameStatusId GameStatusId) (*GameStatus, error)
	Upsert(gameStatus *GameStatus) error
}