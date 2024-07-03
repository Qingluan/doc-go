package docparse

import (
	"io"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

func GetPlainText(reader io.ReadSeeker, code uint16) (paragraphs []string) {
	roff := 0x001c
	reader.Seek(int64(roff), 0)
	sizeBuf := make([]byte, 2)
	reader.Read(sizeBuf)
	size := I16(sizeBuf)
	reader.Seek(2048, 0)
	buf := make([]byte, size)
	// fmt.Println("size:", size-2048)
	n, _ := reader.Read(buf)
	docRaw := ""
	if code == PAGECODE_UTF16LE {
		docRaw = UTF16Decode(buf[:n])
	} else {
		docRaw = string(buf[:n])
	}

	fs := strings.Split(docRaw, "\r")
	if len(fs) > 1 {
		fs = fs[:len(fs)-1]
	}

	isTabel := false
	var T table.Writer

	for _, f := range fs {
		if strings.Contains(f, "\x07") {
			if !isTabel {
				T = table.NewWriter()
				isTabel = true
			}

			row := table.Row{}
			for _, i := range strings.Split(f, "\x07") {
				row = append(row, i)
			}

			if strings.HasSuffix(f, "\x07\x07") {
				T.AppendRow(row[:len(row)-1])
				paragraphs = append(paragraphs, T.Render())
				isTabel = false
			} else {
				T.AppendRow(row)
			}

		} else {
			paragraphs = append(paragraphs, f)
		}

	}
	return
}
