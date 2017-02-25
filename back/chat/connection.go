package chat

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	nrand "math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// ========== addition methods

// random {{{
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func rndStr(n int) string {
	rnd_str := make([]rune, n)
	for i := range rnd_str {
		rnd_str[i] = letterRunes[nrand.Intn(len(letterRunes))]
	}
	return string(rnd_str)
}

// RandToken generates a random @length token.
func RandToken(length int) string {
	thisByte := make([]byte, length)
	rand.Read(thisByte)
	return base64.StdEncoding.EncodeToString(thisByte)
} // }}}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is an middleman between the websocket connection and the hub
type Client struct {
	// Nick name connected client
	nick string
	// hub
	hub *Hub
	// The websocket connection
	conn *websocket.Conn
	// Buffered channel of outbound messages
	send chan []byte
}

// readPump pumps messages from the websocket connection to the hub
func (c *Client) readPump() { // {{{
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

		// fmt.Printf("\n(readPump)nick: %s\nmsg: %s\n", c.nick, message)
		strmsg := []byte(c.nick + ": " + string(message))
		c.hub.broadcast <- strmsg
	}
} // }}}

// writePump pumps messages from the hub to the websocket connection
func (c *Client) writePump() { // {{{
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			// fmt.Printf("\nnick: %v\n", c.nick)
			// strmsg := []byte(c.nick + ": " + string(message))
			w.Write(message)
			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
	/* {{{
		for {
			select {
			case message, ok := <-c.send:
				if !ok {
					// The hub closed the channel.
					c.write(websocket.CloseMessage, []byte{})
					return
				}

				c.ws.SetWriteDeadline(time.Now().Add(writeWait))
				w, err := c.ws.NextWriter(websocket.TextMessage)
				if err != nil {
					return
				}
				w.Write(message)
				// Add queued chat messages to the current websocket message.
				n := len(c.send)
				for i := 0; i < n; i++ {
					w.Write(newline)
					w.Write(<-c.send)
				}
				if err := w.Close(); err != nil {
					return
				}
			case <-ticker.C:
				if err := c.write(websocket.PingMessage, []byte{}); err != nil {
					return
				}
			}
		}
	}}}*/
} // }}}

// write writes a message with the given message type and payload.
func (c *Client) write(mt int, payload []byte) error { // {{{
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	return c.conn.WriteMessage(mt, payload)
} // }}}

// serveWs handles websocket requests from the peer
func serveWs(hub *Hub, c *gin.Context) { // {{{
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	rnd_nick := rndStr(6)
	fmt.Printf("\nnew user connected as: %v\n", rnd_nick)
	client := &Client{hub: hub, nick: rnd_nick, send: make(chan []byte, 256), conn: conn}
	client.hub.register <- client
	go client.writePump()
	client.readPump()
} // }}}
