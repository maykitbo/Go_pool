package ex02

import "sync"

func multiplex(cs ...chan interface{}) chan interface{} {
    out := make(chan interface{})
    go func() {
        defer close(out)
        var wg sync.WaitGroup
        wg.Add(len(cs))
        for _, c := range cs {
            go func(c chan interface{}) {
                // for v := range c {
                    out <- <- c
                // }
                wg.Done()
            }(c)
        }
        wg.Wait()
    }()
    return out
}


