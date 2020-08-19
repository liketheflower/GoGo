// a simple programe to show how to use intergers in go
package main
import "fmt"

func main(){
    var X int8 = 10
    // demostrate overflow
    fmt.Println("X", X, X + 255)
}
