package dp

import (
	"strings"
	"math"
	"sort"
)

func prettyPrint(words string, lineCap int) (alignIdx int, alignedWords string) {
	wordsSlice := strings.Split(words," ")
	//side data structure, store enter position and aligneIdx(blankcnt^3)
	alignIdxArray := make([][]int, len(wordsSlice), cap(wordsSlice))
	enterPosArray := make([][]int, len(wordsSlice), cap(wordsSlice))
	for i := range alignIdxArray {
		alignIdxArray[i] = make([]int, len(wordsSlice), cap(wordsSlice))
		enterPosArray[i] = make([]int, len(wordsSlice), cap(wordsSlice))
	}

	//get blank cnt, if is negtive, the line is exceeded
	getBlankCnt := func (start, end int) (int) {
		charCnt := 0
		for i := start; i <= end; i++ {
			charCnt += len(wordsSlice[i])
		}
		return lineCap - end + start - charCnt
	}
	//length from 1 to max length
	for length := 1;length <= len(wordsSlice); length++ {
		//start position and end position, keek length and from 0 to max value
		for start := 0; start <= len(wordsSlice) - length; start++ {
			end := start + length - 1
			blankCnt := getBlankCnt(start,end)
			//exceed 1 line and spilt to 2 sub problem
			if blankCnt < 0 {
				alignIdxArray[start][end] = math.MaxInt32
				for k := start; k<end; k++ {
					temp := alignIdxArray[start][k] + alignIdxArray[k+1][end]
					if temp < alignIdxArray[start][end] {
						alignIdxArray[start][end] = temp
						enterPosArray[start][end] = k
					}
				}
			} else {
				//not exceed 1 line, so get the end of line
				alignIdxArray[start][end] = int(math.Pow(float64(blankCnt), 3))
				enterPosArray[start][end] = end
			}
		}
	}
	//results
	alignIdx = alignIdxArray[0][len(wordsSlice)-1]
	alignedWords = ""
	//get all enter positions
	enterPos := make([]int ,0 ,0)
	var getEnterPos func (start, end int)
	getEnterPos = func (start, end int){
		if getBlankCnt(start,end) < 0 {
			enter := enterPosArray[start][end]
			enterPos = append(enterPos, enter)
			getEnterPos(start, enter)
			getEnterPos(enter+1, end)
		}
	}
	getEnterPos(0, len(wordsSlice)-1)
	sort.Ints(enterPos)
	//output strings, insert enter in corresponding position
	for i,v := range wordsSlice {
		alignedWords += v
		if len(enterPos) !=0 && i == enterPos[0] {
			alignedWords += "\n"
			enterPos = enterPos[1:]
		} else {
			alignedWords += " "
		}
	}

	return
}