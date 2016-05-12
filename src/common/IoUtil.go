package common

import (
	"bufio"
	"io"
	"os"
)

//存在追加内容，不存在创建文件
func Append_CreateFile(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_APPEND|os.O_CREATE, 0644)
}
func ReadString(reader io.Reader) string {
	var byt [1024]byte
	var builder StringBuilder
	read := bufio.NewReader(reader)
	for {
		offset, err := read.Read(byt[0:])
		if err == io.EOF {
			break
		}
		builder.Concat(string(byt[0:offset]))
	}
	return builder.ToString()
}
