package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	nums []int
	m    map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		nums: make([]int, 0),
		m:    make(map[int]int),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}

	//insert
	this.m[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; !ok {
		return false
	}

	//remove
	index := this.m[val]
	last := len(this.nums) - 1
	if index != last {
		//swap index and last
		this.m[this.nums[last]] = index
		this.nums[index], this.nums[last] = this.nums[last], this.nums[index]
	}
	this.nums = this.nums[:last]
	delete(this.m, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	size := len(this.nums)
	if size > 0 {
		index := rand.Intn(size)
		return this.nums[index]
	}
	return 0
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
func main() {
	obj := Constructor()
	param_1 := obj.Insert(1)
	fmt.Println(obj)
	param_2 := obj.Remove(1)
	param_3 := obj.GetRandom()
	fmt.Printf("param_1: %v, param_2: %v, param_3: %v\n", param_1, param_2, param_3)
}
