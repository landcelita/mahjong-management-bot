ドメインモデル図

局終了時のゲームの進み方:
- 誰かのscoreが負になる <span style='color: #00aa00;'>ゲーム終了</span>
飛び終了なしの場合は後で考える
- オーラス時
  - 親がロンorツモorテンパイをして親が1位になった時: <span style='color: #00aa00;'>ゲーム終了</span>
  - 親がロンorツモorテンパイをして親が1位にならない時: <span style='color: #aa0000;'>honbaが進む</span>
  - 親がロンorツモorテンパイのいずれでもない時: <span style='color: #00aa00;'>ゲーム終了</span>
- オーラスではない時
  - 親がロンorツモorテンパイをした時: <span style='color: #aa0000;'>honbaが進む</span>
  - 親がロンorツモorテンパイのいずれでもない時: <span style='color: #0000aa;'>(ba, kyoku)が進む</span>

```plantuml
@startuml
skinparam {
  PackageStyle rectangle
  defaultFontName Serif
}

package GameStatusAggrigates {
  object GameStatus {
    gameStatusId
    baKyokuHonba
    tonpuOrHanchan
    scoreBoardId
    playerIds[4]
    isActive
  }
}

package ScoreBoardAggrigates {
  object ScoreBoard {
    scoreBoardId
    scores[4]
    kyotaku
  }
}

note top of ScoreBoard
scoreは100点刻み。負もありうる。
endnote

note top of GameStatus
tonpuOrHanchanはtonpu又はhanchan。
baKyokuHonbaは(Ton, 2, 1)
のような値オブジェクト。
baはTon,Nanで、kyokuは1,2,3,4,
honbaは0以上の整数。
honbaは1ずつ増加し、(ba, kyoku)については必ず
(Ton, 1), (Ton, 2), (Ton, 3), (Ton, 4), // tonpuならここまで
(Nan, 1), (Nan, 2), (Nan, 3), (Nan, 4)の順に進む。
さらに、(ba, kyoku)が進む際はhonbaは0にリセットされる。
Olast()はtonpuOrHanchanが、
tonpuの時はbaKyokuHonba == (Ton, 4, *)ならば true
hanchanの時はbaKyokuHonba == (Nan, 4, *)ならば true
それ以外ならばfalse
tonpuの時、baKyokuHonbaのbaがNanになってはならない。
(ただし、これは南入、西入なしな場合で、ある場合はまた後で考える)
局終了時の進み方はかなり複雑なので、上記にmarkdownで記述している。
endnote

GameStatus "1" -> "1" ScoreBoard

package PlayerAggrigate {
  object Player {
    PlayerId
    Name
  }
}

note top of Player
PlayerIdはSlackが各ユーザーに付与しているIDを利用する。
endnote

GameStatus "4" -> "1" Player

package KyokuResultAggrigate {
  object KyokuResult {
    kyokuResultId
    gameStatusId
    baKyokuHonba
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

note top of KyokuResult
isOyaRonTsumoTenpai()は親がロンかツモ、またはテンパイ(流局時)
したことを表す。
endnote

KyokuResult "1" *-- "0" Ron
KyokuResult "1" *-- "0" Tsumo
KyokuResult "1" *-- "0" Ryukyoku
KyokuResult "1" *- "1" Riichi

Player "2" <-- "1" Ron
Player "1" <-- "1" Tsumo
Player "0..4" <---- "1" Ryukyoku
@enduml
```

gameの状況遷移図

```plantuml
@startuml

skinparam {
  PackageStyle rectangle
  defaultFontName Serif
}

[*] --> ゲームスタート: "ゲームスタート"と入力する
ゲームスタート --> 局の終了: プレーヤー選択
局の終了 --> winnerとloserの選択画面: ロンのとき
局の終了 --> winnerの選択画面: ツモのとき
局の終了 --> tenpai者の選択画面: 流局のとき
winnerとloserの選択画面 --> 翻と符の選択画面
winnerの選択画面 --> 翻と符の選択画面
tenpai者の選択画面 --> riichi者の選択画面
翻と符の選択画面 --> riichi者の選択画面
riichi者の選択画面 --> 局の終了: kyoku < 4 (東風戦),\n ba != Nan or kyoku < 4 (半荘戦)
riichi者の選択画面 --> ゲーム終了: kyoku == 4 (東風戦),\n ba == Nan and kyoku == 4 (半荘戦)
ゲーム終了 --> [*]

@enduml
```