package sorter

import "testing"

func TestSorter_BubbleSort(t *testing.T) {
	tests := []struct {
		name           string
		list           []SortItem
		want           bool
		wantedIsSorted bool
		wantedIsASC    bool
	}{
		{
			name: "case1:",
			list: []SortItem{
				{Seq: 1, Val: 7377893500238068435},
				{Seq: 2, Val: 3870084184907221662},
				{Seq: 3, Val: 5816097024187652061},
				{Seq: 4, Val: 9079073840597114598},
				{Seq: 5, Val: 3209982839515082167},
				{Seq: 6, Val: 955314993070246892},
				{Seq: 7, Val: 8579404511464342498},
				{Seq: 8, Val: 8885012081506670610},
				{Seq: 9, Val: 5351312676265698788},
				{Seq: 10, Val: 5351312676265698789},
			},
			wantedIsSorted: true,
			wantedIsASC:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sorter{
				items: tt.list,
			}
			s.BubbleSort(tt.wantedIsASC)
			gotIsSorted, gotIsASC := s.check()
			if !gotIsSorted {
				t.Errorf("BubbleSort() gotIsSorted = %v, wantedIsSorted %v", gotIsSorted, tt.wantedIsSorted)
			}
			if gotIsASC != tt.wantedIsASC {
				t.Errorf("BubbleSort() gotIsASC = %v, wantedIsASC %v", gotIsSorted, tt.wantedIsASC)
			}
		})
	}
}

func TestSorter_check(t *testing.T) {
	tests := []struct {
		name          string
		items         []SortItem
		wantIsOrdered bool
		wantIsAsc     bool
	}{
		{
			name: "case1: ordered, asc",
			items: []SortItem{
				{Seq: 1, Val: 737789350023806},
				{Seq: 2, Val: 3870084184907222},
				{Seq: 3, Val: 58160970241876521},
				{Seq: 4, Val: 907907384059711498},
				{Seq: 5, Val: 3209982839515082167},
			},
			wantIsOrdered: true,
			wantIsAsc:     true,
		},
		{
			name: "case2: ordered, desc",
			items: []SortItem{
				{Seq: 1, Val: 3209982839515082167},
				{Seq: 2, Val: 907907384059711498},
				{Seq: 3, Val: 58160970241876521},
				{Seq: 4, Val: 3870084184907222},
				{Seq: 5, Val: 737789350023806},
			},
			wantIsOrdered: true,
			wantIsAsc:     false,
		},
		{
			name: "case3: not ordered, start with asc",
			items: []SortItem{
				{Seq: 1, Val: 737789350023806},
				{Seq: 2, Val: 3870084184907222},
				{Seq: 3, Val: 907907384059711498},
				{Seq: 4, Val: 58160970241876521},
				{Seq: 5, Val: 3209982839515082167},
			},
			wantIsOrdered: false,
			wantIsAsc:     true,
		},
		{
			name: "case4: not ordered, start with desc",
			items: []SortItem{
				{Seq: 1, Val: 3209982839515082167},
				{Seq: 2, Val: 3209982839515082167},
				{Seq: 3, Val: 3870084184907222},
				{Seq: 4, Val: 58160970241876521},
				{Seq: 5, Val: 737789350023806},
			},
			wantIsOrdered: false,
			wantIsAsc:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sorter{
				items: tt.items,
			}
			gotIsOrdered, gotIsASC := s.check()
			if gotIsOrdered != tt.wantIsOrdered {
				t.Errorf("check() gotIsOrdered = %v, wantIsOrdered %v", gotIsOrdered, tt.wantIsOrdered)
			}
			if gotIsASC != tt.wantIsAsc {
				t.Errorf("check() gotIsASC = %v, wantIsAsc %v", gotIsASC, tt.wantIsAsc)
			}
		})
	}
}
