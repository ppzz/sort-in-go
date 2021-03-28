package sorter

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
	for i := 0; i < len(s.items); i++ {
		for j := i - 1; j >= 0; j-- {
			shouldSwap := s.shouldSwap(isASC, s.items[j], s.items[i])
			if shouldSwap {
				j--
				continue
			}
			if !shouldSwap && j == i-1 {
				break
			}
			var t = s.items[i]
			s.items[i] = s.items[j]
			s.items[j] = t
			break
		}
	}
}

func (s *Sorter) check() (bool, bool) {
	c := NewChecker()
	isOrdered := true
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
