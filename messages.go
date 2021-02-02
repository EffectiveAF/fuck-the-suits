package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type Message string

type ToServer struct {
	ProjectID int `json:"project_id"`
}

type IncomingPayload struct {
	ToServer ToServer `json:"to_server"`
}

type OutgoingPayload struct {
	Changes []Message `json:"changes"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:   3 * 1024,
	WriteBufferSize:  3 * 1024,
	HandshakeTimeout: 45 * time.Second,
}

func WSAllHandler(w http.ResponseWriter, req *http.Request) {
	wsConn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		errStr := "Unable to upgrade to websocket conn"
		realErr := errors.New(errStr + ": " + err.Error())
		WriteErrorStatus(w, errStr, realErr, http.StatusBadRequest)
		return
	}

	client := &Client{
		wsConn: wsConn,
	}
	project.AddClient(client)

	go messageReader(client)
}

func messageReader(client *Client) {
	for {
		messageType, p, err := client.wsConn.ReadMessage()
		if err != nil {
			// TODO: Consider adding more checks
			log.Debugf("Error reading ws message: %s", err)
			project.RemoveClient(client)
			return
		}

		// Respond to message depending on message type
		switch messageType {
		case websocket.TextMessage:
			var payload IncomingPayload
			err := json.Unmarshal(p, &payload)
			if err != nil {
				log.Debugf("Error unmarshalling message `%s` -- %s", p, err)
				continue
			}

		case websocket.BinaryMessage:
			log.Debug("Binary messages are unsupported")

		case websocket.CloseMessage:
			log.Debug("Got close message")
			project.RemoveClient(client)
			return

		default:
			log.Debugf("Unsupport messageType: %d", messageType)

		}

	}
}
