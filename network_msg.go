package main

import "encoding/binary"

type MessageType byte

const (
	MsgBlock MessageType = iota
	MsgTx
	MsgPing
	MsgPong
)

type Message struct {
	Type MessageType
	Data []byte
}

func EncodeMsg(m *Message) []byte {
	buf := make([]byte, 1+len(m.Data))
	buf[0] = byte(m.Type)
	copy(buf[1:], m.Data)
	return buf
}

func DecodeMsg(data []byte) *Message {
	return &Message{
		Type: MessageType(data[0]),
		Data: data[1:],
	}
}

func main() {}
