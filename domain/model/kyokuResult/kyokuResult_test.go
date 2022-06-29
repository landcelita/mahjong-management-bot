package kyokuresult

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	bkh "github.com/landcelita/mahjong-management-bot/domain/model/baKyokuHonba"
	gs "github.com/landcelita/mahjong-management-bot/domain/model/gameStatus"
	hf "github.com/landcelita/mahjong-management-bot/domain/model/hanFu"
	"github.com/landcelita/mahjong-management-bot/domain/model/jicha"
	. "github.com/landcelita/mahjong-management-bot/testutil"
)

var (
	toncha  = jicha.Toncha
	nancha  = jicha.Nancha
	shacha  = jicha.Shacha
	pecha   = jicha.Pecha
	toncha2 = jicha.Toncha
)

func TestNewKyokuResult(t *testing.T) {
	type args struct {
		kyokuResultId KyokuResultId
		gameStatusId  gs.GameStatusId
		baKyokuHonba  bkh.BaKyokuHonba
		riichiers     map[jicha.Jicha]struct{}
		ronWinner     *jicha.Jicha
		ronLoser      *jicha.Jicha
		tsumoWinner   *jicha.Jicha
		tenpaiers     *map[jicha.Jicha]struct{}
		hanFu         *hf.HanFu
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "正常系 ron",
			args: args{
				kyokuResultId: KyokuResultId(uuid.New()),
				gameStatusId:  gs.GameStatusId(uuid.New()),
				baKyokuHonba: FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Ton, 2, 3,
				)),
				riichiers: map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {}, pecha: {},
				},
				ronWinner:   &toncha,
				ronLoser:    &nancha,
				tsumoWinner: nil,
				tenpaiers:   nil,
				hanFu:       FirstPtoP(hf.NewHanFu(hf.Han1, hf.Fu40)),
			},
			wantErr: false,
		},
		{
			name: "正常系 tsumo",
			args: args{
				kyokuResultId: KyokuResultId(uuid.New()),
				gameStatusId:  gs.GameStatusId(uuid.New()),
				baKyokuHonba: FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 4, 1,
				)),
				riichiers:   map[jicha.Jicha]struct{}{},
				ronWinner:   nil,
				ronLoser:    nil,
				tsumoWinner: &nancha,
				tenpaiers:   nil,
				hanFu:       FirstPtoP(hf.NewHanFu(hf.Han1, hf.Fu40)),
			},
			wantErr: false,
		},
		{
			name: "正常系 ryukyoku",
			args: args{
				kyokuResultId: KyokuResultId(uuid.New()),
				gameStatusId:  gs.GameStatusId(uuid.New()),
				baKyokuHonba: FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 1, 1,
				)),
				riichiers: map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {},
				},
				ronWinner:   nil,
				ronLoser:    nil,
				tsumoWinner: nil,
				tenpaiers: &map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {},
				},
				hanFu: nil,
			},
			wantErr: false,
		},
		{
			name: "異常系 riichiersの値が不正",
			args: args{
				kyokuResultId: KyokuResultId(uuid.New()),
				gameStatusId:  gs.GameStatusId(uuid.New()),
				baKyokuHonba: FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 1, 1,
				)),
				riichiers: map[jicha.Jicha]struct{}{
					jicha.Jicha("Ton"): {},
				},
				ronWinner:   nil,
				ronLoser:    nil,
				tsumoWinner: nil,
				tenpaiers: &map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {},
				},
				hanFu: nil,
			},
			wantErr: true,
		},
		{
			name: "異常系 tenpaiersの値が不正",
			args: args{
				kyokuResultId: KyokuResultId(uuid.New()),
				gameStatusId:  gs.GameStatusId(uuid.New()),
				baKyokuHonba: FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 1, 1,
				)),
				riichiers: map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {},
				},
				ronWinner:   nil,
				ronLoser:    nil,
				tsumoWinner: nil,
				tenpaiers: &map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, jicha.Jicha("Sha"): {},
				},
				hanFu: nil,
			},
			wantErr: true,
		},
		{
			name: "異常系 ronWinnerのみに値が入っている",
			args: args{
				kyokuResultId: KyokuResultId(uuid.New()),
				gameStatusId:  gs.GameStatusId(uuid.New()),
				baKyokuHonba: FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Ton, 2, 3,
				)),
				riichiers: map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {}, pecha: {},
				},
				ronWinner:   &toncha,
				ronLoser:    nil,
				tsumoWinner: nil,
				tenpaiers:   nil,
				hanFu:       FirstPtoP(hf.NewHanFu(hf.Han1, hf.Fu40)),
			},
			wantErr: true,
		},
		{
			name: "異常系 ronとtenpaiの両方に値が入っている",
			args: args{
				kyokuResultId: KyokuResultId(uuid.New()),
				gameStatusId:  gs.GameStatusId(uuid.New()),
				baKyokuHonba: FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Ton, 2, 3,
				)),
				riichiers: map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {},
				},
				ronWinner:   &toncha,
				ronLoser:    &nancha,
				tsumoWinner: nil,
				tenpaiers: &map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {},
				},
				hanFu: FirstPtoP(hf.NewHanFu(hf.Han1, hf.Fu40)),
			},
			wantErr: true,
		},
		{
			name: "異常系 ronWinnerとronLoserが同じ",
			args: args{
				kyokuResultId: KyokuResultId(uuid.New()),
				gameStatusId:  gs.GameStatusId(uuid.New()),
				baKyokuHonba: FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Ton, 2, 3,
				)),
				riichiers: map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {}, pecha: {},
				},
				ronWinner:   &toncha,
				ronLoser:    &toncha2,
				tsumoWinner: nil,
				tenpaiers:   nil,
				hanFu:       FirstPtoP(hf.NewHanFu(hf.Han1, hf.Fu40)),
			},
			wantErr: true,
		},
		{
			name: "異常系 ryukyokuしているのにhanfuが入っている",
			args: args{
				kyokuResultId: KyokuResultId(uuid.New()),
				gameStatusId:  gs.GameStatusId(uuid.New()),
				baKyokuHonba: FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 1, 1,
				)),
				riichiers: map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {},
				},
				ronWinner:   nil,
				ronLoser:    nil,
				tsumoWinner: nil,
				tenpaiers: &map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {},
				},
				hanFu: FirstPtoP(hf.NewHanFu(hf.Han1, hf.Fu40)),
			},
			wantErr: true,
		},
		{
			name: "異常系 ronなのにhanfuがない",
			args: args{
				kyokuResultId: KyokuResultId(uuid.New()),
				gameStatusId:  gs.GameStatusId(uuid.New()),
				baKyokuHonba: FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Ton, 2, 3,
				)),
				riichiers: map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {}, pecha: {},
				},
				ronWinner:   &toncha,
				ronLoser:    &nancha,
				tsumoWinner: nil,
				tenpaiers:   nil,
				hanFu:       nil,
			},
			wantErr: true,
		},
		{
			name: "異常系 tenpaiersとriichiersの包含関係がおかしい",
			args: args{
				kyokuResultId: KyokuResultId(uuid.New()),
				gameStatusId:  gs.GameStatusId(uuid.New()),
				baKyokuHonba: FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 1, 1,
				)),
				riichiers: map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, pecha: {},
				},
				ronWinner:   nil,
				ronLoser:    nil,
				tsumoWinner: nil,
				tenpaiers: &map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {},
				},
				hanFu: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewKyokuResult(tt.args.kyokuResultId, tt.args.gameStatusId, tt.args.baKyokuHonba, tt.args.riichiers, tt.args.ronWinner, tt.args.ronLoser, tt.args.tsumoWinner, tt.args.tenpaiers, tt.args.hanFu)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewKyokuResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestKyokuResult_GetKyokuEndType(t *testing.T) {
	type fields struct {
		kyokuResultId KyokuResultId
		gameStatusId  gs.GameStatusId
		baKyokuHonba  bkh.BaKyokuHonba
		riichiers     map[jicha.Jicha]struct{}
		ronWinner     *jicha.Jicha
		ronLoser      *jicha.Jicha
		tsumoWinner   *jicha.Jicha
		tenpaiers     *map[jicha.Jicha]struct{}
		hanFu         *hf.HanFu
	}
	tests := []struct {
		name   string
		fields fields
		want   KyokuEndType
	}{
		{
			name: "Ron Type",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Ton, 2, 3,
				)),
				map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {}, pecha: {},
				},
				&toncha,
				&nancha,
				nil,
				nil,
				FirstPtoP(hf.NewHanFu(hf.Han1, hf.Fu40)),
			},
			want: Ron,
		},
		{
			name: "Tsumo Type",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 4, 1,
				)),
				map[jicha.Jicha]struct{}{},
				nil,
				nil,
				&nancha,
				nil,
				FirstPtoP(hf.NewHanFu(hf.Han1, hf.Fu40)),
			},
			want: Tsumo,
		},
		{
			name: "Ryukyoku Type",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 1, 1,
				)),
				map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {},
				},
				nil,
				nil,
				nil,
				&map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {},
				},
				nil,
			},
			want: Ryukyoku,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kyokuResult := &KyokuResult{
				kyokuResultId: tt.fields.kyokuResultId,
				gameStatusId:  tt.fields.gameStatusId,
				baKyokuHonba:  tt.fields.baKyokuHonba,
				riichiers:     tt.fields.riichiers,
				ronWinner:     tt.fields.ronWinner,
				ronLoser:      tt.fields.ronLoser,
				tsumoWinner:   tt.fields.tsumoWinner,
				tenpaiers:     tt.fields.tenpaiers,
				hanFu:         tt.fields.hanFu,
			}
			if got := kyokuResult.GetKyokuEndType(); got != tt.want {
				t.Errorf("KyokuResult.GetKyokuEndType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKyokuResult_CalcBaseScore(t *testing.T) {
	type fields struct {
		kyokuResultId KyokuResultId
		gameStatusId  gs.GameStatusId
		baKyokuHonba  bkh.BaKyokuHonba
		riichiers     map[jicha.Jicha]struct{}
		ronWinner     *jicha.Jicha
		ronLoser      *jicha.Jicha
		tsumoWinner   *jicha.Jicha
		tenpaiers     *map[jicha.Jicha]struct{}
		hanFu         *hf.HanFu
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint
		wantErr bool
	}{
		{
			name: "正常系 TonchaがRon, Han1Fu40",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Ton, 2, 3,
				)),
				map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {}, pecha: {},
				},
				&toncha,
				&nancha,
				nil,
				nil,
				FirstPtoP(hf.NewHanFu(hf.Han1, hf.Fu40)),
			},
			want:    640,
			wantErr: false,
		},
		{
			name: "正常系 Toncha以外がTsumo 10翻",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 4, 1,
				)),
				map[jicha.Jicha]struct{}{},
				nil,
				nil,
				&nancha,
				nil,
				FirstPtoP(hf.NewHanFu(hf.Han10, hf.FuUndefined)),
			},
			want:    4000,
			wantErr: false,
		},
		{
			name: "正常系 TonchaがTsumo 2翻30符",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 4, 1,
				)),
				map[jicha.Jicha]struct{}{},
				nil,
				nil,
				&toncha,
				nil,
				FirstPtoP(hf.NewHanFu(hf.Han2, hf.Fu30)),
			},
			want:    960,
			wantErr: false,
		},
		{
			name: "異常系 ryukyoku",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 1, 1,
				)),
				map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {},
				},
				nil,
				nil,
				nil,
				&map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {},
				},
				nil,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kyokuResult := &KyokuResult{
				kyokuResultId: tt.fields.kyokuResultId,
				gameStatusId:  tt.fields.gameStatusId,
				baKyokuHonba:  tt.fields.baKyokuHonba,
				riichiers:     tt.fields.riichiers,
				ronWinner:     tt.fields.ronWinner,
				ronLoser:      tt.fields.ronLoser,
				tsumoWinner:   tt.fields.tsumoWinner,
				tenpaiers:     tt.fields.tenpaiers,
				hanFu:         tt.fields.hanFu,
			}
			got, err := kyokuResult.CalcBaseScore()
			if (err != nil) != tt.wantErr {
				t.Errorf("KyokuResult.CalcBaseScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("KyokuResult.CalcBaseScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKyokuResult_isTonchaRonWinner(t *testing.T) {
	type fields struct {
		kyokuResultId KyokuResultId
		gameStatusId  gs.GameStatusId
		baKyokuHonba  bkh.BaKyokuHonba
		riichiers     map[jicha.Jicha]struct{}
		ronWinner     *jicha.Jicha
		ronLoser      *jicha.Jicha
		tsumoWinner   *jicha.Jicha
		tenpaiers     *map[jicha.Jicha]struct{}
		hanFu         *hf.HanFu
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "TonchaがRon",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Ton, 2, 3,
				)),
				map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {}, pecha: {},
				},
				&toncha,
				&nancha,
				nil,
				nil,
				FirstPtoP(hf.NewHanFu(hf.Han1, hf.Fu40)),
			},
			want: true,
		},
		{
			name: "NanchaがRon",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Ton, 2, 3,
				)),
				map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {}, pecha: {},
				},
				&nancha,
				&toncha,
				nil,
				nil,
				FirstPtoP(hf.NewHanFu(hf.Han1, hf.Fu40)),
			},
			want: false,
		},
		{
			name: "TonchaがTsumo",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Ton, 2, 3,
				)),
				map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {}, pecha: {},
				},
				nil,
				nil,
				&toncha,
				nil,
				FirstPtoP(hf.NewHanFu(hf.Han1, hf.Fu40)),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kyokuResult := &KyokuResult{
				kyokuResultId: tt.fields.kyokuResultId,
				gameStatusId:  tt.fields.gameStatusId,
				baKyokuHonba:  tt.fields.baKyokuHonba,
				riichiers:     tt.fields.riichiers,
				ronWinner:     tt.fields.ronWinner,
				ronLoser:      tt.fields.ronLoser,
				tsumoWinner:   tt.fields.tsumoWinner,
				tenpaiers:     tt.fields.tenpaiers,
				hanFu:         tt.fields.hanFu,
			}
			if got := kyokuResult.isTonchaRonWinner(); got != tt.want {
				t.Errorf("KyokuResult.isTonchaRonWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKyokuResult_isTonchaTsumo(t *testing.T) {
	type fields struct {
		kyokuResultId KyokuResultId
		gameStatusId  gs.GameStatusId
		baKyokuHonba  bkh.BaKyokuHonba
		riichiers     map[jicha.Jicha]struct{}
		ronWinner     *jicha.Jicha
		ronLoser      *jicha.Jicha
		tsumoWinner   *jicha.Jicha
		tenpaiers     *map[jicha.Jicha]struct{}
		hanFu         *hf.HanFu
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "TonchaがTsumo",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 4, 1,
				)),
				map[jicha.Jicha]struct{}{},
				nil,
				nil,
				&toncha,
				nil,
				FirstPtoP(hf.NewHanFu(hf.Han10, hf.FuUndefined)),
			},
			want: true,
		},
		{
			name: "nanchaがtsumo",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 4, 1,
				)),
				map[jicha.Jicha]struct{}{},
				nil,
				nil,
				&nancha,
				nil,
				FirstPtoP(hf.NewHanFu(hf.Han10, hf.FuUndefined)),
			},
			want: false,
		},
		{
			name: "tonchaがtenpaiでryukyoku",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 4, 1,
				)),
				map[jicha.Jicha]struct{}{},
				nil,
				nil,
				nil,
				&map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {},
				},
				FirstPtoP(hf.NewHanFu(hf.Han10, hf.FuUndefined)),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kyokuResult := &KyokuResult{
				kyokuResultId: tt.fields.kyokuResultId,
				gameStatusId:  tt.fields.gameStatusId,
				baKyokuHonba:  tt.fields.baKyokuHonba,
				riichiers:     tt.fields.riichiers,
				ronWinner:     tt.fields.ronWinner,
				ronLoser:      tt.fields.ronLoser,
				tsumoWinner:   tt.fields.tsumoWinner,
				tenpaiers:     tt.fields.tenpaiers,
				hanFu:         tt.fields.hanFu,
			}
			if got := kyokuResult.isTonchaTsumo(); got != tt.want {
				t.Errorf("KyokuResult.isTonchaTsumo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKyokuResult_isTonchaTenpai(t *testing.T) {
	type fields struct {
		kyokuResultId KyokuResultId
		gameStatusId  gs.GameStatusId
		baKyokuHonba  bkh.BaKyokuHonba
		riichiers     map[jicha.Jicha]struct{}
		ronWinner     *jicha.Jicha
		ronLoser      *jicha.Jicha
		tsumoWinner   *jicha.Jicha
		tenpaiers     *map[jicha.Jicha]struct{}
		hanFu         *hf.HanFu
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "tonchaがtenpaiでryukyoku",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 4, 1,
				)),
				map[jicha.Jicha]struct{}{},
				nil,
				nil,
				nil,
				&map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {},
				},
				FirstPtoP(hf.NewHanFu(hf.Han10, hf.FuUndefined)),
			},
			want: true,
		},
		{
			name: "tonchaがtenpaiせずryukyoku",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 4, 1,
				)),
				map[jicha.Jicha]struct{}{},
				nil,
				nil,
				nil,
				&map[jicha.Jicha]struct{}{},
				FirstPtoP(hf.NewHanFu(hf.Han10, hf.FuUndefined)),
			},
			want: false,
		},
		{
			name: "TonchaがRon",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Ton, 2, 3,
				)),
				map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {}, pecha: {},
				},
				&toncha,
				&nancha,
				nil,
				nil,
				FirstPtoP(hf.NewHanFu(hf.Han1, hf.Fu40)),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kyokuResult := &KyokuResult{
				kyokuResultId: tt.fields.kyokuResultId,
				gameStatusId:  tt.fields.gameStatusId,
				baKyokuHonba:  tt.fields.baKyokuHonba,
				riichiers:     tt.fields.riichiers,
				ronWinner:     tt.fields.ronWinner,
				ronLoser:      tt.fields.ronLoser,
				tsumoWinner:   tt.fields.tsumoWinner,
				tenpaiers:     tt.fields.tenpaiers,
				hanFu:         tt.fields.hanFu,
			}
			if got := kyokuResult.isTonchaTenpai(); got != tt.want {
				t.Errorf("KyokuResult.isTonchaTenpai() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKyokuResult_WhoRonWinner(t *testing.T) {
	type fields struct {
		kyokuResultId KyokuResultId
		gameStatusId  gs.GameStatusId
		baKyokuHonba  bkh.BaKyokuHonba
		riichiers     map[jicha.Jicha]struct{}
		ronWinner     *jicha.Jicha
		ronLoser      *jicha.Jicha
		tsumoWinner   *jicha.Jicha
		tenpaiers     *map[jicha.Jicha]struct{}
		hanFu         *hf.HanFu
	}
	tests := []struct {
		name    string
		fields  fields
		want    *jicha.Jicha
		wantErr bool
	}{
		{
			name: "正常系 TonchaがRon",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Ton, 2, 3,
				)),
				map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {}, pecha: {},
				},
				&toncha,
				&nancha,
				nil,
				nil,
				FirstPtoP(hf.NewHanFu(hf.Han1, hf.Fu40)),
			},
			want:    &toncha2,
			wantErr: false,
		},
		{
			name: "異常系 tonchaがtenpaiでryukyoku",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 4, 1,
				)),
				map[jicha.Jicha]struct{}{},
				nil,
				nil,
				nil,
				&map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {},
				},
				FirstPtoP(hf.NewHanFu(hf.Han10, hf.FuUndefined)),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kyokuResult := &KyokuResult{
				kyokuResultId: tt.fields.kyokuResultId,
				gameStatusId:  tt.fields.gameStatusId,
				baKyokuHonba:  tt.fields.baKyokuHonba,
				riichiers:     tt.fields.riichiers,
				ronWinner:     tt.fields.ronWinner,
				ronLoser:      tt.fields.ronLoser,
				tsumoWinner:   tt.fields.tsumoWinner,
				tenpaiers:     tt.fields.tenpaiers,
				hanFu:         tt.fields.hanFu,
			}
			got, err := kyokuResult.WhoRonWinner()
			if (err != nil) != tt.wantErr {
				t.Errorf("KyokuResult.WhoRonWinner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KyokuResult.WhoRonWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKyokuResult_WhoRonLoser(t *testing.T) {
	type fields struct {
		kyokuResultId KyokuResultId
		gameStatusId  gs.GameStatusId
		baKyokuHonba  bkh.BaKyokuHonba
		riichiers     map[jicha.Jicha]struct{}
		ronWinner     *jicha.Jicha
		ronLoser      *jicha.Jicha
		tsumoWinner   *jicha.Jicha
		tenpaiers     *map[jicha.Jicha]struct{}
		hanFu         *hf.HanFu
	}
	tests := []struct {
		name    string
		fields  fields
		want    *jicha.Jicha
		wantErr bool
	}{
		{
			name: "正常系 NanchaがRonされた",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Ton, 2, 3,
				)),
				map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {}, pecha: {},
				},
				&toncha,
				&nancha,
				nil,
				nil,
				FirstPtoP(hf.NewHanFu(hf.Han1, hf.Fu40)),
			},
			want:    &nancha,
			wantErr: false,
		},
		{
			name: "異常系 ryukyoku",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 4, 1,
				)),
				map[jicha.Jicha]struct{}{},
				nil,
				nil,
				nil,
				&map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {},
				},
				FirstPtoP(hf.NewHanFu(hf.Han10, hf.FuUndefined)),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kyokuResult := &KyokuResult{
				kyokuResultId: tt.fields.kyokuResultId,
				gameStatusId:  tt.fields.gameStatusId,
				baKyokuHonba:  tt.fields.baKyokuHonba,
				riichiers:     tt.fields.riichiers,
				ronWinner:     tt.fields.ronWinner,
				ronLoser:      tt.fields.ronLoser,
				tsumoWinner:   tt.fields.tsumoWinner,
				tenpaiers:     tt.fields.tenpaiers,
				hanFu:         tt.fields.hanFu,
			}
			got, err := kyokuResult.WhoRonLoser()
			if (err != nil) != tt.wantErr {
				t.Errorf("KyokuResult.WhoRonLoser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KyokuResult.WhoRonLoser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKyokuResult_WhoTsumo(t *testing.T) {
	type fields struct {
		kyokuResultId KyokuResultId
		gameStatusId  gs.GameStatusId
		baKyokuHonba  bkh.BaKyokuHonba
		riichiers     map[jicha.Jicha]struct{}
		ronWinner     *jicha.Jicha
		ronLoser      *jicha.Jicha
		tsumoWinner   *jicha.Jicha
		tenpaiers     *map[jicha.Jicha]struct{}
		hanFu         *hf.HanFu
	}
	tests := []struct {
		name    string
		fields  fields
		want    *jicha.Jicha
		wantErr bool
	}{
		{
			name: "正常系 nanchaがtsumo",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 4, 1,
				)),
				map[jicha.Jicha]struct{}{},
				nil,
				nil,
				&nancha,
				nil,
				FirstPtoP(hf.NewHanFu(hf.Han10, hf.FuUndefined)),
			},
			want:    &nancha,
			wantErr: false,
		},
		{
			name: "異常系 ryukyoku",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 4, 1,
				)),
				map[jicha.Jicha]struct{}{},
				nil,
				nil,
				nil,
				&map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {},
				},
				FirstPtoP(hf.NewHanFu(hf.Han10, hf.FuUndefined)),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kyokuResult := &KyokuResult{
				kyokuResultId: tt.fields.kyokuResultId,
				gameStatusId:  tt.fields.gameStatusId,
				baKyokuHonba:  tt.fields.baKyokuHonba,
				riichiers:     tt.fields.riichiers,
				ronWinner:     tt.fields.ronWinner,
				ronLoser:      tt.fields.ronLoser,
				tsumoWinner:   tt.fields.tsumoWinner,
				tenpaiers:     tt.fields.tenpaiers,
				hanFu:         tt.fields.hanFu,
			}
			got, err := kyokuResult.WhoTsumo()
			if (err != nil) != tt.wantErr {
				t.Errorf("KyokuResult.WhoTsumo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KyokuResult.WhoTsumo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKyokuResult_WhoTenpai(t *testing.T) {
	type fields struct {
		kyokuResultId KyokuResultId
		gameStatusId  gs.GameStatusId
		baKyokuHonba  bkh.BaKyokuHonba
		riichiers     map[jicha.Jicha]struct{}
		ronWinner     *jicha.Jicha
		ronLoser      *jicha.Jicha
		tsumoWinner   *jicha.Jicha
		tenpaiers     *map[jicha.Jicha]struct{}
		hanFu         *hf.HanFu
	}
	tests := []struct {
		name    string
		fields  fields
		want    *map[jicha.Jicha]struct{}
		wantErr bool
	}{
		{
			name: "正常系 toncha, pechaがtenpai",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 4, 1,
				)),
				map[jicha.Jicha]struct{}{
					toncha: {}, pecha: {},
				},
				nil,
				nil,
				nil,
				&map[jicha.Jicha]struct{}{
					toncha: {}, pecha: {},
				},
				FirstPtoP(hf.NewHanFu(hf.Han10, hf.FuUndefined)),
			},
			want: &map[jicha.Jicha]struct{}{
				toncha: {}, pecha: {},
			},
			wantErr: false,
		},
		{
			name: "異常系 Ron",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Ton, 2, 3,
				)),
				map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {}, pecha: {},
				},
				&toncha,
				&nancha,
				nil,
				nil,
				FirstPtoP(hf.NewHanFu(hf.Han1, hf.Fu40)),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kyokuResult := &KyokuResult{
				kyokuResultId: tt.fields.kyokuResultId,
				gameStatusId:  tt.fields.gameStatusId,
				baKyokuHonba:  tt.fields.baKyokuHonba,
				riichiers:     tt.fields.riichiers,
				ronWinner:     tt.fields.ronWinner,
				ronLoser:      tt.fields.ronLoser,
				tsumoWinner:   tt.fields.tsumoWinner,
				tenpaiers:     tt.fields.tenpaiers,
				hanFu:         tt.fields.hanFu,
			}
			got, err := kyokuResult.WhoTenpai()
			if (err != nil) != tt.wantErr {
				t.Errorf("KyokuResult.WhoTenpai() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KyokuResult.WhoTenpai() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKyokuResult_WhoRiichi(t *testing.T) {
	type fields struct {
		kyokuResultId KyokuResultId
		gameStatusId  gs.GameStatusId
		baKyokuHonba  bkh.BaKyokuHonba
		riichiers     map[jicha.Jicha]struct{}
		ronWinner     *jicha.Jicha
		ronLoser      *jicha.Jicha
		tsumoWinner   *jicha.Jicha
		tenpaiers     *map[jicha.Jicha]struct{}
		hanFu         *hf.HanFu
	}
	tests := []struct {
		name   string
		fields fields
		want   *map[jicha.Jicha]struct{}
	}{
		{
			name: "Toncha, Nancha, Pechaがriichi",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 4, 1,
				)),
				map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, pecha: {},
				},
				nil,
				nil,
				&nancha,
				nil,
				FirstPtoP(hf.NewHanFu(hf.Han10, hf.FuUndefined)),
			},
			want: &map[jicha.Jicha]struct{}{
				toncha2: {}, nancha: {}, pecha: {},
			},
		},
		{
			name: "riichierおらず",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 4, 1,
				)),
				map[jicha.Jicha]struct{}{},
				nil,
				nil,
				&nancha,
				nil,
				FirstPtoP(hf.NewHanFu(hf.Han10, hf.FuUndefined)),
			},
			want: &map[jicha.Jicha]struct{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kyokuResult := &KyokuResult{
				kyokuResultId: tt.fields.kyokuResultId,
				gameStatusId:  tt.fields.gameStatusId,
				baKyokuHonba:  tt.fields.baKyokuHonba,
				riichiers:     tt.fields.riichiers,
				ronWinner:     tt.fields.ronWinner,
				ronLoser:      tt.fields.ronLoser,
				tsumoWinner:   tt.fields.tsumoWinner,
				tenpaiers:     tt.fields.tenpaiers,
				hanFu:         tt.fields.hanFu,
			}
			if got := kyokuResult.WhoRiichi(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KyokuResult.WhoRiichi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKyokuResult_IsTonchaRonOrTsumoOrTenpai(t *testing.T) {
	type fields struct {
		kyokuResultId KyokuResultId
		gameStatusId  gs.GameStatusId
		baKyokuHonba  bkh.BaKyokuHonba
		riichiers     map[jicha.Jicha]struct{}
		ronWinner     *jicha.Jicha
		ronLoser      *jicha.Jicha
		tsumoWinner   *jicha.Jicha
		tenpaiers     *map[jicha.Jicha]struct{}
		hanFu         *hf.HanFu
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Toncha以外がtsumo",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 4, 1,
				)),
				map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, pecha: {},
				},
				nil,
				nil,
				&nancha,
				nil,
				FirstPtoP(hf.NewHanFu(hf.Han10, hf.FuUndefined)),
			},
			want: false,
		},
		{
			name: "Tonchaがtsumo",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 4, 1,
				)),
				map[jicha.Jicha]struct{}{
					toncha: {}, nancha: {}, pecha: {},
				},
				nil,
				nil,
				&toncha,
				nil,
				FirstPtoP(hf.NewHanFu(hf.Han10, hf.FuUndefined)),
			},
			want: true,
		},
		{
			name: "Tonchaがtenpai",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 4, 1,
				)),
				map[jicha.Jicha]struct{}{},
				nil,
				nil,
				nil,
				&map[jicha.Jicha]struct{}{
					toncha: {},
				},
				FirstPtoP(hf.NewHanFu(hf.Han10, hf.FuUndefined)),
			},
			want: true,
		},
		{
			name: "Tonchaがron",
			fields: fields{
				KyokuResultId(uuid.New()),
				gs.GameStatusId(uuid.New()),
				FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 4, 1,
				)),
				map[jicha.Jicha]struct{}{},
				&toncha,
				&nancha,
				nil,
				nil,
				FirstPtoP(hf.NewHanFu(hf.Han10, hf.FuUndefined)),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kyokuResult := &KyokuResult{
				kyokuResultId: tt.fields.kyokuResultId,
				gameStatusId:  tt.fields.gameStatusId,
				baKyokuHonba:  tt.fields.baKyokuHonba,
				riichiers:     tt.fields.riichiers,
				ronWinner:     tt.fields.ronWinner,
				ronLoser:      tt.fields.ronLoser,
				tsumoWinner:   tt.fields.tsumoWinner,
				tenpaiers:     tt.fields.tenpaiers,
				hanFu:         tt.fields.hanFu,
			}
			if got := kyokuResult.IsTonchaRonOrTsumoOrTenpai(); got != tt.want {
				t.Errorf("KyokuResult.IsTonchaRonOrTsumoOrTenpai() = %v, want %v", got, tt.want)
			}
		})
	}
}
