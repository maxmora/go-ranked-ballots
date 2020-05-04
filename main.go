package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)


type voter struct {
	timestamp string
	// Candidate names from most preferred (highest-ranked) to least preferred.
	votes []string
}


// indexOfRankingInVoteSlice returns the index of the first instance of ranking
// in voteSlice. -1 if ranking is not found.
//
// In output CSVs, multiple rankings can be given to the same candidate. When
// that occurs, the rankings are semicolon-separated, e.g., a candidate with
// just a 1 will have a vote "1", but a candidate with rankings 1 and 2 will
// have "1;2".
//
// To handle this, we split vote fields on semicolons and only consider the
// first (i.e., the highest) ranking for that candidate. In this way, a
// candidate can only get a single vote, and it will only count when we check
// for the highgest ranking that it has a vote for.
func indexOfRankingInVoteSlice(voteSlice []string, ranking string) (pos int) {
	for i, s := range(voteSlice) {
		if s == ranking {
			return i
		}

		voteParts := strings.Split(s, ";")
		if (len(voteParts) > 0) && (voteParts[0] == ranking) {
			return i
		}
	}
	return -1
}


// possibleRankings is a slice of possible ranking values, from worst to best.
func tabulateVoters(csvRecords [][]string, possibleRankings []string) (voters []voter, err error) {
	err = validateCSVRecords(csvRecords)
	if err != nil {
		return []voter{}, err
	}

	// Just the candidates (CSV headers excluding the first column, which is "Timestamp").
	candidates := csvRecords[0][1:]

	for _, voterSlice := range(csvRecords[1:]) {
		// for each possible ranking (best to worst), get its index and use that to grab the candidate
		// append the candidate to a slice.
		var votesBestToWorst []string
		for _, ranking := range(possibleRankings) {
			// Get the offset (ignoring the timestamp field) of the candidate that this
			// ranked vote corresponds to.
			candidateIdx := indexOfRankingInVoteSlice(voterSlice[1:], ranking)
			if candidateIdx == -1 {
				continue
			}
			votesBestToWorst = append(votesBestToWorst, candidates[candidateIdx])
		}

		newVoter := voter{
			timestamp: voterSlice[0],
			votes: votesBestToWorst,
		}
		voters = append(voters, newVoter)
	}

	return voters, nil
}

func validateCSVRecords(csvRecords [][]string) (err error) {
	if len(csvRecords) < 1 {
		return errors.New("CSV does not appear to have a header line.")
	}

	if len(csvRecords[0]) < 1 || csvRecords[0][0] != "Timestamp" {
		return errors.New("CSV header does not appear to have a \"Timestamp\" field.")
	}

	return nil
}


func main() {
	csvFileName := flag.String("csv-file", "", "CSV file from which to determine winner.")
	flag.Parse()

	fileContents, err := ioutil.ReadFile(*csvFileName)
	if err != nil {
		panic("Could not read CSV file: "+err.Error())
	}
	fmt.Println(fileContents)

	// TODO make this something you can pass in by command line flag?
	possibleRankings := []string{"1", "2", "3"}

	r := csv.NewReader(strings.NewReader(string(fileContents)))
	csvRecords, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	voters, err := tabulateVoters(csvRecords, possibleRankings)
	if err != nil {
		panic(err)
	}

	winningCandidates, err := computeInstantRunoffWinner(voters)
	if err != nil {
		panic(err)
	}
	fmt.Println("Winner(s):")
	for _, wc := range(winningCandidates) {
		fmt.Printf("\t%s\n", wc)
	}
}
