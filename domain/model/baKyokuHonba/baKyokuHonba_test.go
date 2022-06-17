package bakyokuhonba

import (
	"reflect"
	"testing"
)

func TestBa_String(t *testing.T) {
	tests := []struct {
		name string
		e    Ba
		want string
	}{
		{
			name:	"Ton",
			e:		Ton,
			want:	"東",
		},
		{
			name:	"Nan",
			e:		Nan,
			want:	"南",
		},
		{
			name:	"Out of the range 1",
			e:		0,
			want:	"Unknown",
		},
		{
			name:	"Out of the range 2",
			e:		3,
			want:	"Unknown",
		},
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
	bkh := [5]BaKyokuHonba{
		{ba: Ton, kyoku: 1, honba: 10},
		{ba: Ton, kyoku: 2, honba: 10},
		{ba: Ton, kyoku: 3, honba: 0},
		{ba: Ton, kyoku: 4, honba: 0},
		{ba: Nan, kyoku: 4, honba: 10},
	}
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
		{
			name:	 	"Correct 1",
			args: 		args{ba: Ton, kyoku: 1, honba: 10},
			want:		&bkh[0],
			wantErr:	false,
		},
		{
			name:	 	"Correct 2",
			args: 		args{ba: Ton, kyoku: 2, honba: 10},
			want:		&bkh[1],
			wantErr:	false,
		},
		{
			name:	 	"Correct 3",
			args: 		args{ba: Ton, kyoku: 3, honba: 0},
			want:		&bkh[2],
			wantErr:	false,
		},
		{
			name:	 	"Correct 4",
			args: 		args{ba: Ton, kyoku: 4, honba: 0},
			want:		&bkh[3],
			wantErr:	false,
		},
		{
			name:	 	"Correct 5",
			args: 		args{ba: Nan, kyoku: 4, honba: 10},
			want:		&bkh[4],
			wantErr:	false,
		},
		{
			name:		"Incorrect 1",
			args:		args{ba: 0, kyoku: 1, honba: 1},
			want:		nil,
			wantErr:	true,
		},
		{
			name:		"Incorrect 2",
			args:		args{ba: Ton, kyoku: 0, honba: 1},
			want:		nil,
			wantErr:	true,
		},
		{
			name:		"Incorrect 3",
			args:		args{ba: Nan, kyoku: 5, honba: 10},
			want:		nil,
			wantErr:	true,
		},
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
		{
			name:	"true 1",
			fields:	fields{ba: Ton, kyoku: 1, honba: 10},
			args: 	args{otherBaKyokuHonba: BaKyokuHonba{ba: Ton, kyoku: 1, honba: 0}},
			want:	true,
		},
		{
			name:	"true 2",
			fields:	fields{ba: Nan, kyoku: 3, honba: 10},
			args: 	args{otherBaKyokuHonba: BaKyokuHonba{ba: Nan, kyoku: 3, honba: 0}},
			want:	true,
		},
		{
			name:	"true 3",
			fields:	fields{ba: Nan, kyoku: 4, honba: 0},
			args:	args{otherBaKyokuHonba: BaKyokuHonba{ba: Nan, kyoku: 4, honba: 0}},
			want:	true,
		},
		{
			name:	"false 1",
			fields:	fields{ba: Ton, kyoku: 1, honba: 1},
			args:	args{otherBaKyokuHonba: BaKyokuHonba{ba: Ton, kyoku: 4, honba: 1}},
			want:	false,
		},
		{
			name:	"false 2",
			fields:	fields{ba: Nan, kyoku: 2, honba: 1},
			args:	args{otherBaKyokuHonba: BaKyokuHonba{ba: Ton, kyoku: 2, honba: 1}},
			want:	false,
		},
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
	bkh := [...]BaKyokuHonba{
		{ba: Nan, kyoku: 1, honba: 0},
		{ba: Nan, kyoku: 4, honba: 0},
		{ba: Ton, kyoku: 2, honba: 0},
	}
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
		{
			name:		"Correct 1",
			fields:		fields{ba: Ton, kyoku: 4, honba: 10},
			want:		&bkh[0],
			wantErr:	false,
		},
		{
			name:		"Correct 2",
			fields:		fields{ba: Nan, kyoku: 3, honba: 1},
			want:		&bkh[1],
			wantErr:	false,
		},
		{
			name:		"Correct 3",
			fields:		fields{ba: Ton, kyoku: 1, honba: 0},
			want:		&bkh[2],
			wantErr:	false,
		},
		{
			name:		"Incorrect 1",
			fields:		fields{ba: Nan, kyoku: 4, honba: 10},
			want:		nil,
			wantErr:	true,
		},
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
	bkh := [...]BaKyokuHonba{
		{ba: Ton, kyoku: 4, honba: 1},
		{ba: Nan, kyoku: 1, honba: 5},
		{ba: Nan, kyoku: 4, honba: 10},
	}
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
		{
			name:		"Correct 1",
			fields:		fields{ba: Ton, kyoku: 4, honba: 0},
			want:		&bkh[0],
			wantErr:	false,
		},
		{
			name:		"Correct 2",
			fields:		fields{ba: Nan, kyoku: 1, honba: 4},
			want:		&bkh[1],
			wantErr:	false,
		},
		{
			name:		"Correct 3",
			fields:		fields{ba: Nan, kyoku: 4, honba: 9},
			want:		&bkh[2],
			wantErr:	false,
		},
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
