func binSearch(T tabInt, n int, x int) int {
    var left, right, mid, idx int

    left = 0
    right = n - 1
    idx = -1

    // Perulangan selama range ada dan x belum ditemukan (idx masih -1)
    for left <= right && idx == -1 {
        mid = (left + right) / 2
        if x < T[mid] {
            right = mid - 1
        } else if x > T[mid] {
            left = mid + 1
        } else {
            // Jika x == T[mid], maka idx diisi dengan posisi mid
            idx = mid
        }
    }

    return idx
}