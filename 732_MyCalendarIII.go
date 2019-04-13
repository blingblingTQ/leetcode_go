package main

import "fmt"
import "sort"

type MyCalendarThree struct {
	Delta map[int]int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{
		Delta: make(map[int]int),
	}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	this.Delta[start]++
	this.Delta[end]--

	keys := []int{}
	for k := range this.Delta {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	active := 0
	ans := 0
	for _, k := range keys {
		active += this.Delta[k]
		if active > ans {
			ans = active
		}
	}
	return ans
}

func (this *MyCalendarThree) Debug() {
	fmt.Printf("%v\n", this.Delta)
}

func main() {
	calendar := Constructor()
	fmt.Println(calendar.Book(10, 20))
	calendar.Debug()
	fmt.Println(calendar.Book(50, 60))
	calendar.Debug()
	fmt.Println(calendar.Book(10, 40))
	calendar.Debug()
	fmt.Println(calendar.Book(5, 15))
	calendar.Debug()
	fmt.Println(calendar.Book(5, 10))
	calendar.Debug()
	fmt.Println(calendar.Book(25, 55))
	calendar.Debug()
}
