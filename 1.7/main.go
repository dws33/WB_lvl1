package main

import "sync"

func main() {
	aMap := sync.Map{}

	for i := range 5 {
		go func() {
			aMap.Store(i, struct{}{})
		}()
	}

	for i := range 5 {
		go func() {
			aMap.Load(i)
		}()
	}

	//////////////////////////////////////////

	bMap := newFtMap()

	for i := range 5 {
		go func() {
			bMap.Add(i, struct{}{})
		}()
	}

	for i := range 5 {
		go func() {
			bMap.Get(i)
		}()
	}

}

type ftMap struct {
	m  map[int]struct{}
	mu *sync.RWMutex
}

func newFtMap() ftMap {
	return ftMap{
		mu: new(sync.RWMutex),
		m:  make(map[int]struct{}),
	}
}

func (m ftMap) Add(key int, val struct{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[key] = val
}

func (m ftMap) Get(key int) struct{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.m[key]
}
