////////////////////////////////////////////////////////////////////////////////
// name: myFunctions
//
//
////////////////////////////////////////////////////////////////////////////////

package myFunctions

import (
	"fmt"
	"sort"
)

//
func IntSliceDelFunc(inSlice []int, i int) []int {
	return append(inSlice[:i], inSlice[i+1:]...)
}

//
func IntSliceDelByPositionMultiFunc(inSlice []int,  idxs ...int) []int {
	var idxSlice []int

	for i, v := range idxs {
		idxSlice = append(idxSlice, v - i)
	}
	sort.Ints(idxSlice)

	for _, v2 := range idxSlice {
		inSlice = append(inSlice[:v2], inSlice[(v2+1):]...)
	}

	return inSlice
}


//
func IntFindSliceElementPositionFunc(inSlice []int, toFindValue int) int {
	idx := 0

	for j, v := range inSlice {
		if v == toFindValue {
			idx = j
			break
		}
	}

	fmt.Println("found element", toFindValue, "at", idx, "position")
	return idx
}

//
func IntSliceSliceWithLengthFunc(inSlice []int, startPosition int, length int) []int {
	return inSlice[startPosition:(startPosition + length)]
}

//
func IntSliceElementExistFunc(inSlice []int, inValue int) (bool, []int) {

	var idxSlice []int
	var valueExist bool

	for i, v := range inSlice {
		if v == inValue {
			idxSlice = append(idxSlice, i)
			valueExist = true
		}
	}

	return valueExist, idxSlice
}


//
func IntSliceDeduplicateFunc(inSlice []int) []int {
	//fmt.Println("original slice", inSlice)

	for _, v := range inSlice {
		valueExist, valueIdxSlice := IntSliceElementExistFunc(inSlice, v)
		if valueExist && len(valueIdxSlice) > 1 {
			inSlice = IntSliceDelByPositionMultiFunc(inSlice, valueIdxSlice[1:]...)
		}
	}
	return inSlice
}


//
func IntSlicesDiffFunc(inSlice1 []int, inSlice2 []int) []int {
	inSlice1 = IntSliceDeduplicateFunc(inSlice1)
	inSlice2 = IntSliceDeduplicateFunc(inSlice2)

	var idxSlice []int
	for i, v := range inSlice1 {
		existBool, _ := IntSliceElementExistFunc(inSlice2, v)

		if existBool {
			idxSlice = append(idxSlice, i)
		}
	}

	inSlice1 = IntSliceDelByPositionMultiFunc(inSlice1, idxSlice...)

	return inSlice1
}
