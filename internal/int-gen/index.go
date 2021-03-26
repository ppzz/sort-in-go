package int_gen

import (
    "math/rand"
    "sync"
)

var singleton *IntGen

type IntGen struct {
    rand *rand.Rand
}

var once sync.Once

func GetInstance(feed int64) *IntGen {
    once.Do(func() {
        source := rand.NewSource(feed)
        singleton = &IntGen{rand: rand.New(source)}
    })
    return singleton
}

func (intGen *IntGen) Gen() int {
    return intGen.rand.Int()
}
