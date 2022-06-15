package tonpuorhanchan

import "fmt"

type TonpuOrHanchan uint

const (
	Tonpu TonpuOrHanchan = iota + 1
	Hanchan
)

func (e TonpuOrHanchan) Names() []string {
	return []string {
		"Unknown",
		"東風",
		"半荘",
	}
}

func (e TonpuOrHanchan) String() string {
	return e.Names()[e]
}

func NewTonpuOrHanchanFromUint(value uint) (*TonpuOrHanchan, error){
	if value != 1 && value != 2 {
		return nil, fmt.Errorf("ValueError: TonpuOrHanchanの値指定が不正です。")
	}
	
	tonpuOrHanchan := TonpuOrHanchan(value)
	return &tonpuOrHanchan, nil
}
