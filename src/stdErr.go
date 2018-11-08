package main 
import (
	"os"
	"io"
)

func main () {
	io.WriteString(os.Stderr,"Error String")
}