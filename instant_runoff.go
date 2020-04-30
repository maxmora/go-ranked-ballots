package main

import (
	"fmt"
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


// TODO add a verbose flag to print everything?
func computeInstantRunoffWinner(voters []voter) (winningCandidate string, err error) {
	candidateToVoteCount := countWinningVotes(voters)

	sortedResults := sortIterationResults(candidateToVoteCount)
	fmt.Println(sortedResults)

	// TODO check if we have a majority on the winner.
	// TODO what if there is more than one?
	// TODO return the winner(s) if so.
	// TODO determine losers, eliminate them, and iterate if not.

	// TODO Determine how to determine this.
	losingCandidate := "IMPLEMENT ME"

	// Eliminate losing candidates.
	for i, _ := range(voters) {
		// Check if there are any candidates left to count from this voter.
		if len(voters[i].votes) < 1 {
			continue
		}

		// FIXME this may have to eliminate more than one, depending on tie-breaking rules.
		if voters[i].votes[0] == losingCandidate {
			// We have to do this slice manipulation on voters[i] because that rewrites
			// the actual in the slice. If we range over `_, v` and rewrite the loop
			// variable's `v.votes`, it won't mutate the actual voter.
			voters[i].votes = voters[i].votes[1:]
		}
	}
	// TODO then iterate and slice off losing candidates
	// TODO determine how to determine that!
	fmt.Println(candidateToVoteCount)
	return
}

