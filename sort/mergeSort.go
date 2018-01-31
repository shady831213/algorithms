package sort

func mergeSort(arr []int) {
	i := len(arr) / 2
	if i > 0 {
		mergeSort(arr[:i])
		mergeSort(arr[i:])
		func () {
			//copy left and right array
			leftArr, rightArr := make([]int,i,i), make([]int, len(arr) - i , len(arr) - i )
			copy(leftArr,arr[:i])
			copy(rightArr,arr[i:])
			//get first value
			leftIter, rightIter := Ints(leftArr).Iter(),Ints(rightArr).Iter()
			leftValue,leftHasNext:= leftIter()
			rightValue,rightHasNext := rightIter()
			//merge
			for k := range arr {
				if !leftHasNext {//left empty, use right value, in CLRS, use infinity
					arr[k] = rightValue
					rightValue,rightHasNext = rightIter()
				} else if !rightHasNext {//right empty, use left value, in CLRS, use infinity
					arr[k] = leftValue
					leftValue,leftHasNext = leftIter()
				} else {
					if leftValue > rightValue {
						arr[k] = rightValue
						rightValue,rightHasNext = rightIter()
					} else {
						arr[k] = leftValue
						leftValue,leftHasNext = leftIter()
					}
				}
			}
		}()
	}
}
