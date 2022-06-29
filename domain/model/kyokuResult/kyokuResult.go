package kyokuresult

import (
	"fmt"

	"github.com/google/uuid"
	bkh "github.com/landcelita/mahjong-management-bot/domain/model/baKyokuHonba"
	gs "github.com/landcelita/mahjong-management-bot/domain/model/gameStatus"
	hf "github.com/landcelita/mahjong-management-bot/domain/model/hanFu"
	"github.com/landcelita/mahjong-management-bot/domain/model/jicha"
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
		riichiJichas	map[jicha.Jicha]struct{}
		ronWinnerJicha	*jicha.Jicha
		ronLoserJicha	*jicha.Jicha
		tsumoJicha		*jicha.Jicha
		tenpaiJichas	*map[jicha.Jicha]struct{}
		hanFu			*hf.HanFu
	}
)

func NewKyokuResult(
	kyokuResultId	KyokuResultId,
	gameStatusId	gs.GameStatusId,
	baKyokuHonba	bkh.BaKyokuHonba,
	riichiJichas	map[jicha.Jicha]struct{},
	ronWinnerJicha	*jicha.Jicha,
	ronLoserJicha	*jicha.Jicha,
	tsumoJicha		*jicha.Jicha,
	tenpaiJichas	*map[jicha.Jicha]struct{},
	hanFu			*hf.HanFu,
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

func (kyokuResult *KyokuResult) GetKyokuEndType() KyokuEndType {
	if kyokuResult.ronWinnerJicha != nil {
		return Ron
	} else if kyokuResult.tsumoJicha != nil {
		return Tsumo
	} else {
		return Ryukyoku
	}
}

func (kyokuResult *KyokuResult) CalcBaseScore() (uint, error) {
	if kyokuResult.ronWinnerJicha == nil &&
	kyokuResult.tsumoJicha == nil {
		return 0, fmt.Errorf("GetBaseScoreはTsumoかRonの場合のみ機能します。")
	}

	if kyokuResult.isTonchaRonWinner() || kyokuResult.isTonchaTenpai() {
		return kyokuResult.hanFu.CalcBaseScore(true), nil
	} else {
		return kyokuResult.hanFu.CalcBaseScore(false), nil
	}
}

func (kyokuResult *KyokuResult) WhoRonWinner() (*jicha.Jicha, error) {
	if kyokuResult.ronWinnerJicha == nil {
		return nil, fmt.Errorf("ロンでない場合WhoRonWinner()は使えません。")
	}
	ret := *kyokuResult.ronWinnerJicha
	return &ret, nil
}

func (kyokuResult *KyokuResult) WhoRonLoser()  (*jicha.Jicha, error) {
	if kyokuResult.ronLoserJicha == nil {
		return nil, fmt.Errorf("ロンでない場合WhoRonLoser()は使えません。")
	}
	ret := *kyokuResult.ronLoserJicha
	return &ret, nil
}

func (kyokuResult *KyokuResult) WhoTsumo()  (*jicha.Jicha, error) {
	if kyokuResult.tsumoJicha == nil {
		return nil, fmt.Errorf("ツモでない場合WhoTsumo()は使えません。")
	}
	ret := *kyokuResult.tsumoJicha
	return &ret, nil
}

func (kyokuResult *KyokuResult) WhoTenpai()  (*map[jicha.Jicha]struct{}, error) {
	if kyokuResult.tenpaiJichas == nil {
		return nil, fmt.Errorf("流局でない場合WhoTenpai()は使えません。")
	}
	ret := make(map[jicha.Jicha]struct{})
	for key, val := range *kyokuResult.tenpaiJichas {
		ret[key] = val
	}
	return &ret, nil
}

func (kyokuResult *KyokuResult) BKH() bkh.BaKyokuHonba {
	return kyokuResult.baKyokuHonba
}

///// unexported

func (kyokuResult *KyokuResult) isTonchaRonWinner() bool {
	if kyokuResult.ronWinnerJicha != nil && *kyokuResult.ronWinnerJicha == jicha.Toncha {
		return true
	}
	return false
}

func (kyokuResult *KyokuResult) isTonchaTsumo() bool {
	if kyokuResult.tsumoJicha != nil && *kyokuResult.tsumoJicha == jicha.Toncha {
		return true
	}
	return false
}

func (kyokuResult *KyokuResult) isTonchaTenpai() bool {
	if kyokuResult.tenpaiJichas == nil {
		return false
	}
	if _, exist := (*kyokuResult.tenpaiJichas)[jicha.Toncha]; exist {
		return true
	}
	return false
}