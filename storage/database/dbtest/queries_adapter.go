package dbtest

import (
	"github.com/gocraft/dbr"
	"gitlab.com/oiacow/nextfesl/model"
)

type QueriesAdapter struct {
	Server                     model.Server
	ErrFindServerByID          error
	ErrFindServerBySoldierName error
	ErrFindServerByCredentials error

	Player               model.Player
	ErrFindPlayerByToken error
	ErrFindPlayerByID    error

	Heroes                  []model.Hero //does this apply to table name ?
	ErrFindHeroesByPlayerID error
	Hero                    model.Hero
	ErrFindHeroByName       error
	HeroStats               model.HeroStats
	ErrFindHeroStats        error
	ErrUpdateHeroStats      error
}

func (a *QueriesAdapter) FindServerByID(sess *dbr.Session, serverID int) (model.Server, error) {
	return a.Server, a.ErrFindServerByID
}
func (a *QueriesAdapter) FindServerBySoldierName(sess *dbr.Session, soldierName string) (model.Server, error) {
	return a.Server, a.ErrFindServerBySoldierName
}
func (a *QueriesAdapter) FindServerByCredentials(sess *dbr.Session, accountName string) (model.Server, error) {
	return a.Server, a.ErrFindServerByCredentials
}

func (a *QueriesAdapter) FindPlayerByToken(sess *dbr.Session, token string) (model.Player, error) {
	return a.Player, a.ErrFindPlayerByToken
}
func (a *QueriesAdapter) FindPlayerByID(sess *dbr.Session, playerID int) (model.Player, error) {
	return a.Player, a.ErrFindPlayerByID
}

func (a *QueriesAdapter) FindHeroesByPlayerID(sess *dbr.Session, playerID int) ([]model.Hero, error) {
	return a.Heroes, a.ErrFindHeroesByPlayerID
}
func (a *QueriesAdapter) FindHeroByName(sess *dbr.Session, heroName string) (model.Hero, error) {
	return a.Hero, a.ErrFindHeroByName
}
func (a *QueriesAdapter) FindHeroStats(sess *dbr.Session, heroID int) (model.HeroStats, error) {
	return a.HeroStats, a.ErrFindHeroStats
}
func (a *QueriesAdapter) UpdateHeroStats(tx *dbr.Tx, heroID int, pr *model.HeroStats) error {
	return a.ErrUpdateHeroStats
}
