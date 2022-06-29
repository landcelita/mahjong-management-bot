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
	Ron      = KyokuEndType("Ron")
	Tsumo    = KyokuEndType("Tsumo")
	Ryukyoku = KyokuEndType("Ryukyoku")
)

type (
	KyokuResult struct {
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
)

func NewKyokuResult(
	kyokuResultId KyokuResultId,
	gameStatusId gs.GameStatusId,
	baKyokuHonba bkh.BaKyokuHonba,
	riichiers map[jicha.Jicha]struct{},
	ronWinner *jicha.Jicha,
	ronLoser *jicha.Jicha,
	tsumoWinner *jicha.Jicha,
	tenpaiers *map[jicha.Jicha]struct{},
	hanFu *hf.HanFu,
) (*KyokuResult, error) {

	for k := range riichiers {
		if k != jicha.Toncha && k != jicha.Nancha &&
			k != jicha.Shacha && k != jicha.Pecha {
			return nil, fmt.Errorf("riichiersの値が不正です。")
		}
	}

	if tenpaiers != nil {
		for k := range *tenpaiers {
			if k != jicha.Toncha && k != jicha.Nancha &&
				k != jicha.Shacha && k != jicha.Pecha {
				return nil, fmt.Errorf("tenpaiersの値が不正です。")
			}
		}
	}

	if (ronWinner == nil) != (ronLoser == nil) {
		return nil, fmt.Errorf("ronWinnerとronLoserは共にnilか、共に値を持たせる必要があります。")
	}

	{
		cntNotNil := 0
		if ronWinner != nil {
			cntNotNil++
		}
		if tsumoWinner != nil {
			cntNotNil++
		}
		if tenpaiers != nil {
			cntNotNil++
		}

		if cntNotNil != 1 {
			return nil, fmt.Errorf("ron, tsumo, tenpaiのうちどれか一つにインスタンスを渡してください。")
		}
	}

	if ronWinner != nil && *ronWinner == *ronLoser {
		return nil, fmt.Errorf("ronWinnerとronLoserは別の人を選んでください。")
	}

	if (tenpaiers != nil && hanFu != nil) ||
		(tenpaiers == nil && hanFu == nil) {
		return nil, fmt.Errorf("ron又はtsumoの場合、またその場合のみ、hanFuが必要です。")
	}

	if tenpaiers != nil {
		for j := range riichiers {
			if _, exist := (*tenpaiers)[j]; !exist {
				return nil, fmt.Errorf("riichiしている人は必ずtenpaiしていなければいけません。")
			}
		}
	}

	kyokuResult := &KyokuResult{
		kyokuResultId: kyokuResultId,
		gameStatusId:  gameStatusId,
		baKyokuHonba:  baKyokuHonba,
		riichiers:     riichiers,
		ronWinner:     ronWinner,
		ronLoser:      ronLoser,
		tsumoWinner:   tsumoWinner,
		tenpaiers:     tenpaiers,
		hanFu:         hanFu,
	}

	return kyokuResult, nil
}

func (kyokuResult *KyokuResult) GetKyokuEndType() KyokuEndType {
	if kyokuResult.ronWinner != nil {
		return Ron
	} else if kyokuResult.tsumoWinner != nil {
		return Tsumo
	} else {
		return Ryukyoku
	}
}

func (kyokuResult *KyokuResult) CalcBaseScore() (uint, error) {
	if kyokuResult.ronWinner == nil &&
		kyokuResult.tsumoWinner == nil {
		return 0, fmt.Errorf("GetBaseScoreはTsumoかRonの場合のみ機能します。")
	}

	if kyokuResult.isTonchaRonWinner() || kyokuResult.isTonchaTenpai() {
		return kyokuResult.hanFu.CalcBaseScore(true), nil
	} else {
		return kyokuResult.hanFu.CalcBaseScore(false), nil
	}
}

func (kyokuResult *KyokuResult) WhoRonWinner() (*jicha.Jicha, error) {
	if kyokuResult.ronWinner == nil {
		return nil, fmt.Errorf("ロンでない場合WhoRonWinner()は使えません。")
	}
	ret := *kyokuResult.ronWinner
	return &ret, nil
}

func (kyokuResult *KyokuResult) WhoRonLoser() (*jicha.Jicha, error) {
	if kyokuResult.ronLoser == nil {
		return nil, fmt.Errorf("ロンでない場合WhoRonLoser()は使えません。")
	}
	ret := *kyokuResult.ronLoser
	return &ret, nil
}

func (kyokuResult *KyokuResult) WhoTsumo() (*jicha.Jicha, error) {
	if kyokuResult.tsumoWinner == nil {
		return nil, fmt.Errorf("ツモでない場合WhoTsumo()は使えません。")
	}
	ret := *kyokuResult.tsumoWinner
	return &ret, nil
}

func (kyokuResult *KyokuResult) WhoTenpai() (*map[jicha.Jicha]struct{}, error) {
	if kyokuResult.tenpaiers == nil {
		return nil, fmt.Errorf("流局でない場合WhoTenpai()は使えません。")
	}
	ret := make(map[jicha.Jicha]struct{})
	for key, val := range *kyokuResult.tenpaiers {
		ret[key] = val
	}
	return &ret, nil
}

func (kyokuResult *KyokuResult) WhoRiichi() *map[jicha.Jicha]struct{} {
	ret := make(map[jicha.Jicha]struct{})
	for key, val := range kyokuResult.riichiers {
		ret[key] = val
	}
	return &ret
}

func (kyokuResult *KyokuResult) BKH() bkh.BaKyokuHonba {
	return kyokuResult.baKyokuHonba
}

func (kyokuResult *KyokuResult) IsTonchaRonOrTsumoOrTenpai() bool {
	return kyokuResult.isTonchaRonWinner() ||
		kyokuResult.isTonchaTsumo() ||
		kyokuResult.isTonchaTenpai()
}

///// unexported

func (kyokuResult *KyokuResult) isTonchaRonWinner() bool {
	if kyokuResult.ronWinner != nil && *kyokuResult.ronWinner == jicha.Toncha {
		return true
	}
	return false
}

func (kyokuResult *KyokuResult) isTonchaTsumo() bool {
	if kyokuResult.tsumoWinner != nil && *kyokuResult.tsumoWinner == jicha.Toncha {
		return true
	}
	return false
}

func (kyokuResult *KyokuResult) isTonchaTenpai() bool {
	if kyokuResult.tenpaiers == nil {
		return false
	}
	if _, exist := (*kyokuResult.tenpaiers)[jicha.Toncha]; exist {
		return true
	}
	return false
}
