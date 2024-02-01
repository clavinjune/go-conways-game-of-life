package main

import "context"

const (
	CellAliveSymbol = "██"
	CellDeadSymbol  = "  "
)

// Cell will be rendered by board according to its isAlive, see CellAliveSymbol and CellDeadSymbol
type Cell struct {
	_           struct{}
	point       *Point
	isAlive     bool
	isAliveNext bool
}

// IsAlive returns whether a cell is alive or not
func (c *Cell) IsAlive(ctx context.Context) bool {
	return c.isAlive
}

// SetIsAlive set the cell's current alive status
func (c *Cell) SetIsAlive(ctx context.Context, isAlive bool) {
	c.isAlive = isAlive
}

// Update will set the cell's next status whether it will alive or not
func (c *Cell) Update(ctx context.Context, b *Board) {
	neighbours := c.getAliveNeighbours(ctx, b)
	n := len(neighbours)

	if c.isAbleToReproduction(ctx, n) {
		c.isAliveNext = true
	} else if c.isUnderOrOverPopulation(ctx, n) {
		c.isAliveNext = false
	} else {
		c.isAliveNext = c.isAlive
	}
}

// PrepareNextGen set the cell's next status to current status
func (c *Cell) PrepareNextGen(ctx context.Context) {
	c.isAlive = c.isAliveNext
}

// String renders the cells
func (c *Cell) String() string {
	if c.isAlive {
		return CellAliveSymbol
	}

	return CellDeadSymbol
}

func (c *Cell) getAliveNeighbours(ctx context.Context, b *Board) []*Cell {
	neighbours := make([]*Cell, 0, 0)
	for _, moveFn := range c.point.PossibleMoves(ctx) {
		neighbour := b.Cell(ctx, moveFn(ctx))
		if neighbour == nil {
			continue
		}

		if neighbour.IsAlive(ctx) {
			neighbours = append(neighbours, neighbour)
		}
	}

	return neighbours
}

func (c *Cell) isUnderOrOverPopulation(ctx context.Context, n int) bool {
	return c.isAlive && (n < 2 || n > 3)
}

func (c *Cell) isAbleToReproduction(ctx context.Context, n int) bool {
	return !c.isAlive && n == 3
}
