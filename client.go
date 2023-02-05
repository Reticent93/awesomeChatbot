package main

import "github.com/gorilla/websocket"

//client struct represents the connection our user

type client struct {
	//socket is the web socket for this client
	socket *websocket.Conn
	
	//receive is a channel to receive messages from other users
	receive chan []byte
	
	//room is the room the user is chatting in
	room *room
}

func (c *client) read() {
	//close the socket when the function returns
	defer c.socket.Close()
	
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		//if we receive msg, send it to the room
		c.room.forward <- msg
	}
}

//write is a method that writes to the user
func (c *client) write() {
	//close the socket when the function returns
	defer c.socket.Close()
	
	for msg := range c.receive {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
