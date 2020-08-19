package main

import "fmt"

func main() {
const trueConst = true
type myBool  bool
    var defaultBool = trueConst
    var customizedBool myBool = trueConst
    fmt.Println("defaultBool", defaultBool)
    fmt.Println("CustomizedBool", customizedBool)
}



