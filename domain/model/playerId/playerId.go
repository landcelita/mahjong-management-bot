package playerid

import "fmt"

type PlayerId string

func NewPlayerId(value string) (PlayerId, error) {
	if value == "" {
		return "", fmt.Errorf("PlayerIdが空です。")
	}

	id := PlayerId(value)
	return id, nil
}