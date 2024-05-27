package main

import (
	"flag"
	"fmt"
	"github.com/notnil/chess"
	"math/rand"
)

var (
	showBoardFlag bool
	showMoveFlag  bool
)

var algebraicNotation = chess.AlgebraicNotation{}

func main() {
	flag.BoolVar(&showBoardFlag, "board", true, "Show board throughout game")
	flag.BoolVar(&showMoveFlag, "moves", true, "Show valid moves each user turn")
	flag.Parse()
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Flags:")
		flag.PrintDefaults()
	}

	game := chess.NewGame()
	userColor := rand.Int() % 2
	if userColor == 0 {
		fmt.Println("You are White")
	} else {
		fmt.Println("You are Black")
	}

	if showBoardFlag {
		fmt.Println(game.Position().Board().Draw())
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

	// print outcome and game PGN
	fmt.Printf("Game completed. %s by %s.\n", game.Outcome(), game.Method())
	fmt.Println(game.String())
}

func userMove(game *chess.Game) {
	var userMove string
	var err error
	if showMoveFlag {
		printValidMoves(game)
	}
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

func printValidMoves(game *chess.Game) {
	fmt.Println("Valid moves: ")
	validMoves := game.ValidMoves()
	for i, move := range validMoves {
		fmt.Print(algebraicNotation.Encode(game.Position(), move))
		if i == len(validMoves)-1 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
}
