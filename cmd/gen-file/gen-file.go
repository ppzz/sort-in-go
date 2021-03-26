package main

import (
    "github.com/ppzz/sorting-go/internal/rand-gen"
    "path"
    "strconv"
)

func main() {
    assetRoot := "asset/"
    count := 100
    filename := "int-" + strconv.Itoa(100) + ".txt"

    absPath := path.Join(assetRoot, filename)
    fg := RandGen.NewFileGen(absPath)
    fg.GenVisible(count)
}
