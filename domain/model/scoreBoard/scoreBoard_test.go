package scoreboard

import (
	"mahjong/domain/model/score"
	"testing"
	"github.com/google/uuid"
	. "mahjong/testutil"
)

func TestNewScoreBoard(t *testing.T) {
	scoreBoardId := ScoreBoardId(uuid.New())

	type args struct {
		scoreBoardId ScoreBoardId
		scores       [4]score.Score
		kyotaku      score.Score
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    	"ok 1",
			args:    	args{
				scoreBoardId:	scoreBoardId,
				scores:			[4]score.Score{
					First(score.NewScore(-100)),
					First(score.NewScore(0)),
					First(score.NewScore(0)),
					First(score.NewScore(100100)),
				},
				kyotaku:		First(score.NewScore(0)),
			},
			wantErr:	false,
		},
		{
			name:    	"ok 2",
			args:    	args{
				scoreBoardId:	scoreBoardId,
				scores:			[4]score.Score{
					First(score.NewScore(-100)),
					First(score.NewScore(-100000)),
					First(score.NewScore(0)),
					First(score.NewScore(100100)),
				},
				kyotaku:		First(score.NewScore(100000)),
			},
			wantErr:	false,
		},
		{
			name:    	"ng 1 (sum is larger than 100000)",
			args:    	args{
				scoreBoardId:	scoreBoardId,
				scores:			[4]score.Score{
					First(score.NewScore(25000)),
					First(score.NewScore(25000)),
					First(score.NewScore(25000)),
					First(score.NewScore(25000)),
				},
				kyotaku:		First(score.NewScore(1000)),
			},
			wantErr:	true,
		},
		{
			name:    	"ng 2 (kyoutaku is less than 0)",
			args:    	args{
				scoreBoardId:	scoreBoardId,
				scores:			[4]score.Score{
					First(score.NewScore(25000)),
					First(score.NewScore(25000)),
					First(score.NewScore(25000)),
					First(score.NewScore(26000)),
				},
				kyotaku:		First(score.NewScore(-1000)),
			},
			wantErr:	true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewScoreBoard(tt.args.scoreBoardId, tt.args.scores, tt.args.kyotaku)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewScoreBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
