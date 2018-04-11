
package sort
/*
merge sort O(nlgn):
T(n) = 2T(n/2) + O(n)
master theorem:
a = 2, b = 2, f(n) = n
logb(a) = lg2 = 1 f(n) = f(n^logb(a)) = f(n^1)
so, O(n) = O(n^logb(a)lgn) = O(nlgn)
*/
import (
	"sync"
)

func merge(arr []int) {
	i := len(arr) / 2
	//copy left and right array
	leftArr, rightArr := make([]int, i, i), make([]int, len(arr)-i, len(arr)-i)
	copy(leftArr, arr[:i])
	copy(rightArr, arr[i:])
	leftIter, rightIter := ints(leftArr).Iter(), ints(rightArr).Iter()
	leftValue, leftHasNext := leftIter()
	rightValue, rightHasNext := rightIter()
	//merge
	for k := range arr {
		if !leftHasNext { //left empty, use right value, in CLRS, use infinity
			arr[k] = rightValue
			rightValue, rightHasNext = rightIter()
		} else if !rightHasNext { //right empty, use left value, in CLRS, use infinity
			arr[k] = leftValue
			leftValue, leftHasNext = leftIter()
		} else {
			if leftValue > rightValue {
				arr[k] = rightValue
				rightValue, rightHasNext = rightIter()
			} else {
				arr[k] = leftValue
				leftValue, leftHasNext = leftIter()
			}
		}
	}
}

func mergeSort(arr []int) {
	i := len(arr) / 2
	if i > 0 {
		mergeSort(arr[:i])
		mergeSort(arr[i:])
		merge(arr)
	}
}

func mergeSortParallel(arr []int) {
	i := len(arr) / 2
	if i > 0 {
		var wd sync.WaitGroup
		wd.Add(2)
		go func() {
			mergeSortParallel(arr[:i])
			wd.Done()
		}()
		go func() {
			mergeSortParallel(arr[i:])
			wd.Done()
		}()
		wd.Wait()
		merge(arr)
	}
}
