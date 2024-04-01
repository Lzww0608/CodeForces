func nextGreatestLetter(letters []byte, target byte) byte {
    l, r := 0, len(letters)
    for l < r {
        mid := l + ((r - l) >> 1)
        if letters[mid] <= target {
            l = mid + 1
        } else {
            r = mid 
        }
    }
    if r < len(letters) {
        return letters[r]
    }
    return letters[0]
}