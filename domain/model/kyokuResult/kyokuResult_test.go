package kyokuresult

import (
	"testing"

	"github.com/google/uuid"
	bkh "github.com/landcelita/mahjong-management-bot/domain/model/baKyokuHonba"
	gs "github.com/landcelita/mahjong-management-bot/domain/model/gameStatus"
	hf "github.com/landcelita/mahjong-management-bot/domain/model/hanFu"
	jc "github.com/landcelita/mahjong-management-bot/domain/model/jicha"
	. "github.com/landcelita/mahjong-management-bot/testutil"
)

func TestNewKyokuResult(t *testing.T) {
	toncha := jc.Toncha
	nancha := jc.Nancha
	shacha := jc.Shacha
	pecha := jc.Pecha
	toncha2 := jc.Toncha

	type args struct {
		kyokuResultId  KyokuResultId
		gameStatusId   gs.GameStatusId
		baKyokuHonba   bkh.BaKyokuHonba
		riichiJichas   map[jc.Jicha]struct{}
		ronWinnerJicha *jc.Jicha
		ronLoserJicha  *jc.Jicha
		tsumoJicha     *jc.Jicha
		tenpaiJichas   *map[jc.Jicha]struct{}
		hanFu          *hf.HanFu
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
				riichiJichas: map[jc.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {}, pecha:  {},
				},
				ronWinnerJicha: &toncha,
				ronLoserJicha:  &nancha,
				tsumoJicha:     nil,
				tenpaiJichas:   nil,
				hanFu:          FirstPtoP(hf.NewHanFu(hf.Han1, hf.Fu40)),
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
				riichiJichas: map[jc.Jicha]struct{}{
				},
				ronWinnerJicha: nil,
				ronLoserJicha:  nil,
				tsumoJicha:     &nancha,
				tenpaiJichas:   nil,
				hanFu:          FirstPtoP(hf.NewHanFu(hf.Han1, hf.Fu40)),
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
				riichiJichas: map[jc.Jicha]struct{}{
					toncha: {}, nancha: {},
				},
				ronWinnerJicha: nil,
				ronLoserJicha:  nil,
				tsumoJicha:     nil,
				tenpaiJichas:   &map[jc.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {},
				},
				hanFu:          nil,
			},
			wantErr: false,
		},
		{
			name: "異常系 riichiJichasの値が不正",
			args: args{
				kyokuResultId: KyokuResultId(uuid.New()),
				gameStatusId:  gs.GameStatusId(uuid.New()),
				baKyokuHonba: FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 1, 1,
				)),
				riichiJichas: map[jc.Jicha]struct{}{
					jc.Jicha("Ton"): {},
				},
				ronWinnerJicha: nil,
				ronLoserJicha:  nil,
				tsumoJicha:     nil,
				tenpaiJichas:   &map[jc.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {},
				},
				hanFu:          nil,
			},
			wantErr: true,
		},
		{
			name: "異常系 tenpaiJichasの値が不正",
			args: args{
				kyokuResultId: KyokuResultId(uuid.New()),
				gameStatusId:  gs.GameStatusId(uuid.New()),
				baKyokuHonba: FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 1, 1,
				)),
				riichiJichas: map[jc.Jicha]struct{}{
					toncha: {}, nancha: {},
				},
				ronWinnerJicha: nil,
				ronLoserJicha:  nil,
				tsumoJicha:     nil,
				tenpaiJichas:   &map[jc.Jicha]struct{}{
					toncha: {}, nancha: {}, jc.Jicha("Sha"): {},
				},
				hanFu:          nil,
			},
			wantErr: true,
		},
		{
			name: "異常系 ronWinnerJichaのみに値が入っている",
			args: args{
				kyokuResultId: KyokuResultId(uuid.New()),
				gameStatusId:  gs.GameStatusId(uuid.New()),
				baKyokuHonba: FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Ton, 2, 3,
				)),
				riichiJichas: map[jc.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {}, pecha:  {},
				},
				ronWinnerJicha: &toncha,
				ronLoserJicha:  nil,
				tsumoJicha:     nil,
				tenpaiJichas:   nil,
				hanFu:          FirstPtoP(hf.NewHanFu(hf.Han1, hf.Fu40)),
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
				riichiJichas: map[jc.Jicha]struct{}{
					toncha: {}, nancha: {},
				},
				ronWinnerJicha: &toncha,
				ronLoserJicha:  &nancha,
				tsumoJicha:     nil,
				tenpaiJichas:  	&map[jc.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {},
				},
				hanFu:          FirstPtoP(hf.NewHanFu(hf.Han1, hf.Fu40)),
			},
			wantErr: true,
		},
		{
			name: "異常系 ronWinnerJichaとronLoserJichaが同じ",
			args: args{
				kyokuResultId: KyokuResultId(uuid.New()),
				gameStatusId:  gs.GameStatusId(uuid.New()),
				baKyokuHonba: FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Ton, 2, 3,
				)),
				riichiJichas: map[jc.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {}, pecha:  {},
				},
				ronWinnerJicha: &toncha,
				ronLoserJicha:  &toncha2,
				tsumoJicha:     nil,
				tenpaiJichas:   nil,
				hanFu:          FirstPtoP(hf.NewHanFu(hf.Han1, hf.Fu40)),
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
				riichiJichas: map[jc.Jicha]struct{}{
					toncha: {}, nancha: {},
				},
				ronWinnerJicha: nil,
				ronLoserJicha:  nil,
				tsumoJicha:     nil,
				tenpaiJichas:   &map[jc.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {},
				},
				hanFu:          FirstPtoP(hf.NewHanFu(hf.Han1, hf.Fu40)),
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
				riichiJichas: map[jc.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {}, pecha:  {},
				},
				ronWinnerJicha: &toncha,
				ronLoserJicha:  &nancha,
				tsumoJicha:     nil,
				tenpaiJichas:   nil,
				hanFu:          nil,
			},
			wantErr: true,
		},
		{
			name: "異常系 tenpaiJichasとriichiJichasの包含関係がおかしい",
			args: args{
				kyokuResultId: KyokuResultId(uuid.New()),
				gameStatusId:  gs.GameStatusId(uuid.New()),
				baKyokuHonba: FirstPtoV(bkh.NewBaKyokuHonba(
					bkh.Nan, 1, 1,
				)),
				riichiJichas: map[jc.Jicha]struct{}{
					toncha: {}, nancha: {}, pecha: {},
				},
				ronWinnerJicha: nil,
				ronLoserJicha:  nil,
				tsumoJicha:     nil,
				tenpaiJichas:   &map[jc.Jicha]struct{}{
					toncha: {}, nancha: {}, shacha: {},
				},
				hanFu:          nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewKyokuResult(tt.args.kyokuResultId, tt.args.gameStatusId, tt.args.baKyokuHonba, tt.args.riichiJichas, tt.args.ronWinnerJicha, tt.args.ronLoserJicha, tt.args.tsumoJicha, tt.args.tenpaiJichas, tt.args.hanFu)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewKyokuResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
