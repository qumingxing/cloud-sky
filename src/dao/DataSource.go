package dao

import (
	"database/sql"
)

type DataSource interface {
	GetConnection() *sql.DB
	DesConnection(db *sql.DB)
	Connect() (*sql.DB, bool)
}
