package Checker

import "testing"

func TestChecker_IsOrdered(t *testing.T) {
    testCases := []struct {
        name   string
        items  []int
        wanted []bool
    }{
        {
            name:   "case1: 连续多个item相等",
            items:  []int{1, 1, 2, 3, 4},
            wanted: []bool{true, true, true, true, true,},
        },
        {
            name:   "case2: asc,ordered",
            items:  []int{1, 2, 2, 3, 4},
            wanted: []bool{true, true, true, true, true,},
        },
        {
            name:   "case3: asc, not ordered",
            items:  []int{1, 2, 6, 5, 6},
            wanted: []bool{true, true, true, false, true,},
        },
        {
            name:   "case4: desc, ordered",
            items:  []int{6, 6, 6, 4, 3},
            wanted: []bool{true, true, true, true, true,},
        },
        {
            name:   "case5:desc,not_ordered",
            items:  []int{6, 6, 5, 3, 4, 3},
            wanted: []bool{true, true, true, true, false, true,},
        },
    }
    for _, testCase := range testCases {
        t.Run(testCase.name, func(t *testing.T) {
            checker := NewChecker()
            for i := range testCase.items {
                got := checker.IsOrdered(testCase.items[i])
                if got != testCase.wanted[i] {
                    t.Errorf("IsOrdered() = %v, want %v, i = %d", got, testCase.wanted[i], i)
                }
            }
        })
    }
}
