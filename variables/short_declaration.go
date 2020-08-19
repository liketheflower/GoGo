// Go program to show the use of var lexical 
// keyword and short declaration operator 
package main 

import ( 
	"fmt"
) 

// using var keyword to declare 
// and initialize the variable 
// it is package or you can say 
// global level scope 
var myvariable1 = 100 

// using short variable declaration 
// it will give an error as it is not 
// allowed outside the function 
var myvariable3 = 200 

func main() { 

// accessing myvariable1 inside the function 
fmt.Println(myvariable1) 

// using var keyword to declare 
// the variable along with type 
var myvariable2 int
// There is no need to put type. If you do that, it will give error.
//myvariable4 int := 200 
myvariable4 := 200 

fmt.Println(myvariable2) 


fmt.Println(myvariable3) 
fmt.Println(myvariable4) 
	
} 

