package blackjack

import (
	"math/rand"
	"strconv"
	"time"

	randomdata "github.com/Pallinder/go-randomdata"
)

type Card struct {
	Name  string
	Type  string
	Suit  string
	Value int
}

type Player struct {
	Name  string
	User  bool
	Cards []Card
	Sum   int
}

type Game struct {
	Cards   []Card
	Players []Player
	Token   string
}

// Requests

type StartReq struct {
	Players int `form:"players" binding:"required"`
	Decks   int `form:"decks" binding:"required"`
}

// Current

var CurrentGames []Game

// Utility

func TakeCard(game *Game, player *Player) Card {
	num := randomdata.Number(0, len(game.Cards)-1)
	player.Cards = append(player.Cards, game.Cards[num])
	game.Cards = append(game.Cards[:num], game.Cards[num+1:]...)
	player.Sum += player.Cards[len(player.Cards)-1].Value
	return player.Cards[len(player.Cards)-1]
}

// Generators

func GenerateDeck() []Card {
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	numTypes := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	imgTypes := []string{"Jack", "Queen", "King"}
	var deck []Card

	for _, suit := range suits {
		for _, numType := range numTypes {
			deck = append(deck, Card{
				Name:  strconv.Itoa(numType) + " of " + suit,
				Type:  strconv.Itoa(numType),
				Suit:  suit,
				Value: numType,
			})
		}

		for _, imgType := range imgTypes {
			deck = append(deck, Card{
				Name:  imgType + " of " + suit,
				Type:  imgType,
				Suit:  suit,
				Value: 10,
			})
		}

		deck = append(deck, Card{
			Name:  "Ace of " + suit,
			Type:  "Ace",
			Suit:  suit,
			Value: 11,
		})
	}

	return deck
}

func GeneratePlayer(num int) []Player {
	var players []Player
	var emptyDeck []Card
	var fullNameList []string

	players = append(players, Player{
		Name:  "Player",
		User:  true,
		Cards: emptyDeck,
		Sum:   0,
	})
	var name string
	for i := 0; i < num; i++ {
		exist := false
		male := randomdata.Boolean()
		for {
			if male {
				name = randomdata.FirstName(randomdata.Male)
			} else {
				name = randomdata.FirstName(randomdata.Female)
			}
			for _, existName := range fullNameList {
				if existName == name {
					exist = true
				}
			}
			if !exist {
				break
			}
		}
		var emptyDeck []Card
		players = append(players, Player{
			Name:  name,
			User:  false,
			Cards: emptyDeck,
			Sum:   0,
		})
	}
	return players
}

func GenerateToken(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
