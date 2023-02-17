package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

const (
	alive   string = "O"
	dead    string = " "
	numGens int    = 10
)

type Universe struct {
	matrix        [][]string
	size          int
	numGeneration uint
	numAliveCells uint
}

func (u *Universe) String() string {
	var str string
	str += fmt.Sprintf("Generation: #%d\n", u.numGeneration)
	str += fmt.Sprintf("Alive: %d\n", u.numAliveCells)
	for i := 0; i < u.size; i++ {
		for j := 0; j < u.size; j++ {
			str += u.matrix[i][j]
		}
		if i != u.size {
			str += "\n"
		}
	}
	return str
}

func (u *Universe) calculateInitialState() (state string) {
	value := rand.Intn(2)
	if value == 1 {
		state = alive
		u.numAliveCells++
	} else {
		state = dead
	}
	return
}

func (u *Universe) populate() {
	for i := 0; i < u.size; i++ {
		row := make([]string, u.size)
		for j := 0; j < u.size; j++ {
			row[j] = u.calculateInitialState()
		}
		u.matrix[i] = row
	}
}

func (u *Universe) calculateNewState(row, col int) string {
	currentState := u.matrix[row][col]
	numAliveNeighbors := u.calculateAliveNeighbors(row, col)
	switch currentState {
	case alive:
		if numAliveNeighbors == 2 || numAliveNeighbors == 3 {
			return alive
		} else {
			u.numAliveCells--
			return dead
		}
	case dead:
		if numAliveNeighbors == 3 {
			u.numAliveCells++
			return alive
		}
	}
	return dead
}

func (u *Universe) calculateAliveNeighbors(row, col int) uint {
	var numAliveNeighbors uint
	matrixSize := u.size

	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue // ignore current cell
			}

			r := (row + dr + matrixSize) % matrixSize
			c := (col + dc + matrixSize) % matrixSize

			if u.matrix[r][c] == alive {
				numAliveNeighbors++
			}
		}
	}

	return numAliveNeighbors
}

func (u *Universe) nextGeneration() {
	nextGen := make([][]string, u.size)
	for i := 0; i < u.size; i++ {
		row := make([]string, u.size)
		for j := 0; j < u.size; j++ {
			row[j] = u.calculateNewState(i, j)
		}
		nextGen[i] = row
	}
	u.matrix = nextGen
	u.numGeneration++
}

func main() {
	var size int

	_, err := fmt.Scan(&size)
	if err != nil {
		return
	}

	// Deal with invalid values for the matrix size.
	if size < 1 {
		log.Fatal("Error: a matrix size needs to be at least 1x1.")
	}

	// Create the universe
	u := Universe{make([][]string, size), size, 0, 0}
	u.populate()

	// Go through the remaining generations
	for i := 0; i < numGens; i++ {
		u.nextGeneration()
		fmt.Print("\033[H\033[2J")
		fmt.Print(&u)
		time.Sleep(500 * time.Millisecond)
	}
}
