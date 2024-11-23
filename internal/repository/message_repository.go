package repository

import (
	"database/sql"
	"time"

	"github.com/aidosgal/ei-jobs-core/internal/model"
)

type MessageRepository struct {
	DB *sql.DB
}

func NewMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{DB: db}
}

func (r *MessageRepository) SaveMessage(message *model.Message) (int64, error) {
	query := `
		INSERT INTO messages (sender_id, receiver_id, content, resume_id, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?);
	`
	result, err := r.DB.Exec(query, message.SenderID, message.ReceiverID, message.Content, message.ResumeId, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}

	messageID, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	for _, attachment := range message.Attachments {
		err = r.SaveAttachment(messageID, attachment)
		if err != nil {
			return 0, err
		}
	}

	return messageID, nil
}

func (r *MessageRepository) SaveAttachment(messageID int64, attachment *model.MessageAttachment) error {
	query := `
		INSERT INTO message_attachments (message_id, type, url)
		VALUES (?, ?, ?);
	`
	_, err := r.DB.Exec(query, messageID, attachment.Type, attachment.Url)
	return err
}
