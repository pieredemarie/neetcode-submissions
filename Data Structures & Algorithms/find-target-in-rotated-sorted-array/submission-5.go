func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2

		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	pivot := l // this is our minimun (pivot)
	if pivot == 0 {
		return BinarySearchHelper(nums,0, len(nums)-1,target)
	}

	if target >= nums[0] && target <= nums[pivot-1] {
		return BinarySearchHelper(nums, 0, pivot-1, target)
	} else {
		return BinarySearchHelper(nums, pivot, len(nums)-1, target)
	}
}

func BinarySearchHelper(nums []int, l, r, target int) int {
	    for l <= r {
        mid := l + (r-l)/2
        
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {  
            l = mid + 1
        } else {                         
            r = mid - 1
        }
    }
    return -1
}
