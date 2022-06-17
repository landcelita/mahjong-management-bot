package tonpuorhanchan

import (
	"testing"
	"reflect"
)

func TestTonpuOrHanchan_String(t *testing.T) {
	tests := []struct {
		name string
		e    TonpuOrHanchan
		want string
	}{
		{
			name: 	"Tonpu case",
			e:		Tonpu,
			want:	"東風",
		},
		{
			name:	"Hanchan case",
			e:		Hanchan,
			want:	"半荘",
		},
		{
			name:	"Out of the range case",
			e:		0,
			want:	"Unknown",
		},
		{
			name:	"Out of the range case 2",
			e:		3,
			want:	"Unknown",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("TonpuOrHanchan.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTonpuOrHanchanFromUint(t *testing.T) {
	toh := [2]TonpuOrHanchan {1, 2}
	type args struct {
		value uint
	}
	tests := []struct {
		name    string
		args    args
		want    *TonpuOrHanchan
		wantErr bool
	}{
		{
			name: 		"case 0",
			args: 		args{value: 0},
			want:		nil,
			wantErr:	true,
		},
		{
			name:		"case 1",
			args:		args{value: 1},
			want:		&toh[0],
			wantErr:	false,
		},
		{
			name:		"case 2",
			args:		args{value: 2},
			want:		&toh[1],
			wantErr:	false,
		},
		{
			name:		"case 3",
			args:		args{value: 3},
			want:		nil,
			wantErr:	true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTonpuOrHanchanFromUint(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTonpuOrHanchanFromUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTonpuOrHanchanFromUint() = %v, want %v", got, tt.want)
			}
		})
	}
}
