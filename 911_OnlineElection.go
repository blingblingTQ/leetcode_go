package main

type Leader struct {
	Person int
	Time   int
}

type TopVotedCandidate struct {
	Leaders    []Leader
	VoteCounts map[int]int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	leaders := []Leader{}
	voteCounts := make(map[int]int)
	for i := 0; i < len(persons); i++ {
		p := persons[i]
		voteCounts[p]++
		if len(leaders) > 0 {
			last := leaders[len(leaders)-1]
			if voteCounts[last.Person] > voteCounts[p] {
				p = last.Person
			}
		}
		leaders = append(leaders, Leader{Person: p, Time: times[i]})
	}
	return TopVotedCandidate{
		Leaders:    leaders,
		VoteCounts: voteCounts,
	}
}

func (this *TopVotedCandidate) Q(t int) int {
	start, end := 0, len(this.Leaders)-1
	for start <= end {
		mid := start + (end-start)/2
		if this.Leaders[mid].Time <= t {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return this.Leaders[start-1].Person
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
