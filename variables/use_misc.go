package main

import "fmt"

func main(){
    a := 10
    b := &a
    fmt.Println("address of a is ", b)
    fmt.Println("value of a is ", *b)
}
