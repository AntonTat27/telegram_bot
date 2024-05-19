package storage

import (
	"database/sql"
	"fmt"
	"time"
)

// MessagesDB represents a type for managing messages in a database.
//
// It has fields for the database connection, table namespaces for messagesTable and
// filteredMessagesTable.
type MessagesDB struct {
	db                             *sql.DB
	messagesTableNamespace         string
	filteredMessagesTableNamespace string
}

// InitMessagesDB initializes a MessagesDB struct with the provided sql.DB, messagesTableNamespace, and filteredMessagesTableNamespace values.
func InitMessagesDB(db *sql.DB, messagesTableNamespace string, filteredMessagesTableNamespace string) MessagesDB {
	messagesDB := MessagesDB{
		db:                             db,
		messagesTableNamespace:         messagesTableNamespace,
		filteredMessagesTableNamespace: filteredMessagesTableNamespace,
	}

	return messagesDB
}

// AddNewMessage inserts a new message into the messages table in the database.
// sendingDateUnix - time, in unix format, when the message was sent
// message - the text of the message
// senderId - the Id of account, that sent the message
// messageId - the Id of a message
func (m *MessagesDB) AddNewMessage(sendingDateUnix int, message string, senderId int64, messageId int) error {
	// Preparing query
	query := fmt.Sprintf("INSERT INTO %s (created_at, sending_date, message, sender_id, message_id) "+
		"VALUES ($1, $2, $3, $4, $5)", m.messagesTableNamespace)

	// Getting current date and converting sendingDateUnix to a datetime format
	createdAt := time.Now().UTC()
	sendingDate := time.Unix(int64(sendingDateUnix), 0)

	// Executing query
	_, err := m.db.Exec(query, createdAt, sendingDate, message, senderId, messageId)
	if err != nil {
		return err
	}

	return nil
}

// AddFilteredMessage inserts a new message into the messages table in the database.
// sendingDateUnix - time, in unix format, when the message was sent
// message - the text of the message
// senderId - the Id of account, that sent the message
// messageId - the Id of a message
// filteredBy - the word, by which the message was filtered
func (m *MessagesDB) AddFilteredMessage(sendingDateUnix int, message string, senderId int64, messageId int, filteredBy string) error {
	// Preparing query
	query := fmt.Sprintf("INSERT INTO %s (created_at, sending_date, message, sender_id, message_id, word_filtered) "+
		"VALUES ($1, $2, $3, $4, $5, $6)", m.filteredMessagesTableNamespace)

	// Getting current date and converting sendingDateUnix to a datetime format
	createdAt := time.Now().UTC()
	sendingDate := time.Unix(int64(sendingDateUnix), 0)

	// Executing query
	_, err := m.db.Exec(query, createdAt, sendingDate, message, senderId, messageId, filteredBy)
	if err != nil {
		return err
	}

	return nil
}
