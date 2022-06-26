package scoreboard

type ScoreBoardIRepo interface {
	Save(scoreBoard *ScoreBoard) error
	FindById(scoreBoardId ScoreBoardId) (*ScoreBoard, error)
	Update(scoreBoard *ScoreBoard) error
	Delete(scoreBoard *ScoreBoard) error
}