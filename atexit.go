package atexit

import "sync"

var std = new(AtExit)

func Add(fn func()) {
	std.Add(fn)
}

func Cleanup() {
	std.Cleanup()
}

type AtExit struct {
	mut       sync.Mutex
	callbacks []func()
}

func (a *AtExit) Add(fn func()) {
	a.mut.Lock()
	a.callbacks = append(a.callbacks, fn)
	a.mut.Unlock()
}

func (a *AtExit) Cleanup() {
	wg := new(sync.WaitGroup)
	a.mut.Lock()
	for i, fn := range a.callbacks {
		wg.Add(1)
		go func(f func()) {
			f()
			wg.Done()
		}(fn)
		a.callbacks[i] = nil
	}
	a.callbacks = a.callbacks[:0]
	a.mut.Unlock()
	wg.Wait()
}
