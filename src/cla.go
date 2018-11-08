package main 
import (
	"fmt"
	"os"
	"strconv"
)

func main () {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please give me a string")
	}
	min, _ := strconv.ParseFloat(args[1],64)
	max, _ := strconv.ParseFloat(args[2],64)

	if min > max {
		t := max
		max = min
		min = t 
	}

	for i:=3; i < len(args); i++ {
		n, _ := strconv.ParseFloat(args[i],64)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}