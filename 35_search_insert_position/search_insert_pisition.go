package searchinsertposition

func searchInsert(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		if nums[0] > target {
			return 0
		} else {
			return 1
		}
	}
	return search(nums, target)
}

func search(nums []int, target int) int {
	var mid int

	minIndex := 0
	maxIndex := len(nums) - 1

	for minIndex <= maxIndex {
		mid = (minIndex + maxIndex) / 2

		if nums[mid] == target {
			return mid
		}

		if mid+1 < len(nums) {
			if target > nums[mid] && target < nums[mid+1] {
				return mid + 1
			}
		} else {
			return mid + 1
		}

		if mid > 0 {
			if target < nums[mid] && target > nums[mid-1] {
				return mid
			}
		} else {
			if target < nums[mid] {
				return mid
			}
		}

		if nums[mid] < target {
			minIndex = mid + 1
			continue
		}

		if nums[mid] > target {
			maxIndex = mid - 1
			continue
		}
	}
	return 0

}
