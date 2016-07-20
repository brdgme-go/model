package model

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/brdgme/brdgme"
	uuid "github.com/satori/go.uuid"
)

type Game struct {
	Id       string
	Type     string
	Finished bool
	State    []byte
}

// Unmarshal parses the JSON-encoded game state and stores the result in the
// value pointed to by v.
func (g Game) Unmarshal(v interface{}) error {
	return json.Unmarshal(g.State, v)
}

type executor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type rowQuerier interface {
	QueryRow(query string, args ...interface{}) *sql.Row
}

// InsertGame inserts a game into the database.
func InsertGame(db executor, g brdgme.Gamer) (gm Game, err error) {
	gm.Id = uuid.NewV4().String()
	gm.Type = g.Identifier()
	gm.Finished, _ = g.IsFinished()
	gm.State, err = json.Marshal(g)
	if err != nil {
		err = fmt.Errorf("error converting game to JSON for inserting, %s", err)
		return
	}
	_, err = db.Exec(
		`
INSERT INTO games (
  id,
  game_type,
  finished,
  game_state
) VALUES (
  $1,
  $2,
  $3,
  $4
)`,
		gm.Id,
		gm.Type,
		gm.Finished,
		gm.State,
	)
	if err != nil {
		err = fmt.Errorf("error inserting new game into database, %s", err)
		return
	}
	return
}

// LoadGame loads a game from the database specified by the id UUID.
func LoadGame(db rowQuerier, id string) (g Game, ok bool, err error) {
	row := db.QueryRow(
		`
SELECT
  id,
  game_type,
  finished,
  game_state
FROM games
WHERE id=$1`,
		id,
	)
	ok = true
	if err := row.Scan(
		&g.Id,
		&g.Type,
		&g.Finished,
		&g.State,
	); err != nil {
		ok = false
		err = fmt.Errorf("error fetching row, %s", err)
	}
	return
}
