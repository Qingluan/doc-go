package main

import (
	"fmt"
	"os"
	"testdoc/docparse"
)

func main() {
	res := docparse.ReadDoc(os.Args[1])
	for i, p := range res {
		fmt.Println(i, p)
	}
}
