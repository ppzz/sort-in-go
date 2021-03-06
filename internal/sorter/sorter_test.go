package sorter

import (
	"reflect"
	"testing"
)

var sortItems1 = []SortItem{
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
}
var sortItems2 = []SortItem{
	{Seq: 1, Val: 73},
	{Seq: 2, Val: 78},
	{Seq: 3, Val: 58},
	{Seq: 4, Val: 90},
	{Seq: 5, Val: 32},
	{Seq: 6, Val: 37},
	{Seq: 7, Val: 38},
	{Seq: 8, Val: 83},
	{Seq: 9, Val: 9},
	{Seq: 10, Val: 21},
	{Seq: 11, Val: 20},
}
var sortItems3 = []SortItem{
	{Seq: 1, Val: 2222},
	{Seq: 2, Val: 333333},
	{Seq: 3, Val: 11},
}

func TestSorter_BubbleSort(t *testing.T) {
	items1 := make([]SortItem, len(sortItems1))
	items2 := make([]SortItem, len(sortItems2))
	items3 := make([]SortItem, len(sortItems3))
	copy(items1, sortItems1)
	copy(items2, sortItems2)
	copy(items3, sortItems3)

	testCases := []struct {
		name           string
		list           []SortItem
		want           bool
		wantedIsSorted bool
		wantedIsASC    bool
	}{
		{
			name:           "case1:",
			list:           items1,
			wantedIsSorted: true,
			wantedIsASC:    true,
		},
		{
			name:           "case2:",
			list:           items2,
			wantedIsSorted: true,
			wantedIsASC:    true,
		},
		{
			name:           "case3:",
			list:           items3,
			wantedIsSorted: true,
			wantedIsASC:    true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			s := &Sorter{
				items: testCase.list,
			}
			s.BubbleSort(testCase.wantedIsASC)
			gotIsSorted, gotIsASC := s.Check()
			if !gotIsSorted {
				t.Errorf("BubbleSort() gotIsSorted = %v, wantedIsSorted %v", gotIsSorted, testCase.wantedIsSorted)
			}
			if gotIsASC != testCase.wantedIsASC {
				t.Errorf("BubbleSort() gotIsASC = %v, wantedIsASC %v", gotIsSorted, testCase.wantedIsASC)
			}
		})
	}
}

func TestSorter_InsertionSort(t *testing.T) {
	items1 := make([]SortItem, len(sortItems1))
	items2 := make([]SortItem, len(sortItems2))
	copy(items1, sortItems1)
	copy(items2, sortItems2)

	testCases := []struct {
		name           string
		list           []SortItem
		want           bool
		wantedIsSorted bool
		wantedIsASC    bool
	}{
		{
			name:           "case1:",
			list:           items1,
			wantedIsSorted: true,
			wantedIsASC:    true,
		},
		{
			name:           "case2: itemsSimple",
			list:           items2,
			wantedIsSorted: true,
			wantedIsASC:    true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			s := &Sorter{
				items: testCase.list,
			}
			s.InsertionSort(testCase.wantedIsASC)
			gotIsSorted, gotIsASC := s.Check()
			if !gotIsSorted {
				t.Errorf("InsertionSort() gotIsSorted = %v, wantedIsSorted %v", gotIsSorted, testCase.wantedIsSorted)
			}
			if gotIsASC != testCase.wantedIsASC {
				t.Errorf("InsertionSort() gotIsASC = %v, wantedIsASC %v", gotIsSorted, testCase.wantedIsASC)
			}
		})
	}
}

func TestSorter_ShellSort(t *testing.T) {
	items1 := make([]SortItem, len(sortItems1))
	items2 := make([]SortItem, len(sortItems2))
	copy(items1, sortItems1)
	copy(items2, sortItems2)

	testCases := []struct {
		name           string
		list           []SortItem
		want           bool
		wantedIsSorted bool
		wantedIsASC    bool
	}{
		{
			name:           "case1:",
			list:           items1,
			wantedIsSorted: true,
			wantedIsASC:    true,
		},
		{
			name:           "case2: itemsSimple",
			list:           items2,
			wantedIsSorted: true,
			wantedIsASC:    true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			s := &Sorter{
				items: testCase.list,
			}
			s.ShellSort(testCase.wantedIsASC)
			gotIsSorted, gotIsASC := s.Check()
			if !gotIsSorted {
				t.Errorf("ShellSort() gotIsSorted = %v, wantedIsSorted %v", gotIsSorted, testCase.wantedIsSorted)
			}
			if gotIsASC != testCase.wantedIsASC {
				t.Errorf("ShellSort() gotIsASC = %v, wantedIsASC %v", gotIsSorted, testCase.wantedIsASC)
			}
		})
	}
}

func TestSorter_SelectionSort(t *testing.T) {
	items1 := make([]SortItem, len(sortItems1))
	items2 := make([]SortItem, len(sortItems2))
	copy(items1, sortItems1)
	copy(items2, sortItems2)

	testCases := []struct {
		name           string
		list           []SortItem
		want           bool
		wantedIsSorted bool
		wantedIsASC    bool
	}{
		{
			name:           "case1:",
			list:           items1,
			wantedIsSorted: true,
			wantedIsASC:    true,
		},
		{
			name:           "case2: itemsSimple",
			list:           items2,
			wantedIsSorted: true,
			wantedIsASC:    true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			s := &Sorter{
				items: testCase.list,
			}
			s.SelectionSort(testCase.wantedIsASC)
			gotIsSorted, gotIsASC := s.Check()
			if !gotIsSorted {
				t.Errorf("SelectionSort() gotIsSorted = %v, wantedIsSorted %v", gotIsSorted, testCase.wantedIsSorted)
			}
			if gotIsASC != testCase.wantedIsASC {
				t.Errorf("SelectionSort() gotIsASC = %v, wantedIsASC %v", gotIsSorted, testCase.wantedIsASC)
			}
		})
	}
}

func TestSorter_QuickSort(t *testing.T) {
	items1 := make([]SortItem, len(sortItems1))
	items2 := make([]SortItem, len(sortItems2))
	copy(items1, sortItems1)
	copy(items2, sortItems2)

	testCases := []struct {
		name           string
		list           []SortItem
		want           bool
		wantedIsSorted bool
		wantedIsASC    bool
	}{
		{
			name:           "case1:",
			list:           items1,
			wantedIsSorted: true,
			wantedIsASC:    true,
		},
		{
			name:           "case2: itemsSimple",
			list:           items2,
			wantedIsSorted: true,
			wantedIsASC:    true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			s := &Sorter{
				items: testCase.list,
			}
			s.QuickSort(testCase.wantedIsASC)
			gotIsSorted, gotIsASC := s.Check()
			if !gotIsSorted {
				t.Errorf("QuickSort() gotIsSorted = %v, wantedIsSorted %v", gotIsSorted, testCase.wantedIsSorted)
			}
			if gotIsASC != testCase.wantedIsASC {
				t.Errorf("QuickSort() gotIsASC = %v, wantedIsASC %v", gotIsSorted, testCase.wantedIsASC)
			}
		})
	}
}

func TestSorter_MergeSort(t *testing.T) {
	items1 := make([]SortItem, len(sortItems1))
	items2 := make([]SortItem, len(sortItems2))
	copy(items1, sortItems1)
	copy(items2, sortItems2)

	testCases := []struct {
		name           string
		list           []SortItem
		want           bool
		wantedIsSorted bool
		wantedIsASC    bool
	}{
		{
			name:           "case1:",
			list:           items1,
			wantedIsSorted: true,
			wantedIsASC:    true,
		},
		{
			name:           "case2: itemsSimple",
			list:           items2,
			wantedIsSorted: true,
			wantedIsASC:    true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			s := &Sorter{
				items: testCase.list,
			}
			s.MergeSort(testCase.wantedIsASC)
			gotIsSorted, gotIsASC := s.Check()
			if !gotIsSorted {
				t.Errorf("MergeSort() gotIsSorted = %v, wantedIsSorted %v", gotIsSorted, testCase.wantedIsSorted)
			}
			if gotIsASC != testCase.wantedIsASC {
				t.Errorf("MergeSort() gotIsASC = %v, wantedIsASC %v", gotIsSorted, testCase.wantedIsASC)
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
			gotIsOrdered, gotIsASC := s.Check()
			if gotIsOrdered != tt.wantIsOrdered {
				t.Errorf("Check() gotIsOrdered = %v, wantIsOrdered %v", gotIsOrdered, tt.wantIsOrdered)
			}
			if gotIsASC != tt.wantIsAsc {
				t.Errorf("Check() gotIsASC = %v, wantIsAsc %v", gotIsASC, tt.wantIsAsc)
			}
		})
	}
}

func TestSorter_shouldSwap(t *testing.T) {
	type args struct {
		isASC bool
		item1 SortItem
		item2 SortItem
	}
	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1: asc, 12, 12",
			args: args{
				isASC: true,
				item1: SortItem{Seq: 1, Val: 12},
				item2: SortItem{Seq: 2, Val: 12},
			},
			want: false,
		},
		{
			name: "case2: desc, 12, 12",
			args: args{
				isASC: false,
				item1: SortItem{Seq: 1, Val: 12},
				item2: SortItem{Seq: 2, Val: 12},
			},
			want: false,
		},
		{
			name: "case3: asc, 12, 18",
			args: args{
				isASC: true,
				item1: SortItem{Seq: 1, Val: 12},
				item2: SortItem{Seq: 2, Val: 18},
			},
			want: false,
		},
		{
			name: "case4: asc, 18, 12",
			args: args{
				isASC: true,
				item1: SortItem{Seq: 1, Val: 18},
				item2: SortItem{Seq: 2, Val: 12},
			},
			want: true,
		},
		{
			name: "case5: desc, 12, 18",
			args: args{
				isASC: false,
				item1: SortItem{Seq: 1, Val: 12},
				item2: SortItem{Seq: 2, Val: 18},
			},
			want: true,
		},
		{
			name: "case6: desc, 18, 12",
			args: args{
				isASC: false,
				item1: SortItem{Seq: 1, Val: 18},
				item2: SortItem{Seq: 2, Val: 12},
			},
			want: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := shouldSwap(testCase.args.isASC, testCase.args.item1, testCase.args.item2)
			if got != testCase.want {
				t.Errorf("shouldSwap() = %v, want %v", got, testCase.want)
			}
		})
	}
}

func Test_mergeItemsList(t *testing.T) {
	type args struct {
		isASC      bool
		sortItemsA []SortItem
		sortItemsB []SortItem
	}
	testCases := []struct {
		name string
		args args
		want []SortItem
	}{
		{
			name: "case1:",
			args: args{
				isASC: true,
				sortItemsA: []SortItem{
					{Seq: 1, Val: 10},
					{Seq: 2, Val: 21},
					{Seq: 3, Val: 37},
					{Seq: 4, Val: 58},
					{Seq: 5, Val: 78},
					{Seq: 6, Val: 90},
				},
				sortItemsB: []SortItem{
					{Seq: 1, Val: 9},
					{Seq: 2, Val: 20},
					{Seq: 3, Val: 32},
					{Seq: 4, Val: 38},
					{Seq: 5, Val: 73},
					{Seq: 6, Val: 83},
				},
			},
			want: []SortItem{
				{Seq: 1, Val: 9},
				{Seq: 1, Val: 10},
				{Seq: 2, Val: 20},
				{Seq: 2, Val: 21},
				{Seq: 3, Val: 32},
				{Seq: 3, Val: 37},
				{Seq: 4, Val: 38},
				{Seq: 4, Val: 58},
				{Seq: 5, Val: 73},
				{Seq: 5, Val: 78},
				{Seq: 6, Val: 83},
				{Seq: 6, Val: 90},
			},
		},
		{
			name: "case2:",
			args: args{
				isASC: true,
				sortItemsA: []SortItem{
					{Seq: 1, Val: 9},
					{Seq: 1, Val: 10},
					{Seq: 2, Val: 20},
					{Seq: 2, Val: 21},
					{Seq: 3, Val: 32},
					{Seq: 4, Val: 58},
					{Seq: 6, Val: 83},
				},
				sortItemsB: []SortItem{
					{Seq: 3, Val: 37},
					{Seq: 4, Val: 38},
					{Seq: 5, Val: 73},
					{Seq: 5, Val: 78},
					{Seq: 6, Val: 90},
				},
			},
			want: []SortItem{
				{Seq: 1, Val: 9},
				{Seq: 1, Val: 10},
				{Seq: 2, Val: 20},
				{Seq: 2, Val: 21},
				{Seq: 3, Val: 32},
				{Seq: 3, Val: 37},
				{Seq: 4, Val: 38},
				{Seq: 4, Val: 58},
				{Seq: 5, Val: 73},
				{Seq: 5, Val: 78},
				{Seq: 6, Val: 83},
				{Seq: 6, Val: 90},
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := mergeItemsList(testCase.args.isASC, testCase.args.sortItemsA, testCase.args.sortItemsB)

			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("mergeItemsList() = %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestSorter_HeapSort(t *testing.T) {
	items1 := make([]SortItem, len(sortItems1))
	items2 := make([]SortItem, len(sortItems2))
	copy(items1, sortItems1)
	copy(items2, sortItems2)

	testCases := []struct {
		name           string
		items          []SortItem
		isASC          bool
		position       int
		wantedIsSorted bool
		wantedIsASC    bool
	}{
		{
			name: "case1:",
			items: []SortItem{
				{Seq: 1, Val: 1},
				{Seq: 2, Val: 2},
				{Seq: 3, Val: 3},
			},
			isASC:          false,
			position:       0,
			wantedIsSorted: true,
			wantedIsASC:    false,
		},
		{
			name: "case2:",
			items: []SortItem{
				{Seq: 1, Val: 1},
				{Seq: 2, Val: 2},
				{Seq: 3, Val: 3},
				{Seq: 4, Val: 4},
				{Seq: 5, Val: 5},
				{Seq: 6, Val: 6},
				{Seq: 7, Val: 7},
			},
			isASC:          false,
			wantedIsSorted: true,
			wantedIsASC:    false,
		},
		{
			name:           "case3:",
			items:          items1,
			isASC:          true,
			wantedIsSorted: true,
			wantedIsASC:    true,
		},
		{
			name:           "case4:",
			items:          items2,
			isASC:          false,
			wantedIsSorted: true,
			wantedIsASC:    false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			s := &Sorter{
				items: testCase.items,
			}
			s.HeapSort(testCase.isASC)
			gotIsSorted, gotIsASC := s.Check()
			if !gotIsSorted {
				t.Errorf("HeapSort() gotIsSorted = %v, wantedIsSorted %v", gotIsSorted, testCase.wantedIsSorted)
			}
			if gotIsASC != testCase.wantedIsASC {
				t.Errorf("HeapSort() gotIsASC = %v, wantedIsASC %v", gotIsSorted, testCase.wantedIsASC)
			}
		})
	}
}

func TestSorter_heapGetFirstSubNodeIndex(t *testing.T) {
	testCases := []struct {
		name     string
		items    []SortItem
		isASC    bool
		position int
		want     int
	}{
		{
			name: "case1",
			items: []SortItem{
				{Seq: 1, Val: 1},
				{Seq: 2, Val: 2},
				{Seq: 3, Val: 3},
			},
			isASC:    false,
			position: 0,
			want:     2,
		},
		{
			name: "case2",
			items: []SortItem{
				{Seq: 1, Val: 1},
				{Seq: 2, Val: 2},
				{Seq: 3, Val: 3},
				{Seq: 4, Val: 5},
				{Seq: 5, Val: 4},
				{Seq: 6, Val: 3},
			},
			isASC:    false,
			position: 1,
			want:     3,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			s := &Sorter{
				items: testCase.items,
			}
			got := s.heapGetFirstSubNodeIndex(testCase.isASC, testCase.position, len(s.items))
			if got != testCase.want {
				t.Errorf("heapGetFirstSubNodeIndex() = %v, want %v", got, testCase.want)
			}
		})
	}
}
