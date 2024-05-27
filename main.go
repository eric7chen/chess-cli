package main

import (
	"fmt"
	"github.com/notnil/chess"
	"math/rand"
	"reflect"
)

var algebraicNotation = chess.AlgebraicNotation{}

func main() {
	game := chess.NewGame()
	fmt.Println(reflect.TypeOf(game))
	userColor := rand.Int() % 2
	if userColor == 0 {
		fmt.Println("You are White")
	} else {
		fmt.Println("You are Black")
	}

	turnCounter := 0
	for game.Outcome() == chess.NoOutcome {
		fmt.Printf("Turn number %d\n", turnCounter+1)
		if turnCounter%2 == userColor {
			userMove(game)
		} else {
			moves := game.ValidMoves()
			move := moves[rand.Intn(len(moves))]
			moveString := algebraicNotation.Encode(game.Position(), move)
			fmt.Printf("Computer moves: %s\n", moveString)
			game.Move(move)
		}
		fmt.Println(game.Position().Board().Draw())
		turnCounter += 1
	}

	// // generate moves until game is over
	// for game.Outcome() == chess.NoOutcome {
	// 	// select a random move

	// }

	// // print outcome and game PGN
	// fmt.Printf("Game completed. %s by %s.\n", game.Outcome(), game.Method())
	// fmt.Println(game.String())
}

func userMove(game *chess.Game) {
	var userMove string
	var err error
	fmt.Println("Enter Move: ")
	fmt.Scanln(&userMove)
	err = game.MoveStr(userMove)
	for err != nil {
		fmt.Println("Invalid move, try again!")
		fmt.Println("Enter Move: ")
		fmt.Scanln(&userMove)
		err = game.MoveStr(userMove)
	}
}
