package scoreboard

type ScoreBoardIRepo interface {
	FindById(scoreBoardId ScoreBoardId) (*ScoreBoard, error)
	Upsert(scoreBoard *ScoreBoard) error
}