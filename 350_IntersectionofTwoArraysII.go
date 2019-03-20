package main

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	res := make([]int, 0)
	for _, num := range nums1 {
		m[num]++
	}

	for _, num := range nums2 {
		if m[num] > 0 {
			m[num]--
			res = append(res, num)
		}
	}
	return res
}

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	res := make([]int, 0)
	p1, p2 := 0, 0
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] == nums2[p2] {
			res = append(res, nums1[p1])
			p1++
			p2++
		} else if nums1[p1] > nums2[p2] {
			p2++
		} else {
			p1++
		}
	}
	return res
}
