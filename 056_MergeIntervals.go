package main

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

type Intervals []Interval

func (i Intervals) Len() int {
	return len(i)
}
func (i Intervals) Less(a, b int) bool {
	return i[a].Start < i[b].Start
}
func (i Intervals) Swap(a, b int) {
	i[a], i[b] = i[b], i[a]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals []Interval) []Interval {
	result := make([]Interval, 0)
	if len(intervals) == 0 {
		return result
	}
	sort.Sort(Intervals(intervals))
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := len(result) - 1
		if intervals[i].Start > result[last].End {
			result = append(result, intervals[i])
		} else {
			result[last].End = max(result[last].End, intervals[i].End)
		}
	}
	return result
}
