package towptr

func merge(nums1 []int, m int, nums2 []int, n int) {
	index := m + n - 1

	k, j := m-1, n-1

	for k >= 0 && j >= 0 {
		if nums1[k] > nums2[j] {
			nums1[index] = nums1[k]
			k--
		} else {
			nums1[index] = nums2[j]
			j--
		}
		index--
	}

	for j >= 0 {
		nums1[index] = nums2[j]
		index--
		j--
	}
}
