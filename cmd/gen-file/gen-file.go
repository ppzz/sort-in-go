package main

import (
	"fmt"
	"github.com/ppzz/sorting-go/internal/rand-gen"
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
	absPath := path.Join(AssetRoot, filename)
	fg := RandGen.NewFileGen(absPath)
	start := time.Now().UnixNano()
	fg.GenVisible(NumCount)
	end := time.Now().UnixNano()
	fmt.Printf("use time: %d ms", (end-start)/1000/1000)
}
