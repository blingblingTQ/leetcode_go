package main

type Solution struct {
	cululation []int
}

func Constructor(w []int) Solution {
	s := Solution{
		cululation: make([]int, len(w)),
	}
	for i := 0; i < len(w); i++ {
		pre := 0
		if i > 0 {
			pre = s.cululation[i-1]
		}
		s.cululation[i] = pre + w[i]
	}
	return s
}

func (this *Solution) PickIndex() int {
	target := rand.Intn(this.cululation[len(this.cululation)-1]) + 1
	start, end := 0, len(this.cululation)-1
	for start < end {
		mid := start + (end-start)/2
		if this.cululation[mid] < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
