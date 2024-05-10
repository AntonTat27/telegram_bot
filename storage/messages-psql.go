package storage

import (
	"database/sql"
)

type MessagesDB struct {
	db                             *sql.DB
	messagesTableNamespace         string
	filteredMessagesTableNamespace string
}

func InitMessagesDB(db *sql.DB, messagesTableNamespace string, filteredMessagesTableNamespace string) MessagesDB {
	messagesDB := MessagesDB{
		db:                             db,
		messagesTableNamespace:         messagesTableNamespace,
		filteredMessagesTableNamespace: filteredMessagesTableNamespace,
	}

	return messagesDB
}
