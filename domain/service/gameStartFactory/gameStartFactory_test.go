package gamestartfactory

import (
	"reflect"
	"testing"

	bkh "github.com/landcelita/mahjong-management-bot/domain/model/baKyokuHonba"
	gs "github.com/landcelita/mahjong-management-bot/domain/model/gameStatus"
	jicha "github.com/landcelita/mahjong-management-bot/domain/model/jicha"
	pid "github.com/landcelita/mahjong-management-bot/domain/model/playerId"
	"github.com/landcelita/mahjong-management-bot/domain/model/score"
	sb "github.com/landcelita/mahjong-management-bot/domain/model/scoreBoard"
	toh "github.com/landcelita/mahjong-management-bot/domain/model/tonpuOrHanchan"
	gsinmem "github.com/landcelita/mahjong-management-bot/infra/inmem/gameStatus"
	sbinmem "github.com/landcelita/mahjong-management-bot/infra/inmem/scoreBoard"
	"github.com/landcelita/mahjong-management-bot/testutil"
)

func TestGameStartFactory_StartNewGame(t *testing.T) {
	t.Run("gameStart()が正常に機能するか？", func(t *testing.T) {
		// DI
		gsrepo := gsinmem.NewGameStatusRepoInMem()
		sbrepo := sbinmem.NewScoreBoardRepoInMem()
		gsfactory, err := NewGameStartFactory(gsrepo, sbrepo)
		if err != nil {
			panic(err)
		}
		players := map[jicha.Jicha]pid.PlayerId{
			jicha.Toncha:	pid.PlayerId("AAAA"),
			jicha.Nancha:	pid.PlayerId("BBBB"),
			jicha.Shacha:	pid.PlayerId("CCCC"),
			jicha.Pecha:	pid.PlayerId("DDDD"),
		}

		retGameStatus, retScoreBoard, err := gsfactory.StartNewGame(
			toh.Hanchan,
			players[jicha.Toncha],
			players[jicha.Nancha],
			players[jicha.Shacha],
			players[jicha.Pecha],
		)
		if err != nil {
			panic(err)
		}
		gsdata, err := gsrepo.GetAll()
		if err != nil {
			panic(err)
		}
		sbdata, err := sbrepo.GetAll()
		if err != nil {
			panic(err)
		}

		for id, got := range gsdata {
			expected, err := gs.NewGameStatus(
				id,
				testutil.FirstPtoV(bkh.NewBaKyokuHonba(bkh.Ton, 1, 0)),
				toh.Hanchan,
				players,
				true,
			)
			if err != nil {
				panic(err)
			}

			if !reflect.DeepEqual(expected, retGameStatus) {
				t.Errorf("GameStartFactory.StartNewGame() returned = %v, want %v", retGameStatus, expected)
			}
			if !reflect.DeepEqual(expected, got) {
				t.Errorf("GameStartFactory.StartNewGame() got = %v, want %v", got, expected)
			}
		}

		for id, got := range sbdata {
			expectedScore := map[jicha.Jicha]score.Score{
				jicha.Toncha:	testutil.FirstPtoV(score.NewScore(25000)),
				jicha.Nancha:	testutil.FirstPtoV(score.NewScore(25000)),
				jicha.Shacha:	testutil.FirstPtoV(score.NewScore(25000)),
				jicha.Pecha:	testutil.FirstPtoV(score.NewScore(25000)),
			}
			expected, err := sb.NewScoreBoard(
				id,
				expectedScore,
				testutil.FirstPtoV(score.NewScore(0)),
			)
			if err != nil {
				panic(err)
			}

			if !reflect.DeepEqual(expected, retScoreBoard) {
				t.Errorf("GameStartFactory.StartNewGame() returned = %v, want %v", retScoreBoard, expected)
			}
			if !reflect.DeepEqual(expected, got) {
				t.Errorf("GameStartFactory.StartNewGame() got = %v, want %v", got, expected)
			}
		}
	})
}
