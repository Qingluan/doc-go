package main

import (
	"fmt"
	"os"

	"github.com/Qingluan/doc-go/docparse"
)

func main() {
	res := docparse.ReadDoc(os.Args[1])
	for _, p := range res {
		fmt.Println(p)
	}
}
