```plantuml
actor player
left to right direction
rectangle {
    player -- (Select who in lab join the game)
    player -- (Select ron / tsumo / ryukyoku when the kyoku is over)
    player -- (Select who won or lost the kyoku)
    player -- (Select han / fu)
    player -- (Select who declared riichi)
    player -- (Select who completed tenpai)
}
player -- (Select who made chombo)
player -- (Select who received shukugi)
player -- (Input Uma and Oka)
player -- (Input the initial points)
```