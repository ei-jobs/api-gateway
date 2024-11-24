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

func (r *MessageRepository) GetChatsByUserID(userID int) ([]*model.ChatSummary, error) {
	query := `
		WITH latest_messages AS (
		    SELECT
		        CASE
		            WHEN sender_id = ? THEN receiver_id
		            ELSE sender_id
		        END AS other_user_id,
		        MAX(created_at) AS last_sent_time
		    FROM messages
		    WHERE sender_id = ? OR receiver_id = ?
		    GROUP BY other_user_id
		),
		final_result AS (
		    SELECT
		        lm.other_user_id,
		        (
		            SELECT
		                CASE
		                    WHEN role_id = 1 THEN CONCAT(first_name, ' ', last_name)
		                    WHEN role_id = 2 THEN company_name
		                    ELSE 'Unknown'
		                END
		            FROM users
		            WHERE id = lm.other_user_id
		        ) AS receiver_name,
		        m.content,
		        lm.last_sent_time
		    FROM latest_messages lm
		    JOIN messages m
		        ON lm.other_user_id = CASE
		            WHEN m.sender_id = ? THEN m.receiver_id
		            ELSE m.sender_id
		        END
		       AND lm.last_sent_time = m.created_at
		)
		SELECT *
		FROM final_result
		ORDER BY last_sent_time DESC;
	`

	rows, err := r.DB.Query(query, userID, userID, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats []*model.ChatSummary
	for rows.Next() {
		var chat model.ChatSummary
		err := rows.Scan(&chat.ReceiverID, &chat.ReceiverName, &chat.LastMessage, &chat.LastSentTime)
		if err != nil {
			return nil, err
		}
		chats = append(chats, &chat)
	}

	return chats, nil
}

func (r *MessageRepository) GetMessagesByUserAndReceiver(userID, receiverID int) ([]*model.Message, error) {
	query := `
		SELECT id, sender_id, receiver_id, content, resume_id, created_at, updated_at
		FROM messages
		WHERE (sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)
		ORDER BY created_at ASC
	`

	rows, err := r.DB.Query(query, userID, receiverID, receiverID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*model.Message
	for rows.Next() {
		var msg model.Message
		err := rows.Scan(&msg.Id, &msg.SenderID, &msg.ReceiverID, &msg.Content, &msg.ResumeId, &msg.CreatedAt, &msg.UpdatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &msg)
	}

	return messages, nil
}
