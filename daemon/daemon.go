package daemon

import (
	"context"
	"log"
	"time"
	"os"
	"sync"
)

type Server interface {
	Start(ctx context.Context, d time.Duration, msg string, wg sync.WaitGroup)
	Stop()
}

type SampleDaemon struct {

}

func New() SampleDaemon {
	return SampleDaemon{}
}


func (s *SampleDaemon) Start(ctx context.Context, d time.Duration, msg string, wg sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for {
		select {
		case <-time.After(d):
			log.Println(msg)
			sampleProcess()
		case <-ctx.Done():
			log.Print(ctx.Err())
			return
		}
	}	
}

func (s *SampleDaemon) Stop() {
	log.Println("stopping")
	// TODO cleanup
	os.Exit(1)
}

func sampleProcess() {
	log.Println("step...1")
	time.Sleep(10 * time.Second)
	log.Println("step...2")
}