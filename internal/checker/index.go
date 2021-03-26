package Checker

type Checker struct {
    lastVal int
    asc     bool
}

func (c *Checker) IsOrdered(item int) bool {
    return c.lastVal == item || c.lastVal < item && c.asc ||  c.lastVal > item && !c.asc
}
