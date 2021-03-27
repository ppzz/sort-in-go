package main

import (
	"fmt"
	Checker "github.com/ppzz/sorting-go/internal/checker"
	RandGen "github.com/ppzz/sorting-go/internal/rand-gen"
	"github.com/ppzz/sorting-go/internal/sorter"
	"github.com/ppzz/sorting-go/internal/util"
	"github.com/spf13/viper"
	"path"
	"strconv"
	"time"
)

func main() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("dev") // 读取yaml配置文件
	err := viper.ReadInConfig()
	util.NoError(err)

	NumCount := viper.GetInt("randomFile.numCount")
	AssetRoot := viper.GetString("assetRoot")

	filename := "int-" + strconv.Itoa(NumCount) + ".txt"

	fmt.Printf("handle file %s\n", filename)
	start := time.Now().UnixNano() / 1000000
	fc := Checker.NewFileChecker(path.Join(AssetRoot, filename))
	items := fc.LoadToList()

	end := time.Now().UnixNano() / 1000000
	fmt.Printf("\t load file use time: %d ms\n", end-start)
	start = end

	s := sorter.NewSorter(items)
	s.BubbleSort(true)

	end = time.Now().UnixNano() / 1000000
	fmt.Printf("\t sort items use time: %d ms\n", end-start)
	start = end

	filename = fmt.Sprintf("int-%d-%s.txt", NumCount, "asc")
	fg := RandGen.NewFileGen(path.Join(AssetRoot, filename))
	fg.GenVisibleFromList(s.GetItems())

	end = time.Now().UnixNano() / 1000000
	fmt.Printf("\t save sorted file use time: %d ms, file: %s\n", end-start, filename)
	start = end
}
