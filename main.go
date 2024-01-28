package main

import (
	"fmt"
	"os"

	scanner "github.com/QuodEstDubitandum/D--/scanner"
)

func main(){
	test := "₹"
	test2 := "Â"

	fmt.Println(test[0], test[1], test2[0])

	scanner.Start(os.Stdin, os.Stdout)
}