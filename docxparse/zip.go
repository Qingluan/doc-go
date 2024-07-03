package docxparse

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func GetDocxRawText(filename string) (string, error) {
	// 读取zip文件
	// 并解析word/document.xml
	// 返回文本内容

	zipReader, err := zip.OpenReader(filename)
	if err != nil {
		return "", err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		if f.Name == "word/document.xml" {
			rc, err := f.Open()
			if err != nil {
				return "", err
			}
			defer rc.Close()

			// 读取文件内容
			// 并返回文本内容
			// info := f.FileInfo()
			// size := info.Size()
			// fmt.Println(size)
			name := filepath.Join(os.TempDir(), time.Now().String())

			tmpWrite, _ := os.Create(name)
			// buf := make([]byte, 6553535)
			// n, err := rc.Read(buf)
			// if err != nil {
			// 	return "", err
			// }
			io.Copy(tmpWrite, rc)

			return name, nil

		}
	}
	return "", fmt.Errorf("file not found")
}
