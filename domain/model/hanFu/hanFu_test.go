package hanfu

import (
	"reflect"
	"testing"
)

func TestNewHanFu(t *testing.T) {
	type args struct {
		han Han
		fu  Fu
	}
	tests := []struct {
		name    string
		args    args
		want    *HanFu
		wantErr bool
	}{
		{
			name: "異常系 hanがHan1未満",
			args: args{
				han: 0,
				fu: Fu20,
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "異常系 hanがHanTripleYakumanより大きい",
			args: args{
				han: 16,
				fu: Fu20,
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "異常系 fuが10",
			args: args{
				han: Han1,
				fu: Fu(10),
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "異常系 fuが170より大きい",
			args: args{
				han: Han1,
				fu: Fu(200),
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "異常系 fuが28",
			args: args{
				han: Han3,
				fu: Fu(28),
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "異常系 hanがHan4以下で,fuがFuUndefined",
			args: args{
				han: Han2,
				fu: FuUndefined,
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "異常系 hanがHan1でfuがFu25",
			args: args{
				han: Han1,
				fu: Fu25,
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "正常系 hanがHanTripleYakuman, fuがFuUndefined",
			args: args{
				han: HanTripleYakuman,
				fu: FuUndefined,
			},
			want: &HanFu {
				han: HanTripleYakuman,
				fu: FuUndefined,
			},
			wantErr: false,
		},
		{
			name: "正常系 fuが25",
			args: args{
				han: Han2,
				fu: Fu25,
			},
			want: &HanFu{
				han: Han2,
				fu: Fu25,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewHanFu(tt.args.han, tt.args.fu)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHanFu() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHanFu() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHanFu_CalcBaseScore(t *testing.T) {
	type fields struct {
		han Han
		fu  Fu
	}
	type args struct {
		isToncha bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint
	}{
		{
			name: "Han1Fu30 NotToncha",
			fields: fields {
				han: Han1,
				fu: Fu30,
			},
			args: args {
				isToncha: false,
			},
			want: 240,
		},
		{
			name: "Han4Fu25 Toncha",
			fields: fields {
				han: Han4,
				fu: Fu25,
			},
			args: args {
				isToncha: true,
			},
			want: 3200,
		},
		{
			name: "Han4Fu30 NotToncha",
			fields: fields {
				han: Han4,
				fu: Fu30,
			},
			args: args {
				isToncha: false,
			},
			want: 1920,
		},
		{
			name: "Han3Fu70 NotToncha",
			fields: fields {
				han: Han3,
				fu: Fu70,
			},
			args: args {
				isToncha: false,
			},
			want: 2000,
		},
		{
			name: "Han4Fu40 Toncha",
			fields: fields {
				han: Han4,
				fu: Fu40,
			},
			args: args {
				isToncha: true,
			},
			want: 4000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hanFu := HanFu{
				han: tt.fields.han,
				fu:  tt.fields.fu,
			}
			if got := hanFu.CalcBaseScore(tt.args.isToncha); got != tt.want {
				t.Errorf("HanFu.CalcBaseScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
