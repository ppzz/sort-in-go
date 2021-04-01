package sorter

import (
	"log"
	"strconv"
)

type SortItem struct {
	Seq int
	Val int
}

func (m *SortItem) great(ms *SortItem) bool {
	return m.Val > ms.Val
}
func (m *SortItem) greatAndEqual(ms *SortItem) bool {
	return m.Val >= ms.Val
}

func (m *SortItem) less(ms *SortItem) bool {
	return !m.greatAndEqual(ms)
}

func (m *SortItem) lessAndEqual(ms *SortItem) bool {
	return !m.great(ms)
}

type Sorter struct {
	items []SortItem
}

func NewSorter(items []SortItem) *Sorter {
	s := new(Sorter)
	s.items = items
	return s
}

// 冒泡排序
func (s *Sorter) BubbleSort(isASC bool) {
	for i := 0; i < len(s.items); i++ {
		for j := i + 1; j < len(s.items); j++ {
			if shouldSwap(isASC, s.items[i], s.items[j]) {
				var t = s.items[i]
				s.items[i] = s.items[j]
				s.items[j] = t
			}
		}
	}
}

func shouldSwap(isASC bool, item1 SortItem, item2 SortItem) bool {
	return isASC && item1.great(&item2) || !isASC && item2.great(&item1)
}

// 插入排序
func (s *Sorter) InsertionSort(isASC bool) {
	for outIndex := 1; outIndex < len(s.items); outIndex++ {
		temp := s.items[outIndex]                   // 用来暂存外层循环的一个值，
		internalIndex := outIndex - 1               // 初始化内循环的起始位置，为当前位置的前一个值
		for ; internalIndex >= 0; internalIndex-- { // 内循环，用来跟暂存的值比较，并且移动位置
			shouldSwap := shouldSwap(isASC, s.items[internalIndex], temp) // 内循环值跟temp比较，判断是否到了temp应该的位置
			if shouldSwap {                                               // 不是temp的位置，将内循环的值向后移动一个位置
				s.items[internalIndex+1] = s.items[internalIndex]
				continue
			}
			break // 到了合适的位置 跳出循环
		}
		s.items[internalIndex+1] = temp // 将temp复制给该位置的值。
	}
}

func (s *Sorter) ShellSort(isASC bool) {
	count := len(s.items)
	step := count
	for step := step / 3; step >= 1; step /= 3 { // 遍历步长，直到步长等于1
		for groupIndex := 0; groupIndex < step; groupIndex++ { // 步长确定，遍历每个分组
			for outIndex := groupIndex + step; outIndex < count; outIndex += step { // 对每个分组，内的元素排序，外循环
				temp := s.items[outIndex]                         // 用来暂存外层循环的一个值，
				internalIndex := outIndex - step                  // 初始化内循环的起始位置，为当前位置的前一个值
				for ; internalIndex >= 0; internalIndex -= step { // 内循环，用来跟暂存的值比较，并且移动位置
					shouldSwap := shouldSwap(isASC, s.items[internalIndex], temp) // 内循环值跟temp比较，判断是否到了temp应该的位置
					if shouldSwap {                                               // 不是temp的位置，将内循环的值向后移动一个位置
						s.items[internalIndex+step] = s.items[internalIndex]
						continue
					}
					break // 到了合适的位置 跳出循环
				}
				s.items[internalIndex+step] = temp // 将temp复制给该位置的值。
			}
		}
	}
}

func (s *Sorter) Check() (isOrdered bool, isASC bool) {
	c := NewChecker()
	isOrdered = true
	for _, sortItem := range s.items {
		if !c.IsOrdered(sortItem.Val) {
			isOrdered = false
			break
		}
	}
	return isOrdered, c.GetIsAsc()
}

func (s *Sorter) GetItems() []SortItem {
	return s.items
}

func (s *Sorter) printItems() string {
	str := ""
	for _, v := range s.items {
		str += strconv.Itoa(v.Val)
		str += "\t"
	}
	return str
}

func (s *Sorter) SelectionSort(isASC bool) {
	count := len(s.items)
	for outIndex := 0; outIndex < count; outIndex++ {
		currentMinIndex := outIndex
		for internalIndex := outIndex + 1; internalIndex < count; internalIndex++ {
			shouldSwap := shouldSwap(isASC, s.items[currentMinIndex], s.items[internalIndex])
			if shouldSwap {
				currentMinIndex = internalIndex
			}
		}
		temp := s.items[currentMinIndex]
		s.items[currentMinIndex] = s.items[outIndex]
		s.items[outIndex] = temp
	}
}

func (s *Sorter) QuickSort(isASC bool) {
	count := len(s.items)

	s.quickSortRecursive(isASC, 0, count-1)
}

func (s *Sorter) quickSortRecursive(isASC bool, startIndex int, endIndex int) {
	if startIndex >= endIndex {
		return
	}

	leftIndex := startIndex
	rightIndex := endIndex - 1

	midIndex := endIndex // 选定的middle值的index
	for {
		for { // 从左查找比middle值大的元素的下标
			if leftIndex >= rightIndex || shouldSwap(isASC, s.items[leftIndex], s.items[endIndex]) {
				break
			}
			leftIndex++
		}
		for { // 从右查找比middle值 小的元素的下表
			if leftIndex >= rightIndex || shouldSwap(isASC, s.items[endIndex], s.items[rightIndex]) {
				break
			}
			rightIndex--
		}
		if leftIndex >= rightIndex { // 如果两个下表相遇，跳出
			break
		}
		s.swap(leftIndex, rightIndex) // 交换元素
	}
	if shouldSwap(isASC, s.items[leftIndex], s.items[midIndex]) { // 此时left应该等于right
		s.swap(leftIndex, midIndex)
		midIndex = leftIndex
	}
	s.quickSortRecursive(isASC, startIndex, midIndex-1)
	s.quickSortRecursive(isASC, midIndex+1, endIndex)
}

func (s *Sorter) swap(index int, index2 int) {
	if index >= len(s.items) || index < 0 || index2 >= len(s.items) || index2 < 0 {
		log.Fatal("swap: index error:", index, index2)
	}
	temp := s.items[index]
	s.items[index] = s.items[index2]
	s.items[index2] = temp
}

func (s *Sorter) MergeSort(isASC bool) {
	start := 0
	end := len(s.items) - 1
	s.items = s.mergeSortRecursive(isASC, start, end)
}

func (s *Sorter) mergeSortRecursive(isASC bool, start int, end int) []SortItem {
	if start > end {
		return []SortItem{}
	}
	if start == end {
		return []SortItem{s.items[start]}
	}

	midIndex := (start + end) / 2
	sortItemsA := s.mergeSortRecursive(isASC, start, midIndex)
	sortItemsB := s.mergeSortRecursive(isASC, midIndex+1, end)
	return mergeItemsList(isASC, sortItemsA, sortItemsB)
}

func (s *Sorter) HeapSort(isASC bool) {
	// 数组存放堆的结构，左节点： 2i+1，右节点：2i+2
	endIndex := len(s.items)
	for index := endIndex - 1; index >= 0; index -= 2 {
		position := (index - 1) / 2
		s.heapBuildRecursive(!isASC, position, endIndex)
	}

	for endIndex > 0 {
		endIndex -= 1
		s.swap(endIndex, 0)
		s.heapBuildRecursive(!isASC, 0, endIndex)
	}
}

func (s *Sorter) heapGetFirstSubNodeIndex(isASC bool, position int, endIndex int) int {
	leftIndex := 2*position + 1
	rightIndex := 2*position + 2
	first := position

	if leftIndex >= endIndex { // 不存在子节点，返回本身
		return first
	}
	first = leftIndex
	if rightIndex >= endIndex { // 不存在右节点时,返回左节点的值
		return first
	}

	if shouldSwap(isASC, s.items[leftIndex], s.items[rightIndex]) {
		first = rightIndex
	}
	return first
}

func (s *Sorter) heapBuildRecursive(isASC bool, position int, endIndex int) {
	subIndex := s.heapGetFirstSubNodeIndex(isASC, position, endIndex)
	shouldSwap := shouldSwap(isASC, s.items[position], s.items[subIndex])
	if shouldSwap {
		s.swap(position, subIndex)
		s.heapBuildRecursive(isASC, subIndex, endIndex)
	}
}

func mergeItemsList(isASC bool, sortItemsA, sortItemsB []SortItem) []SortItem {
	items := make([]SortItem, len(sortItemsA)+len(sortItemsB))
	indexA, indexB, indexItems := 0, 0, 0
	for {
		if indexA >= len(sortItemsA) && indexB >= len(sortItemsB) {
			break
		}
		if indexA >= len(sortItemsA) {
			items[indexItems] = sortItemsB[indexB]
			indexB++
			indexItems++
			continue
		}
		if indexB >= len(sortItemsB) {
			items[indexItems] = sortItemsA[indexA]
			indexA++
			indexItems++
			continue
		}
		if shouldSwap(isASC, sortItemsA[indexA], sortItemsB[indexB]) {
			items[indexItems] = sortItemsB[indexB]
			indexB++
			indexItems++
			continue
		}
		items[indexItems] = sortItemsA[indexA]
		indexA++
		indexItems++
	}
	return items
}
