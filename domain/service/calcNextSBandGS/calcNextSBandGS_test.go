package calcnxsbgs

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	bkh "github.com/landcelita/mahjong-management-bot/domain/model/baKyokuHonba"
	gs "github.com/landcelita/mahjong-management-bot/domain/model/gameStatus"
	hf "github.com/landcelita/mahjong-management-bot/domain/model/hanFu"
	"github.com/landcelita/mahjong-management-bot/domain/model/jicha"
	kr "github.com/landcelita/mahjong-management-bot/domain/model/kyokuResult"
	pid "github.com/landcelita/mahjong-management-bot/domain/model/playerId"
	"github.com/landcelita/mahjong-management-bot/domain/model/score"
	sb "github.com/landcelita/mahjong-management-bot/domain/model/scoreBoard"
	toh "github.com/landcelita/mahjong-management-bot/domain/model/tonpuOrHanchan"
	. "github.com/landcelita/mahjong-management-bot/testutil"
)

var (
	toncha  = jicha.Toncha
	nancha  = jicha.Nancha
	shacha  = jicha.Shacha
	pecha   = jicha.Pecha
)

func TestCalcNextSBandGS(t *testing.T) {
	id := uuid.New()
	id2 := uuid.New()
	type args struct {
		gameStatus  *gs.GameStatus
		scoreBoard  *sb.ScoreBoard
		kyokuResult *kr.KyokuResult
	}
	tests := []struct {
		name               string
		args               args
		wantNextScoreBoard *sb.ScoreBoard
		wantNextGameStatus *gs.GameStatus
		wantErr            bool
	}{
		{
			name: "Ton,1,2 riichierなし nancha Han1,Fu40 tsumo",
			args: args{
				gameStatus: generate_GS(id, bkh.Ton, 1, 2, toh.Hanchan, [4]string{"A", "B", "C", "D"}, true),
				scoreBoard: generate_SB(id, [4]int{25000, 25000, 25000, 25000}, 0),
				kyokuResult: generate_KR(id2, id, bkh.Ton, 1, 2,
					map[jicha.Jicha]struct{}{}, nil, nil, &nancha, nil, hf.Han1, hf.Fu40),
			},
			wantNextScoreBoard: generate_SB(id, [4]int{27100, 24400, 24400, 24100}, 0),
			wantNextGameStatus: generate_GS(id, bkh.Ton, 2, 0, toh.Hanchan, [4]string{"B", "C", "D", "A"}, true),
			wantErr: false,
		},
		{
			name: "Nan,4,1(Olast) riichierがtoncha,nancha (win)nancha(lose)pecha Han10,Fu2 ron",
			args: args{
				gameStatus: generate_GS(id, bkh.Nan, 4, 1, toh.Hanchan, [4]string{"A", "B", "C", "D"}, true),
				scoreBoard: generate_SB(id, [4]int{30000, 20000, 20000, 20000}, 10000),
				kyokuResult: generate_KR(id2, id, bkh.Nan, 4, 1,
					map[jicha.Jicha]struct{}{toncha: {}, nancha: {}}, &nancha, &pecha,
					nil, nil, hf.Han10, hf.FuUndefined),
			},
			wantNextScoreBoard: generate_SB(id, [4]int{29000, 47300, 20000, 3700}, 0),
			wantNextGameStatus: generate_GS(id, bkh.Nan, 4, 1, toh.Hanchan, [4]string{"A", "B", "C", "D"}, false),
			wantErr: false,
		},
		{
			name: "Ton,4,1(Olast) riichierがpecha (win)toncha(lose)nancha Han2,Fu25 ron",
			args: args{
				gameStatus: generate_GS(id, bkh.Ton, 4, 1, toh.Tonpu, [4]string{"A", "B", "C", "D"}, true),
				scoreBoard: generate_SB(id, [4]int{20000, 40000, 18000, 20000}, 2000),
				kyokuResult: generate_KR(id2, id, bkh.Ton, 4, 1,
					map[jicha.Jicha]struct{}{pecha: {}}, &toncha, &nancha,
					nil, nil, hf.Han2, hf.Fu25),
			},
			wantNextScoreBoard: generate_SB(id, [4]int{25700, 37300, 18000, 19000}, 0),
			wantNextGameStatus: generate_GS(id, bkh.Ton, 4, 2, toh.Tonpu, [4]string{"A", "B", "C", "D"}, true),
			wantErr: false,
		},
		{
			name: "Ton,4,1(Olast) riichier: toncha, tenpaier: toncha",
			args: args{
				gameStatus: generate_GS(id, bkh.Ton, 4, 1, toh.Tonpu, [4]string{"A", "B", "C", "D"}, true),
				scoreBoard: generate_SB(id, [4]int{20000, 40000, 18000, 20000}, 2000),
				kyokuResult: generate_KR(id2, id, bkh.Ton, 4, 1,
					map[jicha.Jicha]struct{}{toncha: {}}, nil, nil,
					nil, &map[jicha.Jicha]struct{}{toncha: {}}, hf.Han(1000), hf.FuUndefined),
			},
			wantNextScoreBoard: generate_SB(id, [4]int{22000, 39000, 17000, 19000}, 3000),
			wantNextGameStatus: generate_GS(id, bkh.Ton, 4, 2, toh.Tonpu, [4]string{"A", "B", "C", "D"}, true),
			wantErr: false,
		},
		{
			name: "Nan,1,0 riichier: toncha, tenpaier: toncha",
			args: args{
				gameStatus: generate_GS(id, bkh.Nan, 1, 0, toh.Hanchan, [4]string{"A", "B", "C", "D"}, true),
				scoreBoard: generate_SB(id, [4]int{0, 0, 0, 0}, 100000),
				kyokuResult: generate_KR(id2, id, bkh.Nan, 1, 0,
					map[jicha.Jicha]struct{}{toncha: {}}, nil, nil,
					nil, &map[jicha.Jicha]struct{}{toncha: {}}, hf.Han(1000), hf.FuUndefined),
			},
			wantNextScoreBoard: generate_SB(id, [4]int{2000, -1000, -1000, -1000}, 101000),
			wantNextGameStatus: generate_GS(id, bkh.Nan, 1, 0, toh.Hanchan, [4]string{"A", "B", "C", "D"}, false),
			wantErr: false,
		},
		{
			name: "Ton,1,0 ryukyoku riichier: Nancha,Shacha tenpaier: 全員",
			args: args{
				gameStatus: generate_GS(id, bkh.Ton, 1, 0, toh.Tonpu, [4]string{"A", "B", "C", "D"}, true),
				scoreBoard: generate_SB(id, [4]int{25000, 25000, 25000, 25000}, 0),
				kyokuResult: generate_KR(id2, id, bkh.Ton, 1, 0,
					map[jicha.Jicha]struct{}{nancha: {}, shacha: {}}, nil, nil,
					nil, &map[jicha.Jicha]struct{}{toncha: {}, nancha: {}, shacha: {}, pecha: {}}, hf.Han(1000), hf.FuUndefined),
			},
			wantNextScoreBoard: generate_SB(id, [4]int{25000, 24000, 24000, 25000}, 2000),
			wantNextGameStatus: generate_GS(id, bkh.Ton, 1, 1, toh.Tonpu, [4]string{"A", "B", "C", "D"}, true),
			wantErr: false,
		},
		{
			name: "Ton,1,0 ryukyoku riichier: Nancha,Shacha,Pecha tenpaier: Nancha,Shacha,Pecha",
			args: args{
				gameStatus: generate_GS(id, bkh.Ton, 1, 0, toh.Tonpu, [4]string{"A", "B", "C", "D"}, true),
				scoreBoard: generate_SB(id, [4]int{25000, 25000, 25000, 25000}, 0),
				kyokuResult: generate_KR(id2, id, bkh.Ton, 1, 0,
					map[jicha.Jicha]struct{}{nancha: {}, shacha: {}, pecha: {}}, nil, nil,
					nil, &map[jicha.Jicha]struct{}{nancha: {}, shacha: {}, pecha: {}}, hf.Han(1000), hf.FuUndefined),
			},
			wantNextScoreBoard: generate_SB(id, [4]int{25000, 25000, 25000, 22000}, 3000),
			wantNextGameStatus: generate_GS(id, bkh.Ton, 2, 0, toh.Tonpu, [4]string{"B", "C", "D", "A"}, true),
			wantErr: false,
		},
		{
			name: "Ton,4,3(Olast) tsumo riichier: Nancha,Shacha,Pecha winner: Toncha",
			args: args{
				gameStatus: generate_GS(id, bkh.Ton, 4, 3, toh.Tonpu, [4]string{"A", "B", "C", "D"}, true),
				scoreBoard: generate_SB(id, [4]int{20000, 20000, 20000, 20000}, 20000),
				kyokuResult: generate_KR(id2, id, bkh.Ton, 4, 3,
					map[jicha.Jicha]struct{}{nancha: {}, shacha: {}, pecha: {}}, nil, nil,
					&toncha, nil, hf.Han2, hf.Fu30),
			},
			wantNextScoreBoard: generate_SB(id, [4]int{46900, 17700, 17700, 17700}, 0),
			wantNextGameStatus: generate_GS(id, bkh.Ton, 4, 3, toh.Tonpu, [4]string{"A", "B", "C", "D"}, false),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNextScoreBoard, gotNextGameStatus, err := CalcNextSBandGS(tt.args.gameStatus, tt.args.scoreBoard, tt.args.kyokuResult)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalcNextSBandGS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNextScoreBoard, tt.wantNextScoreBoard) {
				t.Errorf("CalcNextSBandGS() gotNextScoreBoard = %v, want %v", gotNextScoreBoard, tt.wantNextScoreBoard)
			}
			if !reflect.DeepEqual(gotNextGameStatus, tt.wantNextGameStatus) {
				t.Errorf("CalcNextSBandGS() gotNextGameStatus = %v, want %v", gotNextGameStatus, tt.wantNextGameStatus)
			}
		})
	}
}

func Test_calcNextScoreBoard(t *testing.T) {
	type args struct {
		scoreBoard  *sb.ScoreBoard
		kyokuResult *kr.KyokuResult
	}
	tests := []struct {
		name               string
		args               args
		wantNextScoreBoard *sb.ScoreBoard
		wantErr            bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNextScoreBoard, err := calcNextScoreBoard(tt.args.scoreBoard, tt.args.kyokuResult)
			if (err != nil) != tt.wantErr {
				t.Errorf("calcNextScoreBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNextScoreBoard, tt.wantNextScoreBoard) {
				t.Errorf("calcNextScoreBoard() = %v, want %v", gotNextScoreBoard, tt.wantNextScoreBoard)
			}
		})
	}
}

func Test_calcGameProgressType(t *testing.T) {
	type args struct {
		scoreBoard  *sb.ScoreBoard
		kyokuResult *kr.KyokuResult
		gameStatus  *gs.GameStatus
	}
	tests := []struct {
		name             string
		args             args
		wantGameProgress gameProgressType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotGameProgress := calcGameProgressType(tt.args.scoreBoard, tt.args.kyokuResult, tt.args.gameStatus); !reflect.DeepEqual(gotGameProgress, tt.wantGameProgress) {
				t.Errorf("calcGameProgressType() = %v, want %v", gotGameProgress, tt.wantGameProgress)
			}
		})
	}
}

func Test_ceil100(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ceil100(tt.args.n); got != tt.want {
				t.Errorf("ceil100() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcScoreAndKyotakuDiff(t *testing.T) {
	type args struct {
		kyokuResult *kr.KyokuResult
		scoreBoard  *sb.ScoreBoard
	}
	tests := []struct {
		name            string
		args            args
		wantScoreDiff   map[jicha.Jicha]score.Score
		wantKyotakuDiff score.Score
		wantErr         bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotScoreDiff, gotKyotakuDiff, err := calcScoreAndKyotakuDiff(tt.args.kyokuResult, tt.args.scoreBoard)
			if (err != nil) != tt.wantErr {
				t.Errorf("calcScoreAndKyotakuDiff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotScoreDiff, tt.wantScoreDiff) {
				t.Errorf("calcScoreAndKyotakuDiff() gotScoreDiff = %v, want %v", gotScoreDiff, tt.wantScoreDiff)
			}
			if !reflect.DeepEqual(gotKyotakuDiff, tt.wantKyotakuDiff) {
				t.Errorf("calcScoreAndKyotakuDiff() gotKyotakuDiff = %v, want %v", gotKyotakuDiff, tt.wantKyotakuDiff)
			}
		})
	}
}

func Test_calcScoreAndKyotakuDiffWhenTsumo(t *testing.T) {
	type args struct {
		kyokuResult *kr.KyokuResult
		scoreBoard  *sb.ScoreBoard
	}
	tests := []struct {
		name            string
		args            args
		wantScoreDiff   map[jicha.Jicha]score.Score
		wantKyotakuDiff score.Score
		wantErr         bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotScoreDiff, gotKyotakuDiff, err := calcScoreAndKyotakuDiffWhenTsumo(tt.args.kyokuResult, tt.args.scoreBoard)
			if (err != nil) != tt.wantErr {
				t.Errorf("calcScoreAndKyotakuDiffWhenTsumo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotScoreDiff, tt.wantScoreDiff) {
				t.Errorf("calcScoreAndKyotakuDiffWhenTsumo() gotScoreDiff = %v, want %v", gotScoreDiff, tt.wantScoreDiff)
			}
			if !reflect.DeepEqual(gotKyotakuDiff, tt.wantKyotakuDiff) {
				t.Errorf("calcScoreAndKyotakuDiffWhenTsumo() gotKyotakuDiff = %v, want %v", gotKyotakuDiff, tt.wantKyotakuDiff)
			}
		})
	}
}

func Test_calcScoreAndKyotakuDiffWhenRon(t *testing.T) {
	type args struct {
		kyokuResult *kr.KyokuResult
		scoreBoard  *sb.ScoreBoard
	}
	tests := []struct {
		name            string
		args            args
		wantScoreDiff   map[jicha.Jicha]score.Score
		wantKyotakuDiff score.Score
		wantErr         bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotScoreDiff, gotKyotakuDiff, err := calcScoreAndKyotakuDiffWhenRon(tt.args.kyokuResult, tt.args.scoreBoard)
			if (err != nil) != tt.wantErr {
				t.Errorf("calcScoreAndKyotakuDiffWhenRon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotScoreDiff, tt.wantScoreDiff) {
				t.Errorf("calcScoreAndKyotakuDiffWhenRon() gotScoreDiff = %v, want %v", gotScoreDiff, tt.wantScoreDiff)
			}
			if !reflect.DeepEqual(gotKyotakuDiff, tt.wantKyotakuDiff) {
				t.Errorf("calcScoreAndKyotakuDiffWhenRon() gotKyotakuDiff = %v, want %v", gotKyotakuDiff, tt.wantKyotakuDiff)
			}
		})
	}
}

func Test_calcScoreAndKyotakuDiffWhenRyukyoku(t *testing.T) {
	type args struct {
		kyokuResult *kr.KyokuResult
		scoreBoard  *sb.ScoreBoard
	}
	tests := []struct {
		name            string
		args            args
		wantScoreDiff   map[jicha.Jicha]score.Score
		wantKyotakuDiff score.Score
		wantErr         bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotScoreDiff, gotKyotakuDiff, err := calcScoreAndKyotakuDiffWhenRyukyoku(tt.args.kyokuResult, tt.args.scoreBoard)
			if (err != nil) != tt.wantErr {
				t.Errorf("calcScoreAndKyotakuDiffWhenRyukyoku() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotScoreDiff, tt.wantScoreDiff) {
				t.Errorf("calcScoreAndKyotakuDiffWhenRyukyoku() gotScoreDiff = %v, want %v", gotScoreDiff, tt.wantScoreDiff)
			}
			if !reflect.DeepEqual(gotKyotakuDiff, tt.wantKyotakuDiff) {
				t.Errorf("calcScoreAndKyotakuDiffWhenRyukyoku() gotKyotakuDiff = %v, want %v", gotKyotakuDiff, tt.wantKyotakuDiff)
			}
		})
	}
}

func Test_addRiichiAndKyotakuToDiff(t *testing.T) {
	type args struct {
		kyokuResult *kr.KyokuResult
		scoreBoard  *sb.ScoreBoard
		scoreDiff   *map[jicha.Jicha]score.Score
		winner      jicha.Jicha
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := addRiichiAndKyotakuToDiff(tt.args.kyokuResult, tt.args.scoreBoard, tt.args.scoreDiff, tt.args.winner); (err != nil) != tt.wantErr {
				t.Errorf("addRiichiAndKyotakuToDiff() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func generate_GS(
	id uuid.UUID,
	ba bkh.Ba,
	kyoku uint,
	honba uint,
	toh toh.TonpuOrHanchan,
	players [4]string,
	isActive bool,
) *gs.GameStatus {
	return FirstPtoP(gs.NewGameStatus(
		gs.GameStatusId(id),
		FirstPtoV(bkh.NewBaKyokuHonba(ba, kyoku, honba)),
		toh,
		map[jicha.Jicha]pid.PlayerId{
			toncha: pid.PlayerId(players[0]),
			nancha: pid.PlayerId(players[1]),
			shacha: pid.PlayerId(players[2]),
			pecha:  pid.PlayerId(players[3]),
		},
		isActive,
	))
}

func generate_SB(
	id uuid.UUID,
	scores [4]int,
	kyotaku int,
) *sb.ScoreBoard {
	return FirstPtoP(sb.NewScoreBoard(
		sb.ScoreBoardId(id),
		map[jicha.Jicha]score.Score{
			toncha: FirstPtoV(score.NewScore(scores[0])),
			nancha: FirstPtoV(score.NewScore(scores[1])),
			shacha: FirstPtoV(score.NewScore(scores[2])),
			pecha:  FirstPtoV(score.NewScore(scores[3])),
		},
		FirstPtoV(score.NewScore(kyotaku)),
	))
}

func generate_KR(
	krid uuid.UUID,
	gsid uuid.UUID,
	ba bkh.Ba,
	kyoku uint,
	honba uint,
	riichiers map[jicha.Jicha]struct{},
	ronWinner *jicha.Jicha,
	ronLoser *jicha.Jicha,
	tsumoWinner *jicha.Jicha,
	tenpaiers *map[jicha.Jicha]struct{},
	han hf.Han,
	fu hf.Fu,
) *kr.KyokuResult {
	var hanfu *hf.HanFu
	if han != hf.Han(1000) {
		hanfu = FirstPtoP(hf.NewHanFu(han, fu))
	} else {
		hanfu = nil
	}
	return FirstPtoP(kr.NewKyokuResult(
		kr.KyokuResultId(krid),
		gs.GameStatusId(gsid),
		FirstPtoV(bkh.NewBaKyokuHonba(ba, kyoku, honba)),
		riichiers,
		ronWinner,
		ronLoser,
		tsumoWinner,
		tenpaiers,
		hanfu,
	))
}
