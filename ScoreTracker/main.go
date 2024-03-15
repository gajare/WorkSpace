package main

import (
	"fmt"
	"sort"
)

// ScoreTracker interface defines methods for tracking scores
type ScoreTracker interface {
	AddScore(score int)
	HighScore() int
	AverageScore() float64
	FetchAllScores() []int
	MostCommonScore() (int, int)
}

// scoreTrackerImpl implements the ScoreTracker interface
type scoreTrackerImpl struct {
	scores []int
}

// NewScoreTracker creates a new instance of ScoreTracker
func NewScoreTracker() ScoreTracker {
	return &scoreTrackerImpl{}
}

func (st *scoreTrackerImpl) AddScore(score int) {
	st.scores = append(st.scores, score)
}

func (st *scoreTrackerImpl) HighScore() int {
	if len(st.scores) == 0 {
		return 0
	}
	max := st.scores[0]
	for _, score := range st.scores {
		if score > max {
			max = score
		}
	}
	return max
}

func (st *scoreTrackerImpl) AverageScore() float64 {
	if len(st.scores) == 0 {
		return 0
	}
	sum := 0
	for _, score := range st.scores {
		sum += score
	}
	return float64(sum) / float64(len(st.scores))
}

func (st *scoreTrackerImpl) FetchAllScores() []int {
	return st.scores
}

func (st *scoreTrackerImpl) MostCommonScore() (int, int) {
	if len(st.scores) == 0 {
		return 0, 0
	}
	scoreMap := make(map[int]int)
	for _, score := range st.scores {
		scoreMap[score]++
	}
	var mostCommonScore, count int
	for score, freq := range scoreMap {
		if freq > count {
			mostCommonScore = score
			count = freq
		}
	}
	return mostCommonScore, count
}

func main() {
	tracker := NewScoreTracker()
	tracker.AddScore(10)
	tracker.AddScore(20)
	tracker.AddScore(30)
	tracker.AddScore(10)
	tracker.AddScore(20)

	fmt.Println("High Score:", tracker.HighScore())
	fmt.Println("Average Score:", tracker.AverageScore())
	fmt.Println("All Scores:", tracker.FetchAllScores())
	mostCommonScore, count := tracker.MostCommonScore()
	fmt.Printf("Most Common Score: %d (appeared %d times)\n", mostCommonScore, count)
}
