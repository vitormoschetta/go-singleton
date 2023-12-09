package main

import "sync"

type Singleton struct{}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() { // <-- executado apenas uma vez, o que garante que o singleton seja criado apenas uma vez
		instance = &Singleton{}
	})
	return instance
}
