package model

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Square struct {
	Id       int
	Title    string
	Selected bool
	Needed   bool
}

type Game struct {
	Title   string
	Squares []Square
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GetGame(id string, pool *pgxpool.Pool) (Game, error) {
	// Get Game Info
	var title string
	var session int
	err := pool.QueryRow(context.Background(), "select title, current_session from games where id=$1", id).Scan(&title, &session)

	// TODO: err is not nil when there are 0 rows returned, this results in a 500 that should be a 404
	if err != nil {
		return Game{}, err
	}

	// Get Game Squares
	rows, err := pool.Query(context.Background(), "select squares.id, title, coalesce(selected, false) as selected from squares left join selections on squares.id=square_id and session_id = $1 where game_id = $2 order by squares.id", session, id)

	if err != nil {
		return Game{}, err
	}

	defer rows.Close()

	var squares []Square

	for rows.Next() {
		var id int
		var title string
		var selected bool
		err := rows.Scan(&id, &title, &selected)
		if err != nil {
			return Game{}, err
		}
		squares = append(squares, Square{
			Id:       id,
			Title:    title,
			Selected: selected,
			Needed:   false,
		})
	}

	if len(squares) == 0 {
		return Game{}, err
	}

	// Find squares people need
	rows, err = pool.Query(context.Background(), "select distinct square_id from card_squares join cards on card_id = cards.id where session_id = $1", session)

	if err != nil {
		return Game{}, err
	}

	defer rows.Close()

	var needed []int
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return Game{}, err
		}
		needed = append(needed, id)
	}

	needed_squares := squares[:0]
	for _, square := range squares {
		if contains(needed, square.Id) {
			square.Needed = true
		}

		needed_squares = append(needed_squares, square)
	}

	// Generate game
	return Game{
		Title:   title,
		Squares: needed_squares,
	}, nil
}
