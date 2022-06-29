package scoreboard

import (
	"reflect"
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

func TestScoreBoard_AddKyotakuTo(t *testing.T) {
	type fields struct {
		scoreBoardId ScoreBoardId
		scores       map[jc.Jicha]sc.Score
		kyotaku      sc.Score
	}
	type args struct {
		score sc.Score
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   sc.Score
	}{
		{
			name: "kyotakuが0のとき",
			fields: fields{
				scoreBoardId: ScoreBoardId(uuid.New()),
				scores: map[jc.Jicha]sc.Score{
					jc.Toncha: FirstPtoV(sc.NewScore(25000)),
					jc.Nancha: FirstPtoV(sc.NewScore(25000)),
					jc.Shacha: FirstPtoV(sc.NewScore(25000)),
					jc.Pecha:  FirstPtoV(sc.NewScore(25000)),
				},
				kyotaku: FirstPtoV(sc.NewScore(0)),
			},
			args: args{
				score: FirstPtoV(sc.NewScore(200)),
			},
			want: FirstPtoV(sc.NewScore(200)),
		},
		{
			name: "kyotakuが1000のとき",
			fields: fields{
				scoreBoardId: ScoreBoardId(uuid.New()),
				scores: map[jc.Jicha]sc.Score{
					jc.Toncha: FirstPtoV(sc.NewScore(25000)),
					jc.Nancha: FirstPtoV(sc.NewScore(25000)),
					jc.Shacha: FirstPtoV(sc.NewScore(25000)),
					jc.Pecha:  FirstPtoV(sc.NewScore(25000)),
				},
				kyotaku: FirstPtoV(sc.NewScore(1000)),
			},
			args: args{
				score: FirstPtoV(sc.NewScore(200)),
			},
			want: FirstPtoV(sc.NewScore(1200)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scoreBoard := &ScoreBoard{
				scoreBoardId: tt.fields.scoreBoardId,
				scores:       tt.fields.scores,
				kyotaku:      tt.fields.kyotaku,
			}
			if got := scoreBoard.AddKyotakuTo(tt.args.score); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScoreBoard.AddKyotakuTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScoreBoard_AddedWith(t *testing.T) {
	sbid := ScoreBoardId(uuid.New())
	type fields struct {
		scoreBoardId ScoreBoardId
		scores       map[jc.Jicha]sc.Score
		kyotaku      sc.Score
	}
	type args struct {
		scoreDiff   map[jc.Jicha]sc.Score
		kyotakuDiff sc.Score
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ScoreBoard
	}{
		{
			name: "score: (25000, 25000, 25000, 25000) + (5000, -5000, -1000, -1000), kyotaku: 0 + 2000",
			fields: fields{
				scoreBoardId: sbid,
				scores: map[jc.Jicha]sc.Score{
					jc.Toncha: FirstPtoV(sc.NewScore(25000)),
					jc.Nancha: FirstPtoV(sc.NewScore(25000)),
					jc.Shacha: FirstPtoV(sc.NewScore(25000)),
					jc.Pecha:  FirstPtoV(sc.NewScore(25000)),
				},
				kyotaku: FirstPtoV(sc.NewScore(0)),
			},
			args: args{
				scoreDiff: map[jc.Jicha]sc.Score{
					jc.Toncha: FirstPtoV(sc.NewScore(5000)),
					jc.Nancha: FirstPtoV(sc.NewScore(-5000)),
					jc.Shacha: FirstPtoV(sc.NewScore(-1000)),
					jc.Pecha:  FirstPtoV(sc.NewScore(-1000)),
				},
				kyotakuDiff: FirstPtoV(sc.NewScore(2000)),
			},
			want: &ScoreBoard{
				scoreBoardId: sbid,
				scores: map[jc.Jicha]sc.Score{
					jc.Toncha: FirstPtoV(sc.NewScore(30000)),
					jc.Nancha: FirstPtoV(sc.NewScore(20000)),
					jc.Shacha: FirstPtoV(sc.NewScore(24000)),
					jc.Pecha:  FirstPtoV(sc.NewScore(24000)),
				},
				kyotaku: FirstPtoV(sc.NewScore(2000)),
			},
		},
		{
			name: "score: (0, 0, 0, 90000) + (0, 0, -1000, -1000), kyotaku: 10000 + 2000",
			fields: fields{
				scoreBoardId: sbid,
				scores: map[jc.Jicha]sc.Score{
					jc.Toncha: FirstPtoV(sc.NewScore(0)),
					jc.Nancha: FirstPtoV(sc.NewScore(0)),
					jc.Shacha: FirstPtoV(sc.NewScore(0)),
					jc.Pecha:  FirstPtoV(sc.NewScore(90000)),
				},
				kyotaku: FirstPtoV(sc.NewScore(10000)),
			},
			args: args{
				scoreDiff: map[jc.Jicha]sc.Score{
					jc.Toncha: FirstPtoV(sc.NewScore(0)),
					jc.Nancha: FirstPtoV(sc.NewScore(0)),
					jc.Shacha: FirstPtoV(sc.NewScore(-1000)),
					jc.Pecha:  FirstPtoV(sc.NewScore(-1000)),
				},
				kyotakuDiff: FirstPtoV(sc.NewScore(2000)),
			},
			want: &ScoreBoard{
				scoreBoardId: sbid,
				scores: map[jc.Jicha]sc.Score{
					jc.Toncha: FirstPtoV(sc.NewScore(0)),
					jc.Nancha: FirstPtoV(sc.NewScore(0)),
					jc.Shacha: FirstPtoV(sc.NewScore(-1000)),
					jc.Pecha:  FirstPtoV(sc.NewScore(89000)),
				},
				kyotaku: FirstPtoV(sc.NewScore(12000)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scoreBoard := &ScoreBoard{
				scoreBoardId: tt.fields.scoreBoardId,
				scores:       tt.fields.scores,
				kyotaku:      tt.fields.kyotaku,
			}
			if got := scoreBoard.AddedWith(tt.args.scoreDiff, tt.args.kyotakuDiff); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScoreBoard.AddedWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScoreBoard_IsAnyoneTobi(t *testing.T) {
	type fields struct {
		scoreBoardId ScoreBoardId
		scores       map[jc.Jicha]sc.Score
		kyotaku      sc.Score
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "tobiがいる場合",
			fields: fields{
				scoreBoardId: ScoreBoardId(uuid.New()),
				scores: map[jc.Jicha]sc.Score{
					jc.Toncha: FirstPtoV(sc.NewScore(0)),
					jc.Nancha: FirstPtoV(sc.NewScore(0)),
					jc.Shacha: FirstPtoV(sc.NewScore(-1000)),
					jc.Pecha:  FirstPtoV(sc.NewScore(89000)),
				},
				kyotaku: FirstPtoV(sc.NewScore(12000)),
			},
			want: true,
		},
		{
			name: "tobiがいない場合",
			fields: fields{
				scoreBoardId: ScoreBoardId(uuid.New()),
				scores: map[jc.Jicha]sc.Score{
					jc.Toncha: FirstPtoV(sc.NewScore(0)),
					jc.Nancha: FirstPtoV(sc.NewScore(0)),
					jc.Shacha: FirstPtoV(sc.NewScore(1000)),
					jc.Pecha:  FirstPtoV(sc.NewScore(89000)),
				},
				kyotaku: FirstPtoV(sc.NewScore(10000)),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scoreBoard := &ScoreBoard{
				scoreBoardId: tt.fields.scoreBoardId,
				scores:       tt.fields.scores,
				kyotaku:      tt.fields.kyotaku,
			}
			if got := scoreBoard.IsAnyoneTobi(); got != tt.want {
				t.Errorf("ScoreBoard.IsAnyoneTobi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScoreBoard_No1AtOlast(t *testing.T) {
	type fields struct {
		scoreBoardId ScoreBoardId
		scores       map[jc.Jicha]sc.Score
		kyotaku      sc.Score
	}
	tests := []struct {
		name   string
		fields fields
		want   jc.Jicha
	}{
		{
			name: "全員同点",
			fields: fields{
				scoreBoardId: ScoreBoardId(uuid.New()),
				scores: map[jc.Jicha]sc.Score{
					jc.Toncha: FirstPtoV(sc.NewScore(25000)),
					jc.Nancha: FirstPtoV(sc.NewScore(25000)),
					jc.Shacha: FirstPtoV(sc.NewScore(25000)),
					jc.Pecha:  FirstPtoV(sc.NewScore(25000)),
				},
				kyotaku: FirstPtoV(sc.NewScore(0)),
			},
			want: jc.Nancha,
		},
		{
			name: "PechaとShachaが同点で1,2位",
			fields: fields{
				scoreBoardId: ScoreBoardId(uuid.New()),
				scores: map[jc.Jicha]sc.Score{
					jc.Toncha: FirstPtoV(sc.NewScore(20000)),
					jc.Nancha: FirstPtoV(sc.NewScore(20000)),
					jc.Shacha: FirstPtoV(sc.NewScore(30000)),
					jc.Pecha:  FirstPtoV(sc.NewScore(30000)),
				},
				kyotaku: FirstPtoV(sc.NewScore(0)),
			},
			want: jc.Shacha,
		},
		{
			name: "Tonchaが一位",
			fields: fields{
				scoreBoardId: ScoreBoardId(uuid.New()),
				scores: map[jc.Jicha]sc.Score{
					jc.Toncha: FirstPtoV(sc.NewScore(31000)),
					jc.Nancha: FirstPtoV(sc.NewScore(23000)),
					jc.Shacha: FirstPtoV(sc.NewScore(23000)),
					jc.Pecha:  FirstPtoV(sc.NewScore(23000)),
				},
				kyotaku: FirstPtoV(sc.NewScore(0)),
			},
			want: jc.Toncha,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scoreBoard := &ScoreBoard{
				scoreBoardId: tt.fields.scoreBoardId,
				scores:       tt.fields.scores,
				kyotaku:      tt.fields.kyotaku,
			}
			if got := scoreBoard.No1AtOlast(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScoreBoard.No1AtOlast() = %v, want %v", got, tt.want)
			}
		})
	}
}
