package main

import (
	"fmt"
)

type Interval struct {
	Start int
	End   int
}

func (i Interval) more(other Interval) bool {
	return i.Start >= other.End
}

func (i Interval) overlap(other Interval) bool {
	if i.Start < other.End && i.End > other.Start {
		return true
	}
	return false
}

type MyCalendar struct {
	Intervals []Interval
}

func Constructor() MyCalendar {
	return MyCalendar{Intervals: []Interval{}}
}

func (this *MyCalendar) Book(start int, end int) bool {
	insertInterval := Interval{
		Start: start,
		End:   end,
	}

	i := 0
	for i < len(this.Intervals) {
		if insertInterval.overlap(this.Intervals[i]) {
			return false
		} else if insertInterval.more(this.Intervals[i]) {
			i++
		} else {
			break
		}
	}
	before := this.Intervals[:i]
	after := append([]Interval{}, this.Intervals[i:]...)
	this.Intervals = append(before, insertInterval)
	this.Intervals = append(this.Intervals, after...)
	return true
}

func (this *MyCalendar) Debug() {
	for _, interval := range this.Intervals {
		fmt.Printf("[%d,%d] ", interval.Start, interval.End)
	}
	fmt.Println()
}

func main() {
	calendar := Constructor()
	fmt.Println(calendar.Book(37, 50))
	calendar.Debug()
	fmt.Println(calendar.Book(33, 50))
	calendar.Debug()
	fmt.Println(calendar.Book(4, 17))
	calendar.Debug()
	fmt.Println(calendar.Book(35, 48))
	calendar.Debug()
	fmt.Println(calendar.Book(8, 25))
	calendar.Debug()
}
