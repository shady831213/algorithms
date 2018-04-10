package sort

func countingSort(arr []int, bias int) (retArr []int) {
	countingArr := make([]int, bias+1, bias+1)
	retArr = make([]int, len(arr), cap(arr))
	for _, v := range arr {
		countingArr[v]++
	}
	for i := 1; i < len(countingArr); i++ {
		countingArr[i] += countingArr[i-1]
	}
	for i := len(arr) - 1; i >= 0; i-- {
		retArr[countingArr[arr[i]]-1] = arr[i]
		countingArr[arr[i]]--
	}
	return
}
