package flic

import (
	"encoding/binary"
)

func (c *Client) writeCommand(opcode uint8, connID int32, bdAddr string) error {
	buf := make([]byte, 100)
	pos := 0

	// Write opcode
	buf[pos] = opcode
	pos++

	// Write connection ID
	binary.LittleEndian.PutUint32(buf[pos:], uint32(connID))
	pos += 4

	// Write bdAddr (Bluetooth address)
	for i := 15; i >= 0; i -= 3 {
		b := hexToInt(bdAddr[i:i+2], 16)
		buf[pos] = byte(b)
		pos++
	}

	// Write latency mode (Normal = 0)
	buf[pos] = 0
	pos++

	// Write auto disconnect time (511)
	binary.LittleEndian.PutUint16(buf[pos:], 511)
	pos += 2

	return c.conn.WriteMessage(2, buf[:pos])
}

func hexToInt(s string, base int) int64 {
	var val int64
	for i := 0; i < len(s); i++ {
		val *= int64(base)
		c := s[i]
		switch {
		case c >= '0' && c <= '9':
			val += int64(c - '0')
		case c >= 'a' && c <= 'f':
			val += int64(c - 'a' + 10)
		case c >= 'A' && c <= 'F':
			val += int64(c - 'A' + 10)
		}
	}
	return val
}
