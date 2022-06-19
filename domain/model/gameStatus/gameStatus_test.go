package gamestatus

import (
	"mahjong/domain/model/baKyokuHonba"
	"mahjong/domain/model/playerId"
	"mahjong/domain/model/tonpuOrHanchan"
	"mahjong/domain/model/scoreBoard"
	"testing"
	"reflect"
	"strconv"

	"github.com/google/uuid"
)

const testNum = 5

func generate_TestGameStatus() (
	GameStatusId,
	[testNum]bakyokuhonba.BaKyokuHonba,
	[testNum]tonpuorhanchan.TonpuOrHanchan,
	scoreboard.ScoreBoardId,
	[4]playerid.PlayerId,
	bool,
	) {
	
	gsId := GameStatusId(uuid.New())
	var bkhs [testNum]bakyokuhonba.BaKyokuHonba
	var torh [testNum]tonpuorhanchan.TonpuOrHanchan
	scId := scoreboard.ScoreBoardId(uuid.New())
	var pIds [4]playerid.PlayerId
	isActive := true

	bkhBa := [testNum]bakyokuhonba.Ba{
		bakyokuhonba.Nan,
		bakyokuhonba.Ton,
		bakyokuhonba.Ton,
		bakyokuhonba.Nan,
		bakyokuhonba.Ton,
	}
	bkhKyoku := [testNum]uint{4, 4, 2, 1, 4}
	bkhHonba := [testNum]uint{0, 10, 1, 0, 0}

	for i := 0; i < testNum; i++ {
		bkh, e := bakyokuhonba.NewBaKyokuHonba(bkhBa[i], bkhKyoku[i], bkhHonba[i])
		if e != nil {
			panic(e)
		}
		bkhs[i] = *bkh
	}

	torh[0] = tonpuorhanchan.Hanchan
	torh[1] = tonpuorhanchan.Tonpu
	torh[2] = tonpuorhanchan.Tonpu
	torh[3] = tonpuorhanchan.Hanchan
	torh[4] = tonpuorhanchan.Hanchan

	for i := 0; i < 4; i++ {
		pId, e := playerid.NewPlayerId("AAAAA")
		if e != nil {
			panic(e)
		}
		pIds[i] = *pId
	}

	return gsId, bkhs, torh, scId, pIds, isActive
}

func TestGameStatus_IsOlast(t *testing.T) {
	gsId, bkhs, torh, scId, pIds, isActive := generate_TestGameStatus()
	wants := [testNum]bool{true, true, false, false, false}

	type fields struct {
		gameStatusId   GameStatusId
		baKyokuHonba   bakyokuhonba.BaKyokuHonba
		tonpuOrHanchan tonpuorhanchan.TonpuOrHanchan
		scoreBoardId   scoreboard.ScoreBoardId
		playerIds      [4]playerid.PlayerId
		isActive       bool
	}
	var tests = [testNum]struct {
		name   string
		fields fields
		want   bool
	}{}

	for i := 0; i < testNum; i++ {
		tests[i] = struct {
			name   string
			fields fields
			want   bool
		}{
			name: "test" + strconv.Itoa(i),
			fields: fields{
				gameStatusId:   gsId,
				baKyokuHonba:   bkhs[i],
				tonpuOrHanchan: torh[i],
				scoreBoardId:   scId,
				playerIds:      pIds,
				isActive:       isActive,
			},
			want: wants[i],
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameStatus := GameStatus{
				gameStatusId:   tt.fields.gameStatusId,
				baKyokuHonba:   tt.fields.baKyokuHonba,
				tonpuOrHanchan: tt.fields.tonpuOrHanchan,
				scoreBoardId:   tt.fields.scoreBoardId,
				playerIds:      tt.fields.playerIds,
				isActive:       tt.fields.isActive,
			}
			if got := gameStatus.IsOlast(); got != tt.want {
				t.Errorf("GameStatus.IsOlast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameStatus_AdvanceGameBaKyoku(t *testing.T) {
	gsId, bkhs, torh, scId, pIds, isActive := generate_TestGameStatus()
	wantErrs := [testNum]bool{true, true, false, false, false}
	var wantBKHs [testNum]*bakyokuhonba.BaKyokuHonba
	wantBKHs[0], _ = bakyokuhonba.NewBaKyokuHonba(bakyokuhonba.Nan, 4, 0)
	wantBKHs[1], _ = bakyokuhonba.NewBaKyokuHonba(bakyokuhonba.Ton, 4, 10)
	wantBKHs[2], _ = bakyokuhonba.NewBaKyokuHonba(bakyokuhonba.Ton, 3, 0)
	wantBKHs[3], _ = bakyokuhonba.NewBaKyokuHonba(bakyokuhonba.Nan, 2, 0)
	wantBKHs[4], _ = bakyokuhonba.NewBaKyokuHonba(bakyokuhonba.Nan, 1, 0)

	type fields struct {
		gameStatusId   GameStatusId
		baKyokuHonba   bakyokuhonba.BaKyokuHonba
		tonpuOrHanchan tonpuorhanchan.TonpuOrHanchan
		scoreBoardId   scoreboard.ScoreBoardId
		playerIds      [4]playerid.PlayerId
		isActive       bool
	}
	tests := [testNum]struct {
		name    string
		fields  fields
		wantErr bool
	}{}
	for i := 0; i < testNum; i++ {
		tests[i] = struct {
			name    string
			fields  fields
			wantErr bool
		}{
			name: "test" + strconv.Itoa(i),
			fields: fields{
				gameStatusId:   gsId,
				baKyokuHonba:   bkhs[i],
				tonpuOrHanchan: torh[i],
				scoreBoardId:   scId,
				playerIds:      pIds,
				isActive:       isActive,
			},
			wantErr: wantErrs[i],
		}
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameStatus := &GameStatus{
				gameStatusId:   tt.fields.gameStatusId,
				baKyokuHonba:   tt.fields.baKyokuHonba,
				tonpuOrHanchan: tt.fields.tonpuOrHanchan,
				scoreBoardId:   tt.fields.scoreBoardId,
				playerIds:      tt.fields.playerIds,
				isActive:       tt.fields.isActive,
			}
			err := gameStatus.AdvanceGameBaKyoku()
			if (err != nil) != tt.wantErr {
				t.Errorf("GameStatus.AdvanceGameBaKyoku() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && !reflect.DeepEqual(gameStatus.baKyokuHonba, *wantBKHs[i]) {
				t.Errorf("baKyokuHonba of gameStatus is wrong in %s", tt.name)
			}
		})
	}
}

func TestGameStatus_AdvanceGameHonba(t *testing.T) {
	gsId, bkhs, torh, scId, pIds, isActive := generate_TestGameStatus()
	wantErrs := [testNum]bool{false, false, false, false, false}
	var wantBKHs [testNum]*bakyokuhonba.BaKyokuHonba
	wantBKHs[0], _ = bakyokuhonba.NewBaKyokuHonba(bakyokuhonba.Nan, 4, 1)
	wantBKHs[1], _ = bakyokuhonba.NewBaKyokuHonba(bakyokuhonba.Ton, 4, 11)
	wantBKHs[2], _ = bakyokuhonba.NewBaKyokuHonba(bakyokuhonba.Ton, 2, 2)
	wantBKHs[3], _ = bakyokuhonba.NewBaKyokuHonba(bakyokuhonba.Nan, 1, 1)
	wantBKHs[4], _ = bakyokuhonba.NewBaKyokuHonba(bakyokuhonba.Ton, 4, 1)

	type fields struct {
		gameStatusId   GameStatusId
		baKyokuHonba   bakyokuhonba.BaKyokuHonba
		tonpuOrHanchan tonpuorhanchan.TonpuOrHanchan
		scoreBoardId   scoreboard.ScoreBoardId
		playerIds      [4]playerid.PlayerId
		isActive       bool
	}
	tests := [testNum]struct {
		name    string
		fields  fields
		wantErr bool
	}{}
	for i := 0; i < testNum; i++ {
		tests[i] = struct {
			name    string
			fields  fields
			wantErr bool
		}{
			name: "test" + strconv.Itoa(i),
			fields: fields{
				gameStatusId:   gsId,
				baKyokuHonba:   bkhs[i],
				tonpuOrHanchan: torh[i],
				scoreBoardId:   scId,
				playerIds:      pIds,
				isActive:       isActive,
			},
			wantErr: wantErrs[i],
		}
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameStatus := &GameStatus{
				gameStatusId:   tt.fields.gameStatusId,
				baKyokuHonba:   tt.fields.baKyokuHonba,
				tonpuOrHanchan: tt.fields.tonpuOrHanchan,
				scoreBoardId:   tt.fields.scoreBoardId,
				playerIds:      tt.fields.playerIds,
				isActive:       tt.fields.isActive,
			}
			err := gameStatus.AdvanceGameHonba()
			if (err != nil) != tt.wantErr {
				t.Errorf("GameStatus.AdvanceGameHonba() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && !reflect.DeepEqual(gameStatus.baKyokuHonba, *wantBKHs[i]) {
				t.Errorf("baKyokuHonba of gameStatus is wrong in %s", tt.name)
			}
		})
	}
}

func TestNewGameStatus(t *testing.T) {
	gsId, bkhs, _, scId, pIds, isActive := generate_TestGameStatus()
	torh := [testNum]tonpuorhanchan.TonpuOrHanchan{
		tonpuorhanchan.Hanchan,
		tonpuorhanchan.Tonpu,
		tonpuorhanchan.Tonpu,
		tonpuorhanchan.Tonpu,
		tonpuorhanchan.Hanchan,
	}
	wantErrs := [testNum]bool {
		false, false, false, true, false,
	}

	type args struct {
		gameStatusId   GameStatusId
		baKyokuHonba   bakyokuhonba.BaKyokuHonba
		tonpuOrHanchan tonpuorhanchan.TonpuOrHanchan
		scoreBoardId   scoreboard.ScoreBoardId
		playerIds      [4]playerid.PlayerId
		isActive       bool
	}
	tests := [testNum]struct {
		name    string
		args    args
		wantErr bool
	}{}

	for i := 0; i < testNum; i++ {
		tests[i] = struct {
			name    string
			args	args
			wantErr bool
		}{
			name: "test" + strconv.Itoa(i),
			args: args{
				gameStatusId:   gsId,
				baKyokuHonba:   bkhs[i],
				tonpuOrHanchan: torh[i],
				scoreBoardId:   scId,
				playerIds:      pIds,
				isActive:       isActive,
			},
			wantErr: wantErrs[i],
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewGameStatus(tt.args.gameStatusId, tt.args.baKyokuHonba, tt.args.tonpuOrHanchan, tt.args.scoreBoardId, tt.args.playerIds, tt.args.isActive)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGameStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}