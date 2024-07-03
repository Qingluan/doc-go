package main

import (
	"fmt"
	"os"

	"github.com/Qingluan/doc-go/docparse"
)

func main() {
	res := docparse.ReadDoc(os.Args[1])
	for i, p := range res {
		fmt.Println(i, p)
	}
}
