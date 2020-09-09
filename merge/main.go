package sort

// MergeSort 归并排序算法
func MergeSort(a []int) {
	tmp := make([]int, len(a))
	mergeSort(a, tmp, 0, len(a)-1)
}

func merge(a, tmp []int, left, mid, right int) {
	i, j, index := left, mid+1, 0

	for i <= mid && j <= right {
		if a[i] <= a[j] {
			tmp[index] = a[i]
			i++
		} else {
			tmp[index] = a[j]
			j++
		}
		index++
	}

	for i <= mid {
		tmp[index] = a[i]
		i++
		index++
	}

	for j <= right {
		tmp[index] = a[j]
		j++
		index++
	}

	index = 0
	for left <= right {
		a[left] = tmp[index]
		left++
		index++
	}
}

func mergeSort(a, tmp []int, left, right int) {
	if left >= right {
		return
	}

	mid := (left + right) / 2

	mergeSort(a, tmp, left, mid)
	mergeSort(a, tmp, mid+1, right)

	// do merge
	merge(a, tmp, left, mid, right)
}
