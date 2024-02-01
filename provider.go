package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	boardSyncOnce sync.Once
	board         *Board
)

// MustProvideBoard returns *Board with given configuration
// if pattern is a valid path, it will ignore the rests of the configuration
// otherwise it will create a board with width x height size, with a number of alive cells depending on alivePercentage
func MustProvideBoard(ctx context.Context, width, height, alivePercentage int, pattern string) *Board {
	if board == nil {
		boardSyncOnce.Do(func() {
			if strings.TrimSpace(pattern) != "" {
				board = provideBoardFromPattern(ctx, pattern)
			} else {
				board = provideBoardFromConfig(ctx, width, height, alivePercentage)
			}
		})
	}

	return board
}

func provideBoardFromConfig(ctx context.Context, width, height, alivePercentage int) *Board {
	b := &Board{
		width:  width,
		height: height,
		cells:  make([][]*Cell, width, width),
	}

	for w := range b.cells {
		b.cells[w] = make([]*Cell, height, height)

		for h := range b.cells[w] {
			b.cells[w][h] = ProvideCell(ctx, w, h, MustRandomIsAlive(ctx, alivePercentage))
		}
	}

	return b
}

func provideBoardFromPattern(ctx context.Context, pattern string) *Board {
	if _, err := os.Stat(pattern); errors.Is(err, os.ErrNotExist) {
		panic(pattern + " is not exists")
	}
	f, err := os.Open(pattern)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	var w, h int
	if _, err := fmt.Sscanf(scanner.Text(), "%dx%d", &w, &h); err != nil {
		panic(err.Error())
	}

	b := provideBoardFromConfig(ctx, w, h, 0)

	var x int
	for scanner.Scan() {
		for y, r := range []rune(scanner.Text()) {
			if r != ' ' {
				b.cells[y][x].SetIsAlive(ctx, true)
			}
		}
		x++
	}

	return b
}

// ProvideCell returns a new cell
func ProvideCell(ctx context.Context, x, y int, isAlive bool) *Cell {
	return &Cell{
		point:   ProvidePoint(ctx, x, y),
		isAlive: isAlive,
	}
}

// ProvidePoint returns a new point
func ProvidePoint(ctx context.Context, x, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}
