package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"strings"
)

// TODO read CSV data from a file.
var csvData string = `"Timestamp","Votes [The Blues Brothers (1980)]","Votes [Flight of the Navigator (1986)]","Votes [Ralph Breaks the Internet (2018)]","Votes [Premium Rush (2012)]","Votes [Galaxy Quest (1999)]","Votes [Mystery Men (1999)]"
"2020/04/29 7:23:51 PM AST","1","2","3","","",""
"2020/04/29 7:24:00 PM AST","","3","","","2","1"`


type voter struct {
	timestamp string
	// Candidate names from most preferred (highest-ranked) to least preferred.
	votes []string
}


// indexStringSlice returns the index of the first instance of str in strSlice. -1 if str is not found.
func indexStringSlice(strSlice []string, str string) (pos int) {
	for i, s := range(strSlice) {
		if s == str {
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
	fmt.Println(candidates)

	for _, voterSlice := range(csvRecords[1:]) {
		fmt.Println(voterSlice)

		// for each possible ranking (best to worst), get its index and use that to grab the candidate
		// append the candidate to a slice.
		var votesBestToWorst []string
		for _, ranking := range(possibleRankings) {
			// Get the offset (ignoring the timestamp field) of the candidate that this
			// ranked vote corresponds to.
			candidateIdx := indexStringSlice(voterSlice[1:], ranking)
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


// TODO this needs to first
// TODO add a verbose flag to print everything?
func computeInstatRunoffWinner(csvRecords [][]string) (winningCandidate string, err error) {
	panic("FIXME: computeInstatRunoffWinner; not implemented!")
	return
}



func main() {
	// TODO make this something you can pass in by command line flag?
	possibleRankings := []string{"1", "2", "3"}

	r := csv.NewReader(strings.NewReader(csvData))
	csvRecords, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	voters, err := tabulateVoters(csvRecords, possibleRankings)
	if err != nil {
		panic(err)
	}

	// TODO call computeInstatRunoffWinner

	fmt.Println(voters)
}
