// https://youtu.be/LvgVSSpwND8

package main

import (
    "fmt"
    "time"
)

func main() {
     go  count("sheep")
     go  count("cow")
     time.Sleep(time.Second * 2)
}

func count(thing string) {
    for i := 0; i < 5; i++ {
        fmt.Println(thing)
        time.Sleep(time.Millisecond * 500)
    }
}

