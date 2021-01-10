package main

import (
	"fmt"
	"strconv"

	"github.com/notnil/chess"
)

var searchDepth = 5
var noQ = 2
var endGame = false

var pVal = [8][8]int{
	{0, 0, 0, 0, 0, 0, 0, 0},
	{50, 50, 50, 50, 50, 50, 50, 50},
	{10, 10, 20, 30, 30, 20, 10, 10},
	{5, 5, 10, 25, 25, 10, 5, 5},
	{0, 0, 0, 20, 20, 0, 0, 0},
	{5, -5, -10, 0, 0, -10, -5, 5},
	{5, 10, 10, -20, -20, 10, 10, 5},
	{0, 0, 0, 0, 0, 0, 0, 0}}

var nVal = [8][8]int{
	{-50, -40, -30, -30, -30, -30, -40, -50},
	{-40, -20, 0, 0, 0, 0, -20, -40},
	{-30, 0, 10, 15, 15, 10, 0, -30},
	{-30, 5, 15, 20, 20, 15, 5, -30},
	{-30, 0, 15, 20, 20, 15, 0, -30},
	{-30, 5, 10, 15, 15, 10, 5, -30},
	{-40, -20, 0, 5, 5, 0, -20, -40},
	{-50, -40, -30, -30, -30, -30, -40, -50}}

var bVal = [8][8]int{
	{-20, -10, -10, -10, -10, -10, -10, -20},
	{-10, 0, 0, 0, 0, 0, 0, -10},
	{-10, 0, 5, 10, 10, 5, 0, -10},
	{-10, 5, 5, 10, 10, 5, 5, -10},
	{-10, 0, 10, 10, 10, 10, 0, -10},
	{-10, 10, 10, 10, 10, 10, 10, -10},
	{-10, 5, 0, 0, 0, 0, 5, -10},
	{-20, -10, -10, -10, -10, -10, -10, -20}}

var rVal = [8][8]int{
	{0, 0, 0, 0, 0, 0, 0, 0},
	{5, 10, 10, 10, 10, 10, 10, 5},
	{-5, 0, 0, 0, 0, 0, 0, -5},
	{-5, 0, 0, 0, 0, 0, 0, -5},
	{-5, 0, 0, 0, 0, 0, 0, -5},
	{-5, 0, 0, 0, 0, 0, 0, -5},
	{-5, 0, 0, 0, 0, 0, 0, -5},
	{0, 0, 0, 5, 5, 0, 0, 0}}

var qVal = [8][8]int{
	{-20, -10, -10, -5, -5, -10, -10, -20},
	{-10, 0, 0, 0, 0, 0, 0, -10},
	{-10, 0, 5, 5, 5, 5, 0, -10},
	{-5, 0, 5, 5, 5, 5, 0, -5},
	{0, 0, 5, 5, 5, 5, 0, -5},
	{-10, 5, 5, 5, 5, 5, 0, -10},
	{-10, 0, 5, 0, 0, 0, 0, -10},
	{-20, -10, -10, -5, -5, -10, -10, -20}}

var kmVal = [8][8]int{
	{-30, -40, -40, -50, -50, -40, -40, -30},
	{-30, -40, -40, -50, -50, -40, -40, -30},
	{-30, -40, -40, -50, -50, -40, -40, -30},
	{-30, -40, -40, -50, -50, -40, -40, -30},
	{-20, -30, -30, -40, -40, -30, -30, -20},
	{-10, -20, -20, -20, -20, -20, -20, -10},
	{20, 20, 0, 0, 0, 0, 20, 20},
	{20, 30, 10, 0, 0, 10, 30, 20}}

var keVal = [8][8]int{
	{-50, -40, -30, -20, -20, -30, -40, -50},
	{-30, -20, -10, 0, 0, -10, -20, -30},
	{-30, -10, 20, 30, 30, 20, -10, -30},
	{-30, -10, 30, 40, 40, 30, -10, -30},
	{-30, -10, 30, 40, 40, 30, -10, -30},
	{-30, -10, 20, 30, 30, 20, -10, -30},
	{-30, -30, 0, 0, 0, 0, -30, -30},
	{-50, -30, -30, -30, -30, -30, -30, -50}}

var playerColor chess.Color
var posCheck int

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	game := chess.NewGame()
	var str string

	fmt.Scanln(&str)

	if str == "w" {
		playerColor = chess.White
	} else if str == "b" {
		playerColor = chess.Black
	}

	if playerColor == chess.Black {
		humanMove(game)
	}

	for game.Outcome() == chess.NoOutcome {
		posCheck = 0
		move := search(game)
		game.Move(move)
		fmt.Println(posCheck)
		fmt.Println("Computer's move: " + move.String())

		if game.Outcome() != chess.NoOutcome {
			break
		}

		humanMove(game)
	}

	fmt.Println(game.Position().Board().Draw())
	fmt.Println(evalBoard(game.Position()))
}

func humanMove(game *chess.Game) {
	var str string
	for {
		fmt.Print("Human's move: ")
		fmt.Scanln(&str)
		err := game.MoveStr(str)
		if err == nil {
			break
		}
	}
}

func search(game *chess.Game) *chess.Move {
	moves := game.ValidMoves()
	chosenMove := 0
	alpha := -9000000
	beta := 9000000
	for i, move := range moves {
		val := minMaxMin(game.Position().Update(move), alpha, beta, searchDepth-1)

		if val > alpha {
			alpha = val
			chosenMove = i
		}

	}

	return moves[chosenMove]
}

func minMaxMax(game *chess.Position, alpha, beta, depth int) int {
	moves := game.ValidMoves()

	if depth == 0 /* || len(moves) == 0*/ {
		posCheck++
		return evalBoard(game)
	}

	for _, move := range moves {
		val := minMaxMin(game.Update(move), alpha, beta, depth-1)

		if val >= beta {
			return beta
		}

		if val > alpha {
			alpha = val
		}
	}

	return alpha
}

func minMaxMin(game *chess.Position, alpha, beta, depth int) int {
	moves := game.ValidMoves()

	if depth == 0 /* || len(moves) == 0*/ {
		posCheck++
		return evalBoard(game)
	}

	for _, move := range moves {
		val := minMaxMax(game.Update(move), alpha, beta, depth-1)

		if val <= alpha {
			return alpha
		}

		if val < beta {
			beta = val
		}
	}

	return beta
}

func evalBoard(game *chess.Position) int {
	c := 1
	if game.Turn() != playerColor {
		c = -1
	}

	if game.Status() == chess.Checkmate {
		return c * 1000000
	} else if game.Status() == chess.Stalemate {
		return 0
	} else if game.Status() == chess.DrawOffer {
		return 0
	}

	noQ = 0

	boardState := game.Board().SquareMap()
	sum := 0
	for i, s := range boardState {
		sum += evalPiece(&s, &i)
	}

	if noQ == 0 && !endGame {
		searchDepth += 2
		endGame = true
		kmVal = keVal
	}

	return sum
}

func evalPiece(piece *chess.Piece, sq *chess.Square) int {
	file := []rune(sq.File().String())[0] - 97
	rank, _ := strconv.Atoi(sq.Rank().String())
	rank--
	if playerColor == chess.White {
		rank = 7 - rank
	}
	i := piece.Type()
	c := 1
	if piece.Color() != playerColor {
		c = -1
	}
	switch i {
	case chess.King:
		return c * (20000 + kmVal[rank][file])
	case chess.Queen:
		noQ++
		return c * (900 + qVal[rank][file])
	case chess.Rook:
		return c * (500 + rVal[rank][file])
	case chess.Bishop:
		return c * (330 + bVal[rank][file])
	case chess.Knight:
		return c * (320 + nVal[rank][file])
	case chess.Pawn:
		return c * (100 + pVal[rank][file])
	}
	return 0
}
