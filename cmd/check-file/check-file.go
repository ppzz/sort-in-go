package main

import (
	"fmt"
	Checker "github.com/ppzz/sorting-go/internal/checker"
	"github.com/ppzz/sorting-go/internal/util"
	"github.com/spf13/viper"
	"path"
	"strconv"
)

func main() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("dev") // 读取yaml配置文件
	err := viper.ReadInConfig()
	util.NoError(err)

	NumCount := viper.GetInt("randomFile.numCount")
	AssetRoot := viper.GetString("assetRoot")

	filename := "int-" + strconv.Itoa(NumCount) + ".txt"
	absPath := path.Join(AssetRoot, filename)
	fc := Checker.NewFileChecker(absPath)
	result, count := fc.Check()
	fmt.Printf("file: %s, odered stats: %v, sorted numbers: %d\n", absPath, result, count)
}
