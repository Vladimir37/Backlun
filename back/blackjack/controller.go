package blackjack

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

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

	var decks []Card
	for i := 0; i < request.Decks; i++ {
		decks = append(decks, GenerateDeck()...)
	}

	newGame := Game{
		Cards:   decks,
		Players: GeneratePlayer(request.Players),
		Token:   GenerateToken(26),
	}

	CurrentGames = append(CurrentGames, newGame)
}
