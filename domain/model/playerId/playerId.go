package playerid

import (
	"fmt"
	"regexp"
)

type PlayerId string

func NewPlayerId(value string) (*PlayerId, error) {
	if value == "" {
		return nil, fmt.Errorf("PlayerIdが空です。")
	}

	if ok, _ := regexp.MatchString(`^[A-Z0-9._-]{1,21}$`, value); !ok {
		return nil, fmt.Errorf("PlayerIdがSlackのMember IDの仕様とマッチしていません。")
	}

	id := PlayerId(value)
	return &id, nil
}