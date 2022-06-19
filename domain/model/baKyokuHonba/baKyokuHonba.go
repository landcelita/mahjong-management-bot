package bakyokuhonba

import (
	"fmt"
)

type Ba uint

const (
	Ton Ba = iota + 1
	Nan
)

func (e Ba) names() []string {
	return []string {
		"Unknown",
		"東",
		"南",
	}
}

func (e Ba) String() string {
	if e > Nan {
		return "Unknown"
	}
	return e.names()[e]
}

type (
	BaKyokuHonba struct {
		ba		Ba
		kyoku	uint
		honba	uint
	}
)

func NewBaKyokuHonba(ba Ba, kyoku uint, honba uint) (*BaKyokuHonba, error) {
	if (ba != Nan && ba != Ton) || (kyoku == 0 || kyoku > 4) {
		return nil, fmt.Errorf("ValueError: baKyokuHonbaの値指定が不正です。")
	}

	baKyokuHonba := &BaKyokuHonba {
		ba: ba,
		kyoku: kyoku,
		honba: honba,
	}
	
	return baKyokuHonba, nil
}

func (baKyokuHonba *BaKyokuHonba) EqualsBaKyoku(otherBaKyokuHonba BaKyokuHonba) bool {
	if baKyokuHonba.ba == otherBaKyokuHonba.ba &&
	baKyokuHonba.kyoku == otherBaKyokuHonba.kyoku {
		return true
	} else {
		return false
	}
}

func (baKyokuHonba * BaKyokuHonba) IsLaterThanOrSameFor(otherBaKyokuHonba BaKyokuHonba) bool {
	if baKyokuHonba.ba > otherBaKyokuHonba.ba {
		return true
	} else if baKyokuHonba.ba == otherBaKyokuHonba.ba {
		if baKyokuHonba.kyoku > otherBaKyokuHonba.kyoku {
			return true
		} else if baKyokuHonba.kyoku == otherBaKyokuHonba.kyoku {
			if baKyokuHonba.honba >= otherBaKyokuHonba.honba {
				return true
			}
		}
	}

	return false
}

func (baKyokuHonba BaKyokuHonba) IncrementBaKyoku() (*BaKyokuHonba, error) {
	if baKyokuHonba.ba == Nan && baKyokuHonba.kyoku == 4 {
		return nil, fmt.Errorf("ValueError: これ以上baKyokuを進めることはできません。")
	}

	ret := baKyokuHonba
	ret.honba = 0

	if baKyokuHonba.kyoku <= 3 {
		ret.kyoku++
	} else if baKyokuHonba.ba == Ton && (baKyokuHonba.kyoku == 4) {
		ret.ba = Nan
		ret.kyoku = 1
	}

	return &ret, nil
}

func (baKyokuHonba BaKyokuHonba) IncrementHonba() (*BaKyokuHonba, error) {
	ret := baKyokuHonba
	ret.honba++

	return &ret, nil
}


