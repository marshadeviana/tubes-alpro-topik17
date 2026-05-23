func binSearch(T tabInt, n int, x int) bool {
    var left, right, mid int
	
    left = 0
    right = n - 1
    mid = (left + right) / 2

    for left <= right && T[mid] != x {
        if x < T[mid] {
            right = mid - 1
        } else {
            left = mid + 1
        }
        mid = (left + right) / 2
    }

    return mid > -1 && T[mid] == x
}