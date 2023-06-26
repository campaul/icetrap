package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	"icetrap/model"
)

var pool *pgxpool.Pool

const (
	GameCheckSquare   string = "check_square"
	GameUncheckSquare        = "uncheck_square"
	GameNewSession           = "new_session"
)

type GameEvent struct {
	EventType string
	EventData string
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/api/game/") {
		gameHandler(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/api/play/") {
		playHandler(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/api/card/") {
		cardHandler(w, r)
	} else {
		notFoundHandler(w, r)
	}
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/game/")

	if id == "" {
		notFoundHandler(w, r)
		return
	}

	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			errorHandler(w, r)
			return
		}

		g := GameEvent{}
		if err := json.Unmarshal(body, &g); err != nil {
			fmt.Println(err)
			errorHandler(w, r)
			return
		}

		if g.EventType == GameCheckSquare {
			err = model.ToggleSquare(id, g.EventData, true, pool)
		} else if g.EventType == GameUncheckSquare {
			err = model.ToggleSquare(id, g.EventData, false, pool)
		}

		if err != nil {
			fmt.Println(err)
			errorHandler(w, r)
			return
		}
	} else {
		game, err := model.GetGame(id, pool)

		if err != nil {
			fmt.Println(err)
			errorHandler(w, r)
			return
		}

		j, err := json.Marshal(game)

		if err != nil {
			fmt.Println(err)
			errorHandler(w, r)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	}
}

func playHandler(w http.ResponseWriter, r *http.Request) {
}

func cardHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/card/")

	if id == "" {
		notFoundHandler(w, r)
		return
	}

	card, err := model.GetCard(id, pool)

	if err != nil {
		fmt.Println(err)
		errorHandler(w, r)
		return
	}

	j, err := json.Marshal(card)

	if err != nil {
		fmt.Println(err)
		errorHandler(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	fmt.Fprint(w, "Not Found")
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	fmt.Fprint(w, "Error")
}

func main() {
	var err error
	pool, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer pool.Close()

	http.HandleFunc("/", routeHandler)
	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
