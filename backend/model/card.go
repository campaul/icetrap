package model

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Card struct {
	Squares []Square
}

func GetCard(id string, pool *pgxpool.Pool) (Card, error) {
	rows, err := pool.Query(context.Background(), "select squares.title, card_squares.square_id, coalesce(selected, false) as selected from card_squares join cards on cards.id = card_squares.card_id left join selections on card_squares.square_id = selections.square_id and cards.session_id = selections.session_id join squares on squares.id = card_squares.square_id where card_id = $1", id)
	if err != nil {
		return Card{}, err
	}

	defer rows.Close()

	var squares []Square

	for rows.Next() {
		var title string
		var id int
		var selected bool
		err := rows.Scan(&title, &id, &selected)
		if err != nil {
			return Card{}, err
		}
		squares = append(squares, Square{
			Id:       id,
			Title:    title,
			Selected: selected,
		})
	}

	if len(squares) == 0 {
		return Card{}, err
	}

	return Card{
		Squares: squares,
	}, nil
}
