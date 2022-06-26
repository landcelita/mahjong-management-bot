package gamestatus

type GameStatusIRepo interface {
	Save(gameStatus *GameStatus) error
	FindById(gameStatusId GameStatusId) (*GameStatus, error)
	Update(gameStatus *GameStatus) error
	Delete(gameStatus *GameStatus) error
}