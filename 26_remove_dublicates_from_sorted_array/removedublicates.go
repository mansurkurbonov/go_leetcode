package removedublicatesfromsortedarray

func removeDuplicates(nums []int) int {
	var max struct {
		val   int
		count int
	}

	x := 0
	l := len(nums)

	for i := 0; i < l; i++ {
		if nums[x] == max.val && max.count > 0 {
			max.val = nums[x]
			max.count = max.count + 1
			nums = remove(nums, x)
			x = x - 1
		} else {
			max.val = nums[x]
			max.count = 1
		}
		x++
	}

	return len(nums)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
