package blackjack

import (
	"math/rand"
	"strconv"
	"time"

	"errors"

	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/bradfitz/slice"
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
	Stay  bool
	Cards []Card
	Sum   int
}

type Game struct {
	Ended        bool
	FinalMessage string
	Winner       []Player
	Cards        []Card
	Players      []Player
	Token        string
}

// Requests

type TokenReq struct {
	Token string `form:"token" binding:"required"`
}

type StartReq struct {
	Players int `form:"players" binding:"required"`
	Decks   int `form:"decks"`
}

// Current

var CurrentGames []Game
var EndedGames []Game

// Utility

func TakeCard(game *Game, player *Player) Card {
	num := randomdata.Number(0, len(game.Cards)-1)
	player.Cards = append(player.Cards, game.Cards[num])
	game.Cards = append(game.Cards[:num], game.Cards[num+1:]...)
	player.Sum += player.Cards[len(player.Cards)-1].Value
	return player.Cards[len(player.Cards)-1]
}

func FindPlayer(token string) (int, error) {
	founded := false
	targetIndex := 0

	for index, game := range CurrentGames {
		if game.Token == token {
			founded = true
			targetIndex = index
		}
	}

	if founded {
		return targetIndex, nil
	} else {
		return 0, errors.New("Not found")
	}
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
		Stay:  false,
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
			Stay:  false,
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

func MakeDecision(value int) bool {
	if value < 16 {
		return true
	} else if value < 18 {
		return randomdata.Boolean()
	} else {
		return false
	}
}

func CheckFullWinners(game *Game) bool {
	hasWinner := false
	for _, player := range game.Players {
		if player.Sum == 21 {
			game.Winner = append(game.Winner, player)
			game.Ended = true
			hasWinner = true
		}
	}
	if len(game.Winner) == 1 {
		game.FinalMessage = "One winner"
	} else if len(game.Winner) > 1 {
		game.FinalMessage = "Several winners"
	}

	return hasWinner
}

func checkMaxWinners(game *Game) {
	slice.Sort(game.Players, func(i, j int) bool {
		return game.Players[i].Sum > game.Players[j].Sum
	})

	maxSum := 0

	for _, player := range game.Players {
		if player.Sum > 21 {
			continue
		} else if player.Sum >= maxSum {
			game.Winner = append(game.Winner, player)
			game.Ended = true
			maxSum = player.Sum
		} else {
			break
		}
	}

	if len(game.Winner) == 1 {
		game.FinalMessage = "One winner"
	} else if len(game.Winner) > 1 {
		game.FinalMessage = "Several winners"
	}
}

func DeactivateGame(token string) {
	index, err := FindPlayer(token)
	if err != nil {
		return
	}

	EndedGames = append(EndedGames, CurrentGames[index])
	CurrentGames = append(CurrentGames[:index], CurrentGames[index+1:]...)
}

func CheckCardsLen(game *Game) bool {
	if len(game.Cards) == 0 {
		checkMaxWinners(game)
		return true
	}
	return false
}

func AllPlayersCicle(game *Game) {
	notStay := false
	for index, player := range game.Players {
		if player.User || player.Stay {
			continue
		} else {
			notStay = true
			takeDecision := MakeDecision(player.Sum)
			if takeDecision {
				TakeCard(game, &game.Players[index])
				CheckFullWinners(game)
			} else {
				game.Players[index].Stay = true
			}
		}
	}

	if !notStay {
		checkMaxWinners(game)
	}

	if notStay && game.Players[0].Stay {
		AllPlayersCicle(game)
	}
}
