package gamestatus

import (
	"reflect"
	"strconv"
	"testing"

	bakyokuhonba "github.com/landcelita/mahjong-management-bot/domain/model/baKyokuHonba"
	"github.com/landcelita/mahjong-management-bot/domain/model/jicha"
	playerid "github.com/landcelita/mahjong-management-bot/domain/model/playerId"
	tonpuorhanchan "github.com/landcelita/mahjong-management-bot/domain/model/tonpuOrHanchan"
	. "github.com/landcelita/mahjong-management-bot/testutil"

	"github.com/google/uuid"
)

const testNum = 5

func generate_TestGameStatus() (
	GameStatusId,
	[testNum]bakyokuhonba.BaKyokuHonba,
	[testNum]tonpuorhanchan.TonpuOrHanchan,
	map[jicha.Jicha]playerid.PlayerId,
	bool,
) {

	gsId := GameStatusId(uuid.New())
	bkhs := [testNum]bakyokuhonba.BaKyokuHonba{
		FirstPtoV(bakyokuhonba.NewBaKyokuHonba(bakyokuhonba.Nan, 4, 0)),
		FirstPtoV(bakyokuhonba.NewBaKyokuHonba(bakyokuhonba.Ton, 4, 10)),
		FirstPtoV(bakyokuhonba.NewBaKyokuHonba(bakyokuhonba.Ton, 2, 1)),
		FirstPtoV(bakyokuhonba.NewBaKyokuHonba(bakyokuhonba.Nan, 1, 0)),
		FirstPtoV(bakyokuhonba.NewBaKyokuHonba(bakyokuhonba.Ton, 4, 0)),
	}
	torh := [testNum]tonpuorhanchan.TonpuOrHanchan{
		tonpuorhanchan.Hanchan,
		tonpuorhanchan.Tonpu,
		tonpuorhanchan.Tonpu,
		tonpuorhanchan.Hanchan,
		tonpuorhanchan.Hanchan,
	}
	pIds := make(map[jicha.Jicha]playerid.PlayerId)
	pIds[jicha.Toncha] = FirstPtoV(playerid.NewPlayerId("PLAYER1"))
	pIds[jicha.Nancha] = FirstPtoV(playerid.NewPlayerId("PLAYER2"))
	pIds[jicha.Shacha] = FirstPtoV(playerid.NewPlayerId("PLAYER3"))
	pIds[jicha.Pecha] = FirstPtoV(playerid.NewPlayerId("PLAYER4"))
	isActive := true

	return gsId, bkhs, torh, pIds, isActive
}

func TestGameStatus_IsOlast(t *testing.T) {
	gsId, bkhs, torh, pIds, isActive := generate_TestGameStatus()
	wants := [testNum]bool{true, true, false, false, false}

	type fields struct {
		gameStatusId   GameStatusId
		baKyokuHonba   bakyokuhonba.BaKyokuHonba
		tonpuOrHanchan tonpuorhanchan.TonpuOrHanchan
		playerIds      map[jicha.Jicha]playerid.PlayerId
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
	gsId, bkhs, torh, pIds, isActive := generate_TestGameStatus()
	wantErrs := [testNum]bool{true, true, false, false, false}
	var wantBKHs [testNum]*bakyokuhonba.BaKyokuHonba
	var wantPlayerIds [testNum]map[jicha.Jicha]playerid.PlayerId
	wantBKHs[0], _ = bakyokuhonba.NewBaKyokuHonba(bakyokuhonba.Nan, 4, 0)
	wantBKHs[1], _ = bakyokuhonba.NewBaKyokuHonba(bakyokuhonba.Ton, 4, 10)
	wantBKHs[2], _ = bakyokuhonba.NewBaKyokuHonba(bakyokuhonba.Ton, 3, 0)
	wantBKHs[3], _ = bakyokuhonba.NewBaKyokuHonba(bakyokuhonba.Nan, 2, 0)
	wantBKHs[4], _ = bakyokuhonba.NewBaKyokuHonba(bakyokuhonba.Nan, 1, 0)

	for i := 0; i < testNum; i++ {
		wantPlayerIds[i] = make(map[jicha.Jicha]playerid.PlayerId)
		wantPlayerIds[i][jicha.Toncha] = FirstPtoV(playerid.NewPlayerId("PLAYER4"))
		wantPlayerIds[i][jicha.Nancha] = FirstPtoV(playerid.NewPlayerId("PLAYER1"))
		wantPlayerIds[i][jicha.Shacha] = FirstPtoV(playerid.NewPlayerId("PLAYER2"))
		wantPlayerIds[i][jicha.Pecha] = FirstPtoV(playerid.NewPlayerId("PLAYER3"))
	}

	type fields struct {
		gameStatusId   GameStatusId
		baKyokuHonba   bakyokuhonba.BaKyokuHonba
		tonpuOrHanchan tonpuorhanchan.TonpuOrHanchan
		playerIds      map[jicha.Jicha]playerid.PlayerId
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
	gsId, bkhs, torh, pIds, isActive := generate_TestGameStatus()
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
		playerIds      map[jicha.Jicha]playerid.PlayerId
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
	gsId, bkhs, _, pIds, isActive := generate_TestGameStatus()
	torh := [testNum]tonpuorhanchan.TonpuOrHanchan{
		tonpuorhanchan.Hanchan,
		tonpuorhanchan.Tonpu,
		tonpuorhanchan.Tonpu,
		tonpuorhanchan.Tonpu,
		tonpuorhanchan.Hanchan,
	}
	wantErrs := [testNum]bool{
		false, false, false, true, false,
	}

	type args struct {
		gameStatusId   GameStatusId
		baKyokuHonba   bakyokuhonba.BaKyokuHonba
		tonpuOrHanchan tonpuorhanchan.TonpuOrHanchan
		playerIds      map[jicha.Jicha]playerid.PlayerId
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
			args    args
			wantErr bool
		}{
			name: "test" + strconv.Itoa(i),
			args: args{
				gameStatusId:   gsId,
				baKyokuHonba:   bkhs[i],
				tonpuOrHanchan: torh[i],
				playerIds:      pIds,
				isActive:       isActive,
			},
			wantErr: wantErrs[i],
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := newGameStatus(tt.args.gameStatusId, tt.args.baKyokuHonba, tt.args.tonpuOrHanchan, tt.args.playerIds, tt.args.isActive)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGameStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
