package docparse

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadDoc(file string) []string {

	fp, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	defer fp.Close()
	ole, _ := Open(fp, "utf8")
	ds, _ := ole.ListDir()

	var root *File
	var reader_info io.ReadSeeker
	var plainText io.ReadSeeker
	for _, d := range ds {
		// fmt.Println(no, d.Sstart, d.Size, d.Name())
		if strings.Contains(d.Name(), "DocumentSummaryInformation") {

		} else if strings.Contains(d.Name(), "SummaryInformation") {
			// fmt.Println(d.Sstart, d.Size, d.Name())
			reader_info = ole.OpenFile(d, root)
		} else if strings.Contains(d.Name(), "Root Entry") {
			root = d
		} else if strings.Contains(d.Name(), "WordDocument") {
			plainText = ole.OpenFile(d, root)
		}
	}
	// reader_info = ole.stream_read(0, 448)
	// ole.stream_read()
	// reader_info := ole.stream_read(6, 468)
	// all_buf := make([]byte, 1024)
	// reader_info.Read(all_buf)
	// fmt.Println(all_buf)
	// return
	data := ParseInfomartion("SummaryInformation", reader_info)
	edata := make(map[string]any)
	fmt.Println()
	for ix, v := range data {
		if len(SUMMARY_ATTRIBS) <= int(ix) {
			continue
		}
		name := SUMMARY_ATTRIBS[ix-1]
		edata[name] = v
		// fmt.Println(ix, name, v)
	}

	// for name, v := range edata {
	// 	fmt.Println(name, v)
	// }
	// return

	codepage := edata["codepage"]
	fmt.Println("codepage:", codepage)
	return GetPlainText(plainText, codepage.(uint16))
	// roff := 0x001c
	// reader := ole.stream_read(0, 3626)
	// reader.Seek(int64(roff), 0)
	// size := make([]byte, 2)
	// fmt.Println(size)
	// reader.Read(size)
	// reader.Seek(2048, 2048+int(size[0]))
	// buf := make([]byte, 1024)
	// n, _ := reader.Read(buf)
	// // reader.
	// fmt.Println(string(buf[:n]))
}

func ParseInfomartion(name string, reader_info io.ReadSeeker) (data map[uint32]any) {
	_head := make([]byte, 28)
	reader_info.Read(_head)

	head := make([]byte, 20)
	reader_info.Read(head)
	// fmt.Println("head:", head)
	offset := I32(head[16:])
	// get section
	// fmt.Println("offset:", offset)
	reader_info.Seek(int64(offset), 0)

	sizeBuf := make([]byte, 4)
	reader_info.Read(sizeBuf)
	size := I32(sizeBuf)

	sss := make([]byte, size)
	copy(sss[0:4], []byte("****"))

	reader_info.Read(sss[4:])
	// fmt.Println("sss:", sss)
	num_prop := I32(sss[4:8])
	// fmt.Println("num_prop:", num_prop)
	data = make(map[uint32]any)

	for pid := 0; pid < int(num_prop); pid++ {
		// fmt.Println(8*pid+8, 8*pid+12, sss[8*pid+8:8*pid+12])
		property_id := I32(sss[8*pid+8 : 8*pid+12])
		offset = I32(sss[8*pid+12 : 8*pid+16])
		protype := I32(sss[offset : offset+4])

		// vt_name, ok := VT[protype]
		// if !ok {
		// 	vt_name = "unknown"
		// 	// break
		// }
		value, _ := ParseAsPropertyBasic(sss, int(offset)+4, protype)
		// fmt.Println("property_id:", property_id, "offset:", offset, "protype:", protype, "vt_name:", vt_name, "\n\t\tvalue :", value)
		data[property_id] = value
	}
	return
}
