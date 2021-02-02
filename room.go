package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const (
	POSTGRES_CONNECT = "postgres://superuser:superuser@127.0.0.1:5432/fuckthesuits?sslmode=disable"
)

var (
	project = &Project{}
)

func reportPostgresProblem(ev pq.ListenerEventType, err error) {
	if err != nil {
		fmt.Println("Postgres NOTIFY error: " + err.Error())
	}
}

func initPostgresListener(tableName string) *pq.Listener {
	listener := pq.NewListener(POSTGRES_CONNECT, 10*time.Second, time.Minute,
		reportPostgresProblem)
	err := listener.Listen(tableName)
	if err != nil {
		panic(err)
	}

	log.Printf("LISTENing for all changes to the `%s` table to push to WebSocket clients",
		tableName)

	return listener
}

func waitForNotification(tableName string) {
	listener := initPostgresListener(tableName)

	for {
		select {
		case n := <-listener.Notify:
			if n == nil {
				fmt.Printf("Got nil notification on table `%v`! Postgres may have crashed, or we just reconnected\n", tableName)
				time.Sleep(1 * time.Second)
				continue
			}
			fmt.Println("Received data from channel [", n.Channel, "] :", n.Extra)

			project.BroadcastMessages(nil, Message(n.Extra))

		case <-time.After(90 * time.Second):
			go func() {
				listener.Ping()
			}()
		}
	}
}

func NewProject() *Project {
	tables := []string{
		// TODO(elimisteve): Add fuckthesuits DB table names here
	}

	for _, tableName := range tables {
		tableName := tableName
		go waitForNotification(tableName)
	}

	return &Project{}
}

type Project struct {
	Clients    []*Client
	clientLock sync.RWMutex
}

func (p *Project) AddClient(c *Client) {
	p.clientLock.Lock()
	defer p.clientLock.Unlock()

	p.Clients = append(p.Clients, c)
}

func (p *Project) RemoveClient(c *Client) {
	p.clientLock.Lock()
	defer p.clientLock.Unlock()

	for i, client := range p.Clients {
		if client == c {
			p.Clients = append(p.Clients[:i], p.Clients[i+1:]...)
			break
		}
	}
}

// If it is a message from the project, make the sender nil.
func (p *Project) BroadcastMessages(sender *Client, msgs ...Message) {
	p.clientLock.RLock()
	defer p.clientLock.RUnlock()

	for _, client := range p.Clients {
		go func(client *Client) {
			err := client.SendMessages(msgs...)
			if err != nil {
				log.Debugf("Error sending message. Err: %s", err)
			}
		}(client)
	}
}

type Client struct {
	wsConn    *websocket.Conn
	writeLock sync.Mutex

	project *Project
}

func (c *Client) SendMessages(msgs ...Message) error {
	c.writeLock.Lock()
	defer c.writeLock.Unlock()

	outgoing := OutgoingPayload{Changes: msgs}

	body, err := json.Marshal(outgoing)
	if err != nil {
		return err
	}

	err = c.wsConn.WriteMessage(websocket.TextMessage, body)
	if err != nil {
		log.Debugf("Error sending message to client. Removing client from project. Err: %s", err)
		c.project.RemoveClient(c)
		return err
	}

	return nil
}

func (c *Client) SendError(errStr string, secretErr error) error {
	c.writeLock.Lock()
	defer c.writeLock.Unlock()

	return WSWriteError(c.wsConn, errStr, secretErr)
}
