// https://youtu.be/LvgVSSpwND8

package main

import (
    "fmt"
    "time"
    "sync"
)

func main() {
    var wg  sync.WaitGroup
    wg.Add(1)
    go func() {
        count("sheep")
        wg.Done()
    }()
    wg.Wait()
}

func count(thing string) {
    for i := 0; i < 5; i++ {
        fmt.Println(thing)
    }
    time.Sleep(time.Millisecond * 500)
}

