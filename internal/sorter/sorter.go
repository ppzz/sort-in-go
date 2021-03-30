package sorter

import (
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
			if s.shouldSwap(isASC, s.items[i], s.items[j]) {
				var t = s.items[i]
				s.items[i] = s.items[j]
				s.items[j] = t
			}
		}
	}
}

func (s *Sorter) shouldSwap(isASC bool, item1 SortItem, item2 SortItem) bool {
	return isASC && item1.great(&item2) || !isASC && item2.great(&item1)
}

// 插入排序
func (s *Sorter) InsertionSort(isASC bool) {
	for outIndex := 1; outIndex < len(s.items); outIndex++ {
		temp := s.items[outIndex]                   // 用来暂存外层循环的一个值，
		internalIndex := outIndex - 1               // 初始化内循环的起始位置，为当前位置的前一个值
		for ; internalIndex >= 0; internalIndex-- { // 内循环，用来跟暂存的值比较，并且移动位置
			shouldSwap := s.shouldSwap(isASC, s.items[internalIndex], temp) // 内循环值跟temp比较，判断是否到了temp应该的位置
			if shouldSwap {                                                 // 不是temp的位置，将内循环的值向后移动一个位置
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
					shouldSwap := s.shouldSwap(isASC, s.items[internalIndex], temp) // 内循环值跟temp比较，判断是否到了temp应该的位置
					if shouldSwap {                                                 // 不是temp的位置，将内循环的值向后移动一个位置
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
