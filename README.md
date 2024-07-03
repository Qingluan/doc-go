
## DESC

a pure go to read  doc(97 ) / docx (2007)

### feature

- [x] read text in docx
- [x] read text in doc
- [x] read table within docx
- [x] read table within doc
- [ ] read image within docx
- [ ] read image within doc

### TODO

- [ ] read docx with table and image and link

## USAGE


```golang
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


```
