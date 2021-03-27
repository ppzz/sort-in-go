package Checker

type Checker struct {
    lastVal int
    asc     bool
    ascFlag bool
    counter int
}

func NewChecker() *Checker {
    c := new(Checker)
    c.lastVal = 0
    c.counter = 0
    return c
}

func (c *Checker) IsOrdered(item int) bool {
    if c.counter > 0 && !c.ascFlag && item != c.lastVal { // 确定已有数字的方向
        c.asc = item > c.lastVal
        c.ascFlag = true
    }
    r := c.counter == 0 || !c.ascFlag || c.lastVal == item || c.lastVal < item && c.asc || c.lastVal > item && !c.asc
    c.lastVal = item
    c.counter++
    return r
}

func (c *Checker) setAsc(item int) {
    if item > c.lastVal {
        c.asc = true
    }
}
