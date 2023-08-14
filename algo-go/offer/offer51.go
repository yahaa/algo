package offer

func reversePairs2(nums []int) int {
	return mergeSort2(nums, 0, len(nums)-1)
}

func mergeSort2(nums []int, left, right int) int {
	if left >= right {
		return 0
	}

	mid := (left + right) / 2

	count := mergeSort2(nums, left, mid)    // left --> mid 已经排好序了
	count += mergeSort2(nums, mid+1, right) // mid+1 --> right 已经排好序了

	// do merge 合并两个有序数组

	i, j, k := left, mid+1, 0

	tmp := make([]int, right-left+1)

	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
			count += mid - i + 1
		}

		k++
	}

	for i <= mid {
		tmp[k] = nums[i]
		i++
		k++
	}

	for j <= right {
		tmp[k] = nums[j]
		j++
		k++
	}

	k = 0
	for left <= right {
		nums[left] = tmp[k]
		left++
		k++
	}
	return count
}

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

// leetcode 912 归并排序实现数组排序
func sortArray(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	mergeSort912(nums, 0, len(nums)-1)

	return nums
}

//
func mergeSort912(nums []int, left, right int) {
	if left >= right {
		return
	}

	mid := (left + right) / 2

	mergeSort912(nums, left, mid)    // left --> mid 已经排好序
	mergeSort912(nums, mid+1, right) // mid+1 --> right 已经排好序

	tmp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0

	for i <= mid && j <= right {
		if nums[i] < nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}

		k++
	}

	for i <= mid {
		tmp[k] = nums[i]
		i++
		k++
	}

	for j <= right {
		tmp[k] = nums[j]
		j++
		k++
	}

	k = 0

	for i := left; i <= right; i++ {
		nums[i] = tmp[k]
		k++
	}
}
