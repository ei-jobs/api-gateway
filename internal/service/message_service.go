package service

import (
	"errors"

	"github.com/aidosgal/ei-jobs-core/internal/model"
	"github.com/aidosgal/ei-jobs-core/internal/repository"
)

type MessageService struct {
	Repo *repository.MessageRepository
}

func NewMessageService(repo *repository.MessageRepository) *MessageService {
	return &MessageService{Repo: repo}
}

func (s *MessageService) SendMessage(message *model.Message) (int64, error) {
	if message.SenderID == 0 || message.ReceiverID == 0 {
		return 0, errors.New("invalid sender or receiver")
	}
	if message.Content == nil && len(message.Attachments) == 0 {
		return 0, errors.New("message must have content or attachments")
	}

	messageID, err := s.Repo.SaveMessage(message)
	if err != nil {
		return 0, err
	}

	return messageID, nil
}

func (s *MessageService) GetChatsByUserID(userID int) ([]*model.ChatSummary, error) {
	return s.Repo.GetChatsByUserID(userID)
}

func (s *MessageService) GetMessagesByUserAndReceiver(userID, receiverID int) ([]*model.Message, error) {
	return s.Repo.GetMessagesByUserAndReceiver(userID, receiverID)
}
