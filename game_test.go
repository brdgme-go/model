// +build integration

package model

import (
	"io"
	"testing"

	"github.com/brdgme/brdgme"

	"github.com/stretchr/testify/assert"
)

type fakeGame struct {
	Players int
	Cards   []int
	Hands   map[int][]int
}

func (fg *fakeGame) Name() string {
	return "Fake Game"
}

func (fg *fakeGame) Identifier() string {
	return "fake_game"
}

func (fg *fakeGame) Start(players int) ([]brdgme.Log, error) {
	return nil, nil
}

func (fg *fakeGame) AvailableCommands(player int) []brdgme.CommandDescription {
	return nil
}

func (fg *fakeGame) Command(
	player int,
	input io.Reader,
	playerNames []string,
) (logs []brdgme.Log, remaining io.Reader, err error) {
	return nil, nil, nil
}

func (fg *fakeGame) IsFinished() (finished bool, winners []int) {
	return true, nil
}

func (fg *fakeGame) WhoseTurn() []int {
	return nil
}

var fakeGameInstance = &fakeGame{
	Players: 5,
	Cards:   []int{1, 2, 3, 4, 5},
	Hands: map[int][]int{
		0: {1, 2, 3},
		1: {1, 2, 4},
		2: {},
		3: nil,
		4: {7},
	},
}

func TestInsertGame(t *testing.T) {
	trans, err := db.Begin()
	assert.NoError(t, err)
	defer trans.Rollback()
	game, err := InsertGame(trans, fakeGameInstance)
	assert.NoError(t, err)
	assert.NotEmpty(t, game.Id)
}

func TestLoadGame(t *testing.T) {
	trans, err := db.Begin()
	assert.NoError(t, err)
	defer trans.Rollback()
	game, err := InsertGame(trans, fakeGameInstance)
	assert.NoError(t, err)
	game2, ok, err := LoadGame(trans, game.Id)
	assert.NoError(t, err)
	assert.True(t, ok)
	fakeGameInstance2 := &fakeGame{}
	assert.NoError(t, game2.Unmarshal(fakeGameInstance2))
	assert.Equal(t, fakeGameInstance, fakeGameInstance2)
}
