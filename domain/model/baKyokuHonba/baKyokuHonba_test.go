package bakyokuhonba

import (
	"reflect"
	"testing"
)

func TestBa_Names(t *testing.T) {
	tests := []struct {
		name string
		e    Ba
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ba.Names() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBa_String(t *testing.T) {
	tests := []struct {
		name string
		e    Ba
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("Ba.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBaKyokuHonba(t *testing.T) {
	type args struct {
		ba    Ba
		kyoku uint
		honba uint
	}
	tests := []struct {
		name    string
		args    args
		want    *BaKyokuHonba
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBaKyokuHonba(tt.args.ba, tt.args.kyoku, tt.args.honba)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBaKyokuHonba() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBaKyokuHonba() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaKyokuHonbaFromRepository(t *testing.T) {
	type args struct {
		ba    Ba
		kyoku uint
		honba uint
	}
	tests := []struct {
		name string
		args args
		want *BaKyokuHonba
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BaKyokuHonbaFromRepository(tt.args.ba, tt.args.kyoku, tt.args.honba); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaKyokuHonbaFromRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaKyokuHonba_EqualsBaKyoku(t *testing.T) {
	type fields struct {
		ba    Ba
		kyoku uint
		honba uint
	}
	type args struct {
		otherBaKyokuHonba BaKyokuHonba
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baKyokuHonba := &BaKyokuHonba{
				ba:    tt.fields.ba,
				kyoku: tt.fields.kyoku,
				honba: tt.fields.honba,
			}
			if got := baKyokuHonba.EqualsBaKyoku(tt.args.otherBaKyokuHonba); got != tt.want {
				t.Errorf("BaKyokuHonba.EqualsBaKyoku() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaKyokuHonba_IncrementBaKyoku(t *testing.T) {
	type fields struct {
		ba    Ba
		kyoku uint
		honba uint
	}
	tests := []struct {
		name    string
		fields  fields
		want    *BaKyokuHonba
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baKyokuHonba := BaKyokuHonba{
				ba:    tt.fields.ba,
				kyoku: tt.fields.kyoku,
				honba: tt.fields.honba,
			}
			got, err := baKyokuHonba.IncrementBaKyoku()
			if (err != nil) != tt.wantErr {
				t.Errorf("BaKyokuHonba.IncrementBaKyoku() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaKyokuHonba.IncrementBaKyoku() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaKyokuHonba_IncrementHonba(t *testing.T) {
	type fields struct {
		ba    Ba
		kyoku uint
		honba uint
	}
	tests := []struct {
		name    string
		fields  fields
		want    *BaKyokuHonba
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baKyokuHonba := BaKyokuHonba{
				ba:    tt.fields.ba,
				kyoku: tt.fields.kyoku,
				honba: tt.fields.honba,
			}
			got, err := baKyokuHonba.IncrementHonba()
			if (err != nil) != tt.wantErr {
				t.Errorf("BaKyokuHonba.IncrementHonba() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaKyokuHonba.IncrementHonba() = %v, want %v", got, tt.want)
			}
		})
	}
}
