package offer

func reversePairs(nums []int) int {
	tmp := make([]int, len(nums))
	return mergeSort(nums, tmp, 0, len(nums)-1)
}

func mergeSort(a, tmp []int, left, right int) int {
	if left >= right {
		return 0
	}

	mid := (left + right) / 2

	count := mergeSort(a, tmp, left, mid) + mergeSort(a, tmp, mid+1, right)

	// do merge

	i, j, t := left, mid+1, 0

	for i <= mid && j <= right {
		if a[i] <= a[j] {
			tmp[t] = a[i]
			i++
			count += j - (mid + 1)
		} else {
			tmp[t] = a[j]
			j++
		}
		t++
	}

	for i <= mid {
		tmp[t] = a[i]
		i++
		t++

		count += j - (mid + 1)
	}

	for j <= right {
		tmp[t] = a[j]
		j++
		t++

	}

	t = 0
	for left <= right {
		a[left] = tmp[t]
		left++
		t++
	}

	return count
}
