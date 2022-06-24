package hanfu

import (
	"fmt"
)

type Han uint
type Fu uint

const (
	Han1 Han = iota + 1
	Han2
	Han3
	Han4
	Han5
	Han6
	Han7
	Han8
	Han9
	Han10
	Han11
	Han12
	HanYakuman
	HanDoubleYakuman
	HanTripleYakuman
)

const (
	FuUndefined = Fu(0)
	Fu20 = Fu(20)
	Fu25 = Fu(25)
	Fu30 = Fu(30)
	Fu40 = Fu(40)
	Fu50 = Fu(50)
	Fu60 = Fu(60)
	Fu70 = Fu(70)
	Fu80 = Fu(80)
	Fu90 = Fu(90)
	Fu100 = Fu(100)
	Fu110 = Fu(110)
	Fu120 = Fu(120)
	Fu130 = Fu(130)
	Fu140 = Fu(140)
	Fu150 = Fu(150)
	Fu160 = Fu(160)
	Fu170 = Fu(170)
)

type (
	HanFu struct {
		han		Han
		fu		Fu
	}
)

func NewHanFu(han Han, fu Fu) (*HanFu, error) {
	if han < Han1 || han > HanTripleYakuman {
		return nil, fmt.Errorf("hanはHan1以上HanTripleYakuman以下でなければいけません。")
	}

	if fu != FuUndefined && (fu < Fu20 || fu > Fu170) {
		return nil, fmt.Errorf("fuはFuUndefined、又はFu20以上Fu170以下でなければいけません。")
	}

	if fu != 25 && fu % 10 != 0 {
		return nil, fmt.Errorf("fuはFu25か、Fu10の倍数でなければいけません。")
	}

	if han <= Han4 && fu == FuUndefined {
		return nil, fmt.Errorf("hanがHan4以下の時、fuは指定されている必要があります。")
	}

	if han == Han1 && (fu == Fu20 || fu == Fu25) {
		return nil, fmt.Errorf("hanがHan1のときは必ずfuはFu30以上になります。")
	}

	hanFu := &HanFu{
		han: han,
		fu: fu,
	}

	return hanFu, nil
}

func (hanFu HanFu) CalcBaseScore (isToncha bool) uint {
	if hanFu.han <= Han4 {
		if isToncha {
			bs := uint(hanFu.fu) * (1 << uint(hanFu.han)) * 8
			if bs > 4000 { bs = 4000 }
			return bs
		} else {
			bs := uint(hanFu.fu) * (1 << uint(hanFu.han)) * 4
			if bs > 2000 { bs = 2000 }
			return bs
		}
	} else if hanFu.han == Han5 {
		if isToncha { return 4000 } else { return 2000 }
	} else if hanFu.han == Han6 || hanFu.han == Han7 {
		if isToncha { return 6000 } else { return 3000 }
	} else if hanFu.han == Han8 || hanFu.han == Han9 || hanFu.han == Han10 {
		if isToncha { return 8000 } else { return 4000 }
	} else if hanFu.han == Han11 || hanFu.han == Han12 {
		if isToncha { return 12000 } else { return 6000 }
	} else if hanFu.han == HanYakuman {
		if isToncha { return 16000 } else { return 8000 }
	} else if hanFu.han == HanDoubleYakuman {
		if isToncha { return 32000 } else { return 16000 }
	} else  {
		if isToncha { return 48000 } else { return 24000 }
	}
}
