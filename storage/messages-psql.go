package storage

import (
	"database/sql"
	"fmt"
	"time"
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

func (m *MessagesDB) AddNewMessage(sendingDateUnix int, message string, senderId int64, messageId int) error {
	query := fmt.Sprintf("INSERT INTO %s (created_at, sending_date, message, sender_id, message_id) "+
		"VALUES ($1, $2, $3, $4, $5)", m.messagesTableNamespace)

	createdAt := time.Now().UTC()
	sendingDate := time.Unix(int64(sendingDateUnix), 0)

	_, err := m.db.Exec(query, createdAt, sendingDate, message, senderId, messageId)
	if err != nil {
		return err
	}

	return nil
}

func (m *MessagesDB) AddFilteredMessage(sendingDateUnix int, message string, senderId int64, messageId int, filteredBy string) error {
	query := fmt.Sprintf("INSERT INTO %s (created_at, sending_date, message, sender_id, message_id, word_filtered) "+
		"VALUES ($1, $2, $3, $4, $5, %6)", m.filteredMessagesTableNamespace)

	createdAt := time.Now().UTC()
	sendingDate := time.Unix(int64(sendingDateUnix), 0)

	_, err := m.db.Exec(query, createdAt, sendingDate, message, senderId, messageId, filteredBy)
	if err != nil {
		return err
	}

	return nil
}
