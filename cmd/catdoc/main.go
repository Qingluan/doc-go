package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Qingluan/doc-go/docparse"
	"github.com/Qingluan/doc-go/docxgen"
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
	} else if strings.HasSuffix(f, ".md") {
		create := docxgen.NewDocCreator()
		content, _ := os.ReadFile(f)
		create.FromMarkdown(string(content)).ExportDocx("out.docx")
		fmt.Println("docx文件已生成:", "out.docx")
	}

}
