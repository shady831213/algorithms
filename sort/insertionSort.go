package sort

func insertionSort(arr []int) {
	for i := 1 ; i < len(arr); i++ {
		for j := i; j > 0 ; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}

func bubbleSort(arr []int) {
	for i :=range arr {
		for j := i; j < len(arr) ; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}