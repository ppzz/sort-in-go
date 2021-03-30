package sorter

import "fmt"

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
	fmt.Println(s.items)
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
