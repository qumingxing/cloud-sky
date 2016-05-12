// ParserHtml
package web

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
)

const htmlRoot = "D:\\workspace_java8\\StudentManage\\src\\resources\\"

type ParserHtml struct {
	Path string
}

func (cc *ParserHtml) GetHtml() (byt []byte, err error) {
	if strings.EqualFold(cc.Path, "") {
		return nil, errors.New("Path variable is not empty")
	}
	//dir, _ := os.Getwd()
	file, err := os.Open(htmlRoot + cc.Path)
	if err != nil {
		return
	}
	reader := bufio.NewReader(file)
	defer func() {
		file.Close()
	}()
	var buff [1024]byte
	bufferByte := bytes.NewBuffer(nil)
	for {
		index, err := reader.Read(buff[0:])
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		bufferByte.Write(buff[0:index])
	}
	return bufferByte.Bytes(), nil
}
