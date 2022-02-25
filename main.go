package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"
	"sync"

	"github.com/prakashmirji/examples/daemon"
)

func main() {

	ctx := context.Background()
	var wg sync.WaitGroup
	sampleMsg := "hello daemon"

	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	defer func() {
		signal.Stop(c)
		cancel()
	}()

	go func() {
		select {
		case <-c:
			log.Println("interrupt received, will wait for process to complete before exiting")
			cancel()
		case <-ctx.Done():
		}
	}()

	svr := daemon.New()
	
	svr.Start(ctx, 5*time.Second, sampleMsg, wg)
	wg.Wait()
	svr.Stop()
}