package helpers

import (
	"encoding/json"
	"net/http"
)

type Match struct {
	ID      int    `json:"id"`
	Player1 string `json:"player1"`
	Player2 string `json:"player2"`
	Deck1   string `json:"deck1"`
	Deck2   string `json:"deck2"`
	Winner  string `json:"winner"`
	Turns   string `json:"turns"`
	Date    string `json:"date"`
}

func AddMatch(w http.ResponseWriter, r *http.Request) {
	var match Match
	if err := json.NewDecoder(r.Body).Decode(&match); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := GetDB()
	_, err := db.Exec("INSERT INTO matches (player1, player2, deck1, deck2, winner, turns, date) VALUES (?, ?, ?, ?, ?, ?, ?)",
		match.Player1, match.Player2, match.Deck1, match.Deck2, match.Winner, match.Turns, match.Date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetMatches(w http.ResponseWriter, r *http.Request) {
	db := GetDB()
	rows, err := db.Query("SELECT * FROM matches")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var matches []Match
	for rows.Next() {
		var match Match
		if err := rows.Scan(&match.ID, &match.Player1, &match.Player2, &match.Deck1, &match.Deck2, &match.Winner, &match.Turns, &match.Date); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		matches = append(matches, match)
	}

	json.NewEncoder(w).Encode(matches)
}
