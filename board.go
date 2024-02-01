package main

import (
	"context"
	"fmt"
)

// Board is the game of life board
type Board struct {
	_      struct{}
	width  int
	height int
	cells  [][]*Cell
}

// Cell returns a Cell referred by *Point
func (b *Board) Cell(ctx context.Context, p *Point) *Cell {
	if p.X < 0 || p.X > b.width-1 {
		return nil
	}
	if p.Y < 0 || p.Y > b.height-1 {
		return nil
	}

	return b.cells[p.X][p.Y]
}

// SetCell updates cell's current alive status
func (b *Board) SetCell(ctx context.Context, p *Point, isAlive bool) {
	if c := b.Cell(ctx, p); c != nil {
		c.SetIsAlive(ctx, isAlive)
	}
}

// Update renders the board state
func (b *Board) Update(ctx context.Context) {
	b.renderCurrentGen(ctx)
	b.prepareNextGen(ctx)
}

func (b *Board) renderCurrentGen(ctx context.Context) {
	for x := 0; x < b.height; x++ {
		for y := 0; y < b.width; y++ {
			c := b.cells[y][x]
			fmt.Print(c)
			c.Update(ctx, b)
		}
		fmt.Println()
	}
}

func (b *Board) prepareNextGen(ctx context.Context) {
	for x := 0; x < b.width; x++ {
		for y := 0; y < b.height; y++ {
			b.cells[x][y].PrepareNextGen(ctx)
		}
	}
}
