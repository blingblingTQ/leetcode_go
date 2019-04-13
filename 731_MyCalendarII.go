package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Interval struct {
	Start int
	End   int
}

func (i Interval) Overlap(other Interval) bool {
	if i.Start < other.End && i.End > other.Start {
		return true
	}
	return false
}

type MyCalendarTwo struct {
	Overlaps []Interval
	Calendar []Interval
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		Overlaps: []Interval{},
		Calendar: []Interval{},
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	insertInterval := Interval{
		Start: start,
		End:   end,
	}
	for _, iv := range this.Overlaps {
		if iv.Overlap(insertInterval) {
			return false
		}
	}

	for _, iv := range this.Calendar {
		if iv.Overlap(insertInterval) {
			newOverlap := Interval{
				Start: max(start, iv.Start),
				End:   min(end, iv.End),
			}
			this.Overlaps = append(this.Overlaps, newOverlap)
		}
	}
	this.Calendar = append(this.Calendar, insertInterval)
	return true
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
