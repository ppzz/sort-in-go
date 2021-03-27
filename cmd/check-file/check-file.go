package main

import (
    "fmt"
    Checker "github.com/ppzz/sorting-go/internal/checker"
    "path"
    "strconv"
)

func main() {
    assetRoot := "asset/"
    filename := "int-" + strconv.Itoa(100) + ".txt"
    absPath := path.Join(assetRoot, filename)
    fc := Checker.NewFileChecker(absPath)
    result := fc.Check()
    fmt.Printf("file:%s, odered stats: %v\n", absPath, result)
}
