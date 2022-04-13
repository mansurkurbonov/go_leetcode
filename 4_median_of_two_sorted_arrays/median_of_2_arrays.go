package medianoftwosortedarrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)

	nums3 := make([]int, len1+len2)
	mergeInts(nums1, nums2, nums3, len1, len2)

	if len(nums3)%2 == 0 {
		a := len(nums3) / 2
		a--
		b := a + 1
		c := (float64(nums3[a]) + float64(nums3[b]))
		return c / 2
	} else {
		a := len(nums3) / 2
		return float64(nums3[a])
	}

}

func mergeInts(nums1, nums2, nums3 []int, len1, len2 int) {
	i := 0
	j := 0
	k := 0

	for i < len1 && j < len2 {
		if nums1[i] < nums2[j] {
			nums3[k] = nums1[i]
			k++
			i++
		} else {
			nums3[k] = nums2[j]
			k++
			j++
		}
	}

	for i < len1 {
		nums3[k] = nums1[i]
		k++
		i++
	}

	for j < len2 {
		nums3[k] = nums2[j]
		k++
		j++
	}
}
