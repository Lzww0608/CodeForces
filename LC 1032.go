type StreamChecker struct {
    isEnd bool
    next [26]*StreamChecker
}

var s []byte

func Constructor(words []string) StreamChecker {
    s = []byte{}
    root := &StreamChecker{}
    for _, s := range words {
        cur := root
        for i := len(s) - 1; i >= 0; i-- {
            if cur.next[int(s[i] - 'a')] == nil {
                cur.next[int(s[i] - 'a')] = &StreamChecker{}
            }
            cur = cur.next[int(s[i] - 'a')]
        }
        cur.isEnd = true
    }
    return *root
}



func (this *StreamChecker) Query(letter byte) bool {
    s = append(s, letter)
    n := len(s) - 1
    for this != nil && n >= 0 {
        this = this.next[int(s[n]) - 'a']

        if this != nil && this.isEnd == true {
            return true
        }
        n--
    }
    return false
}


/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */