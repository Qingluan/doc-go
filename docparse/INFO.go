package docparse

import (
	"bytes"
	"io"
	"unicode/utf16"
)

/*

VT


VT_EMPTY=0; VT_NULL=1; VT_I2=2; VT_I4=3; VT_R4=4; VT_R8=5; VT_CY=6;
VT_DATE=7; VT_BSTR=8; VT_DISPATCH=9; VT_ERROR=10; VT_BOOL=11;
VT_VARIANT=12; VT_UNKNOWN=13; VT_DECIMAL=14; VT_I1=16; VT_UI1=17;
VT_UI2=18; VT_UI4=19; VT_I8=20; VT_UI8=21; VT_INT=22; VT_UINT=23;
VT_VOID=24; VT_HRESULT=25; VT_PTR=26; VT_SAFEARRAY=27; VT_CARRAY=28;
VT_USERDEFINED=29; VT_LPSTR=30; VT_LPWSTR=31; VT_FILETIME=64;
VT_BLOB=65; VT_STREAM=66; VT_STORAGE=67; VT_STREAMED_OBJECT=68;
VT_STORED_OBJECT=69; VT_BLOB_OBJECT=70; VT_CF=71; VT_CLSID=72;
VT_VECTOR=0x1000;


*/

const (
	VT_EMPTY           = 0
	VT_NULL            = 1
	VT_I2              = 2
	VT_I4              = 3
	VT_R4              = 4
	VT_R8              = 5
	VT_CY              = 6
	VT_DATE            = 7
	VT_BSTR            = 8
	VT_DISPATCH        = 9
	VT_ERROR           = 10
	VT_BOOL            = 11
	VT_VARIANT         = 12
	VT_UNKNOWN         = 13
	VT_DECIMAL         = 14
	VT_I1              = 16
	VT_UI1             = 17
	VT_UI2             = 18
	VT_UI4             = 19
	VT_I8              = 20
	VT_UI8             = 21
	VT_INT             = 22
	VT_UINT            = 23
	VT_VOID            = 24
	VT_HRESULT         = 25
	VT_PTR             = 26
	VT_SAFEARRAY       = 27
	VT_CARRAY          = 28
	VT_USERDEFINED     = 29
	VT_LPSTR           = 30
	VT_LPWSTR          = 31
	VT_FILETIME        = 64
	VT_BLOB            = 65
	VT_STREAM          = 66
	VT_STORAGE         = 67
	VT_STREAMED_OBJECT = 68
	VT_STORED_OBJECT   = 69
	VT_BLOB_OBJECT     = 70
	VT_CF              = 71
	VT_CLSID           = 72
	VT_VECTOR          = 0x1000

	PAGECODE_UTF16LE = 1200
	PAGECODE_UTF16BE = 1201
	PAGECODE_UTF8    = 65001
	PAGECODE_GB2312  = 956
	PAGECODE_BIG5    = 950
)

var (
	VT = map[uint32]string{
		0:      "VT_EMPTY",
		1:      "VT_NULL",
		2:      "VT_I2",
		3:      "VT_I4",
		4:      "VT_R4",
		5:      "VT_R8",
		6:      "VT_CY",
		7:      "VT_DATE",
		8:      "VT_BSTR",
		9:      "VT_DISPATCH",
		10:     "VT_ERROR",
		11:     "VT_BOOL",
		12:     "VT_VARIANT",
		13:     "VT_UNKNOWN",
		14:     "VT_DECIMAL",
		16:     "VT_I1",
		17:     "VT_UI1",
		18:     "VT_UI2",
		19:     "VT_UI4",
		20:     "VT_I8",
		21:     "VT_UI8",
		22:     "VT_INT",
		23:     "VT_UINT",
		24:     "VT_VOID",
		25:     "VT_HRESULT",
		26:     "VT_PTR",
		27:     "VT_SAFEARRAY",
		28:     "VT_CARRAY",
		29:     "VT_USERDEFINED",
		30:     "VT_LPSTR",
		31:     "VT_LPWSTR",
		64:     "VT_FILETIME",
		65:     "VT_BLOB",
		66:     "VT_STREAM",
		67:     "VT_STORAGE",
		68:     "VT_STREAMED_OBJECT",
		69:     "VT_STORED_OBJECT",
		70:     "VT_BLOB_OBJECT",
		71:     "VT_CF",
		72:     "VT_CLSID",
		0x1000: "VT_VECTOR",
	}

	SUMMARY_ATTRIBS = []string{"codepage", "title", "subject", "author", "keywords", "comments",
		"template", "last_saved_by", "revision_number", "total_edit_time",
		"last_printed", "create_time", "last_saved_time", "num_pages",
		"num_words", "num_chars", "thumbnail", "creating_application",
		"security",
	}

	DOCSUM_ATTRIBS = []string{"codepage_doc", "category", "presentation_target", "bytes", "lines", "paragraphs",
		"slides", "notes", "hidden_slides", "mm_clips",
		"scale_crop", "heading_pairs", "titles_of_parts", "manager",
		"company", "links_dirty", "chars_with_spaces", "unused", "shared_doc",
		"link_base", "hlinks", "hlinks_changed", "version", "dig_sig",
		"content_type", "content_status", "language", "doc_version",
	}
)

func ParseAsPropertyBasic(rawBuf []byte, offset int, ptype uint32) (e any, size int) {
	reader := bytes.NewReader(rawBuf)
	switch ptype {
	case VT_I2:
		size = 2
		buf := make([]byte, size)
		copy(buf, rawBuf[offset:offset+size])
		// fmt.Println("I2:", buf)

		e2 := I16(buf)

		if e2 > 32768 {
			return int(e2) - 65536, size
		} else {
			return e2, size
		}

	case VT_UI2:
		size = 2
		buf := make([]byte, size)
		copy(buf, rawBuf[offset:offset+size])
		e := I16(buf)
		return e, size
	case VT_R4:
		size = 8
		buf := make([]byte, size)
		copy(buf, rawBuf[offset:offset+size])
		// fmt.Println("R4:", buf)
		e := I32(buf)
		return e, size
	case VT_I4, VT_INT, VT_ERROR, VT_UI4, VT_UINT:
		size = 4
		buf := make([]byte, size)
		copy(buf, rawBuf[offset:offset+size])
		e := I32(buf)
		return e, size
	case VT_BSTR, VT_LPSTR:
		countBuf := make([]byte, 4)
		copy(countBuf, rawBuf[offset:offset+4])
		offset += 4
		count := I32(countBuf)
		valueBuf := make([]byte, count)
		copy(valueBuf, rawBuf[offset:offset+int(count)])
		outbuf := bytes.TrimFunc(valueBuf, func(r rune) bool {
			return r == 0
		})
		size = int(4 + count)
		return outbuf, size
	case VT_BLOB:
		countBuf := make([]byte, 4)
		copy(countBuf, rawBuf[offset:offset+4])
		offset += 4
		count := I32(countBuf)
		valueBuf := make([]byte, count)
		copy(valueBuf, rawBuf[offset:offset+int(count)])
		size = int(4 + count)
		return valueBuf, size
	case VT_LPWSTR:
		countBuf := make([]byte, 4)
		copy(countBuf, rawBuf[offset:offset+4])
		offset += 4
		count := I32(countBuf) * 2
		valueBuf := make([]byte, count)
		copy(valueBuf, rawBuf[offset:offset+int(count)])
		size = int(4 + count)
		return UTF16Decode(valueBuf), size
	case VT_FILETIME:
		lbuf := make([]byte, 4)
		hbuf := make([]byte, 4)

		copy(lbuf, rawBuf[offset:offset+4])
		offset += 4
		copy(hbuf, rawBuf[offset:offset+4])
		l := I32(lbuf)
		h := uint64(I32(hbuf))
		cc := (h << 32) + uint64(l)
		// date := time.Unix(int64(cc/100000000), 0)
		// fmt.Println("FILETIME:", date)
		return cc, 8
	case VT_UI1:
		size = 1
		buf := make([]byte, size)
		reader.Seek(int64(offset), 0)
		reader.Read(buf)
		e := buf[0]
		return e, size
	case VT_CLSID:
		countBuf := make([]byte, 4)
		reader.Seek(int64(offset), 0)
		reader.Read(countBuf)
		count := I32(countBuf)
		valueBuf := make([]byte, count)
		reader.Seek(int64(offset)+4, 0)
		reader.Read(valueBuf)
		size = int(4 + count)
		return valueBuf, size
	case VT_BOOL:
		size = 2
		buf := make([]byte, size)
		reader.Seek(int64(offset), 0)
		reader.Read(buf)
		e := I16(buf)
		if e > 0 {
			return true, size
		}
		return false, size
	case VT_CF:
		countBuf := make([]byte, 4)
		reader.Seek(int64(offset), 0)
		reader.Read(countBuf)
		count := I32(countBuf)
		valueBuf := make([]byte, count)
		reader.Seek(int64(offset+4), 0)
		reader.Read(valueBuf)
		size = int(4 + count)
		return valueBuf, size
	}
	return nil, 0
}

func UTF16Decode(data []byte) string {
	runes := make([]uint16, len(data)/2)
	for i := range runes {
		runes[i] = uint16(data[2*i]) + uint16(data[2*i+1])<<8
	}
	return string(utf16.Decode(runes))
}

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
	fs := bytes.Split(buf[:n], []byte("\r"))
	if len(fs) > 1 {
		fs = fs[:len(fs)-1]
	}
	switch code {
	case PAGECODE_UTF16LE:

		for _, f := range fs {
			paragraphs = append(paragraphs, UTF16Decode(f))
		}
	case PAGECODE_UTF8:
		for _, f := range fs {
			paragraphs = append(paragraphs, string(f))
		}
	}
	return
}
