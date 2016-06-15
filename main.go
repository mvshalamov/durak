package main

import (
    "fmt"
    rand "math/rand"
)

type Pair struct {
    Name string
    Weight int
}


type Card struct {
    Lear string
    Data Pair
}


type Player struct {
    Name string
    Type bool
    Cards []Card
}


type Game struct {
    Players []Player
    Cur_step int
    CD CardDeck
    Trump Card
}

func (g *Game) set_trump() {
    g.CD.Cards, g.Trump = g.CD.Cards[:len(g.CD.Cards)-1], g.CD.Cards[len(g.CD.Cards)-1]
}

func (g *Game) pass_card() {
    for it := 0; it < len(g.Players); it++ {
        for i := 0; i < 6; i++ {
            fmt.Println(g.Players[it].Cards)
            g.Players[it].Cards = append(g.Players[it].Cards, g.CD.Cards[len(g.CD.Cards)-1])
            g.CD.Cards = g.CD.Cards[:len(g.CD.Cards)-1]
        }
    }
}

type CardDeck struct {
	Num_card int
    Lear []string
    Names []Pair
    Cards []Card
}

func (cd *CardDeck) Init_carddeck() {
    for it := 0; it < len(cd.Lear); it++ {
        for pair_val := 0; pair_val < len(cd.Names); pair_val++ {
            cd.Cards = append(cd.Cards, Card{Lear: cd.Lear[it], Data: cd.Names[pair_val]})
        }
    }
    cd.mix()
}

func (cd *CardDeck) mix() {
    for i := range cd.Cards {
        j := rand.Intn(i + 1)
        cd.Cards[i], cd.Cards[j] = cd.Cards[j], cd.Cards[i]
    }
}

func main() {
    Cardd := CardDeck{
        Num_card: 36,
        Lear: []string {"K", "T", "B", "C"},
        Names: []Pair{
            Pair{"6", 6}, Pair{"7", 7}, Pair{"8", 8}, Pair{"9", 9},
            Pair{"10", 10} , Pair{"V", 11}, Pair{"D", 12},
            Pair{"K", 13}, Pair{"T", 14},
        },
    }
    Cardd.Init_carddeck()
    fmt.Println("!!", Cardd.Cards)
    game := Game{
        Players: []Player {
            Player{Name: "man", Type: true}, Player{Name: "pc", Type: false},
        },
        Cur_step: 0,
        CD: Cardd,
    }

    game.set_trump()
    fmt.Println(game.Trump)
    fmt.Println(len(game.CD.Cards))
    game.pass_card()
    fmt.Println(len(game.CD.Cards))
}
