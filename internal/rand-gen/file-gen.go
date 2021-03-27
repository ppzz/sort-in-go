package RandGen

import (
	"bytes"
	"encoding/binary"
	"github.com/ppzz/sorting-go/internal/util"
	"log"
	"os"
	"strconv"
	"time"
)

type FileGen struct {
	path string
}

func NewFileGen(path string) *FileGen {
	return &FileGen{path: path}
}

func (fg *FileGen) Gen(count int) string {
	file, err := os.Create(fg.path)
	util.NoError(err)
	defer file.Close()

	intGen := NewIntGen(time.Now().UnixNano())
	for i := 1; i <= count; i++ {
		i := intGen.Gen()
		bytesBuffer := bytes.NewBuffer([]byte{})
		err = binary.Write(bytesBuffer, binary.BigEndian, int64(i))
		util.NoError(err)
		_, err := file.Write(bytesBuffer.Bytes())
		util.NoError(err)
	}
	return fg.path
}

func (fg *FileGen) GenVisible(count int) string {
	file, err := os.Create(fg.path)
	util.NoError(err)
	defer file.Close()
	intGen := NewIntGen(time.Now().UnixNano())
	for index := 1; index <= count; index++ {
		val := intGen.Gen()
		str := strconv.Itoa(index) + "," + strconv.Itoa(val) + "\n"
		b := []byte(str)
		n, err := file.Write(b)
		util.NoError(err)
		if n != len(b) {
			log.Fatal("error: 写入失败: n != len(strBytes) ")
		}
	}
	return fg.path
}
