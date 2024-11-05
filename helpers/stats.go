package helpers

import (
	"encoding/json"
	"net/http"
)

type MatchStats struct {
	Deck1        string  `json:"deck1"`
	Deck2        string  `json:"deck2"`
	Deck1Wins    int     `json:"deck1_wins"`
	Deck2Wins    int     `json:"deck2_wins"`
	TotalMatches int     `json:"total_matches"`
	Deck1WinRate float64 `json:"deck1_win_rate"`
	Deck2WinRate float64 `json:"deck2_win_rate"`
}

func GetStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := GetDB()
	query := `SELECT 
            deck1, 
            deck2, 
            SUM(CASE WHEN winner = player1 THEN 1 ELSE 0 END) AS deck1_wins,
            SUM(CASE WHEN winner = player2 THEN 1 ELSE 0 END) AS deck2_wins,
            COUNT(*) AS total_matches
        FROM 
            matches
        GROUP BY 
            deck1, deck2;`

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var stats []MatchStats

	for rows.Next() {
		var stat MatchStats
		if err := rows.Scan(&stat.Deck1, &stat.Deck2, &stat.Deck1Wins, &stat.Deck2Wins, &stat.TotalMatches); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if stat.TotalMatches > 0 {
			stat.Deck1WinRate = float64(stat.Deck1Wins) / float64(stat.TotalMatches) * 100
			stat.Deck2WinRate = float64(stat.Deck2Wins) / float64(stat.TotalMatches) * 100
		}

		stats = append(stats, stat)
	}
	json.NewEncoder(w).Encode(stats)
}
