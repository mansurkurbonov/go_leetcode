package removeelement

func removeElement(nums []int, val int) int {
	x := 0
	l := len(nums)

	for i := 0; i < l; i++ {
		if nums[x] == val {
			nums = remove(nums, x)
			x = x - 1
		}
		x++
	}

	return len(nums)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
