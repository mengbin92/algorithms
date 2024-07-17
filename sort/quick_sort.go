package sort

import "cmp"

// 快速排序（QuickSort）是一种高效的排序算法，通常用于大型数据集。它使用分治法（Divide and Conquer）将一个数组分成较小的子数组，然后递归地排序这些子数组。其核心思想是通过一个称为“基准”（pivot）的元素，将数组分成两部分，其中一部分元素小于基准，另一部分元素大于基准。

// ### 快速排序的步骤

// 1. **选择基准**：从数组中选择一个元素作为基准。基准的选择方法有多种，可以是数组的第一个元素、最后一个元素、中间元素，或随机选择。
// 2. **划分数组**：将数组重新排序，使得所有比基准小的元素放在基准的左边，所有比基准大的元素放在基准的右边（相等的元素可以放在任意一边）。此过程称为划分（Partition）。
// 3. **递归排序**：递归地对划分后的两个子数组进行快速排序。

// ### 快速排序的复杂度

// - **平均时间复杂度**：O(n log n)
// - **最坏时间复杂度**：O(n^2)（当每次选择的基准都是最大或最小元素时）
// - **空间复杂度**：O(log n)（由于递归调用栈）

func QuickSort[T cmp.Ordered](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}

	// 选择基准
	pivot := arr[len(arr)/2]

	// 划分数组
	var left, right []T
	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		}
	}

	// 递归调用
	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}
