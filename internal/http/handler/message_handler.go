package handler

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/aidosgal/ei-jobs-core/internal/model"
	"github.com/aidosgal/ei-jobs-core/internal/service"
	"github.com/aidosgal/ei-jobs-core/internal/utils"
	"github.com/gorilla/websocket"
)

type MessageHandler struct {
	MessageService *service.MessageService
	Clients        map[*websocket.Conn]bool
	Broadcast      chan *model.Message
}

func NewMessageHandler(service *service.MessageService) *MessageHandler {
	handler := &MessageHandler{
		MessageService: service,
		Clients:        make(map[*websocket.Conn]bool),
		Broadcast:      make(chan *model.Message),
	}
	go handler.BroadcastMessages()
	return handler
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *MessageHandler) HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade WebSocket", http.StatusInternalServerError)
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	log.Printf("New client connected. Total clients: %d", len(h.Clients))
	h.Clients[conn] = true

	for {
		var message model.Message
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Printf("Read error: %v. Removing client. Remaining clients: %d", err, len(h.Clients)-1)
			delete(h.Clients, conn)
			break
		}

		log.Printf("Received message from sender ID: %d to receiver ID: %d", message.SenderID, message.ReceiverID)

		for i, attachment := range message.Attachments {
			fileData, err := base64.StdEncoding.DecodeString(attachment.Data)
			if err != nil {
				log.Printf("Error decoding base64 for attachment %d: %v", i, err)
				continue
			}

			filePath := fmt.Sprintf("./public/messages/%s-%d-%s",
				time.Now().Format("20060102_150405"),
				message.SenderID,
				attachment.Name)

			if err := ioutil.WriteFile(filePath, fileData, 0644); err != nil {
				log.Printf("Error saving file %s: %v", filePath, err)
				continue
			}

			log.Printf("File saved successfully: %s", filePath)
			attachment.Url = filePath
		}

		messageID, err := h.MessageService.SendMessage(&message)
		if err != nil {
			log.Printf("Failed to send message to service: %v", err)
			continue
		}

		message.Id = messageID
		log.Printf("Broadcasting message ID: %d to %d clients", messageID, len(h.Clients))

		h.Broadcast <- &message
	}
}

func (h *MessageHandler) BroadcastMessages() {
	for {
		message := <-h.Broadcast
		log.Printf("Broadcasting message %d to %d clients", message.Id, len(h.Clients))

		for conn := range h.Clients {
			err := conn.WriteJSON(message)
			if err != nil {
				log.Printf("Write error for client: %v", err)
				conn.Close()
				delete(h.Clients, conn)
			}
		}
	}
}

func (h *MessageHandler) GetChatsByUserID(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	chats, err := h.MessageService.GetChatsByUserID(userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, chats)
}

func (h *MessageHandler) GetMessagesByUserAndReceiver(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	receiverIDStr := r.URL.Query().Get("receiver_id")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	receiverID, err := strconv.Atoi(receiverIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	messages, err := h.MessageService.GetMessagesByUserAndReceiver(userID, receiverID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, messages)
}
