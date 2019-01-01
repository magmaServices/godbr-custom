package database

import (
	"github.com/gocraft/dbr"
	"gitlab.com/oiacow/nextfesl/model"
)

type Adapter interface {
	TxAdapter
	model.QueriesAdapter
}

type TxAdapter interface {
	NewSession() *dbr.Session

	Begin(sess *dbr.Session) (*dbr.Tx, error)
	Commit(tx *dbr.Tx) error
	Rollback(tx *dbr.Tx) error
}
