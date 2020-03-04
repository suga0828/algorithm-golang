package bubblesort

// BubbleSort implements bubble sort algorithm
func BubbleSort(d []int, sArr []int) []int {
	l := len(d)
	if len(d) == 0 {
		if len(sArr) > 0 {
			return sArr
		}
		return d
	}
	i := 1
	for i < len(d) {
		if d[i-1] > d[i] {
			min := d[i]
			max := d[i-1]
			d[i-1] = min
			d[i] = max
		}
		i++
	}
	sArr = append(d[len(d)-1], ...sArr)
	d = d[:len(d)-1]
	if len(d) > 0 {
		sArr = BubbleSort(d, sArr)
	}
	if len(sArr) == l {
		sArr = reverseArr(sArr)
	}
	return sArr
}
