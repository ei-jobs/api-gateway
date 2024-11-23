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
	return &MessageHandler{
		MessageService: service,
		Clients:        make(map[*websocket.Conn]bool),
		Broadcast:      make(chan *model.Message),
	}
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
	h.Clients[conn] = true

	for {
		var message model.Message
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Println("Read error:", err)
			delete(h.Clients, conn)
			break
		}
		for _, attachment := range message.Attachments {
			fileData, err := base64.StdEncoding.DecodeString(attachment.Data)
			if err != nil {
				log.Println("Error decoding base64:", err)
				continue
			}

			filePath := fmt.Sprintf("./public/messages/%s-%d-%s", time.Now().Format("20060102_150405"), message.SenderID, attachment.Name)
			if err := ioutil.WriteFile(filePath, fileData, 0644); err != nil {
				log.Println("Error saving file:", err)
				continue
			}

			log.Printf("File saved: %s", filePath)
			attachment.Url = filePath
		}

		messageID, err := h.MessageService.SendMessage(&message)
		if err != nil {
			log.Println("Failed to send message:", err)
			continue
		}

		message.Id = messageID
		h.Broadcast <- &message
	}
}

func (h *MessageHandler) BroadcastMessages() {
	for {
		message := <-h.Broadcast
		for conn := range h.Clients {
			err := conn.WriteJSON(message)
			if err != nil {
				log.Println("Write error:", err)
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
