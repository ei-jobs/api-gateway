package model

import "time"

type Message struct {
	Id          int64                `json:"id"`
	SenderID    int                  `json:"sender_id"`
	ReceiverID  int                  `json:"receiver_id"`
	Content     *string              `json:"content"`
	ResumeId    *string              `json:"resume_id"`
	Attachments []*MessageAttachment `json:"attachments"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}

type MessageAttachment struct {
	MessageId int    `json:"message_id"`
	Type      string `json:"type"`
	Url       string `json:"url"`
	Data      string `json:"data"`
	Name      string `json:"name"`
}

type ChatSummary struct {
	ReceiverID   int       `json:"receiver_id"`
	ReceiverName string    `json:"receiver_name"`
	LastMessage  *string   `json:"last_message"`
	LastSentTime time.Time `json:"last_sent_time"`
}
