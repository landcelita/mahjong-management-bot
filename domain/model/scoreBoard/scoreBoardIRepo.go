package scoreboard

type ScoreBoardIRepo interface {
	GetAll() (map[ScoreBoardId]*ScoreBoard, error,)
	FindById(scoreBoardId ScoreBoardId) (*ScoreBoard, error)
	Upsert(scoreBoard *ScoreBoard) error
}