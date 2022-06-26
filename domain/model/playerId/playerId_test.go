package playerid

import (
	"testing"
	"reflect"
)

func TestNewPlayerId(t *testing.T) {
	s := [7]string {
		"0ABCDEFG92", "1", "", "abcdef100", "あ", "a a a a", "helloworld||",
	}
	w := [2]PlayerId {
		PlayerId(s[0]), PlayerId(s[1]),
	}
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    *PlayerId
		wantErr bool
	}{
		{
			name:		"正常系 数字と英大文字を両方含む",
			args:		args{value: s[0]},
			want:		&w[0],
			wantErr: 	false,
		},
		{
			name:		"正常系 数字のみ",
			args:		args{value: s[1]},
			want:		&w[1],
			wantErr:	false,
		},
		{
			name:		"異常系 空文字列",
			args:		args{value: s[2]},
			want:		nil,
			wantErr:	true,
		},
		{
			name:		"異常系 英小文字を含む",
			args: 		args{value: s[3]},
			want:		nil,
			wantErr:	true,
		},
		{
			name:		"異常系 日本語を含む",
			args:		args{value: s[4]},
			want:		nil,
			wantErr:	true,
		},
		{
			name:		"異常系 空白を含む",
			args:		args{value: s[5]},
			want:		nil,
			wantErr:	true,
		},
		{
			name:		"異常系 記号を含む",
			args:		args{value: s[6]},
			want:		nil,
			wantErr:	true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPlayerId(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPlayerId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlayerId() = %v, want %v", got, tt.want)
			}
		})
	}
}
