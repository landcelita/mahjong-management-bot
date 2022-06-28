package kyokuresult

import (
	"fmt"

	"github.com/google/uuid"
	bkh "github.com/landcelita/mahjong-management-bot/domain/model/baKyokuHonba"
	gs "github.com/landcelita/mahjong-management-bot/domain/model/gameStatus"
	hf "github.com/landcelita/mahjong-management-bot/domain/model/hanFu"
	jc "github.com/landcelita/mahjong-management-bot/domain/model/jicha"
)

type KyokuResultId uuid.UUID

type KyokuEndType string

const (
	Ron			= KyokuEndType("Ron")
	Tsumo		= KyokuEndType("Tsumo")
	Ryukyoku	= KyokuEndType("Ryukyoku")
)

type (
	KyokuResult struct {
		kyokuResultId	KyokuResultId
		gameStatusId	gs.GameStatusId
		baKyokuHonba	bkh.BaKyokuHonba
		riichiJichas	map[jc.Jicha]struct{}
		ronWinnerJicha	*jc.Jicha
		ronLoserJicha	*jc.Jicha
		tsumoJicha		*jc.Jicha
		tenpaiJichas	*map[jc.Jicha]struct{}
		hanFu			*hf.HanFu
	}
)

func NewKyokuResult(
	kyokuResultId	KyokuResultId,
	gameStatusId	gs.GameStatusId,
	baKyokuHonba	bkh.BaKyokuHonba,
	riichiJichas	map[jc.Jicha]struct{},
	ronWinnerJicha	*jc.Jicha,
	ronLoserJicha	*jc.Jicha,
	tsumoJicha		*jc.Jicha,
	tenpaiJichas	*map[jc.Jicha]struct{},
	hanFu			*hf.HanFu,
) (*KyokuResult, error) {

	for k := range riichiJichas {
		if k != jc.Toncha && k != jc.Nancha &&
		k != jc.Shacha && k != jc.Pecha {
			return nil, fmt.Errorf("riichiJichaの値が不正です。")
		}
	}

	if tenpaiJichas != nil {
		for k := range *tenpaiJichas {
			if k != jc.Toncha && k != jc.Nancha &&
			k != jc.Shacha && k != jc.Pecha {
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

func GetKyokuEndType(kyokuResult *KyokuResult) KyokuEndType {
	if kyokuResult.ronWinnerJicha != nil {
		return Ron
	} else if kyokuResult.tsumoJicha != nil {
		return Tsumo
	} else {
		return Ryukyoku
	}
}
