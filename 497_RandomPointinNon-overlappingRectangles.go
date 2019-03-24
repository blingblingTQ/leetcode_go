package main

type Solution struct {
	cumulateAreas []int
	rects         [][]int
}

func getArea(rect []int) int {
	x1, y1, x2, y2 := rect[0], rect[1], rect[2], rect[3]
	return (x2 - x1 + 1) * (y2 - y1 + 1)
}

func Constructor(rects [][]int) Solution {
	solution := Solution{
		cumulateAreas: make([]int, len(rects)),
		rects:         make([][]int, len(rects)),
	}
	preSum := 0
	for i, rect := range rects {
		solution.rects[i] = rects[i]
		area := getArea(rect)
		solution.cumulateAreas[i] = area + preSum
		preSum = solution.cumulateAreas[i]
	}
	return solution
}

func (this *Solution) Pick() []int {
	target := rand.Intn(this.cumulateAreas[len(this.cumulateAreas)-1]) + 1
	start, end := 0, len(this.cumulateAreas)-1
	for start < end {
		mid := start + (end-start)/2
		if this.cumulateAreas[mid] < target {
			start = mid + 1
		} else {
			end = mid
		}
	}

	//从下标是start的area中随机选一个点
	rect := this.rects[start]
	x1, y1, x2, y2 := rect[0], rect[1], rect[2], rect[3]
	x := rand.Intn(x2-x1+1) + x1
	y := rand.Intn(y2-y1+1) + y1
	return []int{x, y}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */
