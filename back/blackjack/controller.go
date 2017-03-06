package blackjack

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetGame(c *gin.Context) {
	var request TokenReq
	err := c.Bind(&request)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect data",
			"body":    nil,
		})
		return
	}

	index, err := FindPlayer(request.Token)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  3,
			"message": "Game not found",
			"body":    nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    CurrentGames[index],
	})
}

func GetAllGames(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    CurrentGames,
	})
}

func GetAllEndedGames(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    EndedGames,
	})
}

func StartGame(c *gin.Context) {
	var request StartReq
	err := c.Bind(&request)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect data",
			"body":    nil,
		})
		return
	}

	if request.Decks == 0 {
		request.Decks = 1
	}

	if (request.Decks < 0 || request.Decks > 10) || (request.Players < 2 || request.Players > 10) {
		c.JSON(400, gin.H{
			"status":  2,
			"message": "Deck can not be less than zero and more than ten. Players can not be less than two and more than ten.",
			"body":    nil,
		})
		return
	}

	var decks []Card
	for i := 0; i < request.Decks; i++ {
		decks = append(decks, GenerateDeck()...)
	}

	var emptyPlayers []Player
	newGame := Game{
		Ended:        false,
		FinalMessage: "",
		Winner:       emptyPlayers,
		Cards:        decks,
		Players:      GeneratePlayer(request.Players),
		Token:        GenerateToken(26),
	}

	for index, _ := range newGame.Players {
		TakeCard(&newGame, &newGame.Players[index])
		TakeCard(&newGame, &newGame.Players[index])
	}

	CurrentGames = append(CurrentGames, newGame)

	CheckFullWinners(&newGame)

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    newGame,
	})

	if newGame.Ended {
		DeactivateGame(newGame.Token)
	}

	return
}

func TakeCardGame(c *gin.Context) {
	var request TokenReq
	err := c.Bind(&request)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect data",
			"body":    nil,
		})
		return
	}

	index, err := FindPlayer(request.Token)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  3,
			"message": "Game not found",
			"body":    nil,
		})
		return
	}

	emptyDeck := CheckCardsLen(&CurrentGames[index])

	if !emptyDeck {
		TakeCard(&CurrentGames[index], &CurrentGames[index].Players[0])
	}

	CheckFullWinners(&CurrentGames[index])

	AllPlayersCicle(&CurrentGames[index])

	if !CurrentGames[index].Ended && CurrentGames[index].Players[0].Sum > 20 {
		CurrentGames[index].Players[0].Stay = true
		AllPlayersCicle(&CurrentGames[index])
	}

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    CurrentGames[index],
	})

	if CurrentGames[index].Ended {
		DeactivateGame(CurrentGames[index].Token)
	}

	return
}

func StopTakeGame(c *gin.Context) {
	var request TokenReq
	err := c.Bind(&request)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect data",
			"body":    nil,
		})
		return
	}

	index, err := FindPlayer(request.Token)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  3,
			"message": "Game not found",
			"body":    nil,
		})
		return
	}

	CurrentGames[index].Players[0].Stay = true

	AllPlayersCicle(&CurrentGames[index])

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    CurrentGames[index],
	})

	if CurrentGames[index].Ended {
		DeactivateGame(CurrentGames[index].Token)
	}

	return
}
