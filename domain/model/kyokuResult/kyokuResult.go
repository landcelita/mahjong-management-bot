package kyokuresult

import (
	"fmt"

	"github.com/google/uuid"
	bakyokuhonba "github.com/landcelita/mahjong-management-bot/domain/model/baKyokuHonba"
	gamestatus "github.com/landcelita/mahjong-management-bot/domain/model/gameStatus"
	hanfu "github.com/landcelita/mahjong-management-bot/domain/model/hanFu"
	"github.com/landcelita/mahjong-management-bot/domain/model/jicha"
)

type KyokuResultId uuid.UUID

type (
	KyokuResult struct {
		kyokuResultId	KyokuResultId
		gameStatusId	gamestatus.GameStatusId
		baKyokuHonba	bakyokuhonba.BaKyokuHonba
		riichiJichas	map[jicha.Jicha]struct{}
		ronWinnerJicha	*jicha.Jicha
		ronLoserJicha	*jicha.Jicha
		tsumoJicha		*jicha.Jicha
		tenpaiJichas	*map[jicha.Jicha]struct{}
		hanFu			*hanfu.HanFu
	}
)

func NewKyokuResult(
	kyokuResultId	KyokuResultId,
	gameStatusId	gamestatus.GameStatusId,
	baKyokuHonba	bakyokuhonba.BaKyokuHonba,
	riichiJichas	map[jicha.Jicha]struct{},
	ronWinnerJicha	*jicha.Jicha,
	ronLoserJicha	*jicha.Jicha,
	tsumoJicha		*jicha.Jicha,
	tenpaiJichas	*map[jicha.Jicha]struct{},
	hanFu			*hanfu.HanFu,
) (*KyokuResult, error) {

	for k := range riichiJichas {
		if k != jicha.Toncha && k != jicha.Nancha &&
		k != jicha.Shacha && k != jicha.Pecha {
			return nil, fmt.Errorf("riichiJichaの値が不正です。")
		}
	}

	if tenpaiJichas != nil {
		for k := range *tenpaiJichas {
			if k != jicha.Toncha && k != jicha.Nancha &&
			k != jicha.Shacha && k != jicha.Pecha {
				return nil, fmt.Errorf("tenpaiJichaの値が不正です。")
			}
		}
	}

	if (ronWinnerJicha == nil) != (ronLoserJicha == nil) {
		return nil, fmt.Errorf("ronWinnerJichaとronLoserJichaは共にnilか、共に値を持たせる必要があります。")
	}

	{
		cntNotNil := 0
		if ronWinnerJicha != nil { cntNotNil++ }
		if tsumoJicha != nil { cntNotNil++ }
		if tenpaiJichas != nil { cntNotNil++ }

		if cntNotNil != 1 {
			return nil, fmt.Errorf("ron, tsumo, tenpaiのうちどれか一つにインスタンスを渡してください。")
		}
	}

	if ronWinnerJicha != nil && *ronWinnerJicha == *ronLoserJicha {
		return nil, fmt.Errorf("ronWinnerJichaとronLoserJichaは別の人を選んでください。")
	}

	if (tenpaiJichas != nil && hanFu != nil) ||
	(tenpaiJichas == nil && hanFu == nil) {
		return nil, fmt.Errorf("ron又はtsumoの場合、またその場合のみ、hanFuが必要です。")
	}

	if tenpaiJichas != nil {
		for j := range riichiJichas {
			if _, exist := (*tenpaiJichas)[j]; !exist {
				return nil, fmt.Errorf("riichiしている人は必ずtenpaiしていなければいけません。")
			}
		}
	}

	kyokuResult := &KyokuResult{
		kyokuResultId:	kyokuResultId,
		gameStatusId:	gameStatusId,
		baKyokuHonba:	baKyokuHonba,
		riichiJichas:	riichiJichas,
		ronWinnerJicha:	ronWinnerJicha,
		ronLoserJicha:	ronLoserJicha,
		tsumoJicha:		tsumoJicha,
		tenpaiJichas:	tenpaiJichas,
		hanFu:			hanFu,
	}

	return kyokuResult, nil
}
