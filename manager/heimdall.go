package manager

import (
	"log"
	"time"
)

const (
	defaultCheckInterval     = time.Second * 10
	defaultChannelBufferSize = 10
)

type heimdallPlugin struct {
	CheckInterval time.Duration
	workerChannel chan bool
}

func (p *heimdallPlugin) Start() error {
	log.Println("Start")

	if p.CheckInterval <= 0 {
		p.CheckInterval = defaultCheckInterval
	}

	p.workerChannel = make(chan bool, defaultChannelBufferSize)

	go p.worker(p.workerChannel)
	go p.invoker(p.workerChannel)

	return nil
}

func (p *heimdallPlugin) worker(req <-chan bool) {
	for {
		_ = <-req

		log.Println("Do something")

		// TODO: Do something here.
		// - check up.
		// - get block height.
		// - alert?
		// - extra actions?
	}
}

func (p *heimdallPlugin) invoker(c chan<- bool) {
	for {
		c <- true
		time.Sleep(p.CheckInterval)
	}
}

func (p *heimdallPlugin) Stop() error {
	return nil
}
