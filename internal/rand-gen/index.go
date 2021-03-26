package RandGen

import (
    "math/rand"
    "sync"
)

var singleton *IntGen

type IntGen struct {
    rand *rand.Rand
}

var once sync.Once

// 使用单例模式
func NewIntGen(feed int64) *IntGen {
    once.Do(func() { // 确保在多个协程竞争的情况下只会被执行一次
        source := rand.NewSource(feed)
        singleton = &IntGen{rand: rand.New(source)}
    })
    return singleton
}

func (intGen *IntGen) Gen() int {
    return intGen.rand.Int()
}
