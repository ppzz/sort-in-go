package multi_sort

import (
	"fmt"
	Checker "github.com/ppzz/sorting-go/internal/checker"
	RandGen "github.com/ppzz/sorting-go/internal/rand-gen"
	"github.com/ppzz/sorting-go/internal/sorter"
	"github.com/ppzz/sorting-go/internal/util"
	"github.com/spf13/viper"
	"log"
	"path"
	"sync"
	"time"
)

func MultiSort() {

	// 加载配置文件
	viper.AddConfigPath("./config")
	viper.SetConfigName("dev") // 读取yaml配置文件
	err := viper.ReadInConfig()
	util.NoError(err)

	// 准备一些变量
	NumCount := viper.GetInt("randomFile.numCount")
	// nameList := []string{"bubble", "insertion", "shell", "quick"} // 现有的算法列表
	nameList := []string{"shell", "quick", "merge"} // 现有的算法列表
	ch := make(chan ResultInfo, len(nameList))      // 用于 goroutine 返回结果

	items := genItems(NumCount)
	itemsList := cloneList(items, len(nameList)) // 克隆多份给goroutine使用

	// 启动协程，等待所有的协程完成工作
	var wg sync.WaitGroup
	for i, name := range nameList {
		wg.Add(1)
		go startOneSort(name, itemsList[i], true, ch, &wg) // 根据名字选择对应的算法排序，并返回统计信息
	}
	wg.Wait()
	close(ch)

	// 打印出每个协程的返回数据
	for i := range ch {
		fmt.Printf("name: %s\t, time: %d, mem:%d, cpu: %d, isOrdered: %v, isASC: %v\n", i.name, i.time, i.mem, i.cpu, i.isOrdered, i.isASC)
	}
}

func genItems(count int) []sorter.SortItem {
	start := time.Now().UnixNano() / 1000000

	s := make([]sorter.SortItem, count)
	rg := RandGen.NewIntGen(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		s[i] = sorter.SortItem{Seq: i, Val: rg.Gen()}
	}

	end := time.Now().UnixNano() / 1000000
	fmt.Printf("gen sortItem list: %d use time: %d ms\n", count, end-start)
	return s
}

func startOneSort(name string, items []sorter.SortItem, isASC bool, ch chan ResultInfo, wg *sync.WaitGroup) {
	start := time.Now().UnixNano()
	s := sorter.NewSorter(items)

	switch name {
	case "bubble":
		s.BubbleSort(isASC)
	case "insertion":
		s.InsertionSort(isASC)
	case "shell":
		s.ShellSort(isASC)
	case "quick":
		s.QuickSort(isASC)
	case "merge":
		s.MergeSort(isASC)
	default:
		log.Fatal("name not found:", name)
	}
	isOrdered, gotASC := s.Check()

	end := time.Now().UnixNano()
	ch <- ResultInfo{
		name:      name,
		time:      int((end - start) / 1000000),
		isOrdered: isOrdered,
		isASC:     gotASC,
	}
	wg.Done()
}

func cloneList(items []sorter.SortItem, count int) [][]sorter.SortItem {
	list := make([][]sorter.SortItem, count)
	for i, _ := range list {
		list[i] = make([]sorter.SortItem, len(items))
		copy(list[i], items)
	}
	return list
}

// 读入文件
func getList(filename string, count int) []sorter.SortItem {

	start := time.Now().UnixNano() / 1000000

	fc := Checker.NewFileChecker(filename)
	items := fc.LoadToList(count)

	end := time.Now().UnixNano() / 1000000
	fmt.Printf("%s:, load file use time: %d ms\n", path.Base(filename), end-start)

	return items
}

type ResultInfo struct {
	name      string
	time      int
	mem       int
	cpu       int
	isOrdered bool
	isASC     bool
}
