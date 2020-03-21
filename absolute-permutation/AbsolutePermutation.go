package absolutepermutation

// AbsolutePermutation return the lexicographically smallest
// absolute permutation. If no absolute permutation exists, print -1.
func AbsolutePermutation(n int32, k int32) []int32 {
	arr := make([]int32, n)
Loop:
	for i := range arr {
		a, b := int32(i)-k+1, int32(i)+k+1
		if a > 0 && arr[a-1] == 0 {
			arr[a-1] = int32(i + 1)
		} else if b <= n && arr[b-1] == 0 {
			arr[b-1] = int32(i + 1)
		} else {
			arr = []int32{-1}
			break Loop
		}
	}
	return arr
}
