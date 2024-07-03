package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Qingluan/doc-go/docparse"
	"github.com/Qingluan/doc-go/docxparse"
)

func main() {
	f := os.Args[1]
	if strings.HasSuffix(f, ".docx") {
		res := docxparse.ReadDocx(f)
		for _, p := range res {
			fmt.Println(p)
		}

	} else if strings.HasSuffix(f, ".doc") {
		//doc
		res := docparse.ReadDoc(f)
		for _, p := range res {
			fmt.Println(p)
		}
	}

}
