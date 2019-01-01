package model

import (
	"github.com/gocraft/dbr"
	"github.com/guregu/null"
)

const (
	tablePlayers = "players" // used in nugetaccount

)

// Player
// table = player -- player or players?
type Player struct {
	ID             int         `db:"user_id"`
	Username       string      `db:"username"`
	Password       string      `db:"password"`
	GameToken      string      `db:"game_token"` // TODO: something like one-time password
	PreferServer   null.String `db:"prefer_server"`
	SelectedHeroID null.Int    `db:"selected_hero_id"`
}

func InsertPlayer(tx *dbr.Tx, player *Player) error {
	r, err := tx.
		InsertInto(tablePlayers).
		Columns(
			//"user_id", // TODO: skip user_id
			"username",
			"password",
			"game_token",
		).
		Record(player).
		Exec()
	if err != nil {
		return err
	}
	i, err := r.LastInsertId()
	if err != nil {
		return err
	}
	player.ID = int(i)
	return nil
}

// FindPlayerByToken returns a player associated with given game_token
func (q *Queries) FindPlayerByToken(sess *dbr.Session, token string) (player Player, err error) {


	_, err = sess.Select("user_id", "username", "password", "game_token", "prefer_server", "selected_hero_id").From(tablePlayers).Where("game_token = ?", token).Load(&player)

	// err = sess.
	// 	Select(
	// 		"user_id", //todo get only one value / but do some encrypt with it //uuid ?
	// 		"username",
	// 		"password",
	// 		"game_token",
	// 		"prefer_server",
	// 		"selected_hero_id",
	// 	).
	// 	From(tablePlayers).
	// 	Where("game_token = ?", token).
	// 	LoadStruct(&player)
	return player, err
}

// FindPlayerByID returns a player associated with given playerID
func (q *Queries) FindPlayerByID(sess *dbr.Session, playerID int) (player Player, err error) {
	_, err = sess.
		Select(
			"user_id",
			"username",
			"password",
			"game_token",
			"prefer_server",
			"selected_hero_id",
		).
		From(tablePlayers).
		Where("user_id = ?", playerID).//IMPROVE THIS
		Load(&player)
	return player, err
}
