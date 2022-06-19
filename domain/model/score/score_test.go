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
			name:    "ok 1",
			args:    args{scoreInt: -100},
			want:    &Score{score: -100},
			wantErr: false,
		},
		{
			name:    "ok 2",
			args:    args{scoreInt: 0},
			want:    &Score{score: 0},
			wantErr: false,
		},
		{
			name:    "ok 3",
			args:    args{scoreInt: 10000000},
			want:    &Score{score: 10000000},
			wantErr: false,
		},
		{
			name:    "ng 1",
			args:    args{scoreInt: 15},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "ng 2",
			args:    args{scoreInt: 1999999},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "ng 3",
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
			name:    "test 1",
			fields:  fields{ score: 100 },
			args:    args{ Score{score: 1000} },
			want:    Score{score: 1100},
		},
		{
			name:    "test 2",
			fields:  fields{ score: -1000 },
			args:    args{ Score{score: 1000} },
			want:    Score{score: 0},
		},
		{
			name:    "test 3",
			fields:  fields{ score: 100000 },
			args:    args{ Score{score: -10000000} },
			want:    Score{score: -9900000},
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
