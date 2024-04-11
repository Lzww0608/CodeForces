func isBoomerang(points [][]int) bool {
    a, b := points[0][0], points[0][1]
    c, d := points[1][0], points[1][1]
    e, f := points[2][0], points[2][1]
    
    if (f - d) * (c - a) == (d - b) * (e - c) {
        return false
    }

    return true
}