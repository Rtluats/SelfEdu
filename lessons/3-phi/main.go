package main

import (
	"fmt"
	"sync"
	"time"
)

type Fork struct {
	status bool
	mutex  sync.RWMutex
}

func (f *Fork) inUse() bool {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return f.status
}

func (f *Fork) changeStatus() {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.status = !f.status
}

type Philosoph struct {
	id      int
	fork1   *Fork
	fork2   *Fork
	counter int32
}

func (p *Philosoph) Start() {
	for {
		if !p.fork1.inUse() && !p.fork2.inUse() {
			p.fork1.changeStatus()
			p.fork2.changeStatus()
			time.Sleep(time.Second / 2)
			p.counter++
			fmt.Printf("Eated %v|%v\n", p.id, p.counter)
			p.fork1.changeStatus()
			p.fork2.changeStatus()
		} else {
			time.Sleep(time.Second / 3)
			fmt.Printf("Was thinking %v|%v\n", p.id, p.counter)
		}
		time.Sleep(time.Second / 3)
	}
}

func main() {
	n := 5
	end := make(chan struct{})

	forks := make([]*Fork, 0, n)

	for i := 0; i < n; i++ {
		forks = append(forks, &Fork{false, sync.RWMutex{}})
	}

	philosophs := make([]*Philosoph, 0, n)

	for i := 0; i < n; i++ {
		philosophs = append(philosophs, &Philosoph{i + 1, forks[i], forks[(i+1)%n], 0})
	}

	for _, p := range philosophs {
		go p.Start()
	}

	<-end
}
