package score

import (
	"reflect"
	"testing"
)

func TestNewScore(t *testing.T) {
	type args struct {
		scoreInt int
	}
	tests := []struct {
		name    string
		args    args
		want    *Score
		wantErr bool
	}{
		{
			name:    "正常系 負",
			args:    args{scoreInt: -100},
			want:    &Score{score: -100},
			wantErr: false,
		},
		{
			name:    "正常系 0",
			args:    args{scoreInt: 0},
			want:    &Score{score: 0},
			wantErr: false,
		},
		{
			name:    "正常系 正",
			args:    args{scoreInt: 10000000},
			want:    &Score{score: 10000000},
			wantErr: false,
		},
		{
			name:    "異常系 100の倍数でない",
			args:    args{scoreInt: 15},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "異常系 100の倍数でない",
			args:    args{scoreInt: 1999999},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "異常系 100の倍数でない",
			args:    args{scoreInt: -199},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewScore(tt.args.scoreInt)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScore_Add(t *testing.T) {
	type fields struct {
		score int
	}
	type args struct {
		otherScore Score
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Score
	}{
		{
			name:   "正と正の和",
			fields: fields{score: 100},
			args:   args{Score{score: 1000}},
			want:   Score{score: 1100},
		},
		{
			name:   "正と負の和が0",
			fields: fields{score: -1000},
			args:   args{Score{score: 1000}},
			want:   Score{score: 0},
		},
		{
			name:   "正と負の和が負",
			fields: fields{score: 100000},
			args:   args{Score{score: -10000000}},
			want:   Score{score: -9900000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := Score{
				score: tt.fields.score,
			}
			if got := score.Add(tt.args.otherScore); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Score.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScore_Equals(t *testing.T) {
	type fields struct {
		score int
	}
	type args struct {
		otherScore Score
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "等しい場合 正",
			fields: fields{score: 10000},
			args:   args{otherScore: Score{score: 10000}},
			want:   true,
		},
		{
			name:   "等しい場合 負",
			fields: fields{score: -100},
			args:   args{otherScore: Score{score: -100}},
			want:   true,
		},
		{
			name:   "異なる場合 正と負",
			fields: fields{score: 10000},
			args:   args{otherScore: Score{score: -10000}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := Score{
				score: tt.fields.score,
			}
			if got := score.Equals(tt.args.otherScore); got != tt.want {
				t.Errorf("Score.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScore_LessThan(t *testing.T) {
	type fields struct {
		score int
	}
	type args struct {
		otherScore Score
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "引数と等しい",
			fields: fields{score: 10000},
			args:   args{otherScore: Score{score: 10000}},
			want:   false,
		},
		{
			name:   "引数より大きい",
			fields: fields{score: -100},
			args:   args{otherScore: Score{score: -1000}},
			want:   false,
		},
		{
			name:   "引数より小さい",
			fields: fields{score: -10000},
			args:   args{otherScore: Score{score: -100}},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := Score{
				score: tt.fields.score,
			}
			if got := score.LessThan(tt.args.otherScore); got != tt.want {
				t.Errorf("Score.LessThan() = %v, want %v", got, tt.want)
			}
		})
	}
}
