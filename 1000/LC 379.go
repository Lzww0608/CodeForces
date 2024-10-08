type PhoneDirectory struct {
    a []int
    used []bool
    id int
}


func Constructor(maxNumbers int) PhoneDirectory {
    a := make([]int, maxNumbers)
    used := make([]bool, maxNumbers)
    for i := range a {
        a[i] = i
    }
    return PhoneDirectory{a, used, 0}
}


func (c *PhoneDirectory) Get() int {
    if c.id == len(c.a) {
        return -1
    }
    c.used[c.a[c.id]] = true
    c.id++
    return c.a[c.id-1]
}


func (c *PhoneDirectory) Check(number int) bool {
    return !c.used[number]
}


func (c *PhoneDirectory) Release(number int)  {
    if c.used[number] {
        c.id--
        c.a[c.id] = number
        c.used[number] = false
    }
    
}


/**
 * Your PhoneDirectory object will be instantiated and called as such:
 * obj := Constructor(maxNumbers);
 * param_1 := obj.Get();
 * param_2 := obj.Check(number);
 * obj.Release(number);
 */