package dao

import (
	"database/sql"
)

type Transaction interface {
	Begin() *TransactionConfig
	Commit(config *TransactionConfig) error
	Rollback(config *TransactionConfig) error
	AddTx(transactionConfig *TransactionConfig, query string, params ...interface{}) (affectCount int, err error)
	UpdateTx(transactionConfig *TransactionConfig, sql string, params ...interface{}) (affectCount int, err error)
	DeleteTx(transactionConfig *TransactionConfig, sql string, params ...interface{}) (affectCount int, err error)
}
type TransactionConfig struct {
	*sql.Tx
	conn *sql.DB
}
