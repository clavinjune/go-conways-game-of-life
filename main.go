package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

const (
	Success int = iota
	Error
)

var (
	flagWidth           = flag.Int("w", 93, "width of the board")
	flagHeight          = flag.Int("h", 53, "height of the board")
	flagAlivePercentage = flag.Int("a", 9, "how many percent of the cell will be alive at the starting point")
	flagInterval        = flag.Duration("i", 250*time.Millisecond, "refresh interval")
	flagPattern         = flag.String("p", "", "read pattern, see example at ./patterns/")
)

func run(ctx context.Context, width, height, alivePercentage int, interval time.Duration, pattern string) int {
	ctx, cancel := context.WithCancel(ctx)
	board := MustProvideBoard(ctx, width, height, alivePercentage, pattern)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)

	go Update(ctx, board, interval)
	<-ch
	cancel()

	return Success
}

func Update(ctx context.Context, b *Board, interval time.Duration) {
	t := time.NewTicker(interval)
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			if err := exec.Command("clear").Run(); err != nil {
				panic(err.Error())
			}
			fmt.Printf("Gen %d\n", i)
			b.Update(ctx)
		}
	}
}

func main() {
	flag.Parse()
	os.Exit(
		run(
			context.Background(),
			*flagWidth,
			*flagHeight,
			*flagAlivePercentage,
			*flagInterval,
			*flagPattern,
		),
	)
}
