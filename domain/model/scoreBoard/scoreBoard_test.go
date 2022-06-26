package scoreboard

import (
	"testing"

	"github.com/google/uuid"
	jc "github.com/landcelita/mahjong-management-bot/domain/model/jicha"
	sc "github.com/landcelita/mahjong-management-bot/domain/model/score"
	. "github.com/landcelita/mahjong-management-bot/testutil"
)

func TestNewScoreBoard(t *testing.T) {
	scoreBoardId := ScoreBoardId(uuid.New())

	type args struct {
		scoreBoardId ScoreBoardId
		scores       map[jc.Jicha]sc.Score
		kyotaku      sc.Score
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "正常系 負の得点の人がいる場合",
			args: args{
				scoreBoardId: scoreBoardId,
				scores: map[jc.Jicha]sc.Score{
					jc.Toncha: FirstPtoV(sc.NewScore(-100)),
					jc.Nancha: FirstPtoV(sc.NewScore(0)),
					jc.Shacha: FirstPtoV(sc.NewScore(0)),
					jc.Pecha:  FirstPtoV(sc.NewScore(100100)),
				},
				kyotaku: FirstPtoV(sc.NewScore(0)),
			},
			wantErr: false,
		},
		{
			name: "正常系 負の得点の人が二人いる場合",
			args: args{
				scoreBoardId: scoreBoardId,
				scores: map[jc.Jicha]sc.Score{
					jc.Toncha: FirstPtoV(sc.NewScore(-100)),
					jc.Nancha: FirstPtoV(sc.NewScore(-100000)),
					jc.Shacha: FirstPtoV(sc.NewScore(0)),
					jc.Pecha:  FirstPtoV(sc.NewScore(100100)),
				},
				kyotaku: FirstPtoV(sc.NewScore(100000)),
			},
			wantErr: false,
		},
		{
			name: "異常系 合計得点が100000でない場合",
			args: args{
				scoreBoardId: scoreBoardId,
				scores: map[jc.Jicha]sc.Score{
					jc.Toncha: FirstPtoV(sc.NewScore(25000)),
					jc.Nancha: FirstPtoV(sc.NewScore(25000)),
					jc.Shacha: FirstPtoV(sc.NewScore(25000)),
					jc.Pecha:  FirstPtoV(sc.NewScore(25000)),
				},
				kyotaku: FirstPtoV(sc.NewScore(1000)),
			},
			wantErr: true,
		},
		{
			name: "異常系 kyoutakuが負な場合",
			args: args{
				scoreBoardId: scoreBoardId,
				scores: map[jc.Jicha]sc.Score{
					jc.Toncha: FirstPtoV(sc.NewScore(25000)),
					jc.Nancha: FirstPtoV(sc.NewScore(25000)),
					jc.Shacha: FirstPtoV(sc.NewScore(25000)),
					jc.Pecha:  FirstPtoV(sc.NewScore(26000)),
				},
				kyotaku: FirstPtoV(sc.NewScore(-1000)),
			},
			wantErr: true,
		},
		{
			name: "異常系 scoreが4つ分指定されていない場合",
			args: args{
				scoreBoardId: scoreBoardId,
				scores: map[jc.Jicha]sc.Score{
					jc.Toncha: FirstPtoV(sc.NewScore(25000)),
					jc.Nancha: FirstPtoV(sc.NewScore(25000)),
					jc.Shacha: FirstPtoV(sc.NewScore(25000)),
				},
				kyotaku: FirstPtoV(sc.NewScore(25000)),
			},
			wantErr: true,
		},
		{
			name: "異常系 playerの形式が間違っている場合",
			args: args{
				scoreBoardId: scoreBoardId,
				scores: map[jc.Jicha]sc.Score{
					jc.Toncha:       FirstPtoV(sc.NewScore(25000)),
					jc.Nancha:       FirstPtoV(sc.NewScore(25000)),
					jc.Shacha:       FirstPtoV(sc.NewScore(25000)),
					jc.Jicha("Nan"): FirstPtoV(sc.NewScore(0)),
				},
				kyotaku: FirstPtoV(sc.NewScore(25000)),
			},
			wantErr: true,
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
