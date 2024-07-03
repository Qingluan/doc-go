
## DESC

a pure go to parse doc(97 ) 

## USAGE


```golang

package main

import (
	"fmt"
	"os"

	"github.com/Qingluan/doc-go/docparse"
)


func main() {
    // even table in doc , can be pretty ansic table
	res := docparse.ReadDoc(os.Args[1])
	for _, p := range res {
		fmt.Println(p)
	}
}

```
