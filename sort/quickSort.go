/*
quichSot worst: O(n^2), expect:O(nlogn)
think about T(n) = T(9n/10) + T(n/10) + O(n)
                n
          9n/10  n/10                                      n
    81n/100 9n/100 9n/100 n/100                            n
log10(n)               ...                                 n
                       ...                                 <=n
log10/9(n)             ...                                 <=n


so T(n) <= nlog10(n) + n(log10/9(n) - log10(n)) = nlog10/9(n) = O(nlogn)
 */

package sort

import "math/rand"

func partition (arr []int)(primeIdx int) {
	primeIdx = 0
	for i := 0; i < len(arr) - 1; i ++ {
		if arr[i] < arr[len(arr) - 1] {
			arr[i], arr[primeIdx] = arr[primeIdx], arr[i]
			primeIdx++
		}
	}
	arr[primeIdx], arr[len(arr)-1] = arr[len(arr)-1], arr[primeIdx]
	return
}

func quickSort(arr []int) {
	if len(arr) > 1 {
		primeIdx := partition(arr)
		if primeIdx < len(arr)/2 {
			quickSort(arr[:primeIdx])
			quickSort(arr[primeIdx+1:])
		} else {
			quickSort(arr[primeIdx+1:])
			quickSort(arr[:primeIdx])
		}
	}
}

func randomQuickSort(arr []int) {
	if len(arr) > 1 {
		primeIdx := rand.Intn(len(arr))
		arr[primeIdx], arr[len(arr)-1] = arr[len(arr)-1], arr[primeIdx]
		primeIdx = partition(arr)
		if primeIdx < len(arr)/2 {
			randomQuickSort(arr[:primeIdx])
			randomQuickSort(arr[primeIdx+1:])
		} else {
			randomQuickSort(arr[primeIdx+1:])
			randomQuickSort(arr[:primeIdx])
		}
	}
}