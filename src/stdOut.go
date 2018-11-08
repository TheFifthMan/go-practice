package main 

import (
	"fmt"
	"os"
)

func main () {
	myString := ""
	arg := os.Args
	if len(arg) == 1 {
		myString = "Please give me a string! "
	} else {
		myString = arg[1]
	} 
	fmt.Println(myString)
}