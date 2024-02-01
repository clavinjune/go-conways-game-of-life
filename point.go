package main

import (
	"context"
)

// MoveFn is a type of function to create a new point
type MoveFn func(context.Context) *Point

// Point is a struct that contains coordinate
type Point struct {
	_ struct{}
	X int
	Y int
}

// PossibleMoves returns the possible moves of a point
func (p *Point) PossibleMoves(ctx context.Context) []MoveFn {
	return []MoveFn{
		p.Up,
		p.Down,
		p.Left,
		p.Right,

		// diagnoal
		p.UpLeft,
		p.DownLeft,
		p.UpRight,
		p.DownRight,
	}
}

// Up returns a new point 1 cell above
func (p *Point) Up(ctx context.Context) *Point {
	return &Point{
		X: p.X,
		Y: p.Y - 1,
	}
}

// Down returns a new point 1 cell below
func (p *Point) Down(ctx context.Context) *Point {
	return &Point{
		X: p.X,
		Y: p.Y + 1,
	}
}

// Left returns a new point 1 cell to the left
func (p *Point) Left(ctx context.Context) *Point {
	return &Point{
		X: p.X - 1,
		Y: p.Y,
	}
}

// Right returns a new point 1 cell to the right
func (p *Point) Right(ctx context.Context) *Point {
	return &Point{
		X: p.X + 1,
		Y: p.Y,
	}
}

// UpLeft returns a new point 1 cell to the left and 1 cell above
func (p *Point) UpLeft(ctx context.Context) *Point {
	return &Point{
		X: p.X - 1,
		Y: p.Y - 1,
	}
}

// DownLeft returns a new point 1 cell to the left and 1 cell below
func (p *Point) DownLeft(ctx context.Context) *Point {
	return &Point{
		X: p.X - 1,
		Y: p.Y + 1,
	}
}

// UpRight returns a new point 1 cell to the right and 1 cell above
func (p *Point) UpRight(ctx context.Context) *Point {
	return &Point{
		X: p.X + 1,
		Y: p.Y - 1,
	}
}

// DownRight returns a new point 1 cell to the right and 1 cell below
func (p *Point) DownRight(ctx context.Context) *Point {
	return &Point{
		X: p.X + 1,
		Y: p.Y + 1,
	}
}
