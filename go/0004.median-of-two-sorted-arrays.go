func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	l := len1 + len2
	mid := (l / 2) + 1
	merged := make([]int, mid)
	n1, n2 := 0, 0
	for i := 0; i < mid; i++ {
		if n1 >= len1 {
			merged[i] = nums2[n2]
			n2++
		} else if n2 >= len2 {
			merged[i] = nums1[n1]
			n1++
		} else if nums2[n2] > nums1[n1] {
			merged[i] = nums1[n1]
			n1++
		} else {
			merged[i] = nums2[n2]
			n2++
		}
	}

	if mid == 1 {
		return float64(merged[0])
	}
	if l%2 == 1 {
		return float64(merged[mid-1])
	}
	return float64(merged[mid-1]+merged[mid-2]) / 2.0
}
