package main

import (
	"aoc-2021/utils/conv"
	"aoc-2021/utils/files"
	"fmt"
	"strconv"
	"strings"
)

const (
	//GridSize is the amount of rows and cells per row
	GridSize = 5
)

// Bingo is the main game
type Bingo struct {
	Boards      []*Board
	DrawNumbers []string
	IsBingo     bool
}

type Board struct {
	rowCells      [GridSize][GridSize]string
	isUserBoard   bool
	hasBingo      bool
	winningNumber int
}

func main() {
	input := files.ReadFile(4, "\n")

	bingo := new(Bingo)

	bingo.gameSetup(input)

	wins := bingo.calcWins()

	fmt.Println("Part 1, first win: ")
	fmt.Println(wins[0].SumUnMarked() * wins[0].winningNumber)
	wins[0].PrintBoard()

	fmt.Println("==================")
	fmt.Println("Part 2, last win: ")
	lastWin := wins[len(wins)-1]
	fmt.Println(lastWin.SumUnMarked() * lastWin.winningNumber)
	lastWin.PrintBoard()

}

func (b *Bingo) gameSetup(input []string) {

	for i := 0; i < len(input); i++ {

		// First line are the draw numbers
		if i == 0 {
			b.DrawNumbers = strings.Split(input[i], ",")
		}

		if len(input[i]) == 0 {
			board := &Board{}

			// probably a better way to do this, but first line is empty, then look ahead at the next 5 rows
			for gridRows := 0; gridRows < GridSize+1; gridRows++ {
				if gridRows == 0 {
					continue
				}

				for index, gridColumn := range conv.ToIntSlice(strings.Split(input[i+gridRows], " ")) {
					board.rowCells[gridRows-1][index] = strconv.Itoa(gridColumn)
				}
				board.hasBingo = false
			}
			b.Boards = append(b.Boards, board)
		}
	}
}

func (b *Bingo) calcWins() []Board {

	winCount := 0
	var wins []Board

	for _, number := range b.DrawNumbers {
		for _, board := range b.Boards {

			if winCount > len(b.Boards) || board.hasBingo {
				continue
			}

			board.MarkNumberAsChecked(number)

			// Check if the board is a winner
			if board.IsBingo() {
				board.winningNumber, _ = strconv.Atoi(number)
				winCount += 1
				wins = append(wins, *board)
			}
		}
	}

	return wins
}

func (b *Board) SumUnMarked() int {
	sum := 0

	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			if b.rowCells[i][j] == "*" {
				continue
			}
			sum += conv.ToInt(b.rowCells[i][j])
		}
	}
	return sum
}

func (b *Board) MarkNumberAsChecked(number string) {
	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			if b.rowCells[i][j] == number {
				b.rowCells[i][j] = "*"
			}
		}
	}
}

func (b *Board) IsBingo() bool {

	countOfBingos := 0

	for i := 0; i < GridSize; i++ {

		numbersInAColumn := 0
		numbersInARow := 0

		for j := 0; j < GridSize; j++ {

			if b.rowCells[i][j] == "*" {
				numbersInAColumn += 1
			}

			if b.rowCells[j][i] == "*" {
				numbersInARow += 1
			}
		}

		if numbersInARow == GridSize || numbersInAColumn == GridSize {
			countOfBingos++
			b.hasBingo = true
			return true
		}

	}
	return b.hasBingo
}

func (b *Board) PrintBoard() {
	for _, cellsArray := range b.rowCells {
		fmt.Println(cellsArray)
	}
}
