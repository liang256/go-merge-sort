package main

import "fmt"

func main() {

	sli := []int{3, 50, 6, 7, 1, 8, 2}

	mergeSort(sli)
	fmt.Println(sli)
}

func mergeSort(sli []int) {
	if len(sli) < 2 {
		return
	}
	mergeSortRecur(sli, 0, len(sli))
}

func mergeSortRecur(sli []int, start_i int, end_i int) []int {
	fmt.Println("sli:", sli[start_i:end_i])
	if start_i == end_i-1 {
		return sli[start_i:end_i]
	}

	leftSli := mergeSortRecur(sli, start_i, (start_i+end_i)/2)
	rightSli := mergeSortRecur(sli, (start_i+end_i)/2, end_i)
	fmt.Println("left:", leftSli, "right:", rightSli)
	sort(sli, leftSli, rightSli, start_i, end_i)
	return sli[start_i:end_i]
}

func sort(sli []int, left []int, right []int, start_i int, end_i int) {
	tmp := make([]int, end_i-start_i)
	m_i := 0
	l_i := 0
	r_i := 0
	for m_i < end_i-start_i {
		if l_i == len(left) {
			// assign the rest right to tmp and end
			tmp = append(tmp[:m_i], right[r_i:]...)
			break
		} else if r_i == len(right) {
			// assign the rest left to tmp and end
			tmp = append(tmp[:m_i], left[l_i:]...)
			break
		}
		if left[l_i] <= right[r_i] {
			// assgin curr left to tmp
			tmp[m_i] = left[l_i]
			l_i++
		} else {
			// assign curr right to tmp
			tmp[m_i] = right[r_i]
			r_i++
		}
		m_i++
	}
	fmt.Println("sorted tmp", tmp)
	tmp = append(sli[:start_i], tmp[:]...)
	sli = append(tmp, sli[end_i:]...)
}
