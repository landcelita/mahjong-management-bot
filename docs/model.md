ドメインモデル図
```plantuml
skinparam PackageStyle rectangle

package GameStatusAggrigate {
  object GameStatus {
    gameStatusId
    ba
    kyoku
    honba
    scoreId
    playerIds[4]
    isActive
  }

  object Score {
    scoreId
    scores[4]
  }
}

GameStatus "1" -* "1" Score

package KyokuResultAggrigate {
  object KyokuResult {
    kyokuResultId
    gameStatusId
    ba
    kyoku
    honba
    ronOrTsumoOrRyukyoku
  }

  object Riichi {
    kyokuResultId
    riichiPlayers
  }

  object Ron {
    kyokuResultId
    winnerId
    loserId
    Han
    Fu
  }

  object Tsumo {
    kyokuResultId
    winnerId
    Han
    Fu
  }

  object Ryukyoku {
    kyokuResultId
    tenpaiPlayers
  }
}

KyokuResult "1" *-- "0" Ron
KyokuResult "1" *-- "0" Tsumo
KyokuResult "1" *-- "0" Ryukyoku
KyokuResult "1" *- "1" Riichi

package PlayerAggrigate {
  object Player {
    PlayerId
    Name
  }
}

GameStatus "4" -right-> "1" Player
Player "2" <-- "1" Ron
Player "1" <-- "1" Tsumo
Player "0..4" <---- "1" Ryukyoku

```



gameの状況遷移図

```plantuml

[*] --> GameStart: players input "game start"
GameStart --> Wait_for_kyoku_end: input players 
Wait_for_kyoku_end --> Select_winner_and_loser: selected ron
Wait_for_kyoku_end --> Select_winner: selected tsumo
Wait_for_kyoku_end --> Select_tenpai: selected ryukyoku
Select_winner_and_loser --> Select_han_and_fu
Select_winner --> Select_han_and_fu
Select_tenpai --> Select_riichi
Select_han_and_fu --> Select_riichi
Select_riichi --> Wait_for_kyoku_end: kyoku < 4 (Tonpu),\n ba != "Nan" or kyoku < 4 (Hanchan)
Select_riichi --> GameOver: kyoku == 4 (Tonpu),\n ba == "Nan" and kyoku == 4 (Hanchan)
GameOver --> [*]
```