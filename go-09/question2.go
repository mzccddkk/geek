// 实现一个从 socket connection 中解码出 goim 协议的解码器。

// Package Length   4 bytes header + body length
// Header Length    2 bytes protocol header length
// Protocol Version 2 bytes protocol version
// Operation 		4 bytes operation for request
// Sequence id 		4 bytes sequence number chosen by client
// Body		PackLen + HeaderLen		binary body bytes

package main

import (
	"encoding/binary"
	"fmt"
)

const (
	packSize   = 4
	headerSize = 2
	verSize    = 2
	opSize     = 4
	seqSize    = 4
	headerLen  = packSize + headerSize + verSize + opSize + seqSize

	packOffsetR   = packSize
	headerOffsetR = packSize + headerSize
	verOffsetR    = headerOffsetR + verSize
	opOffsetR     = verOffsetR + opSize
	seqOffsetR    = opOffsetR + seqSize

	version   = 1
	operation = 2
	sequence  = 3
)

func decoder(data []byte) {
	if len(data) <= headerLen {
		fmt.Printf("data len < %v.", headerLen)
		return
	}

	packLen := binary.BigEndian.Uint32(data[:packOffsetR])
	headerLen := binary.BigEndian.Uint16(data[packOffsetR:headerOffsetR])
	version := binary.BigEndian.Uint16(data[headerOffsetR:verOffsetR])
	operation := binary.BigEndian.Uint32(data[verOffsetR:opOffsetR])
	sequence := binary.BigEndian.Uint32(data[opOffsetR:seqOffsetR])
	body := string(data[headerLen:])

	fmt.Printf("packetLen:%v\nheaderLen:%v\nversion:%v\noperation:%v\nsequence:%v\nbody:%v\n",
		packLen, headerLen, version, operation, sequence, body)
}

func encoder(body string) (buf []byte) {
	packLen := len(body) + headerLen
	buf = make([]byte, packLen)

	binary.BigEndian.PutUint32(buf[:packOffsetR], uint32(packLen))
	binary.BigEndian.PutUint16(buf[packOffsetR:headerOffsetR], uint16(headerLen))
	binary.BigEndian.PutUint16(buf[headerOffsetR:verOffsetR], uint16(version))
	binary.BigEndian.PutUint32(buf[verOffsetR:opOffsetR], uint32(operation))
	binary.BigEndian.PutUint32(buf[opOffsetR:seqOffsetR], uint32(sequence))

	byteBody := []byte(body)
	copy(buf[headerLen:], byteBody)

	return
}

func main() {
	data := encoder("Hello, 世界")
	decoder(data)
}
