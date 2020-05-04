package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type iterationResult struct {
	voteCount uint
	candidate string
}


// runTallyIteration counts votes per candidate for this iteration based on the current
// ranked voter.votes and returns a map of candidate name to vote count.
func countWinningVotes(voters []voter) (candidateToVoteCount map[string]uint) {
	candidateToVoteCount = make(map[string]uint)
	for i, _ := range(voters) {
		// Check if there are any candidates left to count from this voter.
		if len(voters[i].votes) < 1 {
			continue
		}
		candidateToVoteCount[voters[i].votes[0]] += 1
	}
	return candidateToVoteCount
}


func sortIterationResults(candidateToVoteCount map[string]uint) (sortedResults []iterationResult) {
	for candidate, voteCount := range(candidateToVoteCount) {
		sortedResults = append(
			sortedResults,
			iterationResult{
				candidate: candidate,
				voteCount: voteCount,
			},
		)
	}

	// Sort on candidate name to make results consistent regardless of map iteration order.
	sort.SliceStable(sortedResults, func(i, j int) bool {
		return sortedResults[i].candidate < sortedResults[j].candidate
	})
	// Sort by actual vote count.
	sort.SliceStable(sortedResults, func(i, j int) bool {
		return sortedResults[i].voteCount < sortedResults[j].voteCount
	})

	return sortedResults
}


func computeInstantRunoffWinner(voters []voter) (winningCandidates []string, err error) {
	// TODO FIXME do the whole appropriate bit in a loop.

	candidateToVoteCount := countWinningVotes(voters)
	sortedResults := sortIterationResults(candidateToVoteCount)

	fmt.Println("Results of iteration 1:")
	for _, sr := range(sortedResults) {
		fmt.Printf("\t%s: %d\n", sr.candidate, sr.voteCount)
	}

	// Collect all candidate(s) with the most and least votes first choice votes.
	minFirstChoiceVotes := sortedResults[0].voteCount
	maxFirstChoiceVotes := sortedResults[len(sortedResults)-1].voteCount
	var worstCandidates []iterationResult
	var bestCandidates []iterationResult
	for _, res := range(sortedResults) {
		if res.voteCount == minFirstChoiceVotes {
			worstCandidates = append(worstCandidates, res)
		}
		if res.voteCount == maxFirstChoiceVotes {
			bestCandidates = append(bestCandidates, res)
		}
	}

	fmt.Println("The winner(s) is/are:")
	for _, sr := range(bestCandidates) {
		fmt.Printf("\t%s: %d\n", sr.candidate, sr.voteCount)
	}
	if float64(maxFirstChoiceVotes) > float64(len(voters)) {
		fmt.Printf("%d votes is a simple majority of %d voters, so we have (a) winner(s).\n", maxFirstChoiceVotes, len(voters))
		for _, res := range(bestCandidates) {
			winningCandidates = append(winningCandidates, res.candidate)
		}
		return winningCandidates, nil
	} else {
		fmt.Printf("%d votes is not a simple majority of %d voters; beginning next iteration.\n", maxFirstChoiceVotes, len(voters))
		// TODO then we go to the next iteration.
		// TODO pick one of the losers at random to eliminate, elimiate it from all
		// voters for which it was the top candidate, and continue to the next iteration.
		fmt.Println("The candidate(s) with the least first choice votes is/are:")
		for _, sr := range(worstCandidates) {
			fmt.Printf("\t%s: %d\n", sr.candidate, sr.voteCount)
		}
		if len(worstCandidates) < 1 {
			panic("No losing candidates. This doesn't make sense.")
		}
		losingCandidate := worstCandidates[rand.Intn(len(worstCandidates))].candidate
		fmt.Printf("Randomly selected candidate to eliminate: %s\n", losingCandidate)

		// Eliminate losing candidate.
		for i, _ := range(voters) {
			// Check if there are any candidates left to count from this voter.
			if len(voters[i].votes) < 1 {
				continue
			}

			if voters[i].votes[0] == losingCandidate {
				// We have to do this slice manipulation on voters[i] because that rewrites
				// the actual in the slice. If we range over `_, v` and rewrite the loop
				// variable's `v.votes`, it won't mutate the actual voter.
				voters[i].votes = voters[i].votes[1:]
			}
		}
	}

	return []string{}, fmt.Errorf("Something went wrong.")
}

