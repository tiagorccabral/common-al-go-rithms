package arrays

const HOME_TEAM_WON = 1

// Category: arrays
// Difficulty: Easy

// Given an array of matches, and a array of result of those matches, return the winning team.
// A winner gets 3 points, and there will always be at at least and at most one winner.

func TournamentWinner(competitions [][]string, results []int) string {
	teamsScoreMap := make(map[string]int)

	for x, row := range competitions {
		if results[x] == HOME_TEAM_WON {
			current := teamsScoreMap[row[0]]
			teamsScoreMap[row[0]] = current + 3
		} else {
			current := teamsScoreMap[row[1]]
			teamsScoreMap[row[1]] = current + 3
		}
	}

	winTeam := ""
	highestScore := 0
	for team, score := range teamsScoreMap {
		if score > highestScore {
			winTeam = team
			highestScore = score
		}
	}

	return winTeam
}
