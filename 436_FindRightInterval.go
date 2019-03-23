package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

type IntervalSlice []Interval

func (i IntervalSlice) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func (i IntervalSlice) Less(x, y int) bool {
	return i[x].Start < i[y].Start
}

func (i IntervalSlice) Len() int {
	return len(i)
}

//寻找第一个Start大于等于target的interval
func binarySearch(intervals []Interval, target int) int {
	start, end := 0, len(intervals)-1
	for start < end {
		mid := start + (end-start)/2
		if intervals[mid].Start < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

func findRightInterval(intervals []Interval) []int {
	m := make(map[Interval]int)
	for i, interval := range intervals {
		m[interval] = i
	}

	sort.Sort(IntervalSlice(intervals))
	ans := make([]int, len(intervals))
	for _, interval := range intervals {
		index := binarySearch(intervals, interval.End)
		fmt.Println(index)
		if index == len(intervals) || intervals[index].Start < interval.End {
			ans[m[interval]] = -1
		} else {
			ans[m[interval]] = m[intervals[index]]
		}
	}
	return ans
}

func main() {
	intervals := []Interval{Interval{Start: 1, End: 2}}
	fmt.Println(findRightInterval(intervals))
}
