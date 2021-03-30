package Checker

import (
	"github.com/ppzz/sorting-go/internal/sorter"
	"github.com/ppzz/sorting-go/internal/util"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type FileChecker struct {
	path    string
	checker *sorter.Checker
}

func NewFileChecker(filename string) *FileChecker {
	fc := new(FileChecker)
	fc.path = filename
	fc.checker = sorter.NewChecker()
	return fc
}

func (fc *FileChecker) Check() (bool, int) {
	file, err := os.Open(fc.path)
	util.NoError(err)
	defer file.Close()
	result := true

	for {
		n, line, err := readline(file)
		if n > 0 { // 非空行
			arr := strings.Split(string(line), ",")
			val, _ := strconv.Atoi(arr[1])
			if !fc.checker.IsOrdered(val) {
				result = false
				break
			}
			continue
		}
		if n == 0 && err == nil { // 空行
			continue
		}
		if err == io.EOF { // 文件尾
			break
		}
		log.Fatal("dead loop: should nor reach this code")
	}
	return result, fc.checker.GetCount()
}

func (fc *FileChecker) LoadToList(count int) []sorter.SortItem {
	file, err := os.Open(fc.path)
	util.NoError(err)
	defer file.Close()

	items := make([]sorter.SortItem, count)
	i := 0
	for {
		n, line, err := readline(file)
		if n > 0 { // 非空行
			arr := strings.Split(string(line), ",")
			seq, _ := strconv.Atoi(arr[0])
			val, _ := strconv.Atoi(arr[1])
			item := new(sorter.SortItem)
			item.Seq, item.Val = seq, val
			items[i] = *item
			i++
			continue
		}
		if n == 0 && err == nil { // 空行
			continue
		}
		if err == io.EOF { // 文件尾
			break
		}
		log.Fatal("dead loop: should nor reach this code")
	}
	return items
}

func readline(reader io.Reader) (count int, line []byte, err error) {
	line = make([]byte, 0)
	for {
		c := make([]byte, 1)
		n, e := reader.Read(c)
		if e != nil {
			err = e
			return
		}
		if c[0] == '\n' {
			return
		}
		if n > 0 {
			line = append(line, c[0])
			count++
		}
	}
}
