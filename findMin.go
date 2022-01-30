package lesson5

//FindMin: 寻找旋转排序数组中的最小值 II
/*parameter
nums: 有序有重复数组旋转未知次数后的数组
return: 数组中的最小值
*/
func FindMin(nums []int) int {

	left, right := 0, len(nums)-1

	for left < right {

		mid := left + (right-left)>>1
		if nums[mid] > nums[right] {

			left = mid + 1
		} else if nums[mid] < nums[right] {

			right = mid
		} else { //因数组元素可能有重复，所以需判断相等条件，相等则缩减查找的右边界

			right--
		}
	}

	return left
}
