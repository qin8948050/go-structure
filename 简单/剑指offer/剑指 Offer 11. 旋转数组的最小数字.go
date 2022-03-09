/*
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。  

示例 1：

输入：[3,4,5,1,2]
输出：1
示例 2：

输入：[2,2,2,0,1]
输出：0
 */

 /*
 解题思路

使用二分法找最小值，先确定target
☆☆left: 因为递增数组，mid>left或者mid<left都无法确定一个明确的情况。
right: mid<right时，最小值一定在左区间且包含mid;mid>right时，最小值一定在右区间且不包含mid。mid=right时，根据递增数组的性质找下一个right
*/

func minArray(numbers []int) int {
    left, right := 0, len(numbers) - 1
    for left <= right {
        mid := (left + right) / 2 
        if  numbers[mid] > numbers[right] {//说明最小值在右边
            left = mid + 1
        }else {//反之则在左边
            right = right - 1
        }
    }
    return numbers[left]
}

