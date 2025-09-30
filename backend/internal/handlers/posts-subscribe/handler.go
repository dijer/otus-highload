package handler_posts_subscribe

import (
	"fmt"
	"net/http"
	"sync"

	service_posts "github.com/dijer/otus-highload/backend/internal/services/posts"
	"github.com/dijer/otus-highload/backend/internal/utils/httpctx"
	utils_server "github.com/dijer/otus-highload/backend/internal/utils/server"
	"github.com/gorilla/websocket"
	amqp "github.com/rabbitmq/amqp091-go"
)

type PostsSubscribeHandler struct {
	service *service_posts.PostsService
	mu      sync.Mutex
	conns   map[int64][]*websocket.Conn
	rabbit  *amqp.Channel
}

func New(service *service_posts.PostsService, ch *amqp.Channel) *PostsSubscribeHandler {
	return &PostsSubscribeHandler{
		service: service,
		conns:   make(map[int64][]*websocket.Conn),
		rabbit:  ch,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *PostsSubscribeHandler) SubscribeFeed(w http.ResponseWriter, r *http.Request) {
	userID := httpctx.GetUserID(r)
	if userID == 0 {
		utils_server.JsonError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils_server.JsonError(w, http.StatusInternalServerError, "Upgrade error")
		return
	}

	h.mu.Lock()
	h.conns[userID] = append(h.conns[userID], conn)
	h.mu.Unlock()

	queueName := fmt.Sprintf("feed.ws.%d", userID)
	h.rabbit.QueueDeclare(queueName, true, false, false, false, nil)
	h.rabbit.QueueBind(queueName, fmt.Sprintf("feed.deliver.%d", userID), "feeds", false, nil)
	msgs, _ := h.rabbit.Consume(queueName, "", true, false, false, false, nil)

	go func() {
		for d := range msgs {
			h.mu.Lock()
			for _, c := range h.conns[userID] {
				c.WriteMessage(websocket.TextMessage, d.Body)
			}
			h.mu.Unlock()
		}
	}()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}

	h.mu.Lock()
	conns := h.conns[userID]
	for i, c := range conns {
		if c == conn {
			conns = append(conns[:i], conns[i+1:]...)
			break
		}
	}
	if len(conns) == 0 {
		delete(h.conns, userID)
	} else {
		h.conns[userID] = conns
	}
	h.mu.Unlock()

	conn.Close()
}
