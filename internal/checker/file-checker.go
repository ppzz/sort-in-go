package Checker

import (
    "github.com/ppzz/sorting-go/internal/util"
    "io"
    "log"
    "os"
    "strconv"
    "strings"
)

type FileChecker struct {
    path string
}

func NewFileChecker(p string) *FileChecker {
    fc := new(FileChecker)
    fc.path = p
    return fc
}

func (fc *FileChecker) check() bool {
    file, err := os.Open(fc.path)
    util.NoError(err)
    defer file.Close()

    checker := NewChecker()
    result := true

    for {
        n, line, err := readline(file)
        if n > 0 { // 非空行
            arr := strings.Split(string(line), ",")
            val, _ := strconv.Atoi(arr[1])
            if !checker.IsOrdered(val) {
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
        log.Fatal("dead loop: should nor reach this code", )
    }
    return result
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
