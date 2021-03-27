package sorter

import (
	Checker "github.com/ppzz/sorting-go/internal/checker"
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

func (s *Sorter) BubbleSort(isASC bool) {
	for i := 0; i < len(s.items); i++ {
		for j := i + 1; j < len(s.items); j++ {
			if isASC && s.items[i].great(&s.items[j]) || !isASC && s.items[j].great(&s.items[i]) {
				var t SortItem = s.items[i]
				s.items[i] = s.items[j]
				s.items[j] = t
			}
		}
	}
}

func (s *Sorter) check() (bool, bool) {
	c := Checker.NewChecker()
	isOrdered := true
	for _, sortItem := range s.items {
		if !c.IsOrdered(sortItem.Val) {
			isOrdered = false
			break
		}
	}
	return isOrdered, c.GetIsAsc()
}
