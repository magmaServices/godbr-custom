package model

import (
	"encoding/json"

	"github.com/gocraft/dbr"
)

const (
	tableHeroes = "game_heroes" // the name of the table 
)

// Hero defines a single record in heroes table
type Hero struct {
	ID        int    `db:"heroID"`
	HeroName  string `db:"heroName"`
	PlayerID  int    `db:"user_id"`
	HeroStats string `db:"hero_stats"`
}

// func (h *Hero) Stats() (*HeroStats, error) {
// 	var s HeroStats
// 	err := json.Unmarshal([]byte(h.HeroStats), &s)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &s, nil
// }

// FindHeroesByPlayerID returns a list of heroes associated with specified player
func (q *Queries) FindHeroesByPlayerID(sess *dbr.Session, playerID int) (hs []Hero, err error) {
	
_, err = sess.Select("heroID", "heroName", "user_id").From(tableHeroes).Where("user_id = ?", playerID).Load(&hs)


	return hs, err
}

// FindHeroByName returns a hero with specified name
func (q *Queries) FindHeroByName(sess *dbr.Session, heroName string) (h Hero, err error) {

	_, err = sess.Select("heroID", "heroName", "user_id").From(tableHeroes).Where("heroName = ?", heroName).Load(&h)

	return h, err
}

// FindHeroStats returns stats of hero of specified ID
func (q *Queries) FindHeroStats(sess *dbr.Session, heroID int) (pr HeroStats, err error) {
	var payload []byte

	_, err = sess.Select("hero_stats").From(tableHeroes).Where("heroID = ?", heroID).Load(&payload)
	if err != nil {
		return pr, err
	}

	err = json.Unmarshal(payload, &pr)
	return pr, err
}

// UpdateHeroStats changes stats of hero of specified ID
func (q *Queries) UpdateHeroStats(tx *dbr.Tx, heroID int, pr *HeroStats) error {
	by, err := json.Marshal(pr)
	if err != nil {
		return err
	}
	_, err = tx.Update(tableHeroes).Set("hero_stats", string(by)).Where("heroID = ?", heroID).Exec()

	return err
}

func (q *Queries) SubStats(tx *dbr.Tx, heroID int, pr *HeroStats) error {
	by, err := json.Marshal(pr)
	if err != nil {
		return err
	}
	_, err = tx.Update(tableHeroes).Set("hero_stats", string(by)).Where("heroID = ?", heroID).Exec()
	return err
}