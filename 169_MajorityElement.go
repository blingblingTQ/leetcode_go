package main

//HashMap
//O(n)时间复杂度
//O(n)空间复杂度
func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
		if m[num] > len(nums)/2 {
			return num
		}
	}
	return -1
}

//partition
//O(n)时间复杂度
//O(1)空间复杂度
class Solution {
public:
    
    void swap(vector<int>& nums, int a, int b) {
        int tmp = nums[a];
        nums[a] = nums[b];
        nums[b] = tmp;
    }
    
    int partition(vector<int>&nums, int lo, int hi) {
        int i = lo, j = hi + 1;
        int guard = nums[lo];
        while(true) {
            while(nums[++i] < guard) {
                if(i == hi)
                    break;
            }
            while(guard < nums[--j]) {
                if(j == lo)
                    break;
            }
            if (i >= j) {
                break;
            }
            swap(nums, i, j);
        }
        swap(nums, lo, j);
        return j;
    }
    
    int majorityElement(vector<int>& nums) {
        int lo = 0, hi = nums.size()-1;
        int k = nums.size()/2;
        while(lo < hi) {
            int j = partition(nums, lo, hi);
            if (j == k) {
                return nums[k];
            } else if (j < k) {
                lo = j + 1;
            } else {
                hi = j - 1;
            }
        }
        return nums[lo];
    }
};


//摩尔选举算法
//O(n)时间复杂度
//O(1)空间复杂度
func majorityElement(nums []int) int {
    candidate := 0
    count := 0
    for i := 0; i < len(nums); i++ {
        if count == 0 {
            candidate = nums[i]
            count ++
        } else if (nums[i] == candidate) {
            count ++
        } else {
            count --
        }
    }
    return candidate
}
